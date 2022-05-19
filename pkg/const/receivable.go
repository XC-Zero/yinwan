package _const

type ReceivableStatus int

const (
	Finished ReceivableStatus = iota + 56001
	Unfinished
	Partial
)

func (r ReceivableStatus) Display() string {
	switch r {
	case Finished:
		return "已完成"
	case Unfinished:
		return "未完成"
	case Partial:
		return "部分完成"
	default:
		return "未知"
	}
}
