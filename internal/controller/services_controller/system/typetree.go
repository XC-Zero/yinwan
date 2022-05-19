package system

import (
	"github.com/XC-Zero/yinwan/internal/controller/services_controller/common"
	"github.com/XC-Zero/yinwan/pkg/client"
	_const "github.com/XC-Zero/yinwan/pkg/const"
	"github.com/XC-Zero/yinwan/pkg/model/mysql_model"
	"github.com/XC-Zero/yinwan/pkg/utils/errs"
	"github.com/XC-Zero/yinwan/pkg/utils/mysql"
	"github.com/gin-gonic/gin"
)

// CreateTypeTree 创建类型
func CreateTypeTree(ctx *gin.Context) {
	var typeTree mysql_model.TypeTree
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

// SelectTypeTree 类型
func SelectTypeTree(ctx *gin.Context) {
	op := common.SelectMysqlTemplateOptions{
		DB:            client.MysqlClient,
		TableModel:    mysql_model.TypeTree{},
		OrderByColumn: "",
		ResHookFunc:   nil,
	}
	parentTypeId := ctx.PostForm("parent_type_id")
	// 默认为一级类型
	var parentCondition common.MysqlCondition

	if parentTypeId == "" {
		parentCondition = common.MysqlCondition{
			Symbol:      mysql.NULL,
			ColumnName:  "parent_type_id",
			ColumnValue: " ",
		}
	}
	parentCondition = common.MysqlCondition{
		Symbol:      mysql.EQUAL,
		ColumnName:  "parent_type_id",
		ColumnValue: parentTypeId,
	}
	cds := []common.MysqlCondition{
		parentCondition,
		{
			Symbol:      mysql.EQUAL,
			ColumnName:  "rec_id",
			ColumnValue: ctx.PostForm("type_id"),
		},
		{
			Symbol:      mysql.LIKE,
			ColumnName:  "type_name",
			ColumnValue: ctx.PostForm("type_name"),
		},
		{
			Symbol:      mysql.NULL,
			ColumnName:  "deleted_at",
			ColumnValue: " ",
		},
	}
	common.SelectMysqlTableContentWithCountTemplate(ctx, op, cds...)
	return
}

// UpdateTypeTree 更新类型
func UpdateTypeTree(ctx *gin.Context) {
	typeTree := mysql_model.TypeTree{}
	err := ctx.ShouldBind(&typeTree)
	if err != nil || typeTree.RecID == nil {
		common.RequestParamErrorTemplate(ctx, common.REQUEST_PARM_ERROR)
		return
	}
	err = client.MysqlClient.Updates(&typeTree).Where("rec_id = ?", typeTree.RecID).Error
	if err != nil {
		common.InternalDataBaseErrorTemplate(ctx, common.DATABASE_UPDATE_ERROR, typeTree)
		return
	}
	ctx.JSON(_const.OK, errs.CreateSuccessMsg("更新类型成功！"))
	return
}

// DeleteTypeTree 删除类型
func DeleteTypeTree(ctx *gin.Context) {
	var typeTree mysql_model.TypeTree
	id := ctx.PostForm("type_id")
	if id == "" {
		common.RequestParamErrorTemplate(ctx, common.REQUEST_PARM_ERROR)
		return
	}
	err := client.MysqlClient.Delete(typeTree, id).Error
	if err != nil {
		common.InternalDataBaseErrorTemplate(ctx, common.DATABASE_DELETE_ERROR, typeTree)
		return
	}
	ctx.JSON(_const.OK, errs.CreateSuccessMsg("删除类型成功！"))
	return
}
