package email

import (
	"github.com/XC-Zero/yinwan/internal/config"
	"github.com/jordan-wright/email"
	"net/smtp"
	"net/textproto"
)

// SendEmail 发送邮件 ，仅支持SMTP服务
// todo 测一下
func SendEmail(to []string, subject, content string) error {
	emailConfig := config.CONFIG.ApiConfig.EmailConfig
	e := &email.Email{
		To:      to,
		From:    emailConfig.SenderEmail,
		Subject: subject,
		Text:    []byte("Yinwan 来信!"),
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
