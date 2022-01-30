package condition

import (
	"github.com/pkg/errors"
	"reflect"
)

type OperatorSymbol string

const (
	OR  = "or"
	AND = "and"
)

// ModelToMysqlCondition 模型转条件
// 只能输入一个结构体的实例化对象
// 含有多个属性时全是 or
func ModelToMysqlCondition(model interface{}, operatorSymbol OperatorSymbol) (conditions []interface{}, err error) {
	objType, objVal, conditionStr := reflect.TypeOf(model), reflect.ValueOf(model), ""
	if objType.Kind() != reflect.Struct {
		return nil, errors.New("传入的参数类型不是结构体！")
	}

	var conditionList []string
	var valList []interface{}
	// 跳过 BasicModel
	for i := 1; i < objVal.NumField(); i++ {
		if !objVal.Field(i).IsZero() {
			conditionList = append(conditionList, objType.Field(i).Tag.Get("json")+" = ? ")
			valList = append(valList, objVal.Field(i).Interface())
		}
	}
	for i := 0; i < len(conditionList); i++ {
		conditionStr += conditionList[i]
		if i != len(conditionList)-1 {
			conditionStr += " " + string(operatorSymbol) + " "
		}
	}
	// todo valList 也行？ 不需要解开吗？？？
	//conditions = append(conditions, conditionStr, valList)

	conditions = append(conditions, conditionStr)
	for i := range valList {
		conditions = append(conditions, valList[i])
	}
	return
}
