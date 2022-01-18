package model

// AccessControl 权限控制 --- 按模块
type AccessControl struct {
	accessControlID *int   `gorm:"primaryKey;type:int;autoIncrement"json:"id"`
	ModuleID        int    `gorm:"type:int"`
	ModuleName      string `gorm:"type:varchar(50)"`
}

func (a AccessControl) ID() int {
	return *a.accessControlID
}
