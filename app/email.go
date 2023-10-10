package app

type ActivationEmail struct {
	To string
	//other fields
	RawToken string
}

type EmailService interface {
	SendActivationEmail(email ActivationEmail) (*ActivationEmail, error)
}
