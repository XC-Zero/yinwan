package mongo_model

import _const "github.com/XC-Zero/yinwan/pkg/const"

// AssembleTemplate 组装拆卸模板
type AssembleTemplate struct {
	BasicModel            `bson:"inline"`
	BookNameInfo          `bson:"-"`
	AssembleTemplateName  *string           `form:"assemble_template_name,omitempty" json:"assemble_template_name,omitempty" bson:"assemble_template_name,omitempty"`
	AssembleMaterialList  []assembleElement `bson:"assemble_material_list" json:"assemble_material_list" form:"assemble_material_list"`
	AssembleCommodityList []assembleElement `bson:"assemble_commodity_list" json:"assemble_commodity_list" form:"assemble_commodity_list"`
	OperatorID            *int              `bson:"operator_id" json:"operator_id" form:"operator_id"`
	OperatorName          *string           `bson:"operator_name" json:"operator_name" form:"operator_name"`
	Remark                *string           `bson:"remark" json:"remark" form:"remark"`
}

func (a AssembleTemplate) TableCnName() string {
	return "组装拆卸模板"
}

func (a AssembleTemplate) TableName() string {
	return "assemble_templates"
}

type assembleElement struct {
	ContentType _const.StockContentType `json:"content_type" form:"content_type" bson:"content_type" cn:"产品/原材料"`
	RecID       int                     `json:"rec_id" form:"rec_id" bson:"rec_id" cn:"产品或原材料编号"`
	Name        string                  `json:"name" form:"name" bson:"name" cn:"产品或原材料名称"`
	Num         int                     `json:"num" form:"num" bson:"num" cn:"产品或原材料数量"`
}
