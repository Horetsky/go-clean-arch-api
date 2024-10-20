package services

import (
	"seeker/internal/domain/dto"
)

type EmailService interface {
	SendVerificationEmail(to string) error
	SendJobApplicationEmail(to string, input dto.SendJobApplicationEmailDTO) error
}
