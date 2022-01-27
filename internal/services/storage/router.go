package storage

import (
	"github.com/gin-gonic/gin"
)

func InitRouter(services *gin.RouterGroup) {
	services.POST("/storage/stack_out", stackOut)
	services.POST("/storage/stack_in", stackIn)
}

func stackOut(ctx *gin.Context) {

}
func stackIn(ctx *gin.Context) {

}
