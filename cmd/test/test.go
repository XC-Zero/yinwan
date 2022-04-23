package main

import (
	"github.com/XC-Zero/yinwan/pkg/model/common"
	"github.com/XC-Zero/yinwan/pkg/model/mysql_model"
	"github.com/XC-Zero/yinwan/pkg/utils/convert"
)

func main() {
	var dataList []interface{}
	dataList = append(dataList, mysql_model.Staff{
		BasicModel:          common.BasicModel{},
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
	cv, err := convert.SliceConvert(dataList, []interface{}{})
	if err != nil {
		panic(err)
		return
	}
	convert.OneToManyCombinationConvert(dataList, cv.([]interface{}), "")
}
