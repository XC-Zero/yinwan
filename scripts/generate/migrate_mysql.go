package main

import (
	"context"
	"fmt"
	"github.com/XC-Zero/yinwan/pkg/model"
	"github.com/XC-Zero/yinwan/pkg/utils/errs"
	"github.com/XC-Zero/yinwan/pkg/utils/logger"
	"github.com/XC-Zero/yinwan/pkg/utils/mysql"
	"github.com/fwhezfwhez/errorx"
	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/mongo"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

var systemMysqlMigrateList []interface{}
var bookNameMysqlMigrateList []interface{}
var mongoMigrateList []string
var moduleList []model.Module
var roleCapabilities []model.RoleCapabilities
var departmentList []model.Department
var allRole model.Role
var readRole model.Role
var readWriteRole model.Role
var typeTreeList []model.TypeTree

func init() {
	systemMysqlMigrateList = []interface{}{
		&model.Department{},
		&model.Role{},
		&model.RoleCapabilities{},
		&model.Module{},
		&model.Staff{},
		&model.Commodity{},
		&model.CommodityHistoricalCost{},
		&model.CommodityBatch{},
		&model.Material{},
		&model.MaterialBatch{},
		&model.Transaction{},
		&model.Customer{},
		&model.Provider{},
		&model.ManipulationLog{},
		&model.TypeTree{},
	}
	bookNameMysqlMigrateList = []interface{}{
		&model.Payable{},
		&model.Receivable{},
		&model.Purchase{},
		&model.FinanceCredential{},
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
	allRole = model.Role{
		RoleName: "root",
	}
	readRole = model.Role{RoleName: "只读账号"}
	readWriteRole = model.Role{RoleName: "读写账号"}

	roleCapabilities = []model.RoleCapabilities{}
	departmentManagerName, departmentManagerID, finAddr := "超级管理员", 1, "2楼205"
	departmentList = []model.Department{
		{
			BasicModel:            model.BasicModel{},
			DepartmentName:        "技术部",
			DepartmentLocation:    nil,
			DepartmentManagerID:   &departmentManagerID,
			DepartmentManagerName: &departmentManagerName,
		}, {
			BasicModel:            model.BasicModel{},
			DepartmentName:        "财务部",
			DepartmentLocation:    &finAddr,
			DepartmentManagerID:   &departmentManagerID,
			DepartmentManagerName: &departmentManagerName,
		}, {
			BasicModel:            model.BasicModel{},
			DepartmentName:        "生产车间",
			DepartmentLocation:    nil,
			DepartmentManagerID:   &departmentManagerID,
			DepartmentManagerName: &departmentManagerName,
		}, {
			BasicModel:            model.BasicModel{},
			DepartmentName:        "仓库",
			DepartmentLocation:    nil,
			DepartmentManagerID:   &departmentManagerID,
			DepartmentManagerName: &departmentManagerName,
		}, {
			BasicModel:            model.BasicModel{},
			DepartmentName:        "销售部",
			DepartmentLocation:    nil,
			DepartmentManagerID:   &departmentManagerID,
			DepartmentManagerName: &departmentManagerName,
		}, {
			BasicModel:            model.BasicModel{},
			DepartmentName:        "人事部",
			DepartmentLocation:    nil,
			DepartmentManagerID:   &departmentManagerID,
			DepartmentManagerName: &departmentManagerName,
		}, {
			BasicModel:            model.BasicModel{},
			DepartmentName:        "行政部",
			DepartmentLocation:    nil,
			DepartmentManagerID:   &departmentManagerID,
			DepartmentManagerName: &departmentManagerName,
		}}
	typeTreeList = []model.TypeTree{
		{model.BasicModel{}, "固定资产", nil, nil},
		{model.BasicModel{}, "产成品", nil, nil},
		{model.BasicModel{}, "原材料", nil, nil},
		{model.BasicModel{}, "周转材料", nil, nil},
		{model.BasicModel{}, "低值易耗品", nil, nil},
		{model.BasicModel{}, "其他类型 ", nil, nil},
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
	var staffAlias, staffPosition = "超管", "系统开发攻城狮"

	if db == nil {
		return errors.New("Mysql client is not init! ")
	}
	// 初始化建表
	err := db.AutoMigrate(systemMysqlMigrateList...)
	if err != nil {
		return err
	}

	err = db.Transaction(func(tx *gorm.DB) error {
		// 清空所有表
		for i := range systemMysqlMigrateList {
			object, ok := systemMysqlMigrateList[i].(schema.Tabler)
			if !ok {
				panic(fmt.Sprintf("第 %d 张表没实现tableName接口！", i))
			}
			err := tx.Exec(fmt.Sprintf("truncate %s;", object.TableName())).Error
			if err != nil {
				logger.Error(errorx.MustWrap(err), "清空前置表失败！")
				return err
			}
		}
		// 录入系统模块
		err := tx.Model(&model.Module{}).CreateInBatches(moduleList, mysql.CalcMysqlBatchSize(moduleList[0])).Error
		if err != nil {
			logger.Error(errorx.MustWrap(err), "初始化系统模块失败！")
			return err
		}
		// 初始化类型表失败
		err = tx.Model(&model.TypeTree{}).CreateInBatches(typeTreeList, mysql.CalcMysqlBatchSize(typeTreeList[0])).Error
		if err != nil {
			logger.Error(errorx.MustWrap(err), "初始化类型表失败！")
			return err
		}
		// 初始化角色
		err = tx.Model(&model.Role{}).Create(&allRole).Error
		if err != nil {
			logger.Error(errorx.MustWrap(err), "初始化系统超级管理员失败！")
			return err
		}
		err = tx.Model(&model.Role{}).Create(&readWriteRole).Error
		if err != nil {
			logger.Error(errorx.MustWrap(err), "初始化系统读写角色失败！")
			return err
		}
		err = tx.Model(&model.Role{}).Create(&readRole).Error
		if err != nil {
			logger.Error(errorx.MustWrap(err), "初始化系统只读角色失败！")
			return err
		}

		//初始化角色权限
		for i := range moduleList {
			roleCapabilities = append(roleCapabilities, model.RoleCapabilities{
				RoleID:     *allRole.RecID,
				ModuleID:   *moduleList[i].RecID,
				ModuleName: moduleList[i].ModuleName,
				CanRead:    true,
				CanWrite:   true,
				CanDelete:  true,
			}, model.RoleCapabilities{
				RoleID:     *readWriteRole.RecID,
				ModuleID:   *moduleList[i].RecID,
				ModuleName: moduleList[i].ModuleName,
				CanRead:    true,
				CanWrite:   true,
				CanDelete:  false,
			},
				model.RoleCapabilities{
					RoleID:     *readRole.RecID,
					ModuleID:   *moduleList[i].RecID,
					ModuleName: moduleList[i].ModuleName,
					CanRead:    true,
					CanWrite:   false,
					CanDelete:  false,
				})
		}
		err = tx.Model(&model.RoleCapabilities{}).CreateInBatches(roleCapabilities, mysql.CalcMysqlBatchSize(roleCapabilities[0])).Error
		if err != nil {
			logger.Error(errorx.MustWrap(err), "初始化系统预定义角色权限失败！")
		}
		// 初始化部门
		err = tx.Model(&model.Department{}).CreateInBatches(departmentList, mysql.CalcMysqlBatchSize(departmentList[0])).Error
		if err != nil {
			logger.Error(errorx.MustWrap(err), "初始化部门列表失败！")
		}
		// 初始化超管
		err = tx.Model(&model.Staff{}).Create(&model.Staff{
			BasicModel:          model.BasicModel{},
			StaffName:           "超级管理员",
			StaffAlias:          &staffAlias,
			StaffEmail:          "645171033@qq.com",
			StaffPhone:          nil,
			StaffPassword:       "HSHROOT",
			StaffRoleID:         *allRole.RecID,
			StaffRoleName:       "root",
			StaffDepartmentID:   departmentList[0].RecID,
			StaffPosition:       &staffPosition,
			StaffDepartmentName: &departmentList[0].DepartmentName,
		}).Error
		if err != nil {
			logger.Error(errorx.MustWrap(err), "初始化系统超级管理员权限失败！")
		}

		return nil
	})

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

func Fate() {

}
