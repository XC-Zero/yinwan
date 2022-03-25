package mysql

import (
	"fmt"
	"gorm.io/gorm/schema"
	"reflect"
	"strings"
)

type OperatorSymbol string

const (
	EQUAL OperatorSymbol = "="
	LIKE  OperatorSymbol = "like"
	IN    OperatorSymbol = "in"
)

type Selected string

const (
	COUNT        Selected = "count(*)"
	ALL          Selected = "*"
	DISTINCT_ALL Selected = " DISTINCT * "
)

const (
	BASIC_MODEL_CREATED_AT  = "created_at"
	BASIC_MODEL_UPDATED_AT  = "updated_at"
	BASIC_MODEL_DELETED_AT  = "deleted_at"
	BASIC_MODEL_PRIMARY_KEY = "id"
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

// sqlGeneration sql生成器
type sqlGeneration struct {
	sql string
}

// InitSqlGeneration 初始化
func InitSqlGeneration(model schema.Tabler, selected Selected) *sqlGeneration {

	var s sqlGeneration
	s.sql = fmt.Sprintf("select %s from %s where 1=1 ", selected, model.TableName())
	return &s
}

func (s *sqlGeneration) AddConditions(symbol OperatorSymbol, conditions ...string) *sqlGeneration {
	length := len(conditions)
	for i := range conditions {
		if i%2 != 0 {
			continue
		}
		if i+1 >= length {
			break
		}
		if conditions[i+1] != "" {
			tempSql := ""
			switch symbol {
			case LIKE:
				tempSql = fmt.Sprintf(" and %s %s '%%%s%%'", conditions[i], symbol, conditions[i+1])
			case IN:
				tempSql = fmt.Sprintf(" and %s %s (%s)", conditions[i], symbol, conditions[i+1])
			case EQUAL:
				tempSql = fmt.Sprintf(" and %s %s '%s'", conditions[i], symbol, conditions[i+1])

			}
			s.sql += tempSql
		}
	}
	return s
}

func (s *sqlGeneration) AddGroupBy(columnName string) *sqlGeneration {
	s.sql += fmt.Sprintf(" group by %s ", columnName)
	return s
}
func (s *sqlGeneration) AddSuffixOther(text string) *sqlGeneration {
	s.sql += text
	return s
}

func (s *sqlGeneration) HarvestSql() string {
	return s.sql
}

// batchSqlGeneration 批量sql
type batchSqlGeneration struct {
	sqlMap map[string]*sqlGeneration
}

func InitBatchSqlGeneration() *batchSqlGeneration {
	return &batchSqlGeneration{
		make(map[string]*sqlGeneration, 0),
	}
}

func (b *batchSqlGeneration) AddSqlGeneration(name string, s *sqlGeneration) *batchSqlGeneration {
	b.sqlMap[name] = s
	return b
}

func (b *batchSqlGeneration) AddConditions(symbol OperatorSymbol, conditions ...string) *batchSqlGeneration {
	for i := range b.sqlMap {
		b.sqlMap[i].AddConditions(symbol, conditions...)
	}
	return b
}

func (b *batchSqlGeneration) AddSuffixOther(text string) *batchSqlGeneration {
	for i := range b.sqlMap {
		b.sqlMap[i].AddSuffixOther(text)
	}
	return b
}

func (b *batchSqlGeneration) HarvestSql(name string) string {
	if v, ok := b.sqlMap[name]; ok {
		return v.HarvestSql()
	}
	return ""
}

func (b *batchSqlGeneration) Harvest(name string) *sqlGeneration {
	if v, ok := b.sqlMap[name]; ok {
		return v
	}
	return nil
}

func (b *batchSqlGeneration) HarvestAll() []string {
	var sqlList []string
	for key := range b.sqlMap {
		sqlList = append(sqlList, b.sqlMap[key].HarvestSql())
	}
	return sqlList
}
