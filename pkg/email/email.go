package email

import (
	"fmt"
	"net/smtp"
	"strings"
	"uc/pkg/nacos"
)

const MAIL_TYPE_HTML = "html"
const MAIL_TYPE_TEXT = "text"

type Email struct {
	Host     string
	Port     string
	Username string
	Password string
	auth     smtp.Auth
}

var MyEmail = new(Email)

func Init() {
	var data = nacos.Config
	fmt.Println(data)
	MyEmail = &Email{
		Host:     nacos.Config.Email.Host,
		Port:     nacos.Config.Email.Port,
		Username: nacos.Config.Email.Username,
		Password: nacos.Config.Email.Password,
	}
	MyEmail.auth = smtp.PlainAuth("", MyEmail.Username, MyEmail.Password, MyEmail.Host)
}

func (e *Email) SendEmail(subject string, to []string, mailType string, message string) error {
	var contentType = "text/plain; charset=UTF-8"
	if mailType == MAIL_TYPE_HTML {
		contentType = "text/html; charset=UTF-8"
	}
	var msg = "To: " + strings.Join(to, ",") + "\r\n" +
		"From: " + e.Username + "\r\n" +
		"Subject: " + subject + "\r\n" +
		"Content-Type: " + contentType + "\r\n\r\n" +
		message + "\r\n"
	return smtp.SendMail(fmt.Sprintf("%s:%s", e.Host, e.Port), e.auth, e.Username, to, []byte(msg))
}
