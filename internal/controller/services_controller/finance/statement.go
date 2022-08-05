package finance

import (
	"github.com/XC-Zero/yinwan/internal/controller/services_controller/common"
	_const "github.com/XC-Zero/yinwan/pkg/const"
	"github.com/XC-Zero/yinwan/pkg/utils/errs"
	"github.com/XC-Zero/yinwan/pkg/utils/logger"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
)

// 报表

// CreateFixedAssetStatement 生成固定资产报表
func CreateFixedAssetStatement(ctx *gin.Context) {
	bk := common.HarvestClientFromGinContext(ctx)
	if bk == nil {
		return
	}
}

// CreateCashFlowStatement 生成现金流量表
func CreateCashFlowStatement(ctx *gin.Context) {
	bk := common.HarvestClientFromGinContext(ctx)
	if bk == nil {
		return
	}
}

// CreateBalanceStatement 生成资产负债表
func CreateBalanceStatement(ctx *gin.Context) {
	bk := common.HarvestClientFromGinContext(ctx)
	if bk == nil {
		return
	}
}

// CreateIncomeStatement 生成现金流量表
func CreateIncomeStatement(ctx *gin.Context) {
	bk := common.HarvestClientFromGinContext(ctx)
	if bk == nil {
		return
	}
}

func GenerateMaterialStatistics(ctx *gin.Context) {
	bk := common.HarvestClientFromGinContext(ctx)
	if bk == nil {
		return
	}
	var temp []struct {
		MaterialName         string `json:"material_name"`
		RecID                int    `json:"rec_id"`
		MaterialPresentCount int    `json:"material_present_count"`
	}
	n := ctx.PostForm("limit")
	if n == "" {
		n = "5"
	}
	err := bk.MysqlClient.Raw(`
		select material_name, rec_id, material_present_count
			from materials
			where deleted_at is null
			  and material_present_count > 0
			order by material_present_count
			limit ? `, n).Scan(&temp).Error
	if err != nil {
		logger.Error(errors.WithStack(err), "查询前"+n+"原材料失败!")
		ctx.JSON(_const.INTERNAL_ERROR, errs.CreateWebErrorMsg("生成报表失败!"))
		return
	}
	var count int64
	err = bk.MysqlClient.Raw(`
		select sum(material_present_count)
			from materials
			where material_present_count > 0
			and deleted_at is null `).Scan(&count).Error
	if err != nil {
		logger.Error(errors.WithStack(err), "查询前原材料总数!")
		ctx.JSON(_const.INTERNAL_ERROR, errs.CreateWebErrorMsg("生成报表失败!"))
		return
	}
	common.SelectSuccessTemplate(ctx, count, temp)
	return

}

func GenerateCommodityStatistics(ctx *gin.Context) {
	bk := common.HarvestClientFromGinContext(ctx)
	if bk == nil {
		return
	}
	var temp []struct {
		CommodityName        string `json:"commodity_name"`
		RecID                int    `json:"rec_id"`
		MaterialPresentCount int    `json:"material_present_count"`
	}
	n := ctx.PostForm("limit")
	if n == "" {
		n = "5"
	}
	err := bk.MysqlClient.Raw(`
		select material_name, rec_id, material_present_count
			from materials
			where deleted_at is null
			  and material_present_count > 0
			order by material_present_count
			limit ? `, n).Scan(&temp).Error
	if err != nil {
		logger.Error(errors.WithStack(err), "查询前"+n+"原材料失败!")
		ctx.JSON(_const.INTERNAL_ERROR, errs.CreateWebErrorMsg("生成报表失败!"))
		return
	}
	var count int64
	err = bk.MysqlClient.Raw(`
		select sum(material_present_count)
			from materials
			where material_present_count > 0
			and deleted_at is null `).Scan(&count).Error
	if err != nil {
		logger.Error(errors.WithStack(err), "查询前原材料总数!")
		ctx.JSON(_const.INTERNAL_ERROR, errs.CreateWebErrorMsg("生成报表失败!"))
		return
	}
	common.SelectSuccessTemplate(ctx, count, temp)
	return

}

func GeneratePayableStatistics(ctx *gin.Context) {
	bk := common.HarvestClientFromGinContext(ctx)
	if bk == nil {
		return
	}
}

func GenerateReceivableStatistics(ctx *gin.Context) {
	bk := common.HarvestClientFromGinContext(ctx)
	if bk == nil {
		return
	}
}
