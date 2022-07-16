package mongo_model

import (
	"context"
	"github.com/XC-Zero/yinwan/pkg/client"
	_const "github.com/XC-Zero/yinwan/pkg/const"
	"github.com/XC-Zero/yinwan/pkg/utils/convert"
	"github.com/pkg/errors"
)

// StockOutRecord 出库记录
// 存 MongoDB 一份
type StockOutRecord struct {
	BasicModel              `bson:"inline"`
	BookNameInfo            `bson:"-"`
	StockOutRecordOwnerID   int                     `json:"stock_out_record_owner_id" form:"stock_out_record_owner_id" bson:"stock_out_record_owner_id"`
	StockOutRecordOwnerName string                  `json:"stock_out_record_owner_name" form:"stock_out_record_owner_name" bson:"stock_out_record_owner_name"`
	StockOutRecordType      string                  `json:"stock_out_record_type" form:"stock_out_record_type" bson:"stock_out_record_type"`
	StockOutWarehouseID     *int                    `json:"stock_out_warehouse_id,omitempty" form:"stock_out_warehouse_id,omitempty" bson:"stock_out_warehouse_id"`
	StockOutWarehouseName   *string                 `json:"stock_out_warehouse_name,omitempty" form:"stock_out_warehouse_name,omitempty" bson:"stock_out_warehouse_name"`
	StockOutDetailPosition  *string                 `json:"stock_out_detail_position,omitempty" form:"stock_out_detail_position" bson:"stock_out_detail_position"`
	StockOutRecordContent   []stockOutRecordContent `json:"stock_out_record_content" form:"stock_out_record_content" bson:"stock_out_record_content"`
	Remark                  *string                 `json:"remark,omitempty" form:"remark" bson:"remark"`
}
type stockOutRecordContent struct {
	CommodityID      *int    `form:"commodity_id,omitempty" json:"commodity_id,omitempty" bson:"commodity_id"`
	CommodityName    *string `form:"commodity_name,omitempty" json:"commodity_name,omitempty" bson:"commodity_name"`
	CommodityNum     int     `form:"commodity_num" json:"commodity_num" bson:"commodity_num"`
	CommodityPrice   *string `form:"commodity_price,omitempty" json:"commodity_price,omitempty" bson:"commodity_price,omitempty"`
	CommodityBatchID *int    `form:"commodity_batch_id,omitempty" json:"commodity_batch_id,omitempty" bson:"commodity_batch_id,omitempty"`
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
					"type":            "text",   //字符串类型且进行分词, 允许模糊匹配
					"analyzer":        IK_SMART, //设置分词工具
					"search_analyzer": IK_SMART,
				},
				"stock_out_record_type": mapping{
					"type": "keyword",
				},
				"stock_out_owner": mapping{
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

// BeforeInsert todo !!!!
func (m *StockOutRecord) BeforeInsert(ctx context.Context) error {
	bookName := ctx.Value("book_name").(string)
	bk, ok := client.ReadBookMap(bookName)
	if !ok {
		return errors.New("There is no book name!")
	}
	if m.RecID == nil || bk.StorageName == "" {
		return errors.New("缺少主键！")
	}
	return nil
}

// BeforeUpdate todo !!!!
func (m *StockOutRecord) BeforeUpdate(ctx context.Context) error {

	return nil
}

// BeforeRemove todo !!!!
func (m *StockOutRecord) BeforeRemove(ctx context.Context) error {

	return nil
}
