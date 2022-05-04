package _interface

type Excel interface {
	Import(excel Excel) error
	Export() (Excel, error)
}
