package finance

import (
	"github.com/XC-Zero/yinwan/internal/controller/services_controller/common"
	"github.com/gin-gonic/gin"
)

// 报表

// CreateFixedAssetStatement 生成固定资产报表
func CreateFixedAssetStatement(ctx *gin.Context) {
	bk, n := common.HarvestClientFromGinContext(ctx)
	if bk == nil {
		return
	}
}

// CreateCashFlowStatement 生成现金流量表
func CreateCashFlowStatement(ctx *gin.Context) {
	bk, n := common.HarvestClientFromGinContext(ctx)
	if bk == nil {
		return
	}
}

// CreateBalanceStatement 生成资产负债表
func CreateBalanceStatement(ctx *gin.Context) {
	bk, n := common.HarvestClientFromGinContext(ctx)
	if bk == nil {
		return
	}
}

// CreateIncomeStatement 生成现金流量表
func CreateIncomeStatement(ctx *gin.Context) {
	bk, n := common.HarvestClientFromGinContext(ctx)
	if bk == nil {
		return
	}
}
