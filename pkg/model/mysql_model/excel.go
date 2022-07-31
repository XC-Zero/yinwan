package mysql_model

import _const "github.com/XC-Zero/yinwan/pkg/const"

type Excel struct {
	BasicModel
	ExcelName        *string `json:"excel_name" cn:""`
	ExcelURL         string
	CreatorID        *int
	CreatorName      *string
	ExcelExpiredTime int `cn:"过期时间(天)"`
	ExcelStatus      _const.ExcelStatus
}

func (e Excel) TableCnName() string {
	//TODO implement me
	panic("implement me")
}

func (e Excel) TableName() string {
	//TODO implement me
	panic("implement me")
}
