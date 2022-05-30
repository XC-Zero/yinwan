package storage

import (
	"context"
	"fmt"
	"github.com/XC-Zero/yinwan/internal/controller/services_controller/common"
	_const "github.com/XC-Zero/yinwan/pkg/const"
	"github.com/XC-Zero/yinwan/pkg/model/mysql_model"
	"github.com/XC-Zero/yinwan/pkg/utils/errs"
	"github.com/XC-Zero/yinwan/pkg/utils/logger"
	"github.com/XC-Zero/yinwan/pkg/utils/math_plus"
	"github.com/XC-Zero/yinwan/pkg/utils/mysql"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/pkg/errors"
	"strconv"
)

// CreateMaterial 创建原材料
func CreateMaterial(ctx *gin.Context) {
	bk, bookName := common.HarvestClientFromGinContext(ctx)
	if bk == nil {
		return
	}
	var material mysql_model.Material
	err := ctx.ShouldBindBodyWith(&material, binding.JSON)
	if err != nil {
		logger.Error(errors.WithStack(err), "绑定原材料失败!")
		common.RequestParamErrorTemplate(ctx, common.REQUEST_PARM_ERROR)
		return
	}

	err = bk.MysqlClient.WithContext(context.WithValue(context.Background(), "book_name", bookName)).Create(&material).Error
	if err != nil {
		logger.Error(errors.WithStack(err), "创建原材料失败!")
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
	err := ctx.ShouldBindBodyWith(&material, binding.JSON)
	if err != nil || material.RecID == nil {
		logger.Error(errors.WithStack(err), "")
		common.RequestParamErrorTemplate(ctx, common.REQUEST_PARM_ERROR)
		return
	}
	err = bk.MysqlClient.WithContext(context.WithValue(context.Background(), "book_name", bookName)).
		Updates(&material).Where("rec_id", *material.RecID).Error
	if err != nil {
		logger.Error(errors.WithStack(err), "更新原材料失败!")
		common.InternalDataBaseErrorTemplate(ctx, common.DATABASE_UPDATE_ERROR, material)
		return
	}
	ctx.JSON(_const.OK, errs.CreateSuccessMsg("更新原材料成功!"))
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
		logger.Error(errors.WithStack(err), "")
		common.RequestParamErrorTemplate(ctx, common.REQUEST_PARM_ERROR)
		return
	}
	material := mysql_model.Material{BasicModel: mysql_model.BasicModel{
		RecID: &id,
	}}
	err = bk.MysqlClient.WithContext(context.WithValue(context.Background(), "book_name", bookName)).
		Delete(&material).Error
	if err != nil {
		logger.Error(errors.WithStack(err), "")
		common.InternalDataBaseErrorTemplate(ctx, common.DATABASE_DELETE_ERROR, material)
		return
	}
	ctx.JSON(_const.OK, errs.CreateSuccessMsg("删除原材料成功!"))
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
	err := ctx.ShouldBindBodyWith(&materialBatch, binding.JSON)
	if err != nil {
		logger.Error(errors.WithStack(err), "")
		common.RequestParamErrorTemplate(ctx, common.REQUEST_PARM_ERROR)
		return
	}
	err = bk.MysqlClient.WithContext(context.WithValue(context.Background(), "book_name", bookName)).
		Create(&materialBatch).Error
	if err != nil {
		logger.Error(errors.WithStack(err), "")
		common.InternalDataBaseErrorTemplate(ctx, common.DATABASE_INSERT_ERROR, materialBatch)
		return
	}
	ctx.JSON(_const.OK, errs.CreateSuccessMsg("创建批次信息成功!"))
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
	bk, bookName := common.HarvestClientFromGinContext(ctx)
	if bk == nil {
		return
	}
	recID, err := strconv.Atoi(ctx.PostForm("material_batch_id"))
	if err != nil {
		common.RequestParamErrorTemplate(ctx, common.REQUEST_PARM_ERROR)
		return
	}
	var materialBatch mysql_model.MaterialBatch
	materialBatch.RecID = &recID
	err = bk.MysqlClient.WithContext(context.WithValue(context.Background(), "book_name", bookName)).
		Delete(&materialBatch).Error
	if err != nil {
		logger.Error(errors.WithStack(err), "删除批次失败!")
		common.InternalDataBaseErrorTemplate(ctx, common.DATABASE_DELETE_ERROR, materialBatch)
		return
	}
	ctx.JSON(_const.OK, errs.CreateSuccessMsg("删除批次成功!"))
	return
}

// SelectMaterialHistoryCost 详情页历史均价
func SelectMaterialHistoryCost(ctx *gin.Context) {
	bk, bookName := common.HarvestClientFromGinContext(ctx)
	if bk == nil {
		return
	}
	materialID := ctx.PostForm("material_id")
	if materialID == "" {
		common.RequestParamErrorTemplate(ctx, common.REQUEST_PARM_ERROR)
		return
	}
	conditions := []common.MysqlCondition{
		{
			Symbol:      mysql.EQUAL,
			ColumnName:  "material_id",
			ColumnValue: materialID,
		},
		{
			Symbol:      mysql.NULL,
			ColumnName:  "deleted_at",
			ColumnValue: " ",
		},
	}
	res := common.SelectMysqlTableContentWithCountTemplate(ctx, common.SelectMysqlTemplateOptions{
		DB:            bk.MysqlClient.WithContext(context.WithValue(context.Background(), "book_name", bookName)),
		OrderByColumn: "created_at",
		TableModel:    mysql_model.MaterialHistoryCost{},
		NotReturn:     true,
		NotPaginate:   true,
	}, conditions...)
	if res == nil {
		common.InternalDataBaseErrorTemplate(ctx, common.DATABASE_SELECT_ERROR, mysql_model.MaterialHistoryCost{})
		return
	}
	costs := res.([]mysql_model.MaterialHistoryCost)
	var count int64 = 0
	var errorList []error
	var totalPrice = math_plus.NewFromFloatByDecimal(0.0, 1)
	var dataList = make([]mysql_model.MaterialHistoryCost, 0, 15)
	for i := range costs {
		n, err := math_plus.NewFromString(costs[i].Price)
		if err != nil {
			errorList = append(errorList, errors.WithStack(err))
			continue
		}
		one := int64(costs[i].Num)
		totalPrice = totalPrice.Add(n.MulInt64(one))
		count += one
		if len(dataList) < 15 {
			dataList = append(dataList, costs[i])
		}
	}
	m, _ := math_plus.New(count, 1)
	var avg float64
	if count != 0 {
		avg = totalPrice.Div(m).Float64()
	} else {
		avg = 0
	}

	ctx.JSON(_const.OK, gin.H{
		"count":         count,
		"average_price": fmt.Sprintf("%.2f", avg),
		"list":          dataList,
	})
	return
}
