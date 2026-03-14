package controllers

import "github.com/gofiber/fiber/v2"

func CreateExpense(c *fiber.Ctx) error {
	// TODO: Implement based on stories.md Story 5, 6, 7
	return c.Status(fiber.StatusNotImplemented).JSON(fiber.Map{
		"message": "CreateExpense endpoint - not yet implemented",
	})
}

func GetExpenses(c *fiber.Ctx) error {
	// TODO: Implement based on stories.md Story 5
	return c.Status(fiber.StatusNotImplemented).JSON(fiber.Map{
		"message": "GetExpenses endpoint - not yet implemented",
	})
}

func GetBalances(c *fiber.Ctx) error {
	// TODO: Implement based on stories.md Story 8
	return c.Status(fiber.StatusNotImplemented).JSON(fiber.Map{
		"message": "GetBalances endpoint - not yet implemented",
	})
}
