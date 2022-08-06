package mongo_model

import (
	"context"
	"github.com/XC-Zero/yinwan/pkg/client"
	_const "github.com/XC-Zero/yinwan/pkg/const"
	"github.com/XC-Zero/yinwan/pkg/model/mysql_model"
	"github.com/XC-Zero/yinwan/pkg/utils/convert"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

type PayType int

//goland:noinspection GoSnakeCaseUsage
const (
	WX PayType = iota + 35001
	ZFB
	BANK_CARD
	CASH
	OTHER
)

func (p PayType) DisPlay() string {
	switch p {
	case ZFB:
		return "支付宝"
	case CASH:
		return "现金"
	case WX:
		return "微信"
	case BANK_CARD:
		return "银行卡"
	case OTHER:
		return "其他"
	default:
		return "未知"
	}
}

// Transaction 销售
type Transaction struct {
	BasicModel              `bson:"inline"`
	BookNameInfo            `bson:"-"`
	TransactionName         *string              `json:"transaction_name,omitempty" form:"transaction_name,omitempty" bson:"transaction_name,omitempty"`
	TransactionContent      []transactionContent `json:"transaction_content" form:"transaction_content" bson:"transaction_content" cn:"销售详情"`
	TransactionAmount       string               `json:"transaction_amount" form:"transaction_amount" bson:"transaction_amount" cn:"销售金额"`
	TransactionActualAmount string               `json:"transaction_actual_amount" form:"transaction_actual_amount" bson:"transaction_actual_amount" cn:"实际收款金额"`
	CustomerID              *int                 `json:"customer_id" form:"customer_id" bson:"customer_id" cn:"客户编号"`
	CustomerName            *string              `json:"customer_name,omitempty" form:"customer_name,omitempty" json:"customer_name" cn:"客户姓名"`
	TransactionOwnerID      *int                 `json:"transaction_owner_id,omitempty" form:"transaction_owner_id" bson:"transaction_owner_id" cn:"销售管理员ID"`
	TransactionOwnerName    *string              `json:"transaction_owner_name,omitempty" form:"transaction_owner_name" bson:"transaction_owner_name" cn:"销售管理员姓名"`
	TransactionTime         *string              `json:"transaction_time,omitempty" form:"transaction_time" bson:"transaction_time" cn:"交易时间"`
	ReceiveID               *int                 `json:"receive_id,omitempty" form:"receive_id" bson:"receive_id" cn:"关联应收记录"`
	Remark                  *string              `json:"remark,omitempty" form:"remark" bson:"remark" cn:"备注"`
}

type transactionContent struct {
	stockRecordContent
	TransactionPrice      string `json:"transaction_price" form:"transaction_price" bson:"transaction_price" cn:"售价"`
	TransactionTotalPrice string `json:"transaction_total_price" form:"transaction_total_price" bson:"transaction_total_price" cn:"售价合计"`
}

func (t Transaction) Mapping() map[string]interface{} {
	return mapping{
		"settings": mapping{},
		"mappings": mapping{
			"properties": mapping{
				"rec_id": mapping{
					"type": "keyword",
				},
				"transaction_name": mapping{
					"type":            "text",
					"analyzer":        IK_SMART,
					"search_analyzer": IK_SMART,
				},
				"transaction_content": mapping{
					"type":            "text",
					"analyzer":        IK_SMART,
					"search_analyzer": IK_SMART,
				},
				"transaction_owner_name": mapping{
					"type":            "text",
					"analyzer":        IK_SMART,
					"search_analyzer": IK_SMART,
					"fields": mapping{
						"keyword": mapping{
							"type":         "keyword",
							"ignore_above": 256,
						},
					},
				},
				"transaction_amount": mapping{
					"type": "keyword",
				},
				"receive_id": mapping{
					"type": "keyword",
				},
				"remark": mapping{
					"type":            "text",
					"analyzer":        IK_SMART,
					"search_analyzer": IK_SMART,
				},
				"created_at": mapping{
					"type": "text",
				},
				"book_name": mapping{
					"type": "keyword",
				},
				"book_name_id": mapping{
					"type": "keyword",
				},
			},
		},
	}
}

// ToESDoc todo!!!2
func (t Transaction) ToESDoc() map[string]interface{} {
	return map[string]interface{}{
		"rec_id":                 t.RecID,
		"remark":                 t.Remark,
		"created_at":             t.CreatedAt,
		"receive_id":             t.ReceiveID,
		"transaction_amount":     t.TransactionAmount,
		"transaction_owner_name": t.TransactionOwnerName,
		"transaction_content":    convert.StructSliceToTagString(t.TransactionContent, string(_const.CN)),
		"book_name":              t.BookName,
		"book_name_id":           t.BookNameID,
		"transaction_name":       t.TransactionName,
	}
}

func (t Transaction) TableCnName() string {
	return "销售"
}

func (t Transaction) TableName() string {
	return "transactions"
}

// BeforeInsert 创建销售记录
// 如果自动创建应收记录则 mongo 事务套MySQL事务
// 逻辑上保证原子性
// 之所以选择前触发器,是想把mysql的
func (t *Transaction) BeforeInsert(ctx context.Context) error {
	bookName := ctx.Value("book_name").(string)
	bk, ok := client.ReadBookMap(bookName)
	if !ok {
		return errors.New("There is no book name!")
	}
	if t.RecID == nil || bk.StorageName == "" {
		return errors.New("缺少主键！")
	}
	t.BookNameID = bk.StorageName
	t.BookName = bk.BookName
	// 同步创建应收
	flag, ok := ctx.Value("auto").(bool)
	if ok && flag {
		unpaid := _const.Unfinished
		var receivable = mysql_model.Receivable{
			BookNameInfo: mysql_model.BookNameInfo{
				BookNameID: bk.StorageName,
				BookName:   bk.BookName,
			},
			CustomerID:            t.CustomerID,
			CustomerName:          t.CustomerName,
			ReceivableTotalAmount: &t.TransactionActualAmount,
			ReceivableStatus:      &unpaid,
			TransactionID:         t.RecID,
			Remark:                t.Remark,
		}

		err := bk.MysqlClient.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
			err := tx.Create(&receivable).Error
			if err != nil {
				return err
			}
			// 事务里的是并发还是一个一个来?失败了回滚?
			// FIXME 测试下能不能真的插入
			t.ReceiveID = receivable.RecID

			return nil
		})
		if err != nil {
			return err
		}
	}
	return nil

}

