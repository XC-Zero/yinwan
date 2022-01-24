package _interface

// Document 单据
type Document interface {
	ExportDocument() error
	PrintDocument() error
	ImportDocument(doc Document) error
}
