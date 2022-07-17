package mongo_model

import (
	"context"
	"github.com/XC-Zero/yinwan/pkg/client"
	_const "github.com/XC-Zero/yinwan/pkg/const"
	_interface "github.com/XC-Zero/yinwan/pkg/interface"
	"github.com/XC-Zero/yinwan/pkg/model/mysql_model"
	"github.com/XC-Zero/yinwan/pkg/utils/convert"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

// StockOutRecord 出库记录
// 存 MongoDB 一份
type StockOutRecord struct {
	BasicModel              `bson:"inline"`
	BookNameInfo            `bson:"-"`
	StockOutRecordOwnerID   int                  `json:"stock_out_record_owner_id" form:"stock_out_record_owner_id" bson:"stock_out_record_owner_id"`
	StockOutRecordOwnerName string               `json:"stock_out_record_owner_name" form:"stock_out_record_owner_name" bson:"stock_out_record_owner_name"`
	StockOutRecordType      string               `json:"stock_out_record_type" form:"stock_out_record_type" bson:"stock_out_record_type"`
	StockOutWarehouseID     *int                 `json:"stock_out_warehouse_id,omitempty" form:"stock_out_warehouse_id,omitempty" bson:"stock_out_warehouse_id"`
	StockOutWarehouseName   *string              `json:"stock_out_warehouse_name,omitempty" form:"stock_out_warehouse_name,omitempty" bson:"stock_out_warehouse_name"`
	StockOutDetailPosition  *string              `json:"stock_out_detail_position,omitempty" form:"stock_out_detail_position" bson:"stock_out_detail_position"`
	StockOutRecordContent   []stockRecordContent `json:"stock_out_record_content" form:"stock_out_record_content" bson:"stock_out_record_content"`
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
				"stock_out_record_type": mapping{
					"type": "keyword",
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
		//date := time.Now().Format("2006-01-02 15:04")

		for i := 0; i < len(contentList); i++ {
			var model _interface.ChineseTabler
			var surplusColumnName string
			if contentList[i].ContentType == _const.MATERIAL {
				surplusColumnName = "material_batch_surplus_number"
				model = mysql_model.MaterialBatch{
					BasicModel: mysql_model.BasicModel{
						RecID: contentList[i].RelatedBatchID,
					},
				}
			} else if contentList[i].ContentType == _const.COMMODITY {
				surplusColumnName = "commodity_batch_surplus_number"
				model = mysql_model.CommodityBatch{
					BasicModel: mysql_model.BasicModel{
						RecID: contentList[i].RelatedBatchID,
					},
				}
			} else {
				continue
			}
			query := tx.Model(&model).Select(surplusColumnName).Where("rec_id = ?", contentList[i].RelatedBatchID)
			err := tx.WithContext(ctx).UpdateColumn("material_batch_number", gorm.Expr(" ? + ?", query, contentList[i].Num)).
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

// BeforeUpdate todo !!!! 前后对比!!!
func (m *StockOutRecord) BeforeUpdate(ctx context.Context) error {
	//return m.BeforeInsert(ctx)
	return nil
}

// BeforeRemove 撤销出库嘛
func (m *StockOutRecord) BeforeRemove(ctx context.Context) error {
	for i := 0; i < len(m.StockOutRecordContent); i++ {
		m.StockOutRecordContent[i].Num = -m.StockOutRecordContent[i].Num
	}
	return m.BeforeInsert(ctx)
}
