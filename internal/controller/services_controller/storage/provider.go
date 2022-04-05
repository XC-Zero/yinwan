package storage

import (
	"github.com/XC-Zero/yinwan/internal/controller/services_controller/common"
	"github.com/XC-Zero/yinwan/pkg/client"
	"github.com/XC-Zero/yinwan/pkg/model"
	"github.com/gin-gonic/gin"
)

func CreateProvider(ctx *gin.Context) {
	var provider model.Provider
	err := ctx.ShouldBind(&provider)
	if err != nil {
		common.RequestParamErrorTemplate(ctx, common.REQUEST_PARM_ERROR)
		return
	}
	op := common.CreateMysqlTemplateOptions{
		DB:         client.MysqlClient,
		TableModel: provider,
	}
	common.CreateOneMysqlRecordTemplate(ctx, op)
	return
}
func SelectProvider(ctx *gin.Context) {
	conditions := []common.MysqlCondition{
		{},
	}
	op := common.SelectMysqlTemplateOptions{
		DB:         client.MysqlClient,
		TableModel: model.Provider{},
	}
	common.SelectMysqlTableContentWithCountTemplate(ctx, op, conditions...)
	return
}
func UpdateProvider(ctx *gin.Context) {
	return
}
func DeleteProvider(ctx *gin.Context) {
	return
}
