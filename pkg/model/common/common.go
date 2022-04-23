package common

import (
	"gorm.io/gorm"
	"time"
)

// BasicModel 基本模型
type BasicModel struct {
	RecID     *int           `gorm:"primaryKey;type:int;autoIncrement" json:"rec_id,omitempty" bson:"rec_id" cn:"记录ID"`
	CreatedAt time.Time      `gorm:"type:timestamp;not null" json:"created_at" bson:"created_at" cn:"创建时间"`
	UpdatedAt *time.Time     `gorm:"type:timestamp" json:"updated_at" bson:"updated_at" cn:"更新时间"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at,omitempty" bson:"deleted_at" cn:"删除时间"`
}

// TimeOnlyModel 不含主键的基本模型
type TimeOnlyModel struct {
	CreatedAt time.Time      `gorm:"type:timestamp;not null" bson:"created_at" json:"created_at" cn:"创建时间"`
	UpdatedAt *time.Time     `gorm:"type:timestamp" bson:"updated_at" json:"updated_at" cn:"更新时间"`
	DeletedAt gorm.DeletedAt `gorm:"index" bson:"deleted_at" json:"deleted_at" cn:"删除时间"`
}
