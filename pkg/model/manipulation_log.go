package model

import (
	"time"
)

type ManipulationLog struct {
	BasicModel
	ManipulatorID       string `gorm:"type:int"`
	ManipulatorName     string `gorm:"type:varchar(50)"`
	ManipulationContent string `gorm:"type:varchar(500)"`
	ManipulationTime    time.Time
	ManipulationRemark  string `gorm:"type:varchar(500)"`
}
