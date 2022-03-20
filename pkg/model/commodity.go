package model

// CommodityBatch 批次
type CommodityBatch struct {
	BasicModel
}

// Commodity 产品
type Commodity struct {
	BasicModel
	CommodityType      *string                `gorm:"type:varchar(50)" json:"commodity_type,omitempty"`
	CommodityStyle     *string                `gorm:"type:varchar(50)" json:"commodity_style,omitempty"`
	CommodityBatchID   int                    `gorm:"type:int" json:"commodity_batch_id"`
	CommodityOwnerID   int                    `gorm:"type:int" json:"commodity_owner_id"`
	CommodityOwnerName string                 `gorm:"type:varchar(50)" json:"commodity_owner_name"`
	CommodityRemark    *string                `gorm:"type:varchar(200)" json:"commodity_remark,omitempty"`
	CommodityAttribute map[string]interface{} `gorm:"type:varchar(500)" json:"commodity_attribute"`
}
