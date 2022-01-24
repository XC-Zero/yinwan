package model

// CommodityLot 批次
type CommodityLot struct {
}

// Commodity 货品
type Commodity struct {
	CommodityID        *int                   `gorm:"primaryKey;type:int;autoIncrement" json:"id"`
	CommodityType      *string                `gorm:"type:varchar(50)"`
	CommodityStyle     *string                `gorm:"type:varchar(50)"`
	CommodityPrice     float64                `gorm:"type:decimal(20,4)"`
	CommodityLotID     int                    `gorm:"type:int"`
	CommodityManager   int                    `gorm:"type:int"`
	CommodityRemark    *string                `gorm:"type:varchar(200)"`
	CommodityAttribute map[string]interface{} `gorm:"type:varchar(500)"`
}
