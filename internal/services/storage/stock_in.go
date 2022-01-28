package storage

type StockInType int

const (
	INVENTORY_SURPLUS StockInType = iota + 760001
	DISMANTLE
	PURCHASE
)

func (s StockInType) Display() string {
	switch s {
	case INVENTORY_SURPLUS:
		return "盘盈入库"
	case DISMANTLE:
		return "拆解入库"
	case PURCHASE:
		return "采购入库"
	default:
		return "未知"
	}
}
