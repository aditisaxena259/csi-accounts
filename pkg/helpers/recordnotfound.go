package helpers

import (
	"log"

	"github.com/gofiber/fiber/v2"
)

// InternalServerError is a reusable helper for handling internal server errors
func RecordNotFoundError(c *fiber.Ctx, err error, message string) error {
	log.Printf("Record not Found: %s", err.Error()) // Log the error for debugging

	return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
		"status":  "error",
		"message": message,
	})
}