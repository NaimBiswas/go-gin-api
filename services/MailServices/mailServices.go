package mailServices

import (
	Models "NaimBiswas/go-gin-api/models"
	"fmt"
	"github.com/joho/godotenv"
	"io/ioutil"
	"net/smtp"
	"os"
	"strings"
)

func New(subject string, to []string, path string) (Models.IResponse, error) {
	response, err := configureSMTP(subject, to, path)
	fmt.Println("path:=================", path)
	if err != nil {
		return Models.IResponse{}, err
	}
	return response, nil
}

func configureSMTP(subject string, to []string, path string) (Models.IResponse, error) {
	_ = godotenv.Load()
	smtpServer := os.Getenv("SMTP_SERVER")
	smtpPort := os.Getenv("SMTP_PORT")
	smtpUser := os.Getenv("SMTP_USER")
	smtpPass := os.Getenv("SMTP_PASSWORD")
	from := os.Getenv("FROM")

	mime := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"

	auth := smtp.PlainAuth("", smtpUser, smtpPass, smtpServer)
	serverAddr := smtpServer + ":" + smtpPort

	body, err := loadHTMLFile(path)
	if err != nil {
		fmt.Println("Error loading email template:", err)
	}
	// Prepare the email headers and body
	headers := fmt.Sprintf("From: %s\r\nTo: %s\r\nSubject: %s\r\n%s\r\n", from, strings.Join(to, ", "), subject, mime)
	message := []byte(headers + "\r\n" + body)

	err = smtp.SendMail(serverAddr, auth, from, to, message)
	if err != nil {
		return Models.IResponse{}, err
	}
	fmt.Println("Email Sent!")
	return Models.IResponse{
		Message: "Mail Has been sent successfully",
		Details: "mail Sent to flowing mails " + strings.Join(to, ", "),
	}, nil
}

func loadHTMLFile(filename string) (string, error) {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return "", err
	}
	return string(data), nil
}
