package app

import "anthnnygiang/api-template/postmark"

type Email struct {
	From     string
	To       string
	Subject  string
	HtmlBody string
}

type EmailService interface {
	SendActivationEmail(email string) (*postmark.Response, error)
}
