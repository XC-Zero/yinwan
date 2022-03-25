package storage

import (
	"github.com/XC-Zero/yinwan/pkg/client"
	_const "github.com/XC-Zero/yinwan/pkg/const"
	"github.com/XC-Zero/yinwan/pkg/model"
	"github.com/XC-Zero/yinwan/pkg/utils/errs"
	"github.com/XC-Zero/yinwan/pkg/utils/mysql"
	"github.com/gin-gonic/gin"
)

// SelectMaterial 原材料
func SelectMaterial(ctx *gin.Context) {
	materialName := ctx.PostForm("material_name")
	recID := ctx.PostForm("material_id")
	materialTypeID := ctx.PostForm("material_type_id")
	var materialList []model.Material
	var count int
	sqlBatch := mysql.InitBatchSqlGeneration().
		AddSqlGeneration("count", mysql.InitSqlGeneration(model.Material{}, mysql.COUNT)).
		AddSqlGeneration("content", mysql.InitSqlGeneration(model.Material{}, mysql.ALL)).
		AddConditions("like", "material_name", "%"+materialName+"%").
		AddConditions("=", "material_id", recID, "material_type_id", materialTypeID)

	err := client.MysqlClient.Raw(
		sqlBatch.Harvest("content").
			AddGroupBy(mysql.BASIC_MODEL_PRIMARY_KEY).
			AddSuffixOther(client.PaginateSql(ctx)).
			HarvestSql()).
		Scan(&materialList).Error
	if err != nil {
		ctx.JSON(_const.INTERNAL_ERROR, errs.CreateWebErrorMsg("查询原材料列表失败！"))
		return
	}
	err = client.MysqlClient.Raw(sqlBatch.HarvestSql("count")).Scan(&count).Error
	if err != nil {
		ctx.JSON(_const.INTERNAL_ERROR, errs.CreateWebErrorMsg("查询原材料总数失败！"))
		return
	}
	ctx.JSON(_const.OK, gin.H{
		"count":         count,
		"material_list": materialList,
	})
	return
}

// SelectMaterialDetail 原材料批次信息
func SelectMaterialDetail(ctx *gin.Context) {
	var materialBatchList []model.MaterialBatch
	var count int64
	materialID := ctx.PostForm("material_id")
	warehouseID := ctx.PostForm("warehouse_id")
	if materialID == "" {
		return
	}
	err := client.MysqlClient.Model(&model.MaterialBatch{}).
		Scopes(client.Paginate(ctx)).
		Where(" material_id = ? ", materialID).
		Order("created_at").
		Find(&materialBatchList).
		Error
	if err != nil {
		ctx.JSON(_const.INTERNAL_ERROR, errs.CreateWebErrorMsg(""))
		return
	}
	err = client.MysqlClient.Where(" material_id = ? ", materialID).Count(&count).Error
	if err != nil {
		ctx.JSON(_const.INTERNAL_ERROR, errs.CreateWebErrorMsg(""))
		return
	}
	ctx.JSON(_const.OK, gin.H{
		"count":               count,
		"material_batch_list": materialBatchList,
	})
	return
}
