package types

type Sender interface {
	SendHtmlEmail(to []string, subject string, htmlBody string) error
	SendEmail(to []string, subject string, body string) error
}
