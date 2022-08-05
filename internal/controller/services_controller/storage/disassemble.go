package storage

import (
	"github.com/XC-Zero/yinwan/internal/controller/services_controller/common"
	"github.com/XC-Zero/yinwan/pkg/model/mongo_model"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

// CreateDisassemble TODO 组装拆卸
func CreateDisassemble(ctx *gin.Context) {
	bk := common.HarvestClientFromGinContext(ctx)
	var assemble mongo_model.Assemble
	err := ctx.ShouldBindBodyWith(&assemble, binding.JSON)
	if err != nil {
		return
	}
	op := common.CreateMongoDBTemplateOptions{
		DB:         bk.MongoDBClient,
		Context:    ctx,
		TableModel: assemble,
	}
	common.CreateOneMongoDBRecordTemplate(ctx, op)
	return
}

func UpdateDisassemble(ctx *gin.Context) {

}

func SelectDisassemble(ctx *gin.Context) {

}

func DeleteDisassemble(ctx *gin.Context) {

}
