package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/lee-lou2/hub/app/notification/controllers"
	"github.com/lee-lou2/hub/app/notification/models"
)

// PrivateRoutes func for describe group of private routes.
func PrivateRoutes(a *fiber.App) {
	route := a.Group("/api/v1")

	notify := route.Group("/notify")
	{
		notify.Post("/send", func(c *fiber.Ctx) error {
			var payload models.Message
			if err := c.BodyParser(&payload); err != nil {
				return err
			}
			switch payload.MessageType {
			case 0:
				controllers.SendEmail(payload)
			case 1:
				controllers.SenSMS(payload)
			default:
				panic("Not Found Message Type")
			}
			return c.JSON(nil)
		})
	}
}
