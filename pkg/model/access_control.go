package model

// Module 模块表
type Module struct {
	BasicModel
	ModuleName   string  `gorm:"type:varchar(50)"`
	ModuleRemark *string `gorm:"type:varchar(200)"`
}

// Role 职工角色表
type Role struct {
	BasicModel
	RoleName   string  `gorm:"type:varchar(50)"`
	RoleRemark *string `gorm:"type:varchar(200)"`
}

// RoleCapabilities 角色对各模块的权限关系表
type RoleCapabilities struct {
	RoelID    int `gorm:"type:int;index"`
	ModuleID  int `gorm:"type:int;index"`
	CanRead   bool
	CanWrite  bool
	CanDelete bool
}
