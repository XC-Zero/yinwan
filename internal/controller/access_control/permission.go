package access_control

import (
	"github.com/XC-Zero/yinwan/pkg/model"
	"github.com/XC-Zero/yinwan/pkg/utils/errs"
	"github.com/gin-gonic/gin"
)

func Login(ctx *gin.Context) {
	staff := model.Staff{}
	err := ctx.ShouldBindJSON(&staff)
	if err != nil {
		ctx.JSON(errs.NO_AUTH_ERROR, gin.H{
			"?": "?",
		})
	}
	tokenPtr, errMessage := staff.Login()
	if tokenPtr != nil {
		ctx.JSON(errs.NO_AUTH_ERROR, errMessage)
	} else {
		ctx.JSON(errs.SUCCESS, gin.H{
			"token": *tokenPtr,
		})
	}

}

func Logout(ctx *gin.Context) {

}
