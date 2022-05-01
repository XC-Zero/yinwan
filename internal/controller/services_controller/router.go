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
		{
			staff.POST("/create_staff", staff3.CreateStaff)
			staff.POST("/select_staff", staff3.SelectStaff)
			staff.POST("/update_staff", staff3.UpdateStaff)
			staff.POST("/delete_staff", staff3.DeleteStaff)
			staff.POST("/update_password", staff3.UpdatePassword)
		}
		{
			staff.POST("/create_department", staff3.CreateDepartment)
			staff.POST("/select_department", staff3.SelectDepartment)
			staff.POST("/update_department", staff3.UpdateDepartment)
			staff.POST("/delete_department", staff3.DeleteDepartment)
		}
	}

	storage := services.Group("/storage")
	{
		// 原材料
		{
			storage.POST("/create_material", storage3.CreateMaterial)
			storage.POST("/select_material", storage3.SelectMaterial)
			storage.POST("/update_material", storage3.UpdateMaterial)
			storage.POST("/delete_material", storage3.DeleteMaterial)
		}
		// 原材料详情
		{
			storage.POST("/select_material_detail", storage3.SelectMaterialDetail)

		}
		// 供应商
		{
			storage.POST("/create_provider", storage3.CreateProvider)
			storage.POST("/select_provider", storage3.SelectProvider)
			storage.POST("/update_provider", storage3.UpdateProvider)
			storage.POST("/delete_provider", storage3.DeleteProvider)
		}
		// 采购
		{
			storage.POST("/create_purchase", storage3.CreatePurchase)
			storage.POST("/select_purchase", storage3.SelectPurchase)
			storage.POST("/update_purchase", storage3.UpdatePurchase)
			storage.POST("/delete_purchase", storage3.DeletePurchase)
		}
		// 入库
		{
			storage.POST("/create_stock_in", storage3.CreateStockIn)
			storage.POST("/select_stock_in", storage3.SelectStockIn)
			storage.POST("/update_stock_in", storage3.UpdateStockIn)
			storage.POST("/delete_stock_in", storage3.DeleteStockIn)

		}
		// 出库
		{
			storage.POST("/create_stock_out", storage3.CreateStockOut)
			storage.POST("/select_stock_out", storage3.SelectStockOut)
			storage.POST("/update_stock_out", storage3.UpdateStockOut)
			storage.POST("/delete_stock_out", storage3.DeleteStockOut)

		}
		//扫码 小程序用
		{
			storage.POST("/scan_qrcode", storage3.ScanQRCode)
			storage.POST("/create_qrcode", storage3.CreateQRCode)
			storage.POST("/download_qrcode", storage3.DownloadQRCode)
		}

	}

	finance := services.Group("/finance")
	{
		{
			finance.POST("/select_credential", finance3.SelectCredential)
			finance.POST("/create_credential", finance3.CreateCredential)
			finance.POST("/update_credential", finance3.UpdateCredential)
			finance.POST("/delete_credential", finance3.DeleteCredential)
		}
		{
			finance.POST("/select_fixed_asset", finance3.SelectFixedAsset)
			finance.POST("/create_fixed_asset", finance3.CreateFixedAsset)
			finance.POST("/update_fixed_asset", finance3.UpdateFixedAsset)
			finance.POST("/delete_fixed_asset", finance3.DeleteFixedAsset)
		}
	}

	transaction := services.Group("/transaction")
	{
		transaction.POST("")
	}
	system := services.Group("/system")
	{
		system.POST("create_book_name", system2.CreateBookName)
		system.POST("select_all_book_name", system2.SelectAllBookName)
		system.POST("select_module", system2.SelectModule)
		// 角色
		{
			system.POST("/create_role", system2.CreateRole)
			system.POST("/select_role", system2.SelectRole)
			system.POST("/update_role", system2.UpdateRole)
			system.POST("/delete_role", system2.DeleteRole)
		}
		// 类型
		{
			system.POST("/create_type_tree", system2.CreateTypeTree)
			system.POST("/select_type_tree", system2.SelectTypeTree)
			system.POST("/update_type_tree", system2.UpdateTypeTree)
			system.POST("/delete_type_tree", system2.DeleteTypeTree)
		}
		{
			system.POST("/select_material", system2.SelectMaterial)
			system.POST("/select_payable", system2.SelectPayable)
			system.POST("/select_receivable", system2.SelectReceivable)

		}

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
