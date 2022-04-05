package model

// CommodityBatch 批次
type CommodityBatch struct {
	BasicModel
}

func (c CommodityBatch) TableCnName() string {
	return "产品批次"
}
func (c CommodityBatch) TableName() string {
	return "commodity_batches"
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

func (c Commodity) TableCnName() string {
	return "产品"
}
func (c Commodity) TableName() string {
	return "commodities"
}

// CommodityHistoricalCost 历史成本表
type CommodityHistoricalCost struct {
	BasicModel
	CommodityID   int     `gorm:"type:int;index;" json:"commodity_id"`
	CommodityName string  `gorm:"type:varchar(200)" json:"commodity_name"`
	CommodityCost float64 `gorm:"type:decimal(20,2)" json:"commodity_cost"`
}

func (c CommodityHistoricalCost) TableCnName() string {
	return "历史成本"
}
func (c CommodityHistoricalCost) TableName() string {
	return "commodity_historical_costs"
}
