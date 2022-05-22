package storage

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
	"log"
	"strconv"
)

// CreateCommodity 创建产品
func CreateCommodity(ctx *gin.Context) {
	bk, bookName := common.HarvestClientFromGinContext(ctx)
	if bk == nil {
		return
	}
	var commodity mysql_model.Commodity
	err := ctx.ShouldBindBodyWith(&commodity, binding.JSON)
	if err != nil {
		log.Println(err)
		common.RequestParamErrorTemplate(ctx, common.REQUEST_PARM_ERROR)
		return
	}

	err = bk.MysqlClient.WithContext(context.WithValue(context.Background(), "book_name", bookName)).
		Create(&commodity).Error
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
	err := ctx.ShouldBindBodyWith(&commodity, binding.JSON)
	if err != nil || commodity.RecID == nil {
		logger.Error(errors.WithStack(err), "绑定产品失败!")
		common.RequestParamErrorTemplate(ctx, common.REQUEST_PARM_ERROR)
		return
	}
	err = bk.MysqlClient.WithContext(context.WithValue(context.Background(), "book_name", bookName)).
		Updates(&commodity).Where("rec_id", *commodity.RecID).Error
	if err != nil {
		logger.Error(errors.WithStack(err), "更新产品失败!")
		common.InternalDataBaseErrorTemplate(ctx, common.DATABASE_UPDATE_ERROR, commodity)
		return
	}
	ctx.JSON(_const.OK, errs.CreateSuccessMsg("更新产品成功!"))
	return
}

// DeleteCommodity 删除产品
func DeleteCommodity(ctx *gin.Context) {
	bk, bookName := common.HarvestClientFromGinContext(ctx)
	if bk == nil {
		return
	}

	id, err := strconv.Atoi(ctx.PostForm("commodity_id"))
	if err != nil {
		common.RequestParamErrorTemplate(ctx, common.REQUEST_PARM_ERROR)
		return
	}
	commodity := mysql_model.Commodity{BasicModel: mysql_model.BasicModel{
		RecID: &id,
	}}
	err = bk.MysqlClient.WithContext(context.WithValue(context.Background(), "book_name", bookName)).
		Delete(&commodity).Error
	if err != nil {
		common.InternalDataBaseErrorTemplate(ctx, common.DATABASE_DELETE_ERROR, commodity)
		return
	}
	ctx.JSON(_const.OK, errs.CreateSuccessMsg("删除原材料成功!"))
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
		Create(&commodityBatch).Error
	if err != nil {
		common.InternalDataBaseErrorTemplate(ctx, common.DATABASE_INSERT_ERROR, commodityBatch)
		return
	}
	ctx.JSON(_const.OK, errs.CreateSuccessMsg("创建产品批次成功!"))
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

func UpdateCommodityDetail(ctx *gin.Context) {
	bk, bookName := common.HarvestClientFromGinContext(ctx)
	if bk == nil {
		return
	}
	var commodityDetail mysql_model.CommodityBatch
	err := ctx.ShouldBind(&commodityDetail)
	if err != nil || commodityDetail.RecID == nil {
		common.RequestParamErrorTemplate(ctx, common.REQUEST_PARM_ERROR)
		return
	}
	err = bk.MysqlClient.WithContext(context.WithValue(context.Background(), "book_name", bookName)).Updates(&commodityDetail).Where("rec_id = ?", commodityDetail.RecID).Error
	if err != nil {
		common.InternalDataBaseErrorTemplate(ctx, common.DATABASE_UPDATE_ERROR, commodityDetail)
		return
	}
	ctx.JSON(_const.OK, errs.CreateSuccessMsg("更新产品批次信息成功!"))
	return
}

// DeleteCommodityDetail 删除产品批次信息
func DeleteCommodityDetail(ctx *gin.Context) {
	bk, bookName := common.HarvestClientFromGinContext(ctx)
	if bk == nil {
		return
	}
	recID, err := strconv.Atoi(ctx.PostForm("commodity_batch_id"))
	if err != nil {
		common.RequestParamErrorTemplate(ctx, common.REQUEST_PARM_ERROR)
		return
	}

	commodityDetail := mysql_model.CommodityBatch{BasicModel: mysql_model.BasicModel{RecID: &recID}}
	err = bk.MysqlClient.
		WithContext(context.WithValue(context.Background(), "book_name", bookName)).
		Delete(&commodityDetail).Error
	if err != nil {
		logger.Error(errors.WithStack(err), "删除产品批次信息失败!")
		common.InternalDataBaseErrorTemplate(ctx, common.DATABASE_DELETE_ERROR, commodityDetail)
		return
	}
	ctx.JSON(_const.OK, errs.CreateSuccessMsg("删除产品批次成功!"))
	return
}
