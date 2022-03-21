package model

// CommodityBatch 批次
type CommodityBatch struct {
	BasicModel
}

// Commodity 产品
type Commodity struct {
	BasicModel
	CommodityName      string                 `gorm:"type:varchar(200)" json:"commodity_name"`
	CommodityType      *string                `gorm:"type:varchar(50)" json:"commodity_type,omitempty"`
	CommodityStyle     *string                `gorm:"type:varchar(50)" json:"commodity_style,omitempty"`
	CommodityOwnerID   int                    `gorm:"type:int" json:"commodity_owner_id"`
	CommodityOwnerName string                 `gorm:"type:varchar(50)" json:"commodity_owner_name"`
	CommodityRemark    *string                `gorm:"type:varchar(200)" json:"commodity_remark,omitempty"`
	CommodityAttribute map[string]interface{} `gorm:"type:varchar(500)" json:"commodity_attribute"`
}

// CommodityHistoricalCost 历史成本表
type CommodityHistoricalCost struct {
	BasicModel
	CommodityID   int     `gorm:"type:int;index;" json:"commodity_id"`
	CommodityName string  `gorm:"type:varchar(200)" json:"commodity_name"`
	CommodityCost float64 `gorm:"type:decimal(20,2)" json:"commodity_cost"`
}
