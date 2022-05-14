package _const

// StockInType 入库类型
type StockInType int

var StockInTypeList = []map[string]interface{}{
	{
		"stock_in_type": INVENTORY_SURPLUS.Display(),
	}, {
		"stock_in_type": DISMANTLE.Display(),
	}, {
		"stock_in_type": PURCHASE.Display(),
	}, {
		"stock_in_type": DISPATCH.Display(),
	}, {
		"stock_in_type": "未知",
	},
}

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
