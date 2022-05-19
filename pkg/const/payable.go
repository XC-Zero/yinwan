package _const

type PayableStatus int

const (
	Paid PayableStatus = iota + 57001
	UnPaid
	PartialPaid
)

func (r PayableStatus) Display() string {
	switch r {
	case Paid:
		return "已支付"
	case UnPaid:
		return "未支付"
	case PartialPaid:
		return "部分支付"
	default:
		return "未知"
	}
}
