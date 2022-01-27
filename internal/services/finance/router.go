package finance

import (
	"github.com/gin-gonic/gin"
)

func InitRouter(services *gin.RouterGroup) {
	services.POST("", v)
}

// todo  凭证相关
func v(ctx *gin.Context) {

}
