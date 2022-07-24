package common

import (
	"context"
	"fmt"
	"github.com/XC-Zero/yinwan/pkg/client"
	_const "github.com/XC-Zero/yinwan/pkg/const"
	_interface "github.com/XC-Zero/yinwan/pkg/interface"
	"github.com/XC-Zero/yinwan/pkg/utils/convert"
	"github.com/XC-Zero/yinwan/pkg/utils/errs"
	"github.com/XC-Zero/yinwan/pkg/utils/es_tool"
	"github.com/XC-Zero/yinwan/pkg/utils/logger"
	myMongo "github.com/XC-Zero/yinwan/pkg/utils/mongo"
	"github.com/XC-Zero/yinwan/pkg/utils/mysql"
	"github.com/XC-Zero/yinwan/pkg/utils/tools"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/olivere/elastic/v7"
	"github.com/pkg/errors"
	"github.com/qiniu/qmgo"
	"go.mongodb.org/mongo-driver/bson"
	"gorm.io/gorm"
	"log"
	"reflect"
	"strconv"
	"time"
)

const (
	JSON = "application/json"
	FORM = "application/x-www-form-urlencoded"
)

// MysqlCondition MySQL 搜索条件
type MysqlCondition struct {
	Symbol      mysql.OperatorSymbol // 操作符
	ColumnName  string               // 相关字段
	ColumnValue string               // 字段值
}

// MongoCondition MongoDB 搜索条件
type MongoCondition struct {
	Symbol      myMongo.OperatorSymbol
	ColumnName  string
	ColumnValue interface{}
}

// Auto 临时结构
//	是否自动创建/删除/更新应收?
type Auto struct {
	Auto bool `json:"auto" form:"auto"`
}

func (auto Auto) WithContext(ctx context.Context) context.Context {
	return context.WithValue(ctx, "auto", auto.Auto)
}

// SelectMysqlTemplateOptions MySQL 搜索模板配置
type SelectMysqlTemplateOptions struct {
	DB            *gorm.DB                               // db
	TableModel    _interface.ChineseTabler               // 表
	OrderByColumn string                                 // order by 字段
	ResHookFunc   func(data []interface{}) []interface{} // 返回前预处理函数
	NotReturn     bool                                   // 不自动返回
	NotPaginate   bool                                   // 不自动分页
}

// SelectESTemplateOptions ElasticSearch 搜索模板配置
type SelectESTemplateOptions struct {
	TableModel  _interface.EsTabler
	Query       elastic.Query
	Scripts     string
	ResHookFunc func(data []interface{}) []interface{}
}

// CreateMysqlTemplateOptions MySQL 创建模板配置
type CreateMysqlTemplateOptions struct {
	DB         *gorm.DB
	TableModel _interface.ChineseTabler
	PreFunc    func(_interface.ChineseTabler) _interface.ChineseTabler
}

// SelectMongoDBTemplateOptions MongoDB 搜索模板配置
type SelectMongoDBTemplateOptions struct {
	DB            *qmgo.Database
	TableModel    _interface.ChineseTabler
	OrderByColumn string
	ResHookFunc   func(data []interface{}) []interface{}
}

type CreateMongoDBTemplateOptions struct {
	DB         *qmgo.Database
	Context    context.Context
	TableModel _interface.ChineseTabler
	PreFunc    func(_interface.ChineseTabler) _interface.ChineseTabler
	NotSyncES  bool
}

// MongoDBTemplateOptions MongoDB 更新/删除配置项
//	DB 			Mongo数据库连接
//	Context 	上下文传递参数或状态
//	TableModel	要操作的对象
//	PreFunc		操作前执行的函数
//	RecID		对象ID
//	OmitList	对象里忽略的字段
type MongoDBTemplateOptions struct {
	DB         *qmgo.Database
	Context    context.Context
	TableModel _interface.ChineseTabler
	PreFunc    func(_interface.ChineseTabler) _interface.ChineseTabler
	RecID      int
	OmitList   []string
	NotSyncES  bool
}

