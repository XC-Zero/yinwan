package mongo_model

type CredentialTemplate struct {
	BasicModel             `bson:"inline"`
	CredentialTemplateName string `bson:"credential_template_name" json:"credential_template_name" form:"credential_template_name"`
	CredentialMappingList  []CredentialMapping
}

type CredentialMapping struct {
	OriginalColumnEn string `bson:"original_column_en" json:"original_column_en" form:"original_column_en"`
	OriginalColumnCn string `bson:"original_column_cn" json:"original_column_cn" form:"original_column_cn"`
	TargetColumnEn   string `bson:"target_column_en" json:"target_column_en" form:"target_column_en"`
	TargetColumnCn   string `bson:"target_column_cn" json:"target_column_cn" form:"target_column_cn"`
}

func (c CredentialTemplate) TableCnName() string {
	return "凭证模板"
}

func (c CredentialTemplate) TableName() string {
	return "credential_templates"
}
