package storage

import (
	"github.com/XC-Zero/yinwan/internal/controller/services_controller/common"
	"github.com/XC-Zero/yinwan/pkg/model/mongo_model"
	my_mongo "github.com/XC-Zero/yinwan/pkg/utils/mongo"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"go.mongodb.org/mongo-driver/bson/bsontype"
	"strconv"
	"time"
)

// CreateAssemble TODO 组装拆卸
func CreateAssemble(ctx *gin.Context) {
	bk := common.HarvestClientFromGinContext(ctx)
	var assemble mongo_model.Assemble
	err := ctx.ShouldBindBodyWith(&assemble, binding.JSON)
	if err != nil {
		return
	}

	recID := int(time.Now().Unix())
	assemble.RecID = &recID
	assemble.BookName = bk.BookName
	assemble.BookNameID = bk.StorageName
	assemble.CreatedAt = strconv.Itoa(recID)

	op := common.CreateMongoDBTemplateOptions{
		DB:         bk.MongoDBClient,
		Context:    ctx,
		TableModel: assemble,
	}

	common.CreateOneMongoDBRecordTemplate(ctx, op)
	return
}

func UpdateAssemble(ctx *gin.Context) {
	bk := common.HarvestClientFromGinContext(ctx)
	if bk == nil {
		return
	}
	var assemble mongo_model.Assemble
	err := ctx.ShouldBindBodyWith(&assemble, binding.JSON)
	if err != nil || assemble.RecID == nil {
		common.RequestParamErrorTemplate(ctx, common.REQUEST_PARM_ERROR)
		return
	}
	now := strconv.Itoa(int(time.Now().Unix()))
	assemble.UpdatedAt = &now
	common.UpdateOneMongoDBRecordByIDTemplate(ctx, common.MongoDBTemplateOptions{
		DB:         bk.MongoDBClient,
		Context:    ctx,
		TableModel: assemble,
		RecID:      *assemble.RecID,
	})
}

func SelectAssemble(ctx *gin.Context) {
	bk := common.HarvestClientFromGinContext(ctx)
	if bk == nil {
		return
	}
	conditions := []common.MongoCondition{
		{
			my_mongo.EQUAL,
			"rec_id",
			ctx.PostForm("assemble_id"),
		},
		{
			my_mongo.LIKE,
			"assemble_name",
			ctx.PostForm("assemble_name"),
		},
		{
			my_mongo.EQUAL,
			"deleted_at",
			bsontype.Null,
		},
	}
	common.SelectMongoDBTableContentWithCountTemplate(ctx, common.SelectMongoDBTemplateOptions{
		DB:         bk.MongoDBClient,
		TableModel: &mongo_model.Assemble{},
	}, conditions...)
	return
}

func DeleteAssemble(ctx *gin.Context) {
	bk := common.HarvestClientFromGinContext(ctx)
	if bk == nil {
		return
	}
	id, err := strconv.Atoi(ctx.PostForm("assemble_id"))
	if err != nil || id == 0 {
		common.RequestParamErrorTemplate(ctx, common.REQUEST_PARM_ERROR)
		return
	}
	common.DeleteOneMongoDBRecordByIDTemplate(ctx, common.MongoDBTemplateOptions{
		DB:         bk.MongoDBClient,
		Context:    ctx,
		TableModel: &mongo_model.Assemble{},
		RecID:      id,
	})
}
