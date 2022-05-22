package mysql_model

import (
	"github.com/XC-Zero/yinwan/pkg/client"
	"github.com/XC-Zero/yinwan/pkg/utils/es_tool"
	"github.com/pkg/errors"
	"gorm.io/gorm"
	"time"
)

// Material 原材料
type Material struct {
	BasicModel
	BookNameInfo
	MaterialName         string  `gorm:"type:varchar(50);not null" json:"material_name"  form:"material_name" cn:"原材料名称"`
	MaterialTypeID       int     `gorm:"type:int;not null;index" json:"material_type_id" form:"material_type_id" cn:"原材料类型ID"`
	MaterialTypeName     string  `gorm:"type:varchar(50);not null" json:"material_type_name" form:"material_type_name" cn:"原材料类型名称"`
	MaterialStyle        string  `gorm:"type:varchar(50);not null" json:"material_style" form:"material_style" cn:"原材料规格"`
	MaterialOwnerID      *int    `gorm:"type:int" json:"material_owner_id,omitempty" form:"material_owner_id,omitempty" cn:"原材料负责人ID"`
	MaterialOwnerName    *string `gorm:"type:varchar(50)" json:"material_owner_name,omitempty" form:"material_owner_name" cn:"原材料负责人名称"`
	AverageUnitPrice     *string `gorm:"type:varchar(50)" json:"average_unit_price,omitempty" form:"average_unit_price,omitempty" cn:"原材料平均价格"`
	MaterialPresentCount int     `form:"material_present_count" json:"material_present_count" cn:"当前余量"`
	MaterialPicUrl       *string `gorm:"type:varchar(500)" json:"material_pic_url,omitempty" form:"material_pic_url,omitempty" cn:"原材料展示图"`
	Remark               *string `gorm:"type:varchar(200)" json:"remark,omitempty" form:"remark,omitempty" cn:"原材料备注"`
}

func (m *Material) AfterCreate(tx *gorm.DB) error {
	bookName := tx.Statement.Context.Value("book_name").(string)
	bk, ok := client.ReadBookMap(bookName)
	if !ok {
		return errors.New("There is no book name!")
	}
	m.BookNameID = bk.StorageName
	m.BookName = bk.BookName
	err := client.PutIntoIndex(m)
	if err != nil {
		return err
	}
	return nil
}

// AfterUpdate 同步更新原材料
func (m *Material) AfterUpdate(tx *gorm.DB) error {
	bookName := tx.Statement.Context.Value("book_name").(string)
	bk, ok := client.ReadBookMap(bookName)
	if !ok {
		return errors.New("There is no book name!")
	}
	m.BookNameID = bk.StorageName
	m.BookName = bk.BookName
	err := client.UpdateIntoIndex(m, m.RecID, tx.Statement.Context, es_tool.ESDocToUpdateScript(m.ToESDoc()))
	if err != nil {
		return err
	}
	return nil
}

// AfterDelete 同步删除原材料
func (m *Material) AfterDelete(tx *gorm.DB) error {
	err := client.DeleteFromIndex(m, m.RecID, tx.Statement.Context)
	if err != nil {
		return err
	}
	return nil
}

func (m Material) TableName() string {
	return "materials"
}
func (m Material) TableCnName() string {
	return "原材料"
}
func (m Material) Mapping() map[string]interface{} {
	ma := mapping{
		"settings": mapping{},
		"mappings": mapping{
			"properties": mapping{
				"rec_id": mapping{
					"type": "keyword",
				},
				"material_name": mapping{
					"type":            "text",   //字符串类型且进行分词, 允许模糊匹配
					"analyzer":        IK_SMART, //设置分词工具
					"search_analyzer": IK_SMART,
					"fields": mapping{ //当需要对模糊匹配的字符串也允许进行精确匹配时假如此配置
						"keyword": mapping{
							"type":         "keyword",
							"ignore_above": 256,
						},
					},
				},
				"remark": mapping{
					"type":            "text",
					"analyzer":        IK_SMART,
					"search_analyzer": IK_SMART,
				},
				"created_at": mapping{
					"type": "text",
				},
				"material_pic_url": mapping{
					"type": "text",
				},
				"book_name": mapping{
					"type": "keyword",
				},
				"book_name_id": mapping{
					"type": "keyword",
				},
			},
		},
	}
	return ma
}

func (m Material) ToESDoc() map[string]interface{} {
	return map[string]interface{}{
		"rec_id":           m.RecID,
		"created_at":       m.CreatedAt,
		"remark":           m.Remark,
		"material_name":    m.MaterialName,
		"material_pic_url": m.MaterialPicUrl,
		"book_name":        m.BookName,
		"book_name_id":     m.BookNameID,
	}
}

// MaterialBatch 原材料批次
type MaterialBatch struct {
	BasicModel
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
