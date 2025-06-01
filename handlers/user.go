package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/sunnysingha911/user-service/models"
)

// GetMe returns the current authenticated user info
func GetMe(c *fiber.Ctx) error {
	user, ok := c.Locals("user").(*models.User)
	if !ok || user == nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "unauthorized"})
	}

	return c.JSON(fiber.Map{
		"user": buildUserResponse(user),
	})
}
