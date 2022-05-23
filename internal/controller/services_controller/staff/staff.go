package staff

import (
	"github.com/XC-Zero/yinwan/internal/controller/services_controller/common"
	"github.com/XC-Zero/yinwan/pkg/client"
	_const "github.com/XC-Zero/yinwan/pkg/const"
	"github.com/XC-Zero/yinwan/pkg/model/mysql_model"
	"github.com/XC-Zero/yinwan/pkg/utils/email"
	"github.com/XC-Zero/yinwan/pkg/utils/encode"
	"github.com/XC-Zero/yinwan/pkg/utils/errs"
	"github.com/XC-Zero/yinwan/pkg/utils/logger"
	"github.com/XC-Zero/yinwan/pkg/utils/mysql"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-redis/redis/v7"
	"github.com/pkg/errors"
	"math/rand"
	"strconv"
	"time"
)

// CreateStaff
// @Param  model.Staff
func CreateStaff(ctx *gin.Context) {
	staff := mysql_model.Staff{}
	err := ctx.ShouldBindBodyWith(&staff, binding.JSON)
	if err != nil {
		logger.Error(errors.WithStack(err), "")
		ctx.JSON(_const.REQUEST_PARM_ERROR, errs.CreateWebErrorMsg("参数有误"))
		return
	}
	err = client.MysqlClient.Model(&mysql_model.Staff{}).Create(&staff).Error
	if err != nil {
		logger.Error(errors.WithStack(err), "")
		ctx.JSON(_const.REQUEST_PARM_ERROR, errs.CreateWebErrorMsg("创建员工失败！"))
		return
	}
	ctx.JSON(_const.OK, errs.CreateSuccessMsg("创建员工成功!"))
	return
}

// SelectStaff 查询员工
func SelectStaff(ctx *gin.Context) {
	conditions := []common.MysqlCondition{
		{
			Symbol:      mysql.LIKE,
			ColumnName:  "staff_name",
			ColumnValue: ctx.PostForm("staff_name"),
		},
		{
			Symbol:      mysql.LIKE,
			ColumnName:  "staff_position",
			ColumnValue: ctx.PostForm("staff_position"),
		},
		{
			Symbol:      mysql.EQUAL,
			ColumnName:  "staff_department_id",
			ColumnValue: ctx.PostForm("department_id"),
		},
		{
			Symbol:      mysql.EQUAL,
			ColumnName:  "staff_role_id",
			ColumnValue: ctx.PostForm("staff_role_id"),
		},
		{
			Symbol:      mysql.NULL,
			ColumnName:  "deleted_at",
			ColumnValue: " ",
		},
	}
	op := common.SelectMysqlTemplateOptions{
		DB:          client.MysqlClient,
		TableModel:  mysql_model.Staff{},
		ResHookFunc: mysql_model.IgnoreStaffPassword,
	}
	common.SelectMysqlTableContentWithCountTemplate(ctx, op, conditions...)

	return

}

// UpdateStaff 更新员工
func UpdateStaff(ctx *gin.Context) {
	var staff mysql_model.Staff
	err := ctx.ShouldBindBodyWith(&staff, binding.JSON)
	staff.StaffPassword = ""
	if err != nil {
		logger.Error(errors.WithStack(err), "")
		common.RequestParamErrorTemplate(ctx, common.REQUEST_PARM_ERROR)
		return
	}
	err = client.MysqlClient.Model(&mysql_model.Staff{}).Updates(staff).Omit("staff_email", "staff_password").Where("rec_id", staff.RecID).Error
	if err != nil {
		logger.Error(errors.WithStack(err), "")
		common.InternalDataBaseErrorTemplate(ctx, common.DATABASE_UPDATE_ERROR, staff)
		return
	}
	ctx.JSON(_const.OK, errs.CreateSuccessMsg("修改职工信息成功！"))
	return
}

// DeleteStaff 删除员工
func DeleteStaff(ctx *gin.Context) {
	recID := ctx.PostForm("staff_id")
	if recID == "" {
		common.RequestParamErrorTemplate(ctx, common.REQUEST_PARM_ERROR)
		return
	}
	err := client.MysqlClient.Delete(&mysql_model.Staff{}, recID).Error
	if err != nil {
		logger.Error(errors.WithStack(err), "")
		common.InternalDataBaseErrorTemplate(ctx, common.DATABASE_UPDATE_ERROR, mysql_model.Staff{})
		return
	}
	ctx.JSON(_const.OK, errs.CreateSuccessMsg("删除职工成功！"))
	return
}

