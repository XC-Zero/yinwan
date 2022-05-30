package main

import (
	_interface "github.com/XC-Zero/yinwan/pkg/interface"
	"github.com/XC-Zero/yinwan/pkg/model/mongo_model"
	"go.mongodb.org/mongo-driver/bson"
	"log"
	"reflect"
)

func main() {
	d := bson.D{{"stock_in_record_owner_id", nil}, {"rec_id", 1923}}
	//info := mysql_model.BookNameInfo{
	//	BookNameID: "a",
	//	BookName:   "b",
	//}
	a := mongo_model.StockInRecord{}
	var b _interface.ChineseTabler = a
	reflectToStruct(d, b)

	marshal, err := bson.Marshal(d)
	if err != nil {
		panic(err)
	}
	err = bson.Unmarshal(marshal, &a)
	if err != nil {
		panic(err)
	}

	log.Printf("%+v", a)
	// 读取配置文件
	//config.InitConfiguration()
	//// 开协程监听配置文件修改，实现热加载
	//go config.ViperMonitor()
	//client.InitSystemStorage(config.CONFIG.StorageConfig)
	////indexes, err := client.ESClient.CatIndices().Do(context.TODO())
	////if err != nil {
	////	return
	////}
	////for _, index := range indexes {
	////	log.Println(index.Index)
	////}
	//bk, _ := client.ReadBookMap("sss")
	//conditions := []common.MongoCondition{
	//
	//	{
	//		Symbol:      my_mongo.EQUAL,
	//		ColumnName:  "basicmodel.deleted_at",
	//		ColumnValue: bsontype.Null,
	//	},
	//}
	//options := common.SelectMongoDBTemplateOptions{
	//	DB:         bk.MongoDBClient,
	//	TableModel: mongo_model.StockInRecord{},
	//}
	//common.SelectMongoDBTableContentWithCountTemplate(nil, options, conditions...)
	//err := client.MysqlClient.AutoMigrate(mysql_model.MaterialHistoryCost{})
	//if err != nil {
	//	panic(err)
	//}

	//err = client.CreateIndex(&es.Commodity{})
	//if err != nil {
	//	panic(err)
	//}
	//var a = 89764
	//var r = "这他喵又不是测试的？？？？？？"
	//material := es.Material{
	//	BasicModel:   es.BasicModel{RecID: &a},
	//	MaterialName: "他喵???",
	//	Remark:       &r,
	//}
	//err = client.PutIntoIndex(&material)
	//if err != nil {
	//	panic(err)
	//}
	//query := elastic.NewMultiMatchQuery("他喵", "rec_id^999", "remark^2", "material_name^10").Operator("or")
	//
	//data, count, err := client.GetFromIndex(common.Material{}, query, 0, 2)
	//if err != nil {
	//	panic(err)
	//}
	//log.Println(count)
	//log.Printf("%T  \n %+v", data, data)
}

func reflectToStruct(d bson.D, tabler _interface.ChineseTabler) {
	str := reflect.New(reflect.TypeOf(tabler))
	marshal, err := bson.Marshal(d)
	if err != nil {
		panic(err)
	}
	err = bson.Unmarshal(marshal, &str)
	if err != nil {
		panic(err)
	}

	log.Printf("%+v %T", str, str)
}
