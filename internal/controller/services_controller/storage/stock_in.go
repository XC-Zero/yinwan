package storage

import (
	"github.com/XC-Zero/yinwan/internal/controller/services_controller/common"
	"github.com/XC-Zero/yinwan/pkg/client"
	"github.com/XC-Zero/yinwan/pkg/model"
	"github.com/gin-gonic/gin"
)

func CreateStockIn(ctx *gin.Context) {
	temp := model.StockInRecord{}
	bk := client.HarvestClientFromGinContext(ctx)
	if bk == nil {
		common.RequestParamErrorTemplate(ctx, common.BOOK_NAME_LACK_ERROR)
		return
	}
	err := ctx.ShouldBind(&temp)
	if err != nil {
		return
	}
}

func SelectStockIn(ctx *gin.Context) {
	bk := client.HarvestClientFromGinContext(ctx)
	if bk == nil {
		common.RequestParamErrorTemplate(ctx, common.BOOK_NAME_LACK_ERROR)
		return
	}
}
func UpdateStockIn(ctx *gin.Context) {
	bk := client.HarvestClientFromGinContext(ctx)
	if bk == nil {
		common.RequestParamErrorTemplate(ctx, common.BOOK_NAME_LACK_ERROR)
		return
	}
}
func DeleteStockIn(ctx *gin.Context) {
	bk := client.HarvestClientFromGinContext(ctx)
	if bk == nil {
		common.RequestParamErrorTemplate(ctx, common.BOOK_NAME_LACK_ERROR)
		return
	}
}
