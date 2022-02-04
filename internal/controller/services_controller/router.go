package services_controller

import (
	"github.com/XC-Zero/yinwan/internal/config"
	"github.com/XC-Zero/yinwan/internal/controller/services_controller/access_control"
	finance2 "github.com/XC-Zero/yinwan/internal/controller/services_controller/finance"
	staff3 "github.com/XC-Zero/yinwan/internal/controller/services_controller/staff"
	storage3 "github.com/XC-Zero/yinwan/internal/controller/services_controller/storage"
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
		staff.POST("/create_staff", staff3.CreateStaff)
		staff.POST("/delete_staff", staff3.DeleteStaff)
		staff.POST("/update_staff", staff3.UpdateStaff)
		staff.POST("/select_staff", staff3.SelectStaff)
		staff.POST("/validate_staff_email", staff3.ValidateStaffEmail)
	}

	department := services.Group("/department")
	{
		department.POST("/select_department", staff3.SelectDepartment)
		department.POST("/create_department", staff3.CreateDepartment)
		department.POST("/update_department", staff3.UpdateDepartment)
		department.POST("/delete_department", staff3.DeleteDepartment)
	}

	storage := services.Group("/storage")
	{
		storage.POST("/stack_out", storage3.StockIn)
		storage.POST("/stack_in", storage3.StockOut)
		storage.POST("/scan_qrcode", storage3.ScanQRCode)
		storage.POST("/create_qrcode", storage3.CreateQRCode)
		storage.POST("/download_qrcode", storage3.DownloadQRCode)
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
