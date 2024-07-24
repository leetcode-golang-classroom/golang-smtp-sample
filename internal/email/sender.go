package email

import (
	"fmt"
	"net/smtp"

	"github.com/leetcode-golang-classroom/golang-smtp-sample/internal/config"
)

type Sender struct {
	cfg *config.Config
}

func NewSender(cfg *config.Config) *Sender {
	return &Sender{cfg}
}

func (emailer *Sender) SendHtmlEmail(to []string, subject string, htmlBody string) error {
	auth := smtp.PlainAuth(
		"",
		emailer.cfg.FromEmail,
		emailer.cfg.SmtpSecret,
		emailer.cfg.FromEmailSmtp,
	)
	headers := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";"
	message := fmt.Sprintf("Subject: %s\n%s\n\n%s", subject, headers, htmlBody)
	return smtp.SendMail(
		emailer.cfg.SmtpAddr,
		auth,
		emailer.cfg.FromEmail,
		to,
		[]byte(message),
	)
}
func (emailer *Sender) SendEmail(to []string, subject string, body string) error {
	auth := smtp.PlainAuth(
		"",
		emailer.cfg.FromEmail,
		emailer.cfg.SmtpSecret,
		emailer.cfg.FromEmailSmtp,
	)
	message := fmt.Sprintf("Subject: %s\n%s", subject, body)
	return smtp.SendMail(
		emailer.cfg.SmtpAddr,
		auth,
		emailer.cfg.FromEmail,
		to,
		[]byte(message),
	)
}
