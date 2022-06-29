package mongo_model

import (
	"encoding/json"
	"github.com/XC-Zero/yinwan/pkg/client"
	_const "github.com/XC-Zero/yinwan/pkg/const"
	"github.com/XC-Zero/yinwan/pkg/model/mysql_model"
	"github.com/XC-Zero/yinwan/pkg/utils/logger"
	"github.com/pkg/errors"
	"golang.org/x/net/context"
)

// Purchase 采购
type Purchase struct {
	BasicModel               `bson:"inline"`
	BookNameInfo             `bson:"-"`
	ProviderID               *int               `json:"provider_id,omitempty" form:"provider_id"  bson:"provider_id" cn:"供应商编号"`
	ProviderName             *string            `json:"provider_name,omitempty" form:"provider_name"  bson:"provider_name" cn:"供应商名称"`
	ProviderOwner            *string            `json:"provider_owner,omitempty" form:"provider_owner" bson:"provider_owner" cn:"对方联络人"`
	ProviderOwnerPhone       *string            `json:"provider_owner_phone,omitempty" form:"provider_owner_phone" bson:"provider_owner_phone" cn:"对方联系方式"`
	PurchaseOwnerID          *int               `json:"purchase_owner_id,omitempty" form:"purchase_owner_id" bson:"purchase_owner_id" cn:"负责人ID"`
	PurchaseOwnerName        *string            `json:"purchase_owner_name,omitempty" form:"purchase_owner_name" bson:"purchase_owner_name" cn:"负责人名称"`
	PurchaseTotalAmount      string             `json:"purchase_total_amount" form:"purchase_total_amount" bson:"purchase_total_amount"  cn:"总金额"`
	PurchaseTotalIncludedTax string             `json:"purchase_total_included_tax" form:"purchase_total_included_tax" bson:"purchase_total_included_tax"  cn:"含税总额"`
	PurchaseContent          *[]purchaseContent `json:"purchase_content" form:"purchase_content" bson:"purchase_content" cn:"采购内容"`
	PurchaseDate             *string            `json:"purchase_date" form:"purchase_date" bson:"purchase_date"  cn:"采购日期"`
	Remark                   *string            `json:"remark,omitempty" form:"remark,omitempty" bson:"remark" cn:"备注"`
}

type purchaseContent struct {
	MaterialID         int     `bson:"material_id" json:"material_id" form:"material_id" cn:"原材料编号"`
	MaterialName       string  `bson:"material_name" json:"material_name" form:"material_name" cn:"名称"`
	MaterialStyle      string  `bson:"material_style" json:"material_style" form:"material_style" cn:"规格"`
	MaterialUnit       string  `bson:"material_unit" json:"material_unit" form:"material_unit" cn:"单位"`
	MaterialNum        int     `bson:"material_num" json:"material_num" form:"material_num" cn:"数量"`
	MaterialUnitPrice  string  `bson:"material_unit_price" json:"material_unit_price" form:"material_unit_price" cn:"单价"`
	MaterialTotalPrice string  `bson:"material_total_price" json:"material_total_price" form:"material_total_price" cn:"原材料总价"`
	PurchaseTaxRate    string  `bson:"purchase_tax_rate" json:"purchase_tax_rate" form:"purchase_tax_rate"  cn:"采购税率"`
	PurchaseTax        string  `bson:"purchase_tax" json:"purchase_tax" form:"purchase_tax"  cn:"采购税额"`
	PurchaseTotalPrice string  `bson:"purchase_total_price" json:"purchase_total_price" form:"purchase_total_price" cn:"含税总价"`
	Remark             *string `bson:"remark" json:"remark,omitempty" form:"remark,omitempty" cn:"备注"`
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

func (p *Purchase) AfterInsert(ctx context.Context) error {
	flag, ok := ctx.Value("auto_create").(bool)
	if !ok || (ok && !flag) {
		return nil
	}
	bookName := ctx.Value("book_name").(string)
	bk, ok := client.ReadBookMap(bookName)
	if !ok {
		return errors.New("There is no book name!")
	}
	if p.RecID == nil || bk.StorageName == "" {
		return errors.New("缺少主键！")
	}
	payableActualAmount := "0"
	unpaid := _const.UnPaid
	payable := mysql_model.Payable{
		ProviderName:        p.ProviderName,
		ProviderID:          p.ProviderID,
		PayableTotalAmount:  &p.PurchaseTotalAmount,
		PayableActualAmount: &payableActualAmount,
		PayableDebtAmount:   &p.PurchaseTotalAmount,
		PayableStatus:       &unpaid,
	}

	err := bk.MysqlClient.WithContext(ctx).Create(&payable).Error
	if err != nil {
		logger.Error(errors.WithStack(err), "同步创建应收记录失败!")
		return err
	}
	return nil
}
