package email

import (
	"github.com/XC-Zero/yinwan/internal/config"
	"github.com/jordan-wright/email"
	"net/smtp"
	"net/textproto"
)

// SendEmail 仅支持SMTP服务
func SendEmail(to []string, subject, content string) error {
	emailConfig := config.CONFIG.ApiConfig.EmailConfig
	e := &email.Email{
		To:      to,
		From:    emailConfig.SenderEmail,
		Subject: subject,
		Text:    []byte("Text Body is, of course, supported!"),
		HTML:    []byte(content),
		Headers: textproto.MIMEHeader{},
	}
	err := e.Send(emailConfig.EmailServerAddr, smtp.PlainAuth("",
		emailConfig.SenderEmail, emailConfig.EmailSecret, emailConfig.EmailServerAddr))
	if err != nil {
		return err
	}
	return nil
}

//