// UpdateMysqlTemplateOptions Mysql 更新模板配置
type UpdateMysqlTemplateOptions struct {
	CreateMysqlTemplateOptions
	RecID      int
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
	if op.DB == nil {
		logger.Error(errors.New("gorm client is nil"), "表名为:"+op.TableModel.TableName())
		InternalDataBaseErrorTemplate(ctx, BOOK_NAME_LACK_ERROR, op.TableModel)
		return nil
	}
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
	err := op.DB.Raw(contentSql).Scan(&dataList).Error
	if err != nil {
		InternalDataBaseErrorTemplate(ctx, DATABASE_SELECT_ERROR, op.TableModel)
		logger.Error(errors.WithStack(err), "Mysql 查询时错误!表名为: "+op.TableModel.TableName())
		return nil
	}
	err = op.DB.Raw(countSql).Scan(&count).Error
	if err != nil {
		InternalDataBaseErrorTemplate(ctx, DATABASE_COUNT_ERROR, op.TableModel)
		logger.Error(errors.WithStack(err), "Mysql 查询总数时错误!表名为: "+op.TableModel.TableName())
		return nil
	}

	var res = dataList
	if op.ResHookFunc != nil {
		sliceConvert, err := convert.SliceConvert(dataList, []interface{}{})
		if err == nil {
			if slice, ok := sliceConvert.([]interface{}); ok {
				res = op.ResHookFunc(slice)
			}
		} else {
			logger.Error(errors.WithStack(err), "Mysql 执行Hook函数错误!表名为: "+op.TableModel.TableName())
		}
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
	offset, limit := client.Paginate(ctx)
	var filter = bson.D{}
	var list = reflect.MakeSlice(reflect.SliceOf(reflect.TypeOf(op.TableModel)), 0, 0).Interface()

	if op.DB == nil {
		logger.Error(errors.New("gorm client is nil"), "表名为:"+op.TableModel.TableName())
		InternalDataBaseErrorTemplate(ctx, BOOK_NAME_LACK_ERROR, op.TableModel)
		return
	}
	var orderColumn string
	if orderColumn == "" {
		orderColumn = "rec_id"
	}
	if op.DB == nil {
		logger.Error(errors.New("mongo client is nil"), "表名为:"+op.TableModel.TableName())
		RequestParamErrorTemplate(ctx, BOOK_NAME_LACK_ERROR)
		return
	}
	var tx = op.DB.Collection(op.TableModel.TableName())
	if tx == nil {
		logger.Error(errors.New("mongo.Collection is nil! "), fmt.Sprintf("Mongo查询时没有该表或创建失败!(%s)", op.TableModel.TableName()))
		RequestParamErrorTemplate(ctx, BOOK_NAME_LACK_ERROR)
		return
	}
	for _, condition := range conditionList {
		if condition.ColumnValue == nil || reflect.ValueOf(condition.ColumnValue).IsZero() {
			continue
		}
		filter = append(filter, myMongo.TransMysqlOperatorSymbol(condition.Symbol, condition.ColumnName, condition.ColumnValue))
	}
	logger.Info(fmt.Sprintf(" filter is %+v", filter))
	c, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()
	err := tx.Find(c, filter).Sort(orderColumn).Limit(int64(limit)).Skip(int64(offset)).All(&list)

	if err != nil {
		logger.Error(errors.WithStack(err), fmt.Sprintf("Mongo查询时错误!表名为: %s ", op.TableModel.TableName()))
		InternalDataBaseErrorTemplate(ctx, DATABASE_SELECT_ERROR, op.TableModel)
		return
	}

	count, err := tx.Find(c, filter).Count()
	if err != nil {
		logger.Error(errors.WithStack(err), "Mongo查询总数时错误!表名为: "+op.TableModel.TableName())
		InternalDataBaseErrorTemplate(ctx, DATABASE_COUNT_ERROR, op.TableModel)
		return
	}
	var res = list
	if op.ResHookFunc != nil {
		sliceConvert, err := convert.SliceConvert(res, []interface{}{})
		if err == nil {
			if slice, ok := sliceConvert.([]interface{}); ok {
				res = op.ResHookFunc(slice)
			}
			logger.Error(errors.WithStack(err), "Mongo执行Hook函数错误!表名为: "+op.TableModel.TableName())
		}
	}

	SelectSuccessTemplate(ctx, count, res)
	return
}

// CreateOneMongoDBRecordTemplate MongoDB 创建模板
func CreateOneMongoDBRecordTemplate(ctx *gin.Context, op CreateMongoDBTemplateOptions) {
	var data = op.TableModel
	if op.PreFunc != nil {
		data = op.PreFunc(op.TableModel)
	}
	if data == nil {
		RequestParamErrorTemplate(ctx, REQUEST_PARM_ERROR)
		return
	}
	_, err := op.DB.Collection(op.TableModel.TableName()).InsertOne(op.Context, data)

	if err != nil && err != myMongo.CancelError {
		logger.Error(errors.WithStack(err), "Mongo 数据插入失败! 表:"+op.TableModel.TableName())
		InternalDataBaseErrorTemplate(ctx, DATABASE_INSERT_ERROR, op.TableModel)
		return
	}
	v, ok := data.(_interface.EsTabler)
	if ok && !op.NotSyncES {
		err := client.PutIntoIndex(v)
		if err != nil {
			logger.Error(errors.WithStack(err), "Mongo 数据同步插入 ES Data 失败!表:"+op.TableModel.TableName())
			InternalDataBaseErrorTemplate(ctx, DATABASE_INSERT_ERROR, data)
			return
		}
	}
	id := reflect.ValueOf(data).Field(0).FieldByName("RecID").Interface().(*int)
	mes := fmt.Sprintf("新建%s成功,编号为%d", data.TableCnName(), id)
	ctx.JSON(_const.OK, errs.CreateSuccessMsg(mes))
	return
}

// UpdateOneMongoDBRecordByIDTemplate MongoDB 通用更新
//	按REC_ID更新模板
func UpdateOneMongoDBRecordByIDTemplate(ctx *gin.Context, op MongoDBTemplateOptions) {
	var data = op.TableModel
	if op.PreFunc != nil {
		data = op.PreFunc(data)
	}
	filter, update := bson.D{}, bson.D{}
	filter = append(filter, myMongo.TransMysqlOperatorSymbol(myMongo.EQUAL, "rec_id", op.RecID))
	objV, objT := reflect.ValueOf(data), reflect.TypeOf(data)
	if objV.IsZero() {
		logger.Waring(errors.New("This struct is empty! "), "这里尝试更新,但空结构体~")
		ctx.JSON(_const.OK, errs.CreateSuccessMsg(fmt.Sprintf("更新%s信息成功！", data.TableCnName())))
		return
	}
	omitMap := tools.StringSliceToMap(op.OmitList)
	for i := 0; i < objT.NumField(); i++ {
		nowValue := objV.Field(i)
		if !nowValue.IsZero() {
			v, ok := objT.Field(i).Tag.Lookup("bson")
			if !ok {
				continue
			}
			if _, omit := omitMap[v]; omit {
				continue
			}
			if v == "updated_at" {
				update = append(update, bson.E{Key: v, Value: time.Now().String()})
				continue
			}
			update = append(update, bson.E{Key: v, Value: nowValue.Interface()})

		}

	}
	err := op.DB.Collection(op.TableModel.TableName()).UpdateOne(op.Context, filter, bson.D{{"$set", update}})
	if err != nil && err != myMongo.CancelError {
		InternalDataBaseErrorTemplate(ctx, DATABASE_UPDATE_ERROR, data)
		return
	}
	v, ok := op.TableModel.(_interface.EsTabler)
	if ok && !op.NotSyncES {
		err := client.UpdateIntoIndex(v, &op.RecID, op.Context, es_tool.ESDocToUpdateScript(v.ToESDoc()))
		if err != nil {
			logger.Error(errors.WithStack(err), "Mongo 同步更新 es 失败!")
			InternalDataBaseErrorTemplate(ctx, DATABASE_UPDATE_ERROR, data)
			return
		}
	}
	mes := fmt.Sprintf("更新%s信息成功！", data.TableCnName())
	logger.Info(mes)
	ctx.JSON(_const.OK, errs.CreateSuccessMsg(mes))
	return
}

// DeleteOneMongoDBRecordByIDTemplate MongoDB 按REC_ID更新模板
func DeleteOneMongoDBRecordByIDTemplate(ctx *gin.Context, op MongoDBTemplateOptions) {
	var data = op.TableModel
	var en = op.TableModel.TableName()
	if op.PreFunc != nil {
		data = op.PreFunc(data)
	}
	filter := bson.D{}
	filter = append(filter, myMongo.TransMysqlOperatorSymbol(myMongo.EQUAL, "rec_id", op.RecID))
	err := op.DB.Collection(op.TableModel.TableName()).UpdateOne(op.Context, filter, bson.D{{Key: "$set", Value: bson.E{Key: "delete_at", Value: time.Now()}}})

	if err != nil && err != myMongo.CancelError {
		logger.Error(errors.WithStack(err), "Mongo 软删除失败! 表:"+en)
		InternalDataBaseErrorTemplate(ctx, DATABASE_UPDATE_ERROR, data)
		return
	}
	v, ok := op.TableModel.(_interface.EsTabler)
	if ok && !op.NotSyncES {
		err := client.DeleteFromIndex(v, &op.RecID, op.Context)
		if err != nil {
			logger.Error(errors.WithStack(err), "Mongo 同步删除es 失败! 表:"+en)

			InternalDataBaseErrorTemplate(ctx, DATABASE_DELETE_ERROR, data)
			return
		}
	}
	mes := fmt.Sprintf("更新%s信息成功！", data.TableCnName())
	logger.Info(mes)
	ctx.JSON(_const.OK, errs.CreateSuccessMsg(mes))
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
	offset, limit := client.Paginate(ctx)
	list, count, err := client.GetFromIndex(op.TableModel, op.Query, offset, limit)
	log.Println(op.Query.Source())
	if err != nil {
		logger.Error(errors.WithStack(err), " 普适ES查询失败!")
		InternalDataBaseErrorTemplate(ctx, DATABASE_SELECT_ERROR, op.TableModel)
		return
	}
	ctx.JSON(_const.OK, gin.H{
		"count": count,
		"list":  list,
	})
	return

}

type bookNameRequest struct {
	BookName   string `json:"book_name" form:"book_name" `
	BookNameID string `json:"book_name_id" form:"book_name_id" `
}

func HarvestEmailFromHeader(ctx *gin.Context) string {
	return ctx.Request.Header.Get("staff_email")
}

// HarvestClientFromGinContext 从请求体里读取账套信息
func HarvestClientFromGinContext(ctx *gin.Context) (*client.BookName, string) {

	var bookNameJson bookNameRequest
	var bookName string

	if ctx.ContentType() == FORM {
		err := ctx.Request.ParseForm()
		if err != nil {
			logger.Error(errors.WithStack(err), "")
			RequestParamErrorTemplate(ctx, BOOK_NAME_LACK_ERROR)
			return nil, ""
		}
		bookName = ctx.Request.Form.Get("book_name")
	} else {
		err := ctx.ShouldBindBodyWith(&bookNameJson, binding.JSON)
		if err != nil {
			logger.Error(errors.WithStack(err), "")
			RequestParamErrorTemplate(ctx, BOOK_NAME_LACK_ERROR)
			return nil, ""
		}
		bookName = bookNameJson.BookName
	}
	if book, ok := client.ReadBookMap(bookName); ok {
		return &book, bookName
	} else {
		RequestParamErrorTemplate(ctx, BOOK_NAME_LACK_ERROR)
		return nil, ""
	}
}

// HarvestClientFromContext 从Context里读取账套信息
func HarvestClientFromContext(ctx context.Context) (*client.BookName, string) {
	bookName := ctx.Value("book_name").(string)
	bk, ok := client.ReadBookMap(bookName)
	if !ok {
		return nil, ""
	}
	return &bk, bk.BookName
}
