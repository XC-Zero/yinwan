package main

import (
	"github.com/XC-Zero/yinwan/pkg/utils/math_plus"
	//"github.com/devfeel/mapper"
	"github.com/davecgh/go-spew/spew"
	"github.com/devfeel/mapper"
	"github.com/pkg/errors"
	"github.com/xuri/excelize/v2"
	"mime/multipart"
)

//
//func HarvestExcelFromMultiPart(file *multipart.FileHeader) *excelize.File {
//	f := excelize.NewFile()
//	f.
//}
type a struct {
	B struct {
		C struct {
			D string
		}
	}
}

//func tets() {
//	excel := excelize.NewFile()
//	excel.MergeCell()
//}
func StructToMapSlice() {
	var tabler a
	var m = make(map[string]interface{}, 0)
	err := mapper.Mapper(&tabler, &m)
	if err != nil {
		panic(errors.WithStack(err))
	}
	spew.Dump(m)
	//sort.
}

func NewFile() {

}

func OpenExcelFromMultiPart(multi *multipart.FileHeader) (*excelize.File, error) {
	f, err := multi.Open()
	if err != nil {
		return nil, err
	}
	defer f.Close()
	excel, err := excelize.OpenReader(f)
	if err != nil {
		return nil, err
	}
	return excel, nil
}

var mapping map[int]string

func init() {
	mapping = make(map[int]string, 26)
	for i := 1; i < 27; i++ {
		mapping[i] = string(('A') + int32(i-1))
	}
}

func IndexToExcelCol(index int) string {
	return math_plus.TenToAnyWithMapping(index+1, 26, mapping)
}
