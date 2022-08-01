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

// Commodity 产品
type Commodity struct {
	BasicModel
	BookNameInfo
	CommodityName             string  `gorm:"type:varchar(200)" json:"commodity_name" cn:"产品名称"`
	CommodityTypeID           *int    `form:"commodity_type_id" json:"commodity_type_id,omitempty" cn:"产品类型ID"`
	CommodityType             *string `gorm:"type:varchar(50)" json:"commodity_type,omitempty" cn:"产品类型"`
	CommodityStyle            *string `gorm:"type:varchar(50)" json:"commodity_style,omitempty" cn:"产品规格"`
	CommodityOwnerID          int     `gorm:"type:int" json:"commodity_owner_id" cn:"产品负责人ID"`
	CommodityOwnerName        string  `gorm:"type:varchar(50)" json:"commodity_owner_name" cn:"产品负责人"`
	CommodityAverageCostPrice string  `gorm:"type:varchar(50)" form:"commodity_average_cost_price" json:"commodity_average_cost_price" cn:"平均成本价"`
	CommodityPrice            string  `gorm:"type:varchar(50)" form:"commodity_price" json:"commodity_price" cn:"产品定价"`
	CommodityPicUrl           *string `gorm:"type:varchar(500)" json:"commodity_pic_url,omitempty" cn:"产品展示图"`
	Remark                    *string `gorm:"type:varchar(200)" json:"remark,omitempty" cn:"产品备注"`
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
					"type":            "text",
					"analyzer":        IK_SMART,
					"search_analyzer": IK_SMART,
					"fields": mapping{
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

// AfterUpdate 同步更新
func (c *Commodity) AfterUpdate(tx *gorm.DB) error {
	bookName := tx.Statement.Context.Value("book_name").(string)
	bk, ok := client.ReadBookMap(bookName)
	if !ok {
		return errors.New("There is no book name!")
	}
	c.BookNameID = bk.StorageName
	c.BookName = bk.BookName

	err := client.UpdateIntoIndex(c, c.RecID, tx.Statement.Context, es_tool.ESDocToUpdateScript(c.ToESDoc()))

	if err != nil {
		return err
	}
	return nil
}

func (c Commodity) AfterDelete(tx *gorm.DB) error {
	err := client.DeleteFromIndex(c, c.RecID, tx.Statement.Context)
	if err != nil {
		return err
	}
	return nil
}

// CommodityHistoricalCost 历史成本表
type CommodityHistoricalCost struct {
	BasicModel
	CommodityID             int    `gorm:"type:int;not null;index" json:"commodity_id" form:"commodity_id"`
	Price                   string `gorm:"type:varchar(20)" json:"price" form:"price"`
	Num                     int    `json:"num" form:"num"`
	RelatedCommodityBatchID int    `gorm:"type:int;not null;index" json:"related_commodity_batch_id" form:"related_commodity_batch_id"`
}

func (c CommodityHistoricalCost) TableCnName() string {
	return "历史成本"
}

func (c CommodityHistoricalCost) TableName() string {
	return "commodity_historical_costs"
}

// CommodityBatch 批次
type CommodityBatch struct {
	BasicModel
	CommodityID                 int     `gorm:"type:int;not null" json:"commodity_id" cn:"产品ID"`
	CommodityName               string  `gorm:"type:varchar(50);not null" json:"commodity_name" cn:"产品名称"`
	StockInRecordID             int     `gorm:"type:int;not null" json:"stock_in_record_id" cn:"关联入库单号"`
	CommodityBatchOwnerID       *int    `gorm:"type:int;index" json:"commodity_batch_owner_id,omitempty" cn:"负责人ID"`
	CommodityBatchOwnerName     *string `gorm:"type:varchar(50)" json:"commodity_batch_owner_name,omitempty" cn:"负责人名称"`
	CommodityBatchTotalPrice    string  `gorm:"type:varchar(50);not null" json:"commodity_batch_total_price" cn:"批次总价"`
	CommodityBatchNumber        int     `gorm:"type:int;not null" json:"commodity_batch_number" cn:"批次产品总数"`
	CommodityBatchSurplusNumber int     `gorm:"type:int;not null" json:"commodity_batch_surplus_number" cn:"当前批次产品剩余数量"`
	CommodityBatchUnitPrice     string  `gorm:"type:varchar(50);not null" json:"commodity_batch_unit_price" cn:"单价"`
	WarehouseID                 *int    `gorm:"type:int;index" json:"warehouse_id,omitempty" cn:"仓库ID"`
	WarehouseName               *string `gorm:"type:varchar(50)" json:"warehouse_name,omitempty"  cn:"仓库名称"`
	StockInTime                 *string `gorm:"type:timestamp " json:"stock_in_time" cn:"入库时间"`
	Remark                      *string `gorm:"type:varchar(200)" json:"remark,omitempty" cn:"批次备注"`
}

func (c CommodityBatch) TableCnName() string {
	return "产品批次"
}

func (c CommodityBatch) TableName() string {
	return "commodity_batches"
}

// AfterCreate 同步创建历史成本
func (c CommodityBatch) AfterCreate(tx *gorm.DB) error {

	id := tx.Statement.Context.Value("commodity_id")
	rec, ok := id.(int)
	if id == nil || !ok {
		logger.Error(errors.New("Context have not commodity_id"), "")
		return nil
	}
	var cost = CommodityHistoricalCost{
		CommodityID:             rec,
		Num:                     c.CommodityBatchNumber,
		Price:                   c.CommodityBatchUnitPrice,
		RelatedCommodityBatchID: *c.RecID,
	}

	//err := tx.Transaction(func(tx *gorm.DB) error {
	var res CommodityHistoricalCost
	err2 := tx.Where("related_commodity_batch_id = ?", *c.RecID).Find(&res).Error
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

	err2 = tx.Model(CommodityBatch{}).Where("commodity_id = ? and deleted_at is null ", rec).Select(
		"sum(commodity_batch_surplus_number) as surplus," +
			"sum(commodity_batch_total_price) as total_price," +
			"sum(commodity_batch_number) as total_number ").Scan(&tempScan).Error
	if err2 != nil {
		log.Println("[ERROR] Statics ERROR is ", errors.WithStack(err2))
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

	err2 = tx.Raw("update commodities set average_unit_price = ? , "+
		"commodity_present_count = ? where rec_id = ? ",
		averagePrice, tempScan.Surplus, rec).Error
	if err2 != nil {
		log.Println(errors.WithStack(err2))
		return err2
	}
	return nil
	//},
	//)
	//if err != nil {
	//	return err
	//}
	//return nil
}

func (c CommodityBatch) AfterUpdate(tx *gorm.DB) error {
	return c.AfterCreate(tx)
}
func (c CommodityBatch) AfterDelete(tx *gorm.DB) error {
	return c.AfterCreate(tx)
}
