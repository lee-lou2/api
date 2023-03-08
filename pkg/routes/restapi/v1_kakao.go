package restapi

import (
	"github.com/gofiber/fiber/v2"
	kakaoAPI "github.com/lee-lou2/api/platform/kakao"
	"log"
)

func KaKaORoutes(v1 fiber.Router) {
	kakao := v1.Group("/kakao")
	{
		kakao.Post("/token", func(c *fiber.Ctx) error {
			if err := kakaoAPI.CreateKaKaOToken(); err != nil {
				log.Println(err)
			}
			return c.JSON(nil)
		})
		kakao.Post("/token/refresh", func(c *fiber.Ctx) error {
			if err := kakaoAPI.RefreshKaKaoToken(); err != nil {
				log.Println(err)
			}
			return c.JSON(nil)
		})
	}
}
