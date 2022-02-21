package model

// CommodityBatch 批次
type CommodityBatch struct {
	BasicModel
}

// Commodity 产品
type Commodity struct {
	BasicModel
	CommodityType      *string                `gorm:"type:varchar(50)"`
	CommodityStyle     *string                `gorm:"type:varchar(50)"`
	CommodityBatchID   int                    `gorm:"type:int"`
	CommodityOwnerID   int                    `gorm:"type:int"`
	CommodityOwnerName string                 `gorm:"type:varchar(50)"`
	CommodityRemark    *string                `gorm:"type:varchar(200)"`
	CommodityAttribute map[string]interface{} `gorm:"type:varchar(500)"`
}
