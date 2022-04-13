package common

import (
	"fmt"
	_const "github.com/XC-Zero/yinwan/pkg/const"
	_interface "github.com/XC-Zero/yinwan/pkg/interface"
	"github.com/XC-Zero/yinwan/pkg/utils/errs"
	"github.com/gin-gonic/gin"
)

type PredefinedMessage string

const (
	REQUEST_PARM_ERROR    PredefinedMessage = "您的输入有误！"
	BOOK_NAME_LACK_ERROR  PredefinedMessage = "请先选择账套！"
	DATABASE_DELETE_ERROR PredefinedMessage = "删除%s失败！"
	DATABASE_UPDATE_ERROR PredefinedMessage = "更新%s失败！"
	DATABASE_INSERT_ERROR PredefinedMessage = "添加%s失败！"
	DATABASE_SELECT_ERROR PredefinedMessage = "查询%s内容失败！"
	DATABASE_COUNT_ERROR  PredefinedMessage = "查询%s总数失败！"
	OTHER_ERROR           PredefinedMessage = "其他问题导致操作%s失败"
)

// RequestParamErrorTemplate 入参异常
func RequestParamErrorTemplate(ctx *gin.Context, message PredefinedMessage) {
	if message == "" {
		message = REQUEST_PARM_ERROR
	}
	ctx.JSON(_const.REQUEST_PARM_ERROR, errs.CreateWebErrorMsg(string(message)))
	return
}

// InternalDataBaseErrorTemplate 数据库异常
func InternalDataBaseErrorTemplate(ctx *gin.Context, message PredefinedMessage, table _interface.ChineseTabler) {
	ctx.JSON(_const.INTERNAL_ERROR, errs.CreateWebErrorMsg(fmt.Sprintf(string(message), table.TableCnName())))
	return
}

// SelectSuccessTemplate 普适的查询返回
func SelectSuccessTemplate(ctx *gin.Context, count int64, dataList interface{}) {
	ctx.JSON(_const.OK, gin.H{
		"count": count,
		"list":  dataList,
	})
	return
}
