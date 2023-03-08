package middleware

import (
	"context"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/encryptcookie"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/lee-lou2/hub/platform/database"
	"log"
	"os"
	"time"
)

// FiberMiddleware provide Fiber's built-in middlewares.
// See: https://docs.gofiber.io/api/middleware
func FiberMiddleware(a *fiber.App) {
	// cors, logger
	a.Use(
		cors.New(),
		logger.New(),
	)
	// encrypt cookie
	a.Use(encryptcookie.New(encryptcookie.Config{
		Key: os.Getenv("FIBER_ENCRYPT_COOKIE_KEY"),
	}))
	// access log
	a.Use(func(c *fiber.Ctx) error {
		logEntry := map[string]interface{}{
			"method":    c.Method(),
			"uri":       c.Request().URI().String(),
			"ips":       c.IPs(),
			"body":      string(c.Body()),
			"userAgent": c.Get("User-Agent"),
			"now":       time.Now().Unix(),
		}
		// 비동기 처리
		go func(_logEntry map[string]interface{}) {
			// 로그를 MongoDB에 저장
			dbName := os.Getenv("LOG_GROUP_DATABASE")
			client, collection, _ := database.GetCollection(dbName, "access")
			defer client.Disconnect(context.Background())
			_, err := collection.InsertOneDocument(_logEntry)
			if err != nil {
				log.Println("Failed to insert log:", err)
			}
		}(logEntry)
		// 다음 미들웨어 또는 핸들러로 진행
		return c.Next()
	})
}
