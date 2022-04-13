package system

import (
	"github.com/XC-Zero/yinwan/internal/controller/services_controller/common"
	"github.com/XC-Zero/yinwan/pkg/model"
	"github.com/gin-gonic/gin"
)

func CreateTypeTree(ctx *gin.Context) {
	//ctx.ShouldBindWith()
}
func SelectTypeTree(ctx *gin.Context) {
	op := common.SelectMysqlTemplateOptions{
		DB:            nil,
		TableModel:    model.TypeTree{},
		OrderByColumn: "",
		ResHookFunc:   nil,
	}
	cds := []common.MysqlCondition{{}}
	common.SelectMysqlTableContentWithCountTemplate(ctx, op, cds...)
}
