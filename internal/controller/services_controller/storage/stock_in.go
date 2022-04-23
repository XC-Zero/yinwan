package storage

import (
	"github.com/XC-Zero/yinwan/internal/controller/services_controller/common"
	"github.com/XC-Zero/yinwan/pkg/client"
	"github.com/XC-Zero/yinwan/pkg/model/mongo_model"
	"github.com/gin-gonic/gin"
)

func CreateStockIn(ctx *gin.Context) {
	temp := mongo_model.StockInRecord{}
	//bk := client.HarvestClientFromGinContext(ctx)
	//if bk == nil {
	//	common.RequestParamErrorTemplate(ctx, common.BOOK_NAME_LACK_ERROR)
	//	return
	//}
	err := ctx.ShouldBind(&temp)
	if err != nil {
		common.RequestParamErrorTemplate(ctx, common.REQUEST_PARM_ERROR)
		return
	}

}

func SelectStockIn(ctx *gin.Context) {
	//bk := client.HarvestClientFromGinContext(ctx)
	//if bk == nil {
	//	common.RequestParamErrorTemplate(ctx, common.BOOK_NAME_LACK_ERROR)
	//	return
	//}
	conditions := []common.MongoCondition{
		{
			Symbol:      "",
			ColumnName:  "",
			ColumnValue: nil,
		},
	}
	options := common.SelectMongoDBTemplateOptions{
		DB:         client.MongoDBClient,
		TableModel: mongo_model.StockInRecord{},
	}
	common.SelectMongoDBTableContentWithCountTemplate(ctx, options, conditions...)
	return

}
func UpdateStockIn(ctx *gin.Context) {
	//bk := client.HarvestClientFromGinContext(ctx)
	//if bk == nil {
	//	common.RequestParamErrorTemplate(ctx, common.BOOK_NAME_LACK_ERROR)
	//	return
	//}
}
func DeleteStockIn(ctx *gin.Context) {
	//bk := client.HarvestClientFromGinContext(ctx)
	//if bk == nil {
	//	common.RequestParamErrorTemplate(ctx, common.BOOK_NAME_LACK_ERROR)
	//	return
	//}
	//res, err := client.MongoDBClient.Collection(model.StockInRecord{}.TableName()).DeleteOne()
	//if err != nil {
	//	return
	//}
}
