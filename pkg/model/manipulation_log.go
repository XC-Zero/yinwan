package model

import (
	"time"
)

type ManipulationLog struct {
	LogID               string `gorm:"primaryKey;type:int;autoIncrement"json:"id"`
	ManipulatorID       string `gorm:"type:int"`
	ManipulatorName     string `gorm:"type:varchar(50)"`
	ManipulationContent string `gorm:"type:varchar(500)"`
	ManipulationTime    time.Time
	ManipulationRemark  string `gorm:"type:varchar(500)"`
}
