package kakao

import (
	"bytes"
	"fmt"
	"github.com/lee-lou2/api/pkg/convert"
	"github.com/lee-lou2/api/pkg/http"
	"github.com/lee-lou2/api/platform/cache"
	"log"
	netUrl "net/url"
	"os"
)

// setKaKaOToken 카카오톡 토큰 설정
func setKaKaOToken(resp string) error {
	tokenValue, err := convert.StringToMap(resp)
	if err != nil {
		return err
	}

	if _, hasError := tokenValue["error_code"]; hasError {
		// 오류
		err := fmt.Errorf("토큰 조회를 실패하였습니다 : %s", tokenValue["error_code"].(string))
		return err
	}
	if _, hasAccessKey := tokenValue["access_token"]; hasAccessKey {
		redis := cache.RedisClient()
		if _, hasRefreshKey := tokenValue["refresh_token"]; hasRefreshKey {
			// 리프레시 토큰 저장
			if err := redis.SetValue(
				"kakao_refresh_token",
				tokenValue["refresh_token"].(string),
				int(tokenValue["refresh_token_expires_in"].(float64)),
			); err != nil {
				return err
			}
		}
		// 액세스 토큰 저장
		if err := redis.SetValue(
			"kakao_access_token",
			tokenValue["access_token"].(string),
			int(tokenValue["expires_in"].(float64)),
		); err != nil {
			return err
		}
		log.Printf("카카오 엑세스 토큰 발급 완료, 토큰 : %s\n", tokenValue["access_token"].(string))
	} else {
		err := fmt.Errorf("토큰 값이 포함되어있지 않습니다")
		return err
	}
	return nil
}

// CreateKaKaOToken 카카오 토큰 생성
func CreateKaKaOToken() error {
	clientId := os.Getenv("KAKAO_API_CLIENT_ID")
	redirectUri := os.Getenv("KAKAO_API_REDIRECT_URI")
	code := os.Getenv("KAKAO_API_CODE")

	url := "https://kauth.kakao.com/oauth/token"

	params := netUrl.Values{}
	params.Add("grant_type", "authorization_code")
	params.Add("client_id", clientId)
	params.Add("redirect_uri", redirectUri)
	params.Add("code", code)

	resp, _ := http.Request(
		"POST",
		url,
		bytes.NewBufferString(params.Encode()),
		&http.Header{Key: "Content-Type", Value: "application/x-www-form-urlencoded"},
	)
	if err := setKaKaOToken(resp.Body); err != nil {
		return err
	}
	return nil
}

// RefreshKaKaoToken 카카오 토큰 재발급
func RefreshKaKaoToken() error {
	clientId := os.Getenv("KAKAO_API_CLIENT_ID")
	redis := cache.RedisClient()
	refreshToken, err := redis.GetValue("kakao_refresh_token")
	if err != nil {
		return err
	}

	url := "https://kauth.kakao.com/oauth/token"

	params := netUrl.Values{}
	params.Add("grant_type", "refresh_token")
	params.Add("client_id", clientId)
	params.Add("refresh_token", refreshToken)

	resp, _ := http.Request(
		"POST",
		url,
		bytes.NewBufferString(params.Encode()),
	)
	if err := setKaKaOToken(resp.Body); err != nil {
		return err
	}
	return nil
}
