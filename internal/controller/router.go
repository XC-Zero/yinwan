package controller

import (
	"github.com/XC-Zero/yinwan/internal/config"
	"github.com/XC-Zero/yinwan/internal/controller/access_control"
	finance2 "github.com/XC-Zero/yinwan/internal/controller/finance"
	staff2 "github.com/XC-Zero/yinwan/internal/controller/staff"
	storage2 "github.com/XC-Zero/yinwan/internal/controller/storage"

	"github.com/XC-Zero/yinwan/pkg/utils/token"
	"github.com/gin-gonic/gin"
)

func Starter() {

	router := gin.Default()
	router.POST("/login", access_control.Login)
	router.POST("/logout", access_control.Logout)

	// 使用组路由，并添加中间件用于判断token
	services := router.Group("/services", auth)

	staff := services.Group("/staff")
	{
		staff.POST("/create_staff", staff2.CreateStaff)
		staff.POST("/delete_staff", staff2.DeleteStaff)
		staff.POST("/update_staff", staff2.UpdateStaff)
		staff.POST("/select_staff", staff2.SelectStaff)
		staff.POST("/validate_staff_email", staff2.ValidateStaffEmail)
	}

	department := services.Group("/department")
	{
		department.POST("/select_department", staff2.SelectDepartment)
		department.POST("/create_department", staff2.CreateDepartment)
		department.POST("/update_department", staff2.UpdateDepartment)
		department.POST("/delete_department", staff2.DeleteDepartment)
	}

	storage := services.Group("/storage")
	{
		storage.POST("/stack_out", storage2.StockIn)
		storage.POST("/stack_in", storage2.StockOut)
		storage.POST("/scan_qrcode", storage2.ScanQRCode)
		storage.POST("/create_qrcode", storage2.CreateQRCode)
		storage.POST("/download_qrcode", storage2.DownloadQRCode)
	}

	finance := services.Group("/finance")
	{
		finance.POST("/create_credential", finance2.CreateCredential)
	}

	transaction := services.Group("/transaction")
	{
		transaction.POST("")
	}

	err := router.Run(":" + config.CONFIG.ServiceConfig.Port)
	if err != nil {
		panic(err)
	}
}

func auth(ctx *gin.Context) {
	tokenStr := ctx.Request.Header.Get("token")
	if !token.IsExpired(tokenStr) {
		ctx.Abort()
	}
	ctx.Next()
}
