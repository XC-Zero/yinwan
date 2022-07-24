package mongo_model

import (
	"context"
	"fmt"
	"github.com/XC-Zero/yinwan/pkg/client"
	_const "github.com/XC-Zero/yinwan/pkg/const"
	_interface "github.com/XC-Zero/yinwan/pkg/interface"
	"github.com/XC-Zero/yinwan/pkg/model/mysql_model"
	"github.com/XC-Zero/yinwan/pkg/utils/convert"
	"github.com/XC-Zero/yinwan/pkg/utils/math_plus"
	myMongo "github.com/XC-Zero/yinwan/pkg/utils/mongo"
	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/bson"
	"gorm.io/gorm"
	"time"
)

// StockOutRecord 出库记录
type StockOutRecord struct {
	BasicModel              `bson:"inline"`
	BookNameInfo            `bson:"-"`
	StockOutRecordOwnerID   *int                 `json:"stock_out_record_owner_id,omitempty" form:"stock_out_record_owner_id,omitempty" bson:"stock_out_record_owner_id,omitempty"`
	StockOutRecordOwnerName *string              `json:"stock_out_record_owner_name,omitempty" form:"stock_out_record_owner_name,omitempty" bson:"stock_out_record_owner_name,omitempty"`
	StockOutRecordType      string               `json:"stock_out_record_type" form:"stock_out_record_type" bson:"stock_out_record_type"`
	StockOutWarehouseID     *int                 `json:"stock_out_warehouse_id,omitempty" form:"stock_out_warehouse_id,omitempty" bson:"stock_out_warehouse_id"`
	StockOutWarehouseName   *string              `json:"stock_out_warehouse_name,omitempty" form:"stock_out_warehouse_name,omitempty" bson:"stock_out_warehouse_name"`
	StockOutDetailPosition  *string              `json:"stock_out_detail_position,omitempty" form:"stock_out_detail_position" bson:"stock_out_detail_position"`
	StockOutRecordContent   []stockRecordContent `json:"stock_out_record_content" form:"stock_out_record_content" bson:"stock_out_record_content"`
	RelateInvoice           []relatedInvoice     `json:"relate_invoice" form:"relate_invoice" bson:"relate_invoice"`
	Remark                  *string              `json:"remark,omitempty" form:"remark" bson:"remark"`
}

func (m StockOutRecord) ToESDoc() map[string]interface{} {
	return map[string]interface{}{
		"rec_id":                m.RecID,
		"created_at":            m.CreatedAt,
		"remark":                m.Remark,
		"stock_out_content":     convert.StructToTagString(m.StockOutRecordContent, string(_const.CN)),
		"stock_out_record_type": m.StockOutRecordType,
		"stock_out_owner":       m.StockOutRecordOwnerName,
		"book_name":             m.BookName,
		"book_name_id":          m.BookNameID,
	}
}

func (m StockOutRecord) TableName() string {
	return "stock_out_records"
}

func (m StockOutRecord) TableCnName() string {
	return "出库记录"
}

