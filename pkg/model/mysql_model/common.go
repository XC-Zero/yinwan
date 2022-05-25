package mysql_model

import (
	"encoding/json"
	"gorm.io/gorm"
	"log"
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

// BookNameInfo 给ES记录账套用的字段,并不存在mysql里,特意加上了
type BookNameInfo struct {
	BookNameID string `gorm:"-" sql:"-" json:"book_name_id" `
	BookName   string `gorm:"-" sql:"-" json:"book_name" `
}

type Analyzer string

const (
	IK_SMART    = "ik_smart"
	IK_MAX_WORD = "ik_max_word"
)

type mapping map[string]interface{}

func (m mapping) ToString() string {
	marshal, err := json.Marshal(m)
	if err != nil {
		return ""
	}
	res := string(marshal)
	log.Println(res)
	return res
}
