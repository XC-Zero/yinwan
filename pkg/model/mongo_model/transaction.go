package mongo_model

import "github.com/XC-Zero/yinwan/pkg/model/mysql_model"

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

// Transaction 销售
type Transaction struct {
	mysql_model.BasicModel  `bson:"inline"`
	BookNameInfo            `bson:"-"`
	TransactionContent      []map[string]interface{} `json:"transaction_content" form:"transaction_content" bson:"transaction_content" cn:"销售详情"`
	TransactionAmount       string                   `json:"transaction_amount" form:"transaction_amount" bson:"transaction_amount" cn:"销售金额"`
	TransactionActualAmount string                   `json:"transaction_actual_amount" form:"transaction_actual_amount" bson:"transaction_actual_amount" cn:"实际销售金额"`
	PayType                 *PayType                 `json:"pay_type,omitempty" form:"pay_type,omitempty" bson:"pay_type" cn:"支付类型"`
	PayerName               *string                  `json:"payer_name,omitempty" form:"payer_name" bson:"payer_name" cn:"支付人姓名"`
	TransactionOwnerID      *int                     `json:"transaction_owner_id,omitempty" form:"transaction_owner_id" bson:"transaction_owner_id" cn:"销售管理员ID"`
	TransactionOwnerName    *string                  `json:"transaction_owner_name,omitempty" form:"transaction_owner_name" bson:"transaction_owner_name" cn:"销售管理员姓名"`
	TransactionTime         *string                  `json:"transaction_time,omitempty" form:"transaction_time" bson:"transaction_time" cn:"交易时间"`
	ReceiveID               *int                     `json:"receive_id,omitempty" form:"receive_id" bson:"receive_id" cn:"应收记录"`
	Remark                  *string                  `json:"remark,omitempty" form:"remark" bson:"remark" cn:"备注"`
}

func (t Transaction) Mapping() map[string]interface{} {
	return mapping{
		"settings": mapping{},
		"mappings": mapping{
			"properties": mapping{
				"rec_id": mapping{
					"type": "keyword",
				},
				"transaction_content": mapping{
					"type":            "text",
					"analyzer":        IK_SMART,
					"search_analyzer": IK_SMART,
				},
				"transaction_owner_name": mapping{
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
				"transaction_amount": mapping{
					"type": "keyword",
				},
				"receive_id": mapping{
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
}

// ToESDoc todo!!!
func (t Transaction) ToESDoc() map[string]interface{} {
	return map[string]interface{}{
		"rec_id":                 t.RecID,
		"remark":                 t.Remark,
		"created_at":             t.CreatedAt,
		"receive_id":             t.ReceiveID,
		"transaction_amount":     t.TransactionAmount,
		"transaction_owner_name": t.TransactionOwnerName,
		//"transaction_content":t.TransactionContent,
		"book_name":    t.BookName,
		"book_name_id": t.BookNameID,
	}
}

func (t Transaction) TableCnName() string {
	return "交易"
}
func (t Transaction) TableName() string {
	return "transactions"
}
