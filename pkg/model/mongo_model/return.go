package mongo_model

import "github.com/XC-Zero/yinwan/pkg/model/mysql_model"

type Return struct {
	mysql_model.BasicModel `bson:"inline"`
	BookNameInfo           `bson:"-"`
	ReturnContent          map[string]interface{} `json:"return_content" form:"return_content" bson:"return_content" cn:"退款订单内容" `
	ReturnActualAmount     string                 `json:"return_actual_amount" form:"return_actual_amount" bson:"return_actual_amount" cn:"实际退款金额" `
	ReturnAmount           string                 `json:"return_bill_amount" form:"return_bill_amount" cn:"应退金额" `
	ReturnOwnerID          *int                   `json:"return_owner_id,omitempty" form:"return_owner_id,omitempty" json:"return_owner_id" cn:"退款记录负责人ID"`
	ReturnOwnerName        *string                `json:"return_owner_name,omitempty" form:"return_owner_name,omitempty" bson:"return_owner_name" cn:"退款记录负责人名称"`
	TransactionID          *int                   `json:"transaction_id,omitempty" form:"transaction_id" bson:"transaction_id" cn:"关联销售订单ID"`
	ReceiveID              *int                   `json:"receive_id,omitempty" form:"receive_id,omitempty" bson:"receive_id" cn:"关联红字应收记录ID" `
	Remark                 *string                `json:"remark,omitempty" form:"remark,omitempty" bson:"remark" cn:"备注"`
}

func (r Return) TableCnName() string {
	return "退货记录"
}

func (r Return) TableName() string {
	return "returns"
}

func (r Return) Mapping() map[string]interface{} {
	ma := mapping{
		"settings": mapping{},
		"mappings": mapping{
			"properties": mapping{
				"rec_id": mapping{
					"type": "keyword",
				},
				"return_content": mapping{
					"type":            "text",
					"analyzer":        IK_SMART,
					"search_analyzer": IK_SMART,
				},
				"return_owner_name": mapping{
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
				"receive_id": mapping{
					"type": "keyword",
				},
				"transaction_id": mapping{
					"type": "keyword",
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

// ToESDoc todo content转字符串
func (r Return) ToESDoc() map[string]interface{} {
	return map[string]interface{}{
		"rec_id":            r.RecID,
		"created_at":        r.CreatedAt,
		"remark":            r.Remark,
		"receive_id":        r.ReceiveID,
		"transaction_id":    r.TransactionID,
		"return_owner_name": r.ReturnOwnerName,
		//"return_content":    r.ReturnContent,
		"book_name":    r.BookName,
		"book_name_id": r.BookNameID,
	}
}
