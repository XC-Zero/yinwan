package _const

type StockOutType int

const (
	LOSS StockOutType = iota + 770001
	SALE
	ASSEMBLE
)

func (s StockOutType) Display() string {
	switch s {
	case LOSS:
		return "损耗出库"
	case ASSEMBLE:
		return "组装出库"
	case SALE:
		return "销售出库"
	default:
		return "其他"
	}
}
