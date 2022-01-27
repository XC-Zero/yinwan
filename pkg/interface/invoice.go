package _interface

// Invoice 单据
type Invoice interface {
	ExportInvoice() error
	PrintInvoice() error
	ImportInvoice() (error, inv Invoice)
}
