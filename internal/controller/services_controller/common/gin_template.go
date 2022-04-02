package common

import (
	"context"
	"github.com/XC-Zero/yinwan/pkg/client"
	_interface "github.com/XC-Zero/yinwan/pkg/interface"
	my_mongo "github.com/XC-Zero/yinwan/pkg/utils/mongo"
	"github.com/XC-Zero/yinwan/pkg/utils/mysql"
	"github.com/fatih/color"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"gorm.io/gorm"
	"time"
)

type Condition struct {
	Symbol      mysql.OperatorSymbol
	ColumnName  string
	ColumnValue string
}

type MongoCondition struct {
	Symbol      my_mongo.OperatorSymbol
	ColumnName  string
	ColumnValue interface{}
}

// SelectTableContentWithCountMysqlTemplate 通用 Mysql 查询模板
// 入参为
// ctx             略
// db              执行语句的数据库
// tableModel      结构体
// orderByColumn   OrderBy 的字段 默认为 id
// conditionList   条件
// resHookFunc     返回前处理函数
// 返回给前端俩字段
// count
func SelectTableContentWithCountMysqlTemplate(ctx *gin.Context, db *gorm.DB, tableModel _interface.ChineseTabler, orderByColumn string, resHookFunc func(data []_interface.ChineseTabler) []_interface.ChineseTabler, conditionList ...Condition) {
	var count int
	var dataList []_interface.ChineseTabler
	if orderByColumn == "" {
		orderByColumn = "id"
	}
	sqlBatch := mysql.InitBatchSqlGeneration().
		AddSqlGeneration("count", mysql.InitSqlGeneration(tableModel, mysql.COUNT)).
		AddSqlGeneration("content", mysql.InitSqlGeneration(tableModel, mysql.ALL))
	for i := range conditionList {
		sqlBatch.AddConditions(conditionList[i].Symbol, conditionList[i].ColumnName, conditionList[i].ColumnValue)
	}

	contentSql, countSql := sqlBatch.Harvest("content").AddOrderBy(orderByColumn).
		AddSuffixOther(client.MysqlPaginateSql(ctx)).HarvestSql(), sqlBatch.HarvestSql("count")
	c := color.New(color.BgMagenta).Add(color.Underline)
	// 打印成功与否并不重要，error 忽略掉就行
	c.Println(contentSql)
	c.Println(countSql)

	err := db.Raw(contentSql).Scan(&dataList).Error
	if err != nil {
		InternalDataBaseErrorTemplate(ctx, DATABASE_SELECT_ERROR, tableModel)
		return
	}
	err = db.Raw(countSql).Scan(&count).Error
	if err != nil {
		InternalDataBaseErrorTemplate(ctx, DATABASE_COUNT_ERROR, tableModel)
		return
	}
	if resHookFunc != nil {
		dataList = resHookFunc(dataList)
	}

	SelectSuccessTemplate(ctx, int64(count), dataList)
	return
}

func SelectTableContentWithCountMongoDBTemplate(ctx *gin.Context, db *mongo.Database, tableModel _interface.ChineseTabler, orderByColumn string, resHookFunc func(data []_interface.ChineseTabler) []_interface.ChineseTabler, conditionList ...MongoCondition) {
	// 分页参数
	var findOptions = client.MongoPaginate(ctx, &options.FindOptions{})
	var countOptions = options.Count().SetMaxTime(3 * time.Second)
	var tx = db.Collection(tableModel.TableName())
	var filter = bson.D{}
	var list []_interface.ChineseTabler

	if db == nil || tx == nil {
		RequestParamErrorTemplate(ctx, BOOK_NAME_LACK_ERROR)
		return
	}
	for _, condition := range conditionList {
		filter = append(filter, my_mongo.TransMysqlOperatorSymbol(condition.Symbol, condition.ColumnName, condition.ColumnValue))
	}
	findOptions.Sort = bson.D{{orderByColumn, 1}}
	data, err := tx.Find(context.TODO(), filter, findOptions)
	if err != nil {
		InternalDataBaseErrorTemplate(ctx, DATABASE_SELECT_ERROR, tableModel)
		return
	}
	err = data.Decode(&list)
	if err != nil {
		InternalDataBaseErrorTemplate(ctx, DATABASE_SELECT_ERROR, tableModel)
		return
	}
	count, err := tx.CountDocuments(context.TODO(), filter, countOptions)
	if err != nil {
		InternalDataBaseErrorTemplate(ctx, DATABASE_COUNT_ERROR, tableModel)
		return
	}
	if resHookFunc != nil {
		list = resHookFunc(list)
	}
	SelectSuccessTemplate(ctx, count, list)
	return
}
