package transaction

import (
	"context"
	"github.com/XC-Zero/yinwan/internal/controller/services_controller/common"
	"github.com/XC-Zero/yinwan/pkg/model/mysql_model"
	"github.com/XC-Zero/yinwan/pkg/utils/mysql"
	"github.com/gin-gonic/gin"
)

func CreateCustomer(ctx *gin.Context) {
	bk, bookName := common.HarvestClientFromGinContext(ctx)
	if bk == nil {
		return
	}
	var customer mysql_model.Customer
	err := ctx.ShouldBind(&customer)
	if err != nil {
		common.RequestParamErrorTemplate(ctx, common.REQUEST_PARM_ERROR)
		return
	}
	op := common.CreateMysqlTemplateOptions{
		DB:         bk.MysqlClient.WithContext(context.WithValue(context.Background(), "book_name", bookName)),
		TableModel: customer,
	}
	common.CreateOneMysqlRecordTemplate(ctx, op)
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
	err := ctx.ShouldBind(&customer)
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
	return
}
func DeleteCustomer(ctx *gin.Context) {
	bk, bookName := common.HarvestClientFromGinContext(ctx)
	if bk == nil {
		return
	}
	var customer mysql_model.Customer
	var recID = ctx.PostForm("customer_id")
	if recID == "" {
		common.RequestParamErrorTemplate(ctx, common.REQUEST_PARM_ERROR)
		return
	}
	err := bk.MysqlClient.WithContext(
		context.WithValue(context.Background(), "book_name", bookName)).
		Delete(&customer, recID).Error
	if err != nil {
		common.InternalDataBaseErrorTemplate(ctx, common.DATABASE_DELETE_ERROR, customer)
		return
	}
	return
}
