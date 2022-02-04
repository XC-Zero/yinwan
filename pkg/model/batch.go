package model

import (
	_interface "github.com/XC-Zero/yinwan/pkg/interface"
	"time"
)

type Batch struct {
	BasicModel
	BatchTime  time.Time
	BatchName  string
	BatchOwner _interface.Staff
}
