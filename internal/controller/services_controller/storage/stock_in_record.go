package storage

import (
	"context"
	"github.com/XC-Zero/yinwan/internal/controller/services_controller/common"
	_const "github.com/XC-Zero/yinwan/pkg/const"
	"github.com/XC-Zero/yinwan/pkg/model/mongo_model"
	"github.com/XC-Zero/yinwan/pkg/utils/logger"
	my_mongo "github.com/XC-Zero/yinwan/pkg/utils/mongo"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/pkg/errors"
	"strconv"
	"time"
)

// CreateStockIn 创建入库单
// 入库会同步触发新增原材料批次信息
func CreateStockIn(ctx *gin.Context) {
	bk, n := common.HarvestClientFromGinContext(ctx)
	if bk == nil {
		return
	}
	temp := mongo_model.StockInRecord{}

	err := ctx.ShouldBindBodyWith(&temp, binding.JSON)
	if err != nil {
		logger.Error(errors.WithStack(err), "")
		common.RequestParamErrorTemplate(ctx, common.REQUEST_PARM_ERROR)
		return
	}
	recID := int(time.Now().Unix())
	temp.RecID = &recID
	temp.BookName = n
	temp.BookNameID = bk.StorageName
	temp.CreatedAt = strconv.FormatInt(time.Now().Unix(), 10)
	common.CreateOneMongoDBRecordTemplate(ctx, common.CreateMongoDBTemplateOptions{
		DB:         bk.MongoDBClient,
		Context:    context.WithValue(context.Background(), "book_name", n),
		TableModel: temp,
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
			my_mongo.EQUAL,
			"rec_id",
			ctx.PostForm("stock_in_record_id"),
		},
		{
			my_mongo.LESS_THAN_EQUAL,
			"deleted_at",
			nil,
		},
		{
			my_mongo.EQUAL,
			"stock_in_record_type",
			ctx.PostForm("stock_in_record_type"),
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
func DeleteStockIn(ctx *gin.Context) {
	bk, n := common.HarvestClientFromGinContext(ctx)
	if bk == nil {
		return
	}
	var stockOutRecord mongo_model.StockInRecord
	recID, err := strconv.Atoi(ctx.PostForm("stock_in_record_id"))
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

func SelectStockInType(ctx *gin.Context) {
	common.SelectSuccessTemplate(ctx, int64(len(_const.StockInTypeList)), _const.StockInTypeList)
	return
}
