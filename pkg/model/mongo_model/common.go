package mongo_model

import (
	"encoding/json"
	"github.com/XC-Zero/yinwan/pkg/model/mysql_model"
	"log"
)

type Analyzer string

const (
	IK_SMART    = "ik_smart"
	IK_MAX_WORD = "ik_max_word"
)

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
