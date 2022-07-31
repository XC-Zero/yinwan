package storage

import (
	"github.com/XC-Zero/yinwan/internal/controller/services_controller/common"
	_const "github.com/XC-Zero/yinwan/pkg/const"
	"github.com/XC-Zero/yinwan/pkg/model/mysql_model"
	"github.com/XC-Zero/yinwan/pkg/utils/errs"
	"github.com/XC-Zero/yinwan/pkg/utils/logger"
	"github.com/XC-Zero/yinwan/pkg/utils/mysql"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/pkg/errors"
	"strconv"
)

func CreateWarehouse(ctx *gin.Context) {
	bk := common.HarvestClientFromGinContext(ctx)
	if bk == nil {
		return
	}
	var warehouse mysql_model.Warehouse
	err := ctx.ShouldBindBodyWith(&warehouse, binding.JSON)
	if err != nil {
		logger.Error(errors.WithStack(err), "")
		common.RequestParamErrorTemplate(ctx, common.REQUEST_PARM_ERROR)
		return
	}
	err = bk.MysqlClient.WithContext(ctx).Create(&warehouse).Error
	if err != nil {
		logger.Error(errors.WithStack(err), "")
		common.InternalDataBaseErrorTemplate(ctx, common.DATABASE_INSERT_ERROR, warehouse)
		return
	}
	ctx.JSON(_const.OK, errs.CreateSuccessMsg("创建仓库成功!"))
	return
}

func SelectWarehouse(ctx *gin.Context) {
	bk := common.HarvestClientFromGinContext(ctx)
	if bk == nil {
		return
	}
	conditions := []common.MysqlCondition{
		{
			Symbol:      mysql.LIKE,
			ColumnName:  "warehouse_name",
			ColumnValue: ctx.PostForm("warehouse_name"),
		},
		{
			Symbol:      mysql.EQUAL,
			ColumnName:  "rec_id",
			ColumnValue: ctx.PostForm("warehouse_id"),
		},
		{
			Symbol:      mysql.EQUAL,
			ColumnName:  "warehouse_owner_id",
			ColumnValue: ctx.PostForm("warehouse_owner_id"),
		},
		{
			Symbol:      mysql.NULL,
			ColumnName:  "deleted_at",
			ColumnValue: " ",
		},
	}
	op := common.SelectMysqlTemplateOptions{
		DB:         bk.MysqlClient.WithContext(ctx),
		TableModel: mysql_model.Warehouse{},
	}
	common.SelectMysqlTableContentWithCountTemplate(ctx, op, conditions...)
	return

}
func UpdateWarehouse(ctx *gin.Context) {
	bk := common.HarvestClientFromGinContext(ctx)
	if bk == nil {
		return
	}
	var warehouse mysql_model.Warehouse
	err := ctx.ShouldBindBodyWith(&warehouse, binding.JSON)
	if err != nil || warehouse.RecID == nil {
		logger.Error(errors.WithStack(err), "")
		common.RequestParamErrorTemplate(ctx, common.REQUEST_PARM_ERROR)
		return
	}
	err = bk.MysqlClient.WithContext(ctx).
		Updates(&warehouse).Where("rec_id", *warehouse.RecID).Error
	if err != nil {
		logger.Error(errors.WithStack(err), "")
		common.InternalDataBaseErrorTemplate(ctx, common.DATABASE_UPDATE_ERROR, warehouse)
		return
	}
	ctx.JSON(_const.OK, errs.CreateSuccessMsg("更新仓库成功!"))

	return
}

func DeleteWarehouse(ctx *gin.Context) {
	bk := common.HarvestClientFromGinContext(ctx)
	if bk == nil {
		return
	}

	recID, err := strconv.Atoi(ctx.PostForm("warehouse_id"))
	if err != nil {
		logger.Error(errors.WithStack(err), "")
		common.RequestParamErrorTemplate(ctx, common.REQUEST_PARM_ERROR)
		return
	}
	var warehouse = mysql_model.Warehouse{BasicModel: mysql_model.BasicModel{
		RecID: &recID,
	}}

	err = bk.MysqlClient.WithContext(
		ctx).Delete(&warehouse).Error
	if err != nil {
		logger.Error(errors.WithStack(err), "删除仓库失败!")
		common.InternalDataBaseErrorTemplate(ctx, common.DATABASE_DELETE_ERROR, warehouse)
		return
	}
	ctx.JSON(_const.OK, errs.CreateSuccessMsg("删除仓库成功!"))
	return
}
