package staff

import (
	"github.com/XC-Zero/yinwan/internal/controller/services_controller/common"
	"github.com/XC-Zero/yinwan/pkg/client"
	_const "github.com/XC-Zero/yinwan/pkg/const"
	"github.com/XC-Zero/yinwan/pkg/model/mysql_model"
	"github.com/XC-Zero/yinwan/pkg/utils/errs"
	"github.com/XC-Zero/yinwan/pkg/utils/logger"
	"github.com/XC-Zero/yinwan/pkg/utils/mysql"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/pkg/errors"
)

// SelectDepartment 查询部门
func SelectDepartment(ctx *gin.Context) {

	conditions := []common.MysqlCondition{
		{
			Symbol:      mysql.LIKE,
			ColumnName:  "department_name",
			ColumnValue: ctx.PostForm("department_name"),
		},
		{
			Symbol:      mysql.LIKE,
			ColumnName:  "department_manger_name",
			ColumnValue: ctx.PostForm("department_manger_name"),
		},
		{
			Symbol:      mysql.EQUAL,
			ColumnName:  "department_manager_id",
			ColumnValue: ctx.PostForm("department_manager_id"),
		},
		{
			Symbol:      mysql.NULL,
			ColumnName:  "deleted_at",
			ColumnValue: " ",
		},
	}
	op := common.SelectMysqlTemplateOptions{
		DB:            client.MysqlClient,
		TableModel:    mysql_model.Department{},
		OrderByColumn: "",
		ResHookFunc:   nil,
	}
	common.SelectMysqlTableContentWithCountTemplate(ctx, op, conditions...)

	return
}

// CreateDepartment 创建部门
func CreateDepartment(ctx *gin.Context) {
	var department mysql_model.Department
	err := ctx.ShouldBindBodyWith(&department, binding.JSON)
	if err != nil {
		logger.Error(errors.WithStack(err), "")
		common.RequestParamErrorTemplate(ctx, common.REQUEST_PARM_ERROR)
		return
	}
	err = client.MysqlClient.Model(&mysql_model.Department{}).Create(&department).Error
	if err != nil {
		logger.Error(errors.WithStack(err), "")
		common.InternalDataBaseErrorTemplate(ctx, common.DATABASE_INSERT_ERROR, department)
		return
	}

	ctx.JSON(_const.OK, errs.CreateSuccessMsg("创建部门成功！"))
	return

}

// UpdateDepartment 更新部门
func UpdateDepartment(ctx *gin.Context) {
	var department mysql_model.Department
	err := ctx.ShouldBindBodyWith(&department, binding.JSON)
	if err != nil || department.RecID == nil {
		logger.Error(errors.WithStack(err), "")
		common.RequestParamErrorTemplate(ctx, common.REQUEST_PARM_ERROR)
		return
	}
	err = client.MysqlClient.Model(&mysql_model.Department{}).Where(" rec_id = ? ", department.RecID).Omit("rec_id").Updates(department).Error
	if err != nil {
		logger.Error(errors.WithStack(err), "")
		common.InternalDataBaseErrorTemplate(ctx, common.DATABASE_UPDATE_ERROR, department)
		return
	}
	ctx.JSON(_const.OK, errs.CreateSuccessMsg("更新部门成功！"))
	return
}

// DeleteDepartment 删除部门
func DeleteDepartment(ctx *gin.Context) {
	departmentID := ctx.PostForm("department_id")
	if departmentID == "" {
		common.RequestParamErrorTemplate(ctx, common.REQUEST_PARM_ERROR)
		return
	}
	err := client.MysqlClient.Delete(&mysql_model.Department{}, departmentID).Error
	if err != nil {
		logger.Error(errors.WithStack(err), "")
		common.InternalDataBaseErrorTemplate(ctx, common.DATABASE_DELETE_ERROR, mysql_model.Department{})
		return
	}
	ctx.JSON(_const.OK, errs.CreateSuccessMsg("删除部门成功！"))
	return
}
