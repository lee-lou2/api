package validations

import (
	"fmt"
	"regexp"
)

func NormalizePhoneNumber(phoneNumber string) (string, error) {
	// "+"와 "82"를 제거
	phoneNumber = regexp.MustCompile(`\+|82`).ReplaceAllString(phoneNumber, "")

	// "-"를 제거
	phoneNumber = regexp.MustCompile(`-`).ReplaceAllString(phoneNumber, "")

	// 숫자가 아닌 문자가 포함되어 있는지 검사합니다.
	if regexp.MustCompile(`[^\d]`).MatchString(phoneNumber) {
		return "", fmt.Errorf("휴대폰 번호는 숫자로만 이루어져야 합니다: %s", phoneNumber)
	}

	// 전화번호의 길이가 10자리 이상이고, 11자리 이하인지 검사합니다.
	if len(phoneNumber) < 10 || len(phoneNumber) > 11 {
		return "", fmt.Errorf("휴대폰 번호는 10~11자리여야 합니다: %s", phoneNumber)
	}

	// 숫자 외 문자가 있는 경우 제거
	phoneNumber = regexp.MustCompile(`[^\d]`).ReplaceAllString(phoneNumber, "")

	// "10"로 시작하는 경우에만 "010"으로 시작하도록 합니다.
	if (len(phoneNumber) == 9 || len(phoneNumber) == 10) && phoneNumber[:1] == "1" {
		phoneNumber = "0" + phoneNumber
	}

	// 휴대폰 번호 정규식에 맞는지 확인
	if !regexp.MustCompile(`^01[0-9]{1}[0-9]{3,4}[0-9]{4}$`).MatchString(phoneNumber) {
		return "", fmt.Errorf("휴대폰번호가 올바르지 않습니다: %s", phoneNumber)
	}

	// 숫자만으로 이루어진 휴대폰 번호를 반환합니다.
	return regexp.MustCompile(`[^\d]`).ReplaceAllString(phoneNumber, ""), nil
}
