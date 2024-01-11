package email

import (
	"fmt"
	"log"
	"net/smtp"
	"os"
	"strconv"
)

type EmailServiceInterface interface {
	SendEmail(to []string, subject, body string) error
	SendVerificationEmail(to []string, otp, email string) error
}

type SMTPEmailService struct {
	SMTPHost      string
	SMTPPort      int
	Email         string
	Password      string
	EmailUrl      string
	CodeGenerator VerificationCodeGenerator
}

func NewSMTPEmailService() *SMTPEmailService {
	smtpHost := os.Getenv("SMTP_HOST")
	smtpPortStr := os.Getenv("SMTP_PORT")
	smtpPort, err := strconv.Atoi(smtpPortStr)
	if err != nil {
		log.Fatal("SMTP_PORT harus merupakan bilangan bulat")
	}

	email := os.Getenv("EMAIL")
	password := os.Getenv("PASSWORD")
	emailUrl := os.Getenv("EMAIL_URL")

	codeGen := DefaultCodeGenerator{}

	return &SMTPEmailService{
		SMTPHost:      smtpHost,
		SMTPPort:      smtpPort,
		Email:         email,
		Password:      password,
		EmailUrl:      emailUrl,
		CodeGenerator: codeGen,
	}
}

func (s *SMTPEmailService) SendEmail(to []string, subject, body string) error {
	auth := smtp.PlainAuth("", s.Email, s.Password, s.SMTPHost)

	msg := []byte(fmt.Sprintf("To: %s\r\n"+
		"Subject: %s\r\n"+
		"\r\n"+
		"%s\r\n", to[0], subject, body))

	if err := smtp.SendMail(fmt.Sprintf("%s:%d", s.SMTPHost, s.SMTPPort), auth, s.Email, to, msg); err != nil {
		return fmt.Errorf("gagal mengirim email: %s", err)
	}

	return nil
}

func (s *SMTPEmailService) SendVerificationEmail(to []string, otp, email string) error {
	verificationLink := fmt.Sprintf("%s/verify-email?email=%s&code=%s", s.EmailUrl, email, otp)

	subject := "Verifikasi Email"
	body := fmt.Sprintf("Klik tautan ini untuk verifikasi email Anda: \n %s", verificationLink)

	if err := s.SendEmail(to, subject, body); err != nil {
		return err
	}

	fmt.Println("Email verifikasi berhasil dikirim.")
	return nil
}
