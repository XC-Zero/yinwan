package common

import (
	"context"
	"fmt"
	"github.com/XC-Zero/yinwan/pkg/client"
	_const "github.com/XC-Zero/yinwan/pkg/const"
	_interface "github.com/XC-Zero/yinwan/pkg/interface"
	"github.com/XC-Zero/yinwan/pkg/utils/errs"
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

// MysqlCondition MySQL 搜索条件
type MysqlCondition struct {
	Symbol      mysql.OperatorSymbol
	ColumnName  string
	ColumnValue string
}

// MongoCondition MongoDB 搜索条件
type MongoCondition struct {
	Symbol      my_mongo.OperatorSymbol
	ColumnName  string
	ColumnValue interface{}
}

// SelectMysqlTemplateOptions MySQL 搜索模板配置
type SelectMysqlTemplateOptions struct {
	DB            *gorm.DB
	TableModel    _interface.ChineseTabler
	OrderByColumn string
	ResHookFunc   func(data []_interface.ChineseTabler) []_interface.ChineseTabler
}

// SelectMongoDBTemplateOptions MongoDB 搜索模板配置
type SelectMongoDBTemplateOptions struct {
	DB            *mongo.Database
	TableModel    _interface.ChineseTabler
	OrderByColumn string
	ResHookFunc   func(data []_interface.ChineseTabler) []_interface.ChineseTabler
}

// CreateMysqlTemplateOptions MySQL 创建模板配置
type CreateMysqlTemplateOptions struct {
	DB         *gorm.DB
	TableModel _interface.ChineseTabler
	PreFunc    func(_interface.ChineseTabler) _interface.ChineseTabler
}

// CreateMongoDBTemplateOptions MongoDB 创建模板配置
type CreateMongoDBTemplateOptions struct {
	DB         *mongo.Database
	TableModel _interface.ChineseTabler
	PreFunc    func(_interface.ChineseTabler) _interface.ChineseTabler
}

// UpdateMysqlTemplateOptions Mysql 更新模板配置
type UpdateMysqlTemplateOptions struct {
	CreateMysqlTemplateOptions
	OmitColumn []string
}

/*
	----------------------------------    华丽的分割线   ---------------------------------
*/

// SelectMysqlTableContentWithCountTemplate  Mysql 搜索模板
func SelectMysqlTableContentWithCountTemplate(ctx *gin.Context, op SelectMysqlTemplateOptions, conditionList ...MysqlCondition) {
	var count int
	var dataList []_interface.ChineseTabler
	if op.OrderByColumn == "" {
		op.OrderByColumn = "id"
	}
	sqlBatch := mysql.InitBatchSqlGeneration().
		AddSqlGeneration("count", mysql.InitSqlGeneration(op.TableModel, mysql.COUNT)).
		AddSqlGeneration("content", mysql.InitSqlGeneration(op.TableModel, mysql.ALL))
	for i := range conditionList {
		sqlBatch.AddConditions(conditionList[i].Symbol, conditionList[i].ColumnName, conditionList[i].ColumnValue)
	}

	contentSql, countSql := sqlBatch.Harvest("content").AddOrderBy(op.OrderByColumn).
		AddSuffixOther(client.MysqlPaginateSql(ctx)).HarvestSql(), sqlBatch.HarvestSql("count")
	c := color.New(color.BgMagenta).Add(color.Underline)
	// 打印成功与否并不重要，error 忽略掉就行
	c.Println(contentSql)
	c.Println(countSql)

	err := op.DB.Raw(contentSql).Scan(&dataList).Error
	if err != nil {
		InternalDataBaseErrorTemplate(ctx, DATABASE_SELECT_ERROR, op.TableModel)
		return
	}
	err = op.DB.Raw(countSql).Scan(&count).Error
	if err != nil {
		InternalDataBaseErrorTemplate(ctx, DATABASE_COUNT_ERROR, op.TableModel)
		return
	}
	if op.ResHookFunc != nil {
		dataList = op.ResHookFunc(dataList)
	}

	SelectSuccessTemplate(ctx, int64(count), dataList)
	return
}

// SelectMongoDBTableContentWithCountTemplate MongoDB 搜索模板
func SelectMongoDBTableContentWithCountTemplate(ctx *gin.Context, op SelectMongoDBTemplateOptions, conditionList ...MongoCondition) {
	// 分页参数
	var findOptions = client.MongoPaginate(ctx, &options.FindOptions{})
	var countOptions = options.Count().SetMaxTime(3 * time.Second)
	var filter = bson.D{}
	var list []_interface.ChineseTabler

	if op.DB == nil {
		RequestParamErrorTemplate(ctx, BOOK_NAME_LACK_ERROR)
		return
	}
	var tx = op.DB.Collection(op.TableModel.TableName())
	if tx == nil {
		RequestParamErrorTemplate(ctx, BOOK_NAME_LACK_ERROR)
		return
	}
	for _, condition := range conditionList {
		filter = append(filter, my_mongo.TransMysqlOperatorSymbol(condition.Symbol, condition.ColumnName, condition.ColumnValue))
	}
	findOptions.Sort = bson.D{{op.OrderByColumn, 1}}
	data, err := tx.Find(context.TODO(), filter, findOptions)
	if err != nil {
		InternalDataBaseErrorTemplate(ctx, DATABASE_SELECT_ERROR, op.TableModel)
		return
	}
	err = data.Decode(&list)
	if err != nil {
		InternalDataBaseErrorTemplate(ctx, DATABASE_SELECT_ERROR, op.TableModel)
		return
	}
	count, err := tx.CountDocuments(context.TODO(), filter, countOptions)
	if err != nil {
		InternalDataBaseErrorTemplate(ctx, DATABASE_COUNT_ERROR, op.TableModel)
		return
	}
	if op.ResHookFunc != nil {
		list = op.ResHookFunc(list)
	}
	SelectSuccessTemplate(ctx, count, list)
	return
}

// CreateOneMysqlRecordTemplate MySQL 创建模板
func CreateOneMysqlRecordTemplate(ctx *gin.Context, op CreateMysqlTemplateOptions) {
	var data = op.TableModel
	if op.PreFunc != nil {
		data = op.PreFunc(op.TableModel)
	}
	err := op.DB.Create(op.TableModel).Error
	if err != nil {
		InternalDataBaseErrorTemplate(ctx, DATABASE_INSERT_ERROR, data)
		return
	}
	ctx.JSON(_const.OK, errs.CreateSuccessMsg(fmt.Sprintf("新建%s成功", data.TableCnName())))
	return
}

// CreateOneMongoDBRecordTemplate MongoDB 创建模板
func CreateOneMongoDBRecordTemplate(ctx *gin.Context, op CreateMongoDBTemplateOptions) {
	var data = op.TableModel
	if op.PreFunc != nil {
		data = op.PreFunc(op.TableModel)
	}
	res, err := op.DB.Collection(op.TableModel.TableName()).InsertOne(context.TODO(), data)

	if err != nil {
		InternalDataBaseErrorTemplate(ctx, DATABASE_INSERT_ERROR, op.TableModel)
		return
	}
	ctx.JSON(_const.OK, errs.CreateSuccessMsg(fmt.Sprintf("新建%s成功,编号为%s", data.TableCnName(), res.InsertedID)))
	return
}

// UpdateOneMysqlRecordTemplate MySQL 更新模板
func UpdateOneMysqlRecordTemplate(ctx *gin.Context, op UpdateMysqlTemplateOptions) {
	var data = op.TableModel
	if op.PreFunc != nil {
		data = op.PreFunc(data)
	}
	err := op.DB.Updates(data).Omit(op.OmitColumn...).Error
	if err != nil {
		InternalDataBaseErrorTemplate(ctx, DATABASE_UPDATE_ERROR, data)
		return
	}
	ctx.JSON(_const.OK, errs.CreateSuccessMsg(fmt.Sprintf("更新%s信息成功！", data.TableCnName())))
	return
}
