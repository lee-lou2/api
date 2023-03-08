package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/lee-lou2/hub/pkg/configs"
	"github.com/lee-lou2/hub/pkg/core"
	"github.com/lee-lou2/hub/pkg/middleware"
	"github.com/lee-lou2/hub/pkg/routes/restapi"
	"os"
)

// @title API
// @version 1.0
// @description This is an auto-generated API Docs.
// @termsOfService http://swagger.io/terms/
// @contact.name API Support
// @contact.email your@mail.com
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @BasePath /api
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
func main() {
	// 기본 설정
	configs.SetApplicationConfig()

	// Fiber 설정 파일 조회
	config := configs.FiberConfig()
	// Fiber 신규 생성
	app := fiber.New(config)

	// 미들웨어
	middleware.FiberMiddleware(app) // Register Fiber's middleware for app.

	// 라우트
	restapi.SwaggerRoute(app)  // Register a route for API Docs (Swagger).
	restapi.MainRoutes(app)    // Register a public routes for app.
	restapi.NotFoundRoute(app) // Register route for 404 Error.

	// 서버 실행
	if os.Getenv("STAGE_STATUS") == "dev" {
		core.StartServer(app)
	} else {
		core.StartServerWithGracefulShutdown(app)
	}
}
