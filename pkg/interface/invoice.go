package _interface

import _const "github.com/XC-Zero/yinwan/pkg/const"

// Invoice 单据
type Invoice interface {

	// ImportInvoice 导入单据
	ImportInvoice(invoiceType _const.InvoiceType) (error, inv Invoice)

	// ExportInvoice 导出单据
	ExportInvoice() error

	// TransferToCredential 单据转凭证
	// 根据类型去找对应的转换模板
	//TransferToCredential(_const.CredentialType) Credential
}
