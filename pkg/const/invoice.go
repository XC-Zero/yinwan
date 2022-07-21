package _const

import (
	"github.com/XC-Zero/yinwan/pkg/model/mongo_model"
	"github.com/XC-Zero/yinwan/pkg/model/mysql_model"
)

type InvoiceType string

var (
	PURCHASE_INVOICE    = InvoiceType(mongo_model.Purchase{}.TableName())
	TRANSACTION_INVOICE = InvoiceType(mongo_model.Transaction{}.TableName())
	RETURN_INVOICE      = InvoiceType(mongo_model.Return{}.TableName())
	PAYABLE_INVOICE     = InvoiceType(mysql_model.Payable{}.TableName())
	RECIVEABLE_INVOICE  = InvoiceType(mysql_model.Receivable{}.TableName())
	ASSEMBLE_INVOICE    = InvoiceType(mongo_model.Assemble{}.TableName())
	//PURCHASE_INVOICE    = mongo_model.Purchase{}.TableName()
	//PURCHASE_INVOICE    = mongo_model.Purchase{}.TableName()
)

func (i InvoiceType) Display() string {
	switch i {
	case PURCHASE_INVOICE:
		return "采购"
	case TRANSACTION_INVOICE:
		return "销售"
	case RETURN_INVOICE:
		return "退货"
	case PAYABLE_INVOICE:
		return "应付"
	case RECIVEABLE_INVOICE:
		return "应收"
	case ASSEMBLE_INVOICE:
		return "组装拆卸"
	default:
		return "其他"
	}
}

func TransToInvoiceType(tableName string) InvoiceType {
	switch tableName {
	case "purchases":
		return PURCHASE_INVOICE
	case "transactions":
		return TRANSACTION_INVOICE
	case "returns":
		return RETURN_INVOICE
	case "payables":
		return PAYABLE_INVOICE
	case "reciveables":
		return RECIVEABLE_INVOICE
	case "assembles":
		return ASSEMBLE_INVOICE
	default:
		return "其他"
	}
}