// BeforeUpdate TODO !
func (t *Transaction) BeforeUpdate(ctx context.Context) error {
	flag, ok := ctx.Value("auto").(bool)
	if ok && flag {
		bookName := ctx.Value("book_name").(string)
		bk, ok := client.ReadBookMap(bookName)
		if !ok {
			return errors.New("There is no book name!")
		}
		if t.RecID == nil || bk.StorageName == "" {
			return errors.New("缺少主键！")
		}
		t.BookNameID = bk.StorageName
		t.BookName = bk.BookName
		var receivable = mysql_model.Receivable{
			BookNameInfo: mysql_model.BookNameInfo{
				BookNameID: bk.StorageName,
				BookName:   bk.BookName,
			},
			CustomerID:            t.CustomerID,
			CustomerName:          t.CustomerName,
			ReceivableTotalAmount: &t.TransactionActualAmount,
			TransactionID:         t.RecID,
			Remark:                t.Remark,
		}
		err := bk.MysqlClient.WithContext(ctx).Transaction(func(tx *gorm.DB) error {

			var receivables []mysql_model.Receivable
			err := tx.Model(receivable).Where("transaction_id = ?", t.RecID).Find(&receivables).Error
			if err != nil {
				return err
			}
			if len(receivables) != 0 {
				err = tx.Updates(&receivable).Where("transaction_id = ?", t.RecID).Error
				if err != nil {
					return err
				}
			}
			t.ReceiveID = receivable.RecID
			return nil
		})
		if err != nil {
			return err
		}
	}
	return nil
}

// BeforeDelete TODO !
func (t *Transaction) BeforeDelete(ctx context.Context) error {
	bookName := ctx.Value("book_name").(string)
	bk, ok := client.ReadBookMap(bookName)
	if !ok {
		return errors.New("There is no book name!")
	}
	if t.RecID == nil || bk.StorageName == "" {
		return errors.New("缺少主键！")
	}
	t.BookNameID = bk.StorageName
	t.BookName = bk.BookName

	err := bk.MysqlClient.WithContext(ctx).Transaction(func(tx *gorm.DB) error {

		var receivables []mysql_model.Receivable
		err := tx.Model(mysql_model.Receivable{}).Where("transaction_id = ?", t.RecID).Find(&receivables).Error
		if err != nil {
			return err
		}
		if len(receivables) == 0 {
			return nil
		}
		flag, ok := ctx.Value("auto").(bool)
		if ok && flag {
			err = bk.MysqlClient.WithContext(ctx).Delete(mysql_model.Receivable{}).Where("transaction_id = ?", t.RecID).Error
			if err != nil {
				return err
			}
		} else {
			var receivable = mysql_model.Receivable{
				BookNameInfo: mysql_model.BookNameInfo{
					BookNameID: bk.StorageName,
					BookName:   bk.BookName,
				},
				CustomerID:            t.CustomerID,
				CustomerName:          t.CustomerName,
				ReceivableTotalAmount: &t.TransactionActualAmount,
				ReceivableDebtAmount:  &t.TransactionActualAmount,
				TransactionID:         t.RecID,
				Remark:                t.Remark,
			}
			err = bk.MysqlClient.WithContext(ctx).Where("transaction_id = ?", t.RecID).
				Updates(&receivable).Error
			if err != nil {
				return err
			}
		}
		return nil
	})
	if err != nil {
		return err
	}
	return nil
}
