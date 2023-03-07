package configs

import (
	"github.com/lee-lou2/hub/pkg/core"
)

// SetApplicationConfig 어플리케이션 설정
func SetApplicationConfig() {
	// 로그 설정
	core.SetLogger()
}
