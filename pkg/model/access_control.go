package model

// Module 模块表
type Module struct {
	BasicModel
	ModuleName   string  `gorm:"type:varchar(50)" json:"module_name"`
	ModuleRemark *string `gorm:"type:varchar(200)" json:"module_remark,omitempty"`
}

func (r Module) TableName() string {
	return "modules"
}

// Role 职工角色表
type Role struct {
	BasicModel
	RoleName   string  `gorm:"type:varchar(50)"  binding:"required,dive"`
	RoleRemark *string `gorm:"type:varchar(200)"  json:"role_remark,omitempty"`
}

func (r Role) TableName() string {
	return "roles"
}

// RoleCapabilities 角色对各模块的权限关系表
type RoleCapabilities struct {
	BasicModel
	RoleID     int    `gorm:"type:int;index" json:"role_id" binding:"required"`
	ModuleID   int    `gorm:"type:int;index" json:"module_id" binding:"required"`
	ModuleName string `gorm:"type:varchar(50)" json:"module_name" binding:"required"`
	CanRead    bool   `json:"can_read"`
	CanWrite   bool   `json:"can_write"`
	CanDelete  bool   `json:"can_delete"`
}

func (r RoleCapabilities) TableName() string {
	return "role_capabilities"
}
