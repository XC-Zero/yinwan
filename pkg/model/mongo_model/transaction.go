package mongo_model

import (
	"encoding/json"
	"github.com/XC-Zero/yinwan/pkg/model/common"
	"time"
)

type PayType int

//goland:noinspection GoSnakeCaseUsage
const (
	WX PayType = iota + 35001
	ZFB
	BANK_CARD
	CASH
	OTHER
)

func (p PayType) DisPlay() string {
	switch p {
	case ZFB:
		return "支付宝"
	case CASH:
		return "现金"
	case WX:
		return "微信"
	case BANK_CARD:
		return "银行卡"
	case OTHER:
		return "其他"
	default:
		return "未知"
	}
}

type Transaction struct {
	common.BasicModel
	TransactionContent string    `gorm:"type:varchar(500)"`
	TransactionAmount  float64   `gorm:"type:decimal(20,4);not null;"`
	PayType            *PayType  `gorm:"type:int"`
	PayerID            *int      `gorm:"type:int"`
	PayerName          string    `gorm:"type:varchar(50)"`
	PayeeID            *int      `gorm:"type:int"`
	PayeeName          string    `gorm:"type:varchar(50)"`
	TransactionTime    time.Time `gorm:"not null;"`
	TransactionRemark  *string   `gorm:"type:varchar(200)"`
}

func (c Transaction) TableCnName() string {
	return "交易"
}
func (c Transaction) TableName() string {
	return "transactions"
}

func (t *Transaction) SetTransactionContent(m map[string]interface{}) {
	marshal, _ := json.Marshal(m)
	t.TransactionContent = string(marshal)
}

func (t Transaction) TransactionContentMap() (map[string]interface{}, error) {
	m := make(map[string]interface{}, 0)
	err := json.Unmarshal([]byte(t.TransactionContent), &m)
	if err != nil {
		return nil, err
	}
	return m, nil
}
