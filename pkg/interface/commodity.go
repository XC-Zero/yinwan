package _interface

// Commodity 货品
type Commodity interface {
	// StockIn 入库
	StockIn() error
	// StockOut 出库
	StockOut() error
	// FractionalPrice 价格以形式分数
	FractionalPrice(price float64)
	// GetBatch 获取批次信息
	GetBatch() Batch
	// SetBatch 设置批次信息
	SetBatch(batch Batch)
}
