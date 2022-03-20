package staff

import (
	"fmt"
	"github.com/XC-Zero/yinwan/pkg/client"
	_const "github.com/XC-Zero/yinwan/pkg/const"
	"github.com/XC-Zero/yinwan/pkg/model"
	"github.com/XC-Zero/yinwan/pkg/utils/email"
	"github.com/XC-Zero/yinwan/pkg/utils/errs"
	"github.com/XC-Zero/yinwan/pkg/utils/logger"
	"github.com/fwhezfwhez/errorx"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v7"
	"log"
	"math/rand"
	"strconv"
	"time"
)

// CreateStaff
// @Param  model.Staff
func CreateStaff(ctx *gin.Context) {
	temp := model.Staff{}
	err := ctx.ShouldBindJSON(&temp)
	if err != nil {
		mes := "参数有误"
		logger.Error(errorx.MustWrap(err), mes)
		ctx.JSON(_const.REQUEST_PARM_ERROR, errs.CreateWebErrorMsg(mes))
		return
	}
	err = client.MysqlClient.Model(&model.Staff{}).Create(&temp).Error
	if err != nil {
		mes := "创建员工失败！"
		logger.Error(errorx.MustWrap(err), mes)
		ctx.JSON(_const.REQUEST_PARM_ERROR, errs.CreateWebErrorMsg(mes))
		return
	}
	ctx.JSON(_const.OK, errs.CreateWebErrorMsg("创建员工成功!"))
	return
}

// SelectStaff 查询员工
func SelectStaff(ctx *gin.Context) {

	departmentID := ctx.PostForm("department_id")
	staffPosition := ctx.PostForm("staff_position")
	staffRoleId := ctx.PostForm("staff_role_id")
	staffName := ctx.PostForm("staff_name")
	log.Printf("%+v", staffName)

	var staffList []model.Staff

	sql := `select %s from staffs where 1 = 1 `
	if departmentID != "" {
		sql += fmt.Sprintf(" and staff_department_id = '%s'", departmentID)
	}
	if staffPosition != "" {
		sql += fmt.Sprintf(" and staff_position = '%s'", staffPosition)
	}
	if staffRoleId != "" {
		sql += fmt.Sprintf(" and staff_role_id = '%s'", staffRoleId)
	}
	contentSql, countSql := fmt.Sprintf(sql, "*"), fmt.Sprintf(sql, "count(*)")
	// 不能前置，% 符会被 sprintf 解析成 (MISSING)
	if staffName != "" {
		suffixSql := fmt.Sprintf(" and staff_name like '%%%s%%'", staffName)
		contentSql += suffixSql
		countSql += suffixSql
	}
	contentSql += " order by  " + _const.PRIMARY_KEY_NAME + client.PaginateSql(ctx)
	err := client.MysqlClient.Raw(contentSql).Scan(&staffList).Error
	if err != nil {
		logger.Error(errorx.MustWrap(err), "查询时员工失败！")
		ctx.JSON(_const.INTERNAL_ERROR, gin.H(errs.CreateWebErrorMsg("查询时员工内容失败！")))
		return
	}
	count := 0
	err = client.MysqlClient.Raw(countSql).Scan(&count).Error
	if err != nil {
		logger.Error(errorx.MustWrap(err), "查询时员工总数失败！")
		ctx.JSON(_const.INTERNAL_ERROR, gin.H(errs.CreateWebErrorMsg("查询时员工失败！")))
		return
	}
	ctx.JSON(_const.OK, gin.H{
		"count":      count,
		"staff_list": model.IgnoreStaffPassword(staffList),
	})
	return

}

// UpdateStaff todo  !!!
func UpdateStaff(ctx *gin.Context) {
	var staff model.Staff
	err := ctx.ShouldBind(&staff)
	if err != nil {
		ctx.JSON(_const.REQUEST_PARM_ERROR, errs.CreateWebErrorMsg("参数有误"))
		return
	}
	err = client.MysqlClient.Model(&model.Staff{}).Omit("staff_password", "rec_id  ").Updates(staff).Error
	if err != nil {
		mes := fmt.Sprintf("更新职工信息出错，职工ID为 %d ！", staff.RecID)
		logger.Error(errorx.MustWrap(err), mes)
		ctx.JSON(_const.INTERNAL_ERROR, errs.CreateWebErrorMsg(mes))
		return
	}
	ctx.JSON(_const.OK, errs.CreateSuccessMsg("修改职工信息成功！"))
	return
}

// DeleteStaff todo  !!!
func DeleteStaff(ctx *gin.Context) {
	recID := ctx.PostForm("id")
	if recID == "" {
		ctx.JSON(_const.REQUEST_PARM_ERROR, errs.CreateWebErrorMsg("请输入职工ID哦！"))
		return
	}
	err := client.MysqlClient.Delete(&model.Staff{}, recID).Error
	if err != nil {
		mes := fmt.Sprintf("删除职工失败！职工ID: %s!", recID)
		logger.Error(errorx.MustWrap(err), mes)
		ctx.JSON(_const.INTERNAL_ERROR, errs.CreateWebErrorMsg(mes))
		return
	}
	ctx.JSON(_const.OK, errs.CreateSuccessMsg("删除职工成功！"))
	return
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
		mes := "Redis存储邮件验证码失败!"
		logger.Error(errorx.MustWrap(err), mes)
		ctx.JSON(_const.INTERNAL_ERROR, gin.H(errs.CreateWebErrorMsg(mes)))
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
		mes := "发送邮件验证码失败!"
		logger.Error(errorx.MustWrap(err), mes)
		ctx.JSON(_const.INTERNAL_ERROR, gin.H(errs.CreateWebErrorMsg(mes)))
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
		logger.Error(errorx.MustWrap(err), mes)
		ctx.JSON(_const.INTERNAL_ERROR, gin.H(errs.CreateWebErrorMsg(mes)))
		return
	}
	if staffCaptcha == redisCaptcha {
		mes := "邮箱验证通过!"
		logger.Info(mes)
		ctx.JSON(_const.OK, gin.H(errs.CreateSuccessMsg(mes)))
		return
	} else {
		mes := "您的邮箱验证码不正确!"
		logger.Info(mes)
		ctx.JSON(_const.INTERNAL_ERROR, gin.H(errs.CreateWebErrorMsg(mes)))
		return
	}
}

func SelectStaffRole(ctx *gin.Context) {

}
