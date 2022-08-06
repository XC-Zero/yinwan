package mongo_model

import (
	"context"
	"github.com/XC-Zero/yinwan/pkg/utils/convert"
)

// Disassemble 拆卸
type Disassemble struct {
	BasicModel           `bson:"inline"`
	BookNameInfo         `bson:"-"`
	DisassembleName      *string              `form:"disassemble_name,omitempty" json:"disassemble_name,omitempty" bson:"disassemble_name,omitempty"`
	DisassembleOwnerID   *int                 `form:"disassemble_owner_id,omitempty" json:"disassemble_owner_id,omitempty" bson:"disassemble_owner_id,omitempty"`
	DisassembleOwnerName *string              `form:"disassemble_owner_name,omitempty" json:"disassemble_owner_name,omitempty" bson:"disassemble_owner_name,omitempty"`
	AssembleTemplateID   int                  `form:"assemble_template_id" json:"assemble_template_id" bson:"assemble_template_id"`
	DisassembleContent   []stockRecordContent `form:"disassemble_content" json:"disassemble_content" bson:"disassemble_content"`
	Remark               *string              `form:"remark,omitempty" json:"remark,omitempty" bson:"remark,omitempty"`
}

func (d Disassemble) Mapping() map[string]interface{} {
	ma := mapping{
		"settings": mapping{},
		"mappings": mapping{
			"properties": mapping{
				"rec_id": mapping{
					"type": "keyword",
				},
				"disassemble_content": mapping{
					"type":            "text",
					"analyzer":        IK_SMART,
					"search_analyzer": IK_SMART,
				},
				"disassemble_name": mapping{
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
				"disassemble_owner_name": mapping{
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

func (d Disassemble) ToESDoc() map[string]interface{} {
	return map[string]interface{}{
		"rec_id":                 d.RecID,
		"created_at":             d.CreatedAt,
		"remark":                 d.Remark,
		"disassemble_content":    convert.StructToTagString(d.DisassembleContent, "cn"),
		"disassemble_owner_name": d.DisassembleOwnerName,
		"disassemble_name":       d.DisassembleName,
		"book_name":              d.BookName,
		"book_name_id":           d.BookNameID,
	}
}

func (d Disassemble) TableName() string {
	return "disassembles"
}
func (d Disassemble) TableCnName() string {
	return "拆卸"
}

func (d *Disassemble) BeforeInsert(ctx context.Context) error {
	return nil
}

func (d *Disassemble) BeforeUpdate(ctx context.Context) error {
	return nil
}

func (d *Disassemble) BeforeRemove(ctx context.Context) error {
	return nil
}
