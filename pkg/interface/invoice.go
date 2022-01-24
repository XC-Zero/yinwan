package _interface

// Invoice 单据
type Invoice interface {
	ExportDocument() error
	PrintDocument() error
	ImportDocument(inv Invoice) error
}
