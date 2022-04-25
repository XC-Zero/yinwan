package mysql_model

import (
	"github.com/XC-Zero/yinwan/pkg/model/mysql_model/common"
	"time"
)

// ManipulationLog 操作日志
type ManipulationLog struct {
	common.BasicModel
	ManipulatorID       string `gorm:"type:int"`
	ManipulatorName     string `gorm:"type:varchar(50)"`
	ManipulationContent string `gorm:"type:varchar(500)"`
	ManipulationTime    time.Time
	ManipulationRemark  string `gorm:"type:varchar(500)"`
}

func (c ManipulationLog) TableCnName() string {
	return "操作日志"
}
func (c ManipulationLog) TableName() string {
	return "manipulation_logs"
}
