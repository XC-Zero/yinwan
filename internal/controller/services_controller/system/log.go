package system

import (
	_const "github.com/XC-Zero/yinwan/pkg/const"
	"github.com/XC-Zero/yinwan/pkg/utils/errs"
	"github.com/XC-Zero/yinwan/pkg/utils/logger"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
)

func OperateLog(ctx *gin.Context) {
	ws, err := logger.SocketClient.Upgrade(ctx.Writer, ctx.Request, nil)
	if err != nil {
		logger.Error(errors.WithStack(err), "开启Socket失败!")
		ctx.JSON(_const.INTERNAL_ERROR, errs.CreateWebErrorMsg("读取日志失败!"))
		return
	}
	logger.OperateLogSocketClient = ws
}
