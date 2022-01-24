package router

import (
	"github.com/XC-Zero/yinwan/internal/services/access_control"
	"github.com/XC-Zero/yinwan/internal/services/finance"
	"github.com/XC-Zero/yinwan/internal/services/staff"
	"github.com/XC-Zero/yinwan/internal/services/storage"
	"github.com/XC-Zero/yinwan/internal/services/transaction"
	"github.com/gin-gonic/gin"
)

var Router *gin.Engine

func init() {
	Router = gin.Default()
	access_control.InitRouter()
	storage.InitRouter()
	staff.InitRouter()
	finance.InitRouter()
	transaction.InitRouter()
}
