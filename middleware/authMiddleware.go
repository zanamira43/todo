package middleware

import (
	"github.com/gofiber/fiber/v2"
	"github.com/zanamira43/todo/util"
)

func IsAuthenticated(c *fiber.Ctx) error {
	cookie := c.Cookies("jwt")

	if _, err := util.ParseJwt(cookie); err != nil {
		c.SendStatus(fiber.StatusUnauthorized)
		return c.JSON(fiber.Map{
			"message": "User UnAuthenticated",
		})
	}

	return c.Next()
}
