package middlewares

import (
	"github/sing3demons/go-fiber-admin/util"

	"github.com/gofiber/fiber/v2"
)

func IsAuthenticated(c *fiber.Ctx) error {
	cookie := c.Cookies("jwt")

	_, err := util.ParseJwt(cookie)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"message": err.Error()})
	}
	return c.Next()

}
