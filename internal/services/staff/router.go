package staff

import (
	"github.com/XC-Zero/yinwan/internal/gateway/router"
	"github.com/gin-gonic/gin"
)

func InitRouter() {
	router.Router.POST("/staff/create_staff", createStaff)
	router.Router.POST("/staff/update_staff", updateStaff)
	router.Router.POST("/staff/delete_staff", deleteStaff)
	router.Router.POST("/staff/select_staff", selectStaff)
	router.Router.POST("/staff/select_department", selectDepartment)
	router.Router.POST("/staff/create_department", createDepartment)
	router.Router.POST("/staff/update_department", updateDepartment)
	router.Router.POST("/staff/delete_department", deleteDepartment)

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
