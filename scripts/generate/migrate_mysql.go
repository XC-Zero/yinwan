package main

import (
	"context"
	"github.com/XC-Zero/yinwan/pkg/model"
	"github.com/XC-Zero/yinwan/pkg/utils/errs"
	"github.com/XC-Zero/yinwan/pkg/utils/logger"
	"github.com/XC-Zero/yinwan/pkg/utils/mysql"
	"github.com/fwhezfwhez/errorx"
	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/mongo"
	"gorm.io/gorm"
)

var systemMysqlMigrateList []interface{}
var bookNameMysqlMigrateList []interface{}
var mongoMigrateList []string
var moduleList []model.Module
var roleCapabilities []model.RoleCapabilities

// todo 供应商、客户是各个账套 独立？ or 共享？ (目前共享)
func init() {
	systemMysqlMigrateList = []interface{}{
		&model.Department{},
		&model.Role{},
		&model.RoleCapabilities{},
		&model.Module{},
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
		&model.StockInRecord{},
		&model.StockInRecord{},
		&model.FinanceCredentialEvent{},
		&model.EventItem{},
	}
	moduleList = []model.Module{
		{
			ModuleName: "storage",
		},
		{
			ModuleName: "finance",
		},
		{
			ModuleName: "system",
		},
		{
			ModuleName: "staff",
		},
		{
			ModuleName: "transaction",
		},
	}
	roleCapabilities = []model.RoleCapabilities{}
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
		return err
	}
	err = db.Transaction(func(tx *gorm.DB) error {
		err := tx.Exec("truncate modules").Error
		if err != nil {
			return err
		}
		err = tx.Exec("truncate role_capabilities ").Error
		if err != nil {
			return err
		}
		err = tx.Exec("truncate roles").Error
		if err != nil {
			return err
		}
		err = tx.Exec("truncate staffs").Error
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		logger.Error(errorx.MustWrap(err), "清空前置表失败！")
	}

	err = db.Model(&model.Module{}).CreateInBatches(moduleList, mysql.CalcMysqlBatchSize(moduleList[0])).Error
	if err != nil {
		logger.Error(errorx.MustWrap(err), "初始化系统模块失败！")
	}
	role := model.Role{
		RoleName: "root",
	}
	err = db.Model(&model.Role{}).Create(&role).Error
	if err != nil {
		logger.Error(errorx.MustWrap(err), "初始化系统超级管理员失败！")
	}
	for i := range moduleList {
		roleCapabilities = append(roleCapabilities, model.RoleCapabilities{
			RoelID:    *role.RecID,
			ModuleID:  *moduleList[i].RecID,
			CanRead:   true,
			CanWrite:  true,
			CanDelete: true,
		})
	}

	err = db.Model(&model.RoleCapabilities{}).CreateInBatches(roleCapabilities, mysql.CalcMysqlBatchSize(roleCapabilities[0])).Error
	if err != nil {
		logger.Error(errorx.MustWrap(err), "初始化系统超级管理员权限失败！")
	}

	err = db.Model(&model.Staff{}).Create(&model.Staff{
		BasicModel:    model.BasicModel{},
		StaffName:     "超级管理员",
		StaffAlias:    nil,
		StaffEmail:    "645171033@qq.com",
		StaffPhone:    nil,
		StaffPassword: "HSHROOT",
		StaffRoleID:   *role.RecID,
		StaffRoleName: "root",
	}).Error
	if err != nil {
		logger.Error(errorx.MustWrap(err), "初始化系统超级管理员权限失败！")
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
