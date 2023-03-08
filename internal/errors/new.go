package errors

import (
	"log"
)

// Error 오류 스키마
type Error struct {
	Code       int
	StatusCode int
	Message    string
}

// New 오류 생성
func New(err Error, defaultErrors ...error) error {
	// 오류 알람
	err.Alert()

	// 기본 오류들 알람
	for _, _err := range defaultErrors {
		log.Println(_err)
	}
	return &err
}
