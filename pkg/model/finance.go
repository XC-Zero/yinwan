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
