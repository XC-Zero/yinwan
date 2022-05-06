package finance

import (
	"github.com/XC-Zero/yinwan/internal/controller/services_controller/common"
	_const "github.com/XC-Zero/yinwan/pkg/const"
	"github.com/XC-Zero/yinwan/pkg/model/mysql_model"
	"github.com/XC-Zero/yinwan/pkg/utils/errs"
	"github.com/gin-gonic/gin"
	"golang.org/x/net/context"
)

// CreateFixedAsset 创建固定资产
func CreateFixedAsset(ctx *gin.Context) {
	bk, n := common.HarvestClientFromGinContext(ctx)
	if bk == nil {
		return
	}
	var fixedAsset mysql_model.FixedAsset
	err := ctx.ShouldBind(&fixedAsset)
	if err != nil {
		common.RequestParamErrorTemplate(ctx, common.REQUEST_PARM_ERROR)
		return
	}

	err = bk.MysqlClient.WithContext(context.WithValue(context.Background(), "book_name", n)).Create(&fixedAsset).Error
	if err != nil {
		common.InternalDataBaseErrorTemplate(ctx, common.DATABASE_INSERT_ERROR, fixedAsset)
		return
	}
	ctx.JSON(_const.OK, errs.CreateSuccessMsg("创建固定资产成功！"))
	return
}

func SelectFixedAsset(ctx *gin.Context) {
	bk, n := common.HarvestClientFromGinContext(ctx)
	if bk == nil {
		return
	}
	conditions := []common.MysqlCondition{{}}
	op := common.SelectMysqlTemplateOptions{
		DB:         bk.MysqlClient.WithContext(context.WithValue(context.Background(), "book_name", n)),
		TableModel: mysql_model.FixedAsset{},
	}
	common.SelectMysqlTableContentWithCountTemplate(ctx, op, conditions...)
	return
}

// UpdateFixedAsset TODO !
func UpdateFixedAsset(ctx *gin.Context) {
	bk, n := common.HarvestClientFromGinContext(ctx)
	if bk == nil {
		return
	}
	var fixedAsset mysql_model.FixedAsset
	err := ctx.ShouldBind(&fixedAsset)
	if err != nil {
		common.RequestParamErrorTemplate(ctx, common.REQUEST_PARM_ERROR)
		return
	}
	//bk.MysqlClient.WithContext(context.WithValue(context.Background(), "book_name", n)).Updates()
}

func DeleteFixedAsset(ctx *gin.Context) {
	id := ctx.PostForm("fixed_asset_id")
	asset := mysql_model.FixedAsset{}
	if id == "" {
		common.RequestParamErrorTemplate(ctx, common.REQUEST_PARM_ERROR)
		return
	}
	bk, n := common.HarvestClientFromGinContext(ctx)
	if bk == nil {
		return
	}

	err := bk.MysqlClient.WithContext(context.WithValue(context.Background(), "book_name", n)).
		Delete(asset, id).
		Error
	if err != nil {
		common.InternalDataBaseErrorTemplate(ctx, common.DATABASE_DELETE_ERROR, asset)
		return
	}
	return
}
