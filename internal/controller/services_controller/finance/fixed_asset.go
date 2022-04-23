package finance

import (
	"github.com/XC-Zero/yinwan/internal/controller/services_controller/common"
	"github.com/XC-Zero/yinwan/pkg/client"
	_const "github.com/XC-Zero/yinwan/pkg/const"
	"github.com/XC-Zero/yinwan/pkg/model/mysql_model"
	"github.com/XC-Zero/yinwan/pkg/utils/errs"
	"github.com/gin-gonic/gin"
)

func CreateFixedAsset(ctx *gin.Context) {
	var fixedAsset mysql_model.FixedAsset
	err := ctx.ShouldBind(&fixedAsset)
	if err != nil {
		common.RequestParamErrorTemplate(ctx, common.REQUEST_PARM_ERROR)
		return
	}

	//bk := client.HarvestClientFromGinContext(ctx)
	//if bk == nil {
	//	common.RequestParamErrorTemplate(ctx, common.BOOK_NAME_LACK_ERROR)
	//	return
	//}
	ctx.JSON(_const.OK, errs.CreateSuccessMsg("创建固定资产成功！"))
	return
}

func SelectFixedAsset(ctx *gin.Context) {
	//bk := client.HarvestClientFromGinContext(ctx)
	//if bk == nil {
	//	common.RequestParamErrorTemplate(ctx, common.BOOK_NAME_LACK_ERROR)
	//	return
	//}
	conditions := []common.MysqlCondition{{}}
	op := common.SelectMysqlTemplateOptions{
		DB:         client.MysqlClient,
		TableModel: mysql_model.FixedAsset{},
	}
	common.SelectMysqlTableContentWithCountTemplate(ctx, op, conditions...)
	return
}

func UpdateFixedAsset(ctx *gin.Context) {

}

func DeleteFixedAsset(ctx *gin.Context) {

}
