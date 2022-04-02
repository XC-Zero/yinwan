package system

import (
	"context"
	"fmt"
	"github.com/XC-Zero/yinwan/internal/config"
	"github.com/XC-Zero/yinwan/pkg/client"
	config2 "github.com/XC-Zero/yinwan/pkg/config"
	"github.com/XC-Zero/yinwan/pkg/utils/logger"
	"github.com/fwhezfwhez/errorx"
	"github.com/minio/minio-go/v7"
	"github.com/mozillazg/go-pinyin"
	"math/rand"
	"strings"
	"time"
)

// AddBookName 新建账套
//  其中 bookName 用于展示和Map的Key
//  可能为中文或其他
//  如需处理请转 []rune
func AddBookName(bookName string) (status bool) {

	var storage = config.CONFIG.StorageConfig
	var mysqlCfg = storage.MysqlConfig
	var minioCfg = storage.MinioConfig
	var mongoCfg = storage.MongoDBConfig
	var name = fmt.Sprintf("%s-%s-%d", Translate(bookName), time.Now().Format("2006-01-02"), rand.Int63n(1000000))
	mysqlCfg.DBName, minioCfg.Bucket, mongoCfg.DBName = name, name, name
	cfg := config2.BookConfig{
		BookName:      bookName,
		StorageName:   name,
		MysqlConfig:   mysqlCfg,
		MinioConfig:   minioCfg,
		MongoDBConfig: mongoCfg,
	}
	// 如果有失败的，删掉刚创建的
	defer func(status bool) {
		if !status {
			deleteMysqlDataBase(name)
			deleteBucket(name)
		}
	}(status)

	// 创建 mysql database
	err := createMysqlDataBase(name)
	if err != nil {
		return false
	}
	// MongoDB 的 database并不需要创建，直接用就完了

	//创建 minio bucket
	err = createBucket(name)
	if err != nil {
		return false
	}

	// 保存配置
	config.CONFIG.BookNameConfig = append(config.CONFIG.BookNameConfig, cfg)
	err = config.SaveConfig("book_name_config", config.CONFIG.BookNameConfig)
	if err != nil {
		logger.Error(errorx.MustWrap(err), fmt.Sprintf("新建账套时写入配置文件失败! "))
		return false
	}
	logger.Info(fmt.Sprintf("新建账套成功，账套名为 %s ", cfg.BookName))
	err = client.InitBookMap(config.CONFIG.BookNameConfig)
	if err != nil {
		logger.Error(errorx.MustWrap(err), fmt.Sprintf("初始化账套 (%s) 客户端失败! ", cfg.BookName))
	}
	return true
}

func UpdateBookName(oldBookName, NewBookName string) bool {
	bookName, ok := client.BookNameMap[oldBookName]
	if  !ok {
		return false
	}else {
		if _, exist := client.BookNameMap[NewBookName]; exist {
			return false
		}
		client.
	}


}

// Translate
//  将中文转驼峰拼音
func Translate(chineseSentence string) string {
	p, str := pinyin.NewArgs(), ""
	list := pinyin.Pinyin(chineseSentence, p)
	for i := range list {
		str += strings.ToLower(list[i][0])
	}
	return str
}

// StrFirstToUpper
//  首字母大写
//  minio 的 bucket 名不支持大写 只能包含 小写英文 数字和 - .
//  所以这个方法没啥用
func StrFirstToUpper(str string) string {
	list := strings.SplitN(str, "", 2)
	list[0] = strings.ToUpper(list[0])
	return strings.Join(list, "")
}

func createMysqlDataBase(name string) error {
	// 创建 Mysql Database
	createSQL := fmt.Sprintf(
		"CREATE DATABASE IF NOT EXISTS `%s` CHARACTER SET utf8mb4;",
		name,
	)
	err := client.MysqlClient.Exec(createSQL).Error
	if err != nil {
		logger.Error(errorx.MustWrap(err), fmt.Sprintf("新建账套时创建 Mysql Database 失败! "))
		return err
	}
	logger.Info(fmt.Sprintf("新建账套时创建 Mysql Database（%s） 成功! ", name))
	return nil
}

func deleteMysqlDataBase(name string) {

}

func createBucket(name string) error {
	// 创建 minio bucket
	exists, err := client.MinioClient.BucketExists(context.TODO(), name)
	if err != nil {
		logger.Error(errorx.MustWrap(err), fmt.Sprintf("新建账套时查询 Minio 桶失败! "))
		return err
	}
	if !exists {
		err = client.MinioClient.MakeBucket(context.TODO(), name, minio.MakeBucketOptions{
			Region:        "cn-north-1",
			ObjectLocking: false,
		})
		if err != nil {
			logger.Error(errorx.MustWrap(err), fmt.Sprintf("新建账套时创建 Minio 桶失败! "))
			return err
		}
	} else {
		logger.Waring(errorx.Empty(), fmt.Sprintf("新建账套时创建 Minio 桶已存在! "))
	}

	logger.Info(fmt.Sprintf("新建账套时创建 Minio 桶(%s)成功! ", name))
	return nil
}

func deleteBucket(name string) {
	exists, err := client.MinioClient.BucketExists(context.TODO(), name)
	if err != nil {
		logger.Error(errorx.MustWrap(err), fmt.Sprintf("新建账套时查询 Minio 桶失败! "))
		return
	}
	if exists {
		err = client.MinioClient.RemoveBucket(context.TODO(), name)
		if err != nil {
			return
		}
	}
}