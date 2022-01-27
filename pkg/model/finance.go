package model

type Currency struct {
	BasicModel
	CurrencyName string `gorm:"type:varchar(50)"`
}
