package _interface

// Commodity 货品
type Commodity interface {
	// StockIn 入库
	StockIn() error
	// StockOut 出库
	StockOut() error

	// FractionalPrice 价格以形式分数
	FractionalPrice()
}
