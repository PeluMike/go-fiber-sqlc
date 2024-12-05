package middlewares

import (
	"context"
	"strings"

	"github.com/PeluMike/blog/database"
	"github.com/PeluMike/blog/src/config"
	"github.com/gofiber/fiber/v2"
)

func IsAuth(c *fiber.Ctx) error {
	authHeader := c.Get("Authorization")

	if authHeader == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Your are not authenticated",
		})
	}

	parts := strings.Split(authHeader, " ")
	claim, err := config.VerifyJWT(parts[1])

	if len(parts) != 2 {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Your are not authenticated",
		})
	}

	if err != nil {
		return c.Status(401).JSON(fiber.Map{
			"error": "Your are not authenticated",
		})
	}

	user, err1 := database.Queries.GetUserByEmail(context.Background(), claim.UserEmail)

	if err1 != nil {
		return c.Status(401).JSON(fiber.Map{
			"error": "Your are not authenticated",
		})
	}

	c.Locals("user", user)
	return c.Next()
}
