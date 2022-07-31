package _const

type ExcelStatus int

const (
	EXPORTING ExcelStatus = iota + 1
	DONE
	EXPIRED
)
