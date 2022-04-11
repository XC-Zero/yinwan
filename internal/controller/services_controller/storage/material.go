package storage

import (
	"github.com/XC-Zero/yinwan/internal/controller/services_controller/common"
	"github.com/XC-Zero/yinwan/pkg/client"
	"github.com/XC-Zero/yinwan/pkg/model"
	"github.com/XC-Zero/yinwan/pkg/utils/mysql"
	"github.com/gin-gonic/gin"
)

func CreateMaterial(ctx *gin.Context) {

}

// SelectMaterial 原材料
func SelectMaterial(ctx *gin.Context) {

	conditions := []common.MysqlCondition{
		{
			Symbol:      mysql.LIKE,
			ColumnName:  "material_name",
			ColumnValue: ctx.PostForm("material_name"),
		},
		{
			Symbol:      mysql.EQUAL,
			ColumnName:  "material_id",
			ColumnValue: ctx.PostForm("material_id"),
		},
		{
			Symbol:      mysql.EQUAL,
			ColumnName:  "material_type_id",
			ColumnValue: ctx.PostForm("material_type_id"),
		},
		{
			Symbol:      mysql.NULL,
			ColumnName:  "deleted_at",
			ColumnValue: " ",
		},
	}
	op := common.SelectMysqlTemplateOptions{DB: client.MysqlClient, TableModel: model.Material{}}
	common.SelectMysqlTableContentWithCountTemplate(ctx, op, conditions...)

	return
}

func UpdateMaterial(ctx *gin.Context) {

}

func DeleteMaterial(ctx *gin.Context) {

}

func CreateMaterialBatch(ctx *gin.Context) {

}

// SelectMaterialDetail 原材料批次信息
func SelectMaterialDetail(ctx *gin.Context) {

	conditions := []common.MysqlCondition{
		{
			Symbol:      mysql.EQUAL,
			ColumnName:  "warehouse_id",
			ColumnValue: ctx.PostForm("warehouse_id"),
		},
		{
			Symbol:      mysql.EQUAL,
			ColumnName:  "material_id",
			ColumnValue: ctx.PostForm("material_id"),
		},
		{
			Symbol:      mysql.EQUAL,
			ColumnName:  "material_type_id",
			ColumnValue: ctx.PostForm("material_type_id"),
		},
		{
			Symbol:      mysql.NULL,
			ColumnName:  "deleted_at",
			ColumnValue: " ",
		},
	}
	op := common.SelectMysqlTemplateOptions{DB: client.MysqlClient, TableModel: model.MaterialBatch{}}

	common.SelectMysqlTableContentWithCountTemplate(ctx, op, conditions...)

	return
}
