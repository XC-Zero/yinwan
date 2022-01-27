package router

import (
	"github.com/XC-Zero/yinwan/internal/services/access_control"
	"github.com/XC-Zero/yinwan/internal/services/finance"
	"github.com/XC-Zero/yinwan/internal/services/staff"
	"github.com/XC-Zero/yinwan/internal/services/storage"
	"github.com/XC-Zero/yinwan/internal/services/transaction"
	"github.com/gin-gonic/gin"
)

func init() {
	router := gin.Default()
	access_control.InitRouter(router)

	services := router.Group("/services", auth)

	storage.InitRouter(services)
	staff.InitRouter(services)
	finance.InitRouter(services)
	transaction.InitRouter(services)
}

func auth(ctx *gin.Context) {
	//token := ctx.Request.Header.Get("token")
	//parse, err := jwt.Parse(token)
	//if err != nil {
	//	return
	//}
}
