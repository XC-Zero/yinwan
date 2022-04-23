package mysql_model

import "github.com/XC-Zero/yinwan/pkg/model/common"

// Provider 供应商
type Provider struct {
	common.BasicModel
	ProviderName       string  `gorm:"type:varchar(200);not null" json:"provider_name" form:"provider_name" binding:"required"`
	AccumulatedAmount  float64 `gorm:"type:decimal(20,2); " json:"accumulated_amount" form:"accumulated_amount"`
	ProviderLogoUrl    *string `gorm:"type:varchar(500); " json:"provider_pic_url,omitempty" form:"provider_pic_url" `
	ProviderOwner      *string `gorm:"type:varchar(50);" json:"provider_owner" form:"provider_owner"`
	ProviderOwnerPhone *string `gorm:"type:varchar(50);"  json:"provider_owner_phone" form:"provider_owner_phone"`
	OurOwnerID         *int    `json:"our_owner_id,omitempty" form:"our_owner_id"`
	OurOwnerName       *string `gorm:"type:varchar(50);"  json:"our_owner_name" form:"our_owner_name"`
}

func (p Provider) TableCnName() string {
	return "供应商"
}

func (p Provider) TableName() string {
	return "providers"

}
