package model

type Department struct {
	BasicModel
	DepartmentName        string  `gorm:"type:varchar(50)" json:"department_name" cn:"部门名称"`
	DepartmentLocation    string  `gorm:"type:varchar(50)" json:"department_location" cn:"部门地点"`
	DepartmentManagerID   int     `json:"department_manager_id" cn:"部门主管ID"`
	DepartmentManagerName string  `gorm:"type:varchar(50)" json:"department_manager_name" cn:"部门主管名称"`
	DepartmentPhone       *string `gorm:"type:varchar(50)" json:"department_phone,omitempty" cn:"部门联系电话"`
}
