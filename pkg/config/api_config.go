package config

// ApiConfig 外部API配置
type ApiConfig struct {
	EmailConfig       EmailConfig       `json:"email_config"`
	TextMessageConfig TextMessageConfig `json:"text_message_config"`
}

// TextMessageConfig 短信服务
type TextMessageConfig struct {
	AppID      string `json:"app_id"`
	AppSecret  string `json:"app_secret"`
	ApiAddress string `json:"api_address"`
}

// EmailConfig 邮件服务
type EmailConfig struct {
	EmailServerAddr string `json:"email_server_addr"`
	SenderEmail     string `json:"sender_email"`
	EmailSecret     string `json:"email_secret"`
}
