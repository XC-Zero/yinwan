package mysql_model

import (
	"fmt"
	"github.com/XC-Zero/yinwan/pkg/client"
	"github.com/XC-Zero/yinwan/pkg/utils/es_tool"
	"github.com/XC-Zero/yinwan/pkg/utils/logger"
	"github.com/XC-Zero/yinwan/pkg/utils/math_plus"
	"github.com/pkg/errors"
	"gorm.io/gorm"
	"log"
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
func (m Material) AfterDelete(tx *gorm.DB) error {
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
	MaterialID                 int     `gorm:"type:int;not null" json:"material_id" cn:"原材料ID"`
	MaterialName               string  `gorm:"type:varchar(50);not null" json:"material_name" cn:"原材料名称"`
	StockInRecordID            int     `gorm:"type:int;not null" json:"stock_in_record_id" cn:"关联入库单号"`
	MaterialBatchOwnerID       *int    `gorm:"type:int;index" json:"material_batch_owner_id,omitempty" cn:"负责人ID"`
	MaterialBatchOwnerName     *string `gorm:"type:varchar(50)" json:"material_batch_owner_name,omitempty" cn:"负责人名称"`
	MaterialBatchTotalPrice    string  `gorm:"type:varchar(50);not null" json:"material_batch_total_price" cn:"批次总价"`
	MaterialBatchNumber        int     `gorm:"type:int;not null" json:"material_batch_number" cn:"批次原材料总数"`
	MaterialBatchSurplusNumber int     `gorm:"type:int;not null" json:"material_batch_surplus_number" cn:"当前批次原材料剩余数量"`
	MaterialBatchUnitPrice     string  `gorm:"type:varchar(50);not null" json:"material_batch_unit_price" cn:"单价"`
	WarehouseID                *int    `gorm:"type:int;index" json:"warehouse_id,omitempty" cn:"仓库ID"`
	WarehouseName              *string `gorm:"type:varchar(50)" json:"warehouse_name,omitempty"  cn:"仓库名称"`
	StockInTime                *string `gorm:"type:timestamp " json:"stock_in_time" cn:"入库时间"`
	Remark                     *string `gorm:"type:varchar(200)" json:"remark,omitempty" cn:"批次备注"`
}
type tempScan struct {
	Surplus     int64 `json:"surplus"`
	TotalPrice  int64 `json:"total_price"`
	TotalNumber int64 `json:"total_number"`
}

// AfterCreate 同步创建历史成本
func (m MaterialBatch) AfterCreate(tx *gorm.DB) error {

	id := tx.Statement.Context.Value("material_id")
	rec, ok := id.(int)
	if id == nil || !ok {
		logger.Error(errors.New("Context have not material_id"), "")
		return nil
	}
	var cost = MaterialHistoryCost{
		MaterialID:             rec,
		Num:                    m.MaterialBatchNumber,
		Price:                  m.MaterialBatchTotalPrice,
		RelatedMaterialBatchID: *m.RecID,
	}

	err := tx.Transaction(func(tx *gorm.DB) error {
		var res MaterialHistoryCost
		err2 := tx.Where("related_material_batch_id = ?", *m.RecID).Find(&res).Error
		if err2 != nil {
			log.Println(errors.WithStack(err2))
			return err2
		}
		if res.RecID == nil {
			err2 := tx.Model(cost).Create(&cost).Error
			if err2 != nil {
				log.Println("@@@@ Create ERROR is ", errors.WithStack(err2))
				return err2
			}
		} else {
			err2 := tx.Updates(&cost).Where("rec_id", *res.RecID).Error
			if err2 != nil {
				log.Println("@@@@ Updates ERROR is ", errors.WithStack(err2))
				return err2
			}
		}
		var tempScan tempScan

		err2 = tx.Model(MaterialBatch{}).Where("material_id = ? and deleted_at is null ", rec).Select(
			"sum(material_batch_surplus_number) as surplus," +
				"sum(material_batch_total_price) as total_price," +
				"sum(material_batch_number) as total_number ").Scan(&tempScan).Error
		if err2 != nil {
			log.Println("@@@@ Statics ERROR is ", errors.WithStack(err2))

			return err2
		}
		var averagePrice string
		if tempScan.TotalNumber == 0 {
			averagePrice = "0"
		} else {
			fraction, err := math_plus.New(tempScan.TotalPrice, tempScan.TotalNumber)
			if err != nil {
				averagePrice = "0"
			}
			averagePrice = fmt.Sprintf("%.2f", fraction.Float64())
		}

		err2 = tx.Raw("update materials set average_unit_price = ? , "+
			"material_present_count = ? where rec_id = ? ",
			averagePrice, tempScan.Surplus, rec).Error
		if err2 != nil {
			log.Println(errors.WithStack(err2))
			return err2
		}
		return nil
	},
	)
	if err != nil {
		return err
	}
	return nil
}

func (m MaterialBatch) AfterUpdate(tx *gorm.DB) error {
	return m.AfterCreate(tx)
}

func (m MaterialBatch) TableName() string {
	return "material_batches"
}
func (m MaterialBatch) TableCnName() string {
	return "原材料批次"
}

// MaterialHistoryCost 原材料历史进货价
type MaterialHistoryCost struct {
	BasicModel
	MaterialID             int    `gorm:"type:int;not null;index" json:"material_id" form:"material_id"`
	Price                  string `gorm:"type:varchar(20)" json:"price" form:"price"`
	Num                    int    `json:"num" form:"num"`
	RelatedMaterialBatchID int    `gorm:"type:int;not null;index" json:"related_material_batch_id" form:"related_material_batch_id"`
}

func (m MaterialHistoryCost) TableName() string {
	return "material_history_costs"
}
func (m MaterialHistoryCost) TableCnName() string {
	return "原材料历史进货价"
}
