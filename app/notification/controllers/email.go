package controllers

import (
	"github.com/lee-lou2/hub/app/notification/models"
	"gopkg.in/gomail.v2"
	"log"
	"os"
	"strconv"
)

// SendEmail 이메일 일괄 전송
func SendEmail(message models.Message) {
	// 기본 변수
	fromEmail := os.Getenv("EMAIL_USER")
	emailUser := os.Getenv("EMAIL_USER")
	emailPassword := os.Getenv("EMAIL_PASSWORD")
	emailHost := os.Getenv("EMAIL_HOST")
	emailPortString := os.Getenv("EMAIL_PORT")
	emailPort, _ := strconv.Atoi(emailPortString)

	// 메시지 전송
	for _, target := range message.Targets {
		go sendEmail(
			fromEmail,
			emailUser,
			emailPassword,
			emailHost,
			target,
			message.Subject,
			message.Message,
			emailPort,
		)
	}
}

// sendEmail 이메일 개별 전송
func sendEmail(
	fromEmail, emailUser, emailPassword, emailHost, to, subject, message string,
	emailPort int,
) {
	// 이메일 발송
	m := gomail.NewMessage()
	m.SetHeader("From", fromEmail)
	m.SetHeader("To", to)
	m.SetHeader("Subject", subject)
	m.SetBody("text/html", message)
	d := gomail.NewDialer(emailHost, emailPort, emailUser, emailPassword)
	if err := d.DialAndSend(m); err != nil {
		log.Println(err)
	}
}
