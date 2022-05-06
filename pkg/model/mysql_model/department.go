package mysql_model

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
