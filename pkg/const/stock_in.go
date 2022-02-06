package _const

// StockInType 入库类型
type StockInType int

//goland:noinspection GoSnakeCaseUsage
const (
	INVENTORY_SURPLUS StockInType = iota + 760001
	DISMANTLE
	PURCHASE
	DISPATCH
)

func (s StockInType) Display() string {
	switch s {
	case INVENTORY_SURPLUS:
		return "盘盈入库"
	case DISMANTLE:
		return "拆解入库"
	case PURCHASE:
		return "采购入库"
	case DISPATCH:
		return "调拨入库"

	default:
		return "其他"
	}
}
