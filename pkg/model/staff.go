package model

import (
	"fmt"
	"github.com/XC-Zero/yinwan/pkg/client"
	"github.com/XC-Zero/yinwan/pkg/utils/logger"
	"github.com/XC-Zero/yinwan/pkg/utils/token"
	"github.com/fwhezfwhez/errorx"
	"strconv"
	"time"
)

//goland:noinspection GoSnakeCaseUsage
const EXPIRE_TIME = time.Hour * 2

// Staff 职工表
type Staff struct {
	BasicModel
	StaffName     string  `gorm:"type:varchar(50);not null;" json:"staff_name"`
	StaffAlias    *string `gorm:"type:varchar(50)" json:"staff_alias"` // 职工别名
	StaffEmail    string  `gorm:"type:varchar(50);not null;index;unique" json:"staff_email"`
	StaffPhone    *string `gorm:"type:varchar(50);index;" json:"staff_phone"`
	StaffPassword string  `gorm:"type:varchar(20)" json:"staff_password"`
	StaffRoleID   int     `gorm:"type:int" json:"staff_role_id"`
	StaffRoleName string  `gorm:"type:varchar(50)" json:"staff_role_name"`
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
	tokenStr, err := token.GenerateToken(strconv.Itoa(*temp.RecID))
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

//func (s Staff) HarvestRoleResponse() {
//	client.MysqlClient.Model(&Role{}).Find()
//}
