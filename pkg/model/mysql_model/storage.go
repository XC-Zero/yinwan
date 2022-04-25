package mysql_model

import (
	"github.com/XC-Zero/yinwan/pkg/model/mysql_model/common"
	"time"
)

// Warehouse 仓库
type Warehouse struct {
	common.BasicModel
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

// Material 原材料
type Material struct {
	common.BasicModel
	MaterialName      string  `gorm:"type:varchar(50);not null" json:"material_name" cn:"原材料名称"`
	MaterialTypeID    int     `gorm:"type:int;not null;index" json:"material_type_id" cn:"原材料类型ID"`
	MaterialTypeName  string  `gorm:"type:varchar(50);not null" json:"material_type_name" cn:"原材料类型名称"`
	MaterialStyle     string  `gorm:"type:varchar(50);not null" json:"material_style" cn:"原材料规格"`
	MaterialOwnerID   *int    `gorm:"type:int" json:"material_owner_id,omitempty" cn:"原材料负责人ID"`
	MaterialOwnerName *string `gorm:"type:varchar(50)" json:"material_owner_name,omitempty" cn:"原材料负责人名称"`
	AverageUnitPrice  *string `gorm:"type:varchar(50)" json:"average_unit_price,omitempty" cn:"原材料平均价格"`
	MaterialPicUrl    *string `gorm:"type:varchar(500)" json:"material_pic_url,omitempty" cn:"原材料展示图"`
	Remark            *string `gorm:"type:varchar(200)" json:"remark,omitempty" cn:"原材料备注"`
}

func (m Material) TableName() string {
	return "materials"
}
func (m Material) TableCnName() string {
	return "原材料"
}

// MaterialBatch 原材料批次
type MaterialBatch struct {
	common.BasicModel
	MaterialID                 int        `gorm:"type:int;not null" json:"material_id" cn:"原材料ID"`
	MaterialName               string     `gorm:"type:varchar(50);not null" json:"material_name" cn:"原材料名称"`
	StockInRecordID            int        `gorm:"type:int;not null" json:"stock_in_record_id" cn:"关联入库单号"`
	MaterialBatchOwnerID       *int       `gorm:"type:int;index" json:"material_batch_owner_id,omitempty" cn:"负责人ID"`
	MaterialBatchOwnerName     *string    `gorm:"type:varchar(50)" json:"material_batch_owner_name,omitempty" cn:"负责人名称"`
	MaterialBatchTotalPrice    string     `gorm:"type:varchar(50);not null" json:"material_batch_total_price" cn:"批次总价"`
	MaterialBatchNumber        int        `gorm:"type:int;not null" json:"material_batch_number" cn:"批次原材料总数"`
	MaterialBatchSurplusNumber int        `gorm:"type:int;not null" json:"material_batch_surplus_number" cn:"当前批次原材料剩余数量"`
	MaterialBatchUnitPrice     string     `gorm:"type:varchar(50);not null" json:"material_batch_unit_price" cn:"单价"`
	WarehouseID                *int       `gorm:"type:int;index" json:"warehouse_id,omitempty" cn:"仓库ID"`
	WarehouseName              *string    `gorm:"type:int" json:"warehouse_name,omitempty"  cn:"仓库名称"`
	StockInTime                *time.Time `gorm:"type:timestamp " json:"stock_in_time" cn:"入库时间"`
	Remark                     *string    `gorm:"type:varchar(200)" json:"remark,omitempty" cn:"批次备注"`
}

func (m MaterialBatch) TableName() string {
	return "material_batches"
}
func (m MaterialBatch) TableCnName() string {
	return "原材料批次"
}
