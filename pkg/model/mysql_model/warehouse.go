package mysql_model

// Warehouse 仓库
type Warehouse struct {
	BasicModel
	WarehouseName      string  `gorm:"type:varchar(50);not null;" json:"warehouse_name" cn:"仓库名称"`
	WarehouseLocation  *string `gorm:"type:varchar(200)" json:"warehouse_location,omitempty" cn:"仓库位置"`
	WarehouseOwnerID   *int    `gorm:"type:int" json:"warehouse_owner_id,omitempty" cn:"仓库管理员ID"`
	WarehouseOwnerName *string `gorm:"type:varchar(50)" json:"warehouse_owner_name,omitempty" cn:"仓库管理员名称"`
	Remark             *string `gorm:"type:varchar(200)" json:"remark,omitempty" cn:"仓库备注"`
}

func (m Warehouse) TableName() string {
	return "warehouses"
}

func (m Warehouse) TableCnName() string {
	return "仓库"
}
