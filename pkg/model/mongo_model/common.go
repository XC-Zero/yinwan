package mongo_model

import (
	"encoding/json"
	_const "github.com/XC-Zero/yinwan/pkg/const"
	"log"
)

type Analyzer string

const (
	IK_SMART    = "ik_smart"
	IK_MAX_WORD = "ik_max_word"
)

type BasicModel struct {
	RecID     *int    `json:"rec_id,omitempty" bson:"rec_id" cn:"记录ID"`
	CreatedAt string  `json:"created_at" bson:"created_at"`
	UpdatedAt *string `json:"updated_at,omitempty" bson:"updated_at"`
	DeletedAt *string `json:"deleted_at,omitempty" bson:"deleted_at" cn:"删除时间"`
}

type mapping map[string]interface{}
type BookNameInfo struct {
	BookNameID string `gorm:"-" sql:"-" json:"book_name_id" cn:"账套编号"`
	BookName   string `gorm:"-" sql:"-" json:"book_name" cn:"账套名称"`
}

// 出入库通用
type stockRecordContent struct {
	RecID          int                     `json:"rec_id" form:"rec_id" bson:"rec_id" cn:"产品或原材料编号"`
	Name           string                  `json:"name" form:"name" bson:"name" cn:"产品或原材料名称"`
	Num            int                     `json:"num" form:"num" bson:"num" cn:"产品或原材料编号"`
	Price          string                  `json:"price" form:"price" bson:"price" cn:"产品或原材料单价"`
	TotalPrice     string                  `json:"total_price" form:"total_price" bson:"total_price" cn:"产品或原材料总价"`
	RelatedBatchID *int                    `json:"related_batch_id,omitempty" form:"related_batch_id,omitempty" bson:"related_batch_id,omitempty" cn:"产品或原材料批次编号"`
	ContentType    _const.StockContentType `json:"content_type" form:"content_type" bson:"content_type" cn:"产品/原材料"`
}

// RelatedInvoice 关联相关单据
type RelatedInvoice struct {
	InvoiceType _const.InvoiceType `json:"invoice_type" form:"invoice_tye" bson:"invoice_tye" cn:"单据类型"`
	InvoiceID   int                `json:"invoice_id" form:"invoice_id" bson:"invoice_id" cn:"单据编号"`
}

//func (s *stockRecordContent) TransferByContentType() interface{} {
//	var content interface{}
//	switch s.ContentType {
//	case _const.COMMODITY:
//
//		content = mysql_model.Commodity{
//			BasicModel: mysql_model.BasicModel{
//				RecID: &s.RecID,
//			},
//		}
//	case _const.MATERIAL:
//		content = mysql_model.Material{
//			BasicModel: mysql_model.BasicModel{
//				RecID: &s.RecID,
//			},
//			MaterialName: s.Name,
//
//		}
//
//	}
//
//}

func (m mapping) String() string {
	marshal, err := json.Marshal(m)
	if err != nil {
		return ""
	}
	res := string(marshal)
	log.Println(res)
	return res
}
