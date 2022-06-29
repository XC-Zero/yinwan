package transaction

import (
	"context"
	"github.com/XC-Zero/yinwan/internal/controller/services_controller/common"
	"github.com/XC-Zero/yinwan/pkg/model/mongo_model"
	"github.com/XC-Zero/yinwan/pkg/utils/logger"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/pkg/errors"
	"time"
)

// 是否自动创建应收?
type autoCreate struct {
	AutoCreate bool `json:"auto_create" form:"auto_create"`
}

func CreateTransaction(ctx *gin.Context) {
	bk, n := common.HarvestClientFromGinContext(ctx)
	var transaction mongo_model.Transaction
	err := ctx.ShouldBindBodyWith(&transaction, binding.JSON)
	if err != nil {
		logger.Error(errors.WithStack(err), "")
		common.RequestParamErrorTemplate(ctx, common.REQUEST_PARM_ERROR)
		return
	}
	var autoCreate autoCreate
	ctx.ShouldBindBodyWith(&autoCreate, binding.JSON)
	recID := int(time.Now().Unix())
	transaction.RecID = &recID
	transaction.BookName = n
	transaction.BookNameID = bk.StorageName
	common.CreateOneMongoDBRecordTemplate(ctx, common.CreateMongoDBTemplateOptions{
		DB:         bk.MongoDBClient,
		Context:    context.WithValue(context.WithValue(context.Background(), "book_name", n), "auto_create", autoCreate.AutoCreate),
		TableModel: transaction,
	})
	return
}

func UpdateTransaction(ctx *gin.Context) {

}

func SelectTransaction(ctx *gin.Context) {

}

func DeleteTransaction(ctx *gin.Context) {

}
