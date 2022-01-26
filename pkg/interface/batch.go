package _interface

import "time"

type Batch interface {
	GetBatchTime() time.Time
	GetBatchOwner() Staff
}
