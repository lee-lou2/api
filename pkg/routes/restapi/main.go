package restapi

import (
	swagger "github.com/arsmn/fiber-swagger/v2"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/monitor"
)

// MainRoutes 기본 라우터
func MainRoutes(a *fiber.App) {
	a.Get("/metrics", monitor.New(monitor.Config{Title: "MyService Metrics Page"}))

	a.Get("/ping", func(c *fiber.Ctx) error {
		return c.SendString("pong")
	})

	// Create v1 routes group.
	v1 := a.Group("/v1")
	{
		// 메시지
		NotifyRoutes(v1)
		// 카카오
		KaKaORoutes(v1)
	}
}

// SwaggerRoute func for describe group of API Docs routes.
func SwaggerRoute(a *fiber.App) {
	// Create routes group.
	route := a.Group("/swagger")

	// Routes for GET method:
	route.Get("*", swagger.HandlerDefault) // get one user by ID
}

// NotFoundRoute func for describe 404 Error route.
func NotFoundRoute(a *fiber.App) {
	// Register new special route.
	a.Use(
		// Anonymous function.
		func(c *fiber.Ctx) error {
			// Return HTTP 404 status and JSON response.
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"error": true,
				"msg":   "sorry, endpoint is not found",
			})
		},
	)
}
