package access_control

import (
	"github.com/XC-Zero/yinwan/pkg/model"
	"github.com/gin-gonic/gin"
)

func InitRouter(router *gin.Engine) {
	router.POST("login", login)
}

func login(ctx *gin.Context) {
	userID := ctx.Param("userID")
	userName := ctx.Param("userName")
	userPhone := ctx.Param("userPhone")
	pass := ctx.Param("password")
	staff := model.Staff{
		BasicModel:    model.BasicModel{},
		StaffName:     "",
		StaffAlias:    nil,
		StaffEmail:    nil,
		StaffPhone:    nil,
		StaffPassword: "",
		StaffRoleID:   0,
	}
}
