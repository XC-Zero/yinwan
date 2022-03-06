package model

// StockInRecord 入库记录
type StockInRecord struct {
	BasicModel
	RecordOwnerID   int                    `gorm:"type:int;not null;index"`
	RecordOwnerName string                 `gorm:"type:int;not null;index"`
	RecordType      string                 `gorm:"type:varchar(50);not null"`
	RecordContent   map[string]interface{} `gorm:"type:varchar(500);not null"`
	RecordRemark    *string                `gorm:"type:varchar(200);"`
}

// StockOutRecord 出库记录
type StockOutRecord struct {
	BasicModel
	RecordOwnerID   int                    `gorm:"type:int;not null;index"`
	RecordOwnerName string                 `gorm:"type:int;not null;index"`
	RecordType      string                 `gorm:"type:varchar(50);not null"`
	RecordContent   map[string]interface{} `gorm:"type:varchar(500);not null"`
	RecordRemark    *string                `gorm:"type:varchar(200);"`
}
