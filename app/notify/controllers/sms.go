package controllers

import (
	"fmt"
	twilio "github.com/kevinburke/twilio-go"
	"github.com/lee-lou2/api/app/notify/models"
	"github.com/lee-lou2/api/internal/validations"
	"log"
	"os"
)

// SenSMS 문자 발송
func SenSMS(message models.Message) {
	// 기본 변수
	fromNumber := os.Getenv("SMS_TWILIO_FROM_NUMBER")
	sid := os.Getenv("SMS_TWILIO_SID")
	token := os.Getenv("SMS_TWILIO_AUTH_TOKEN")

	// 메시지 전송
	for _, target := range message.Targets {
		phone, err := validations.NormalizePhoneNumber(target)
		if err != nil {
			log.Println(err)
			continue
		}
		go sendSMSTwilio(
			sid,
			token,
			fromNumber,
			phone,
			message.Message,
		)
	}
}

// sendSMSTwilio 트윌리오를 이용한 문자 발송
func sendSMSTwilio(sid, token, fromNumber, toNumber, message string) {
	client := twilio.NewClient(sid, token, nil)

	if _, err := client.Messages.SendMessage(
		fromNumber,
		fmt.Sprintf("+82%s", toNumber[1:]),
		message,
		nil,
	); err != nil {
		log.Println(err)
	}
}
