package model

// Module 模块表
type Module struct {
	BasicModel
	ModuleName   string  `gorm:"type:varchar(50)" json:"module_name"`
	ModuleRemark *string `gorm:"type:varchar(200)" json:"module_remark,omitempty"`
}

func (r Module) TableCnName() string {
	return "系统模块"
}

func (r Module) TableName() string {
	return "modules"
}

// Role 职工角色表
type Role struct {
	BasicModel
	RoleName   string  `gorm:"type:varchar(50)" form:"role_name" json:"role_name"  binding:"required"`
	RoleRemark *string `gorm:"type:varchar(200)" form:"role_remark" json:"role_remark,omitempty"`
}

func (r Role) TableCnName() string {
	return "角色"
}

func (r Role) TableName() string {
	return "roles"
}

// RoleCapabilities 角色对各模块的权限关系表
type RoleCapabilities struct {
	BasicModel
	RoleID     int    `gorm:"type:int;index" form:"role_id" json:"role_id"`
	ModuleID   int    `gorm:"type:int;index" form:"module_id" json:"module_id" binding:"required"`
	ModuleName string `gorm:"type:varchar(50)" form:"module_name" json:"module_name" binding:"required"`
	CanRead    bool   `form:"can_read" json:"can_read"`
	CanWrite   bool   `form:"can_write" json:"can_write"`
	CanDelete  bool   `form:"can_delete" json:"can_delete"`
}

func (r RoleCapabilities) TableCnName() string {
	return "角色权限"
}

func (r RoleCapabilities) TableName() string {
	return "role_capabilities"
}
