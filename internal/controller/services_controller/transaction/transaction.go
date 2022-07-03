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
	"time"
)

// auto 临时结构
//	是否自动创建/删除/更新应收?
type auto struct {
	Auto bool `json:"auto" form:"auto"`
}

func CreateTransaction(ctx *gin.Context) {
	bk, n := common.HarvestClientFromGinContext(ctx)
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
	var autoCreate auto
	ctx.ShouldBindBodyWith(&autoCreate, binding.JSON)
	recID := int(time.Now().Unix())
	transaction.RecID = &recID
	transaction.BookName = n
	transaction.BookNameID = bk.StorageName
	common.CreateOneMongoDBRecordTemplate(ctx, common.CreateMongoDBTemplateOptions{
		DB: bk.MongoDBClient,
		Context: context.WithValue(context.WithValue(context.Background(), "book_name", n),
			"auto_create", autoCreate.Auto),
		TableModel: transaction,
	})
	return
}
func SelectTransaction(ctx *gin.Context) {
	bk, _ := common.HarvestClientFromGinContext(ctx)
	if bk == nil {
		return
	}
	conditions := []common.MongoCondition{
		{
			my_mongo.LESS_THAN_EQUAL,
			"deleted_at",
			nil,
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
	bk, n := common.HarvestClientFromGinContext(ctx)
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
	var autoUpdate auto
	//ctx.ShouldBindBodyWith(&autoUpdate, binding.JSON) TODO 考虑是否自动级联更新?
	common.UpdateOneMongoDBRecordByIDTemplate(ctx, common.MongoDBTemplateOptions{
		DB: bk.MongoDBClient,
		Context: context.WithValue(context.WithValue(context.Background(), "book_name", n),
			"auto_update", autoUpdate.Auto),
		TableModel: transaction,
		RecID:      *transaction.RecID,
	})
	return
}

func DeleteTransaction(ctx *gin.Context) {
	bk, n := common.HarvestClientFromGinContext(ctx)
	if bk == nil {
		return
	}

	var autoDelete auto
	ctx.ShouldBindBodyWith(&autoDelete, binding.JSON)
	var transaction mongo_model.Transaction
	err := ctx.ShouldBindBodyWith(&transaction, binding.JSON)
	if err != nil || transaction.RecID == nil {
		logger.Error(errors.WithStack(err), "")
		common.RequestParamErrorTemplate(ctx, common.REQUEST_PARM_ERROR)
		return
	}
	common.DeleteOneMongoDBRecordByIDTemplate(ctx, common.MongoDBTemplateOptions{
		DB: bk.MongoDBClient,
		Context: context.WithValue(context.WithValue(context.Background(), "book_name", n),
			"auto_create", autoDelete.Auto),
		TableModel: transaction,
		RecID:      *transaction.RecID,
	})
	return
}
