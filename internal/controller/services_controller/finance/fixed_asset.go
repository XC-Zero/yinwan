package finance

import (
	"github.com/XC-Zero/yinwan/internal/controller/services_controller/common"
	"github.com/XC-Zero/yinwan/pkg/client"
	_const "github.com/XC-Zero/yinwan/pkg/const"
	"github.com/XC-Zero/yinwan/pkg/model"
	"github.com/XC-Zero/yinwan/pkg/utils/errs"
	"github.com/gin-gonic/gin"
)

func CreateFixedAsset(ctx *gin.Context) {
	var fixedAsset model.FixedAsset
	err := ctx.ShouldBind(&fixedAsset)
	if err != nil {
		common.RequestParamErrorTemplate(ctx, common.REQUEST_PARM_ERROR)
		return
	}

	bk := client.HarvestClientFromGinContext(ctx)
	if bk == nil {
		return
	}
	err = bk.MysqlClient.Create(&fixedAsset).Error
	if err != nil {
		common.InternalDataBaseErrorTemplate(ctx, common.DATABASE_INSERT_ERROR, fixedAsset)
		return
	}
	ctx.JSON(_const.OK, errs.CreateSuccessMsg("创建固定资产成功！"))
	return
}

func SelectFixedAsset(ctx *gin.Context) {
	bk := client.HarvestClientFromGinContext(ctx)
	if bk == nil {
		return
	}
	conditions := []common.MysqlCondition{{}}
	op := common.SelectMysqlTemplateOptions{
		DB:         bk.MysqlClient,
		TableModel: model.FixedAsset{},
	}
	common.SelectMysqlTableContentWithCountTemplate(ctx, op, conditions...)
	return
}

func UpdateFixedAsset(ctx *gin.Context) {

}

func DeleteFixedAsset(ctx *gin.Context) {

}
