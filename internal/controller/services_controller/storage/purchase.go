package storage

import (
	"context"
	"github.com/XC-Zero/yinwan/internal/controller/services_controller/common"
	"github.com/XC-Zero/yinwan/pkg/client"
	"github.com/XC-Zero/yinwan/pkg/model/mongo_model"
	my_mongo "github.com/XC-Zero/yinwan/pkg/utils/mongo"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"go.mongodb.org/mongo-driver/bson/bsontype"
	"strconv"
	"time"
)

func CreatePurchase(ctx *gin.Context) {
	bk, n := common.HarvestClientFromGinContext(ctx)
	if bk == nil {
		return
	}
	var auto common.Auto
	ctx.ShouldBindBodyWith(&auto, binding.JSON)
	var purchase mongo_model.Purchase
	err := ctx.ShouldBindBodyWith(&purchase, binding.JSON)
	if err != nil {
		common.RequestParamErrorTemplate(ctx, common.REQUEST_PARM_ERROR)
		return
	}
	recID := int(time.Now().Unix())
	purchase.RecID = &recID
	purchase.BookName = n
	purchase.BookNameID = bk.StorageName
	op := common.CreateMongoDBTemplateOptions{
		DB:         client.MongoDBClient,
		Context:    auto.WithContext(context.WithValue(context.Background(), "book_name", n)),
		TableModel: purchase,
	}
	common.CreateOneMongoDBRecordTemplate(ctx, op)
	return
}
func SelectPurchase(ctx *gin.Context) {
	bk, _ := common.HarvestClientFromGinContext(ctx)
	if bk == nil {
		return
	}

	conditions := []common.MongoCondition{
		{
			my_mongo.EQUAL,
			"rec_id",
			ctx.PostForm("stock_in_record_id"),
		},
		{
			my_mongo.EQUAL,
			"deleted_at",
			bsontype.Null,
		},
		{
			my_mongo.EQUAL,
			"provider_id",
			ctx.PostForm("provider_id"),
		},
		{
			my_mongo.EQUAL,
			"stock_in_warehouse_id",
			ctx.PostForm("stock_in_warehouse_id"),
		},
		{
			my_mongo.EQUAL,
			"stock_in_record_owner_id",
			ctx.PostForm("stock_in_record_owner_id"),
		},
	}
	options := common.SelectMongoDBTemplateOptions{
		DB:         bk.MongoDBClient,
		TableModel: mongo_model.Purchase{},
	}
	common.SelectMongoDBTableContentWithCountTemplate(ctx, options, conditions...)
	return

}
func UpdatePurchase(ctx *gin.Context) {
	bk, n := common.HarvestClientFromGinContext(ctx)
	if bk == nil {
		return
	}
	var auto common.Auto
	ctx.ShouldBindBodyWith(&auto, binding.JSON)
	temp := mongo_model.Purchase{}
	err := ctx.ShouldBindBodyWith(&temp, binding.JSON)
	if err != nil || temp.RecID == nil {
		common.RequestParamErrorTemplate(ctx, common.REQUEST_PARM_ERROR)
		return
	}
	common.UpdateOneMongoDBRecordByIDTemplate(ctx, common.MongoDBTemplateOptions{
		DB:         bk.MongoDBClient,
		Context:    auto.WithContext(context.WithValue(context.Background(), "book_name", n)),
		RecID:      *temp.RecID,
		TableModel: temp,
	})
	return

}
func DeletePurchase(ctx *gin.Context) {
	bk, n := common.HarvestClientFromGinContext(ctx)
	if bk == nil {
		return
	}
	var auto common.Auto
	ctx.ShouldBindBodyWith(&auto, binding.JSON)

	var stockOutRecord mongo_model.Purchase

	recID, err := strconv.Atoi(ctx.PostForm("purchase_id"))
	if err != nil {
		common.RequestParamErrorTemplate(ctx, common.REQUEST_PARM_ERROR)
		return
	}
	stockOutRecord.RecID = &recID
	common.DeleteOneMongoDBRecordByIDTemplate(ctx, common.MongoDBTemplateOptions{
		DB:         bk.MongoDBClient,
		Context:    auto.WithContext(context.WithValue(context.Background(), "book_name", n)),
		RecID:      recID,
		TableModel: stockOutRecord,
	})
	return
}
