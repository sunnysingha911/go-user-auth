package handlers

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/sunnysingha911/user-service/models"
	"github.com/sunnysingha911/user-service/services"
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

func GetAllUsers(c *fiber.Ctx) error {
	// Parse query params
	page, err := strconv.Atoi(c.Query("page", "1"))
	if err != nil || page < 1 {
		page = 1
	}
	limit, err := strconv.Atoi(c.Query("limit", "10"))
	if err != nil || limit < 1 {
		limit = 10
	}

	// Call service
	res, err := services.GetUserList(page, limit)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "failed to fetch users",
		})
	}

	// Format response
	users := make([]fiber.Map, 0)
	for _, user := range res.Users {
		users = append(users, buildUserResponse(user))
	}

	return c.JSON(fiber.Map{
		"users": users,
		"meta": fiber.Map{
			"total": res.Total,
			"page":  res.Page,
			"limit": res.Limit,
		},
	})
}
