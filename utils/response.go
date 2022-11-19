package utils

import (
	"github.com/gofiber/fiber/v2"
)

func MakeResponse(c *fiber.Ctx, data interface{}, status int) error {
	return c.Status(status).JSON(fiber.Map{
		"success": true,
		"data":    data,
	})
}

func MakeErrorResponse(c *fiber.Ctx, message string, status int) error {
	return c.Status(status).JSON(fiber.Map{
		"error":   true,
		"message": message,
	})
}
