package common

import (
	"context"
	"fmt"
	"github.com/XC-Zero/yinwan/pkg/client"
	_const "github.com/XC-Zero/yinwan/pkg/const"
	_interface "github.com/XC-Zero/yinwan/pkg/interface"
	"github.com/XC-Zero/yinwan/pkg/utils/convert"
	"github.com/XC-Zero/yinwan/pkg/utils/errs"
	my_mongo "github.com/XC-Zero/yinwan/pkg/utils/mongo"
	"github.com/XC-Zero/yinwan/pkg/utils/mysql"
	"github.com/fatih/color"
	"github.com/gin-gonic/gin"
	"github.com/olivere/elastic/v7"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"gorm.io/gorm"
	"log"
	"reflect"
	"strconv"
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

type EsCondition struct {
}

// SelectMysqlTemplateOptions MySQL 搜索模板配置
type SelectMysqlTemplateOptions struct {
	DB            *gorm.DB
	TableModel    _interface.ChineseTabler
	OrderByColumn string
	ResHookFunc   func(data []interface{}) []interface{}
	NotReturn     bool
	NotPaginate   bool
}

// SelectMongoDBTemplateOptions MongoDB 搜索模板配置
type SelectMongoDBTemplateOptions struct {
	DB            *mongo.Database
	TableModel    _interface.ChineseTabler
	OrderByColumn string
	ResHookFunc   func(data []interface{}) []interface{}
}

// SelectESTemplateOptions ElasticSearch 搜索模板配置
type SelectESTemplateOptions struct {
	DB           *elastic.Client
	TableModel   _interface.EsTabler
	Scripts      string
	NotHighLight bool
	ResHookFunc  func(data []interface{}) []interface{}
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
	----------------------------------    华丽的分割线   ----------------------------------------
*/

// SelectMysqlTableContentWithCountTemplate  Mysql 搜索模板
func SelectMysqlTableContentWithCountTemplate(ctx *gin.Context, op SelectMysqlTemplateOptions, conditionList ...MysqlCondition) interface{} {

	var count int
	// 根据传入的类型决定创建对应类型的切片
	var dataList = reflect.MakeSlice(reflect.SliceOf(reflect.TypeOf(op.TableModel)), 0, 0).Interface()

	if op.OrderByColumn == "" {
		op.OrderByColumn = "rec_id"
	}
	sqlBatch := mysql.InitBatchSqlGeneration().
		AddSqlGeneration("count", mysql.InitSqlGeneration(op.TableModel, mysql.COUNT)).
		AddSqlGeneration("content", mysql.InitSqlGeneration(op.TableModel, mysql.ALL))
	for i := range conditionList {
		sqlBatch.AddConditions(conditionList[i].Symbol, conditionList[i].ColumnName, conditionList[i].ColumnValue)
	}
	sg := sqlBatch.Harvest("content").AddOrderBy(op.OrderByColumn)
	if !op.NotPaginate {
		sg = sg.AddSuffixOther(client.MysqlPaginateSql(ctx))
	}
	contentSql, countSql := sg.HarvestSql(), sqlBatch.HarvestSql("count")
	c := color.New(color.BgMagenta).Add(color.Underline)
	// 打印成功与否并不重要，error 忽略掉就行
	c.Println(contentSql)
	c.Println(countSql)

	err := op.DB.Raw(contentSql).Scan(&dataList).Error
	if err != nil {
		InternalDataBaseErrorTemplate(ctx, DATABASE_SELECT_ERROR, op.TableModel)

		return nil
	}
	err = op.DB.Raw(countSql).Scan(&count).Error
	if err != nil {
		InternalDataBaseErrorTemplate(ctx, DATABASE_COUNT_ERROR, op.TableModel)
		return nil
	}

	var res = dataList
	if op.ResHookFunc != nil {
		sliceConvert, err := convert.SliceConvert(dataList, []interface{}{})
		if err == nil {
			if slice, ok := sliceConvert.([]interface{}); ok {
				res = op.ResHookFunc(slice)
			}
		}
		log.Println(err)
	}

	if !op.NotReturn {
		SelectSuccessTemplate(ctx, int64(count), res)
		return nil
	}
	return res
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
	var res interface{} = list
	if op.ResHookFunc != nil {
		sliceConvert, err := convert.SliceConvert(list, []interface{}{})
		if err == nil {
			res = op.ResHookFunc(sliceConvert.([]interface{}))
		}
		log.Println(err)
	}
	SelectSuccessTemplate(ctx, count, res)
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

/*
	----------------------------------    华丽的分割线   ----------------------------------------
*/

// GinPaginate 在service层面分页
func GinPaginate(ctx *gin.Context, data []interface{}) {
	pageNumber := ctx.PostForm("page_number")
	pageSize := ctx.PostForm("page_size")
	n, limit := 0, 0
	if pn, err := strconv.Atoi(pageNumber); err != nil {
		n = 1
	} else {
		n = pn
	}
	if ps, err := strconv.Atoi(pageSize); err != nil {
		limit = 10
	} else {
		limit = ps
	}

	offset := (n - 1) * limit
	end := offset + limit
	length := len(data)
	if offset > length {
		data = nil
	} else if end > length {
		data = data[offset:]
	} else {
		data = data[offset:end]
	}
	ctx.JSON(_const.OK, gin.H{
		"count": length,
		"list":  data,
	})
	return
}

/*
	----------------------------------    华丽的分割线   ----------------------------------------
*/

func SelectESTableContentWithCountTemplate(ctx *gin.Context, op SelectESTemplateOptions) {
	//query := make(map[string]interface{}, 0)
	//
	//if op.DB == nil {
	//	InternalDataBaseErrorTemplate(ctx, DATABASE_SELECT_ERROR, op.TableModel)
	//	return
	//}
	//do, err := client.ESPaginate(ctx,
	//	op.DB.Search(op.TableModel.TableName()).Query(elastic.NewQueryStringQuery())).
	//	Pretty(true).
	//	Do(context.Background())
	//if err != nil {
	//	return
	//}

}
func HarvestClientFromGinContext(ctx *gin.Context) *client.BookName {
	bookName := ctx.PostForm("book_name")
	if bookName == "" {
		RequestParamErrorTemplate(ctx, BOOK_NAME_LACK_ERROR)
		return nil
	}
	if book, ok := client.ReadBookMap(bookName); ok {
		return &book
	} else {
		RequestParamErrorTemplate(ctx, BOOK_NAME_LACK_ERROR)

		return nil
	}
}
