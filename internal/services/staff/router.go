package staff

import (
	"github.com/gin-gonic/gin"
)

func InitRouter(services *gin.RouterGroup) {
	// 增
	services.POST("/staff/create_staff", createStaff)
	// 删
	services.POST("/staff/delete_staff", deleteStaff)
	// 改
	services.POST("/staff/update_staff", updateStaff)
	// 查
	services.POST("/staff/select_staff", selectStaff)
	// 校验邮箱
	services.POST("/staff/validate_staff_email", validateStaffEmail)
	services.POST("/staff/select_department", selectDepartment)
	services.POST("/staff/create_department", createDepartment)
	services.POST("/staff/update_department", updateDepartment)
	services.POST("/staff/delete_department", deleteDepartment)

}

func createStaff(ctx *gin.Context) {

}
func selectStaff(ctx *gin.Context) {

}
func updateStaff(ctx *gin.Context) {

}

func deleteStaff(ctx *gin.Context) {

}
func validateStaffEmail(ctx *gin.Context) {

}

func selectDepartment(ctx *gin.Context) {

}

func createDepartment(ctx *gin.Context) {

}

func updateDepartment(ctx *gin.Context) {
}

func deleteDepartment(ctx *gin.Context) {

}
