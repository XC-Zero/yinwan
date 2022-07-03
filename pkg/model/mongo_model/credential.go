package mongo_model

import (
	"fmt"
)

// Credential 财务凭证
type Credential struct {
	BasicModel               `bson:"inline"`
	BookNameInfo             `bson:"-"`
	CredentialLabel          string            `json:"credential_label" form:"credential_label" bson:"credential_label" cn:"凭证标题"`
	CredentialOwnerID        int               `json:"credential_owner_id" form:"credential_owner_id" bson:"credential_owner_id" cn:"凭证责任人编号"`
	CredentialOwnerName      string            `json:"credential_owner_name" form:"credential_owner_name" bson:"credential_owner_name" cn:"凭证责任人名称"`
	CredentialEvents         []CredentialEvent `json:"credential_events" form:"credential_events" bson:"credential_events" cn:"凭证分录"`
	CredentialMakerID        *int              `json:"credential_maker_id" form:"credential_maker_id" bson:"credential_maker_id" cn:"制单人编号"`
	CredentialMakerName      *string           `json:"credential_maker_name" form:"credential_maker_name" bson:"credential_maker_name" cn:"制单人名称"`
	CredentialAccountantID   *int              `json:"credential_accountant_id" form:"credential_accountant_id" bson:"credential_accountant_id" cn:"会计编号"`
	CredentialAccountantName *string           `json:"credential_accountant_name" form:"credential_accountant_name" bson:"credential_accountant_name" cn:"会计名称"`
	CredentialCashierID      *int              `json:"credential_cashier_id" form:"credential_cashier_id" bson:"credential_cashier_id" cn:"出纳编号"`
	CredentialCashierName    *string           `json:"credential_cashier_name" form:"credential_cashier_name" bson:"credential_cashier_name" cn:"出纳名称"`
	CredentialCheckerID      *int              `json:"credential_checker_id" form:"credential_checker_id" bson:"credential_checker_id" cn:"复核编号"`
	CredentialCheckerName    *string           `json:"credential_checker_name" form:"credential_checker_name" bson:"credential_checker_name" cn:"复核名称"`
	Remark                   *string           `json:"remark" form:"remark" bson:"remark" cn:"备注"`
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
	var credentialContent string
	for _, event := range c.CredentialEvents {
		for _, decrease := range event.DecreaseEvent {
			credentialContent += fmt.Sprintf(
				"贷: 变动类型 :%s 变动对象: %s 变动金额:%s \n",
				decrease.EventItemType,
				decrease.EventItemObject,
				decrease.EventItemAmount)
		}

		for _, increase := range event.IncreaseEvent {
			credentialContent += fmt.Sprintf(
				"借: 变动类型 :%s 变动对象: %s 变动金额:%s \n",
				increase.EventItemType,
				increase.EventItemObject,
				increase.EventItemAmount)
		}
	}

	return map[string]interface{}{
		"rec_id":             c.RecID,
		"created_at":         c.CreatedAt,
		"remark":             c.Remark,
		"related_person":     fmt.Sprintf("制单人:%s  会计:%s  出纳:%s  复核:%s ", maker, account, cashier, checker),
		"credential_content": credentialContent,
		//"stock_in_owner":             m.StockInRecordOwnerName,
		"book_name":    c.BookName,
		"book_name_id": c.BookNameID,
	}
}

// CredentialEvent 凭证条目
type CredentialEvent struct {
	// 借
	IncreaseEvent []EventItem `json:"increase_event" bson:"increase_event"`
	// 贷
	DecreaseEvent []EventItem `json:"decrease_event" bson:"decrease_event"`
}

// EventItem 具体条目
type EventItem struct {
	// 变动类型
	EventItemType string `json:"event_item_type" bson:"event_item_type"`
	// 变动对象
	EventItemObject string `json:"event_item_object" bson:"event_item_object"`
	// 变动金额
	EventItemAmount string `json:"event_item_amount" bson:"event_item_amount"`
}

// CalculateTotalAmount todo 通过凭证中各条目计算总金额
// todo 需不需要，有待商榷
func (c *Credential) CalculateTotalAmount() {

}

// CreateCredential 创建财务凭证
//  todo
func CreateCredential() {

}

//
//// TransferToInvoice 转为单据
//// todo
//func (c *Credential) TransferToInvoice(invoiceType _const.InvoiceType) _interface.Invoice {
//	return nil
//}
