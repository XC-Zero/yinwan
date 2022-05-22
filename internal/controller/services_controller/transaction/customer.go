package transaction

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

func CreateCustomer(ctx *gin.Context) {
	bk, bookName := common.HarvestClientFromGinContext(ctx)
	if bk == nil {
		return
	}
	var customer mysql_model.Customer
	err := ctx.ShouldBindBodyWith(&customer, binding.JSON)
	if err != nil {
		logger.Error(errors.WithStack(err), "")
		common.RequestParamErrorTemplate(ctx, common.REQUEST_PARM_ERROR)
		return
	}
	err = bk.MysqlClient.WithContext(context.WithValue(context.Background(), "book_name", bookName)).
		Create(&customer).Error
	if err != nil {
		common.InternalDataBaseErrorTemplate(ctx, common.DATABASE_INSERT_ERROR, customer)
		return
	}
	ctx.JSON(_const.OK, errs.CreateSuccessMsg("创建客户成功!"))
	return
}
func SelectCustomer(ctx *gin.Context) {
	bk, bookName := common.HarvestClientFromGinContext(ctx)
	if bk == nil {
		return
	}
	conditions := []common.MysqlCondition{
		{
			Symbol:      mysql.LIKE,
			ColumnName:  "customer_name",
			ColumnValue: ctx.PostForm("customer_name"),
		},
		{
			Symbol:      mysql.EQUAL,
			ColumnName:  "rec_id",
			ColumnValue: ctx.PostForm("customer_id"),
		},
		{
			Symbol:      mysql.EQUAL,
			ColumnName:  "customer_social_credit_code",
			ColumnValue: ctx.PostForm("customer_social_credit_code"),
		},
		{
			Symbol:      mysql.NULL,
			ColumnName:  "deleted_at",
			ColumnValue: " ",
		},
	}
	op := common.SelectMysqlTemplateOptions{
		DB:         bk.MysqlClient.WithContext(context.WithValue(context.Background(), "book_name", bookName)),
		TableModel: mysql_model.Customer{},
	}
	common.SelectMysqlTableContentWithCountTemplate(ctx, op, conditions...)
	return
}

func UpdateCustomer(ctx *gin.Context) {
	bk, bookName := common.HarvestClientFromGinContext(ctx)
	if bk == nil {
		return
	}
	var customer mysql_model.Customer
	err := ctx.ShouldBindBodyWith(&customer, binding.JSON)
	if err != nil || customer.RecID == nil {
		common.RequestParamErrorTemplate(ctx, common.REQUEST_PARM_ERROR)
		return
	}
	err = bk.MysqlClient.WithContext(context.WithValue(context.Background(), "book_name", bookName)).
		Updates(&customer).Where("rec_id", *customer.RecID).Error
	if err != nil {
		common.InternalDataBaseErrorTemplate(ctx, common.DATABASE_UPDATE_ERROR, customer)
		return
	}
	ctx.JSON(_const.OK, errs.CreateSuccessMsg("修改客户成功!"))
	return
}
func DeleteCustomer(ctx *gin.Context) {
	bk, bookName := common.HarvestClientFromGinContext(ctx)
	if bk == nil {
		return
	}

	recID, err := strconv.Atoi(ctx.PostForm("customer_id"))
	if err != nil {
		common.RequestParamErrorTemplate(ctx, common.REQUEST_PARM_ERROR)
		return
	}
	var customer = mysql_model.Customer{BasicModel: mysql_model.BasicModel{
		RecID: &recID,
	}}

	err = bk.MysqlClient.WithContext(
		context.WithValue(context.Background(), "book_name", bookName)).Delete(&customer).Error
	if err != nil {
		logger.Error(errors.WithStack(err), "删除客户失败!")
		common.InternalDataBaseErrorTemplate(ctx, common.DATABASE_DELETE_ERROR, customer)
		return
	}
	ctx.JSON(_const.OK, errs.CreateSuccessMsg("删除客户成功!"))
	return
}
