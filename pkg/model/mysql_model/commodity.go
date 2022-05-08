package mysql_model

import (
	"github.com/XC-Zero/yinwan/pkg/client"
	"github.com/olivere/elastic/v7"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

// Commodity 产品
type Commodity struct {
	BasicModel
	BookNameInfo
	CommodityName      string  `gorm:"type:varchar(200)" json:"commodity_name"`
	CommodityType      *string `gorm:"type:varchar(50)" json:"commodity_type,omitempty"`
	CommodityStyle     *string `gorm:"type:varchar(50)" json:"commodity_style,omitempty"`
	CommodityOwnerID   int     `gorm:"type:int" json:"commodity_owner_id"`
	CommodityOwnerName string  `gorm:"type:varchar(50)" json:"commodity_owner_name"`
	CommodityRemark    *string `gorm:"type:varchar(200)" json:"commodity_remark,omitempty"`
	CommodityAttribute string  `gorm:"type:varchar(500)" json:"commodity_attribute"`
	CommodityPicUrl    *string `gorm:"type:varchar(500)" json:"commodity_pic_url,omitempty" cn:"产品展示图"`
	Remark             *string `gorm:"type:varchar(200)" json:"remark,omitempty" cn:"产品备注"`
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
func (c *Commodity) AfterUpdate(tx *gorm.DB) error {
	err := client.UpdateIntoIndex(c, c.RecID, tx,
		elastic.NewScriptInline("ctx._source.nickname=params.nickname;ctx._source.ancestral=params.ancestral").
			Params(c.ToESDoc()))
	if err != nil {
		return err
	}
	return nil
}
func (c *Commodity) AfterDelete(tx *gorm.DB) error {
	err := client.DeleteFromIndex(c, c.RecID, tx)
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
