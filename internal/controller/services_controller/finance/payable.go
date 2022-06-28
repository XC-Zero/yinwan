package finance

import (
	"context"
	"fmt"
	"github.com/XC-Zero/yinwan/internal/controller/services_controller/common"
	_const "github.com/XC-Zero/yinwan/pkg/const"
	"github.com/XC-Zero/yinwan/pkg/model/mysql_model"
	"github.com/XC-Zero/yinwan/pkg/utils/errs"
	"github.com/XC-Zero/yinwan/pkg/utils/logger"
	"github.com/XC-Zero/yinwan/pkg/utils/mysql"
	"github.com/devfeel/mapper"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/pkg/errors"
	"strconv"
)

func CreatePayable(ctx *gin.Context) {
	bk, n := common.HarvestClientFromGinContext(ctx)
	if bk == nil {
		return
	}
	temp := mysql_model.Payable{}

	err := ctx.ShouldBindBodyWith(&temp, binding.JSON)
	if err != nil {
		logger.Error(errors.WithStack(err), "绑定模型失败!")
		common.RequestParamErrorTemplate(ctx, common.REQUEST_PARM_ERROR)
		return
	}
	err = bk.MysqlClient.WithContext(context.WithValue(context.Background(), "book_name", n)).
		Create(&temp).Error
	if err != nil {
		logger.Error(errors.WithStack(err), "绑定模型失败!")
		common.InternalDataBaseErrorTemplate(ctx, common.DATABASE_INSERT_ERROR, temp)
		return
	}
	ctx.JSON(_const.OK, errs.CreateSuccessMsg("新建应付成功!"))
	return
}

func SelectPayable(ctx *gin.Context) {
	bk, n := common.HarvestClientFromGinContext(ctx)
	if bk == nil {
		return
	}
	condition := []common.MysqlCondition{
		{
			Symbol:      mysql.NOT_EQUAL,
			ColumnName:  "deleted_at",
			ColumnValue: " ",
		},
	}
	op := common.SelectMysqlTemplateOptions{
		DB:         bk.MysqlClient.WithContext(context.WithValue(context.Background(), "book_name", n)),
		TableModel: mysql_model.Payable{},
		ResHookFunc: func(data []interface{}) []interface{} {
			var datum []interface{}
			m := mapper.NewMapper()
			for i := 0; i < len(data); i++ {
				payable, ok := data[i].(mysql_model.Payable)
				if !ok {
					continue
				}
				valMap := make(map[string]interface{})
				m.SetEnabledJsonTag(true)
				err := m.Mapper(payable, valMap)
				if err != nil {
					continue
				}
				valMap["payable_status"] = payable.PayableStatus.Display()
				datum = append(datum, valMap)
			}
			if len(datum) == 0 {
				return data
			}
			return datum
		},
	}
	common.SelectMysqlTableContentWithCountTemplate(ctx, op, condition...)
	return
}

func UpdatePayable(ctx *gin.Context) {
	bk, n := common.HarvestClientFromGinContext(ctx)
	if bk == nil {
		return
	}
	temp := mysql_model.Payable{}

	err := ctx.ShouldBindBodyWith(&temp, binding.JSON)
	if err != nil || temp.RecID == nil {
		logger.Error(errors.WithStack(err), "更新应付失败!")
		common.RequestParamErrorTemplate(ctx, common.REQUEST_PARM_ERROR)
		return
	}
	err = bk.MysqlClient.WithContext(context.WithValue(context.Background(), "book_name", n)).
		Updates(&temp).Where("rec_id = ?", temp.RecID).Error
	if err != nil {
		logger.Error(errors.WithStack(err), "更新应付记录失败!")
		common.InternalDataBaseErrorTemplate(ctx, common.DATABASE_UPDATE_ERROR, temp)
		return
	}
	ctx.JSON(_const.OK, errs.CreateSuccessMsg("更新应付记录成功!"))
	return
}

func DeletePayable(ctx *gin.Context) {
	bk, n := common.HarvestClientFromGinContext(ctx)
	if bk == nil {
		return
	}
	var payable mysql_model.Payable
	recID, err := strconv.Atoi(ctx.PostForm("payable_id"))
	if err != nil {
		common.RequestParamErrorTemplate(ctx, common.REQUEST_PARM_ERROR)
		return
	}
	payable.RecID = &recID
	err = bk.MysqlClient.WithContext(context.WithValue(context.Background(), "book_name", n)).
		Delete(&payable).Where("rec_id = ? ", recID).Error
	if err != nil {
		common.RequestParamErrorTemplate(ctx, common.REQUEST_PARM_ERROR)
		return
	}
	logger.Info(fmt.Sprintf("删除应付记录成功!记录ID:%d 操作人: %s", recID, common.HarvestEmailFromHeader(ctx)))
	ctx.JSON(_const.OK, errs.CreateSuccessMsg("删除应付记录成功!"))
	return
}

func CreatePayableDetail(ctx *gin.Context) {

}
func UpdatePayableDetail(ctx *gin.Context) {

}
func DeletePayableDetail(ctx *gin.Context) {

}
func SelectPayableDetail(ctx *gin.Context) {
	bk, n := common.HarvestClientFromGinContext(ctx)
	if bk == nil {
		return
	}
	condition := []common.MysqlCondition{
		{
			Symbol:      mysql.EQUAL,
			ColumnName:  "payable_id",
			ColumnValue: ctx.PostForm("payable_id"),
		},
		{
			Symbol:      mysql.NOT_EQUAL,
			ColumnName:  "deleted_at",
			ColumnValue: " ",
		},
	}
	op := common.SelectMysqlTemplateOptions{
		DB:         bk.MysqlClient.WithContext(context.WithValue(context.Background(), "book_name", n)),
		TableModel: mysql_model.PayableDetail{},
	}
	common.SelectMysqlTableContentWithCountTemplate(ctx, op, condition...)
	return
}
