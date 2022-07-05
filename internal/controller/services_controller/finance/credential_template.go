package finance

import (
	"github.com/XC-Zero/yinwan/internal/controller/services_controller/common"
	"github.com/XC-Zero/yinwan/pkg/client"
	"github.com/XC-Zero/yinwan/pkg/model/mongo_model"
	"github.com/XC-Zero/yinwan/pkg/utils/mysql"
	"github.com/gin-gonic/gin"
)

// todo !!! 凭证模板

func CreateCredentialTemplate(ctx *gin.Context) {

}
func SelectCredentialTemplate(ctx *gin.Context) {

	conditions := []common.MysqlCondition{
		{
			Symbol:      mysql.EQUAL,
			ColumnName:  "rec_id",
			ColumnValue: ctx.PostForm("credential_template_id"),
		},
		{
			Symbol:      mysql.NULL,
			ColumnName:  "deleted_at",
			ColumnValue: " ",
		},
	}

	common.SelectMysqlTableContentWithCountTemplate(ctx, common.SelectMysqlTemplateOptions{
		DB:         client.MysqlClient,
		TableModel: mongo_model.CredentialTemplate{},
	}, conditions...)

}
func UpdateCredentialTemplate(ctx *gin.Context) {

}
func DeleteCredentialTemplate(ctx *gin.Context) {

}
