package es_model

import (
	"encoding/json"
)

type mapping map[string]interface{}

func (m mapping) ToString() string {
	marshal, err := json.Marshal(m)
	if err != nil {
		return ""
	}
	return string(marshal)
}

type Payable struct {
}

func (p Payable) TableCnName() string {
	return "应付"
}

func (p Payable) TableName() string {
	return "payables"

}
func (p Payable) Mapping() string {
	m := mapping{
		"": mapping{},
	}
	return m.ToString()
}
