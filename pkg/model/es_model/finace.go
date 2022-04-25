package es_model

type Payable struct {
}

func (p Payable) TableCnName() string {
	return "应付"
}

func (p Payable) TableName() string {
	return "payables"

}
func (p Payable) Mapping() map[string]interface{} {
	m := mapping{
		"settings": mapping{},
		"mappings": mapping{
			"properties": mapping{
				"rec_id": mapping{
					"type": "integer",
				},
			},
		},
	}
	return m
}
