package finance

import (
	"github.com/XC-Zero/yinwan/internal/controller/services_controller/common"
	"github.com/XC-Zero/yinwan/pkg/client"
	"github.com/XC-Zero/yinwan/pkg/model"
	"github.com/gin-gonic/gin"
)

func CreatePayable(ctx *gin.Context) {

}

func SelectPayable(ctx *gin.Context) {
	//bk := client.HarvestClientFromGinContext(ctx)
	//if bk == nil {
	//	common.RequestParamErrorTemplate(ctx, common.BOOK_NAME_LACK_ERROR)
	//	return
	//}
	condition := []common.MongoCondition{
		{},
	}
	op := common.SelectMongoDBTemplateOptions{
		DB:         client.MongoDBClient,
		TableModel: model.Payable{},
	}
	common.SelectMongoDBTableContentWithCountTemplate(ctx, op, condition...)
}

func UpdatePayable(ctx *gin.Context) {

}

func DeletePayable(ctx *gin.Context) {

}
