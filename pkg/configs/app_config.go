package configs

import (
	_ "github.com/joho/godotenv/autoload"
	"github.com/lee-lou2/hub/pkg/core"
	"github.com/lee-lou2/hub/platform/aws"
)

// SetApplicationConfig 어플리케이션 설정
func SetApplicationConfig() {
	// 환경 변수 설정
	aws.SetEnvironments()
	// 로그 설정
	core.SetLogger()
}
