package _const

type Tag string

const (
	CN Tag = "cn"
)

func (t Tag) Display() string {
	switch t {
	case CN:
		return "中文备注"
	default:
		return "未知"
	}
}
