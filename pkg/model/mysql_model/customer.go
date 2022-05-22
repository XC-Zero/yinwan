package mysql_model

import (
	"github.com/XC-Zero/yinwan/pkg/client"
	"github.com/XC-Zero/yinwan/pkg/utils/es_tool"
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
	CustomerDetailAddress    *string `gorm:"type:varchar(500);" json:"customer_detail_address,omitempty" form:"customer_detail_address" cn:"客户详细地址"`
	CustomerSocialCreditCode *string `form:"customer_social_credit_code" json:"customer_social_credit_code,omitempty" gorm:"type:varchar(50)" cn:"社会信用代码"`
	CustomerContact          *string `form:"customer_contact" json:"customer_contact,omitempty" gorm:"type:varchar(50)" cn:"客户方联系人"`
	CustomerContactPhone     *string `form:"customer_contact_phone" json:"customer_contact_phone,omitempty" gorm:"type:varchar(20)" cn:"联系人电话"`
	CustomerContactWechat    *string `form:"customer_contact_wechat" json:"customer_contact_wechat,omitempty" gorm:"type:varchar(50)" cn:"联系人微信"`
	CustomerOwnerID          *int    `form:"customer_owner_id" json:"customer_owner_id,omitempty" gorm:"type:varchar(50)" cn:"客户负责人ID"`
	CustomerOwnerName        *string `form:"customer_owner_name" json:"customer_owner_name,omitempty" gorm:"type:varchar(50)" cn:"客户负责人名称"`
	AccumulateReceiveAmount  string  `gorm:"type:varchar(50);default:0"  json:"accumulate_receive_amount" form:"accumulate_receive_amount" cn:"累计收款"`
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
				"provider_social_credit_code": mapping{
					"type":         "keyword",
					"ignore_above": 256,
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
		"rec_id":                      c.RecID,
		"remark":                      c.Remark,
		"created_at":                  c.CreatedAt,
		"provider_social_credit_code": c.CustomerSocialCreditCode,
		"customer_alias":              c.CustomerAlias,
		"customer_name":               c.CustomerName,
		"customer_contact":            c.CustomerContact,
	}
}
func (c *Customer) AfterCreate(tx *gorm.DB) error {
	bookName := tx.Statement.Context.Value("book_name").(string)
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

// AfterUpdate 同步更新
func (c *Customer) AfterUpdate(tx *gorm.DB) error {
	bookName := tx.Statement.Context.Value("book_name").(string)
	bk, ok := client.ReadBookMap(bookName)
	if !ok {
		return errors.New("There is no book name!")
	}
	c.BookNameID = bk.StorageName
	c.BookName = bk.BookName
	err := client.UpdateIntoIndex(c, c.RecID, tx.Statement.Context, es_tool.ESDocToUpdateScript(c.ToESDoc()))

	if err != nil {
		return err
	}
	return nil
}
func (c *Customer) AfterDelete(tx *gorm.DB) error {
	err := client.DeleteFromIndex(c, c.RecID, tx.Statement.Context)
	if err != nil {
		return err
	}
	return nil
}
