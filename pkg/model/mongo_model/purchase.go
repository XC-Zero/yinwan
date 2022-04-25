package mongo_model

import (
	"github.com/XC-Zero/yinwan/pkg/model/mysql_model/common"
	"time"
)

// Purchase 采购
//  存在MongoDB里
//  凡是什么什么单含有不确定数据结构的均存在MongoDB里
type Purchase struct {
	common.BasicModel
	ProviderID               int                     `json:"provider_id" form:"provider_id" binding:"required" bson:"provider_id" cn:"供应商编号"`
	ProviderName             string                  `json:"provider_name" form:"provider_name" binding:"required" bson:"provider_name" cn:"供应商名称"`
	ProviderOwner            *string                 `json:"provider_owner,omitempty" form:"provider_owner" bson:"provider_owner" cn:"对方负责人"`
	ProviderOwnerPhone       *string                 `json:"provider_owner_phone,omitempty" form:"provider_owner_phone" bson:"provider_owner_phone" cn:"对方负责人联系方式"`
	PurchaseOwnerID          *int                    `json:"purchase_owner_id,omitempty" form:"purchase_owner_id" bson:"purchase_owner_id" cn:"我方负责人编号"`
	PurchaseOwnerName        *string                 `json:"owner_id,omitempty" form:"owner_id" bson:"owner_id" cn:"我方负责人编号"`
	PurchaseTaxRate          float64                 `json:"purchase_tax_rate" form:"purchase_tax_rate" bson:"purchase_tax_rate" binding:"required" cn:"采购税率"`
	PurchaseTax              float64                 `json:"purchase_tax" form:"purchase_tax" bson:"purchase_tax" binding:"required" cn:"采购税额"`
	PurchaseTotalAmount      float64                 `json:"purchase_total_amount" form:"purchase_total_amount" bson:"purchase_total_amount" binding:"required" cn:"总金额"`
	PurchaseTotalIncludedTax float64                 `json:"purchase_total_included_tax" form:"purchase_total_included_tax" bson:"purchase_total_included_tax" binding:"required" cn:"含税总额"`
	PurchaseContent          *map[string]interface{} `json:"purchase_content" form:"purchase_content" bson:"purchase_content" cn:"采购内容"`
	PurchaseDate             time.Time               `json:"purchase_date" form:"purchase_date" bson:"purchase_date" binding:"required" cn:"采购日期"`
}

func (p Purchase) TableCnName() string {
	return "采购"
}

func (p Purchase) TableName() string {
	return "purchases"

}
