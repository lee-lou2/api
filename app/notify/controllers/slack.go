package controllers

import (
	"github.com/lee-lou2/api/app/notify/models"
	"github.com/slack-go/slack"
	"log"
	"os"
)

// SendSlack 슬랙 일괄 전송
func SendSlack(message models.Message) {
	// 기본 변수
	accessToken := os.Getenv("SLACK_API_TOKEN")

	// 메시지 전송
	for _, target := range message.Targets {
		go sendSlack(
			accessToken,
			target,
			message.Message,
		)
	}
}

// sendSlack 슬랙 전송
func sendSlack(accessToken, channel, message string) {
	// APP 생성
	api := slack.New(accessToken)
	if _, _, err := api.PostMessage(
		channel,
		slack.MsgOptionText(message, false),
		slack.MsgOptionAsUser(true),
	); err != nil {
		log.Println(err)
	}
}
