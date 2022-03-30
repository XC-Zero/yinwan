package staff

import (
	"github.com/XC-Zero/yinwan/internal/controller/services_controller/common"
	"github.com/XC-Zero/yinwan/pkg/client"
	_const "github.com/XC-Zero/yinwan/pkg/const"
	"github.com/XC-Zero/yinwan/pkg/model"
	"github.com/XC-Zero/yinwan/pkg/utils/errs"
	"github.com/XC-Zero/yinwan/pkg/utils/mysql"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

// SelectDepartment 查询部门
func SelectDepartment(ctx *gin.Context) {

	conditions := []common.Condition{
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
	}
	common.SelectTableContentWithCountMysqlTemplate(ctx, client.MysqlClient, model.Department{}, "", nil, conditions...)

	return
}

// CreateDepartment todo !!!
func CreateDepartment(ctx *gin.Context) {
	var department model.Department
	err := ctx.ShouldBindWith(&department, binding.Form)
	if err != nil {
		ctx.JSON(_const.REQUEST_PARM_ERROR, errs.CreateWebErrorMsg("部门参数有误！"))
		return
	}
	err = client.MysqlClient.Model(&model.Department{}).Create(&department).Error
	if err != nil {
		ctx.JSON(_const.INTERNAL_ERROR, errs.CreateWebErrorMsg("创建部门失败！"))
		return
	}

	ctx.JSON(_const.OK, errs.CreateSuccessMsg("创建部门成功！"))
	return

}

// UpdateDepartment todo !!!
func UpdateDepartment(ctx *gin.Context) {
	var department model.Department
	err := ctx.ShouldBind(&department)
	if err != nil {
		ctx.JSON(_const.REQUEST_PARM_ERROR, "输入有误！")
		return
	}
	client.MysqlClient.Model(&model.Department{}).Updates(department)
}

// DeleteDepartment todo !!!
func DeleteDepartment(ctx *gin.Context) {
	staffEmail := ctx.PostForm("staff_email")
	var staffList []model.Staff

	err := client.MysqlClient.Model(&model.Staff{}).Raw("delete from staff where staff_email = ?", staffEmail).Scan(staffList).Error
	if err != nil {
		ctx.JSON(_const.INTERNAL_ERROR, gin.H(errs.CreateWebErrorMsg("查询失败！")))
		return
	}
	ctx.JSON(_const.OK, staffList)
	return
}
