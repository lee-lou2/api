package restapi

import (
	"github.com/gofiber/fiber/v2"
	"github.com/lee-lou2/api/app/notify/controllers"
	"github.com/lee-lou2/api/app/notify/models"
)

// NotifyRoutes func for describe group of private routes.
func NotifyRoutes(v1 fiber.Router) {
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
			case "slack":
				controllers.SendSlack(payload)
			case "kakao-to-me":
				controllers.SendKaKaOToMe(payload)
			case "kakao-to-friend":
				controllers.SendKaKaOToFriend(payload)
			default:
				panic("Not Found Message Type")
			}
			return c.JSON(nil)
		})
	}
}
