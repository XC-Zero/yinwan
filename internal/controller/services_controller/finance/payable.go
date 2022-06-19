package finance

import (
	"context"
	"github.com/XC-Zero/yinwan/internal/controller/services_controller/common"
	"github.com/XC-Zero/yinwan/pkg/client"
	"github.com/XC-Zero/yinwan/pkg/model/mysql_model"
	"github.com/XC-Zero/yinwan/pkg/utils/logger"
	my_mongo "github.com/XC-Zero/yinwan/pkg/utils/mongo"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/pkg/errors"
	"strconv"
)

// todo mysql 的用啥MongoDB啊!!!

func CreatePayable(ctx *gin.Context) {
	bk, n := common.HarvestClientFromGinContext(ctx)
	if bk == nil {
		return
	}
	temp := mysql_model.Payable{}

	err := ctx.ShouldBindBodyWith(&temp, binding.JSON)
	if err != nil {
		logger.Error(errors.WithStack(err), "绑定模型失败!")
		common.RequestParamErrorTemplate(ctx, common.REQUEST_PARM_ERROR)
		return
	}
	common.CreateOneMongoDBRecordTemplate(ctx, common.CreateMongoDBTemplateOptions{
		DB:         bk.MongoDBClient,
		Context:    context.WithValue(context.Background(), "book_name", n),
		TableModel: temp,
		PreFunc:    nil,
	})
	return
}

func SelectPayable(ctx *gin.Context) {
	bk, _ := common.HarvestClientFromGinContext(ctx)
	if bk == nil {
		return
	}
	condition := []common.MongoCondition{
		{
			Symbol:      my_mongo.NOT_EQUAL,
			ColumnName:  "deleted_at",
			ColumnValue: nil,
		},
	}
	op := common.SelectMongoDBTemplateOptions{
		DB:         client.MongoDBClient,
		TableModel: mysql_model.Payable{},
	}
	common.SelectMongoDBTableContentWithCountTemplate(ctx, op, condition...)
}

func UpdatePayable(ctx *gin.Context) {
	bk, n := common.HarvestClientFromGinContext(ctx)
	if bk == nil {
		return
	}
	temp := mysql_model.Payable{}

	err := ctx.ShouldBindBodyWith(&temp, binding.JSON)
	if err != nil {
		logger.Error(errors.WithStack(err), "绑定模型失败!")
		common.RequestParamErrorTemplate(ctx, common.REQUEST_PARM_ERROR)
		return
	}
	common.UpdateOneMongoDBRecordByIDTemplate(ctx, common.MongoDBTemplateOptions{
		DB:         bk.MongoDBClient,
		Context:    context.WithValue(context.Background(), "book_name", n),
		TableModel: temp,
		PreFunc:    nil,
	})
	return
}

func DeletePayable(ctx *gin.Context) {
	bk, n := common.HarvestClientFromGinContext(ctx)
	if bk == nil {
		return
	}
	var stockOutRecord mysql_model.Payable
	recID, err := strconv.Atoi(ctx.PostForm("payable_id"))
	if err != nil {
		common.RequestParamErrorTemplate(ctx, common.REQUEST_PARM_ERROR)
		return
	}
	common.DeleteOneMongoDBRecordByIDTemplate(ctx, common.MongoDBTemplateOptions{
		DB:         bk.MongoDBClient,
		Context:    context.WithValue(context.Background(), "book_name", n),
		RecID:      recID,
		TableModel: stockOutRecord,
		PreFunc:    nil,
	})
	return
}
