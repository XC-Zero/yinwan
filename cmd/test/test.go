package main

import (
	"github.com/fatih/color"
)

func main() {
	//config.InitConfiguration()
	//go config.ViperMonitor()
	//
	////client.MysqlClient = client.InitMysqlGormV2(config.CONFIG.StorageConfig.MysqlConfig)
	//minio, err := client.InitMinio(config.CONFIG.StorageConfig.MinioConfig)
	//if err != nil {
	//	panic(err)
	//}
	//exists, err := minio.BucketExists(context.TODO(), "test01x")
	//if err != nil {
	//	panic(err)
	//}
	//log.Println(exists)
	//
	//services_controller.Starter()
	//select {}
	c := color.New(color.BgHiRed).Add(color.Underline)
	c.Println("Prints cyan text with an underline.")
	//err := client.MysqlClient.Raw(sqlBatch[1] ).Scan(&materialList).Error
	//if err != nil {
	//	return
	//}
	//err = client.MysqlClient.Raw(sqlBatch[0]).Scan(&count).Error
	//if err != nil {
	//	panic(err)
	//	return
	//}
	//ctx.JSON(_const.OK, gin.H{
	//	"count":         count,
	//	"material_list": materialList,
	//})
	return
}
