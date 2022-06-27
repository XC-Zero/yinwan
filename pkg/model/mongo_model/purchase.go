package mongo_model

import "encoding/json"

// Purchase 采购
type Purchase struct {
	BasicModel               `bson:"inline"`
	BookNameInfo             `bson:"-"`
	ProviderID               *int                      `json:"provider_id,omitempty" form:"provider_id"  bson:"provider_id" cn:"供应商编号"`
	ProviderName             *string                   `json:"provider_name,omitempty" form:"provider_name"  bson:"provider_name" cn:"供应商名称"`
	ProviderOwner            *string                   `json:"provider_owner,omitempty" form:"provider_owner" bson:"provider_owner" cn:"对方联络人"`
	ProviderOwnerPhone       *string                   `json:"provider_owner_phone,omitempty" form:"provider_owner_phone" bson:"provider_owner_phone" cn:"对方联系方式"`
	PurchaseOwnerID          *int                      `json:"purchase_owner_id,omitempty" form:"purchase_owner_id" bson:"purchase_owner_id" cn:"负责人ID"`
	PurchaseOwnerName        *string                   `json:"purchase_owner_name,omitempty" form:"purchase_owner_name" bson:"purchase_owner_name" cn:"负责人名称"`
	PurchaseTotalAmount      string                    `json:"purchase_total_amount" form:"purchase_total_amount" bson:"purchase_total_amount"  cn:"总金额"`
	PurchaseTotalIncludedTax string                    `json:"purchase_total_included_tax" form:"purchase_total_included_tax" bson:"purchase_total_included_tax"  cn:"含税总额"`
	PurchaseContent          *[]map[string]interface{} `json:"purchase_content" form:"purchase_content" bson:"purchase_content" cn:"采购内容"`
	PurchaseDate             *string                   `json:"purchase_date" form:"purchase_date" bson:"purchase_date"  cn:"采购日期"`
	Remark                   *string                   `json:"remark,omitempty" form:"remark,omitempty" bson:"remark" cn:"备注"`
	//PurchaseTaxRate          string                    `json:"purchase_tax_rate" form:"purchase_tax_rate" bson:"purchase_tax_rate"  cn:"采购税率"`
	//PurchaseTax              string                    `json:"purchase_tax" form:"purchase_tax" bson:"purchase_tax"  cn:"采购税额"`
}

func (p Purchase) Mapping() map[string]interface{} {
	ma := mapping{
		"settings": mapping{},
		"mappings": mapping{
			"properties": mapping{
				"rec_id": mapping{
					"type": "keyword",
				},
				"purchase_content": mapping{
					"type":            "text",
					"analyzer":        IK_SMART,
					"search_analyzer": IK_SMART,
				},
				"purchase_owner_name": mapping{
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
				"provider_name": mapping{
					"type":            "text",
					"analyzer":        IK_SMART,
					"search_analyzer": IK_SMART,
				},
				"remark": mapping{
					"type":            "text",
					"analyzer":        IK_SMART,
					"search_analyzer": IK_SMART,
				},
				"created_at": mapping{
					"type": "text",
				},
				"book_name": mapping{
					"type": "keyword",
				},
				"book_name_id": mapping{
					"type": "keyword",
				},
			},
		},
	}
	return ma
}

func (p Purchase) ToESDoc() map[string]interface{} {
	var str string
	bytes, err := json.Marshal(p.PurchaseContent)
	str = string(bytes)
	if err != nil {
		str = ""
	}

	return map[string]interface{}{
		"rec_id":              p.RecID,
		"created_at":          p.CreatedAt,
		"remark":              p.Remark,
		"purchase_content":    str,
		"provider_name":       p.ProviderName,
		"purchase_owner_name": p.PurchaseOwnerName,
		"book_name":           p.BookName,
		"book_name_id":        p.BookNameID,
	}
}

func (p Purchase) TableCnName() string {
	return "采购"
}

func (p Purchase) TableName() string {
	return "purchases"

}
