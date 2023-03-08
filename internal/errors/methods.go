package errors

import (
	"fmt"
	"log"
)

// Error 오류 메세지
func (e *Error) Error() string {
	return fmt.Sprintf("%d, [%d] %s", e.Code, e.StatusCode, e.Message)
}

// Alert 오류 알람
func (e *Error) Alert() {
	log.Println(e.Error())
}
