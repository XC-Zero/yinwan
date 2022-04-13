package main

import (
	"github.com/XC-Zero/yinwan/pkg/model"
	"github.com/XC-Zero/yinwan/pkg/utils/tools"
	"log"
	"reflect"
)

func main() {
	var dataList []model.Staff
	dataList = append(dataList, model.Staff{
		BasicModel:          model.BasicModel{},
		StaffName:           "",
		StaffAlias:          nil,
		StaffEmail:          "",
		StaffPhone:          nil,
		StaffPassword:       "",
		StaffPosition:       nil,
		StaffDepartmentID:   nil,
		StaffDepartmentName: nil,
		StaffRoleID:         0,
		StaffRoleName:       "",
	})
	convert, err := tools.SliceConvert(dataList, []interface{}{})
	if err != nil {
		return
	}
	log.Println(reflect.TypeOf(convert))
	i := convert.([]interface{})
	log.Println(i)
}
