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

func CreateAssembleTemplate(ctx *gin.Context) {
	bk := common.HarvestClientFromGinContext(ctx)
	if bk == nil {
		return
	}
	var assemble mongo_model.AssembleTemplate
	err := ctx.ShouldBindBodyWith(&assemble, binding.JSON)
	if err != nil {
		common.RequestParamErrorTemplate(ctx, common.REQUEST_PARM_ERROR)
		return
	}

	recID := int(time.Now().Unix())
	assemble.RecID = &recID
	assemble.BookName = bk.BookName
	assemble.BookNameID = bk.StorageName
	assemble.CreatedAt = strconv.Itoa(recID)
	common.CreateOneMongoDBRecordTemplate(ctx, common.CreateMongoDBTemplateOptions{
		DB:         bk.MongoDBClient,
		Context:    ctx,
		TableModel: &assemble,
		NotSyncES:  true,
	})
	return
}

func UpdateAssembleTemplate(ctx *gin.Context) {
	bk := common.HarvestClientFromGinContext(ctx)
	if bk == nil {
		return
	}
	var assemble mongo_model.AssembleTemplate
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
		NotSyncES:  true,
	})
}

func SelectAssembleTemplate(ctx *gin.Context) {
	bk := common.HarvestClientFromGinContext(ctx)
	if bk == nil {
		return
	}
	conditions := []common.MongoCondition{
		{
			my_mongo.EQUAL,
			"rec_id",
			ctx.PostForm("assemble_template_id"),
		},
		{
			my_mongo.LIKE,
			"assemble_template_name",
			ctx.PostForm("assemble_template_name"),
		},
		{
			my_mongo.EQUAL,
			"deleted_at",
			bsontype.Null,
		},
	}
	common.SelectMongoDBTableContentWithCountTemplate(ctx, common.SelectMongoDBTemplateOptions{
		DB:         bk.MongoDBClient,
		TableModel: &mongo_model.AssembleTemplate{},
	}, conditions...)
	return
}

func DeleteAssembleTemplate(ctx *gin.Context) {
	bk := common.HarvestClientFromGinContext(ctx)
	if bk == nil {
		return
	}

	id, err := strconv.Atoi(ctx.PostForm("assemble_template_id"))
	if err != nil || id == 0 {
		common.RequestParamErrorTemplate(ctx, common.REQUEST_PARM_ERROR)
		return
	}

	common.DeleteOneMongoDBRecordByIDTemplate(ctx, common.MongoDBTemplateOptions{
		DB:         bk.MongoDBClient,
		Context:    ctx,
		TableModel: &mongo_model.AssembleTemplate{},
		RecID:      id,
		NotSyncES:  true,
	})
}
