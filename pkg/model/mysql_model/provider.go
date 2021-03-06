package mysql_model

import (
	"github.com/XC-Zero/yinwan/pkg/client"
	"github.com/XC-Zero/yinwan/pkg/utils/es_tool"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

// Provider 供应商
type Provider struct {
	BasicModel
	BookNameInfo
	ProviderName             string  `form:"provider_name" json:"provider_name" gorm:"type:varchar(200);not null;" cn:"供应商名称"`
	ProviderLegalName        *string `form:"provider_legal_name" json:"provider_legal_name,omitempty" gorm:"type:varchar(50)" cn:"供应商公司全称"`
	ProviderAlias            *string `form:"provider_alias" json:"provider_alias,omitempty" gorm:"type:varchar(50)" cn:"供应商简称"`
	ProviderLogoUrl          *string `gorm:"type:varchar(500); " form:"provider_logo_url" json:"provider_logo_url,omitempty" cn:"供应商头像地址"`
	ProviderAddress          *string `form:"provider_address" json:"provider_address,omitempty"  gorm:"type:varchar(500);" cn:"供应商省市区"`
	ProviderDetailAddress    *string `gorm:"type:varchar(500);" json:"provider_detail_address,omitempty" form:"provider_detail_address,omitempty" cn:"供应商详细地址"`
	ProviderSocialCreditCode *string `form:"provider_social_credit_code" json:"provider_social_credit_code,omitempty" gorm:"type:varchar(50)" cn:"社会信用代码"`
	ProviderContact          *string `form:"provider_contact" json:"provider_contact,omitempty" gorm:"type:varchar(50)" cn:"供应商方联系人"`
	ProviderContactPhone     *string `form:"provider_contact_phone" json:"provider_contact_phone,omitempty" gorm:"type:varchar(20)" cn:"联系人电话"`
	ProviderContactWechat    *string `form:"provider_contact_wechat" json:"provider_contact_wechat,omitempty" gorm:"type:varchar(50)" cn:"联系人微信"`
	ProviderOwnerID          *int    `form:"provider_owner_id" json:"provider_owner_id,omitempty" gorm:"type:varchar(50)" cn:"供应商负责人ID"`
	ProviderOwnerName        *string `form:"provider_owner_name" json:"provider_owner_name,omitempty" gorm:"type:varchar(50)" cn:"供应商负责人名称"`
	AccumulatePayAmount      string  `json:"accumulate_pay_amount" gorm:"type:varchar(50);default:0" cn:"累计付款"`
	Remark                   *string `form:"remark" json:"remark,omitempty" gorm:"type:varchar(200)"  cn:"备注"`
}

func (p Provider) TableCnName() string {
	return "供应商"
}

func (p Provider) TableName() string {
	return "providers"

}

func (p Provider) Mapping() map[string]interface{} {
	m := mapping{
		"settings": mapping{},
		"mappings": mapping{
			"properties": mapping{
				"rec_id": mapping{
					"type": "keyword",
				},
				"provider_name": mapping{
					"type":            "text",
					"analyzer":        IK_SMART,
					"search_analyzer": IK_SMART,
					"fields": mapping{
						"keyword": mapping{
							"type":         "keyword",
							"ignore_above": 256,
						},
					},
				},
				"provider_social_credit_code": mapping{
					"type":         "keyword",
					"ignore_above": 256,
				},

				"provider_contact": mapping{
					"type":            "text",
					"analyzer":        IK_SMART,
					"search_analyzer": IK_SMART,
				},
				"provider_alias": mapping{
					"type":            "text",
					"analyzer":        IK_SMART,
					"search_analyzer": IK_SMART,
				},
				"provider_address": mapping{
					"type":            "text",
					"analyzer":        IK_SMART,
					"search_analyzer": IK_SMART,
				},

				"remark": mapping{
					"type":            "text",
					"analyzer":        IK_SMART,
					"search_analyzer": IK_SMART,
				},
				"created_at": mapping{
					"type": "text",
				},
				"book_name": mapping{
					"type": "keyword",
				},
				"book_name_id": mapping{
					"type": "keyword",
				},
			},
		},
	}
	return m
}
func (p Provider) ToESDoc() map[string]interface{} {
	var address string
	if p.ProviderAddress != nil {
		address += *p.ProviderAddress
	}
	if p.ProviderDetailAddress != nil {
		address += *p.ProviderDetailAddress
	}
	return map[string]interface{}{
		"rec_id":                      p.RecID,
		"remark":                      p.Remark,
		"created_at":                  p.CreatedAt,
		"provider_social_credit_code": p.ProviderSocialCreditCode,
		"provider_pic_url":            p.ProviderLogoUrl,
		"provider_name":               p.ProviderName,
		"provider_contact":            p.ProviderContact,
		"provider_address":            address,
		"book_name":                   p.BookName,
		"book_name_id":                p.BookNameID,
	}
}
func (p *Provider) AfterCreate(tx *gorm.DB) error {
	bookName := tx.Statement.Context.Value("book_name").(string)
	bk, ok := client.ReadBookMap(bookName)
	if !ok {
		return errors.New("There is no book name!")
	}
	p.BookNameID = bk.StorageName
	p.BookName = bk.BookName
	err := client.PutIntoIndex(p)
	if err != nil {
		return err
	}
	return nil
}

// AfterUpdate 同步更新
func (p *Provider) AfterUpdate(tx *gorm.DB) error {
	bookName := tx.Statement.Context.Value("book_name").(string)
	bk, ok := client.ReadBookMap(bookName)
	if !ok {
		return errors.New("There is no book name!")
	}
	p.BookNameID = bk.StorageName
	p.BookName = bk.BookName
	err := client.UpdateIntoIndex(p, p.RecID, tx.Statement.Context, es_tool.ESDocToUpdateScript(p.ToESDoc()))

	if err != nil {
		return err
	}
	return nil
}
func (p *Provider) AfterDelete(tx *gorm.DB) error {

	err := client.DeleteFromIndex(p, p.RecID, tx.Statement.Context)
	if err != nil {
		return err
	}
	return nil
}
