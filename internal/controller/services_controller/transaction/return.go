package transaction

import (
	"context"
	"github.com/XC-Zero/yinwan/internal/controller/services_controller/common"
	"github.com/XC-Zero/yinwan/pkg/model/mongo_model"
	my_mongo "github.com/XC-Zero/yinwan/pkg/utils/mongo"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"strconv"
	"time"
)

func CreateReturn(ctx *gin.Context) {
	bk, n := common.HarvestClientFromGinContext(ctx)
	if bk == nil {
		return
	}

	var ret mongo_model.Return
	err := ctx.ShouldBindBodyWith(&ret, binding.JSON)
	if err != nil {
		common.RequestParamErrorTemplate(ctx, common.REQUEST_PARM_ERROR)
		return
	}
	recID := int(time.Now().Unix())
	ret.RecID = &recID
	ret.BookName = n
	ret.BookNameID = bk.StorageName
	common.CreateOneMongoDBRecordTemplate(ctx, common.CreateMongoDBTemplateOptions{
		DB:         bk.MongoDBClient,
		Context:    context.WithValue(context.Background(), "book_name", n),
		TableModel: ret,
	})
	return
}

func UpdateReturn(ctx *gin.Context) {
	bk, n := common.HarvestClientFromGinContext(ctx)
	if bk == nil {
		return
	}
	var ret mongo_model.Return
	err := ctx.ShouldBindBodyWith(&ret, binding.JSON)
	if err != nil || ret.RecID == nil {
		common.RequestParamErrorTemplate(ctx, common.REQUEST_PARM_ERROR)
		return
	}
	common.UpdateOneMongoDBRecordByIDTemplate(ctx, common.MongoDBTemplateOptions{
		DB:         bk.MongoDBClient,
		Context:    context.WithValue(context.Background(), "book_name", n),
		TableModel: ret,
		RecID:      *ret.RecID,
	})
	return
}

func SelectReturn(ctx *gin.Context) {
	bk, _ := common.HarvestClientFromGinContext(ctx)
	if bk == nil {
		return
	}
	conditions := []common.MongoCondition{
		{
			my_mongo.LESS_THAN_EQUAL,
			"deleted_at",
			nil,
		},
		{
			Symbol:      my_mongo.EQUAL,
			ColumnName:  "rec_id",
			ColumnValue: ctx.PostForm("return_id"),
		},
	}

	common.SelectMongoDBTableContentWithCountTemplate(ctx, common.SelectMongoDBTemplateOptions{
		DB:         bk.MongoDBClient,
		TableModel: mongo_model.Return{},
	}, conditions...)
	return
}

func DeleteReturn(ctx *gin.Context) {
	bk, n := common.HarvestClientFromGinContext(ctx)
	if bk == nil {
		return
	}
	var ret mongo_model.Return
	recID, err := strconv.Atoi(ctx.PostForm("return_id"))
	if err != nil || recID == 0 {
		common.RequestParamErrorTemplate(ctx, common.REQUEST_PARM_ERROR)
		return
	}
	ret.RecID = &recID
	common.DeleteOneMongoDBRecordByIDTemplate(ctx, common.MongoDBTemplateOptions{
		DB:         bk.MongoDBClient,
		Context:    context.WithValue(context.Background(), "book_name", n),
		TableModel: ret,
		RecID:      recID,
	})
}
