package storage

import (
	"context"
	"github.com/XC-Zero/yinwan/internal/controller/services_controller/common"
	"github.com/XC-Zero/yinwan/pkg/client"
	"github.com/XC-Zero/yinwan/pkg/model/mongo_model"
	my_mongo "github.com/XC-Zero/yinwan/pkg/utils/mongo"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"strconv"
	"time"
)

func CreatePurchase(ctx *gin.Context) {
	bk, bookName := common.HarvestClientFromGinContext(ctx)
	if bk == nil {
		return
	}
	var purchase mongo_model.Purchase
	err := ctx.ShouldBindBodyWith(&purchase, binding.JSON)
	if err != nil {
		common.RequestParamErrorTemplate(ctx, common.REQUEST_PARM_ERROR)
		return
	}
	recID := int(time.Now().Unix())
	purchase.RecID = &recID
	purchase.BookName = bookName
	purchase.BookNameID = bk.StorageName
	op := common.CreateMongoDBTemplateOptions{
		DB:         client.MongoDBClient,
		Context:    context.WithValue(context.Background(), "book_name", bookName),
		TableModel: mongo_model.Purchase{},
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
			nil,
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

	temp := mongo_model.Purchase{}

	err := ctx.ShouldBindBodyWith(&temp, binding.JSON)
	if err != nil || temp.RecID == nil {
		common.RequestParamErrorTemplate(ctx, common.REQUEST_PARM_ERROR)
		return
	}
	common.UpdateOneMongoDBRecordByIDTemplate(ctx, common.MongoDBTemplateOptions{
		DB:         bk.MongoDBClient,
		Context:    context.WithValue(context.Background(), "book_name", n),
		RecID:      *temp.RecID,
		TableModel: temp,
		PreFunc:    nil,
	})

	return

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
