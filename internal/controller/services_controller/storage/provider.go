package storage

import (
	"context"
	"github.com/XC-Zero/yinwan/internal/controller/services_controller/common"
	_const "github.com/XC-Zero/yinwan/pkg/const"
	"github.com/XC-Zero/yinwan/pkg/model/mysql_model"
	"github.com/XC-Zero/yinwan/pkg/utils/errs"
	"github.com/XC-Zero/yinwan/pkg/utils/logger"
	"github.com/XC-Zero/yinwan/pkg/utils/mysql"
	"github.com/fwhezfwhez/errorx"
	"github.com/gin-gonic/gin"
)

func CreateProvider(ctx *gin.Context) {
	bk, bookName := common.HarvestClientFromGinContext(ctx)
	if bk == nil {
		return
	}
	var provider mysql_model.Provider
	err := ctx.ShouldBind(&provider)
	if err != nil {
		logger.Error(errorx.MustWrap(err), "provider errors ")
		common.RequestParamErrorTemplate(ctx, common.REQUEST_PARM_ERROR)
		return
	}
	err = bk.MysqlClient.WithContext(context.WithValue(context.Background(), "book_name", bookName)).Create(&provider).Error
	if err != nil {
		common.InternalDataBaseErrorTemplate(ctx, common.DATABASE_INSERT_ERROR, provider)
		return
	}
	ctx.JSON(_const.OK, errs.CreateSuccessMsg("创建供应商成功!"))

	//common.CreateOneMysqlRecordTemplate(ctx, op)
	return
}
func SelectProvider(ctx *gin.Context) {
	bk, bookName := common.HarvestClientFromGinContext(ctx)
	if bk == nil {
		return
	}
	conditions := []common.MysqlCondition{
		{
			Symbol:      mysql.LIKE,
			ColumnName:  "provider_name",
			ColumnValue: ctx.PostForm("provider_name"),
		},
		{
			Symbol:      mysql.EQUAL,
			ColumnName:  "rec_id",
			ColumnValue: ctx.PostForm("provider_id"),
		},
		{
			Symbol:      mysql.EQUAL,
			ColumnName:  "provider_social_credit_code",
			ColumnValue: ctx.PostForm("provider_social_credit_code"),
		},
		{
			Symbol:      mysql.NULL,
			ColumnName:  "deleted_at",
			ColumnValue: " ",
		},
	}
	op := common.SelectMysqlTemplateOptions{
		DB:         bk.MysqlClient.WithContext(context.WithValue(context.Background(), "book_name", bookName)),
		TableModel: mysql_model.Provider{},
	}
	common.SelectMysqlTableContentWithCountTemplate(ctx, op, conditions...)
	return
}
func UpdateProvider(ctx *gin.Context) {
	bk, bookName := common.HarvestClientFromGinContext(ctx)
	if bk == nil {
		return
	}
	var provider mysql_model.Provider
	err := ctx.ShouldBind(&provider)
	if err != nil || provider.RecID == nil {
		common.RequestParamErrorTemplate(ctx, common.REQUEST_PARM_ERROR)
		return
	}
	err = bk.MysqlClient.WithContext(context.WithValue(context.Background(), "book_name", bookName)).
		Updates(&provider).Where("rec_id", *provider.RecID).Error
	if err != nil {
		common.InternalDataBaseErrorTemplate(ctx, common.DATABASE_UPDATE_ERROR, provider)
		return
	}
	return
}
func DeleteProvider(ctx *gin.Context) {
	bk, bookName := common.HarvestClientFromGinContext(ctx)
	if bk == nil {
		return
	}
	var provider mysql_model.Provider
	var recID = ctx.PostForm("provider_id")
	if recID == "" {
		common.RequestParamErrorTemplate(ctx, common.REQUEST_PARM_ERROR)
		return
	}
	err := bk.MysqlClient.WithContext(
		context.WithValue(context.Background(), "book_name", bookName)).
		Delete(&provider, recID).Error
	if err != nil {
		common.InternalDataBaseErrorTemplate(ctx, common.DATABASE_DELETE_ERROR, provider)
		return
	}
	return
}
