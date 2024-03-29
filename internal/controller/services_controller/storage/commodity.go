package storage

import (
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

// CreateCommodity 创建产品
func CreateCommodity(ctx *gin.Context) {
	bk := common.HarvestClientFromGinContext(ctx)
	if bk == nil {
		return
	}
	var commodity mysql_model.Commodity
	err := ctx.ShouldBindBodyWith(&commodity, binding.JSON)
	if err != nil {
		logger.Error(errors.WithStack(err), "")
		common.RequestParamErrorTemplate(ctx, common.REQUEST_PARM_ERROR)
		return
	}

	err = bk.MysqlClient.WithContext(ctx).
		Create(&commodity).Error
	if err != nil {
		logger.Error(errors.WithStack(err), "")
		common.InternalDataBaseErrorTemplate(ctx, common.DATABASE_INSERT_ERROR, commodity)
		return
	}
	ctx.JSON(_const.OK, errs.CreateSuccessMsg("创建产品成功！"))
	return

}

// SelectCommodity 产品
func SelectCommodity(ctx *gin.Context) {
	bk := common.HarvestClientFromGinContext(ctx)
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
		DB:         bk.MysqlClient.WithContext(ctx),
		TableModel: mysql_model.Commodity{},
	}
	common.SelectMysqlTableContentWithCountTemplate(ctx, op, conditions...)

	return
}

func UpdateCommodity(ctx *gin.Context) {
	bk := common.HarvestClientFromGinContext(ctx)
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
	err = bk.MysqlClient.WithContext(ctx).Where("rec_id", *commodity.RecID).
		Updates(&commodity).Error
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
	bk := common.HarvestClientFromGinContext(ctx)
	if bk == nil {
		return
	}

	id, err := strconv.Atoi(ctx.PostForm("commodity_id"))
	if err != nil {
		logger.Error(errors.WithStack(err), "")
		common.RequestParamErrorTemplate(ctx, common.REQUEST_PARM_ERROR)
		return
	}
	commodity := mysql_model.Commodity{BasicModel: mysql_model.BasicModel{
		RecID: &id,
	}}
	err = bk.MysqlClient.WithContext(ctx).
		Delete(&commodity).Error
	if err != nil {
		logger.Error(errors.WithStack(err), "")
		common.InternalDataBaseErrorTemplate(ctx, common.DATABASE_DELETE_ERROR, commodity)
		return
	}
	ctx.JSON(_const.OK, errs.CreateSuccessMsg("删除原材料成功!"))
	return
}

// CreateCommodityBatch  创建产品批次
//	(预留的,正常不应该有的,正常都应该走入库)
func CreateCommodityBatch(ctx *gin.Context) {
	bk := common.HarvestClientFromGinContext(ctx)
	if bk == nil {
		return
	}

	var commodityBatch mysql_model.CommodityBatch
	err := ctx.ShouldBindBodyWith(&commodityBatch, binding.JSON)
	if err != nil {
		logger.Error(errors.WithStack(err), "")
		common.RequestParamErrorTemplate(ctx, common.REQUEST_PARM_ERROR)
		return
	}
	err = bk.MysqlClient.WithContext(ctx).
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
	bk := common.HarvestClientFromGinContext(ctx)
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
			mysql.GREATER_THEN,
			"commodity_batch_surplus_number",
			ctx.PostForm("commodity_batch_surplus_number"),
		},
		{
			Symbol:      mysql.NULL,
			ColumnName:  "deleted_at",
			ColumnValue: " ",
		},
	}
	op := common.SelectMysqlTemplateOptions{
		DB:         bk.MysqlClient.WithContext(ctx),
		TableModel: mysql_model.CommodityBatch{},
	}

	common.SelectMysqlTableContentWithCountTemplate(ctx, op, conditions...)

	return
}

func UpdateCommodityDetail(ctx *gin.Context) {
	bk := common.HarvestClientFromGinContext(ctx)
	if bk == nil {
		return
	}
	var commodityDetail mysql_model.CommodityBatch
	err := ctx.ShouldBindBodyWith(&commodityDetail, binding.JSON)
	if err != nil || commodityDetail.RecID == nil {
		logger.Error(errors.WithStack(err), "")
		common.RequestParamErrorTemplate(ctx, common.REQUEST_PARM_ERROR)
		return
	}
	err = bk.MysqlClient.WithContext(ctx).Where("rec_id = ?", commodityDetail.RecID).Updates(&commodityDetail).Error
	if err != nil {
		logger.Error(errors.WithStack(err), "")
		common.InternalDataBaseErrorTemplate(ctx, common.DATABASE_UPDATE_ERROR, commodityDetail)
		return
	}
	ctx.JSON(_const.OK, errs.CreateSuccessMsg("更新产品批次信息成功!"))
	return
}

// DeleteCommodityDetail 删除产品批次信息
func DeleteCommodityDetail(ctx *gin.Context) {
	bk := common.HarvestClientFromGinContext(ctx)
	if bk == nil {
		return
	}
	recID, err := strconv.Atoi(ctx.PostForm("commodity_batch_id"))
	if err != nil {
		logger.Error(errors.WithStack(err), "")
		common.RequestParamErrorTemplate(ctx, common.REQUEST_PARM_ERROR)
		return
	}

	commodityDetail := mysql_model.CommodityBatch{BasicModel: mysql_model.BasicModel{RecID: &recID}}
	err = bk.MysqlClient.
		WithContext(ctx).
		Delete(&commodityDetail).Error
	if err != nil {
		logger.Error(errors.WithStack(err), "删除产品批次信息失败!")
		common.InternalDataBaseErrorTemplate(ctx, common.DATABASE_DELETE_ERROR, commodityDetail)
		return
	}
	ctx.JSON(_const.OK, errs.CreateSuccessMsg("删除产品批次成功!"))
	return
}

// SelectCommodityHistoryCost 详情页历史均价
func SelectCommodityHistoryCost(ctx *gin.Context) {
	bk := common.HarvestClientFromGinContext(ctx)
	if bk == nil {
		return
	}
	commodityID := ctx.PostForm("commodity_id")
	if commodityID == "" {
		common.RequestParamErrorTemplate(ctx, common.REQUEST_PARM_ERROR)
		return
	}
	conditions := []common.MysqlCondition{
		{
			Symbol:      mysql.EQUAL,
			ColumnName:  "commodity_id",
			ColumnValue: commodityID,
		},
		{
			Symbol:      mysql.NULL,
			ColumnName:  "deleted_at",
			ColumnValue: " ",
		},
	}
	res := common.SelectMysqlTableContentWithCountTemplate(ctx, common.SelectMysqlTemplateOptions{
		DB:            bk.MysqlClient.WithContext(ctx),
		OrderByColumn: "created_at",
		TableModel:    mysql_model.CommodityHistoricalCost{},
		NotReturn:     true,
		NotPaginate:   true,
	}, conditions...)
	if res == nil {
		common.InternalDataBaseErrorTemplate(ctx, common.DATABASE_SELECT_ERROR, mysql_model.CommodityHistoricalCost{})
		return
	}
	costs := res.([]mysql_model.CommodityHistoricalCost)
	var count int64 = 0
	var errorList []error
	var totalPrice = math_plus.NewFromFloatByDecimal(0.0, 1)
	var dataList = make([]mysql_model.CommodityHistoricalCost, 0, 15)
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
