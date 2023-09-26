package app

type Email struct {
	From     string
	To       string
	Subject  string
	HtmlBody string
}

type EmailService interface {
	SendActivationEmail(email string) error
}
