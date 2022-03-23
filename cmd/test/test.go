package main

import (
	"github.com/XC-Zero/yinwan/pkg/model"
	"log"
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

	mapping := model.QRCodeMapping{}
	var i = 8

	err := mapping.GenerateSql(model.Staff{
		BasicModel:          model.BasicModel{RecID: &i},
		StaffName:           "test01x",
		StaffAlias:          nil,
		StaffEmail:          "",
		StaffPhone:          nil,
		StaffPassword:       "",
		StaffPosition:       nil,
		StaffDepartmentID:   nil,
		StaffDepartmentName: nil,
		StaffRoleID:         0,
		StaffRoleName:       "",
	})
	if err != nil {
		panic(err)
	}
	log.Printf("%+v", mapping)
}
