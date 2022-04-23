package mongo_model

import (
	_const "github.com/XC-Zero/yinwan/pkg/const"
	_interface "github.com/XC-Zero/yinwan/pkg/interface"
	"github.com/XC-Zero/yinwan/pkg/model/common"
)

// FinanceCredential 财务凭证
type FinanceCredential struct {
	common.BasicModel
	// 凭证标题
	FinanceCredentialLabel string
	// 凭证责任人ID
	FinanceCredentialOwnerID int
	// 凭证责任人名称
	FinanceCredentialOwnerName string
	// 凭证分录
	FinanceCredentialEvents []FinanceCredentialEvent
	//	凭证备注
	FinanceCredentialRemark string
}

// FinanceCredentialEvent 凭证条目
type FinanceCredentialEvent struct {
	// 借
	IncreaseEvent []EventItem
	// 贷
	DecreaseEvent []EventItem
}

// EventItem 具体条目
type EventItem struct {
	// 变动类型
	EventItemType string
	// 变动对象
	EventItemObject string
	// 变动金额
	EventItemAmount string
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
