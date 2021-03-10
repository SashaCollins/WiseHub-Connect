package viewmodel

import (
	"bytes"
	"crypto/tls"
	"errors"
	"gopkg.in/gomail.v2"
	"log"
	"path/filepath"
	"strconv"
	"text/template"
)

type MailServer struct {

	// Server credentials
	ServerHost string
	ServerPort string
	LoginName string
	LoginPassword string

	// Url for validation
	BodyUrl string
}

/*
Parse the mail body.
 */
func parseTemplate(templateFileName string, data interface{}) (string, error) {
	_, err := filepath.Abs(templateFileName)
	if err != nil {
		return "", errors.New("invalid template name")
	}
	t, err := template.New("validate").Parse("Dear WiseHub-User,\n\n    To authenticate your account\n\n    {{.Url}}\n\n    click here.\n\nAlternatively, you can paste the link in your browser's address bar.\n\n\nIf you have not requested a authentication simply ignore this message.\n\nSincerely,\nThe WiseHub Team")
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

/*
Send email per smtp.
 */
func (ms *MailServer) SendEmailSMTP(to string, data interface{}, temp string) (bool, error) {
	//Generate email body
	emailBody, err := parseTemplate(temp, data)
	if err != nil {
		return false, errors.New("unable to parse email template")
	}

	// Set new connection to email server
	port, _ := strconv.Atoi(ms.ServerPort)
	dialer := gomail.NewDialer(ms.ServerHost, port, ms.LoginName, ms.LoginPassword)
	dialer.TLSConfig = &tls.Config{InsecureSkipVerify: true}
	s, err := dialer.Dial()
	if err != nil {
		log.Println(err)
		return false, err
	}

	// Build message
	m := gomail.NewMessage()
	m.SetHeader("From", ms.LoginName)
	m.SetHeader("To", to)
	m.SetAddressHeader("Cc", ms.LoginName, "Wisehub-Connect")
	m.SetHeader("Subject", "[WiseHub-Connect] EMail Validator")
	m.SetBody("text/plain", emailBody)

	// Send the email
	if err := gomail.Send(s, m); err != nil {
		log.Printf("Could not send email to %s: %v", to, err)
		return false, err
	}
	return true, nil
}
