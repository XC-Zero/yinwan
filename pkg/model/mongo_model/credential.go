package mongo_model

import (
	"fmt"
	_const "github.com/XC-Zero/yinwan/pkg/const"
	_interface "github.com/XC-Zero/yinwan/pkg/interface"
	"github.com/XC-Zero/yinwan/pkg/model/mysql_model"
)

// FinanceCredential 财务凭证
type FinanceCredential struct {
	mysql_model.BasicModel
	BookNameInfo
	// 凭证标题
	FinanceCredentialLabel string `json:"finance_credential_label" bson:"finance_credential_label"`
	// 凭证责任人ID
	FinanceCredentialOwnerID int `json:"finance_credential_owner_id" bson:"finance_credential_owner_id"`
	// 凭证责任人名称
	FinanceCredentialOwnerName string `json:"finance_credential_owner_name" bson:"finance_credential_owner_name"`
	// 凭证分录
	FinanceCredentialEvents []FinanceCredentialEvent `json:"finance_credential_events" bson:"finance_credential_events"`
	// 制单人
	FinanceCredentialMakerID *int `json:"finance_credential_maker_id" bson:"finance_credential_maker_id"`
	// 会计
	FinanceCredentialAccountantID *int `json:"finance_credential_accountant_id" bson:"finance_credential_accountant_id"`
	// 出纳
	FinanceCredentialCashierID *int `json:"finance_credential_cashier_id,omitempty" bson:"finance_credential_cashier_id"`
	// 复核
	FinanceCredentialCheckerID *int `json:"finance_credential_checker_id,omitempty" bson:"finance_credential_checker_id"`
	// 制单人
	FinanceCredentialMakerName *string `json:"finance_credential_maker_name,omitempty" bson:"finance_credential_maker_name"`
	// 会计
	FinanceCredentialAccountantName *string `json:"finance_credential_accountant_name,omitempty" bson:"finance_credential_accountant_name"`
	// 出纳
	FinanceCredentialCashierName *string `json:"finance_credential_cashier_name,omitempty" bson:"finance_credential_cashier_name"`
	// 复核
	FinanceCredentialCheckerName *string `json:"finance_credential_checker_name,omitempty" bson:"finance_credential_checker_name"`
	//	凭证备注
	Remark *string `json:"remark,omitempty" bson:"remark"`
}

func (c FinanceCredential) TableCnName() string {
	return "财务凭证"
}
func (c FinanceCredential) TableName() string {
	return "finance_credentials"
}
func (c FinanceCredential) Mapping() map[string]interface{} {
	ma := mapping{
		"settings": mapping{},
		"mappings": mapping{
			"properties": mapping{
				"rec_id": mapping{
					"type": "keyword",
				},
				"related_person": mapping{
					"type":            "text",
					"analyzer":        IK_SMART,
					"search_analyzer": IK_SMART,
				},
				"finance_credential_content": mapping{
					"type":            "text",
					"analyzer":        IK_SMART,
					"search_analyzer": IK_SMART,
				},
				"stock_in_record_type": mapping{
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
	return ma
}

// ToESDoc todo !!!!
func (c FinanceCredential) ToESDoc() map[string]interface{} {

	maker, account, cashier, checker := "", "", "", ""
	if c.FinanceCredentialMakerName != nil {
		maker = *c.FinanceCredentialMakerName
	}
	if c.FinanceCredentialCashierName != nil {
		cashier = *c.FinanceCredentialCashierName
	}
	if c.FinanceCredentialAccountantName != nil {
		account = *c.FinanceCredentialAccountantName
	}

	if c.FinanceCredentialCheckerName != nil {
		checker = *c.FinanceCredentialCheckerName
	}
	var credentialContent string
	for _, event := range c.FinanceCredentialEvents {
		for _, decrease := range event.DecreaseEvent {
			credentialContent += fmt.Sprintf(
				"贷: 变动类型 :%s 变动对象: %s 变动金额:%s \n",
				decrease.EventItemType,
				decrease.EventItemObject,
				decrease.EventItemAmount)
		}

		for _, increase := range event.IncreaseEvent {
			credentialContent += fmt.Sprintf(
				"借: 变动类型 :%s 变动对象: %s 变动金额:%s \n",
				increase.EventItemType,
				increase.EventItemObject,
				increase.EventItemAmount)
		}
	}

	return map[string]interface{}{
		"rec_id":                     c.RecID,
		"created_at":                 c.CreatedAt,
		"remark":                     c.Remark,
		"related_person":             fmt.Sprintf("制单人:%s  会计:%s  出纳:%s  复核:%s ", maker, account, cashier, checker),
		"finance_credential_content": credentialContent,
		//"stock_in_owner":             m.StockInRecordOwnerName,
		//"book_name":                  m.BookName,
		//"book_name_id":               m.BookNameID,
	}
}

// FinanceCredentialEvent 凭证条目
type FinanceCredentialEvent struct {
	// 借
	IncreaseEvent []EventItem `json:"increase_event" bson:"increase_event"`
	// 贷
	DecreaseEvent []EventItem `json:"decrease_event" bson:"decrease_event"`
}

// EventItem 具体条目
type EventItem struct {
	// 变动类型
	EventItemType string `json:"event_item_type" bson:"event_item_type"`
	// 变动对象
	EventItemObject string `json:"event_item_object" bson:"event_item_object"`
	// 变动金额
	EventItemAmount string `json:"event_item_amount" bson:"event_item_amount"`
}

// CalculateTotalAmount todo 通过凭证中各条目计算总金额
// todo 需不需要，有待商榷
func (c *FinanceCredential) CalculateTotalAmount() {

}

// CreateFinanceCredential 创建财务凭证
//  todo
func CreateFinanceCredential() {

}

// TransferToInvoice 转为单据
// todo
func (c *FinanceCredential) TransferToInvoice(invoiceType _const.InvoiceType) _interface.Invoice {
	return nil
}
