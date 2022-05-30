package mongo_model

import "github.com/XC-Zero/yinwan/pkg/model/mysql_model"

// Purchase 采购
//  存在MongoDB里
//  凡是什么什么单含有不确定数据结构的均存在MongoDB里
type Purchase struct {
	mysql_model.BasicModel   `bson:"inline"`
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
	//TODO implement me
	panic("implement me")
}

func (p Purchase) ToESDoc() map[string]interface{} {
	//TODO implement me
	panic("implement me")
}

func (p Purchase) TableCnName() string {
	return "采购"
}

func (p Purchase) TableName() string {
	return "purchases"

}
