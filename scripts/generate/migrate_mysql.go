package main

import (
	"context"
	"fmt"
	"github.com/XC-Zero/yinwan/pkg/model/common"
	"github.com/XC-Zero/yinwan/pkg/model/mongo_model"
	"github.com/XC-Zero/yinwan/pkg/model/mysql_model"
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
var moduleList = mysql_model.GetModuleList()

var roleCapabilities []mysql_model.RoleCapabilities
var departmentList []mysql_model.Department
var allRole mysql_model.Role
var readRole mysql_model.Role
var readWriteRole mysql_model.Role
var typeTreeList []mysql_model.TypeTree

func init() {
	systemMysqlMigrateList = []interface{}{
		&mysql_model.Role{},
		&mysql_model.RoleCapabilities{},
		&mysql_model.Module{},
		&mysql_model.Department{},
		&mysql_model.Staff{},
		&mysql_model.Commodity{},
		&mysql_model.CommodityHistoricalCost{},
		&mysql_model.CommodityBatch{},
		&mysql_model.Material{},
		&mysql_model.MaterialBatch{},
		&mongo_model.Transaction{},
		&mysql_model.Customer{},
		&mysql_model.Provider{},
		&mysql_model.ManipulationLog{},
		&mysql_model.TypeTree{},
		&mysql_model.Payable{},
		&mysql_model.Receivable{},
		&mongo_model.Purchase{},
		&mongo_model.FinanceCredential{},
		&mongo_model.FinanceCredentialEvent{},
		&mongo_model.EventItem{},
	}
	bookNameMysqlMigrateList = []interface{}{
		&mysql_model.Payable{},
		&mysql_model.Receivable{},
		&mongo_model.Purchase{},
		&mongo_model.FinanceCredential{},
		&mongo_model.FinanceCredentialEvent{},
		&mongo_model.EventItem{},
	}
	allRole = mysql_model.Role{
		RoleName: "root",
	}
	readRole = mysql_model.Role{RoleName: "只读账号"}
	readWriteRole = mysql_model.Role{RoleName: "读写账号"}
	roleCapabilities = []mysql_model.RoleCapabilities{}
	departmentManagerName, departmentManagerID, finAddr := "超级管理员", 1, "2楼205"
	departmentList = []mysql_model.Department{
		{
			DepartmentName:        "技术部",
			DepartmentManagerID:   &departmentManagerID,
			DepartmentManagerName: &departmentManagerName,
		}, {
			BasicModel:            common.BasicModel{},
			DepartmentName:        "财务部",
			DepartmentLocation:    &finAddr,
			DepartmentManagerID:   &departmentManagerID,
			DepartmentManagerName: &departmentManagerName,
		}, {
			DepartmentName:        "生产车间",
			DepartmentManagerID:   &departmentManagerID,
			DepartmentManagerName: &departmentManagerName,
		}, {
			DepartmentName:        "仓库",
			DepartmentManagerID:   &departmentManagerID,
			DepartmentManagerName: &departmentManagerName,
		}, {
			DepartmentName:        "销售部",
			DepartmentManagerID:   &departmentManagerID,
			DepartmentManagerName: &departmentManagerName,
		}, {
			DepartmentName:        "人事部",
			DepartmentManagerID:   &departmentManagerID,
			DepartmentManagerName: &departmentManagerName,
		}, {
			DepartmentName:        "行政部",
			DepartmentManagerID:   &departmentManagerID,
			DepartmentManagerName: &departmentManagerName,
		}}
	typeTreeList = []mysql_model.TypeTree{
		{common.BasicModel{}, "固定资产", nil, nil},
		{common.BasicModel{}, "产成品", nil, nil},
		{common.BasicModel{}, "原材料", nil, nil},
		{common.BasicModel{}, "周转材料", nil, nil},
		{common.BasicModel{}, "低值易耗品", nil, nil},
		{common.BasicModel{}, "其他类型 ", nil, nil},
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
		err := tx.Model(&mysql_model.Module{}).CreateInBatches(moduleList, mysql.CalcMysqlBatchSize(moduleList[0])).Error
		if err != nil {
			logger.Error(errorx.MustWrap(err), "初始化系统模块失败！")
			return err
		}
		// 初始化类型表失败
		err = tx.Model(&mysql_model.TypeTree{}).CreateInBatches(typeTreeList, mysql.CalcMysqlBatchSize(typeTreeList[0])).Error
		if err != nil {
			logger.Error(errorx.MustWrap(err), "初始化类型表失败！")
			return err
		}
		// 初始化角色
		err = tx.Model(&mysql_model.Role{}).Create(&allRole).Error
		if err != nil {
			logger.Error(errorx.MustWrap(err), "初始化系统超级管理员失败！")
			return err
		}
		err = tx.Model(&mysql_model.Role{}).Create(&readWriteRole).Error
		if err != nil {
			logger.Error(errorx.MustWrap(err), "初始化系统读写角色失败！")
			return err
		}
		err = tx.Model(&mysql_model.Role{}).Create(&readRole).Error
		if err != nil {
			logger.Error(errorx.MustWrap(err), "初始化系统只读角色失败！")
			return err
		}

		//初始化角色权限
		for i := range moduleList {
			roleCapabilities = append(roleCapabilities, mysql_model.RoleCapabilities{
				RoleID:     *allRole.RecID,
				ModuleID:   *moduleList[i].RecID,
				ModuleName: moduleList[i].ModuleName,
				CanRead:    true,
				CanWrite:   true,
				CanDelete:  true,
			}, mysql_model.RoleCapabilities{
				RoleID:     *readWriteRole.RecID,
				ModuleID:   *moduleList[i].RecID,
				ModuleName: moduleList[i].ModuleName,
				CanRead:    true,
				CanWrite:   true,
				CanDelete:  false,
			},
				mysql_model.RoleCapabilities{
					RoleID:     *readRole.RecID,
					ModuleID:   *moduleList[i].RecID,
					ModuleName: moduleList[i].ModuleName,
					CanRead:    true,
					CanWrite:   false,
					CanDelete:  false,
				})
		}
		err = tx.Model(&mysql_model.RoleCapabilities{}).CreateInBatches(roleCapabilities, mysql.CalcMysqlBatchSize(roleCapabilities[0])).Error
		if err != nil {
			logger.Error(errorx.MustWrap(err), "初始化系统预定义角色权限失败！")
		}
		// 初始化部门
		err = tx.Model(&mysql_model.Department{}).CreateInBatches(departmentList, mysql.CalcMysqlBatchSize(departmentList[0])).Error
		if err != nil {
			logger.Error(errorx.MustWrap(err), "初始化部门列表失败！")
		}
		// 初始化超管
		err = tx.Model(&mysql_model.Staff{}).Create(&mysql_model.Staff{
			BasicModel:          common.BasicModel{},
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
