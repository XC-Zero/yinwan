package main

import (
	"github.com/XC-Zero/yinwan/pkg/model"
	"github.com/XC-Zero/yinwan/pkg/utils/condition"
	"log"
)

func main() {

	toCondition, err := condition.ModelToMysqlCondition(
		model.Staff{
			BasicModel:    model.BasicModel{},
			StaffName:     "123",
			StaffAlias:    nil,
			StaffEmail:    "645171033@qq.com",
			StaffPhone:    nil,
			StaffPassword: "",
			StaffRoleID:   0,
			StaffRoleName: "",
		}, "or")

	if err != nil {
		panic(err)
	}

	log.Printf("%+v ", toCondition)
}
