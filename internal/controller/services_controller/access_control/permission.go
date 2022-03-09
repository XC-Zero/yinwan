package access_control

import (
	"github.com/XC-Zero/yinwan/pkg/client"
	_const "github.com/XC-Zero/yinwan/pkg/const"
	"github.com/XC-Zero/yinwan/pkg/model"
	"github.com/XC-Zero/yinwan/pkg/utils/email"
	"github.com/XC-Zero/yinwan/pkg/utils/errs"
	"github.com/XC-Zero/yinwan/pkg/utils/logger"
	"github.com/fwhezfwhez/errorx"
	"github.com/gin-gonic/gin"
	"math/rand"
	"strconv"
	"time"
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

func SendForgetPasswordEmail(ctx *gin.Context) {
	staffEmail := ctx.PostForm("staff_email")
	if len(staffEmail) == 0 {
		ctx.JSON(_const.REQUEST_PARM_ERROR, gin.H(errs.CreateWebErrorMsg("未输入邮箱哦")))
	}
	rand.Seed(time.Now().Unix())
	n := strconv.Itoa(rand.Intn(999999))
	err := client.RedisClient.Set(staffEmail, n, time.Minute*10).Err()
	if err != nil {
		mes := "Redis存储邮件验证码失败!"
		logger.Error(errorx.MustWrap(err), mes)
		ctx.JSON(_const.INTERNAL_ERROR, gin.H(errs.CreateWebErrorMsg(mes)))
		return
	}
	err = email.SendEmail("好上好系统验证邮件,您正在执行找回密码操作", "<div style=\"width: 100vw;height: 100vh;background-color: #ffffff;display: flex;flex-direction: column; justify-content: center;align-items: center\">"+
		"<div style=\"width: 80%;text-align:left;\"><h1 style=\"color: #4F709B\">HSH</h1>"+
		"<h2>欢迎您，您的验证码为：</h2></div>"+
		"<div style=\" margin-top:50px;margin-bottom:50px;background-color: #ffffff;width: 300px ;height: 150px;box-shadow: 0px 3px 6px #999999;border-radius: 5px ;line-height:150px;text-align:center;letter-spacing:5px;font-size: 40px;\">"+
		n+
		"</div>"+
		"<h3 style=\"width: 80%;text-align:right;\">验证码十分钟有效~</h3>"+
		"</div>", staffEmail)
	if err != nil {
		mes := "发送邮件验证码失败!"
		logger.Error(errorx.MustWrap(err), mes)
		ctx.JSON(_const.INTERNAL_ERROR, gin.H(errs.CreateWebErrorMsg(mes)))
		return
	}
	ctx.JSON(200, gin.H(errs.CreateWebErrorMsg("邮箱验证码发送成功！")))
	return
}

// ForgetPassword 忘记密码
func ForgetPassword(ctx *gin.Context) {
	//ctx.
}
