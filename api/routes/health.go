package routes

import (
	"github.com/gofiber/fiber/v2"
)

type healthResponse struct {
	Message string	`json:"message"`
}

func HealthCheck(c *fiber.Ctx) error {
	resp := healthResponse{
		Message: "Snip working fine",
	}

	return c.Status(fiber.StatusOK).JSON(resp)
}