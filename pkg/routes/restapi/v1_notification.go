package restapi

import (
	"github.com/gofiber/fiber/v2"
	"github.com/lee-lou2/hub/app/notification/controllers"
	"github.com/lee-lou2/hub/app/notification/models"
)

// NotificationRoutes func for describe group of private routes.
func NotificationRoutes(v1 fiber.Router) {
	notify := v1.Group("/notify")
	{
		notify.Post("/send/:app", func(c *fiber.Ctx) error {
			var payload models.Message
			if err := c.BodyParser(&payload); err != nil {
				return err
			}

			// 라우팅
			switch c.Params("app") {
			case "email":
				controllers.SendEmail(payload)
			case "sms":
				controllers.SenSMS(payload)
			default:
				panic("Not Found Message Type")
			}
			return c.JSON(nil)
		})
	}
}
