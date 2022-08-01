package mongo_model

import (
	"fmt"
	"github.com/XC-Zero/yinwan/internal/controller/services_controller/common"
	"github.com/XC-Zero/yinwan/pkg/client"
	_const "github.com/XC-Zero/yinwan/pkg/const"
	"github.com/XC-Zero/yinwan/pkg/model/mysql_model"
	"github.com/XC-Zero/yinwan/pkg/utils/convert"
	"github.com/XC-Zero/yinwan/pkg/utils/errs"
	"github.com/XC-Zero/yinwan/pkg/utils/logger"
	myMongo "github.com/XC-Zero/yinwan/pkg/utils/mongo"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"golang.org/x/net/context"
	"gorm.io/gorm"
	"log"
	"strconv"
	"time"
)

// StockInRecord 入库记录
type StockInRecord struct {
	BasicModel             `bson:"inline"`
	BookNameInfo           `bson:"-"`
	StockInRecordName      *string              `json:"stock_in_record_name,omitempty" form:"stock_in_record_name,omitempty" bson:"stock_in_record_name,omitempty"`
	StockInRecordOwnerID   *int                 `json:"stock_in_record_owner_id,omitempty"  form:"stock_in_record_owner_id" bson:"stock_in_record_owner_id" `
	StockInRecordOwnerName *string              `json:"stock_in_record_owner_name,omitempty" form:"stock_in_record_owner_name" bson:"stock_in_record_owner_name" `
	StockInWarehouseID     *int                 `json:"stock_in_warehouse_id,omitempty" form:"stock_in_warehouse_id,omitempty" bson:"stock_in_warehouse_id"`
	StockInWarehouseName   *string              `json:"stock_in_warehouse_name,omitempty" form:"stock_in_warehouse_name,omitempty" bson:"stock_in_warehouse_name"`
	StockInDetailPosition  *string              `json:"stock_in_detail_position,omitempty" form:"stock_in_detail_position" bson:"stock_in_detail_position"`
	StockInRecordType      string               `json:"stock_in_record_type" form:"stock_in_record_type" bson:"stock_in_record_type" `
	StockInRecordContent   []stockRecordContent `json:"stock_in_record_content" form:"stock_in_record_content" bson:"stock_in_record_content" `
	RelateInvoice          []RelatedInvoice     `json:"relate_invoice" form:"relate_invoice" bson:"relate_invoice"`
	Remark                 *string              `json:"remark,omitempty" form:"remark" bson:"remark"`
}

func (s *StockInRecord) ToCredential(ctx *gin.Context) {
	book := common.HarvestClientFromGinContext(ctx)
	if book == nil {
		return
	}
	content := s.StockInRecordContent
	sprintf := fmt.Sprintf("系统根据入库记录 [单号]:%d 生成", s.RecID)
	recID := int(time.Now().Unix())
	events := make([]CredentialEvent, 0, len(content)*2)
	name := "入库单凭证"
	var cre = Credential{
		BasicModel: BasicModel{
			RecID:     &recID,
			CreatedAt: strconv.Itoa(int(time.Now().Unix())),
		},
		BookNameInfo:        s.BookNameInfo,
		CredentialName:      &name,
		CredentialOwnerID:   0,
		CredentialOwnerName: "系统",
		Remark:              &sprintf,
	}
	for _, c := range content {
		events = append(events,
			CredentialEvent{
				Abstract:       fmt.Sprintf("入库-%s", c.Name),
				Classify:       fmt.Sprintf("借: %s-%s", c.ContentType.Display(), c.Name),
				DetailClassify: "",
				LoanAmount:     c.TotalPrice,
				DebitAmount:    "",
			},
			CredentialEvent{
				Abstract:       fmt.Sprintf("入库-%s", c.Name),
				Classify:       fmt.Sprintf("贷: %s-%s", c.ContentType.Display(), c.Name),
				DetailClassify: "",
				LoanAmount:     "",
				DebitAmount:    c.TotalPrice,
			})
	}
	cre.CredentialEvents = events
	common.CreateOneMongoDBRecordTemplate(ctx, common.CreateMongoDBTemplateOptions{
		DB:         book.MongoDBClient,
		Context:    ctx,
		TableModel: cre,
		NotSyncES:  false,
	})
	common.UpdateOneMongoDBRecordByIDTemplate(ctx, common.MongoDBTemplateOptions{
		DB:         book.MongoDBClient,
		Context:    ctx,
		TableModel: s,
		NotSyncES:  true,
	})

	ctx.JSON(_const.OK, errs.CreateSuccessMsg("生成凭证成功!", map[string]int{
		"credential_id": recID,
	}))

}

