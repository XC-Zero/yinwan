package access_control

import (
	_const "github.com/XC-Zero/yinwan/pkg/const"
	"github.com/XC-Zero/yinwan/pkg/model"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

func Login(ctx *gin.Context) {
	staff := model.Staff{}
	err := ctx.ShouldBindWith(&staff, binding.JSON)
	if err != nil {
		ctx.JSON(_const.REQUEST_PARM_ERROR, gin.H{
			"?": "?",
		})
	}
	tokenPtr, errMessage := staff.Login()
	if tokenPtr != nil {
		ctx.JSON(_const.REQUEST_PARM_ERROR, errMessage)
	} else {
		ctx.JSON(_const.OK, gin.H{
			"token": *tokenPtr,
		})
	}

}

func Logout(ctx *gin.Context) {

}
