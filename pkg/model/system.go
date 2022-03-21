package model

// TypeTree 类型表
type TypeTree struct {
	BasicModel
	TypeName     string  `gorm:"type:varchar(50);not null;" json:"type_name"`
	ParentTypeID *int    `gorm:"int;index" json:"parent_type_id,omitempty"`
	Remark       *string `gorm:"type:varchar(200)" json:"remark,omitempty"`
}
