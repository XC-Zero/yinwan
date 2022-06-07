package mongo_model

import (
	"encoding/json"
	"github.com/XC-Zero/yinwan/pkg/model/mysql_model"
	"log"
	"time"
)

type Analyzer string

const (
	IK_SMART    = "ik_smart"
	IK_MAX_WORD = "ik_max_word"
)

type BasicModel struct {
	RecID     *int       ` json:"rec_id,omitempty" bson:"rec_id" cn:"记录ID"`
	CreatedAt time.Time  ` json:"created_at" bson:"created_at" cn:"创建时间"`
	UpdatedAt *time.Time ` json:"updated_at,omitempty" bson:"updated_at" cn:"更新时间"`
	DeletedAt *time.Time ` json:"deleted_at,omitempty" bson:"deleted_at" cn:"删除时间"`
}

type mapping map[string]interface{}
type BookNameInfo mysql_model.BookNameInfo

func (m mapping) ToString() string {
	marshal, err := json.Marshal(m)
	if err != nil {
		return ""
	}
	res := string(marshal)
	log.Println(res)
	return res
}