// UpdatePersonalInfo TODO 读请求头验证是不是本人
func UpdatePersonalInfo(ctx *gin.Context) {

}

func SendStaffValidateEmail(ctx *gin.Context) {
	staffEmail := ctx.PostForm("staff_email")
	if len(staffEmail) == 0 {
		ctx.JSON(_const.REQUEST_PARM_ERROR, gin.H(errs.CreateWebErrorMsg("未输入邮箱哦")))
	}
	rand.Seed(time.Now().Unix())
	n := strconv.Itoa(rand.Intn(999999))
	err := client.RedisClient.Set(staffEmail, n, time.Minute*10).Err()
	if err != nil {
		logger.Error(errors.WithStack(err), "")
		ctx.JSON(_const.INTERNAL_ERROR, errs.CreateWebErrorMsg("Redis存储邮件验证码失败!"))
		return
	}
	err = email.SendEmail("好上好系统验证邮件", "<div style=\"width: 100vw;height: 100vh;background-color: #ffffff;display: flex;flex-direction: column; justify-content: center;align-items: center\">"+
		"<div style=\"width: 80%;text-align:left;\"><h1 style=\"color: #4F709B\">HSH</h1>"+
		"<h2>欢迎您，您的验证码为：</h2></div>"+
		"<div style=\" margin-top:50px;margin-bottom:50px;background-color: #ffffff;width: 300px ;height: 150px;box-shadow: 0px 3px 6px #999999;border-radius: 5px ;line-height:150px;text-align:center;letter-spacing:5px;font-size: 40px;\">"+
		n+
		"</div>"+
		"<h3 style=\"width: 80%;text-align:right;\">验证码十分钟有效~</h3>"+
		"</div>", staffEmail)
	if err != nil {
		logger.Error(errors.WithStack(err), "")
		ctx.JSON(_const.INTERNAL_ERROR, gin.H(errs.CreateWebErrorMsg("发送邮件验证码失败!")))
		return
	}
	ctx.JSON(_const.OK, gin.H(errs.CreateSuccessMsg("邮箱验证码发送成功！")))
	return
}

// ValidateStaffEmail
// @Description 验证邮箱验证码
// @Accept  json
// @Produce  json
// @Router  /services/staff/validate_staff_email  [post]
// @Param  staff_email string "用户邮箱",captcha string  "用户输入的验证码"
// @Success 200
// @Failure otherCode
func ValidateStaffEmail(ctx *gin.Context) {
	staffEmail := ctx.PostForm("staff_email")
	if len(staffEmail) == 0 {
		ctx.JSON(_const.REQUEST_PARM_ERROR, gin.H(errs.CreateWebErrorMsg("未输入邮箱哦")))
		return
	}
	staffCaptcha := ctx.PostForm("captcha")
	redisCaptcha, err := client.RedisClient.Get(staffEmail).Result()
	if err == redis.Nil {
		mes := "邮件验证码不存在或已过期!"
		logger.Info(mes)
		ctx.JSON(_const.INTERNAL_ERROR, gin.H(errs.CreateWebErrorMsg(mes)))
		return
	} else if err != nil {
		mes := "Redis 读取失败!"
		logger.Error(errors.WithStack(err), "")
		ctx.JSON(_const.INTERNAL_ERROR, gin.H(errs.CreateWebErrorMsg(mes)))
		return
	}
	if staffCaptcha == redisCaptcha {
		aes, err := encode.EncryptByAes(staffEmail)
		if err != nil {
			logger.Error(errors.WithStack(err), "")
		}
		ctx.JSON(_const.OK, gin.H(errs.CreateSuccessMsg("邮箱验证通过!", map[string]interface{}{
			"secret_key": aes,
		})))
		return
	} else {
		mes := "您的邮箱验证码不正确!"
		ctx.JSON(_const.INTERNAL_ERROR, gin.H(errs.CreateWebErrorMsg(mes)))
		return
	}
}

func SelectStaffRole(ctx *gin.Context) {

}
