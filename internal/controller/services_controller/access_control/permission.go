package access_control

import (
	"fmt"
	"github.com/XC-Zero/yinwan/pkg/client"
	_const "github.com/XC-Zero/yinwan/pkg/const"
	"github.com/XC-Zero/yinwan/pkg/model"
	"github.com/XC-Zero/yinwan/pkg/utils/errs"
	"github.com/XC-Zero/yinwan/pkg/utils/logger"
	"github.com/fwhezfwhez/errorx"
	"github.com/gin-gonic/gin"
)

// Login
// @Param  staff_email string,staff_password
func Login(ctx *gin.Context) {
	staffEmail, staffPassword := ctx.PostForm("staff_email"), ctx.PostForm("staff_password")
	staff := model.Staff{
		StaffEmail:    staffEmail,
		StaffPassword: staffPassword,
	}

	tokenPtr, errMessage := staff.Login()
	if tokenPtr == nil {
		ctx.JSON(_const.REQUEST_PARM_ERROR, gin.H(errs.CreateWebErrorMsg(errMessage)))
	} else {

		//client.MysqlClient.Raw("").Scan()
		ctx.JSON(_const.OK, gin.H{
			"token": *tokenPtr,
			"role":  "",
		})

	}

}

func HarvestRole(ctx *gin.Context) {

}

// ForgetPassword 忘记密码
func ForgetPassword(ctx *gin.Context) {
	staffEmail, staffPass := ctx.PostForm("staff_email"), ctx.PostForm("staff_password")
	err := client.MysqlClient.Model(&model.Staff{}).
		Update("staff_password", staffPass).
		Where("staff_email = ?", staffEmail).
		Error

	if err != nil {
		logger.Error(errorx.MustWrap(err), fmt.Sprintf("邮箱: %s 更改密码失败！", staffEmail))
		ctx.JSON(_const.INTERNAL_ERROR, gin.H(errs.CreateWebErrorMsg("更改密码失败")))
		return
	}
	logger.Info(fmt.Sprintf("邮箱: %s 更改密码成功！", staffEmail))
	ctx.JSON(_const.OK, gin.H{
		"message": "更改密码成功",
	})

	return
}
