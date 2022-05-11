package mongo_model

import (
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
	FinanceCredentialCashierID *int `json:"finance_credential_cashier_id" bson:"finance_credential_cashier_id"`
	// 复核
	FinanceCredentialCheckerID *int `json:"finance_credential_checker_id" bson:"finance_credential_checker_id"`
	// 制单人
	FinanceCredentialMakerName *string `json:"finance_credential_maker_name" bson:"finance_credential_maker_name"`
	// 会计
	FinanceCredentialAccountantName *string `json:"finance_credential_accountant_name" bson:"finance_credential_accountant_name"`
	// 出纳
	FinanceCredentialCashierName *string `json:"finance_credential_cashier_name" bson:"finance_credential_cashier_name"`
	// 复核
	FinanceCredentialCheckerName *string `json:"finance_credential_checker_name" bson:"finance_credential_checker_name"`
	//	凭证备注
	FinanceCredentialRemark string `json:"finance_credential_remark" bson:"finance_credential_remark"`
}

func (c FinanceCredential) TableCnName() string {
	return "财务凭证"
}

func (c FinanceCredential) TableName() string {
	return "finance_credentials"
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
