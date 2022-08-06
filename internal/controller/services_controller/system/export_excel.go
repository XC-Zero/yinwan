package system

import (
	"github.com/XC-Zero/yinwan/internal/controller/services_controller/common"
	"github.com/XC-Zero/yinwan/pkg/client"
	"github.com/XC-Zero/yinwan/pkg/model/mysql_model"
	"github.com/XC-Zero/yinwan/pkg/utils/mysql"
	"github.com/gin-gonic/gin"
)

func GetExportExcel(ctx *gin.Context) {
	conditions := []common.MysqlCondition{
		{
			Symbol:      mysql.EQUAL,
			ColumnName:  "creator_id",
			ColumnValue: ctx.PostForm("creator_id"),
		},
		{
			Symbol:      mysql.NULL,
			ColumnName:  "deleted_at",
			ColumnValue: " ",
		},
	}
	common.SelectMysqlTableContentWithCountTemplate(ctx, common.SelectMysqlTemplateOptions{
		DB:         client.MysqlClient.WithContext(ctx),
		TableModel: mysql_model.Excel{},
	}, conditions...)
	return
}

func DownloadExportExcel(ctx *gin.Context) {

}
