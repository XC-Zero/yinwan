package mongo_model

import (
	"github.com/XC-Zero/yinwan/pkg/client"
	_const "github.com/XC-Zero/yinwan/pkg/const"
	"github.com/XC-Zero/yinwan/pkg/model/mysql_model"
	"github.com/XC-Zero/yinwan/pkg/utils/convert"
	"github.com/XC-Zero/yinwan/pkg/utils/logger"
	"github.com/pkg/errors"
	"golang.org/x/net/context"
	"gorm.io/gorm"
	"time"
)

// StockInRecord 入库记录
// 存 MongoDB
type StockInRecord struct {
	BasicModel                `bson:"inline"`
	BookNameInfo              `bson:"-"`
	StockInRecordOwnerID      *int                 ` json:"stock_in_record_owner_id,omitempty"  form:"stock_in_record_owner_id" bson:"stock_in_record_owner_id" `
	StockInRecordOwnerName    *string              ` json:"stock_in_record_owner_name,omitempty" form:"stock_in_record_owner_name" bson:"stock_in_record_owner_name" `
	StockInRecordProviderID   *int                 ` json:"stock_in_record_provider_id,omitempty" bson:"stock_in_record_provider_id"`
	StockInRecordProviderName *string              ` json:"stock_in_record_provider_name,omitempty" bson:"stock_in_record_provider_name"`
	StockInWarehouseID        *int                 ` json:"stock_in_warehouse_id,omitempty" form:"stock_in_warehouse_id,omitempty" bson:"stock_in_warehouse_id"`
	StockInWarehouseName      *string              ` json:"stock_in_warehouse_name,omitempty" form:"stock_in_warehouse_name,omitempty" bson:"stock_in_warehouse_name"`
	StockInDetailPosition     *string              ` json:"stock_in_detail_position,omitempty" form:"stock_in_detail_position" bson:"stock_in_detail_position"`
	StockInRecordType         string               ` json:"stock_in_record_type" form:"stock_in_record_type" bson:"stock_in_record_type" `
	StockInRecordContent      []stockRecordContent ` json:"stock_in_record_content" form:"stock_in_record_content" bson:"stock_in_record_content" `
	RelatePurchaseID          []*int               ` json:"relate_purchase_id,omitempty" form:"relate_purchase_id,omitempty" bson:"relate_purchase_id"`
	Remark                    *string              ` json:"remark,omitempty" form:"remark" bson:"remark"`
}
type stockInRecordContent struct {
	MaterialID          int    `bson:"material_id" json:"material_id" form:"material_id"`
	MaterialName        string `bson:"material_name" json:"material_name" form:"material_name"`
	MaterialNum         int    `bson:"material_num" json:"material_num" form:"material_num"`
	MaterialAmount      string `bson:"material_amount" json:"material_amount" form:"material_amount"`
	MaterialTotalAmount string `bson:"material_total_amount" json:"material_total_amount" form:"material_total_amount"`
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
					"type":            "text",   //字符串类型且进行分词, 允许模糊匹配
					"analyzer":        IK_SMART, //设置分词工具
					"search_analyzer": IK_SMART,
				},
				"stock_in_owner": mapping{
					"type":            "text",   //字符串类型且进行分词, 允许模糊匹配
					"analyzer":        IK_SMART, //设置分词工具
					"search_analyzer": IK_SMART,
					"fields": mapping{ //当需要对模糊匹配的字符串也允许进行精确匹配时假如此配置
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
		date := time.Now().Format("2006-01-02 15:04")
		for _, content := range s.StockInRecordContent {
			batch := mysql_model.MaterialBatch{
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
				StockInTime:                &date,
				Remark:                     s.Remark,
			}
			err := tx.WithContext(context.WithValue(context.Background(), "material_id", content.RecID)).
				Create(&batch).Error
			if err != nil {
				logger.Error(errors.WithStack(err), "同步创建批次信息失败!")
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
