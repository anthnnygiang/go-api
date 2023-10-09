package app

type ActivationEmail struct {
	To string
	//other fields
	ActivationURL string
	Token
}

type EmailService interface {
	SendActivationEmail(email ActivationEmail) (*ActivationEmail, error)
}
