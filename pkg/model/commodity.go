package model

// CommodityBatch 批次
type CommodityBatch struct {
	BasicModel
}

// Commodity 材料
type Commodity struct {
	BasicModel
	CommodityType      *string                `gorm:"type:varchar(50)"`
	CommodityStyle     *string                `gorm:"type:varchar(50)"`
	CommodityPrice     float64                `gorm:"type:decimal(20,4)"`
	CommodityBatchID   int                    `gorm:"type:int"`
	CommodityManager   int                    `gorm:"type:int"`
	CommodityRemark    *string                `gorm:"type:varchar(200)"`
	CommodityAttribute map[string]interface{} `gorm:"type:varchar(500)"`
}
