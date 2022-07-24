package _const

type InvoiceType string

var InvoiceTypeList = []map[string]interface{}{
	{
		"invoice_type": PURCHASE_INVOICE,
		"label":        PURCHASE_INVOICE.Display(),
	}, {
		"invoice_type": TRANSACTION_INVOICE,
		"label":        TRANSACTION_INVOICE.Display(),
	}, {
		"invoice_type": RETURN_INVOICE,
		"label":        RETURN_INVOICE.Display(),
	}, {
		"invoice_type": PAYABLE_INVOICE,
		"label":        PAYABLE_INVOICE.Display(),
	}, {
		"invoice_type": RECIVEABLE_INVOICE,
		"label":        RECIVEABLE_INVOICE.Display(),
	}, {
		"invoice_type": ASSEMBLE_INVOICE,
		"label":        ASSEMBLE_INVOICE.Display(),
	},
}

var (
	PURCHASE_INVOICE    = InvoiceType("purchases")
	TRANSACTION_INVOICE = InvoiceType("transactions")
	RETURN_INVOICE      = InvoiceType("returns")
	PAYABLE_INVOICE     = InvoiceType("payables")
	RECIVEABLE_INVOICE  = InvoiceType("reciveables")
	ASSEMBLE_INVOICE    = InvoiceType("assembles")
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
