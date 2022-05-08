package storage

import (
	"github.com/gin-gonic/gin"
)

func CreatePurchase(ctx *gin.Context) {
	//bk, bookName := common.HarvestClientFromGinContext(ctx)
	//if bk == nil {
	//	return
	//}
	//var purchase mongo_model.Purchase
	//
	//err := ctx.ShouldBind(&purchase)
	//if err != nil {
	//	common.RequestParamErrorTemplate(ctx, common.REQUEST_PARM_ERROR)
	//	return
	//}
	//op := common.CreateMongoDBTemplateOptions{
	//	DB:         client.MongoDBClient,
	//	TableModel: mongo_model.Purchase{},
	//}
	//common.CreateOneMongoDBRecordTemplate(ctx, op)
}
func SelectPurchase(ctx *gin.Context) {
	//bk, bookName := common.HarvestClientFromGinContext(ctx)
	//if bk == nil {
	//	return
	//}

}
func UpdatePurchase(ctx *gin.Context) {
	//bk, bookName := common.HarvestClientFromGinContext(ctx)
	//if bk == nil {
	//	return
	//}

}
func DeletePurchase(ctx *gin.Context) {
	//bk, bookName := common.HarvestClientFromGinContext(ctx)
	//if bk == nil {
	//	return
	//}

}
