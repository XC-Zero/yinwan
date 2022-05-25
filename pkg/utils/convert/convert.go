package convert

import (
	"encoding/json"
	"fmt"
	"github.com/pkg/errors"
	"log"
	"reflect"
	"strconv"
)

// SliceConvert 数组类型转换
//  此处用了大量反射，高并发情况下将导致性能瓶颈
func SliceConvert(slice interface{}, newSlice interface{}) (interface{}, error) {
	ot, nt := reflect.TypeOf(slice), reflect.TypeOf(newSlice)
	ote, nte := ot.Elem(), nt.Elem()
	log.Println(ot, ote, ote.Name())
	if ot.Kind() != reflect.Slice {
		return nil, errors.New(fmt.Sprintf("Slice called with non-slice value of type %T", ot))
	}
	if nt.Kind() != reflect.Slice {
		return nil, errors.New(fmt.Sprintf("Slice called with non-slice type of type %T", nt))
	}

	var v = reflect.ValueOf(slice)
	var l = v.Len()
	var dv = reflect.MakeSlice(nt, 0, l)

	if !ote.ConvertibleTo(nte) {
		return nil, errors.New(fmt.Sprintf("Type %T can not convert to type %T", ote, nte))
	}
	for i := 0; i < l; i++ {
		dv = reflect.Append(dv, v.Index(i).Convert(nte))
	}

	return dv.Interface(), nil
}

// OneToManyCombinationConvert 一对多合并
func OneToManyCombinationConvert(one, many []interface{}, combColumnName string) []interface{} {
	var res []interface{}
	if len(one) == 0 || len(many) == 0 {
		return nil
	}
	//oneKey := string_plus.ToSnakeString(reflect.TypeOf(reflect.ValueOf(one).Index(0).Elem().Interface()).Name())
	//manyKey := string_plus.ToSnakeString(reflect.TypeOf(reflect.ValueOf(many).Index(0).Elem().Interface()).Name())
	//var tempMap = make(map[interface{}]interface{}, len(one))
	//for i := range many {
	//	many[i]
	//}
	return res

}

func GetInterfaceToString(value interface{}) string {
	// interface 转 string
	var key string
	if value == nil {
		return key
	}

	switch value.(type) {
	case float64:
		ft := value.(float64)
		key = strconv.FormatFloat(ft, 'f', -1, 64)
	case float32:
		ft := value.(float32)
		key = strconv.FormatFloat(float64(ft), 'f', -1, 64)
	case int:
		it := value.(int)
		key = strconv.Itoa(it)
	case uint:
		it := value.(uint)
		key = strconv.Itoa(int(it))
	case int8:
		it := value.(int8)
		key = strconv.Itoa(int(it))
	case uint8:
		it := value.(uint8)
		key = strconv.Itoa(int(it))
	case int16:
		it := value.(int16)
		key = strconv.Itoa(int(it))
	case uint16:
		it := value.(uint16)
		key = strconv.Itoa(int(it))
	case int32:
		it := value.(int32)
		key = strconv.Itoa(int(it))
	case uint32:
		it := value.(uint32)
		key = strconv.Itoa(int(it))
	case int64:
		it := value.(int64)
		key = strconv.FormatInt(it, 10)
	case uint64:
		it := value.(uint64)
		key = strconv.FormatUint(it, 10)
	case string:
		key = value.(string)
	case []byte:
		key = string(value.([]byte))
	default:
		newValue, _ := json.Marshal(value)
		key = string(newValue)
	}

	return key
}
