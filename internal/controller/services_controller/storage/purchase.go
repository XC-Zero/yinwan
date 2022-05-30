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
			Symbol:      my_mongo.EQUAL,
			ColumnName:  "basicmodel.rec_id",
			ColumnValue: ctx.PostForm("stock_in_record_id"),
		},
		{
			Symbol:      my_mongo.EQUAL,
			ColumnName:  "basicmodel.deleted_at",
			ColumnValue: "null",
		},
		{
			Symbol:      my_mongo.EQUAL,
			ColumnName:  "provider_id",
			ColumnValue: ctx.PostForm("provider_id"),
		},
		{
			Symbol:      my_mongo.EQUAL,
			ColumnName:  "stock_in_warehouse_id",
			ColumnValue: ctx.PostForm("stock_in_warehouse_id"),
		},
		{
			Symbol:      my_mongo.EQUAL,
			ColumnName:  "stock_in_record_owner_id",
			ColumnValue: ctx.PostForm("stock_in_record_owner_id"),
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
