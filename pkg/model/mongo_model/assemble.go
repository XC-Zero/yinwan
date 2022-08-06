package mongo_model

import (
	"context"
	"github.com/XC-Zero/yinwan/pkg/utils/convert"
)

// Assemble 组装
type Assemble struct {
	BasicModel         `bson:"inline"`
	BookNameInfo       `bson:"-"`
	AssembleName       *string              `form:"assemble_name,omitempty" json:"assemble_name,omitempty" bson:"assemble_name,omitempty" cn:"组装单名称"`
	AssembleOwnerID    *int                 `form:"assemble_owner_id,omitempty" json:"assemble_owner_id,omitempty" bson:"assemble_owner_id,omitempty" cn:"组装负责人编号"`
	AssembleOwnerName  *string              `form:"assemble_owner_name,omitempty" json:"assemble_owner_name,omitempty" bson:"assemble_owner_name,omitempty" cn:"组装负责人名称"`
	AssembleTemplateID int                  `form:"assemble_template_id" json:"assemble_template_id" bson:"assemble_template_id" cn:"组装拆卸模板编号"`
	AssembleContent    []stockRecordContent `form:"assemble_content" json:"assemble_content" bson:"assemble_content" cn:"组装单内容"`
	Remark             *string              `form:"remark,omitempty" json:"remark,omitempty" bson:"remark,omitempty" cn:"备注"`
}

func (a Assemble) Mapping() map[string]interface{} {
	ma := mapping{
		"settings": mapping{},
		"mappings": mapping{
			"properties": mapping{
				"rec_id": mapping{
					"type": "keyword",
				},
				"assemble_content": mapping{
					"type":            "text",
					"analyzer":        IK_SMART,
					"search_analyzer": IK_SMART,
				},
				"assemble_name": mapping{
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
				"assemble_owner_name": mapping{
					"type":            "text",
					"analyzer":        IK_SMART,
					"search_analyzer": IK_SMART,
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

func (a Assemble) ToESDoc() map[string]interface{} {
	return map[string]interface{}{
		"rec_id":              a.RecID,
		"created_at":          a.CreatedAt,
		"remark":              a.Remark,
		"assemble_content":    convert.StructSliceToTagString(a.AssembleContent, "cn"),
		"assemble_owner_name": a.AssembleOwnerName,
		"assemble_name":       a.AssembleName,
		"book_name":           a.BookName,
		"book_name_id":        a.BookNameID,
	}
}

func (a Assemble) TableName() string {
	return "assembles"
}
func (a Assemble) TableCnName() string {
	return "组装拆卸"
}

func (a *Assemble) BeforeInsert(ctx context.Context) error {
	return nil
}

func (a *Assemble) BeforeUpdate(ctx context.Context) error {
	return nil
}

func (a *Assemble) BeforeRemove(ctx context.Context) error {
	return nil
}
