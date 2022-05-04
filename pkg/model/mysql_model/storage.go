package mysql_model

import (
	"github.com/XC-Zero/yinwan/pkg/client"
	"github.com/olivere/elastic/v7"
	"github.com/pkg/errors"
	"gorm.io/gorm"
	"time"
)

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

// Material 原材料
type Material struct {
	BasicModel
	BookNameInfo
	MaterialName      string  `gorm:"type:varchar(50);not null" json:"material_name"  form:"material_name" cn:"原材料名称"`
	MaterialTypeID    int     `gorm:"type:int;not null;index" json:"material_type_id" form:"material_type_id" cn:"原材料类型ID"`
	MaterialTypeName  string  `gorm:"type:varchar(50);not null" json:"material_type_name" form:"material_type_name" cn:"原材料类型名称"`
	MaterialStyle     string  `gorm:"type:varchar(50);not null" json:"material_style" form:"material_style" cn:"原材料规格"`
	MaterialOwnerID   *int    `gorm:"type:int" json:"material_owner_id,omitempty" form:"material_owner_id,omitempty" cn:"原材料负责人ID"`
	MaterialOwnerName *string `gorm:"type:varchar(50)" json:"material_owner_name,omitempty" form:"material_owner_name" cn:"原材料负责人名称"`
	AverageUnitPrice  *string `gorm:"type:varchar(50)" json:"average_unit_price,omitempty" form:"average_unit_price,omitempty" cn:"原材料平均价格"`
	MaterialPicUrl    *string `gorm:"type:varchar(500)" json:"material_pic_url,omitempty" form:"material_pic_url,omitempty" cn:"原材料展示图"`
	Remark            *string `gorm:"type:varchar(200)" json:"remark,omitempty" form:"remark,omitempty" cn:"原材料备注"`
}

func (m *Material) AfterCreate(db *gorm.DB) error {
	bookName := db.Statement.Context.Value("book_name").(string)
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

// AfterUpdate todo !!!
func (m *Material) AfterUpdate(db *gorm.DB) error {
	err := client.UpdateIntoIndex(m, m.RecID,
		elastic.NewScriptInline("ctx._source.nickname=params.nickname;ctx._source.ancestral=params.ancestral").
			Params(m.ToESDoc()))
	if err != nil {
		return err
	}
	return nil
}
func (m *Material) AfterDelete(db *gorm.DB) error {
	err := client.DeleteFromIndex(m, m.RecID)
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

// CommodityBatch 批次
type CommodityBatch struct {
	BasicModel
}

func (c CommodityBatch) TableCnName() string {
	return "产品批次"
}
func (c CommodityBatch) TableName() string {
	return "commodity_batches"
}

// Commodity 产品
type Commodity struct {
	BasicModel
	BookNameInfo
	CommodityName      string                 `gorm:"type:varchar(200)" json:"commodity_name"`
	CommodityType      *string                `gorm:"type:varchar(50)" json:"commodity_type,omitempty"`
	CommodityStyle     *string                `gorm:"type:varchar(50)" json:"commodity_style,omitempty"`
	CommodityOwnerID   int                    `gorm:"type:int" json:"commodity_owner_id"`
	CommodityOwnerName string                 `gorm:"type:varchar(50)" json:"commodity_owner_name"`
	CommodityRemark    *string                `gorm:"type:varchar(200)" json:"commodity_remark,omitempty"`
	CommodityAttribute map[string]interface{} `gorm:"type:varchar(500)" json:"commodity_attribute"`
	CommodityPicUrl    *string                `gorm:"type:varchar(500)" json:"commodity_pic_url,omitempty" cn:"产品展示图"`
	Remark             *string                `gorm:"type:varchar(200)" json:"remark,omitempty" cn:"产品备注"`
}

func (c Commodity) TableCnName() string {
	return "产品"
}
func (c Commodity) TableName() string {
	return "commodities"
}
func (c Commodity) Mapping() map[string]interface{} {
	ma := mapping{
		"settings": mapping{},
		"mappings": mapping{
			"properties": mapping{
				"rec_id": mapping{
					"type": "keyword",
				},
				"commodity_pic_url": mapping{
					"type": "text",
				},
				"commodity_name": mapping{
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

func (c Commodity) ToESDoc() map[string]interface{} {
	return map[string]interface{}{
		"rec_id":            c.RecID,
		"created_at":        c.CreatedAt,
		"remark":            c.Remark,
		"commodity_name":    c.CommodityName,
		"commodity_pic_url": c.CommodityPicUrl,
		"book_name":         c.BookName,
		"book_name_id":      c.BookNameID,
	}
}
func (c *Commodity) AfterCreate(db *gorm.DB) error {
	bookName := db.Statement.Context.Value("book_name").(string)
	bk, ok := client.ReadBookMap(bookName)
	if !ok {
		return errors.New("There is no book name!")
	}
	c.BookNameID = bk.StorageName
	c.BookName = bk.BookName
	err := client.PutIntoIndex(c)
	if err != nil {
		return err
	}
	return nil
}

// AfterUpdate todo !!!
func (c *Commodity) AfterUpdate(db *gorm.DB) error {
	err := client.UpdateIntoIndex(c, c.RecID,
		elastic.NewScriptInline("ctx._source.nickname=params.nickname;ctx._source.ancestral=params.ancestral").
			Params(c.ToESDoc()))
	if err != nil {
		return err
	}
	return nil
}
func (c *Commodity) AfterDelete(db *gorm.DB) error {
	err := client.DeleteFromIndex(c, c.RecID)
	if err != nil {
		return err
	}
	return nil
}

// CommodityHistoricalCost 历史成本表
type CommodityHistoricalCost struct {
	BasicModel
	CommodityID   int     `gorm:"type:int;index;" json:"commodity_id"`
	CommodityName string  `gorm:"type:varchar(200)" json:"commodity_name"`
	CommodityCost float64 `gorm:"type:decimal(20,2)" json:"commodity_cost"`
}

func (c CommodityHistoricalCost) TableCnName() string {
	return "历史成本"
}
func (c CommodityHistoricalCost) TableName() string {
	return "commodity_historical_costs"
}
