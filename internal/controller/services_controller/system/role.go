package system

import (
	"github.com/XC-Zero/yinwan/pkg/client"
	_const "github.com/XC-Zero/yinwan/pkg/const"
	"github.com/XC-Zero/yinwan/pkg/model"
	"github.com/XC-Zero/yinwan/pkg/utils/errs"
	"github.com/XC-Zero/yinwan/pkg/utils/mysql"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"gorm.io/gorm"
)

type roleResponse struct {
	Role             model.Role               `json:"role" binding:"required,dive"`
	RoleCapabilities []model.RoleCapabilities `json:"role_capabilities" binding:"required,dive"`
}

func CreateRole(ctx *gin.Context) {
	var roleResponse roleResponse
	err := ctx.ShouldBindWith(&roleResponse, binding.Form)
	if err != nil {
		ctx.JSON(_const.REQUEST_PARM_ERROR, errs.CreateWebErrorMsg("输入有误！"))
		return
	}
	err = client.MysqlClient.Model(&model.Role{}).Create(&roleResponse.Role).Error
	if err != nil {
		ctx.JSON(_const.INTERNAL_ERROR, errs.CreateWebErrorMsg("创建角色失败！"))
		return
	}
	for i := range roleResponse.RoleCapabilities {
		roleResponse.RoleCapabilities[i].RoleID = *roleResponse.Role.RecID
	}
	err = client.MysqlClient.Model(&model.Role{}).CreateInBatches(roleResponse.RoleCapabilities, mysql.CalcMysqlBatchSize(roleResponse.RoleCapabilities[0])).Error
	if err != nil {
		ctx.JSON(_const.INTERNAL_ERROR, errs.CreateWebErrorMsg("创建角色内权限失败！"))
		return
	}

	ctx.JSON(_const.OK, errs.CreateSuccessMsg("创建角色成功！"))
	return
}

func SelectRole(ctx *gin.Context) {

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