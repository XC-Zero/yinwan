package finance

import (
	"context"
	"github.com/XC-Zero/yinwan/internal/controller/services_controller/common"
	"github.com/XC-Zero/yinwan/pkg/client"
	"github.com/XC-Zero/yinwan/pkg/model/mongo_model"
	"github.com/XC-Zero/yinwan/pkg/utils/mysql"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"strconv"
	"time"
)

// todo !!! 凭证模板

func CreateCredentialTemplate(ctx *gin.Context) {
	var credentialTemplate mongo_model.CredentialTemplate
	err := ctx.ShouldBindBodyWith(&credentialTemplate, binding.JSON)
	if err != nil || credentialTemplate.RecID == nil {
		common.RequestParamErrorTemplate(ctx, common.REQUEST_PARM_ERROR)
		return
	}
	recID := int(time.Now().Unix())
	credentialTemplate.RecID = &recID
	credentialTemplate.CreatedAt = strconv.FormatInt(time.Now().Unix(), 10)
	common.CreateOneMongoDBRecordTemplate(ctx, common.CreateMongoDBTemplateOptions{
		DB:         client.MongoDBClient,
		Context:    context.TODO(),
		TableModel: mongo_model.CredentialTemplate{},
		NotSyncES:  true,
	})
}
func SelectCredentialTemplate(ctx *gin.Context) {

	conditions := []common.MysqlCondition{
		{
			Symbol:      mysql.EQUAL,
			ColumnName:  "rec_id",
			ColumnValue: ctx.PostForm("credential_template_id"),
		},
		{
			Symbol:      mysql.NULL,
			ColumnName:  "deleted_at",
			ColumnValue: " ",
		},
	}

	common.SelectMysqlTableContentWithCountTemplate(ctx, common.SelectMysqlTemplateOptions{
		DB:         client.MysqlClient,
		TableModel: mongo_model.CredentialTemplate{},
	}, conditions...)

}
func UpdateCredentialTemplate(ctx *gin.Context) {
	var credentialTemplate mongo_model.CredentialTemplate
	err := ctx.ShouldBindBodyWith(&credentialTemplate, binding.JSON)
	if err != nil || credentialTemplate.RecID == nil {
		common.RequestParamErrorTemplate(ctx, common.REQUEST_PARM_ERROR)
		return
	}
	common.UpdateOneMongoDBRecordByIDTemplate(ctx, common.MongoDBTemplateOptions{
		DB:         client.MongoDBClient,
		Context:    context.TODO(),
		TableModel: credentialTemplate,
		RecID:      *credentialTemplate.RecID,
		NotSyncES:  true,
	})
	return

}
func DeleteCredentialTemplate(ctx *gin.Context) {
	var credentialTemplate mongo_model.CredentialTemplate
	recID, err := strconv.Atoi(ctx.PostForm("credential_template_id"))
	if err != nil || recID == 0 {
		common.RequestParamErrorTemplate(ctx, common.REQUEST_PARM_ERROR)
		return
	}
	credentialTemplate.RecID = &recID
	common.DeleteOneMongoDBRecordByIDTemplate(ctx, common.MongoDBTemplateOptions{
		DB:         client.MongoDBClient,
		Context:    context.TODO(),
		TableModel: credentialTemplate,
		RecID:      recID,
		NotSyncES:  true,
	})
}
