package mysql_model

import _const "github.com/XC-Zero/yinwan/pkg/const"

type Excel struct {
	BasicModel
	ExcelName        string             `gorm:"type:varchar(50);not null;" form:"excel_name,omitempty" json:"excel_name,omitempty" `
	ExcelURL         string             `gorm:"type:varchar(500);not null;" form:"excel_url,omitempty" json:"excel_url,omitempty"`
	CreatorID        *int               `gorm:"type:int" form:"creator_id,omitempty" json:"creator_id,omitempty" `
	CreatorName      *string            `gorm:"type:varchar(50)" form:"creator_name,omitempty" json:"creator_name,omitempty" `
	ExcelExpiredTime int                `gorm:"type:int" form:"excel_expired_time,omitempty" json:"excel_expired_time,omitempty"  cn:"过期时间(天)"`
	ExcelStatus      _const.ExcelStatus `gorm:"type:int" form:"excel_status" json:"excel_status" bson:"excel_status"`
}

func (e Excel) TableCnName() string {
	return "Excel表格"
}

func (e Excel) TableName() string {
	return "excels"
}
