package controllers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/lee-lou2/api/app/notify/models"
	"github.com/lee-lou2/api/pkg/http"
	"github.com/lee-lou2/api/platform/cache"
	"log"
	netUrl "net/url"
	"os"
)

// SimpleMessageTemplate 기본 메시지 템플릿
func SimpleMessageTemplate(message string) string {
	templateObject := map[string]interface{}{
		"object_type": "text",
		"text":        message,
		"link": map[string]string{
			"web_url":        "",
			"mobile_web_url": "",
		},
	}
	templateObjectBytes, _ := json.Marshal(templateObject)
	return string(templateObjectBytes)
}

// SendKaKaOToMe 나에게 카카오톡 보내기
func SendKaKaOToMe(message models.Message) {
	go sendKaKaOToMe(message.Message)
}

// SendKaKaOToFriend 친구들에게 카카오톡 보내기
func SendKaKaOToFriend(message models.Message) {
	// 메시지 전송
	if len(message.Targets) == 0 {
		go sendKaKaOToFriend(message.Message, "")
	}
	for _, target := range message.Targets {
		friendUUID := target
		if friendUUID == "" {
			friendUUID = os.Getenv("KAKAO_API_DEFAULT_FRIEND_UUID")
		}
		go sendKaKaOToFriend(
			message.Message,
			friendUUID,
		)
	}
}

// sendKaKaOToMe 나에게 카카오톡 보내기
func sendKaKaOToMe(message string) {
	templateObject := SimpleMessageTemplate(message)

	url := "https://kapi.kakao.com/v2/api/talk/memo/default/send"

	redis := cache.RedisClient()
	accessToken, err := redis.GetValue("kakao_access_token")
	if err != nil {
		log.Println(err)
		return
	}

	params := netUrl.Values{}
	params.Add("template_object", templateObject)

	_, err = http.Request(
		"POST",
		url,
		bytes.NewBufferString(params.Encode()),
		&http.Header{Key: "Content-Type", Value: "application/x-www-form-urlencoded"},
		&http.Header{Key: "Authorization", Value: "Bearer " + accessToken},
	)
	if err != nil {
		log.Println(err)
	}
}

// sendKaKaOToFriend 친구에게 카카오톡 보내기
func sendKaKaOToFriend(message, friendUuid string) {
	if friendUuid == "" {
		_friendUuid := os.Getenv("KAKAO_API_DEFAULT_FRIEND_UUID")
		friendUuid = _friendUuid
	}
	templateObject := SimpleMessageTemplate(message)

	url := "https://kapi.kakao.com/v1/api/talk/friends/message/default/send"

	redis := cache.RedisClient()
	accessToken, err := redis.GetValue("kakao_access_token")
	if err != nil {
		log.Println(err)
		return
	}

	params := netUrl.Values{}
	params.Add("receiver_uuids", fmt.Sprintf(`["%s"]`, friendUuid))
	params.Add("template_object", templateObject)

	_, err = http.Request(
		"POST",
		url,
		bytes.NewBufferString(params.Encode()),
		&http.Header{Key: "Content-Type", Value: "application/x-www-form-urlencoded"},
		&http.Header{Key: "Authorization", Value: "Bearer " + accessToken},
	)
	if err != nil {
		log.Println(err)
	}
}
