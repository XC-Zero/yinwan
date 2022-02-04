package main

import (
	"github.com/XC-Zero/yinwan/pkg/client"
	"github.com/XC-Zero/yinwan/pkg/model"
	"github.com/pkg/errors"
)

var DataBase []interface{}

func init() {
	DataBase = append(DataBase, model.Staff{})

}
func main() {

	GenerateMysqlLogTables()

}

func GenerateMysqlSchema(schemaName string) error {
	client := client.MysqlClient
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

func GenerateMysqlTables() error {
	client := client.MysqlClient
	if client == nil {
		return errors.New("Mysql client is not init! ")
	}
	err := client.AutoMigrate()
	if err != nil {
		return errors.New("Mysql client is not init! ")
	}
	return nil
}

func GenerateMysqlLogTables() error {
	client := client.MysqlClient
	if client == nil {
		return errors.New("Mysql client is not init! ")
	}
	err := client.AutoMigrate(DataBase...)
	if err != nil {
		return errors.New("Mysql client is not init! ")
	}
	return nil
}
