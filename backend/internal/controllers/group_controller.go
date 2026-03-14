package controllers

import "github.com/gofiber/fiber/v2"

func CreateGroup(c *fiber.Ctx) error {
	// TODO: Implement based on stories.md Story 3
	return c.Status(fiber.StatusNotImplemented).JSON(fiber.Map{
		"message": "CreateGroup endpoint - not yet implemented",
	})
}

func JoinGroup(c *fiber.Ctx) error {
	// TODO: Implement based on stories.md Story 4
	return c.Status(fiber.StatusNotImplemented).JSON(fiber.Map{
		"message": "JoinGroup endpoint - not yet implemented",
	})
}

func GetGroup(c *fiber.Ctx) error {
	// TODO: Implement based on stories.md Story 3
	return c.Status(fiber.StatusNotImplemented).JSON(fiber.Map{
		"message": "GetGroup endpoint - not yet implemented",
	})
}
