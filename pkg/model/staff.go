package model

import (
	"strconv"
	"strings"
)

const SPLIT_SYMBOL = "|"

// Staff 职工表
type Staff struct {
	BasicModel
	StaffName     string  `gorm:"type:varchar(50);not null;"`
	StaffAlias    *string `gorm:"type:varchar(50)"` // 职工别名
	StaffEmail    *string `gorm:"type:varchar(50)"`
	StaffPhone    *string `gorm:"type:varchar(50)"`
	StaffPassword string  `gorm:"type:varchar(20)"`
	StaffRoleID   int     `gorm:"type:int"`
}

// Role 职工角色表
type Role struct {
	BasicModel
	RoleName    string `gorm:"type:varchar(50)"`
	RoleContent string `gorm:"type:varchar(500)"`
}

func (r *Role) SetRoleContent(controlList []AccessControl) {
	content := ""
	for i := 0; i < len(controlList)-1; i++ {
		ac := controlList[i]
		if ac.RecID != nil {
			content += strconv.Itoa(*controlList[i].RecID) + SPLIT_SYMBOL
		}
	}
	r.RoleContent = content
}

func (r Role) GetRoleAccessControlList() (controlList []AccessControl, err error) {
	arr := strings.Split(r.RoleContent, SPLIT_SYMBOL)
	if len(arr) == 0 {
		return nil, nil
	}
	// todo  根据ID从数据库找
	return nil, nil
}
