package access_control

import (
	"encoding/json"
	"github.com/XC-Zero/yinwan/internal/config"
	"github.com/XC-Zero/yinwan/internal/controller/services_controller/common"
	"github.com/XC-Zero/yinwan/pkg/client"
	_const "github.com/XC-Zero/yinwan/pkg/const"
	"github.com/XC-Zero/yinwan/pkg/model/mysql_model"
	"github.com/XC-Zero/yinwan/pkg/utils/encode"
	"github.com/XC-Zero/yinwan/pkg/utils/errs"
	"github.com/XC-Zero/yinwan/pkg/utils/logger"
	"github.com/XC-Zero/yinwan/pkg/utils/mysql"
	"github.com/dollarkillerx/urllib"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/pkg/errors"
	"io/ioutil"
	"log"
)

// Login
// @Param  staff_email string,staff_password
func Login(ctx *gin.Context) {
	staffEmail, staffPassword := ctx.PostForm("staff_email"), ctx.PostForm("staff_password")
	staff := mysql_model.Staff{
		StaffEmail:    staffEmail,
		StaffPassword: staffPassword,
	}

	tokenPtr, errMessage, _ := staff.Login()
	if tokenPtr == nil {
		ctx.JSON(_const.REQUEST_PARM_ERROR, gin.H(errs.CreateWebErrorMsg(errMessage)))
	} else {
		var rc []mysql_model.RoleCapabilities
		role := mysql_model.Role{}
		err1 := client.MysqlClient.Model(&mysql_model.Role{}).Where("rec_id = (select staff_role_id from staffs where staff_email = ? limit 1) ", staffEmail).Find(&role).Error
		err2 := client.MysqlClient.Raw("select * from "+mysql_model.RoleCapabilities{}.TableName()+" where role_id = (select staff_role_id from staffs where staff_email = ? limit 1)", staffEmail).Scan(&rc).Error

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
		err := client.MysqlClient.Model(&mysql_model.Staff{}).
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

type wxLoginRequest struct {
	JsCode string `json:"js_code" bson:"js_code"  form:"js_code"`
}

type wxLoginResponse struct {
	OpenID     string `json:"open_id" bson:"open_id"  form:"open_id"`
	SessionKey string `json:"session_key" form:"session_key" bson:"session_key"`
	UnionID    string `json:"unionid" form:"unionid" bson:"unionid"`
	ErrCode    string `json:"errcode" form:"errcode" bson:"errcode"`
	ErrMsg     string `json:"errmsg" form:"errmsg" bson:"errmsg"`
}

func WxLogin(ctx *gin.Context) {
	wxConfig := config.CONFIG.ApiConfig.WxConfig
	if wxConfig.AppSecrete == "" || wxConfig.AppID == "" || wxConfig.ApiUrl == "" {
		logger.Error(errors.WithStack(errors.New("There have not wx config ! ")), "微信小程序暂未配置!!")
	}
	var temp wxLoginRequest
	err := ctx.ShouldBindBodyWith(&temp, binding.JSON)
	if err != nil {
		logger.Error(errors.WithStack(err), "绑定微信登录请求体失败! ")
		common.RequestParamErrorTemplate(ctx, common.REQUEST_PARM_ERROR)
		return
	}

	var requestWechatMap = make(map[string]string, 4)
	requestWechatMap["js_code"] = temp.JsCode
	requestWechatMap["appid"] = wxConfig.AppID
	requestWechatMap["secret"] = wxConfig.AppSecrete
	requestWechatMap["grant_type"] = "authorization_code"
	res, err := urllib.Get(wxConfig.ApiUrl).ParamsMap(requestWechatMap).AlloverTLS().Body()
	if err != nil {
		logger.Error(errors.WithStack(err), "绑定微信登录请求体失败! ")
		ctx.JSON(_const.INTERNAL_ERROR, errs.CreateWebErrorMsg("请求微信小程序接口失败!"))
		return
	}

	all, err := ioutil.ReadAll(res.Body)
	if err != nil {
		logger.Error(errors.WithStack(err), "读取小程序返回接口body失败! ")
		ctx.JSON(_const.INTERNAL_ERROR, errs.CreateWebErrorMsg("请求微信小程序接口失败!"))
		return
	}
	var response wxLoginResponse
	err = json.Unmarshal(all, &response)
	if err != nil {
		logger.Error(errors.WithStack(err), "绑定微信登录返回值失败! ")
		ctx.JSON(_const.INTERNAL_ERROR, errs.CreateWebErrorMsg("请求微信小程序接口失败!"))
		return
	}
	data := common.SelectMysqlTableContentWithCountTemplate(ctx, common.SelectMysqlTemplateOptions{
		DB:            client.MysqlClient,
		TableModel:    mysql_model.WxStaffRelationship{},
		OrderByColumn: "",
		ResHookFunc:   nil,
		NotReturn:     true,
		NotPaginate:   true,
	}, common.MysqlCondition{
		Symbol:      mysql.EQUAL,
		ColumnName:  "open_id",
		ColumnValue: response.OpenID,
	}, common.MysqlCondition{
		Symbol:      mysql.NOT_NULL,
		ColumnName:  "deleted_at",
		ColumnValue: " ",
	})
	relationships := data.([]mysql_model.WxStaffRelationship)

	tempResponse := struct {
		Token      string `json:"token" form:"token"`
		IsNewStaff bool   `json:"is_new_staff" form:"is_new_staff"`
		wxLoginResponse
	}{
		wxLoginResponse: response,
	}
	if len(relationships) == 0 {
		tempResponse.IsNewStaff = true
	} else {
		var staff mysql_model.Staff
		err := client.MysqlClient.Model(mysql_model.Staff{}).Where("rec_id = ?", relationships[0].StaffID).First(&staff).Error
		if err != nil {
			logger.Error(errors.WithStack(err), "微信小程序登录绑定时,查无此人!")
			tempResponse.IsNewStaff = true
			ctx.JSON(_const.OK, tempResponse)
			return
		}
		tk, errMsg, _ := staff.Login()
		if tk == nil {
			logger.Error(errors.WithStack(err), "此人登录失败!")
			ctx.JSON(_const.INTERNAL_ERROR, errs.CreateWebErrorMsg(errMsg))
			return
		} else {
			tempResponse.Token = *tk
		}
	}
	ctx.JSON(_const.OK, tempResponse)
	return
}

type wxBindRequest struct {
	StaffEmail    string `json:"staff_email" form:"staff_email"`
	StaffPassword string `json:"staff_password" form:"staff_password"`
	OpenID        string `json:"open_id" form:"open_id"`
}

func WxBindingStaff(ctx *gin.Context) {
	var temp wxBindRequest
	err := ctx.ShouldBindBodyWith(&temp, binding.JSON)
	if err != nil {
		logger.Error(errors.WithStack(err), "微信绑定系统账号请求体失败! ")
		common.RequestParamErrorTemplate(ctx, common.REQUEST_PARM_ERROR)
		return
	}
	staff := mysql_model.Staff{StaffEmail: temp.StaffEmail, StaffPassword: temp.StaffPassword}
	tk, errMsg, id := staff.Login()
	if tk == nil {
		logger.Error(errors.WithStack(err), "此人登录失败!")
		ctx.JSON(_const.INTERNAL_ERROR, errs.CreateWebErrorMsg(errMsg))
		return
	} else {
		err = client.MysqlClient.Model(mysql_model.WxStaffRelationship{}).Create(&mysql_model.WxStaffRelationship{
			StaffID: *id,
			OpenID:  temp.OpenID,
		}).Error
		if err != nil {
			common.InternalDataBaseErrorTemplate(ctx, common.DATABASE_INSERT_ERROR, mysql_model.WxStaffRelationship{})
			return
		}

	}
	ctx.JSON(_const.OK, gin.H{
		"token":   *tk,
		"status":  "success",
		"message": "绑定成功!",
	})
	return
}

// RemoveBindingStaff todo 解绑账号
func RemoveBindingStaff(ctx *gin.Context) {

}
