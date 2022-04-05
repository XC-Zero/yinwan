package storage

import (
	"github.com/XC-Zero/yinwan/internal/controller/services_controller/common"
	"github.com/XC-Zero/yinwan/pkg/model"
	"github.com/gin-gonic/gin"
)

func CreatePurchase(ctx *gin.Context) {
	var purchase model.Purchase
	err := ctx.ShouldBind(&purchase)
	if err != nil {
		common.RequestParamErrorTemplate(ctx, common.REQUEST_PARM_ERROR)
		return
	}
	op := common.CreateMongoDBTemplateOptions{
		DB:         nil,
		TableModel: nil,
		PreFunc:    nil,
	}
	common.CreateOneMongoDBRecordTemplate(ctx, op)
}
func SelectPurchase(ctx *gin.Context) {

}
func UpdatePurchase(ctx *gin.Context) {

}
func DeletePurchase(ctx *gin.Context) {

}
