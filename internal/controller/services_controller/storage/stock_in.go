package storage

import (
	"context"
	"github.com/XC-Zero/yinwan/internal/controller/services_controller/common"
	"github.com/XC-Zero/yinwan/pkg/model/mongo_model"
	my_mongo "github.com/XC-Zero/yinwan/pkg/utils/mongo"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/bsontype"
)

func CreateStockIn(ctx *gin.Context) {
	bk, n := common.HarvestClientFromGinContext(ctx)
	if bk == nil {
		return
	}
	temp := mongo_model.StockInRecord{}

	err := ctx.ShouldBind(&temp)
	if err != nil {
		common.RequestParamErrorTemplate(ctx, common.REQUEST_PARM_ERROR)
		return
	}
	common.CreateOneMongoDBRecordTemplate(ctx, common.CreateMongoDBTemplateOptions{
		DB:         bk.MongoDBClient,
		Context:    context.WithValue(context.Background(), "book_name", n),
		TableModel: temp,
		PreFunc:    nil,
	})
	return
}

func SelectStockIn(ctx *gin.Context) {
	bk, _ := common.HarvestClientFromGinContext(ctx)
	if bk == nil {
		return
	}

	conditions := []common.MongoCondition{
		{
			Symbol:      my_mongo.EQUAL,
			ColumnName:  "rec_id",
			ColumnValue: ctx.PostForm("stock_in_record_id"),
		},
		{
			Symbol:      my_mongo.EQUAL,
			ColumnName:  "rec_id",
			ColumnValue: ctx.PostForm("stock_in_record_id"),
		},
		{
			Symbol:      my_mongo.NOT_EQUAL,
			ColumnName:  "deleted_at",
			ColumnValue: bsontype.Null,
		},
	}
	options := common.SelectMongoDBTemplateOptions{
		DB:         bk.MongoDBClient,
		TableModel: mongo_model.StockInRecord{},
	}
	common.SelectMongoDBTableContentWithCountTemplate(ctx, options, conditions...)
	return

}
func UpdateStockIn(ctx *gin.Context) {
	bk, n := common.HarvestClientFromGinContext(ctx)
	if bk == nil {
		return
	}

	temp := mongo_model.StockInRecord{}

	err := ctx.ShouldBind(&temp)
	if err != nil || temp.RecID == nil {
		common.RequestParamErrorTemplate(ctx, common.REQUEST_PARM_ERROR)
		return
	}
	common.UpdateOneMongoDBRecordByIDTemplate(ctx, common.MongoDBTemplateOptions{
		bk.MongoDBClient,
		context.WithValue(context.Background(), "book_name", n),
		*temp.RecID,
		temp,
		nil,
	})
	return
}
func DeleteStockIn(ctx *gin.Context) {
	bk, n := common.HarvestClientFromGinContext(ctx)
	if bk == nil {
		return
	}

	temp := mongo_model.StockInRecord{}

	err := ctx.ShouldBind(&temp)
	if err != nil || temp.RecID == nil {
		common.RequestParamErrorTemplate(ctx, common.REQUEST_PARM_ERROR)
		return
	}
	common.DeleteOneMongoDBRecordByIDTemplate(ctx, common.MongoDBTemplateOptions{
		DB:         bk.MongoDBClient,
		Context:    context.WithValue(context.Background(), "book_name", n),
		RecID:      *temp.RecID,
		TableModel: temp,
		PreFunc:    nil,
	})
	return
}
