package model

import (
	"gorm.io/gorm"
	"time"
)

// BasicModel 基本模型
type BasicModel struct {
	// 记录ID
	RecID *int `gorm:"primaryKey;type:int;autoIncrement" json:"id"`
	//创建时间
	CreatedAt time.Time `gorm:"type:timestamp;not null" json:"created_at"`
	// 更新时间
	UpdatedAt *time.Time `gorm:"type:timestamp" json:"updated_at"`
	// 删除时间
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}

// TimeOnlyModel 不含主键的基本模型
type TimeOnlyModel struct {
	CreatedAt time.Time      `gorm:"type:timestamp;not null" json:"created_at"`
	UpdatedAt *time.Time     `gorm:"type:timestamp" json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}
