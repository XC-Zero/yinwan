package staff

import (
	"github.com/XC-Zero/yinwan/pkg/client"
	"github.com/XC-Zero/yinwan/pkg/model"
)

// SelectStaff todo 通过条件找员工 先只支持 `或`
func SelectStaff(staff model.Staff) (staffList []model.Staff, err error) {
	s := client.MysqlClient.Model(&model.Staff{})

	if staff.RecID != nil {
		err = s.Find(&staff, "rec_id = ?", *staff.RecID).Error
		if err != nil {
			staffList = append(staffList, staff)
		}
		return
	}
	//s.Find()
	return nil, nil
}

// UpdateStaff 更新员工信息
func UpdateStaff(staff model.Staff) bool {
	//s := client.MysqlClient.Model(&model.Staff{})
	//s.Update()
	return false
}
