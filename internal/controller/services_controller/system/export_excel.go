package system

import (
	"github.com/XC-Zero/yinwan/internal/controller/services_controller/common"
	"github.com/XC-Zero/yinwan/pkg/client"
	"github.com/XC-Zero/yinwan/pkg/model/mysql_model"
	"github.com/XC-Zero/yinwan/pkg/utils/mysql"
	"github.com/gin-gonic/gin"
)

func SelectExportExcel(ctx *gin.Context) {

	conditions := []common.MysqlCondition{
		{
			mysql.EQUAL,
			"rec_id",
			ctx.PostForm("excel_id"),
		},
		{
			mysql.LIKE,
			"excel_name",
			ctx.PostForm("excel_name"),
		},
		{
			mysql.EQUAL,
			"creator_id",
			ctx.PostForm("creator_id"),
		},
		{
			mysql.EQUAL,
			"excel_status",
			ctx.PostForm("excel_status"),
		},
	}
	common.SelectMysqlTableContentWithCountTemplate(ctx, common.SelectMysqlTemplateOptions{
		DB:         client.MysqlClient,
		TableModel: mysql_model.Excel{},
	}, conditions...)
}

// TODO 创建Excel
