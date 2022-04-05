package model

// Provider 供应商
type Provider struct {
	BasicModel
	ProviderName      string  `gorm:"type:varchar(200);not null" json:"provider_name" form:"provider_name" binding:"required"`
	AccumulatedAmount float64 `gorm:"type:decimal(20,2);not null" json:"accumulated_amount" form:"accumulated_amount"`
}

func (p Provider) TableCnName() string {
	return "供应商"
}

func (p Provider) TableName() string {
	return "providers"

}
