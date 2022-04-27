package es_model

type Material struct {
	RecID        int    `json:"rec_id"`
	MaterialName string `json:"material_name"`
	Remark       string `json:"remark"`
	CreateAt     string `json:"create_at"`
}

func (m Material) TableName() string {
	return "materials"
}
func (m Material) TableCnName() string {
	return "原材料"
}
func (m Material) Mapping() map[string]interface{} {
	ma := mapping{
		"settings": mapping{},
		"mappings": mapping{
			"properties": mapping{
				"rec_id": mapping{
					"type": "keyword",
				},
				"material_name": mapping{
					"type":            "text",   //字符串类型且进行分词, 允许模糊匹配
					"analyzer":        IK_SMART, //设置分词工具
					"search_analyzer": IK_SMART,
					"fields": mapping{ //当需要对模糊匹配的字符串也允许进行精确匹配时假如此配置
						"keyword": mapping{
							"type":         "keyword",
							"ignore_above": 256,
						},
					},
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
	return ma
}

type Commodity struct {
	RecID         int    `json:"rec_id"`
	CommodityName string `json:"commodity_name"`
	Remark        string `json:"remark"`
	CreateAt      string `json:"create_at"`
}

func (m Commodity) TableName() string {
	return "commodity"
}
func (m Commodity) TableCnName() string {
	return "产品"
}
func (m Commodity) Mapping() map[string]interface{} {
	ma := mapping{
		"settings": mapping{},
		"mappings": mapping{
			"properties": mapping{
				"rec_id": mapping{
					"type": "integer",
				},
				"commodity_name": mapping{
					"type":            "text",   //字符串类型且进行分词, 允许模糊匹配
					"analyzer":        IK_SMART, //设置分词工具
					"search_analyzer": IK_SMART,
					"fields": mapping{ //当需要对模糊匹配的字符串也允许进行精确匹配时假如此配置
						"keyword": mapping{
							"type":         "keyword",
							"ignore_above": 256,
						},
					},
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
	return ma
}

type StockOutRecord struct {
	RecID           int    `json:"rec_id"`
	StockOutContent string `json:"stock_out_content"`
	StockOutOwner   string `json:"stock_out_owner"`
	Remark          string `json:"remark"`
	CreateAt        string `json:"create_at"`
}

func (m StockOutRecord) TableName() string {
	return "stock_out_records"
}
func (m StockOutRecord) TableCnName() string {
	return "出库记录"
}
func (m StockOutRecord) Mapping() map[string]interface{} {
	ma := mapping{
		"settings": mapping{},
		"mappings": mapping{
			"properties": mapping{
				"rec_id": mapping{
					"type": "integer",
				},
				"stock_out_content": mapping{
					"type":            "text",   //字符串类型且进行分词, 允许模糊匹配
					"analyzer":        IK_SMART, //设置分词工具
					"search_analyzer": IK_SMART,
				},
				"stock_out_owner": mapping{
					"type":            "text",   //字符串类型且进行分词, 允许模糊匹配
					"analyzer":        IK_SMART, //设置分词工具
					"search_analyzer": IK_SMART,
					"fields": mapping{ //当需要对模糊匹配的字符串也允许进行精确匹配时假如此配置
						"keyword": mapping{
							"type":         "keyword",
							"ignore_above": 256,
						},
					},
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
	return ma
}

type StockInRecord struct {
	RecID           int    `json:"rec_id"`
	StockOutContent string `json:"stock_in_content"`
	StockOutOwner   string `json:"stock_in_owner"`
	Remark          string `json:"remark"`
	CreateAt        string `json:"create_at"`
}

func (m StockInRecord) TableName() string {
	return "stock_in_records"
}
func (m StockInRecord) TableCnName() string {
	return "入库记录"
}
func (m StockInRecord) Mapping() map[string]interface{} {
	ma := mapping{
		"settings": mapping{},
		"mappings": mapping{
			"properties": mapping{
				"rec_id": mapping{
					"type": "integer",
				},
				"stock_in_content": mapping{
					"type":            "text",   //字符串类型且进行分词, 允许模糊匹配
					"analyzer":        IK_SMART, //设置分词工具
					"search_analyzer": IK_SMART,
				},
				"stock_in_owner": mapping{
					"type":            "text",   //字符串类型且进行分词, 允许模糊匹配
					"analyzer":        IK_SMART, //设置分词工具
					"search_analyzer": IK_SMART,
					"fields": mapping{ //当需要对模糊匹配的字符串也允许进行精确匹配时假如此配置
						"keyword": mapping{
							"type":         "keyword",
							"ignore_above": 256,
						},
					},
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
	return ma
}
