package excel

import (
	"encoding/json"
	"github.com/xuri/excelize/v2"
	"reflect"
	"sync"
)

func A() {
	excelize.NewFile()
}

// JsonListToMapData todo 一条条JSON 转 map
func JsonListToMapData(jsonList []string) {
	var wg sync.WaitGroup
	var mutex sync.Mutex
	num, length, dataMap := 50, len(jsonList), make(map[string]interface{}, 0)
	if length < 50 {
		num = length
	}
	wg.Add(num)
	for i := range jsonList {
		var errorList []error
		go func(index int) {
			tempMap := map[string]interface{}{}
			err := json.Unmarshal([]byte(jsonList[i]), &tempMap)
			if err != nil {
				mutex.Lock()
				errorList = append(errorList, err)
				mutex.Unlock()
			} else {
				for key, value := range tempMap {
					if data, ok := dataMap[key]; ok {
						tempType, tempVal := reflect.TypeOf(value), reflect.ValueOf(value)
						if tempType.Kind() == reflect.Slice {
							dataMap[key] = append(tempVal.Interface().([]interface{}), data)
						}
					} else {
						dataMap[key] = value
					}
				}
			}
			wg.Done()
		}(i)
		wg.Wait()

	}
}

//
func MapToSheet(data map[string]interface{}) {

}
