package model

// StockInRecord 入库记录
// 存 MongoDB
type StockInRecord struct {
	BasicModel
	StockInRecordOwnerID   int                    ` json:"stock_in_record_owner_id" bson:"stock_in_record_owner_id"`
	StockInRecordOwnerName string                 ` json:"stock_in_record_owner_name" bson:"stock_in_record_owner_name"`
	StockInRecordType      string                 ` json:"stock_in_record_type" bson:"stock_in_record_type"`
	StockInRecordContent   map[string]interface{} ` json:"stock_in_record_content" bson:"stock_in_record_content"`
	Remark                 *string                ` json:"remark,omitempty" bson:"remark"`
}

func (m StockInRecord) TableName() string {
	return "stock_in_records"
}

// StockOutRecord 出库记录
// 存 MongoDB 一份
type StockOutRecord struct {
	BasicModel
	StockOutRecordOwnerID   int                    `json:"stock_out_record_owner_id" bson:"stock_out_record_owner_id"`
	StockOutRecordOwnerName string                 `json:"stock_out_record_owner_name" bson:"stock_out_record_owner_name"`
	StockOutRecordType      string                 `json:"stock_out_record_type" bson:"stock_out_record_type"`
	StockOutRecordContent   map[string]interface{} `json:"stock_out_record_content" bson:"stock_out_record_content"`
	Remark                  *string                `json:"remark,omitempty" bson:"remark"`
}

func (m StockOutRecord) TableName() string {
	return "stock_out_records"
}
