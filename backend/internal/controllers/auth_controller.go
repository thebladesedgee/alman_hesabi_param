package controllers

import "github.com/gofiber/fiber/v2"

func Register(c *fiber.Ctx) error {
	// TODO: Implement based on stories.md Story 1
	return c.Status(fiber.StatusNotImplemented).JSON(fiber.Map{
		"message": "Register endpoint - not yet implemented",
	})
}

func Login(c *fiber.Ctx) error {
	// TODO: Implement based on stories.md Story 2
	return c.Status(fiber.StatusNotImplemented).JSON(fiber.Map{
		"message": "Login endpoint - not yet implemented",
	})
}
