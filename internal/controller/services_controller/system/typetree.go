package system

import (
	"github.com/XC-Zero/yinwan/internal/controller/services_controller/common"
	"github.com/XC-Zero/yinwan/pkg/client"
	_const "github.com/XC-Zero/yinwan/pkg/const"
	"github.com/XC-Zero/yinwan/pkg/model"
	"github.com/XC-Zero/yinwan/pkg/utils/errs"
	"github.com/gin-gonic/gin"
)

func CreateTypeTree(ctx *gin.Context) {
	var typeTree model.TypeTree
	err := ctx.ShouldBind(&typeTree)
	if err != nil {
		common.RequestParamErrorTemplate(ctx, common.REQUEST_PARM_ERROR)
		return
	}
	err = client.MysqlClient.Create(typeTree).Error
	if err != nil {
		common.InternalDataBaseErrorTemplate(ctx, common.DATABASE_INSERT_ERROR, typeTree)
		return
	}
	ctx.JSON(_const.OK, errs.CreateSuccessMsg("创建类型成功！"))
	return
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

func UpdateTypeTree(ctx *gin.Context) {

}

func DeleteTypeTree(ctx *gin.Context) {

}
