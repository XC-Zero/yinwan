package mysql_model

import (
	"fmt"
	"github.com/XC-Zero/yinwan/pkg/client"
	"github.com/XC-Zero/yinwan/pkg/model/mysql_model/common"
	"github.com/pkg/errors"
	"gorm.io/gorm/schema"
	"reflect"
	"strings"
)

// TypeTree 类型表
type TypeTree struct {
	common.BasicModel
	TypeName     string  `gorm:"type:varchar(50);not null;" json:"type_name"`
	ParentTypeID *int    `gorm:"int;index" json:"parent_type_id,omitempty"`
	Remark       *string `gorm:"type:varchar(200)" json:"remark,omitempty"`
}

func (m TypeTree) TableName() string {
	return "type_tree"
}
func (m TypeTree) TableCnName() string {
	return "类型"
}

// QRCodeMapping 二维码映射表
type QRCodeMapping struct {
	common.TimeOnlyModel
	QRCodeMD5  string `gorm:"type:varchar(100);index" json:"qr_code_md5"`
	contentSql string `gorm:"type:varchar(500)" `
}

func (q *QRCodeMapping) GenerateSql(table schema.Tabler) error {
	var id *int
	baseSql, asSql := "select %s from %s where id = %d ;", ""
	objType, objVal := reflect.TypeOf(table), reflect.ValueOf(table)

	for i := 0; i < objType.NumField(); i++ {
		if i == 0 {
			id = reflect.ValueOf(objVal.Field(i).Interface()).Field(0).Interface().(*int)
			asSql += " id AS 编号,"
			continue
		}
		en := strings.ReplaceAll(objType.Field(i).Tag.Get("json"), ",omitempty", "")
		cn := objType.Field(i).Tag.Get("cn")

		if en == "" || cn == "" {
			continue
		}
		asSql += en + " AS '" + cn + "' ,"

	}
	if id == nil {
		return errors.New("Primary key is nil!")
	}

	q.contentSql = strings.ReplaceAll(fmt.Sprintf(baseSql, asSql, table.TableName(), *id), ", from", "from")
	return nil
}
func (q *QRCodeMapping) ExecSql(res map[string]interface{}) (err error) {
	return client.MysqlClient.Raw(q.contentSql).Scan(&res).Error
}

func (q QRCodeMapping) TableName() string {
	return "qrcode_mappings"
}
