package finance

import (
	"context"
	"github.com/XC-Zero/yinwan/internal/controller/services_controller/common"
	"github.com/XC-Zero/yinwan/pkg/model/mongo_model"
	my_mongo "github.com/XC-Zero/yinwan/pkg/utils/mongo"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"go.mongodb.org/mongo-driver/bson/bsontype"
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
	credential.CreatedAt = time.Now().String()
	credential.BookName = n
	credential.BookNameID = bk.StorageName
	common.CreateOneMongoDBRecordTemplate(ctx, common.CreateMongoDBTemplateOptions{
		DB:         bk.MongoDBClient,
		TableModel: &credential,
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
			my_mongo.EQUAL,
			"rec_id",
			ctx.PostForm("credential_id"),
		},
		{
			my_mongo.LIKE,
			"credential_name",
			ctx.PostForm("credential_name"),
		},
		{
			my_mongo.EQUAL,
			"deleted_at",
			bsontype.Null,
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
		TableModel: &credential,
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
		TableModel: &credential,
		Context:    context.WithValue(context.Background(), "book_name", n),
	})
	return
}