func (m StockOutRecord) Mapping() map[string]interface{} {
	ma := mapping{
		"settings": mapping{},
		"mappings": mapping{
			"properties": mapping{
				"rec_id": mapping{
					"type": "keyword",
				},
				"stock_out_content": mapping{
					"type":            "text",
					"analyzer":        IK_SMART,
					"search_analyzer": IK_SMART,
				},
				"stock_out_owner": mapping{
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
				"stock_out_record_type": mapping{
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
	return ma
}

// BeforeInsert 出库 同步产品或者其他
// 	TODO 前端或后端实现 同一原材料,同一批次只可选择一次!!!
func (m *StockOutRecord) BeforeInsert(ctx context.Context) error {
	bookName := ctx.Value("book_name").(string)
	bk, ok := client.ReadBookMap(bookName)
	if !ok {
		return errors.New("There is no book name!")
	}
	if m.RecID == nil || bk.StorageName == "" {
		return errors.New("缺少主键！")
	}
	err := bk.MysqlClient.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		var contentList = m.StockOutRecordContent
		for i := 0; i < len(contentList); i++ {
			var model _interface.ChineseTabler
			var surplusColumnName string
			if contentList[i].ContentType == _const.MATERIAL {
				surplusColumnName = "material_batch_surplus_number"
				model = mysql_model.MaterialBatch{
					BasicModel: mysql_model.BasicModel{
						RecID: &contentList[i].RelatedBatchID,
					},
				}
			} else if contentList[i].ContentType == _const.COMMODITY {
				surplusColumnName = "commodity_batch_surplus_number"
				model = mysql_model.CommodityBatch{
					BasicModel: mysql_model.BasicModel{
						RecID: &contentList[i].RelatedBatchID,
					},
				}
			} else {
				continue
			}
			query := tx.Model(&model).Select(surplusColumnName).Where("rec_id = ?", contentList[i].RelatedBatchID)
			err := tx.WithContext(ctx).UpdateColumn(surplusColumnName, gorm.Expr(" ? - ?", query, contentList[i].Num)).
				Where("rec_id = ?", contentList[i].RelatedBatchID).Error
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

// BeforeUpdate
//	 todo 前端禁止! 修改时不允许新增!!!! 不允许修改批次号!!! 除了数量啥也不能改!!!! 可以删!!!
func (m *StockOutRecord) BeforeUpdate(ctx context.Context) error {
	var redStockRecord = *m
	redStockRecord.StockOutRecordContent = []stockRecordContent{}
	bookName := ctx.Value("book_name").(string)
	bk, ok := client.ReadBookMap(bookName)
	if !ok {
		return errors.New("There is no book name!")
	}
	if m.RecID == nil || bk.StorageName == "" {
		return errors.New("缺少主键！")
	}
	c, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()
	var filter = bson.D{}
	filter = append(filter, myMongo.TransMysqlOperatorSymbol(myMongo.EQUAL, "rec_id", m.RecID))
	var beforeStockOutRecord StockOutRecord
	err := bk.MongoDBClient.Collection(m.TableName()).Find(c, filter).Limit(1).All(&beforeStockOutRecord)
	if err != nil {
		return err
	}
	for _, before := range beforeStockOutRecord.StockOutRecordContent {
		for _, now := range m.StockOutRecordContent {
			// 如果是同一批次就更新
			if before.RelatedBatchID == now.RelatedBatchID {
				// 改了数量
				if before.Num != now.Num {
					tmp := now
					tmp.Num = now.Num - before.Num
					fraction, err := math_plus.NewFromString(tmp.Price)
					if err != nil {
						return err
					}
					tmp.TotalPrice = fraction.MulInt64(int64(tmp.Num)).String()
					redStockRecord.StockOutRecordContent = append(redStockRecord.StockOutRecordContent, tmp)
				}
				continue
			}
		}
		//	 找不到这个批次号也就是删了
		tmp := before
		tmp.TotalPrice = "-" + tmp.TotalPrice
		tmp.Num = -tmp.Num
		redStockRecord.StockOutRecordContent = append(redStockRecord.StockOutRecordContent, tmp)
	}
	remark := fmt.Sprintf("此为系统根据 [修改] 编号为: [%d] 所创建的(红字)出库单", *m.RecID)
	id, name := 0, "系统"
	redStockRecord.Remark = &remark
	redStockRecord.StockOutRecordOwnerID = &id
	redStockRecord.StockOutRecordOwnerName = &name
	_, err = bk.MongoDBClient.Collection(m.TableName()).InsertOne(ctx, redStockRecord)
	if err != nil {
		return err
	}
	// 撤销修改,换为新增红字
	return myMongo.CancelError
}

// BeforeRemove 撤销/删除 出库嘛 换为新增红字
// 	FIXME 万一他删除我的删除的记录呢??? 套娃???
func (m *StockOutRecord) BeforeRemove(ctx context.Context) error {
	bookName := ctx.Value("book_name").(string)
	bk, ok := client.ReadBookMap(bookName)
	if !ok {
		return errors.New("There is no book name!")
	}
	if m.RecID == nil || bk.StorageName == "" {
		return errors.New("缺少主键！")
	}

	for i := 0; i < len(m.StockOutRecordContent); i++ {
		m.StockOutRecordContent[i].Num = -m.StockOutRecordContent[i].Num
		m.StockOutRecordContent[i].TotalPrice = "-" + m.StockOutRecordContent[i].TotalPrice
	}

	remark := fmt.Sprintf("此为系统根据 [删除] 编号为: [%d] 所创建的(红字)出库单", *m.RecID)
	id, name := 0, "系统"
	m.Remark = &remark
	m.StockOutRecordOwnerID = &id
	m.StockOutRecordOwnerName = &name
	_, err := bk.MongoDBClient.Collection(m.TableName()).InsertOne(ctx, m)
	if err != nil {
		return err
	}
	// 撤销删除,换为新增红字 TODO 真的会取消吗??
	return myMongo.CancelError
}
