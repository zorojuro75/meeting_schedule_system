package utils

import (
	"fmt"
	"net/smtp"
	"strings"
)

// SendEmail sends a simple text email using SMTP. Callers should run this in a goroutine if async.
func SendEmail(host, port, user, pass, from string, to []string, subject, body string) error {
	addr := fmt.Sprintf("%s:%s", host, port)
	msg := buildMessage(from, to, subject, body)
	auth := smtp.PlainAuth("", user, pass, host)
	return smtp.SendMail(addr, auth, from, to, []byte(msg))
}

func buildMessage(from string, to []string, subject, body string) string {
	headers := []string{
		fmt.Sprintf("From: %s", from),
		fmt.Sprintf("To: %s", strings.Join(to, ",")),
		fmt.Sprintf("Subject: %s", subject),
		"MIME-Version: 1.0",
		"Content-Type: text/plain; charset=UTF-8",
		"",
	}
	return strings.Join(headers, "\r\n") + body
}
