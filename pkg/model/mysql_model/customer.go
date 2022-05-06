package mysql_model

import (
	"github.com/XC-Zero/yinwan/pkg/client"
	"github.com/olivere/elastic/v7"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

// Customer 客户
type Customer struct {
	BasicModel
	BookNameInfo
	CustomerName             string  `form:"customer_name" json:"customer_name" gorm:"type:varchar(200);not null;" cn:"客户名称"`
	CustomerLegalName        *string `form:"customer_legal_name" json:"customer_legal_name,omitempty" gorm:"type:varchar(50)" cn:"客户公司全称"`
	CustomerAlias            *string `form:"customer_alias" json:"customer_alias,omitempty" gorm:"type:varchar(50)" cn:"客户简称"`
	CustomerLogoUrl          *string `gorm:"type:varchar(500); " form:"customer_logo_url" json:"customer_logo_url,omitempty" cn:"客户头像地址"`
	CustomerAddress          *string `form:"customer_address" json:"customer_address,omitempty"  gorm:"type:varchar(500);" cn:"客户地址"`
	CustomerSocialCreditCode *string `form:"customer_social_credit_code" json:"customer_social_credit_code,omitempty" gorm:"type:varchar(50)" cn:"社会信用代码"`
	CustomerContact          *string `form:"customer_contact" json:"customer_contact,omitempty" gorm:"type:varchar(50)" cn:"客户方联系人"`
	CustomerContactPhone     *string `form:"customer_contact_phone" json:"customer_contact_phone,omitempty" gorm:"type:varchar(20)" cn:"联系人电话"`
	CustomerContactWechat    *string `form:"customer_contact_wechat" json:"customer_contact_wechat,omitempty" gorm:"type:varchar(50)" cn:"联系人微信"`
	CustomerOwnerID          *int    `form:"customer_owner_id" json:"customer_owner_id,omitempty" gorm:"type:varchar(50)" cn:"客户负责人ID"`
	CustomerOwnerName        *string `form:"customer_owner_name" json:"customer_owner_name,omitempty" gorm:"type:varchar(50)" cn:"客户负责人名称"`
	Remark                   *string `form:"remark" json:"remark,omitempty" gorm:"type:varchar(200)"  cn:"备注"`
}

func (c Customer) TableCnName() string {
	return "客户"
}
func (c Customer) TableName() string {
	return "customers"
}
func (c Customer) Mapping() map[string]interface{} {
	m := mapping{
		"settings": mapping{},
		"mappings": mapping{
			"properties": mapping{
				"rec_id": mapping{
					"type": "keyword",
				},
				"customer_name": mapping{
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

				"customer_contact": mapping{
					"type":            "text",
					"analyzer":        IK_SMART,
					"search_analyzer": IK_SMART,
				},
				"customer_alias": mapping{
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
func (c Customer) ToESDoc() map[string]interface{} {
	return map[string]interface{}{
		"rec_id":           c.RecID,
		"remark":           c.Remark,
		"created_at":       c.CreatedAt,
		"customer_alias":   c.CustomerAlias,
		"customer_name":    c.CustomerName,
		"customer_contact": c.CustomerContact,
	}
}
func (c *Customer) AfterCreate(db *gorm.DB) error {
	bookName := db.Statement.Context.Value("book_name").(string)
	bk, ok := client.ReadBookMap(bookName)
	if !ok {
		return errors.New("There is no book name!")
	}
	c.BookNameID = bk.StorageName
	c.BookName = bk.BookName
	err := client.PutIntoIndex(c)
	if err != nil {
		return err
	}
	return nil
}

// AfterUpdate todo !!!
func (c *Customer) AfterUpdate(_ *gorm.DB) error {
	err := client.UpdateIntoIndex(c, c.RecID,
		elastic.NewScriptInline("ctx._source.nickname=params.nickname;ctx._source.ancestral=params.ancestral").
			Params(c.ToESDoc()))
	if err != nil {
		return err
	}
	return nil
}
func (c *Customer) AfterDelete(_ *gorm.DB) error {
	err := client.DeleteFromIndex(c, c.RecID)
	if err != nil {
		return err
	}
	return nil
}
