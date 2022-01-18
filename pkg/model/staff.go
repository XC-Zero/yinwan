package model

// Staff 职工表
type Staff struct {
	staffID     *int    `gorm:"primaryKey;type:int;autoIncrement;"json:"id"`
	StaffName   string  `gorm:"type:varchar(50);not null;"`
	StaffAlias  *string `gorm:"type:varchar(50)"` // 职工别名
	StaffRoleID int     `gorm:"type:varchar(50)"`
}

func (s *Staff) GetStaffID() int {
	return *s.staffID
}

// StaffRole 职工角色表
type StaffRole struct {
	staffRoleID *int `gorm:"primaryKey;type:int;autoIncrement"json:"id"`
}
