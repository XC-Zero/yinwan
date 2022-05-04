package client

import (
	"fmt"
	config2 "github.com/XC-Zero/yinwan/pkg/config"
	"gorm.io/gorm"

	"gorm.io/driver/mysql"
	"gorm.io/gorm/logger"
)

// InitMysqlGormV2 GormV2 returns a MySQL DB engine from configs
func InitMysqlGormV2(config config2.MysqlConfig) (*gorm.DB, error) {
	url := fmt.Sprintf("%s:%s@(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local&multiStatements=True",
		config.UserName,
		config.Password,
		config.Host,
		config.Port,
		config.DBName,
	)
	var log logger.Interface
	if config.LogMode == config2.None {
		log = logger.Default.LogMode(logger.Silent)
	} else {
		log = logger.Default.LogMode(logger.Info)
	}

	db, err := gorm.Open(mysql.Open(url), &gorm.Config{
		Logger: log,
	})
	// TODO 添加普罗米修斯!
	//db.Use()
	if err != nil {
		return nil, err
	}

	sqlDB, err := db.DB()
	if err != nil {
		return nil, err
	}
	sqlDB.SetMaxIdleConns(5)
	sqlDB.SetMaxOpenConns(20)
	return db, nil
}
