package finance

import (
	"context"
	"github.com/XC-Zero/yinwan/internal/controller/services_controller/common"
	"github.com/XC-Zero/yinwan/pkg/model/mongo_model"
	my_mongo "github.com/XC-Zero/yinwan/pkg/utils/mongo"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"strconv"
	"time"
)

// CreateCredential 创建凭证
func CreateCredential(ctx *gin.Context) {
	bk, n := common.HarvestClientFromGinContext(ctx)
	if bk == nil {
		return
	}
	var credential mongo_model.Credential

	err := ctx.ShouldBindBodyWith(&credential, binding.JSON)
	if err != nil {
		common.RequestParamErrorTemplate(ctx, common.REQUEST_PARM_ERROR)
		return
	}
	recID := int(time.Now().Unix())
	credential.RecID = &recID
	credential.CreatedAt = strconv.FormatInt(time.Now().Unix(), 10)
	credential.BookName = n
	credential.BookNameID = bk.StorageName
	common.CreateOneMongoDBRecordTemplate(ctx, common.CreateMongoDBTemplateOptions{
		DB:         bk.MongoDBClient,
		TableModel: credential,
		Context:    context.WithValue(context.Background(), "book_name", n),
	})
	return
}

// SelectCredential 查找凭证
func SelectCredential(ctx *gin.Context) {
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

	op := common.SelectMongoDBTemplateOptions{
		DB:         bk.MongoDBClient,
		TableModel: mongo_model.Credential{},
	}
	common.SelectMongoDBTableContentWithCountTemplate(ctx, op, conditions...)
	return
}

// UpdateCredential 更新凭证
func UpdateCredential(ctx *gin.Context) {

	bk, n := common.HarvestClientFromGinContext(ctx)
	if bk == nil {
		return
	}
	var credential mongo_model.Credential
	err := ctx.ShouldBindBodyWith(&credential, binding.JSON)
	if err != nil || credential.RecID == nil {
		common.RequestParamErrorTemplate(ctx, common.REQUEST_PARM_ERROR)
		return
	}
	common.UpdateOneMongoDBRecordByIDTemplate(ctx, common.MongoDBTemplateOptions{
		DB:         bk.MongoDBClient,
		RecID:      *credential.RecID,
		TableModel: credential,
		Context:    context.WithValue(context.Background(), "book_name", n),
	})
	return
}

// DeleteCredential 删除凭证
func DeleteCredential(ctx *gin.Context) {
	bk, n := common.HarvestClientFromGinContext(ctx)
	if bk == nil {
		return
	}
	var credential mongo_model.Credential
	err := ctx.ShouldBindBodyWith(&credential, binding.JSON)
	if err != nil || credential.RecID == nil {
		common.RequestParamErrorTemplate(ctx, common.REQUEST_PARM_ERROR)
		return
	}

	common.DeleteOneMongoDBRecordByIDTemplate(ctx, common.MongoDBTemplateOptions{
		DB:         bk.MongoDBClient,
		RecID:      *credential.RecID,
		TableModel: credential,
		Context:    context.WithValue(context.Background(), "book_name", n),
	})
	return
}
