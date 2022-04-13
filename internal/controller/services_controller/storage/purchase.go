package storage

import (
	"github.com/XC-Zero/yinwan/internal/controller/services_controller/common"
	"github.com/XC-Zero/yinwan/pkg/client"
	"github.com/XC-Zero/yinwan/pkg/model"
	"github.com/gin-gonic/gin"
)

func CreatePurchase(ctx *gin.Context) {
	var purchase model.Purchase
	bk := client.HarvestClientFromGinContext(ctx)
	if bk == nil {
		common.RequestParamErrorTemplate(ctx, common.BOOK_NAME_LACK_ERROR)
		return
	}
	err := ctx.ShouldBind(&purchase)
	if err != nil {
		common.RequestParamErrorTemplate(ctx, common.REQUEST_PARM_ERROR)
		return
	}
	op := common.CreateMongoDBTemplateOptions{
		DB:         bk.MongoDBClient,
		TableModel: model.Purchase{},
	}
	common.CreateOneMongoDBRecordTemplate(ctx, op)
}
func SelectPurchase(ctx *gin.Context) {
	bk := client.HarvestClientFromGinContext(ctx)
	if bk == nil {
		common.RequestParamErrorTemplate(ctx, common.BOOK_NAME_LACK_ERROR)
		return
	}
}
func UpdatePurchase(ctx *gin.Context) {
	bk := client.HarvestClientFromGinContext(ctx)
	if bk == nil {
		common.RequestParamErrorTemplate(ctx, common.BOOK_NAME_LACK_ERROR)
		return
	}
}
func DeletePurchase(ctx *gin.Context) {
	bk := client.HarvestClientFromGinContext(ctx)
	if bk == nil {
		common.RequestParamErrorTemplate(ctx, common.BOOK_NAME_LACK_ERROR)
		return
	}
}
