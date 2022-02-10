package _interface

import _const "github.com/XC-Zero/yinwan/pkg/const"

type Credential interface {
	TransferToInvoice(invoiceType _const.InvoiceType) Invoice
}
