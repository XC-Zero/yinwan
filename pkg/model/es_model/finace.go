package es_model

type Payable struct {
	RecID                    int     `json:"rec_id"`
	PayableAmount            float64 `json:"payable_amount"`
	PayableEnterprise        string  `json:"payable_enterprise"`
	PayableEnterpriseAddress string  `json:"payable_enterprise_address"`
	PayableContact           string  `json:"payable_contact"`
	Remark                   string  `json:"remark"`
	CreateAt                 string  `json:"create_at"`
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
				"payable_enterprise": mapping{
					"type":            "text",
					"analyzer":        IK_SMART,
					"search_analyzer": IK_SMART,
					"fields": mapping{
						"keyword": mapping{
							"type":         "keyword",
							"ignore_above": 256,
						},
					},
				},
				"payable_enterprise_address": mapping{
					"type":            "text",
					"analyzer":        IK_SMART,
					"search_analyzer": IK_SMART,
				},
				"payable_contact": mapping{
					"type":            "text",
					"analyzer":        IK_SMART,
					"search_analyzer": IK_SMART,
				},
				"payable_amount": mapping{
					"type":            "text",
					"analyzer":        IK_SMART,
					"search_analyzer": IK_SMART,
				},
				"remark": mapping{
					"type":            "text",
					"analyzer":        IK_SMART,
					"search_analyzer": IK_SMART,
				},
				"create_at": mapping{
					"type": "text",
				},
			},
		},
	}
	return m
}

type Receivable struct {
	RecID                       int     `json:"rec_id"`
	ReceivableAmount            float64 `json:"receivable_amount"`
	ReceivableEnterprise        string  `json:"receivable_enterprise"`
	ReceivableEnterpriseAddress string  `json:"receivable_enterprise_address"`
	ReceivableContact           string  `json:"receivable_contact"`
	Remark                      string  `json:"remark"`
	CreateAt                    string  `json:"create_at"`
}

func (p Receivable) TableCnName() string {
	return "应收"
}
func (p Receivable) TableName() string {
	return "receivables"

}
func (p Receivable) Mapping() map[string]interface{} {
	m := mapping{
		"settings": mapping{},
		"mappings": mapping{
			"properties": mapping{
				"rec_id": mapping{
					"type": "integer",
				},
				"receivable_enterprise": mapping{
					"type":            "text",
					"analyzer":        IK_SMART,
					"search_analyzer": IK_SMART,
					"fields": mapping{
						"keyword": mapping{
							"type":         "keyword",
							"ignore_above": 256,
						},
					},
				},
				"receivable_enterprise_address": mapping{
					"type":            "text",
					"analyzer":        IK_SMART,
					"search_analyzer": IK_SMART,
				},
				"receivable_contact": mapping{
					"type":            "text",
					"analyzer":        IK_SMART,
					"search_analyzer": IK_SMART,
				},
				"receivable_amount": mapping{
					"type":            "text",
					"analyzer":        IK_SMART,
					"search_analyzer": IK_SMART,
				},
				"remark": mapping{
					"type":            "text",
					"analyzer":        IK_SMART,
					"search_analyzer": IK_SMART,
				},
				"create_at": mapping{
					"type": "text",
				},
			},
		},
	}
	return m
}

type FixedAsset struct {
	RecID            int     `json:"rec_id"`
	FixedAssetName   string  `json:"fixed_asset_name"`
	FixedAssetAmount float64 `json:"fixed_asset_amount"`
	Remark           string  `json:"remark"`
	CreateAt         string  `json:"create_at"`
}

func (p FixedAsset) TableCnName() string {
	return "固定资产"
}
func (p FixedAsset) TableName() string {
	return "fixed_assets"

}
func (p FixedAsset) Mapping() map[string]interface{} {
	m := mapping{
		"settings": mapping{},
		"mappings": mapping{
			"properties": mapping{
				"rec_id": mapping{
					"type": "integer",
				},
				"fixed_asset_name": mapping{
					"type":            "text",
					"analyzer":        IK_SMART,
					"search_analyzer": IK_SMART,
					"fields": mapping{
						"keyword": mapping{
							"type":         "keyword",
							"ignore_above": 256,
						},
					},
				},

				"fixed_asset_amount": mapping{
					"type":            "text",
					"analyzer":        IK_SMART,
					"search_analyzer": IK_SMART,
				},
				"remark": mapping{
					"type":            "text",
					"analyzer":        IK_SMART,
					"search_analyzer": IK_SMART,
				},
				"create_at": mapping{
					"type": "text",
				},
			},
		},
	}
	return m
}