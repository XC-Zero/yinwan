package model

// Customer 客户
// todo add gorm tag
type Customer struct {
	BasicModel
	CustomerName          string  `json:"customer_name"`
	CustomerLegalName     *string `json:"customer_legal_name"`
	CustomerAlias         *string `json:"customer_alias"`
	CustomerContact       *string `json:"customer_contact"`
	CustomerContactPhone  *string `json:"customer_contact_phone"`
	CustomerContactWechat *string `json:"customer_contact_wechat"`
	CustomerOwnerID       *int    `json:"customer_owner_id"`
	CustomerOwnerName     string  `json:"customer_owner_name"`
}
