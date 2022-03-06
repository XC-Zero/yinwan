package access_control

import (
	_const "github.com/XC-Zero/yinwan/pkg/const"
	"github.com/XC-Zero/yinwan/pkg/model"
	"github.com/XC-Zero/yinwan/pkg/utils/errs"
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
		ctx.JSON(_const.OK, gin.H{
			"token": *tokenPtr,
		})

	}

}
