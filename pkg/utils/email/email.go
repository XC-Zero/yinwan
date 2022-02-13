package email

import (
	"github.com/XC-Zero/yinwan/internal/config"
	"github.com/jordan-wright/email"
	"net/smtp"
	"strings"
)

// SendEmail 发送邮件 ，仅支持SMTP服务
func SendEmail(subject, content string, to ...string) error {
	emailConfig := config.CONFIG.ApiConfig.EmailConfig
	e := email.NewEmail()
	e.From = emailConfig.SenderEmail
	e.To = to
	e.Subject = subject
	e.Text = []byte("YinWan 来信！")
	e.HTML = []byte(content)
	err := e.Send(emailConfig.EmailServerAddr, smtp.PlainAuth("", emailConfig.SenderEmail,
		emailConfig.EmailSecret, strings.Split(emailConfig.EmailServerAddr, ":")[0]))
	if err != nil {
		return err
	}
	return nil
}
