package finance

import (
	"context"
	"github.com/XC-Zero/yinwan/internal/controller/services_controller/common"
	"github.com/XC-Zero/yinwan/pkg/model/mongo_model"
	"github.com/XC-Zero/yinwan/pkg/utils/mongo"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"strconv"
	"time"
)

// todo !!! 凭证模板

func CreateCredentialTemplate(ctx *gin.Context) {
	bk, n := common.HarvestClientFromGinContext(ctx)
	if bk == nil {
		return
	}

	var credentialTemplate mongo_model.CredentialTemplate
	err := ctx.ShouldBindBodyWith(&credentialTemplate, binding.JSON)
	if err != nil || credentialTemplate.RecID == nil {
		common.RequestParamErrorTemplate(ctx, common.REQUEST_PARM_ERROR)
		return
	}
	recID := int(time.Now().Unix())
	credentialTemplate.RecID = &recID
	credentialTemplate.CreatedAt = time.Now().String()
	common.CreateOneMongoDBRecordTemplate(ctx, common.CreateMongoDBTemplateOptions{
		DB:         bk.MongoDBClient,
		Context:    context.WithValue(context.Background(), "book_name", n),
		TableModel: &credentialTemplate,
		NotSyncES:  true,
	})
}
func SelectCredentialTemplate(ctx *gin.Context) {
	bk, _ := common.HarvestClientFromGinContext(ctx)
	if bk == nil {
		return
	}

	conditions := []common.MongoCondition{
		{
			Symbol:      mongo.EQUAL,
			ColumnName:  "rec_id",
			ColumnValue: ctx.PostForm("credential_template_id"),
		},
		{
			Symbol:      mongo.NOT_EQUAL,
			ColumnName:  "deleted_at",
			ColumnValue: " ",
		},
	}

	common.SelectMongoDBTableContentWithCountTemplate(ctx, common.SelectMongoDBTemplateOptions{
		DB:         bk.MongoDBClient,
		TableModel: mongo_model.CredentialTemplate{},
	}, conditions...)

}
func UpdateCredentialTemplate(ctx *gin.Context) {
	bk, n := common.HarvestClientFromGinContext(ctx)
	if bk == nil {
		return
	}

	var credentialTemplate mongo_model.CredentialTemplate
	err := ctx.ShouldBindBodyWith(&credentialTemplate, binding.JSON)
	if err != nil || credentialTemplate.RecID == nil {
		common.RequestParamErrorTemplate(ctx, common.REQUEST_PARM_ERROR)
		return
	}
	now := time.Now().String()
	credentialTemplate.UpdatedAt = &now

	common.UpdateOneMongoDBRecordByIDTemplate(ctx, common.MongoDBTemplateOptions{
		DB:         bk.MongoDBClient,
		Context:    context.WithValue(context.Background(), "book_name", n),
		TableModel: &credentialTemplate,
		RecID:      *credentialTemplate.RecID,
		NotSyncES:  true,
	})
	return

}
func DeleteCredentialTemplate(ctx *gin.Context) {
	bk, n := common.HarvestClientFromGinContext(ctx)
	if bk == nil {
		return
	}
	var credentialTemplate mongo_model.CredentialTemplate
	recID, err := strconv.Atoi(ctx.PostForm("credential_template_id"))
	if err != nil || recID == 0 {
		common.RequestParamErrorTemplate(ctx, common.REQUEST_PARM_ERROR)
		return
	}
	credentialTemplate.RecID = &recID
	common.DeleteOneMongoDBRecordByIDTemplate(ctx, common.MongoDBTemplateOptions{
		DB:         bk.MongoDBClient,
		Context:    context.WithValue(context.Background(), "book_name", n),
		TableModel: &credentialTemplate,
		RecID:      recID,
		NotSyncES:  true,
	})
}