func (s StockInRecord) ToESDoc() map[string]interface{} {
	return map[string]interface{}{
		"rec_id":               s.RecID,
		"created_at":           s.CreatedAt,
		"remark":               s.Remark,
		"stock_in_content":     convert.StructSliceToTagString(s.StockInRecordContent, string(_const.CN)),
		"stock_in_record_type": s.StockInRecordType,
		"stock_in_owner":       s.StockInRecordOwnerName,
		"book_name":            s.BookName,
		"book_name_id":         s.BookNameID,
		"stock_in_record_name": s.StockInRecordName,
	}
}

func (s StockInRecord) TableName() string {
	return "stock_in_records"
}
func (s StockInRecord) TableCnName() string {
	return "入库记录"
}
func (s StockInRecord) Mapping() map[string]interface{} {
	ma := mapping{
		"settings": mapping{},
		"mappings": mapping{
			"properties": mapping{
				"rec_id": mapping{
					"type": "keyword",
				},
				"stock_in_record_name": mapping{
					"type":            "text",
					"analyzer":        IK_SMART,
					"search_analyzer": IK_SMART,
				},
				"stock_in_content": mapping{
					"type":            "text",
					"analyzer":        IK_SMART,
					"search_analyzer": IK_SMART,
				},
				"stock_in_owner": mapping{
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
				"stock_in_record_type": mapping{
					"type": "keyword",
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

// BeforeInsert 创建入库
//
func (s *StockInRecord) BeforeInsert(ctx context.Context) error {
	bookName := ctx.Value("book_name").(string)
	bk, ok := client.ReadBookMap(bookName)
	if !ok {
		return errors.New("There is no book name!")
	}
	if s.RecID == nil || bk.StorageName == "" {
		return errors.New("缺少主键！")
	}
	err := bk.MysqlClient.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		var contentList = s.StockInRecordContent
		var now = time.Now().Format("2006-01-02")
		var remark = fmt.Sprintf("根据入库单[编号]:%d 自动创建", *s.RecID)
		for i := 0; i < len(contentList); i++ {
			content := contentList[i]

			if content.ContentType == _const.MATERIAL {
				relateID := "material_id"
				model := mysql_model.MaterialBatch{
					BasicModel: mysql_model.BasicModel{
						CreatedAt: time.Now(),
					},
					MaterialID:                 content.RecID,
					MaterialName:               content.Name,
					StockInRecordID:            *s.RecID,
					MaterialBatchOwnerID:       s.StockInRecordOwnerID,
					MaterialBatchOwnerName:     s.StockInRecordOwnerName,
					MaterialBatchTotalPrice:    content.TotalPrice,
					MaterialBatchNumber:        content.Num,
					MaterialBatchSurplusNumber: content.Num,
					MaterialBatchUnitPrice:     content.Price,
					WarehouseID:                s.StockInWarehouseID,
					WarehouseName:              s.StockInWarehouseName,
					StockInTime:                &now,
					Remark:                     &remark,
				}
				log.Printf("model type is %T", model)
				err := tx.WithContext(context.WithValue(context.Background(), relateID, content.RecID)).
					Create(&model).Error

				if err != nil {
					return err
				}
			} else if content.ContentType == _const.COMMODITY {
				relateID := "commodity_id"
				model := mysql_model.CommodityBatch{
					BasicModel: mysql_model.BasicModel{
						CreatedAt: time.Now(),
					},
					CommodityID:                 content.RecID,
					CommodityName:               content.Name,
					StockInRecordID:             *s.RecID,
					CommodityBatchOwnerID:       s.StockInRecordOwnerID,
					CommodityBatchOwnerName:     s.StockInRecordOwnerName,
					CommodityBatchTotalPrice:    content.TotalPrice,
					CommodityBatchNumber:        content.Num,
					CommodityBatchSurplusNumber: content.Num,
					CommodityBatchUnitPrice:     content.Price,
					WarehouseID:                 s.StockInWarehouseID,
					WarehouseName:               s.StockInWarehouseName,
					StockInTime:                 &now,
					Remark:                      &remark,
				}
				log.Printf("model type is %T", model)
				err := tx.WithContext(context.WithValue(context.Background(), relateID, content.RecID)).
					Create(&model).Error

				if err != nil {
					return err
				}

			} else {
				return errors.New("未选择入库物品类型!")
			}

		}
		return nil
	})
	if err != nil {
		logger.Error(errors.WithStack(err), "?????")
		return err
	}
	return nil
}

func (s *StockInRecord) BeforeUpdate(ctx context.Context) error {
	return myMongo.CancelError
}

func (s *StockInRecord) BeforeRemove(ctx context.Context) error {
	return myMongo.CancelError
}
