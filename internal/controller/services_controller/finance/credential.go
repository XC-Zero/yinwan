package finance

import (
	"context"
	"github.com/XC-Zero/yinwan/internal/controller/services_controller/common"
	_const "github.com/XC-Zero/yinwan/pkg/const"
	_interface "github.com/XC-Zero/yinwan/pkg/interface"
	"github.com/XC-Zero/yinwan/pkg/model/mongo_model"
	"github.com/XC-Zero/yinwan/pkg/model/mysql_model"
	my_mongo "github.com/XC-Zero/yinwan/pkg/utils/mongo"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/bsontype"
	"gorm.io/gorm/callbacks"
	"strconv"
	"time"
)

// CreateCredential 创建凭证
func CreateCredential(ctx *gin.Context) {
	bk := common.HarvestClientFromGinContext(ctx)
	if bk == nil {
		return
	}
	var credential mongo_model.Credential

	err := ctx.ShouldBindBodyWith(&credential, binding.JSON)
	if err != nil {
		common.RequestParamErrorTemplate(ctx, common.REQUEST_PARM_ERROR)
		return
	}
	recID := int(time.Now().Unix())
	credential.RecID = &recID
	credential.CreatedAt = strconv.Itoa(int(time.Now().Unix()))
	credential.BookName = bk.BookName
	credential.BookNameID = bk.StorageName
	common.CreateOneMongoDBRecordTemplate(ctx, common.CreateMongoDBTemplateOptions{
		DB:         bk.MongoDBClient,
		TableModel: credential,
		Context:    ctx,
	})
	return
}

// SelectCredential 查找凭证
func SelectCredential(ctx *gin.Context) {
	bk := common.HarvestClientFromGinContext(ctx)
	if bk == nil {
		return
	}
	conditions := []common.MongoCondition{
		{
			my_mongo.EQUAL,
			"rec_id",
			ctx.PostForm("credential_id"),
		},
		{
			my_mongo.LIKE,
			"credential_name",
			ctx.PostForm("credential_name"),
		},
		{
			my_mongo.EQUAL,
			"deleted_at",
			bsontype.Null,
		},
	}

	op := common.SelectMongoDBTemplateOptions{
		DB:         bk.MongoDBClient,
		TableModel: mongo_model.Credential{},
		ResHookFunc: func(data []interface{}) []interface{} {
			var res []interface{}
			for _, datum := range data {
				if c, ok := datum.(mongo_model.Credential); ok {
					if c.CredentialStatus != nil {
						temp := c.CredentialStatus.Display()
						c.CredentialStatus = (*_const.CredentialStatus)(&temp)
					}
					res = append(res, c)
				}
			}
			return res
		},
	}
	common.SelectMongoDBTableContentWithCountTemplate(ctx, op, conditions...)
	return
}

// UpdateCredential 更新凭证
func UpdateCredential(ctx *gin.Context) {
	bk := common.HarvestClientFromGinContext(ctx)
	if bk == nil {
		return
	}
	var credential mongo_model.Credential
	err := ctx.ShouldBindBodyWith(&credential, binding.JSON)
	if err != nil || credential.RecID == nil {
		common.RequestParamErrorTemplate(ctx, common.REQUEST_PARM_ERROR)
		return
	}
	common.UpdateOneMongoDBRecordByIDTemplate(ctx, common.MongoDBTemplateOptions{
		DB:         bk.MongoDBClient,
		RecID:      *credential.RecID,
		TableModel: credential,
		Context:    ctx,
	})
	return
}

// DeleteCredential 删除凭证
func DeleteCredential(ctx *gin.Context) {
	bk := common.HarvestClientFromGinContext(ctx)
	if bk == nil {
		return
	}
	var credential mongo_model.Credential
	err := ctx.ShouldBindBodyWith(&credential, binding.JSON)
	if err != nil || credential.RecID == nil {
		common.RequestParamErrorTemplate(ctx, common.REQUEST_PARM_ERROR)
		return
	}

	common.DeleteOneMongoDBRecordByIDTemplate(ctx, common.MongoDBTemplateOptions{
		DB:         bk.MongoDBClient,
		RecID:      *credential.RecID,
		TableModel: credential,
		Context:    ctx,
	})
	return
}

// GenerateCredential 生成凭证
func GenerateCredential(ctx *gin.Context) {
	bk := common.HarvestClientFromGinContext(ctx)
	if bk == nil {
		return
	}
	recID := ctx.PostForm("rec_id")
	srcType := ctx.PostForm("src_type")
	if recID == "" || srcType == "" {
		common.RequestParamErrorTemplate(ctx, common.REQUEST_PARM_ERROR)
		return
	}
	filter := bson.D{}
	filter = append(filter,
		my_mongo.TransMysqlOperatorSymbol(my_mongo.EQUAL, "rec_id", recID),
		my_mongo.TransMysqlOperatorSymbol(my_mongo.EQUAL, "deleted_at", bsontype.Null),
	)
	c, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()
	var temp _interface.ChineseTabler
	switch srcType {
	case "stock_in":
		temp = mongo_model.StockInRecord{}
	case "stock_out":
		temp = mongo_model.StockOutRecord{}
	case "purchase":
		temp = mongo_model.Purchase{}
	case "payable":
		temp = mysql_model.Payable{}
	case "receivable":
		temp = mysql_model.Receivable{}
	case "transaction":
		temp = mongo_model.Transaction{}
	case "assemble":
		temp = mongo_model.Assemble{}
	default:
		common.RequestParamErrorTemplate(ctx, "这个暂时还不能转为凭证哦~")
		return
	}
	_, ok := temp.(callbacks.AfterCreateInterface)
	_, ok2 := temp.(callbacks.AfterUpdateInterface)
	if ok || ok2 {
		err := bk.MysqlClient.WithContext(ctx).Where("rec_id =?", recID).Find(&temp).Error
		if err != nil {
			common.InternalDataBaseErrorTemplate(ctx, "出了点问题!", temp)
			return
		}
	} else {
		err := bk.MongoDBClient.Collection(temp.TableName()).Find(c, filter).Limit(1).All(&temp)
		if err != nil && err != my_mongo.CancelError {
			common.InternalDataBaseErrorTemplate(ctx, "出了点问题!", temp)
			return
		}
	}

	cTemp, ok := temp.(_interface.Credential)
	if !ok {
		common.RequestParamErrorTemplate(ctx, "这个暂时还不能转为凭证哦~")
	}
	cTemp.ToCredential(ctx)
	return
}
