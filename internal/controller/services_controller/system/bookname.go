package system

import (
	"github.com/XC-Zero/yinwan/internal/services/system"
	"github.com/XC-Zero/yinwan/pkg/client"
	_const "github.com/XC-Zero/yinwan/pkg/const"
	"github.com/XC-Zero/yinwan/pkg/utils/errs"
	"github.com/gin-gonic/gin"
)

type CreateBookNameRequest struct {
	BookName string `json:"book_name" binding:"required"`
}

func CreateBookName(ctx *gin.Context) {
	var req CreateBookNameRequest
	err := ctx.BindJSON(&req)
	if err != nil {
		ctx.JSON(_const.EXPECTATION_FAILED_ERROR, errs.CreateWebErrorMsg("请输入账套名称哦"))
		return
	} else {
		if system.AddBookName(req.BookName) {
			ctx.JSON(_const.OK, nil)
			return
		}
		ctx.JSON(_const.INTERNAL_ERROR, errs.CreateWebErrorMsg("创建账套失败！"))
	}
}

func SelectAllBookName(ctx *gin.Context) {
	var bookNameList = make([]string, 0, len(client.BookNameMap))
	for key := range client.BookNameMap {
		bookNameList = append(bookNameList, key)
	}
	ctx.JSON(_const.OK, gin.H{
		"count":          len(bookNameList),
		"book_name_list": bookNameList,
	})
	return
}
