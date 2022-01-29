package router

import (
	"github.com/XC-Zero/yinwan/internal/config"
	"github.com/XC-Zero/yinwan/internal/services/access_control"
	"github.com/XC-Zero/yinwan/internal/services/finance"
	"github.com/XC-Zero/yinwan/internal/services/staff"
	"github.com/XC-Zero/yinwan/internal/services/storage"
	"github.com/XC-Zero/yinwan/internal/services/transaction"
	"github.com/XC-Zero/yinwan/pkg/utils/token"
	"github.com/gin-gonic/gin"
)

func Starter() {
	router := gin.Default()
	access_control.InitRouter(router)
	// 使用组路由，并添加中间件用于判断token
	services := router.Group("/services", auth)

	storage.InitRouter(services)
	staff.InitRouter(services)
	finance.InitRouter(services)
	transaction.InitRouter(services)
	router.Run(":" + config.CONFIG.ServiceConfig.Port)
}

func auth(ctx *gin.Context) {
	tokenStr := ctx.Request.Header.Get("token")
	if !token.IsExpired(tokenStr) {
		ctx.Abort()
	}
	ctx.Next()
}
