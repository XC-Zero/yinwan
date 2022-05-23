package mongo_model

import (
	"encoding/json"
	"github.com/XC-Zero/yinwan/pkg/model/mysql_model"
)

// StockInRecord 入库记录
// 存 MongoDB
type StockInRecord struct {
	mysql_model.BasicModel
	BookNameInfo
	StockInRecordOwnerID      *int                     ` json:"stock_in_record_owner_id,omitempty"  form:"stock_in_record_owner_id" bson:"stock_in_record_owner_id" `
	StockInRecordOwnerName    *string                  ` json:"stock_in_record_owner_name,omitempty" form:"stock_in_record_owner_name" bson:"stock_in_record_owner_name" `
	StockInRecordProviderID   *int                     ` json:"stock_in_record_provider_id,omitempty" bson:"stock_in_record_provider_id"`
	StockInRecordProviderName *string                  ` json:"stock_in_record_provider_name,omitempty" bson:"stock_in_record_provider_name"`
	StockInWarehouseID        *int                     ` json:"stock_in_warehouse_id,omitempty" form:"stock_in_warehouse_id,omitempty" bson:"stock_in_warehouse_id"`
	StockInWarehouseName      *string                  ` json:"stock_in_warehouse_name,omitempty" form:"stock_in_warehouse_name,omitempty" bson:"stock_in_warehouse_name"`
	StockInDetailPosition     *string                  ` json:"stock_in_detail_position,omitempty" form:"stock_in_detail_position" bson:"stock_in_detail_position"`
	StockInRecordType         string                   ` json:"stock_in_record_type" form:"stock_in_record_type" bson:"stock_in_record_type" `
	StockInRecordContent      []map[string]interface{} ` json:"stock_in_record_content" form:"stock_in_record_content" bson:"stock_in_record_content" `
	RelatePurchaseID          *int                     ` json:"relate_purchase_id,omitempty" form:"relate_purchase_id,omitempty" bson:"relate_purchase_id"`
	Remark                    *string                  ` json:"remark,omitempty" form:"remark" bson:"remark"`
}

func (m StockInRecord) TableName() string {
	return "stock_in_records"
}
func (m StockInRecord) TableCnName() string {
	return "入库记录"
}
func (m StockInRecord) Mapping() map[string]interface{} {
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
func (m StockInRecord) ToESDoc() map[string]interface{} {
	var str string
	bytes, err := json.Marshal(m.StockInRecordContent)
	str = string(bytes)
	if err != nil {
		str = ""
	}

	return map[string]interface{}{
		"rec_id":               m.RecID,
		"created_at":           m.CreatedAt,
		"remark":               m.Remark,
		"stock_in_content":     str,
		"stock_in_record_type": m.StockInRecordType,
		"stock_in_owner":       m.StockInRecordOwnerName,
		"book_name":            m.BookName,
		"book_name_id":         m.BookNameID,
	}
}

// StockOutRecord 出库记录
// 存 MongoDB 一份
type StockOutRecord struct {
	mysql_model.BasicModel
	BookNameInfo
	StockOutRecordOwnerID   int                      `json:"stock_out_record_owner_id" form:"stock_out_record_owner_id" bson:"stock_out_record_owner_id"`
	StockOutRecordOwnerName string                   `json:"stock_out_record_owner_name" form:"stock_out_record_owner_name" bson:"stock_out_record_owner_name"`
	StockOutRecordType      string                   `json:"stock_out_record_type" form:"stock_out_record_type" bson:"stock_out_record_type"`
	StockOutWarehouseID     *int                     ` json:"stock_out_warehouse_id,omitempty" form:"stock_out_warehouse_id,omitempty" bson:"stock_out_warehouse_id"`
	StockOutWarehouseName   *string                  ` json:"stock_out_warehouse_name,omitempty" form:"stock_out_warehouse_name,omitempty" bson:"stock_out_warehouse_name"`
	StockOutDetailPosition  *string                  ` json:"stock_out_detail_position,omitempty" form:"stock_out_detail_position" bson:"stock_out_detail_position"`
	StockOutRecordContent   []map[string]interface{} `json:"stock_out_record_content" form:"stock_out_record_content" bson:"stock_out_record_content"`
	Remark                  *string                  `json:"remark,omitempty" form:"remark" bson:"remark"`
}

func (m StockOutRecord) ToESDoc() map[string]interface{} {
	return map[string]interface{}{
		"rec_id":                m.RecID,
		"created_at":            m.CreatedAt,
		"remark":                m.Remark,
		"stock_out_content":     m.StockOutRecordContent,
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
