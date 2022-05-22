package config

// ApiConfig 外部API配置
type ApiConfig struct {
	EmailConfig        EmailConfig        `json:"email_config" yaml:"email_config"`
	TextMessageConfig  TextMessageConfig  `json:"text_message_config" yaml:"text_message_config"`
	CurrencyRateConfig CurrencyRateConfig `json:"currency_rate_config" yaml:"currency_rate_config"`
}

// TextMessageConfig 短信服务
type TextMessageConfig struct {
	AppID      string `json:"app_id" yaml:"app_id"`
	AppSecret  string `json:"app_secret" yaml:"app_secret"`
	ApiAddress string `json:"api_address" yaml:"api_address"`
}

// EmailConfig 邮件服务
type EmailConfig struct {
	EmailServerAddr string `json:"email_server_addr" yaml:"email_server_addr"`
	SenderEmail     string `json:"sender_email" yaml:"sender_email"`
	EmailSecret     string `json:"email_secret" yaml:"email_secret"`
}

// CurrencyRateConfig 货币汇率服务
type CurrencyRateConfig struct {
	ListURL string `json:"list_url" yaml:"list_url"` // https://sapi.k780.com/?app=finance.rate_curlist&curType=rateRealtime
	RateURL string `json:"rate_url" yaml:"rate_url"`
	AppKey  string `json:"app_key" yaml:"app_key"`
	Sign    string `json:"sign" yaml:"sign"`
}
