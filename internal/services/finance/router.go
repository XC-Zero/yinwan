package finance

import (
	"github.com/XC-Zero/yinwan/internal/gateway/router"
	"github.com/gin-gonic/gin"
)

func InitRouter() {
	router.Router.POST("", v)
}

// todo  凭证相关
func v(ctx *gin.Context) {

}
