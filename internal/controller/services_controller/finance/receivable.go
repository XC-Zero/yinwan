package finance

import (
	"github.com/XC-Zero/yinwan/internal/controller/services_controller/common"
	"github.com/gin-gonic/gin"
)

func CreateReceivable(ctx *gin.Context) {
	bk, _ := common.HarvestClientFromGinContext(ctx)
	if bk == nil {
		return
	}
}

func SelectReceivable(ctx *gin.Context) {
	bk, _ := common.HarvestClientFromGinContext(ctx)
	if bk == nil {
		return
	}
}

func UpdateReceivable(ctx *gin.Context) {
	bk, _ := common.HarvestClientFromGinContext(ctx)
	if bk == nil {
		return
	}
}

func DeleteReceivable(ctx *gin.Context) {
	bk, _ := common.HarvestClientFromGinContext(ctx)
	if bk == nil {
		return
	}
}
