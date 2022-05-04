package mysql_model

import (
	"github.com/XC-Zero/yinwan/pkg/client"
	"github.com/olivere/elastic/v7"
	"gorm.io/gorm"
)

// Currency 货币
type Currency struct {
	BasicModel
	// 货币名称
	CurrencyName string `gorm:"type:varchar(50)"`
	// 货币符号
	CurrencySymbol string `gorm:"type:varchar(50)"`
	// 对比人民币的汇率
	CurrencyExchangeRate float64 `gorm:"type:decimal(20,4)"`
}

// Payable 应付
type Payable struct {
	BasicModel
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
func (p *Payable) AfterCreate(_ *gorm.DB) error {
	err := client.PutIntoIndex(p)
	if err != nil {
		return err
	}
	return nil
}

// AfterUpdate todo !!!
func (p *Payable) AfterUpdate(_ *gorm.DB) error {
	err := client.UpdateIntoIndex(p, p.RecID,
		elastic.NewScriptInline("ctx._source.nickname=params.nickname;ctx._source.ancestral=params.ancestral").
			Params(p.ToESDoc()))
	if err != nil {
		return err
	}
	return nil
}
func (p *Payable) AfterDelete(_ *gorm.DB) error {
	err := client.DeleteFromIndex(p, p.RecID)
	if err != nil {
		return err
	}
	return nil
}

// Receivable 应收
type Receivable struct {
	BasicModel
	//	关联凭证ID
	CredentialID *int    `gorm:"type:int" json:"credential_id,omitempty"`
	Remark       *string `gorm:"type:varchar(200)" json:"remark,omitempty"`
}

func (p Receivable) TableCnName() string {
	return "应收"
}
func (p Receivable) TableName() string {
	return "receivables"

}
func (p Receivable) Mapping() map[string]interface{} {
	m := mapping{
		"settings": mapping{},
		"mappings": mapping{
			"properties": mapping{
				"rec_id": mapping{
					"type": "keyword",
				},
				"receivable_enterprise": mapping{
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
				"receivable_enterprise_address": mapping{
					"type":            "text",
					"analyzer":        IK_SMART,
					"search_analyzer": IK_SMART,
				},
				"receivable_contact": mapping{
					"type":            "text",
					"analyzer":        IK_SMART,
					"search_analyzer": IK_SMART,
				},
				"receivable_amount": mapping{
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
func (p Receivable) ToESDoc() map[string]interface{} {
	return map[string]interface{}{
		"rec_id": p.RecID,
		"remark": p.Remark,
		//"payable_enterprise":p
	}
}
func (p *Receivable) AfterCreate(_ *gorm.DB) error {
	err := client.PutIntoIndex(p)
	if err != nil {
		return err
	}
	return nil
}

// AfterUpdate todo !!!
func (p *Receivable) AfterUpdate(_ *gorm.DB) error {
	err := client.UpdateIntoIndex(p, p.RecID,
		elastic.NewScriptInline("ctx._source.nickname=params.nickname;ctx._source.ancestral=params.ancestral").
			Params(p.ToESDoc()))
	if err != nil {
		return err
	}
	return nil
}
func (p *Receivable) AfterDelete(_ *gorm.DB) error {
	err := client.DeleteFromIndex(p, p.RecID)
	if err != nil {
		return err
	}
	return nil
}

// FixedAsset 固定资产
// 固定资产类型存放于 TypeTree
type FixedAsset struct {
	BasicModel
	FixedAssetTypeID          *int    `gorm:"type:int" json:"fixed_asset_type_id,omitempty" cn:"固定资产类型ID"`
	FixedAssetTypeName        *string `gorm:"type:varchar(50)" json:"fixed_asset_type_name,omitempty" cn:"固定资产类型名称"`
	DepreciationPeriod        int     `gorm:"type:int;not null" json:"depreciation_period" cn:"折旧期限（月）"`
	TotalPrice                float64 `gorm:"type:decimal(20,2);not null" json:"total_price" cn:"原价"`
	CurrentPrice              float64 `gorm:"type:decimal(20,2);not null" json:"current_price" cn:"现价"`
	MonthlyDepreciationAmount float64 `gorm:"type:decimal(20,2);not null" json:"monthly_depreciation_amount" cn:"每月折旧额"`
	Remark                    *string `gorm:"type:varchar(200)" json:"remark" cn:"备注"`
}

func (p FixedAsset) TableCnName() string {
	return "固定资产"
}
func (p FixedAsset) TableName() string {
	return "fixed_assets"

}
func (p FixedAsset) Mapping() map[string]interface{} {
	m := mapping{
		"settings": mapping{},
		"mappings": mapping{
			"properties": mapping{
				"rec_id": mapping{
					"type": "keyword",
				},
				"fixed_asset_name": mapping{
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

				"fixed_asset_amount": mapping{
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
func (p FixedAsset) ToESDoc() map[string]interface{} {
	return map[string]interface{}{
		"rec_id": p.RecID,
		"remark": p.Remark,
		//"payable_enterprise":p
	}
}
func (p *FixedAsset) AfterCreate(_ *gorm.DB) error {
	err := client.PutIntoIndex(p)
	if err != nil {
		return err
	}
	return nil
}

// AfterUpdate todo !!!
func (p *FixedAsset) AfterUpdate(_ *gorm.DB) error {
	err := client.UpdateIntoIndex(p, p.RecID,
		elastic.NewScriptInline("ctx._source.nickname=params.nickname;ctx._source.ancestral=params.ancestral").
			Params(p.ToESDoc()))
	if err != nil {
		return err
	}
	return nil
}
func (p *FixedAsset) AfterDelete(_ *gorm.DB) error {
	err := client.DeleteFromIndex(p, p.RecID)
	if err != nil {
		return err
	}
	return nil
}

// FixedAssetStatement 固定资产月度表
type FixedAssetStatement struct {
	BasicModel
}

func (m FixedAssetStatement) TableName() string {
	return "fixed_asset_statements"
}
