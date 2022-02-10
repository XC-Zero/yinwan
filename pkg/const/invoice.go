package _const

type InvoiceType int

func (i *InvoiceType) Display() string {
	switch i {
	default:
		return "其他"
	}
}
