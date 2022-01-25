package storage

import (
	"github.com/XC-Zero/yinwan/internal/gateway/router"
	"github.com/gin-gonic/gin"
)

func InitRouter() {
	router.Router.POST("/storage/stack_out", stackOut)
	router.Router.POST("/storage/stack_in", stackIn)
}

func stackOut(ctx *gin.Context) {

}
func stackIn(ctx *gin.Context) {

}
