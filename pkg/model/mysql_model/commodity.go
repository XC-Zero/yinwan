package mysql_model

import (
	"github.com/XC-Zero/yinwan/pkg/client"
	"github.com/XC-Zero/yinwan/pkg/utils/es_tool"
	"github.com/pkg/errors"
	"gorm.io/gorm"
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
	CommodityID   int    `gorm:"type:int;index;" json:"commodity_id" form:"commodity_id"`
	CommodityName string `gorm:"type:varchar(200)" json:"commodity_name" form:"commodity_name"`
	CommodityCost string `gorm:"type:varchar(20)" json:"commodity_cost" form:"commodity_cost"`
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
}

func (c CommodityBatch) TableCnName() string {
	return "产品批次"
}
func (c CommodityBatch) TableName() string {
	return "commodity_batches"
}
