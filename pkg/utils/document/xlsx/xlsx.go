package xlsx

import (
	"github.com/XC-Zero/yinwan/pkg/utils/errs"
	//"github.com/XC-Zero/common/errs"
	"github.com/pkg/errors"
	"github.com/tealeg/xlsx"
)

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
