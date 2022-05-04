package mongo_model

import (
	"github.com/XC-Zero/yinwan/pkg/model/mysql_model"
	"gorm.io/gorm"
)

// StockInRecord 入库记录
// 存 MongoDB
type StockInRecord struct {
	mysql_model.BasicModel
	StockInRecordOwnerID   int                    ` json:"stock_in_record_owner_id"  form:"stock_in_record_owner_id" bson:"stock_in_record_owner_id" binding:"required"`
	StockInRecordOwnerName string                 ` json:"stock_in_record_owner_name" form:"stock_in_record_owner_name" bson:"stock_in_record_owner_name" binding:"required"`
	StockInRecordType      string                 ` json:"stock_in_record_type" form:"stock_in_record_type" bson:"stock_in_record_type" binding:"required"`
	StockInRecordContent   map[string]interface{} ` json:"stock_in_record_content" form:"stock_in_record_content" bson:"stock_in_record_content" binding:"required"`
	RelatePurchaseID       *int                   ` json:"relate_purchase_id,omitempty" form:"relate_purchase_id,omitempty"`
	Remark                 *string                ` json:"remark,omitempty" form:"remark" bson:"remark"`
}

func (m StockInRecord) ToESDoc() map[string]interface{} {
	//TODO implement me
	panic("implement me")
}

func (m *StockInRecord) AfterCreate(db *gorm.DB) error {
	//TODO implement me
	panic("implement me")
}

func (m *StockInRecord) AfterUpdate(db *gorm.DB) error {
	//TODO implement me
	panic("implement me")
}

func (m *StockInRecord) AfterDelete(db *gorm.DB) error {
	//TODO implement me
	panic("implement me")
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

// StockOutRecord 出库记录
// 存 MongoDB 一份
type StockOutRecord struct {
	mysql_model.BasicModel
	StockOutRecordOwnerID   int                    `json:"stock_out_record_owner_id" form:"stock_out_record_owner_id" bson:"stock_out_record_owner_id"`
	StockOutRecordOwnerName string                 `json:"stock_out_record_owner_name" form:"stock_out_record_owner_name" bson:"stock_out_record_owner_name"`
	StockOutRecordType      string                 `json:"stock_out_record_type" form:"stock_out_record_type" bson:"stock_out_record_type"`
	StockOutRecordContent   map[string]interface{} `json:"stock_out_record_content" form:"stock_out_record_content" bson:"stock_out_record_content"`
	Remark                  *string                `json:"remark,omitempty" form:"remark" bson:"remark"`
}

func (m StockOutRecord) ToESDoc() map[string]interface{} {
	//TODO implement me
	panic("implement me")
}

func (m *StockOutRecord) AfterCreate(db *gorm.DB) error {
	//TODO implement me
	panic("implement me")
}

func (m *StockOutRecord) AfterUpdate(db *gorm.DB) error {
	//TODO implement me
	panic("implement me")
}

func (m *StockOutRecord) AfterDelete(db *gorm.DB) error {
	//TODO implement me
	panic("implement me")
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
