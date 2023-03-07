package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/monitor"
)

// PublicRoutes func for describe group of public routes.
func PublicRoutes(a *fiber.App) {
	a.Get("/metrics", monitor.New(monitor.Config{Title: "MyService Metrics Page"}))

	// Create routes group.
	route := a.Group("/api/v1")

	// Routes for GET method:
	route.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("pong")
	})
}
