package mysql_model

import (
	"github.com/XC-Zero/yinwan/pkg/client"
	"github.com/olivere/elastic/v7"
	"github.com/pkg/errors"
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

// FixedAsset 固定资产
// 固定资产类型存放于 TypeTree
type FixedAsset struct {
	BasicModel
	BookNameInfo
	FixedAssetName            string  `gorm:"type:varchar(100);not null'" form:"fixed_asset_name" json:"fixed_asset_name"`
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
		"rec_id":             p.RecID,
		"remark":             p.Remark,
		"created_at":         p.CreatedAt,
		"fixed_asset_amount": p.TotalPrice,
		"fixed_asset_name":   p.FixedAssetName,
	}
}
func (p *FixedAsset) AfterCreate(tx *gorm.DB) error {
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
func (p *FixedAsset) AfterUpdate(tx *gorm.DB) error {
	err := client.UpdateIntoIndex(p, p.RecID, tx.Statement.Context,
		elastic.NewScriptInline("ctx._source.nickname=params.nickname;ctx._source.ancestral=params.ancestral").
			Params(p.ToESDoc()))
	if err != nil {
		return err
	}
	return nil
}
func (p *FixedAsset) AfterDelete(tx *gorm.DB) error {
	err := client.DeleteFromIndex(p, p.RecID, tx.Statement.Context)
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
