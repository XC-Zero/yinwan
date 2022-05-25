package mysql_model

import (
	"github.com/XC-Zero/yinwan/pkg/client"
	"github.com/XC-Zero/yinwan/pkg/utils/es_tool"
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
	CurrencyExchangeRate string `gorm:"type:varchar(20)"`
}

// FixedAsset 固定资产
// 固定资产类型存放于 TypeTree
type FixedAsset struct {
	BasicModel
	BookNameInfo
	FixedAssetPicUrl          *string `gorm:"type:varchar(500)" json:"fixed_asset_pic_url,omitempty" form:"fixed_asset_pic_url,omitempty" `
	FixedAssetName            string  `gorm:"type:varchar(100);not null'" form:"fixed_asset_name" json:"fixed_asset_name"`
	FixedAssetTypeID          *int    `gorm:"type:int;index" json:"fixed_asset_type_id,omitempty"  form:"fixed_asset_type_id" cn:"固定资产类型ID"`
	FixedAssetTypeName        *string `gorm:"type:varchar(50)" json:"fixed_asset_type_name,omitempty"  form:"fixed_asset_type_name" cn:"固定资产类型名称"`
	DepreciationPeriod        int     `gorm:"type:int;not null" json:"depreciation_period"  form:"depreciation_period" cn:"折旧期限（月）"`
	TotalPrice                *string `gorm:"type:varchar(20);not null" json:"total_price,omitempty"  form:"total_price" cn:"原价"`
	CurrentPrice              *string `gorm:"type:varchar(20);not null" json:"current_price,omitempty"  form:"current_price" cn:"残值"`
	MonthlyDepreciationAmount *string `gorm:"type:varchar(20);not null" json:"monthly_depreciation_amount,omitempty"  form:"monthly_depreciation_amount" cn:"每月折旧额"`
	Remark                    *string `gorm:"type:varchar(200)" json:"remark"  form:"remark" cn:"备注"`
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
		"book_name":          p.BookName,
		"book_name_id":       p.BookNameID,
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
