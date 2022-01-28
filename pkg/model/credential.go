package model

import "github.com/XC-Zero/yinwan/pkg/tools/math_plus"

// Credential 财务凭证
type Credential struct {
	BasicModel
	CredentialLabel     string
	CredentialOwnerID   int
	CredentialOwnerName string
	CredentialEvents    []CredentialEvent
}

// CalculateTotalAmount todo 通过凭证中各条目计算总金额
func (c *Credential) CalculateTotalAmount() {

}

// CredentialEvent 凭证条目
type CredentialEvent struct {
	IncreaseEvent []EventItem
	DecreaseEvent []EventItem
}

// EventItem 具体条目
type EventItem struct {
	// 变动类型
	EventItemType string
	// 变动对象
	EventItemObject string
	// 变动金额
	EventItemAmount math_plus.Fraction
}
