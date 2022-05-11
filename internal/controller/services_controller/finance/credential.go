package finance

import (
	"github.com/XC-Zero/yinwan/internal/controller/services_controller/common"
	"github.com/XC-Zero/yinwan/pkg/model/mongo_model"
	"github.com/gin-gonic/gin"
)

// CreateCredential 创建凭证
func CreateCredential(ctx *gin.Context) {
	bk, _ := common.HarvestClientFromGinContext(ctx)
	if bk == nil {
		return
	}
	var credential mongo_model.FinanceCredential

	err := ctx.ShouldBind(&credential)
	if err != nil {
		common.RequestParamErrorTemplate(ctx, common.REQUEST_PARM_ERROR)
		return
	}
	common.CreateOneMongoDBRecordTemplate(ctx, common.CreateMongoDBTemplateOptions{
		DB:         bk.MongoDBClient,
		TableModel: credential,
		PreFunc:    nil,
	})
	return
}

// SelectCredential 查找凭证
func SelectCredential(ctx *gin.Context) {
	bk, _ := common.HarvestClientFromGinContext(ctx)
	if bk == nil {
		return
	}
}

// UpdateCredential 更新凭证
func UpdateCredential(ctx *gin.Context) {

	bk, _ := common.HarvestClientFromGinContext(ctx)
	if bk == nil {
		return
	}
	var credential mongo_model.FinanceCredential
	err := ctx.ShouldBind(&credential)
	if err != nil {
		common.RequestParamErrorTemplate(ctx, common.REQUEST_PARM_ERROR)
		return
	}
	common.UpdateOneMongoDBRecordByIDTemplate(ctx, common.UpdateMongoDBTemplateOptions{
		DB:         bk.MongoDBClient,
		RecID:      0,
		TableModel: credential,
		PreFunc:    nil,
	})
	return
}

// DeleteCredential 删除凭证
func DeleteCredential(ctx *gin.Context) {
	bk, _ := common.HarvestClientFromGinContext(ctx)
	if bk == nil {
		return
	}
}
