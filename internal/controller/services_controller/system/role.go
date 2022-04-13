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
	"log"
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
	roleCapList = model.RoleCapabilitiesMerge(roleCapList)
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

type TempRole struct {
	Role model.Role               `json:"role"`
	Rcs  []model.RoleCapabilities `json:"role_capabilities"`
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
			TableModel:  model.Role{},
			NotReturn:   true,
			NotPaginate: true,
		}, conditions...).([]model.Role)

	rcs := common.SelectMysqlTableContentWithCountTemplate(ctx,
		common.SelectMysqlTemplateOptions{
			DB:          client.MysqlClient,
			TableModel:  model.RoleCapabilities{},
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
		}).([]model.RoleCapabilities)
	log.Printf("%+v\n\n", role)
	log.Printf("%+v\n", rcs)

	if role != nil && rcs != nil {
		count := len(role)
		data := make([]interface{}, 0, count)

		for i := range role {
			r := role[i]
			tr := TempRole{
				Role: r,
				Rcs:  []model.RoleCapabilities{},
			}
			for j := range rcs {
				rc := rcs[j]
				if r.RecID != nil && rc.RoleID == *r.RecID {
					tr.Rcs = append(tr.Rcs, rc)
				}
			}
			log.Println(tr)
			data = append(data, tr)
		}
		log.Println(data)
		common.GinPaginate(ctx, data)
		return
	} else {
		common.InternalDataBaseErrorTemplate(ctx, common.OTHER_ERROR, model.Role{})
		return
	}
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
