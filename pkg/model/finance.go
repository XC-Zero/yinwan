package model

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
	CredentialID *int `gorm:"type:int" json:"credential_id,omitempty"`
}

// Receivable 应收
type Receivable struct {
	BasicModel
	//	关联凭证ID
	CredentialID *int `gorm:"type:int" json:"credential_id,omitempty"`
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

// FixedAssetStatement 固定资产月度表
type FixedAssetStatement struct {
	BasicModel
}
