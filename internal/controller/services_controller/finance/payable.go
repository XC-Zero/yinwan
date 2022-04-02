package finance

import (
	"github.com/XC-Zero/yinwan/internal/controller/services_controller/common"
	"github.com/XC-Zero/yinwan/pkg/client"
	"github.com/gin-gonic/gin"
)

func CreatePayable(ctx *gin.Context) {

}

func SelectPayable(ctx *gin.Context) {
	bk := client.HarvestClientFromGinContext(ctx)
	if bk == nil {
		ctx.JSON("")
		return
	}
	condition := []common.Condition{
		{},
	}
	common.SelectTableContentWithCountMongoDBTemplate(ctx)
}

func UpdatePayable(ctx *gin.Context) {

}

func DeletePayable(ctx *gin.Context) {

}
