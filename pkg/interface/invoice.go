package _interface

import "github.com/XC-Zero/yinwan/pkg/model"

// Invoice 单据
type Invoice interface {
	// ExportInvoice 导出单据
	ExportInvoice() error
	// ImportInvoice 导入单据
	ImportInvoice() (error, inv Invoice)
	// TransferToCredential 转换为凭证
	TransferToCredential() (model.Credential, error)
}
