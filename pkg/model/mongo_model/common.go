package mongo_model

import (
	"encoding/json"
	"log"
)

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
