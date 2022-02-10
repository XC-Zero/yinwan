package other

import (
	"fmt"
	"github.com/XC-Zero/yinwan/internal/config"
	config2 "github.com/XC-Zero/yinwan/pkg/config"
	"github.com/XC-Zero/yinwan/pkg/utils/logger"
	"github.com/fwhezfwhez/errorx"
	"github.com/mozillazg/go-pinyin"
	"strings"
	"time"
)

// AddBookName 新建账套
// bookName 用于展示，可能为中文或其他，如需处理请转 []rune
func AddBookName(bookName string) bool {
	storage := config.CONFIG.StorageConfig
	mysqlCfg, minioCfg, mongoCfg := storage.MysqlConfig, storage.MinioConfig, storage.MongoDBConfig
	name := Translate(bookName) + "_" + time.Now().Format("2006_01_02")
	mysqlCfg.DBName, minioCfg.Bucket, mongoCfg.DBName = name, name, name
	cfg := config2.BookConfig{
		BookName:      bookName,
		MysqlConfig:   mysqlCfg,
		MinioConfig:   minioCfg,
		MongoDBConfig: mongoCfg,
	}

	// 保存配置
	config.CONFIG.BookNameConfig = append(config.CONFIG.BookNameConfig, cfg)
	err := config.SaveConfig("book_name_config", config.CONFIG.BookNameConfig)
	if err != nil {
		logger.Error(errorx.MustWrap(err), fmt.Sprintf("新建账套时写入配置文件失败! "))
		return false
	}
	logger.Info(fmt.Sprintf("新建账套成功，账套名为 %s ", cfg.BookName))
	return true
}

// Translate 将中文转驼峰拼音
func Translate(chineseSentence string) string {
	p, str := pinyin.NewArgs(), ""
	list := pinyin.Pinyin(chineseSentence, p)
	for i := range list {
		str += StrFirstToUpper(list[i][0])
	}
	return str
}

// StrFirstToUpper 首字母大写
func StrFirstToUpper(str string) string {
	temp := strings.Split(str, "")
	var upperStr string
	for y := 0; y < len(temp); y++ {
		vv := []rune(temp[y])
		if y != 0 {
			for i := 0; i < len(vv); i++ {
				if i == 0 {
					vv[i] -= 32
					upperStr += string(vv[i])
				} else {
					upperStr += string(vv[i])
				}
			}
		}
	}
	return temp[0] + upperStr
}
