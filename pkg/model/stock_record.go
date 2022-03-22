package model

// StockInRecord 入库记录
// 存 MongoDB 一份
type StockInRecord struct {
	BasicModel
	StockInRecordOwnerID   int                    `gorm:"type:int;not null;index" json:"stock_in_record_owner_id" bson:"stock_in_record_owner_id"`
	StockInRecordOwnerName string                 `gorm:"type:int;not null" json:"stock_in_record_owner_name" bson:"stock_in_record_owner_name"`
	StockInRecordType      string                 `gorm:"type:varchar(50);not null" json:"stock_in_record_type" bson:"stock_in_record_type"`
	StockInRecordContent   map[string]interface{} `gorm:"type:varchar(500);not null" json:"stock_in_record_content" bson:"stock_in_record_content"`
	Remark                 *string                `gorm:"type:varchar(200);" json:"remark,omitempty" bson:"remark"`
}

// StockOutRecord 出库记录
// 存 MongoDB 一份
type StockOutRecord struct {
	BasicModel
	StockOutRecordOwnerID   int                    `gorm:"type:int;not null;index" json:"stock_out_record_owner_id" bson:"stock_out_record_owner_id"`
	StockOutRecordOwnerName string                 `gorm:"type:int;not null" json:"stock_out_record_owner_name" bson:"stock_out_record_owner_name"`
	StockOutRecordType      string                 `gorm:"type:varchar(50);not null" json:"stock_out_record_type" bson:"stock_out_record_type"`
	StockOutRecordContent   map[string]interface{} `gorm:"type:varchar(500);not null" json:"stock_out_record_content" bson:"stock_out_record_content"`
	Remark                  *string                `gorm:"type:varchar(200);" json:"remark,omitempty" bson:"remark"`
}
