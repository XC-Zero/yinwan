package mysql

import (
	"reflect"
	"strings"
)

func CalcMysqlBatchSize(data interface{}) int {
	count := countStructFields(reflect.ValueOf(data))
	return 60000 / count
}

func countStructFields(v reflect.Value) int {
	var count int
	for i := 0; i < v.NumField(); i++ {
		field := v.Type().Field(i)
		if field.Type.Kind() == reflect.Struct && IsInnerStruct(field) {
			num := countStructFields(v.Field(i))
			count += num
		} else if !strings.Contains(field.Tag.Get("gorm"), "foreignKey") {
			count += 1
		}
	}
	return count
}
