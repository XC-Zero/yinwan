package storage

import (
	"context"
	"github.com/XC-Zero/yinwan/internal/controller/services_controller/common"
	_const "github.com/XC-Zero/yinwan/pkg/const"
	"github.com/XC-Zero/yinwan/pkg/model/mysql_model"
	"github.com/XC-Zero/yinwan/pkg/utils/errs"
	"github.com/XC-Zero/yinwan/pkg/utils/mysql"
	"github.com/gin-gonic/gin"
	"log"
	"strconv"
)

// CreateMaterial 创建原材料
func CreateMaterial(ctx *gin.Context) {
	bk, bookName := common.HarvestClientFromGinContext(ctx)
	if bk == nil {
		return
	}
	var material mysql_model.Material
	err := ctx.ShouldBind(&material)
	if err != nil {
		log.Println(err)
		common.RequestParamErrorTemplate(ctx, common.REQUEST_PARM_ERROR)
		return
	}

	err = bk.MysqlClient.WithContext(context.WithValue(context.Background(), "book_name", bookName)).Create(&material).Error
	if err != nil {
		log.Println(err)
		common.InternalDataBaseErrorTemplate(ctx, common.DATABASE_INSERT_ERROR, material)
		return
	}
	ctx.JSON(_const.OK, errs.CreateSuccessMsg("创建原材料成功！"))
	return

}

// SelectMaterial 原材料
func SelectMaterial(ctx *gin.Context) {
	bk, bookName := common.HarvestClientFromGinContext(ctx)
	if bk == nil {
		return
	}
	conditions := []common.MysqlCondition{
		{
			Symbol:      mysql.LIKE,
			ColumnName:  "material_name",
			ColumnValue: ctx.PostForm("material_name"),
		},
		{
			Symbol:      mysql.EQUAL,
			ColumnName:  "rec_id",
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
	op := common.SelectMysqlTemplateOptions{
		DB:         bk.MysqlClient.WithContext(context.WithValue(context.Background(), "book_name", bookName)),
		TableModel: mysql_model.Material{},
	}
	common.SelectMysqlTableContentWithCountTemplate(ctx, op, conditions...)

	return
}

func UpdateMaterial(ctx *gin.Context) {
	bk, bookName := common.HarvestClientFromGinContext(ctx)
	if bk == nil {
		return
	}
	material := mysql_model.Material{}
	err := ctx.ShouldBind(&material)
	if err != nil || material.RecID == nil {
		common.RequestParamErrorTemplate(ctx, common.REQUEST_PARM_ERROR)
		return
	}
	err = bk.MysqlClient.WithContext(context.WithValue(context.Background(), "book_name", bookName)).
		Updates(&material).Where("rec_id", *material.RecID).Error
	if err != nil {
		common.InternalDataBaseErrorTemplate(ctx, common.DATABASE_UPDATE_ERROR, material)
		return
	}
	return
}

// DeleteMaterial 删除原材料
func DeleteMaterial(ctx *gin.Context) {
	bk, bookName := common.HarvestClientFromGinContext(ctx)
	if bk == nil {
		return
	}

	id, err := strconv.Atoi(ctx.PostForm("material_id"))
	if err != nil {
		common.RequestParamErrorTemplate(ctx, common.REQUEST_PARM_ERROR)
		return
	}
	material := mysql_model.Material{BasicModel: mysql_model.BasicModel{
		RecID: &id,
	}}
	err = bk.MysqlClient.WithContext(context.WithValue(context.Background(), "book_name", bookName)).
		Delete(material, id).Error
	if err != nil {
		common.InternalDataBaseErrorTemplate(ctx, common.DATABASE_DELETE_ERROR, material)
		return
	}
	ctx.JSON(_const.OK, "删除原材料成功!")
	return
}

// CreateMaterialBatch  创建原材料批次
//	(预留的,正常不应该有的,正常都应该走入库)
func CreateMaterialBatch(ctx *gin.Context) {
	bk, bookName := common.HarvestClientFromGinContext(ctx)
	if bk == nil {
		return
	}

	var materialBatch mysql_model.MaterialBatch
	err := ctx.ShouldBind(&materialBatch)
	if err != nil {
		common.RequestParamErrorTemplate(ctx, common.REQUEST_PARM_ERROR)
		return
	}
	err = bk.MysqlClient.WithContext(context.WithValue(context.Background(), "book_name", bookName)).
		Create(materialBatch).Error
	if err != nil {
		common.InternalDataBaseErrorTemplate(ctx, common.DATABASE_INSERT_ERROR, materialBatch)
		return
	}
	return
}

// SelectMaterialDetail 原材料批次信息
func SelectMaterialDetail(ctx *gin.Context) {
	bk, bookName := common.HarvestClientFromGinContext(ctx)
	if bk == nil {
		return
	}
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
	op := common.SelectMysqlTemplateOptions{
		DB:         bk.MysqlClient.WithContext(context.WithValue(context.Background(), "book_name", bookName)),
		TableModel: mysql_model.MaterialBatch{},
	}

	common.SelectMysqlTableContentWithCountTemplate(ctx, op, conditions...)

	return
}

// DeleteMaterialDetail 删除原材料批次信息
func DeleteMaterialDetail(ctx *gin.Context) {

}
