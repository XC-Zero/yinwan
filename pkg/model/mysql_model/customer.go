package mysql_model

// Customer 客户
type Customer struct {
	BasicModel
	CustomerName             string  `json:"customer_name" gorm:"type:varchar(200);not null;" cn:"客户名称"`
	CustomerLegalName        *string `json:"customer_legal_name" gorm:"type:varchar(50)" cn:"客户公司全称"`
	CustomerAlias            *string `json:"customer_alias" gorm:"type:varchar(50)" cn:"客户简称"`
	CustomerAddress          *string `json:"customer_address"  gorm:"type:varchar(500);" cn:"客户地址"`
	CustomerSocialCreditCode *string `json:"customer_social_credit_code" gorm:"type:varchar(50)" cn:"社会信用代码"`
	CustomerContact          *string `json:"customer_contact" gorm:"type:varchar(50)" cn:"客户方联系人"`
	CustomerContactPhone     *string `json:"customer_contact_phone" gorm:"type:varchar(20)" cn:"联系人电话"`
	CustomerContactWechat    *string `json:"customer_contact_wechat" gorm:"type:varchar(50)" cn:"联系人微信"`
	CustomerOwnerID          *int    `json:"customer_owner_id" gorm:"type:varchar(50)" cn:"客户负责人ID"`
	CustomerOwnerName        *string `json:"customer_owner_name"  gorm:"type:varchar(50)" cn:"客户负责人名称"`
}

func (c Customer) TableCnName() string {
	return "客户"
}
func (c Customer) TableName() string {
	return "customers"
}
