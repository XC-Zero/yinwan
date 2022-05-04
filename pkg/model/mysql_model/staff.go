package mysql_model

import (
	"fmt"
	"github.com/XC-Zero/yinwan/pkg/client"
	"github.com/XC-Zero/yinwan/pkg/utils/logger"
	"github.com/XC-Zero/yinwan/pkg/utils/token"
	"github.com/fwhezfwhez/errorx"
	"time"
)

//goland:noinspection GoSnakeCaseUsage
const EXPIRE_TIME = time.Hour * 2

// Staff 职工表

type Staff struct {
	BasicModel
	StaffName           string  `gorm:"type:varchar(50);not null;" json:"staff_name" form:"staff_name" cn:"职工名称"`
	StaffAlias          *string `gorm:"type:varchar(50)" json:"staff_alias,omitempty" form:"staff_alias" cn:"职工别名"` // 职工别名
	StaffEmail          string  `gorm:"type:varchar(50);not null;index;unique" json:"staff_email"  form:"staff_email" binding:"required" cn:"职工邮箱"`
	StaffPhone          *string `gorm:"type:varchar(50);index;" json:"staff_phone,omitempty" form:"staff_phone"`
	StaffPassword       string  `gorm:"type:varchar(20)" json:"staff_password" form:"staff_password" `
	StaffPosition       *string `gorm:"type:varchar(50)" json:"staff_position,omitempty"  form:"staff_position"`
	StaffDepartmentID   *int    `json:"staff_department_id,omitempty"  form:"staff_department_id"`
	StaffDepartmentName *string `gorm:"type:varchar(50)" json:"staff_department_name,omitempty" form:"staff_department_name"  cn:"职工部门名称"`
	StaffRoleID         int     `gorm:"type:int" json:"staff_role_id" form:"staff_role_id" `
	StaffRoleName       string  `gorm:"type:varchar(50)" json:"staff_role_name" form:"staff_role_name"  `
}

func (s Staff) TableName() string {
	return "staffs"
}
func (s Staff) TableCnName() string {
	return "员工"
}

// Login 登录
// 查mysql ,校验一下，生成个 token 丢 redis ，设置 2 小时过期
// 返回 token 指针 和 错误信息
func (s Staff) Login() (tokenPtr *string, errorMessage string) {
	temp := Staff{}
	err := client.MysqlClient.Model(&Staff{}).Find(&temp, "  staff_email =? ", s.StaffEmail).Error
	if err != nil || temp.RecID == nil {
		logger.Error(errorx.MustWrap(err), fmt.Sprintf("用户登录失败，数据库内找不到此用户, staff is %+v\n ,error is %s", s, err))
		return nil, "无此用户"
	}
	if s.StaffPassword != temp.StaffPassword {
		return nil, "抱歉，密码不正确"
	}
	tokenStr, err := token.GenerateToken(temp.StaffEmail)
	if err != nil {
		logger.Error(errorx.MustWrap(err), fmt.Sprintf("生成 token 失败, error is %s", err))
	}
	err = client.RedisClient.Set(tokenStr, s.StaffEmail, EXPIRE_TIME).Err()
	if err != nil {
		logger.Error(errorx.MustWrap(err), fmt.Sprintf("redis 存储 token 失败, error is %s", err))
	}
	tokenPtr = &tokenStr
	return
}

// LogOut 退出登录
func (s Staff) LogOut() {
	err := client.RedisClient.Del(s.StaffEmail).Err()
	if err != nil {
		logger.Error(errorx.MustWrap(err), fmt.Sprintf("redis 删除 key 为 %s 的 token 失败, error is %s", s.StaffEmail, err))
	}
}

// IgnoreStaffPassword 去除密码
func IgnoreStaffPassword(staffList []interface{}) []interface{} {
	staffs := make([]interface{}, 0, len(staffList))
	for i := range staffList {
		staff, _ := staffList[i].(Staff)
		staff.StaffPassword = "********"
		staffs = append(staffs, staff)
	}
	return staffs
}

type Department struct {
	BasicModel
	DepartmentName        string  `gorm:"type:varchar(50)" json:"department_name" form:"department_name" binding:"required" cn:"部门名称"`
	DepartmentLocation    *string `gorm:"type:varchar(50)" json:"department_location,omitempty" form:"department_location" cn:"部门地点"`
	DepartmentManagerID   *int    `gorm:"type:int;index"  json:"department_manager_id,omitempty" form:"department_manager_id" cn:"部门主管ID"`
	DepartmentManagerName *string `gorm:"type:varchar(50)"  json:"department_manager_name,omitempty" form:"department_manager_name" cn:"部门主管名称"`
	DepartmentPhone       *string `gorm:"type:varchar(50)" json:"department_phone,omitempty" form:"department_phone" cn:"部门联系电话"`
}

func (d Department) TableName() string {
	return "departments"
}

func (d Department) TableCnName() string {
	return "部门"
}
