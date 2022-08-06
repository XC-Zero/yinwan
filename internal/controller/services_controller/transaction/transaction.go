package transaction

import (
	"context"
	"github.com/XC-Zero/yinwan/internal/controller/services_controller/common"
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

func CreateTransaction(ctx *gin.Context) {
	bk := common.HarvestClientFromGinContext(ctx)
	if bk == nil {
		return
	}
	var transaction mongo_model.Transaction
	err := ctx.ShouldBindBodyWith(&transaction, binding.JSON)
	if err != nil {
		logger.Error(errors.WithStack(err), "")
		common.RequestParamErrorTemplate(ctx, common.REQUEST_PARM_ERROR)
		return
	}
	var auto common.Auto
	ctx.ShouldBindBodyWith(&auto, binding.JSON)
	recID := int(time.Now().Unix())
	transaction.RecID = &recID
	transaction.BookName = bk.BookName
	transaction.BookNameID = bk.StorageName
	transaction.CreatedAt = strconv.Itoa(int(time.Now().Unix()))
	common.CreateOneMongoDBRecordTemplate(ctx, common.CreateMongoDBTemplateOptions{
		DB:         bk.MongoDBClient,
		Context:    auto.WithContext(ctx),
		TableModel: transaction,
	})
	return
}
func SelectTransaction(ctx *gin.Context) {
	bk := common.HarvestClientFromGinContext(ctx)
	if bk == nil {
		return
	}
	conditions := []common.MongoCondition{
		{
			my_mongo.EQUAL,
			"rec_id",
			ctx.PostForm("transaction_id"),
		},
		{
			my_mongo.LIKE,
			"transaction_name",
			ctx.PostForm("transaction_name"),
		},
		{
			my_mongo.EQUAL,
			"transaction_owner_id",
			ctx.PostForm("transaction_owner_id"),
		},
		{
			my_mongo.GREATER_THEN_EQUAL,
			"transaction_time",
			ctx.PostForm("transaction_time"),
		},
		{
			my_mongo.LESS_THAN_EQUAL,
			"transaction_time",
			ctx.PostForm("transaction_time2"),
		},
		{
			my_mongo.EQUAL,
			"deleted_at",
			bsontype.Null,
		},
	}

	common.SelectMongoDBTableContentWithCountTemplate(ctx, common.SelectMongoDBTemplateOptions{
		DB:         bk.MongoDBClient,
		TableModel: mongo_model.Transaction{},
	}, conditions...)
	return
}

// UpdateTransaction TODO 考虑是否自动级联更新?
func UpdateTransaction(ctx *gin.Context) {
	bk := common.HarvestClientFromGinContext(ctx)
	if bk == nil {
		return
	}
	var transaction mongo_model.Transaction
	err := ctx.ShouldBindBodyWith(&transaction, binding.JSON)
	if err != nil || transaction.RecID == nil {
		logger.Error(errors.WithStack(err), "")
		common.RequestParamErrorTemplate(ctx, common.REQUEST_PARM_ERROR)
		return
	}
	now := strconv.Itoa(int(time.Now().Unix()))
	transaction.UpdatedAt = &now
	var auto common.Auto
	//TODO 考虑是否自动级联更新?
	ctx.ShouldBindBodyWith(&auto, binding.JSON)
	common.UpdateOneMongoDBRecordByIDTemplate(ctx, common.MongoDBTemplateOptions{
		DB:         bk.MongoDBClient,
		Context:    auto.WithContext(ctx),
		TableModel: transaction,
		RecID:      *transaction.RecID,
	})
	return
}

func DeleteTransaction(ctx *gin.Context) {
	bk := common.HarvestClientFromGinContext(ctx)
	if bk == nil {
		return
	}

	var transaction mongo_model.Transaction
	recID, err := strconv.Atoi(ctx.PostForm("transaction_id"))
	transaction.RecID = &recID
	if err != nil || transaction.RecID == nil || recID == 0 {
		logger.Error(errors.WithStack(err), "")
		common.RequestParamErrorTemplate(ctx, common.REQUEST_PARM_ERROR)
		return
	}
	common.DeleteOneMongoDBRecordByIDTemplate(ctx, common.MongoDBTemplateOptions{
		DB:      bk.MongoDBClient,
		Context: context.WithValue(ctx, "auto", ctx.PostForm("auto")),

		TableModel: transaction,
		RecID:      *transaction.RecID,
	})
	return
}
