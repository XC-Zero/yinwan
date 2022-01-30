package staff

import (
	"github.com/XC-Zero/yinwan/pkg/client"
	"github.com/XC-Zero/yinwan/pkg/model"
)

// GetStaff todo 通过条件找员工 先只支持 `或`
func GetStaff(staff model.Staff) (staffList []model.Staff, err error) {
	s := client.MysqlClient.Model(&model.Staff{})

	if staff.RecID != nil {
		err = s.Find(&staff, "rec_id = ?", *staff.RecID).Error
		if err != nil {
			staffList = append(staffList, staff)
		}
		return
	}
	s.Find()
}
