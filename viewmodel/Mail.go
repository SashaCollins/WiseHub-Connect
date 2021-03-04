package viewmodel

import (
	"bytes"
	"errors"
	"fmt"
	"net/smtp"
	"os"
	"path/filepath"
	"text/template"
)

type Credentials struct {
	Username string `env:"MAIL_USER"`
	Password string `env:"MAIL_PASSWORD"`
}

type Server struct {
	Host string `env:"MAIL_HOST"`
	Port string `env:"MAIL_PORT"`
}

type Mail struct {}

func parseTemplate(templateFileName string, data interface{}) (string, error) {
	_, err := filepath.Abs(templateFileName)
	if err != nil {
		return "", errors.New("invalid template name")
	}
	t, err := template.New("validate").Parse("Dear WiseHub-User,\n\n    To authenticate your account\n\n    {{.Url}}\n\n    click here.\n\nAlternatively, you can paste the following link in your browser's address bar.\n\n\nIf you have not requested a authentication simply ignore this message.\n\nSincerely,\nThe WiseHub Team")
	if err != nil {
		return "", err
	}
	buf := new(bytes.Buffer)
	if err = t.Execute(buf, data); err != nil {
		return "", err
	}
	body := buf.String()
	return body, nil
}

var emailAuth smtp.Auth

func (m *Mail) SendEmailSMTP(to []string, data interface{}, temp string) (bool, error) {
	emailHost := os.Getenv("MAIL_HOST")
	emailFrom := os.Getenv("MAIL_USERNAME")
	emailPassword := os.Getenv("MAIL_PASSWORD")
	emailPort := os.Getenv("MAIL_PORT")

	emailAuth = smtp.PlainAuth("", emailFrom, emailPassword, emailHost)

	emailBody, err := parseTemplate(temp, data)
	if err != nil {
		return false, errors.New("unable to parse email template")
	}

	mime := "MIME-version: 1.0;\nContent-Type: text/plain; charset=\"UTF-8\";\n\n"
	subject := "Subject: WiseHub-Connect\n"
	msg := []byte(subject + mime + "\n" + emailBody)
	addr := fmt.Sprintf("%s:%s", emailHost, emailPort)

	if err := smtp.SendMail(addr, emailAuth, emailFrom, to, msg); err != nil {
		return false, err
	}
	return true, nil
}

