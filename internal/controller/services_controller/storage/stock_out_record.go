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
	"go.mongodb.org/mongo-driver/bson/bsontype"
	"strconv"
	"time"
)

// CreateStockOut
//	出库
func CreateStockOut(ctx *gin.Context) {
	bk, n := common.HarvestClientFromGinContext(ctx)
	if bk == nil {
		return
	}
	temp := mongo_model.StockOutRecord{}

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
	common.CreateOneMongoDBRecordTemplate(ctx, common.CreateMongoDBTemplateOptions{
		DB:         bk.MongoDBClient,
		Context:    context.WithValue(context.Background(), "book_name", n),
		TableModel: temp,
	})
	return

}

// SelectStockOut 查询出库
func SelectStockOut(ctx *gin.Context) {
	bk, _ := common.HarvestClientFromGinContext(ctx)
	if bk == nil {
		return
	}

	conditions := []common.MongoCondition{
		{
			my_mongo.EQUAL,
			"rec_id",
			ctx.PostForm("stock_out_record_id"),
		},
		{
			my_mongo.NOT_EQUAL,
			"deleted_at",
			bsontype.Null,
		},
		{
			my_mongo.EQUAL,
			"stock_out_record_type",
			ctx.PostForm("stock_out_record_type"),
		},
		{
			my_mongo.EQUAL,
			"stock_out_warehouse_id",
			ctx.PostForm("stock_out_warehouse_id"),
		},
		{
			my_mongo.EQUAL,
			"stock_out_record_owner_id",
			ctx.PostForm("stock_out_record_owner_id"),
		},
	}
	options := common.SelectMongoDBTemplateOptions{
		DB:         bk.MongoDBClient,
		TableModel: mongo_model.StockOutRecord{},
	}
	common.SelectMongoDBTableContentWithCountTemplate(ctx, options, conditions...)
	return

}

// UpdateStockOut 更新出库
func UpdateStockOut(ctx *gin.Context) {
	bk, n := common.HarvestClientFromGinContext(ctx)
	if bk == nil {
		return
	}

	temp := mongo_model.StockOutRecord{}

	err := ctx.ShouldBindBodyWith(&temp, binding.JSON)
	if err != nil || temp.RecID == nil {
		logger.Error(errors.WithStack(err), "")
		common.RequestParamErrorTemplate(ctx, common.REQUEST_PARM_ERROR)
		return
	}
	common.UpdateOneMongoDBRecordByIDTemplate(ctx, common.MongoDBTemplateOptions{
		DB:         bk.MongoDBClient,
		Context:    context.WithValue(context.Background(), "book_name", n),
		RecID:      *temp.RecID,
		TableModel: temp,
	})
	return
}

// DeleteStockOut 删除出库
func DeleteStockOut(ctx *gin.Context) {
	bk, n := common.HarvestClientFromGinContext(ctx)
	if bk == nil {
		return
	}
	var stockOutRecord mongo_model.StockOutRecord
	recID, err := strconv.Atoi(ctx.PostForm("stock_out_record_id"))
	if err != nil {
		logger.Error(errors.WithStack(err), "")
		common.RequestParamErrorTemplate(ctx, common.REQUEST_PARM_ERROR)
		return
	}
	common.DeleteOneMongoDBRecordByIDTemplate(ctx, common.MongoDBTemplateOptions{
		DB:         bk.MongoDBClient,
		Context:    context.WithValue(context.Background(), "book_name", n),
		RecID:      recID,
		TableModel: stockOutRecord,
	})
	return
}

// SelectStockOutType 查询出库类型
func SelectStockOutType(ctx *gin.Context) {
	common.SelectSuccessTemplate(ctx, int64(len(_const.StockOutTypeList)), _const.StockInTypeList)
	return
}

// SelectInvoiceType 查询单据类型
func SelectInvoiceType(ctx *gin.Context) {
	common.SelectSuccessTemplate(ctx, int64(len(_const.InvoiceTypeList)), _const.InvoiceTypeList)
	return
}
