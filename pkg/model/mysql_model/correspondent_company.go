package mysql_model

import "github.com/XC-Zero/yinwan/pkg/model/common"

// CorrespondentCompany 往来公司
type CorrespondentCompany struct {
	common.BasicModel
	CompanyName             string  `json:"company_name" gorm:"type:varchar(200);not null;" cn:"往来公司名称"`
	CompanyLegalName        *string `json:"company_legal_name,omitempty" gorm:"type:varchar(50)" cn:"往来公司全称"`
	CompanyAlias            *string `json:"company_alias,omitempty" gorm:"type:varchar(50)" cn:"往来公司简称"`
	CompanyAddress          *string `json:"company_address,omitempty"  gorm:"type:varchar(500);" cn:"往来公司地址"`
	CompanySocialCreditCode *string `json:"company_social_credit_code,omitempty" gorm:"type:varchar(50)" cn:"社会信用代码"`
	CompanyContact          *string `json:"company_contact,omitempty" gorm:"type:varchar(50)" cn:"往来公司方联系人"`
	CompanyContactPhone     *string `json:"company_contact_phone,omitempty" gorm:"type:varchar(20)" cn:"联系人电话"`
	CompanyContactWechat    *string `json:"company_contact_wechat,omitempty" gorm:"type:varchar(50)" cn:"联系人微信"`
	CompanyOwnerID          *int    `json:"company_owner_id,omitempty" gorm:"type:varchar(50)" cn:"往来公司负责人ID"`
	CompanyOwnerName        *string `json:"company_owner_name,omitempty"  gorm:"type:varchar(50)" cn:"往来公司负责人名称"`
	IsCustomer              bool    `json:"is_customer"`
	IsProvider              bool    `json:"is_provider"`
}
