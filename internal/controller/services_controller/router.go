package services_controller

import (
	"github.com/XC-Zero/yinwan/internal/config"
	"github.com/XC-Zero/yinwan/internal/controller/services_controller/access_control"
	finance3 "github.com/XC-Zero/yinwan/internal/controller/services_controller/finance"
	staff3 "github.com/XC-Zero/yinwan/internal/controller/services_controller/staff"
	storage3 "github.com/XC-Zero/yinwan/internal/controller/services_controller/storage"
	system2 "github.com/XC-Zero/yinwan/internal/controller/services_controller/system"
	_const "github.com/XC-Zero/yinwan/pkg/const"
	"github.com/XC-Zero/yinwan/pkg/utils/errs"
	"github.com/XC-Zero/yinwan/pkg/utils/token"
	"github.com/gin-gonic/gin"
)

func Starter() {

	router := gin.Default()
	//router.Use() todo 使用自定义的日志输出！
	router.POST("/login", access_control.Login)
	router.POST("/forget_password", access_control.ForgetPassword)
	router.POST("/send_to_staff_email", staff3.SendStaffValidateEmail)
	router.POST("/validate_staff_email", staff3.ValidateStaffEmail)
	// 使用组路由，并添加中间件用于判断token
	services := router.Group("/services", auth)

	staff := services.Group("/staff")
	{
		staff.POST("/create_staff", staff3.CreateStaff)
		staff.POST("/select_staff", staff3.SelectStaff)
		staff.POST("/update_staff", staff3.UpdateStaff)
		staff.POST("/delete_staff", staff3.DeleteStaff)

		staff.POST("/create_department", staff3.CreateDepartment)
		staff.POST("/select_department", staff3.SelectDepartment)
		staff.POST("/update_department", staff3.UpdateDepartment)
		staff.POST("/delete_department", staff3.DeleteDepartment)
	}

	storage := services.Group("/storage")
	{

		storage.POST("/select_material", storage3.SelectMaterial)

		storage.POST("/select_material_detail", storage3.SelectMaterialDetail)

		storage.POST("/stack_out", storage3.CreateStockIn)
		storage.POST("/stack_in", storage3.CreateStockOut)
		storage.POST("/scan_qrcode", storage3.ScanQRCode)
		storage.POST("/create_qrcode", storage3.CreateQRCode)
		storage.POST("/download_qrcode", storage3.DownloadQRCode)
	}

	finance := services.Group("/finance")
	{
		finance.POST("/select_credential", finance3.SelectCredential)
		finance.POST("/create_credential", finance3.CreateCredential)
		finance.POST("/update_credential", finance3.UpdateCredential)
		finance.POST("/delete_credential", finance3.DeleteCredential)
	}

	transaction := services.Group("/transaction")
	{
		transaction.POST("")
	}
	system := services.Group("/system")
	{
		system.POST("create_book_name", system2.CreateBookName)
		system.POST("select_all_book_name", system2.SelectAllBookName)

	}
	err := router.Run(":" + config.CONFIG.ServiceConfig.Port)
	if err != nil {
		panic(err)
	}
}

func auth(ctx *gin.Context) {
	tokenStr := ctx.Request.Header.Get("token")
	staffEmail := ctx.Request.Header.Get("staff_email")
	if token.IsExpired(tokenStr, staffEmail) {
		ctx.JSON(_const.UNAUTHORIZED_ERROR, gin.H(errs.CreateWebErrorMsg("登录过期了哦，重新登录呢")))
		ctx.Abort()
	}
	ctx.Next()
}
