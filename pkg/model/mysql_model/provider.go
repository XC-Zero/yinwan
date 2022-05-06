package mysql_model

import (
	"github.com/XC-Zero/yinwan/pkg/client"
	"github.com/olivere/elastic/v7"
	"gorm.io/gorm"
)

// Provider 供应商
type Provider struct {
	BasicModel
	ProviderName       string  `gorm:"type:varchar(200);not null" json:"provider_name" form:"provider_name" binding:"required"`
	ProviderAlias      *string `gorm:"type:varchar(200);not null" form:"provider_alias" json:"provider_alias,omitempty" binding:"required"`
	AccumulatedAmount  float64 `gorm:"type:decimal(20,2); " json:"accumulated_amount" form:"accumulated_amount"`
	ProviderLogoUrl    *string `gorm:"type:varchar(500); " json:"provider_pic_url,omitempty" form:"provider_pic_url" `
	ProviderOwner      *string `gorm:"type:varchar(50);" json:"provider_owner" form:"provider_owner"`
	ProviderOwnerPhone *string `gorm:"type:varchar(50);"  json:"provider_owner_phone" form:"provider_owner_phone"`
	OurOwnerID         *int    `json:"our_owner_id,omitempty" form:"our_owner_id"`
	OurOwnerName       *string `gorm:"type:varchar(50);"  json:"our_owner_name" form:"our_owner_name"`
	Remark             *string `form:"remark" json:"remark,omitempty" gorm:"type:varchar(200)"  cn:"备注"`
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
func (p Provider) ToESDoc() map[string]interface{} {
	return map[string]interface{}{
		"rec_id":           p.RecID,
		"remark":           p.Remark,
		"created_at":       p.CreatedAt,
		"provider_pic_url": p.ProviderLogoUrl,
		"customer_name":    p.CustomerName,
		"customer_contact": p.CustomerContact,
	}
}
func (p *Provider) AfterCreate(db *gorm.DB) error {
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
func (p *Provider) AfterUpdate(_ *gorm.DB) error {
	err := client.UpdateIntoIndex(c, c.RecID,
		elastic.NewScriptInline("ctx._source.nickname=params.nickname;ctx._source.ancestral=params.ancestral").
			Params(c.ToESDoc()))
	if err != nil {
		return err
	}
	return nil
}
func (p *Provider) AfterDelete(_ *gorm.DB) error {
	err := client.DeleteFromIndex(c, c.RecID)
	if err != nil {
		return err
	}
	return nil
}
