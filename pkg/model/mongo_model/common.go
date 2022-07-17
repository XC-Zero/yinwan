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
	RecID     *int    ` json:"rec_id,omitempty" bson:"rec_id" cn:"记录ID"`
	CreatedAt string  `json:"created_at" bson:"created_at"`
	UpdatedAt *string `json:"updated_at,omitempty" bson:"updated_at"`
	DeletedAt *string `json:"deleted_at,omitempty" bson:"deleted_at" cn:"删除时间"`
}

type mapping map[string]interface{}
type BookNameInfo struct {
	BookNameID string `gorm:"-" sql:"-" json:"book_name_id" `
	BookName   string `gorm:"-" sql:"-" json:"book_name" `
}

type stockRecordContent struct {
	RecID          int                     `json:"rec_id" form:"rec_id" bson:"rec_id"`
	Name           string                  `json:"name" form:"name" bson:"name"`
	Num            int                     `json:"out_num" form:"out_num" bson:"out_num"`
	Price          string                  `json:"out_price" form:"out_price" bson:"out_price"`
	TotalPrice     string                  `json:"total_price" form:"total_price" bson:"total_price"`
	RelatedBatchID *int                    `json:"related_batch_id,omitempty" form:"related_batch_id,omitempty" bson:"related_batch_id,omitempty"`
	ContentType    _const.StockContentType `json:"content_type" form:"content_type" bson:"content_type"`
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
