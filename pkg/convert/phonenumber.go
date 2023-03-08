package convert

import (
	"fmt"
	"regexp"
)

func NormalizePhoneNumber(phoneNumber string) (string, error) {
	// "+"와 "82"를 제거합니다.
	phoneNumber = regexp.MustCompile(`\+|82`).ReplaceAllString(phoneNumber, "")

	// "-"를 제거합니다.
	phoneNumber = regexp.MustCompile(`-`).ReplaceAllString(phoneNumber, "")

	// 숫자가 아닌 문자가 포함되어 있는지 검사합니다.
	if regexp.MustCompile(`[^\d]`).MatchString(phoneNumber) {
		return "", fmt.Errorf("휴대폰 번호는 숫자로만 이루어져야 합니다: %s", phoneNumber)
	}

	// 전화번호의 길이가 10자리 이상이고, 11자리 이하인지 검사합니다.
	if len(phoneNumber) < 10 || len(phoneNumber) > 11 {
		return "", fmt.Errorf("휴대폰 번호는 10~11자리여야 합니다: %s", phoneNumber)
	}

	// "+82"로 시작하는 경우에만 "010"으로 시작하도록 합니다.
	if len(phoneNumber) == 11 && phoneNumber[:3] == "820" {
		phoneNumber = "010" + phoneNumber[3:]
	}

	// 그 외의 경우에는 무조건 "010"으로 시작하도록 합니다.
	if len(phoneNumber) == 10 || (len(phoneNumber) == 11 && phoneNumber[:3] != "010") {
		phoneNumber = "010" + phoneNumber[3:]
	}

	// 숫자만으로 이루어진 휴대폰 번호를 반환합니다.
	return regexp.MustCompile(`[^\d]`).ReplaceAllString(phoneNumber, ""), nil
}
