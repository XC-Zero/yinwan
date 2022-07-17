package finance

import (
	"context"
	"fmt"
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

func CreatePayable(ctx *gin.Context) {
	bk, n := common.HarvestClientFromGinContext(ctx)
	if bk == nil {
		return
	}
	temp := mysql_model.Payable{}

	err := ctx.ShouldBindBodyWith(&temp, binding.JSON)
	if err != nil {
		logger.Error(errors.WithStack(err), "绑定模型失败!")
		common.RequestParamErrorTemplate(ctx, common.REQUEST_PARM_ERROR)
		return
	}
	err = bk.MysqlClient.WithContext(context.WithValue(context.Background(), "book_name", n)).
		Create(&temp).Error
	if err != nil {
		logger.Error(errors.WithStack(err), "绑定模型失败!")
		common.InternalDataBaseErrorTemplate(ctx, common.DATABASE_INSERT_ERROR, temp)
		return
	}
	ctx.JSON(_const.OK, errs.CreateSuccessMsg("新建应付成功!"))
	return
}

func SelectPayable(ctx *gin.Context) {
	bk, n := common.HarvestClientFromGinContext(ctx)
	if bk == nil {
		return
	}
	condition := []common.MysqlCondition{
		{
			Symbol:      mysql.EQUAL,
			ColumnName:  "rec_id",
			ColumnValue: ctx.PostForm("payable_id"),
		},
		{
			Symbol:      mysql.LIKE,
			ColumnName:  "provider_name",
			ColumnValue: ctx.PostForm("provider_name"),
		},
		{
			Symbol:      mysql.GREATER_THEN_EQUAL,
			ColumnName:  "payable_date",
			ColumnValue: ctx.PostForm("payable_date"),
		},
		{
			Symbol:      mysql.LESS_THAN_EQUAL,
			ColumnName:  "payable_date",
			ColumnValue: ctx.PostForm("payable_date2"),
		},
		{
			Symbol:      mysql.NULL,
			ColumnName:  "deleted_at",
			ColumnValue: " ",
		},
	}
	op := common.SelectMysqlTemplateOptions{
		DB:         bk.MysqlClient.WithContext(context.WithValue(context.Background(), "book_name", n)),
		TableModel: mysql_model.Payable{},
	}
	common.SelectMysqlTableContentWithCountTemplate(ctx, op, condition...)
	return
}

func UpdatePayable(ctx *gin.Context) {
	bk, n := common.HarvestClientFromGinContext(ctx)
	if bk == nil {
		return
	}
	temp := mysql_model.Payable{}

	err := ctx.ShouldBindBodyWith(&temp, binding.JSON)
	if err != nil || temp.RecID == nil {
		logger.Error(errors.WithStack(err), "更新应付失败!")
		common.RequestParamErrorTemplate(ctx, common.REQUEST_PARM_ERROR)
		return
	}
	err = bk.MysqlClient.WithContext(context.WithValue(context.Background(), "book_name", n)).
		Updates(&temp).Where("rec_id = ?", temp.RecID).Error
	if err != nil {
		logger.Error(errors.WithStack(err), "更新应付记录失败!")
		common.InternalDataBaseErrorTemplate(ctx, common.DATABASE_UPDATE_ERROR, temp)
		return
	}
	ctx.JSON(_const.OK, errs.CreateSuccessMsg("更新应付记录成功!"))
	return
}

func DeletePayable(ctx *gin.Context) {
	bk, n := common.HarvestClientFromGinContext(ctx)
	if bk == nil {
		return
	}
	var payable mysql_model.Payable
	recID, err := strconv.Atoi(ctx.PostForm("payable_id"))
	if err != nil {
		common.RequestParamErrorTemplate(ctx, common.REQUEST_PARM_ERROR)
		return
	}
	payable.RecID = &recID
	err = bk.MysqlClient.WithContext(context.WithValue(context.Background(), "book_name", n)).
		Delete(&payable).Where("rec_id = ? ", recID).Error
	if err != nil {
		common.RequestParamErrorTemplate(ctx, common.REQUEST_PARM_ERROR)
		return
	}
	logger.Info(fmt.Sprintf("删除应付记录成功!记录ID:%d 操作人: %s", recID, common.HarvestEmailFromHeader(ctx)))
	ctx.JSON(_const.OK, errs.CreateSuccessMsg("删除应付记录成功!"))
	return
}

func CreatePayableDetail(ctx *gin.Context) {
	bk, n := common.HarvestClientFromGinContext(ctx)
	if bk == nil {
		return
	}
	temp := mysql_model.PayableDetail{}

	err := ctx.ShouldBindBodyWith(&temp, binding.JSON)
	if err != nil || temp.PayableID == nil {
		logger.Error(errors.WithStack(err), "绑定模型失败!")
		common.RequestParamErrorTemplate(ctx, common.REQUEST_PARM_ERROR)
		return
	}
	err = bk.MysqlClient.WithContext(context.WithValue(context.Background(), "book_name", n)).
		Create(&temp).Error
	if err != nil {
		logger.Error(errors.WithStack(err), "绑定模型失败!")
		common.InternalDataBaseErrorTemplate(ctx, common.DATABASE_INSERT_ERROR, temp)
		return
	}
	ctx.JSON(_const.OK, errs.CreateSuccessMsg("新建应付详情成功!"))
	return
}

func UpdatePayableDetail(ctx *gin.Context) {
	bk, n := common.HarvestClientFromGinContext(ctx)
	if bk == nil {
		return
	}
	temp := mysql_model.PayableDetail{}

	err := ctx.ShouldBindBodyWith(&temp, binding.JSON)
	if err != nil || temp.RecID == nil {
		logger.Error(errors.WithStack(err), "更新应付详情失败!")
		common.RequestParamErrorTemplate(ctx, common.REQUEST_PARM_ERROR)
		return
	}
	err = bk.MysqlClient.WithContext(context.WithValue(context.Background(), "book_name", n)).
		Updates(&temp).Where("rec_id = ?", temp.RecID).Error
	if err != nil {
		logger.Error(errors.WithStack(err), "更新应付详情记录失败!")
		common.InternalDataBaseErrorTemplate(ctx, common.DATABASE_UPDATE_ERROR, temp)
		return
	}
	ctx.JSON(_const.OK, errs.CreateSuccessMsg("更新应付详情记录成功!"))
	return
}

func DeletePayableDetail(ctx *gin.Context) {
	bk, n := common.HarvestClientFromGinContext(ctx)
	if bk == nil {
		return
	}
	var payableDetail mysql_model.PayableDetail
	recID, err := strconv.Atoi(ctx.PostForm("payable_detail_id"))
	if err != nil {
		common.RequestParamErrorTemplate(ctx, common.REQUEST_PARM_ERROR)
		return
	}
	payableDetail.RecID = &recID
	err = bk.MysqlClient.WithContext(context.WithValue(context.Background(), "book_name", n)).
		Delete(&payableDetail).Where("rec_id = ? ", recID).Error
	if err != nil {
		common.RequestParamErrorTemplate(ctx, common.REQUEST_PARM_ERROR)
		return
	}
	logger.Info(fmt.Sprintf("删除应付记录成功!记录ID:%d 操作人: %s", recID, common.HarvestEmailFromHeader(ctx)))
	ctx.JSON(_const.OK, errs.CreateSuccessMsg("删除应付记录成功!"))
	return
}

func SelectPayableDetail(ctx *gin.Context) {
	bk, n := common.HarvestClientFromGinContext(ctx)
	if bk == nil {
		return
	}
	condition := []common.MysqlCondition{
		{
			mysql.EQUAL,
			"rec_id",
			ctx.PostForm("payable_detail_id"),
		},
		{
			mysql.EQUAL,
			"payable_id",
			ctx.PostForm("payable_id"),
		},
		{
			mysql.NOT_EQUAL,
			"deleted_at",
			" ",
		},
		{
			mysql.NOT_NULL,
			"deleted_at",
			ctx.PostForm("is_deleted"),
		},
	}
	op := common.SelectMysqlTemplateOptions{
		DB:         bk.MysqlClient.WithContext(context.WithValue(context.Background(), "book_name", n)),
		TableModel: mysql_model.PayableDetail{},
	}
	common.SelectMysqlTableContentWithCountTemplate(ctx, op, condition...)
	return
}
