package _interface

import "time"

// Batch 批次
type Batch interface {
	// BatchTime 批次时间
	BatchTime() time.Time
	// BatchName 批次名
	BatchName() string
	// BatchOwner 批次责任人
	BatchOwner() Staff
}
