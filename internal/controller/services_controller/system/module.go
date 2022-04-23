package system

import (
	"github.com/XC-Zero/yinwan/internal/controller/services_controller/common"
	"github.com/XC-Zero/yinwan/pkg/client"
	"github.com/XC-Zero/yinwan/pkg/model/mysql_model"
	"github.com/XC-Zero/yinwan/pkg/utils/mysql"
	"github.com/gin-gonic/gin"
)

func SelectModule(ctx *gin.Context) {
	condition := []common.MysqlCondition{
		{
			Symbol:      mysql.NULL,
			ColumnName:  "deleted_at",
			ColumnValue: " ",
		},
	}
	op := common.SelectMysqlTemplateOptions{
		DB:         client.MysqlClient,
		TableModel: mysql_model.Module{},
	}
	common.SelectMysqlTableContentWithCountTemplate(ctx, op, condition...)
	return
}
