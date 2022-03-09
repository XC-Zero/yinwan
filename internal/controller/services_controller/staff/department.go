package staff

import (
	"github.com/XC-Zero/yinwan/pkg/client"
	_const "github.com/XC-Zero/yinwan/pkg/const"
	"github.com/XC-Zero/yinwan/pkg/model"
	"github.com/XC-Zero/yinwan/pkg/utils/errs"
	"github.com/gin-gonic/gin"
)

func SelectDepartment(ctx *gin.Context) {
	staffEmail := ctx.PostForm("staff_email")
	var staffList []model.Staff
	err := client.MysqlClient.Model(&model.Staff{}).Raw("select * from staff where staff_email = ?", staffEmail).Scan(staffList).Error
	if err != nil {
		ctx.JSON(_const.INTERNAL_ERROR, gin.H(errs.CreateWebErrorMsg("查询失败！")))
		return
	}
	ctx.JSON(_const.OK, staffList)
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
