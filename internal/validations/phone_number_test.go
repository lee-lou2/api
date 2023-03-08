package validations

import "testing"

func TestNormalizePhoneNumber(t *testing.T) {
	testCases := []struct {
		input          string
		expectedOutput string
		expectedErr    error
	}{
		// 정상적인 입력값들
		{"01012345678", "01012345678", nil},
		{"010-1234-5678", "01012345678", nil},
		{"+82-10-1234-5678", "01012345678", nil},
		{"+8201012345678", "01012345678", nil},
		{"82-10-1234-5678", "01012345678", nil},
		{"10-1234-5678", "01012345678", nil},

		// 입력값이 비어있는 경우
		{"", "", &ErrorString{"휴대폰 번호는 10~11자리여야 합니다: "}},

		// "-"와 "+" 이외의 문자가 포함된 경우
		{"0101234a5678", "", &ErrorString{"휴대폰 번호는 숫자로만 이루어져야 합니다: 0101234a5678"}},

		// 전화번호의 길이가 10자리 미만인 경우
		{"010123456", "", &ErrorString{"휴대폰 번호는 10~11자리여야 합니다: 010123456"}},

		// 전화번호의 길이가 11자리 초과인 경우
		{"010123456789", "", &ErrorString{"휴대폰 번호는 10~11자리여야 합니다: 010123456789"}},
	}

	for _, tc := range testCases {
		output, err := NormalizePhoneNumber(tc.input)
		if output != tc.expectedOutput {
			t.Errorf("normalizePhoneNumber(%v) = %v, want %v", tc.input, output, tc.expectedOutput)
		}
		if tc.expectedErr != nil {
			if err == nil {
				t.Errorf("normalizePhoneNumber(%v) err = nil, want %v", tc.input, tc.expectedErr)
			} else if err.Error() != tc.expectedErr.Error() {
				t.Errorf("normalizePhoneNumber(%v) err = %v, want %v", tc.input, err, tc.expectedErr)
			}
		} else {
			if err != nil {
				t.Errorf("normalizePhoneNumber(%v) err = %v, want nil", tc.input, err)
			}
		}
	}
}

type ErrorString struct {
	s string
}

func (e *ErrorString) Error() string {
	return e.s
}
