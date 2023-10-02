package app

type ActivationEmail struct {
	To string
	//other fields
	OtherFields string
}

type EmailService interface {
	SendActivationEmail(email ActivationEmail) (*ActivationEmail, error)
}
