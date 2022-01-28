package model

// Credential 财务凭证
type Credential struct {
	BasicModel
	// 凭证标题
	CredentialLabel string
	// 凭证责任人ID
	CredentialOwnerID int
	// 凭证责任人名称
	CredentialOwnerName string
	// 凭证分录
	CredentialEvents []CredentialEvent
	//	凭证备注
	CredentialRemark string
}

// CredentialEvent 凭证条目
type CredentialEvent struct {
	// 借
	IncreaseEvent []EventItem
	// 贷
	DecreaseEvent []EventItem
}

// EventItem 具体条目
type EventItem struct {
	// 变动类型
	EventItemType string
	// 变动对象
	EventItemObject string
	// 变动金额
	EventItemAmount string
}

// CalculateTotalAmount todo 通过凭证中各条目计算总金额
func (c *Credential) CalculateTotalAmount() {

}
