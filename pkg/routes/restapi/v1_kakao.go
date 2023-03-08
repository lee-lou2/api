package restapi

import (
	"github.com/gofiber/fiber/v2"
	kakaoAPI "github.com/lee-lou2/api/platform/kakao"
)

func KaKaORoutes(v1 fiber.Router) {
	kakao := v1.Group("/kakao")
	{
		kakao.Post("/token", func(c *fiber.Ctx) error {
			if err := kakaoAPI.CreateKaKaOToken(); err != nil {
				return err
			}
			return c.JSON(map[string]string{"data": "카카오톡 토큰 발급 완료"})
		})
		kakao.Post("/token/refresh", func(c *fiber.Ctx) error {
			if err := kakaoAPI.RefreshKaKaoToken(); err != nil {
				return err
			}
			return c.JSON(map[string]string{"data": "카카오톡 토큰 재발급 완료"})
		})
	}
}
