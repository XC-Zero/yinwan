package model

// AccessControl 权限控制 --- 按模块
type AccessControl struct {
	AccessControlID *int   `gorm:"primaryKey;type:int;autoIncrement" json:"id"`
	ModuleID        int    `gorm:"type:int"`
	ModuleName      string `gorm:"type:varchar(50)"`
}
