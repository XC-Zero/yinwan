package _interface

import "time"

type Batch interface {
	BatchTime() time.Time
	BatchName() string
	BatchOwner() Staff
}
