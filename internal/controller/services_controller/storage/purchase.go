package storage

import (
	"context"
	"github.com/XC-Zero/yinwan/internal/controller/services_controller/common"
	"github.com/XC-Zero/yinwan/pkg/model/mongo_model"
	"github.com/gin-gonic/gin"
	"strconv"
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
	bk, n := common.HarvestClientFromGinContext(ctx)
	if bk == nil {
		return
	}
	var stockOutRecord mongo_model.Purchase
	recID, err := strconv.Atoi(ctx.PostForm("purchase_id"))
	if err != nil {
		common.RequestParamErrorTemplate(ctx, common.REQUEST_PARM_ERROR)
		return
	}
	common.DeleteOneMongoDBRecordByIDTemplate(ctx, common.MongoDBTemplateOptions{
		DB:         bk.MongoDBClient,
		Context:    context.WithValue(context.Background(), "book_name", n),
		RecID:      recID,
		TableModel: stockOutRecord,
		PreFunc:    nil,
	})
	return
}
