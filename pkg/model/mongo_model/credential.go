package mongo_model

import (
	"fmt"
	_const "github.com/XC-Zero/yinwan/pkg/const"
	"github.com/XC-Zero/yinwan/pkg/utils/convert"
)

// Credential 财务凭证
type Credential struct {
	BasicModel               `bson:"inline"`
	BookNameInfo             `bson:"-"`
	CredentialName           *string                  `json:"credential_name,omitempty" form:"credential_name,omitempty" bson:"credential_name,omitempty" cn:"凭证标题"`
	CredentialOwnerID        int                      `json:"credential_owner_id" form:"credential_owner_id" bson:"credential_owner_id" cn:"凭证责任人编号"`
	CredentialOwnerName      string                   `json:"credential_owner_name" form:"credential_owner_name" bson:"credential_owner_name" cn:"凭证责任人名称"`
	CredentialEvents         []CredentialEvent        `json:"credential_events" form:"credential_events" bson:"credential_events" cn:"凭证分录"`
	CredentialMakerID        *int                     `json:"credential_maker_id" form:"credential_maker_id" bson:"credential_maker_id" cn:"制单人编号"`
	CredentialMakerName      *string                  `json:"credential_maker_name" form:"credential_maker_name" bson:"credential_maker_name" cn:"制单人名称"`
	CredentialAccountantID   *int                     `json:"credential_accountant_id" form:"credential_accountant_id" bson:"credential_accountant_id" cn:"会计编号"`
	CredentialAccountantName *string                  `json:"credential_accountant_name" form:"credential_accountant_name" bson:"credential_accountant_name" cn:"会计名称"`
	CredentialCashierID      *int                     `json:"credential_cashier_id" form:"credential_cashier_id" bson:"credential_cashier_id" cn:"出纳编号"`
	CredentialCashierName    *string                  `json:"credential_cashier_name" form:"credential_cashier_name" bson:"credential_cashier_name" cn:"出纳名称"`
	CredentialCheckerID      *int                     `json:"credential_checker_id" form:"credential_checker_id" bson:"credential_checker_id" cn:"复核编号"`
	CredentialCheckerName    *string                  `json:"credential_checker_name" form:"credential_checker_name" bson:"credential_checker_name" cn:"复核名称"`
	CredentialStatus         *_const.CredentialStatus `json:"credential_status,omitempty" form:"credential_status,omitempty" bson:"credential_status,omitempty"`
	Remark                   *string                  `json:"remark" form:"remark" bson:"remark" cn:"备注"`
}

func (c Credential) TableCnName() string {
	return "凭证"
}
func (c Credential) TableName() string {
	return "credentials"
}
func (c Credential) Mapping() map[string]interface{} {
	ma := mapping{
		"settings": mapping{},
		"mappings": mapping{
			"properties": mapping{
				"rec_id": mapping{
					"type": "keyword",
				},
				"credential_name": mapping{
					"type":            "text",
					"analyzer":        IK_SMART,
					"search_analyzer": IK_SMART,
				},
				"related_person": mapping{
					"type":            "text",
					"analyzer":        IK_SMART,
					"search_analyzer": IK_SMART,
				},
				"credential_content": mapping{
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

// ToESDoc todo !!!!
func (c Credential) ToESDoc() map[string]interface{} {

	maker, account, cashier, checker := "", "", "", ""
	if c.CredentialMakerName != nil {
		maker = *c.CredentialMakerName
	}
	if c.CredentialCashierName != nil {
		cashier = *c.CredentialCashierName
	}
	if c.CredentialAccountantName != nil {
		account = *c.CredentialAccountantName
	}

	if c.CredentialCheckerName != nil {
		checker = *c.CredentialCheckerName
	}
	return map[string]interface{}{
		"rec_id":             c.RecID,
		"created_at":         c.CreatedAt,
		"remark":             c.Remark,
		"credential_name":    c.CredentialName,
		"related_person":     fmt.Sprintf("制单人:%s  会计:%s  出纳:%s  复核:%s ", maker, account, cashier, checker),
		"credential_content": convert.StructSliceToTagString(c.CredentialEvents, "cn"),
		"book_name":          c.BookName,
		"book_name_id":       c.BookNameID,
	}
}

// CredentialEvent 凭证条目
type CredentialEvent struct {
	Abstract       string `json:"abstract" form:"abstract" bson:"abstract" cn:"摘要"`
	Classify       string `json:"classify" form:"classify" bson:"classify" cn:"科目"`
	DetailClassify string `json:"detail_classify" form:"detail_classify" bson:"detail_classify" cn:"明细科目"`
	LoanAmount     string `json:"loan_amount" form:"loan_amount" bson:"loan_amount" cn:"贷方金额"`
	DebitAmount    string `json:"debit_amount" form:"debit_amount" bson:"debit_amount" cn:"借方金额"`
}
