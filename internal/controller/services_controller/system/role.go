package system

import (
	"encoding/json"
	"github.com/XC-Zero/yinwan/internal/controller/services_controller/common"
	"github.com/XC-Zero/yinwan/pkg/client"
	_const "github.com/XC-Zero/yinwan/pkg/const"
	"github.com/XC-Zero/yinwan/pkg/model/mysql_model"
	"github.com/XC-Zero/yinwan/pkg/utils/errs"
	"github.com/XC-Zero/yinwan/pkg/utils/mysql"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"gorm.io/gorm"
)

// CreateRole 创建角色
func CreateRole(ctx *gin.Context) {
	var postData map[string]interface{}
	err := ctx.ShouldBindBodyWith(&postData, binding.JSON)
	r, ok := postData["role"]
	rc, rcOK := postData["role_capabilities"]

	if !ok || !rcOK || err != nil {
		common.RequestParamErrorTemplate(ctx, common.REQUEST_PARM_ERROR)
		return
	}
	var role mysql_model.Role
	var roleCapList []mysql_model.RoleCapabilities

	rb, _ := json.Marshal(r)
	rcb, _ := json.Marshal(rc)
	err = json.Unmarshal(rb, &role)
	if err != nil || role.RoleName == "" {
		common.RequestParamErrorTemplate(ctx, common.REQUEST_PARM_ERROR)
		return
	}
	err = json.Unmarshal(rcb, &roleCapList)
	if err != nil {
		common.RequestParamErrorTemplate(ctx, common.REQUEST_PARM_ERROR)
		return
	}
	// 补全
	roleCapList = mysql_model.RoleCapabilitiesMerge(roleCapList)
	err = client.MysqlClient.Transaction(func(tx *gorm.DB) error {
		err := client.MysqlClient.Model(&mysql_model.Role{}).Create(&role).Error
		if err != nil {
			return err
		}
		for i := range roleCapList {
			roleCapList[i].RoleID = *role.RecID
		}
		err = client.MysqlClient.Model(&mysql_model.RoleCapabilities{}).CreateInBatches(roleCapList, mysql.CalcMysqlBatchSize(roleCapList[0])).Error
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		common.InternalDataBaseErrorTemplate(ctx, common.DATABASE_INSERT_ERROR, mysql_model.Role{})
		return
	}
	ctx.JSON(_const.OK, errs.CreateSuccessMsg("创建角色成功！"))
	return
}

type TempRole struct {
	Role mysql_model.Role               `json:"role"`
	Rcs  []mysql_model.RoleCapabilities `json:"role_capabilities"`
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
		}}
	role := common.SelectMysqlTableContentWithCountTemplate(ctx,
		common.SelectMysqlTemplateOptions{
			DB:          client.MysqlClient,
			TableModel:  mysql_model.Role{},
			NotReturn:   true,
			NotPaginate: true,
		}, conditions...).([]mysql_model.Role)

	rcs := common.SelectMysqlTableContentWithCountTemplate(ctx,
		common.SelectMysqlTemplateOptions{
			DB:          client.MysqlClient,
			TableModel:  mysql_model.RoleCapabilities{},
			NotReturn:   true,
			NotPaginate: true,
		},
		common.MysqlCondition{
			Symbol:      mysql.EQUAL,
			ColumnName:  "rec_id",
			ColumnValue: roleID,
		},
		common.MysqlCondition{
			Symbol:      mysql.NULL,
			ColumnName:  "deleted_at",
			ColumnValue: " ",
		}).([]mysql_model.RoleCapabilities)

	if role != nil && rcs != nil {
		count := len(role)
		data := make([]interface{}, 0, count)

		for i := range role {
			r := role[i]
			tr := TempRole{
				Role: r,
				Rcs:  []mysql_model.RoleCapabilities{},
			}
			for j := range rcs {
				rc := rcs[j]
				if r.RecID != nil && rc.RoleID == *r.RecID {
					tr.Rcs = append(tr.Rcs, rc)
				}
			}
			data = append(data, tr)
		}
		common.GinPaginate(ctx, data)
		return
	} else {
		common.InternalDataBaseErrorTemplate(ctx, common.OTHER_ERROR, mysql_model.Role{})
		return
	}
}

type updateRoleRequest struct {
	Role             mysql_model.Role               `form:"role" json:"role"  `
	RoleCapabilities []mysql_model.RoleCapabilities `form:"role_capabilities" json:"role_capabilities" `
}

func UpdateRole(ctx *gin.Context) {
	var postData updateRoleRequest
	err := ctx.ShouldBindBodyWith(&postData, binding.JSON)
	if err != nil {
		common.RequestParamErrorTemplate(ctx, common.REQUEST_PARM_ERROR)
		return
	}
	roleCapList := mysql_model.RoleCapabilitiesMerge(postData.RoleCapabilities)

	err = client.MysqlClient.Transaction(func(tx *gorm.DB) error {
		err2 := tx.Updates(postData.Role).Error
		if err2 != nil {
			return err2
		}
		for i := range roleCapList {
			err2 := tx.Updates(roleCapList[i]).Error
			if err2 != nil {
				return err2
			}
		}
		return err2
	})
	if err != nil {
		common.InternalDataBaseErrorTemplate(ctx, common.DATABASE_UPDATE_ERROR, mysql_model.Role{})
		return
	}
	return
}

func DeleteRole(ctx *gin.Context) {
	recID := ctx.PostForm("role_id")
	if recID == "" {
		return
	}
	err := client.MysqlClient.Transaction(func(tx *gorm.DB) error {
		err := tx.Delete(&mysql_model.Role{}, recID).Error
		if err != nil {
			return err
		}
		err = tx.Delete(&mysql_model.RoleCapabilities{}, "role_id = ?", recID).Error
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
