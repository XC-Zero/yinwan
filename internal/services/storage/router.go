package storage

import (
	"github.com/XC-Zero/yinwan/internal/gateway/router"
	"github.com/gin-gonic/gin"
)

func InitRouter() {
	router.Router.POST("", stackOut)
	router.Router.POST("", stackIn)
}

func stackOut(ctx *gin.Context) {

}
func stackIn(ctx *gin.Context) {

}
