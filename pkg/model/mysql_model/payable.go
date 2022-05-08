package mysql_model

import (
	"github.com/XC-Zero/yinwan/pkg/client"
	"github.com/olivere/elastic/v7"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

// Payable 应付
type Payable struct {
	BasicModel
	BookNameInfo
	//	关联凭证ID
	CredentialID *int    `gorm:"type:int" json:"credential_id,omitempty"`
	Remark       *string `gorm:"type:varchar(200)" json:"remark,omitempty"`
}

func (p Payable) TableCnName() string {
	return "应付"
}
func (p Payable) TableName() string {
	return "payables"

}
func (p Payable) ToESDoc() map[string]interface{} {
	return map[string]interface{}{
		"rec_id": p.RecID,
		"remark": p.Remark,
		//"payable_enterprise":p
	}
}
func (p Payable) Mapping() map[string]interface{} {
	m := mapping{
		"settings": mapping{},
		"mappings": mapping{
			"properties": mapping{
				"rec_id": mapping{
					"type": "keyword",
				},
				"payable_enterprise": mapping{
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
				"payable_enterprise_address": mapping{
					"type":            "text",
					"analyzer":        IK_SMART,
					"search_analyzer": IK_SMART,
				},
				"payable_contact": mapping{
					"type":            "text",
					"analyzer":        IK_SMART,
					"search_analyzer": IK_SMART,
				},
				"payable_amount": mapping{
					"type": "keyword",
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
func (p *Payable) AfterCreate(tx *gorm.DB) error {
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

// AfterUpdate todo !!!
func (p *Payable) AfterUpdate(tx *gorm.DB) error {

	err := client.UpdateIntoIndex(p, p.RecID, tx,
		elastic.NewScriptInline("ctx._source.nickname=params.nickname;ctx._source.ancestral=params.ancestral").
			Params(p.ToESDoc()))
	if err != nil {
		return err
	}
	return nil
}
func (p *Payable) AfterDelete(tx *gorm.DB) error {

	err := client.DeleteFromIndex(p, p.RecID, tx)
	if err != nil {
		return err
	}
	return nil
}
