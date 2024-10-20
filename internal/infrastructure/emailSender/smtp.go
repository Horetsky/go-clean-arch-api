package emailSender

import (
	"bytes"
	_ "embed"
	"fmt"
	"html/template"
	"seeker/internal/app/config"
	"seeker/internal/domain/dto"
	"seeker/internal/domain/services"

	"gopkg.in/gomail.v2"
)

//go:embed templates/verify-email.html
var verificationEmailTemplate string

//go:embed templates/job-application.html
var jobApplicationTemplate string

type smtpSender struct {
	emailFrom string
	password  string
	publicUrl string
}

func NewSmtpSender() services.EmailService {
	cfg := config.Load()

	return &smtpSender{
		emailFrom: cfg.EmailSender.EmailFrom,
		password:  cfg.EmailSender.Password,
		publicUrl: cfg.PublicUrl,
	}
}

func (s *smtpSender) SendVerificationEmail(to string) error {
	message := gomail.NewMessage()
	message.SetHeader("From", s.emailFrom)
	message.SetHeader("To", to)
	message.SetHeader("Subject", "Verify your email")

	content, err := parseTemplateToStr(verificationEmailTemplate, struct {
		Website string
		Name    string
		Link    string
	}{
		Website: "Seeker",
		Name:    to,
		Link:    fmt.Sprintf("%s/auth/verify-email", s.publicUrl),
	})

	if err != nil {
		return err
	}

	message.SetBody("text/html", content)

	return s.send(message)
}

func (s *smtpSender) SendJobApplicationEmail(to string, input dto.SendJobApplicationEmailDTO) error {
	message := gomail.NewMessage()
	message.SetHeader("From", s.emailFrom)
	message.SetHeader("To", to)
	message.SetHeader("Subject", "New job application")

	content, err := parseTemplateToStr(jobApplicationTemplate, struct {
		JobTitle      string
		RecruiterName string
		ApplicantName string
		CompanyName   string
		Link          string
		Website       string
	}{
		JobTitle:      input.JobTitle,
		RecruiterName: input.RecruiterName,
		ApplicantName: input.ApplicantName,
		CompanyName:   input.CompanyName,
		Website:       "Seeker",
		Link:          fmt.Sprintf("%s/auth/verify-email", s.publicUrl),
	})

	if err != nil {
		return err
	}

	message.SetBody("text/html", content)

	return s.send(message)
}

func (s *smtpSender) send(msg *gomail.Message) error {
	d := gomail.NewDialer("smtp.gmail.com", 587, s.emailFrom, s.password)
	return d.DialAndSend(msg)
}

func parseTemplateToStr(file string, data any) (string, error) {
	var buffer bytes.Buffer

	tmpl, err := template.New("email").Parse(file)
	if err != nil {
		return "", err
	}

	err = tmpl.Execute(&buffer, data)

	if err != nil {
		return "", err
	}

	return buffer.String(), nil
}
