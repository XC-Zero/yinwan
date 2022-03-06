package mysql

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strings"
	"time"

	"gorm.io/gorm/schema"
)

/*
原因：
	1。由于从 mysql 到kafka 的 json 数据映射，并不能很好的兼容 go 的反序列化，例如 boolean 类型的 is_company 在 json 数据中为：
	{"is_company": 1}; 该数据无法直接调用  	json.Unmarshal()
	2。直接调用 json.Unmarshal() 还需要在 pevc 的 model 中添加 json tag, 维护起来比较麻烦
实现：
	1。定义一个 map[string]interface{} 用于存储 kafka 中的 json 数据
	2。使用 pevc 中的 gorm tag 作为 key 值，在上述 的 map 中获取值
	3。使用反射给 pevc 的字段赋值
缺陷：
	1。自己实现的逻辑，可能没有考虑到所有的情况；不如目前只实现了 bool ,int,float,string,date time ,以及其指针的反序列化，如涉及其他类型的数据，则跳过
建议：如果能寻找到 kafka connector convert 兼容 go 的序列化，反序列方法的，此方法可以废弃
*/

// Deserialize ...
func Deserialize(sourceData []byte, destStruct interface{}) error {
	// 使用 map 存储 kafka 中的 json 数据
	var mapData map[string]interface{}
	err := json.Unmarshal(sourceData, &mapData)
	if err != nil {
		fmt.Println(err)
		return err
	}
	// 获取对象实例,必须是指针类型
	v := reflect.ValueOf(destStruct).Elem()
	fields := getModelFields(v)
	for _, field := range fields {
		gTag := field.Tag.Get("gorm")
		column := getColumnNameFromTag(gTag)
		if column == "" {
			column = schema.NamingStrategy{}.ColumnName("", field.Name)
		}
		value := mapData[column]
		if value == nil {
			continue
		}

		// ======================================= 反序列化 ===============================================
		// 字段的变量名
		fieldName := field.Name
		// 如果是时间类型: created_at
		if field.Type == reflect.TypeOf(time.Time{}) {
			if isTimestampType(gTag) {
				t, err := formatTime(value.(string))
				if err != nil {
					return err
				}
				v.FieldByName(fieldName).Set(reflect.ValueOf(*t))
			} else {
				v.FieldByName(fieldName).Set(reflect.ValueOf(time.Unix(int64(value.(float64))/1000, 0)))
			}
			continue
		}
		// 指针类型的  updated_at, deleted_at
		if field.Type == reflect.TypeOf(&time.Time{}) {
			v.FieldByName(fieldName).Set(reflect.New(v.FieldByName(fieldName).Type().Elem()))
			if isTimestampType(gTag) {
				t, err := formatTime(value.(string))
				if err != nil {
					return err
				}
				v.FieldByName(fieldName).Elem().Set(reflect.ValueOf(*t))
			} else {
				v.FieldByName(fieldName).Elem().Set(reflect.ValueOf(time.Unix(int64(value.(float64))/1000, 0)))
			}
			continue
		}
		// 根据字段的类型，进行赋值
		switch field.Type.Kind() {
		// 这里如何判断 bool 为 true,false;需要根据具体的情况进行判断
		case reflect.Bool:
			if value.(float64) == 1 {
				v.FieldByName(fieldName).SetBool(true)
			}
		case reflect.String:
			v.FieldByName(fieldName).SetString(value.(string))
		case reflect.Int, reflect.Int8, reflect.Int32, reflect.Int64:
			v.FieldByName(fieldName).SetInt(int64(value.(float64)))
		case reflect.Float32, reflect.Float64:
			v.FieldByName(fieldName).SetFloat(value.(float64))
			// 指针类型
		case reflect.Ptr:
			v.FieldByName(fieldName).Set(reflect.New(v.FieldByName(fieldName).Type().Elem()))
			switch field.Type.Elem().Kind() {
			case reflect.Bool:
				if value.(float64) == 1 {
					v.FieldByName(fieldName).Elem().SetBool(true)
				}
			case reflect.String:
				v.FieldByName(fieldName).Elem().SetString(value.(string))
			case reflect.Int, reflect.Int8, reflect.Int32, reflect.Int64:
				v.FieldByName(fieldName).Elem().SetInt(int64(value.(float64)))
			case reflect.Float32, reflect.Float64:
				v.FieldByName(fieldName).Elem().SetFloat(value.(float64))
			}
		}
		// ======================================= 反序列化 ===============================================
	}
	return nil
}

func getModelFields(v reflect.Value) []reflect.StructField {
	sf := make([]reflect.StructField, 0)
	if v.Type().Kind() != reflect.Struct {
		panic("CDC Connector Deserialization Please pass in the structure ！！！！")
	}

	for i := 0; i < v.NumField(); i++ {
		field := v.Type().Field(i)
		if field.Type.Kind() == reflect.Struct && IsInnerStruct(field) {
			innerSF := getModelFields(v.Field(i))
			sf = append(sf, innerSF...)
		} else {
			sf = append(sf, field)
		}
	}
	return sf
}

func IsInnerStruct(structField reflect.StructField) bool {
	if structField.Type == reflect.TypeOf(time.Time{}) || structField.Type == reflect.TypeOf(&time.Time{}) {
		return false
	}
	if strings.Contains(structField.Tag.Get("gorm"), "foreignKey") {
		return false
	}

	return true
}

func getColumnNameFromTag(tag string) string {
	// 按分号分割
	tags := strings.Split(tag, ";")
	for _, value := range tags {
		v := strings.Split(value, ":")
		k := strings.TrimSpace(strings.ToUpper(v[0]))
		if len(v) >= 2 {
			if k == "COLUMN" {
				return strings.Join(v[1:], ":")
			}
		}
	}
	return ""
}

func getColumnTypeFromTag(tag string) string {
	// 按分号分割
	tags := strings.Split(tag, ";")
	for _, value := range tags {
		v := strings.Split(value, ":")
		k := strings.TrimSpace(strings.ToUpper(v[0]))
		if len(v) >= 2 && k == "TYPE" {
			return strings.TrimSpace(strings.ToUpper(v[1]))
		}
	}
	return ""
}

func isTimestampType(gTag string) bool {
	typ := getColumnTypeFromTag(gTag)
	if typ == "TIMESTAMP" {
		return true
	}
	return false
}

// formatTime ...
func formatTime(date string) (*time.Time, error) {
	layout := "2006-01-02T15:04:05Z"

	t, err := time.ParseInLocation(layout, date, time.FixedZone("BJ", 8*3600))
	if err != nil {
		return nil, err
	}
	return &t, nil
}
