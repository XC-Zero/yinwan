package finance

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
	"golang.org/x/net/context"
	"strconv"
)

// CreateFixedAsset 创建固定资产
func CreateFixedAsset(ctx *gin.Context) {
	bk, n := common.HarvestClientFromGinContext(ctx)
	if bk == nil {
		return
	}
	var fixedAsset mysql_model.FixedAsset
	err := ctx.ShouldBindBodyWith(&fixedAsset, binding.JSON)
	if err != nil {
		common.RequestParamErrorTemplate(ctx, common.REQUEST_PARM_ERROR)
		return
	}

	err = bk.MysqlClient.WithContext(context.WithValue(context.Background(), "book_name", n)).Create(&fixedAsset).Error
	if err != nil {
		logger.Error(errors.WithStack(err), "创建固定资产失败!")
		common.InternalDataBaseErrorTemplate(ctx, common.DATABASE_INSERT_ERROR, fixedAsset)
		return
	}
	ctx.JSON(_const.OK, errs.CreateSuccessMsg("创建固定资产成功！"))
	return
}

func SelectFixedAsset(ctx *gin.Context) {
	bk, n := common.HarvestClientFromGinContext(ctx)
	if bk == nil {
		return
	}
	conditions := []common.MysqlCondition{
		{
			Symbol:      mysql.EQUAL,
			ColumnName:  "rec_id",
			ColumnValue: ctx.PostForm("fixed_asset_id"),
		},
		{
			Symbol:      mysql.LIKE,
			ColumnName:  "fixed_asset_name",
			ColumnValue: ctx.PostForm("fixed_asset_name"),
		},
		{
			Symbol:      mysql.EQUAL,
			ColumnName:  "fixed_asset_type_id",
			ColumnValue: ctx.PostForm("fixed_asset_type_id"),
		},
		{
			Symbol:      mysql.NULL,
			ColumnName:  "deleted_at",
			ColumnValue: " ",
		}}
	op := common.SelectMysqlTemplateOptions{
		DB:         bk.MysqlClient.WithContext(context.WithValue(context.Background(), "book_name", n)),
		TableModel: mysql_model.FixedAsset{},
	}
	common.SelectMysqlTableContentWithCountTemplate(ctx, op, conditions...)
	return
}

// UpdateFixedAsset 更新固定资产
func UpdateFixedAsset(ctx *gin.Context) {
	bk, n := common.HarvestClientFromGinContext(ctx)
	if bk == nil {
		return
	}
	var fixedAsset mysql_model.FixedAsset
	err := ctx.ShouldBindBodyWith(&fixedAsset, binding.JSON)
	if err != nil || fixedAsset.RecID == nil {
		common.RequestParamErrorTemplate(ctx, common.REQUEST_PARM_ERROR)
		return
	}
	err = bk.MysqlClient.WithContext(context.WithValue(context.Background(), "book_name", n)).
		Updates(&fixedAsset).Where("rec_id = ?", fixedAsset.RecID).Error
	if err != nil {
		logger.Error(errors.WithStack(err), "更新固定资产失败!")
		common.InternalDataBaseErrorTemplate(ctx, common.DATABASE_UPDATE_ERROR, fixedAsset)
		return
	}
	ctx.JSON(_const.OK, errs.CreateSuccessMsg("更新固定资产成功!"))
	return
}

// DeleteFixedAsset 删除固定资产
func DeleteFixedAsset(ctx *gin.Context) {
	bk, n := common.HarvestClientFromGinContext(ctx)
	if bk == nil {
		return
	}

	id, err := strconv.Atoi(ctx.PostForm("fixed_asset_id"))

	if err != nil {
		common.RequestParamErrorTemplate(ctx, common.REQUEST_PARM_ERROR)
		return
	}
	asset := mysql_model.FixedAsset{}
	asset.RecID = &id
	err = bk.MysqlClient.WithContext(context.WithValue(context.Background(), "book_name", n)).
		Delete(&asset).Error
	if err != nil {
		logger.Error(errors.WithStack(err), "删除固定资产失败!")
		common.InternalDataBaseErrorTemplate(ctx, common.DATABASE_DELETE_ERROR, asset)
		return
	}
	ctx.JSON(_const.OK, errs.CreateSuccessMsg("删除固定资产成功!"))
	return
}
