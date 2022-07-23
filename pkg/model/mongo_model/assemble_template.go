package mongo_model

// AssembleTemplate 组装拆卸模板
type AssembleTemplate struct {
	BasicModel            `bson:"inline"`
	BookNameInfo          `bson:"-"`
	AssembleTemplateName  string
	AssembleMaterialList  []assembleMaterial  `bson:"assemble_material_list" json:"assemble_material_list" form:"assemble_material_list"`
	AssembleCommodityList []assembleCommodity `bson:"assemble_commodity_list" json:"assemble_commodity_list" form:"assemble_commodity_list"`
	OperatorID            *int                `bson:"operator_id" json:"operator_id" form:"operator_id"`
	OperatorName          *string             `bson:"operator_name" json:"operator_name" form:"operator_name"`
	Remark                *string             `bson:"remark" json:"remark" form:"remark"`
}

type assembleMaterial struct {
	MaterialID   int    `bson:"material_id" json:"material_id" form:"material_id"`
	MaterialName string `bson:"material_name" json:"material_name" form:"material_name"`
	Num          int    `bson:"num" json:"num" form:"num"`
}
type assembleCommodity struct {
	CommodityID   int    `bson:"commodity_id" json:"commodity_id" form:"commodity_id"`
	CommodityName string `bson:"commodity_name" json:"commodity_name" form:"commodity_name"`
	Num           int    `bson:"num" json:"num" form:"num"`
}
