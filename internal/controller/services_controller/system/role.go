package system

import (
	"encoding/json"
	"github.com/XC-Zero/yinwan/internal/controller/services_controller/common"
	"github.com/XC-Zero/yinwan/pkg/client"
	_const "github.com/XC-Zero/yinwan/pkg/const"
	"github.com/XC-Zero/yinwan/pkg/model"
	"github.com/XC-Zero/yinwan/pkg/utils/errs"
	"github.com/XC-Zero/yinwan/pkg/utils/mysql"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// CreateRole 创建角色
func CreateRole(ctx *gin.Context) {
	var postData map[string]interface{}
	err := ctx.ShouldBind(&postData)
	r, ok := postData["role"]
	rc, rcOK := postData["role_capabilities"]

	if !ok || !rcOK || err != nil {
		common.RequestParamErrorTemplate(ctx, common.REQUEST_PARM_ERROR)
		return
	}
	var role model.Role
	var roleCapList []model.RoleCapabilities

	rb, _ := json.Marshal(r)
	rcb, _ := json.Marshal(rc)
	err = json.Unmarshal(rb, &role)
	if err != nil {
		common.RequestParamErrorTemplate(ctx, common.REQUEST_PARM_ERROR)
		return
	}
	err = json.Unmarshal(rcb, &roleCapList)
	if err != nil {
		common.RequestParamErrorTemplate(ctx, common.REQUEST_PARM_ERROR)
		return
	}

	err = client.MysqlClient.Transaction(func(tx *gorm.DB) error {
		err := client.MysqlClient.Model(&model.Role{}).Create(&role).Error
		if err != nil {
			return err
		}
		for i := range roleCapList {
			roleCapList[i].RoleID = *role.RecID
		}
		err = client.MysqlClient.Model(&model.RoleCapabilities{}).CreateInBatches(roleCapList, mysql.CalcMysqlBatchSize(roleCapList[0])).Error
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		common.InternalDataBaseErrorTemplate(ctx, common.DATABASE_INSERT_ERROR, model.Role{})
		return
	}

	ctx.JSON(_const.OK, errs.CreateSuccessMsg("创建角色成功！"))
	return
}

// SelectRole 查询角色
func SelectRole(ctx *gin.Context) {
	roleID := ctx.PostForm("role_id")
	conditions := []common.MysqlCondition{
		{
			Symbol:      mysql.EQUAL,
			ColumnName:  "rec_id",
			ColumnValue: roleID,
		},
		{
			Symbol:      mysql.LIKE,
			ColumnName:  "role_name",
			ColumnValue: ctx.PostForm("role_name"),
		},
		{
			Symbol:      mysql.NULL,
			ColumnName:  "deleted_at",
			ColumnValue: " ",
		},
	}
	op := common.SelectMysqlTemplateOptions{
		DB:         client.MysqlClient,
		TableModel: model.Role{},
	}
	common.SelectMysqlTableContentWithCountTemplate(ctx, op, conditions...)

	// 查询角色内置权限  todo 测试！
	op2 := common.SelectMysqlTemplateOptions{
		DB:         client.MysqlClient,
		TableModel: model.RoleCapabilities{},
	}

	common.SelectMysqlTableContentWithCountTemplate(ctx, op2,
		common.MysqlCondition{
			Symbol:      mysql.EQUAL,
			ColumnName:  "rec_id",
			ColumnValue: roleID,
		},
		common.MysqlCondition{
			Symbol:      mysql.NULL,
			ColumnName:  "deleted_at",
			ColumnValue: " ",
		})

	return
}

func UpdateRole(ctx *gin.Context) {

}

func DeleteRole(ctx *gin.Context) {
	recID := ctx.PostForm("role_id")
	if recID == "" {
		return
	}
	err := client.MysqlClient.Transaction(func(tx *gorm.DB) error {
		err := tx.Delete(&model.Role{}, recID).Error
		if err != nil {
			return err
		}
		err = tx.Delete(&model.RoleCapabilities{}, "role_id = ?", recID).Error
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		ctx.JSON(_const.INTERNAL_ERROR, errs.CreateWebErrorMsg("删除角色失败！"))
		return
	}

	ctx.JSON(_const.OK, errs.CreateSuccessMsg("删除成功！"))
	return
}
