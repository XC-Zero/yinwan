package main

import (
	"github.com/XC-Zero/yinwan/internal/config"
	"github.com/XC-Zero/yinwan/pkg/model"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

var DataBase []interface{}

func init() {
	DataBase = append(DataBase, model.Staff{})

}
func main() {
	config.InitConfiguration()
	go config.ViperMonitor()
	// todo 初始化配置文件
	//client.MysqlClient = client.InitMysqlGormV2(config.CONFIG.StorageConfig.MysqlConfig)
	//GenerateMysqlSchema()

}

func GenerateMysqlSchema(client *gorm.DB, schemaName string) error {
	if client == nil {
		return errors.New("Mysql client is not init! ")
	}

	err := client.Raw(`CREATE DATABASE IF NOT EXISTS ? default charset utf8 COLLATE utf8_general_ci `, schemaName).Error
	if err != nil {
		return err
	}
	client.Raw("C")
	return nil
}

func GenerateMysqlTables(client *gorm.DB) error {

	if client == nil {
		return errors.New("Mysql client is not init! ")
	}
	err := client.AutoMigrate(DataBase...)
	if err != nil {
		return errors.New("Mysql client is not init! ")
	}
	return nil
}

func GenerateMysqlLogTables(client *gorm.DB) error {
	if client == nil {
		return errors.New("Mysql client is not init! ")
	}
	err := client.AutoMigrate()
	if err != nil {
		return errors.New("Mysql client is not init! ")
	}
	return nil
}
