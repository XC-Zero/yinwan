package mysql

import (
	"fmt"
	"gorm.io/gorm/schema"
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

type SqlGeneration struct {
	sql string
}

func InitSql(model schema.Tabler, isCount bool) *SqlGeneration {
	str := "*"
	if true {
		str = "count(*)"
	}
	var s SqlGeneration
	s.sql = fmt.Sprintf("select %s from %s where 1=1 ", str, model.TableName())
	return &s
}

func (s *SqlGeneration) AddConditions(symbol string, conditions ...string) *SqlGeneration {
	length := len(conditions)
	for i := range conditions {
		if i%2 != 0 {
			continue
		}
		if i+1 >= length {
			break
		}
		if conditions[i+1] != "" {
			s.sql += fmt.Sprintf(" and %s %s '%s'", conditions[i], symbol, conditions[i+1])
		}
	}
	return s
}
func (s *SqlGeneration) AddOther(text string) *SqlGeneration {
	s.sql += text
	return s
}
func (s *SqlGeneration) Harvest() string {
	return s.sql
}

type BatchSqlGeneration struct {
	sqlList []SqlGeneration
}
