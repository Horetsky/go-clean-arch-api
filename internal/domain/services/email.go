package services

type EmailService interface {
	SendVerificationEmail(to string) error
}
