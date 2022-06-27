package finance

import (
	"context"
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

func CreateReceivable(ctx *gin.Context) {
	bk, n := common.HarvestClientFromGinContext(ctx)
	if bk == nil {
		return
	}
	var receivable mysql_model.Receivable
	err := ctx.ShouldBindBodyWith(&receivable, binding.JSON)
	if err != nil {
		common.RequestParamErrorTemplate(ctx, common.REQUEST_PARM_ERROR)
		return
	}
	err = bk.MysqlClient.WithContext(context.WithValue(context.Background(), "book_name", n)).Create(&receivable).Error
	if err != nil {
		logger.Error(errors.WithStack(err), "创建应收记录失败!")
		common.InternalDataBaseErrorTemplate(ctx, common.DATABASE_INSERT_ERROR, receivable)
		return
	}
	ctx.JSON(_const.OK, errs.CreateSuccessMsg("创建应收记录成功！"))
	return
}

func SelectReceivable(ctx *gin.Context) {
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
			ColumnName:  "customer_name",
			ColumnValue: ctx.PostForm("customer_name"),
		},
		{
			Symbol:      mysql.GREATER_THEN_EQUAL,
			ColumnName:  "receivable_date",
			ColumnValue: ctx.PostForm("receivable_date"),
		},
		{
			Symbol:      mysql.LESS_THAN_EQUAL,
			ColumnName:  "receivable_date",
			ColumnValue: ctx.PostForm("receivable_date2"),
		},
		{
			Symbol:      mysql.NULL,
			ColumnName:  "deleted_at",
			ColumnValue: " ",
		}}
	op := common.SelectMysqlTemplateOptions{
		DB:         bk.MysqlClient.WithContext(context.WithValue(context.Background(), "book_name", n)),
		TableModel: mysql_model.Receivable{},
	}
	common.SelectMysqlTableContentWithCountTemplate(ctx, op, conditions...)
	return

}

func UpdateReceivable(ctx *gin.Context) {
	bk, n := common.HarvestClientFromGinContext(ctx)
	if bk == nil {
		return
	}

	var receivable mysql_model.Receivable
	err := ctx.ShouldBindBodyWith(&receivable, binding.JSON)
	if err != nil || receivable.RecID == nil {
		common.RequestParamErrorTemplate(ctx, common.REQUEST_PARM_ERROR)
		return
	}

	err = bk.MysqlClient.WithContext(context.WithValue(context.Background(), "book_name", n)).
		Updates(&receivable).Where("rec_id = ?", receivable.RecID).Error
	if err != nil {
		logger.Error(errors.WithStack(err), "更新应收记录失败!")
		common.InternalDataBaseErrorTemplate(ctx, common.DATABASE_UPDATE_ERROR, receivable)
		return
	}
	ctx.JSON(_const.OK, errs.CreateSuccessMsg("更新应收记录成功!"))
	return
}

func DeleteReceivable(ctx *gin.Context) {
	bk, n := common.HarvestClientFromGinContext(ctx)
	if bk == nil {
		return
	}
	var receivable mysql_model.Receivable
	recID, err := strconv.Atoi(ctx.PostForm("receivable_id"))
	if err != nil {
		common.RequestParamErrorTemplate(ctx, common.REQUEST_PARM_ERROR)
		return
	}
	receivable.RecID = &recID
	err = bk.MysqlClient.WithContext(context.WithValue(context.Background(), "book_name", n)).
		Delete(&receivable).Error
	if err != nil {
		logger.Error(errors.WithStack(err), "删除应收记录失败!")
		common.InternalDataBaseErrorTemplate(ctx, common.DATABASE_DELETE_ERROR, receivable)
		return
	}
	ctx.JSON(_const.OK, errs.CreateSuccessMsg("删除应收记录成功!"))
	return
}
