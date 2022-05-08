package finance

import (
	"github.com/XC-Zero/yinwan/internal/controller/services_controller/common"
	"github.com/XC-Zero/yinwan/pkg/client"
	"github.com/XC-Zero/yinwan/pkg/model/mysql_model"
	"github.com/gin-gonic/gin"
)

func CreatePayable(ctx *gin.Context) {
	bk, _ := common.HarvestClientFromGinContext(ctx)
	if bk == nil {
		return
	}
}

func SelectPayable(ctx *gin.Context) {
	bk, _ := common.HarvestClientFromGinContext(ctx)
	if bk == nil {
		return
	}
	condition := []common.MongoCondition{
		{},
	}
	op := common.SelectMongoDBTemplateOptions{
		DB:         client.MongoDBClient,
		TableModel: mysql_model.Payable{},
	}
	common.SelectMongoDBTableContentWithCountTemplate(ctx, op, condition...)
}

func UpdatePayable(ctx *gin.Context) {
	bk, _ := common.HarvestClientFromGinContext(ctx)
	if bk == nil {
		return
	}
}

func DeletePayable(ctx *gin.Context) {
	bk, _ := common.HarvestClientFromGinContext(ctx)
	if bk == nil {
		return
	}
}
