package _interface

import "github.com/XC-Zero/yinwan/pkg/model"

// Invoice 单据
type Invoice interface {
	ExportInvoice() error
	ImportInvoice() (error, inv Invoice)
	TransferToCredential() (model.Credential, error)
}
