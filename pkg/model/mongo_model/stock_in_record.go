package mongo_model

import (
	"fmt"
	"github.com/XC-Zero/yinwan/pkg/client"
	_const "github.com/XC-Zero/yinwan/pkg/const"
	_interface "github.com/XC-Zero/yinwan/pkg/interface"
	"github.com/XC-Zero/yinwan/pkg/model/mysql_model"
	"github.com/XC-Zero/yinwan/pkg/utils/convert"
	myMongo "github.com/XC-Zero/yinwan/pkg/utils/mongo"
	"github.com/pkg/errors"
	"golang.org/x/net/context"
	"gorm.io/gorm"
	"time"
)

// StockInRecord 入库记录
type StockInRecord struct {
	BasicModel             `bson:"inline"`
	BookNameInfo           `bson:"-"`
	StockInRecordOwnerID   *int                 `json:"stock_in_record_owner_id,omitempty"  form:"stock_in_record_owner_id" bson:"stock_in_record_owner_id" `
	StockInRecordOwnerName *string              `json:"stock_in_record_owner_name,omitempty" form:"stock_in_record_owner_name" bson:"stock_in_record_owner_name" `
	StockInWarehouseID     *int                 `json:"stock_in_warehouse_id,omitempty" form:"stock_in_warehouse_id,omitempty" bson:"stock_in_warehouse_id"`
	StockInWarehouseName   *string              `json:"stock_in_warehouse_name,omitempty" form:"stock_in_warehouse_name,omitempty" bson:"stock_in_warehouse_name"`
	StockInDetailPosition  *string              `json:"stock_in_detail_position,omitempty" form:"stock_in_detail_position" bson:"stock_in_detail_position"`
	StockInRecordType      string               `json:"stock_in_record_type" form:"stock_in_record_type" bson:"stock_in_record_type" `
	StockInRecordContent   []stockRecordContent `json:"stock_in_record_content" form:"stock_in_record_content" bson:"stock_in_record_content" `
	RelateInvoice          []relatedInvoice     `json:"relate_invoice" form:"relate_invoice" bson:"relate_invoice"`
	Remark                 *string              `json:"remark,omitempty" form:"remark" bson:"remark"`
}

func (s StockInRecord) ToESDoc() map[string]interface{} {
	return map[string]interface{}{
		"rec_id":               s.RecID,
		"created_at":           s.CreatedAt,
		"remark":               s.Remark,
		"stock_in_content":     convert.StructSliceToTagString(s.StockInRecordContent, string(_const.CN)),
		"stock_in_record_type": s.StockInRecordType,
		"stock_in_owner":       s.StockInRecordOwnerName,
		"book_name":            s.BookName,
		"book_name_id":         s.BookNameID,
	}
}

func (s StockInRecord) TableName() string {
	return "stock_in_records"
}
func (s StockInRecord) TableCnName() string {
	return "入库记录"
}
func (s StockInRecord) Mapping() map[string]interface{} {
	ma := mapping{
		"settings": mapping{},
		"mappings": mapping{
			"properties": mapping{
				"rec_id": mapping{
					"type": "keyword",
				},
				"stock_in_content": mapping{
					"type":            "text",
					"analyzer":        IK_SMART,
					"search_analyzer": IK_SMART,
				},
				"stock_in_owner": mapping{
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
				"stock_in_record_type": mapping{
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

// BeforeInsert 创建入库
//
func (s *StockInRecord) BeforeInsert(ctx context.Context) error {
	bookName := ctx.Value("book_name").(string)
	bk, ok := client.ReadBookMap(bookName)
	if !ok {
		return errors.New("There is no book name!")
	}
	if s.RecID == nil || bk.StorageName == "" {
		return errors.New("缺少主键！")
	}
	err := bk.MysqlClient.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		var contentList = s.StockInRecordContent
		var now = time.Now().Format("2006-01-02")
		var remark = fmt.Sprintf("根据入库单[编号]:%d 自动创建", *s.RecID)
		for i := 0; i < len(contentList); i++ {
			content := contentList[i]
			var model _interface.ChineseTabler
			var relateID string
			if content.ContentType == _const.MATERIAL {
				relateID = "material_id"
				model = mysql_model.MaterialBatch{
					BasicModel: mysql_model.BasicModel{
						CreatedAt: time.Now(),
					},
					MaterialID:                 content.RecID,
					MaterialName:               content.Name,
					StockInRecordID:            *s.RecID,
					MaterialBatchOwnerID:       s.StockInRecordOwnerID,
					MaterialBatchOwnerName:     s.StockInRecordOwnerName,
					MaterialBatchTotalPrice:    content.TotalPrice,
					MaterialBatchNumber:        content.Num,
					MaterialBatchSurplusNumber: content.Num,
					MaterialBatchUnitPrice:     content.Price,
					WarehouseID:                s.StockInWarehouseID,
					WarehouseName:              s.StockInWarehouseName,
					StockInTime:                &now,
					Remark:                     &remark,
				}
			} else if content.ContentType == _const.COMMODITY {
				relateID = "commodity_id"
				model = mysql_model.CommodityBatch{
					BasicModel: mysql_model.BasicModel{
						CreatedAt: time.Now(),
					},
					CommodityID:                 content.RecID,
					CommodityName:               content.Name,
					StockInRecordID:             *s.RecID,
					CommodityBatchOwnerID:       s.StockInRecordOwnerID,
					CommodityBatchOwnerName:     s.StockInRecordOwnerName,
					CommodityBatchTotalPrice:    content.TotalPrice,
					CommodityBatchNumber:        content.Num,
					CommodityBatchSurplusNumber: content.Num,
					CommodityBatchUnitPrice:     content.Price,
					WarehouseID:                 s.StockInWarehouseID,
					WarehouseName:               s.StockInWarehouseName,
					StockInTime:                 &now,
					Remark:                      &remark,
				}
			} else {
				continue
			}
			err := tx.WithContext(context.WithValue(context.Background(), relateID, content.RecID)).
				Create(&model).Error

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

func (s *StockInRecord) BeforeUpdate(ctx context.Context) error {
	return myMongo.CancelError
}

func (s *StockInRecord) BeforeRemove(ctx context.Context) error {
	return myMongo.CancelError
}
