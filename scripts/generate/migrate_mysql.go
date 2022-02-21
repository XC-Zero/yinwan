package main

import (
	"context"
	"github.com/XC-Zero/yinwan/pkg/model"
	"github.com/XC-Zero/yinwan/pkg/utils/errs"
	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/mongo"
	"gorm.io/gorm"
)

var systemMysqlMigrateList []interface{}
var bookNameMysqlMigrateList []interface{}
var mongoMigrateList []string

// todo 供应商、客户是各个账套 独立？ or 共享？ (目前共享)
func init() {
	systemMysqlMigrateList = []interface{}{
		&model.Department{},
		&model.Staff{},
		&model.Commodity{},
		&model.CommodityBatch{},
		&model.Material{},
		&model.MaterialBatch{},
		&model.Transaction{},
		&model.Customer{},
		&model.Provider{},
		&model.ManipulationLog{},
	}

	bookNameMysqlMigrateList = []interface{}{
		&model.Payable{},
		&model.Receivable{},
		&model.FinanceCredential{},
		&model.FinanceCredentialEvent{},
		&model.EventItem{},
	}
}

func GenerateBookNameMigrateMysql(db *gorm.DB) error {
	if db == nil {
		return errors.New("BookName Mysql client is not init! ")
	}
	err := db.AutoMigrate(bookNameMysqlMigrateList)
	if err != nil {
		return err
	}
	return nil
}

func GenerateSystemMysqlTables(db *gorm.DB) error {

	if db == nil {
		return errors.New("Mysql client is not init! ")
	}
	err := db.AutoMigrate(systemMysqlMigrateList...)
	if err != nil {
		return errors.New("Mysql client is not init! ")
	}
	return nil
}

func MigrateMongo(db *mongo.Database) error {
	var errorList []error
	for i := range mongoMigrateList {
		err := db.CreateCollection(context.TODO(), mongoMigrateList[i])
		if err != nil {
			errorList = append(errorList, err)
		}
	}
	return errs.ErrorListToError(errorList)
}

func MigrateMinio() {

}
