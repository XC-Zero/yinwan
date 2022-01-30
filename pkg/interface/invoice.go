package _interface

// Invoice 单据
type Invoice interface {
	// ExportInvoice 导出单据
	ExportInvoice() error
	// ImportInvoice 导入单据
	ImportInvoice() (error, inv Invoice)
}
