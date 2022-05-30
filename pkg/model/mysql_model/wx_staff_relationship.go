package mysql_model

type WxStaffRelationship struct {
	BasicModel
	StaffID        int     `json:"staff_id" form:"staff_id" gorm:"type:int;index;not null"`
	OpenID         string  `json:"open_id" form:"open_id" gorm:"type:varchar(200);not null;index"`
	WxBindingPhone *string `json:"wx_binding_phone" form:"wx_binding_phone" gorm:"type:varchar(20)"`
}

func (w WxStaffRelationship) TableCnName() string {
	return "微信与员工关联表"
}

func (w WxStaffRelationship) TableName() string {
	return "wx_staff_relationships"
}
