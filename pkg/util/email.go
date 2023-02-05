package util

import (
	"fmt"
	"net/smtp"

	"github.com/jordan-wright/email"
)

type IEmail interface {
	Host() string
	Port() int
	Secret() string
	From() string
	Recipients() []string
	Subject() string
	Body() []byte
}

// SendEmail 发送邮件
func SendEmail(mail IEmail) error {
	e := email.NewEmail()
	e.From = mail.From()
	e.To = mail.Recipients()
	e.Subject = mail.Subject()
	e.Text = mail.Body()
	return e.Send(fmt.Sprintf("%s:%d", mail.Host(), mail.Port()),
		smtp.PlainAuth("", mail.From(), mail.Secret(), mail.Host()))
}
