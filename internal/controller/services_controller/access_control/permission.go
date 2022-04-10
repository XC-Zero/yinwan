package access_control

import (
	"github.com/XC-Zero/yinwan/internal/controller/services_controller/common"
	"github.com/XC-Zero/yinwan/pkg/client"
	_const "github.com/XC-Zero/yinwan/pkg/const"
	"github.com/XC-Zero/yinwan/pkg/model"
	"github.com/XC-Zero/yinwan/pkg/utils/encode"
	"github.com/XC-Zero/yinwan/pkg/utils/errs"
	"github.com/gin-gonic/gin"
	"log"
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
		var rc []model.RoleCapabilities
		role := model.Role{}
		err1 := client.MysqlClient.Model(&model.Role{}).Where("rec_id = (select staff_role_id from staffs where staff_email = ? limit 1) ", staffEmail).Find(&role).Error
		err2 := client.MysqlClient.Raw("select * from "+model.RoleCapabilities{}.TableName()+" where role_id = (select staff_role_id from staffs where staff_email = ? limit 1)", staffEmail).Scan(&rc).Error

		if err1 != nil || err2 != nil {
			common.InternalDataBaseErrorTemplate(ctx, common.DATABASE_SELECT_ERROR, role)
			return
		}
		ctx.JSON(_const.OK, gin.H{
			"token":             *tokenPtr,
			"role":              role,
			"role_capabilities": rc,
		})
		return
	}

}

func HarvestRole(ctx *gin.Context) {

}

// ForgetPassword 忘记密码
func ForgetPassword(ctx *gin.Context) {
	staffEmail := ctx.PostForm("staff_email")
	staffPass := ctx.PostForm("staff_password")
	secretKey := ctx.PostForm("secret_key")
	if secretKey == "" {
		ctx.JSON(_const.FORBIDDEN_ERROR, errs.CreateWebErrorMsg("禁止访问！"))
		return
	}
	log.Println(secretKey)
	email, err := encode.DecryptByAes(secretKey)
	if err != nil {
		ctx.JSON(_const.FORBIDDEN_ERROR, errs.CreateWebErrorMsg("禁止访问！"))
		return
	}
	if string(email) == staffEmail && staffEmail != "" && staffPass != "" {
		err := client.MysqlClient.Model(&model.Staff{}).
			Where("staff_email = ?", staffEmail).
			Update("staff_password", staffPass).
			Error

		if err != nil {
			ctx.JSON(_const.INTERNAL_ERROR, gin.H(errs.CreateWebErrorMsg("更改密码失败")))
			return
		}
		ctx.JSON(_const.OK, gin.H(errs.CreateSuccessMsg("更改密码成功 !")))
		return
	} else {
		ctx.JSON(_const.REQUEST_PARM_ERROR, errs.CreateWebErrorMsg("输入有误！"))
		return
	}

}
