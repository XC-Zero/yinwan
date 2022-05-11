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
)

// CreateCommodity 创建产品
func CreateCommodity(ctx *gin.Context) {
	bk, bookName := common.HarvestClientFromGinContext(ctx)
	if bk == nil {
		return
	}
	var commodity mysql_model.Commodity
	err := ctx.ShouldBind(&commodity)
	if err != nil {
		log.Println(err)
		common.RequestParamErrorTemplate(ctx, common.REQUEST_PARM_ERROR)
		return
	}

	err = bk.MysqlClient.WithContext(context.WithValue(context.Background(), "book_name", bookName)).Create(&commodity).Error
	if err != nil {
		log.Println(err)
		common.InternalDataBaseErrorTemplate(ctx, common.DATABASE_INSERT_ERROR, commodity)
		return
	}
	ctx.JSON(_const.OK, errs.CreateSuccessMsg("创建产品成功！"))
	return

}

// SelectCommodity 产品
func SelectCommodity(ctx *gin.Context) {
	bk, bookName := common.HarvestClientFromGinContext(ctx)
	if bk == nil {
		return
	}
	conditions := []common.MysqlCondition{
		{
			Symbol:      mysql.LIKE,
			ColumnName:  "commodity_name",
			ColumnValue: ctx.PostForm("commodity_name"),
		},
		{
			Symbol:      mysql.EQUAL,
			ColumnName:  "rec_id",
			ColumnValue: ctx.PostForm("commodity_id"),
		},
		{
			Symbol:      mysql.EQUAL,
			ColumnName:  "commodity_type_id",
			ColumnValue: ctx.PostForm("commodity_type_id"),
		},
		{
			Symbol:      mysql.NULL,
			ColumnName:  "deleted_at",
			ColumnValue: " ",
		},
	}
	op := common.SelectMysqlTemplateOptions{
		DB:         bk.MysqlClient.WithContext(context.WithValue(context.Background(), "book_name", bookName)),
		TableModel: mysql_model.Commodity{},
	}
	common.SelectMysqlTableContentWithCountTemplate(ctx, op, conditions...)

	return
}

func UpdateCommodity(ctx *gin.Context) {
	bk, bookName := common.HarvestClientFromGinContext(ctx)
	if bk == nil {
		return
	}
	commodity := mysql_model.Commodity{}
	err := ctx.ShouldBind(&commodity)
	if err != nil {
		common.RequestParamErrorTemplate(ctx, common.REQUEST_PARM_ERROR)
		return
	}
	err = bk.MysqlClient.WithContext(context.WithValue(context.Background(), "book_name", bookName)).
		Updates(commodity).Error
	if err != nil {
		common.InternalDataBaseErrorTemplate(ctx, common.DATABASE_UPDATE_ERROR, commodity)
		return
	}
	return
}

// DeleteCommodity 删除产品
func DeleteCommodity(ctx *gin.Context) {
	bk, bookName := common.HarvestClientFromGinContext(ctx)
	if bk == nil {
		return
	}
	commodity := mysql_model.Commodity{}
	id := ctx.PostForm("commodity_id")
	if id == "" {
		common.RequestParamErrorTemplate(ctx, common.REQUEST_PARM_ERROR)
		return
	}
	err := bk.MysqlClient.WithContext(context.WithValue(context.Background(), "book_name", bookName)).
		Delete(commodity, id).Error
	if err != nil {
		common.InternalDataBaseErrorTemplate(ctx, common.DATABASE_DELETE_ERROR, commodity)
		return
	}
	return
}

// CreateCommodityBatch  创建产品批次
//	(预留的,正常不应该有的,正常都应该走入库)
func CreateCommodityBatch(ctx *gin.Context) {
	bk, bookName := common.HarvestClientFromGinContext(ctx)
	if bk == nil {
		return
	}

	var commodityBatch mysql_model.CommodityBatch
	err := ctx.ShouldBind(&commodityBatch)
	if err != nil {
		common.RequestParamErrorTemplate(ctx, common.REQUEST_PARM_ERROR)
		return
	}
	err = bk.MysqlClient.WithContext(context.WithValue(context.Background(), "book_name", bookName)).
		Create(commodityBatch).Error
	if err != nil {
		common.InternalDataBaseErrorTemplate(ctx, common.DATABASE_INSERT_ERROR, commodityBatch)
		return
	}
	return
}

// SelectCommodityDetail 产品批次信息
func SelectCommodityDetail(ctx *gin.Context) {
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
			ColumnName:  "commodity_id",
			ColumnValue: ctx.PostForm("commodity_id"),
		},
		{
			Symbol:      mysql.EQUAL,
			ColumnName:  "commodity_type_id",
			ColumnValue: ctx.PostForm("commodity_type_id"),
		},
		{
			Symbol:      mysql.NULL,
			ColumnName:  "deleted_at",
			ColumnValue: " ",
		},
	}
	op := common.SelectMysqlTemplateOptions{
		DB:         bk.MysqlClient.WithContext(context.WithValue(context.Background(), "book_name", bookName)),
		TableModel: mysql_model.CommodityBatch{},
	}

	common.SelectMysqlTableContentWithCountTemplate(ctx, op, conditions...)

	return
}

// DeleteCommodityDetail 删除产品批次信息
func DeleteCommodityDetail(ctx *gin.Context) {

}
