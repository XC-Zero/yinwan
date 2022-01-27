package staff

import (
	"github.com/gin-gonic/gin"
)

func InitRouter(services *gin.RouterGroup) {
	services.POST("/staff/create_staff", createStaff)
	services.POST("/staff/update_staff", updateStaff)
	services.POST("/staff/delete_staff", deleteStaff)
	services.POST("/staff/select_staff", selectStaff)
	services.POST("/staff/select_department", selectDepartment)
	services.POST("/staff/create_department", createDepartment)
	services.POST("/staff/update_department", updateDepartment)
	services.POST("/staff/delete_department", deleteDepartment)

}

func createStaff(ctx *gin.Context) {

}

func updateStaff(ctx *gin.Context) {

}
func selectStaff(ctx *gin.Context) {

}
func deleteStaff(ctx *gin.Context) {

}

func selectDepartment(ctx *gin.Context) {

}

func createDepartment(ctx *gin.Context) {

}

func updateDepartment(ctx *gin.Context) {
}

func deleteDepartment(ctx *gin.Context) {

}
