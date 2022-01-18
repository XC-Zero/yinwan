package xlsx

import (
	"github.com/XC-Zero/yinwan/pkg/utils/errs"
	"os"
	"reflect"
	"strconv"

	//"github.com/XC-Zero/common/errs"
	"github.com/pkg/errors"
	"github.com/tealeg/xlsx"
)

// ReadSheetRow 从xlsx文件的制定Sheet中按行读取
// fun 是对每一行数据特殊处理的函数
func ReadSheetRow(path, sheetName string, fun func(row *xlsx.Row) (interface{}, error)) (dataList []interface{}, err error) {
	xlFile, err := xlsx.OpenFile(path)
	if err != nil {
		return nil, errors.New("Read xlsx file failed ! ")
	}
	sheet := xlFile.Sheet[sheetName]
	length := sheet.MaxRow
	errList := make([]error, 0)

	for i := 0; i < length; i++ {
		row := sheet.Row(i)
		data, err := fun(row)
		if err != nil {
			errList = append(errList, err)
			panic(err)
		}
		dataList = append(dataList, data)
	}
	return dataList, errs.ErrorListToError(errList)
}

// SaveXlsx 保存处理的文档，会覆盖源文档
func SaveXlsx(xlsxFile *xlsx.File, path string) error {
	err := DeleteFileIfExist(path)
	if err != nil {
		return err
	}
	err = xlsxFile.Save(path)
	if err != nil {
		return err
	}
	return nil
}

// AddDataToNewSheet 添加数据到新的sheet中
func AddDataToNewSheet(xlsxFile *xlsx.File, sheetName string, dataList []interface{}, fun func(data *interface{})) error {
	println("+++++++++ Start add data to sheet+++++++++")
	if dataList == nil {
		return errors.New("DataList is empty ! ")
	}
	var sheet *xlsx.Sheet
	sheet, err := xlsxFile.AddSheet(sheetName)
	if err != nil {
		return err
	}
	//AddColumnTitle(reflect.ValueOf(dataList).Field(0).Elem(), sheet)
	if len(dataList) == 0 {
		return nil
	}
	for i := range dataList {
		data := dataList[i]
		row := sheet.AddRow()
		if fun != nil {
			fun(&data)
		}
		TransferDataToRow(data, row)
	}
	return nil
}

// AddColumnTitle 通过tag添加首行
func AddColumnTitle(structModel interface{}, sheet *xlsx.Sheet) {
	objValue := reflect.ValueOf(structModel)
	row := sheet.AddRow()
	for i := 0; i < objValue.NumField(); i++ {
		fieldInfo := objValue.Type().Field(i)
		tag := fieldInfo.Tag
		name := tag.Get("json")
		cell := row.AddCell()
		cell.SetString(name)
	}
}

// TransferDataToRow 转换结构体类型数据成一行数据
func TransferDataToRow(data interface{}, row *xlsx.Row) {
	objValue := reflect.ValueOf(data)
	objType := reflect.TypeOf(data)
	for i := 0; i < objValue.NumField(); i++ {
		cell := row.AddCell()
		str := TransferInterfaceToString(objType.Field(i), objValue.Field(i).Interface())
		cell.SetString(str)
	}
}

// PathExists 文件是否存在
func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

// DeleteFileIfExist 删除文件
func DeleteFileIfExist(path string) error {
	flag, err := PathExists(path)
	if err != nil {
		return err
	}
	if flag {
		err := os.Remove(path)
		if err != nil {
			return err
		}
	}
	return nil
}

// TransferInterfaceToString 将interface转换为String类型
// 目前只支持 int,float,string,和上述的指针类型
func TransferInterfaceToString(structField reflect.StructField, data interface{}) string {
	var (
		float32Ptr *float32
		float64Ptr *float64
		intPtr     *int
		int8Ptr    *int8
		int16Ptr   *int16
		int32Ptr   *int32
		int64Ptr   *int64
		stringPtr  *string
	)

	switch structField.Type {
	case reflect.TypeOf(float64Ptr):
		d := data.(*float64)
		if d == nil {
			return ""
		} else {
			return strconv.FormatFloat(*d, 'f', 6, 64)
		}
	case reflect.TypeOf(float32Ptr):
		d := data.(*float32)
		if d == nil {
			return ""
		} else {
			return strconv.FormatFloat(float64(*d), 'f', 6, 32)
		}
	case reflect.TypeOf(float32(1)):
		return strconv.FormatFloat(float64(data.(float32)), 'f', 6, 32)
	case reflect.TypeOf(float64(1)):
		return strconv.FormatFloat(data.(float64), 'f', 6, 64)
	case reflect.TypeOf(intPtr):
		d := data.(*int)
		if d == nil {
			return ""
		} else {
			return strconv.Itoa(*d)
		}
	case reflect.TypeOf(int8Ptr):
		d := data.(*int)
		if d == nil {
			return ""
		} else {
			return strconv.Itoa(*d)
		}
	case reflect.TypeOf(int16Ptr):
		d := data.(*int)
		if d == nil {
			return ""
		} else {
			return strconv.Itoa(*d)
		}
	case reflect.TypeOf(int32Ptr):
		d := data.(*int)
		if d == nil {
			return ""
		} else {
			return strconv.Itoa(*d)
		}
	case reflect.TypeOf(int64Ptr):
		d := data.(*int)
		if d == nil {
			return ""
		} else {
			return strconv.Itoa(*d)
		}
	case reflect.TypeOf(1):
		return strconv.Itoa(data.(int))
	case reflect.TypeOf(int8(1)):
		return strconv.Itoa(int(data.(int8)))
	case reflect.TypeOf(int16(1)):
		return strconv.Itoa(int(data.(int16)))
	case reflect.TypeOf(int32(1)):
		return strconv.Itoa(int(data.(int32)))
	case reflect.TypeOf(int64(1)):
		return strconv.Itoa(int(data.(int64)))
	case reflect.TypeOf(stringPtr):
		d := data.(*string)
		if d == nil {
			return ""
		} else {
			return *d
		}

	case reflect.TypeOf(""):
		return data.(string)

	default:
		return ""
	}
}
