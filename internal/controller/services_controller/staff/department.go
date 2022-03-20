package staff

import (
	"fmt"
	"github.com/XC-Zero/yinwan/pkg/client"
	_const "github.com/XC-Zero/yinwan/pkg/const"
	"github.com/XC-Zero/yinwan/pkg/model"
	"github.com/XC-Zero/yinwan/pkg/utils/errs"
	"github.com/gin-gonic/gin"
	"strings"
)

// SelectDepartment 查询部门
func SelectDepartment(ctx *gin.Context) {
	departmentName := ctx.PostForm("department_name")
	departmentMangerName := ctx.PostForm("department_manger_name")
	departmentManagerID := ctx.PostForm("department_manager_id")

	var departmentList []model.Department
	var count int
	sql := `select * from departments where 1 = 1 `
	if departmentName != "" {
		sql += fmt.Sprintf(" and department_name like '%%%s%%' ", departmentName)
	}
	if departmentMangerName != "" {
		sql += fmt.Sprintf(" and department_manager_name like '%%%s%%' ", departmentMangerName)
	}
	if departmentManagerID != "" {
		sql += fmt.Sprintf(" and department_manager_id = '%s' ", departmentManagerID)
	}
	countSql := strings.Replace(sql, "*", "count(*)", 1)

	sql += " order by rec_id " + client.PaginateSql(ctx)
	err := client.MysqlClient.Raw(sql).Scan(&departmentList).Error
	if err != nil {
		ctx.JSON(_const.INTERNAL_ERROR, gin.H(errs.CreateWebErrorMsg("查询部门内容失败！")))
		return
	}
	err = client.MysqlClient.Raw(countSql).Scan(&count).Error
	if err != nil {
		ctx.JSON(_const.INTERNAL_ERROR, gin.H(errs.CreateWebErrorMsg("查询部门总数失败！")))
		return
	}
	ctx.JSON(_const.OK, gin.H{
		"count":           count,
		"department_list": departmentList,
	})
	return
}

// CreateDepartment todo !!!
func CreateDepartment(ctx *gin.Context) {

}

// UpdateDepartment todo !!!
func UpdateDepartment(ctx *gin.Context) {
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
