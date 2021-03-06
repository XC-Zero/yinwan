package mysql_model

import (
	"github.com/XC-Zero/yinwan/pkg/client"
	_const "github.com/XC-Zero/yinwan/pkg/const"
	"github.com/XC-Zero/yinwan/pkg/utils/es_tool"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

// Payable 应付
type Payable struct {
	BasicModel
	BookNameInfo
	ProviderID           *int                  `gorm:"type:int" json:"provider_id,omitempty" form:"provider_id,omitempty"`
	ProviderName         *string               `gorm:"type:varchar(50)" json:"provider_name,omitempty" form:"provider_name"`
	PayableDate          *string               `gorm:"type:varchar(50)" json:"payable_date,omitempty" form:"payable_date"`
	PayableTotalAmount   *string               `gorm:"type:varchar(50);" json:"payable_total_amount,omitempty" form:"payable_total_amount"`
	PayableActualAmount  *string               `gorm:"type:varchar(50)" json:"payable_actual_amount,omitempty" form:"payable_actual_amount"`
	PayableDiscount      *string               `gorm:"type:varchar(50)" json:"payable_discount,omitempty" form:"payable_discount"`
	PayableHedgingAmount *string               `gorm:"type:varchar(50)" json:"payable_hedging_amount,omitempty" form:"payable_hedging_amount"`
	PayableDebtAmount    *string               `gorm:"type:varchar(50)" json:"payable_debt_amount,omitempty" form:"payable_debt_amount"`
	PayableStatus        *_const.PayableStatus `gorm:"type:int" json:"payable_status,omitempty" form:"payable_status"`
	PurchaseID           *int                  `gorm:"type:int" json:"purchase_id,omitempty" form:"purchase_id,omitempty" cn:"关联采购单ID"`
	CredentialID         *int                  `gorm:"type:int" json:"credential_id,omitempty" form:"credential_id,omitempty"`
	Remark               *string               `gorm:"type:varchar(200)" json:"remark,omitempty" form:"remark"`
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

// AfterUpdate 同步更新
func (p *Payable) AfterUpdate(tx *gorm.DB) error {
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
func (p *Payable) AfterDelete(tx *gorm.DB) error {

	err := client.DeleteFromIndex(p, p.RecID, tx.Statement.Context)
	if err != nil {
		return err
	}
	return nil
}
