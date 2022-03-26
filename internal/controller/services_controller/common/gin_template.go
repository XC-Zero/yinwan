package common

import (
	"fmt"
	"github.com/XC-Zero/yinwan/pkg/client"
	_const "github.com/XC-Zero/yinwan/pkg/const"
	_interface "github.com/XC-Zero/yinwan/pkg/interface"
	"github.com/XC-Zero/yinwan/pkg/utils/errs"
	"github.com/XC-Zero/yinwan/pkg/utils/mysql"
	"github.com/fatih/color"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Condition struct {
	Symbol      mysql.OperatorSymbol
	ColumnName  string
	ColumnValue string
}

// SelectTableContentWithCountTemplate 通用查询模板
// 入参为
// ctx             略
// db              执行语句的数据库
// tableModel      结构体
// conditionList   条件
// 返回给前端俩字段
// count
func SelectTableContentWithCountTemplate(ctx *gin.Context, db *gorm.DB, tableModel _interface.ChineseTabler, conditionList ...Condition) {
	var count int
	var dataList []_interface.ChineseTabler

	sqlBatch := mysql.InitBatchSqlGeneration().
		AddSqlGeneration("count", mysql.InitSqlGeneration(tableModel, mysql.COUNT)).
		AddSqlGeneration("content", mysql.InitSqlGeneration(tableModel, mysql.ALL))
	for i := range conditionList {
		sqlBatch.AddConditions(conditionList[i].Symbol, conditionList[i].ColumnName, conditionList[i].ColumnValue)
	}

	contentSql, countSql := sqlBatch.Harvest("content").AddOrderBy("id").
		AddSuffixOther(client.PaginateSql(ctx)).HarvestSql(), sqlBatch.HarvestSql("count")
	c := color.New(color.BgMagenta).Add(color.Underline)
	// 打印成功与否并不重要，error 忽略掉就行
	c.Println(contentSql)
	c.Println(countSql)

	err := db.Raw(contentSql).Scan(&dataList).Error
	if err != nil {
		ctx.JSON(_const.INTERNAL_ERROR, gin.H(errs.CreateWebErrorMsg(fmt.Sprintf("查询%s内容失败！", tableModel.TableCnName()))))
		return
	}
	err = db.Raw(countSql).Scan(&count).Error
	if err != nil {
		ctx.JSON(_const.INTERNAL_ERROR, gin.H(errs.CreateWebErrorMsg(fmt.Sprintf("查询%s总数失败！", tableModel.TableCnName()))))
		return
	}
	ctx.JSON(_const.OK, gin.H{
		"count": count,
		"list":  dataList,
	})
	return
}
