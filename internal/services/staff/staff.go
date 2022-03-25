package staff

import (
	"github.com/XC-Zero/yinwan/pkg/model"
	"gorm.io/gorm"
)

// SelectStaff 查找员工
func SelectStaff(db *gorm.DB, staff model.Staff) (staffList []model.Staff, bool2 bool) {

	return nil, false
}

// UpdateStaff 更新员工信息
func UpdateStaff(db *gorm.DB, oldStaff, newStaff model.Staff) bool {

	return false
}

// AddStaff 添加员工
func AddStaff(db *gorm.DB, staff model.Staff) {

}
