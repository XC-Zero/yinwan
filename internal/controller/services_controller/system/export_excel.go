package system

import (
	"github.com/XC-Zero/yinwan/internal/controller/services_controller/common"
	"github.com/XC-Zero/yinwan/pkg/client"
	"github.com/XC-Zero/yinwan/pkg/model/mysql_model"
	"github.com/gin-gonic/gin"
)

func GetExportExcel(ctx *gin.Context) {
	conditions := []common.MysqlCondition{
		{},
	}
	common.SelectMysqlTableContentWithCountTemplate(ctx, common.SelectMysqlTemplateOptions{
		DB:         client.MysqlClient.WithContext(ctx),
		TableModel: mysql_model.Excel{},
	}, conditions...)
	return
}

func DownloadExportExcel(ctx *gin.Context) {

}
