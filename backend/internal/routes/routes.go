package routes

import (
	"alman-hesabi-backend/internal/controllers"

	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App) {
	api := app.Group("/api")

	api.Get("/health", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{"status": "ok"})
	})

	auth := api.Group("/auth")
	auth.Post("/register", controllers.Register)
	auth.Post("/login", controllers.Login)

	groups := api.Group("/groups")
	groups.Post("/", controllers.CreateGroup)
	groups.Post("/join", controllers.JoinGroup)
	groups.Get("/:groupId", controllers.GetGroup)

	expenses := api.Group("/groups/:groupId/expenses")
	expenses.Post("/", controllers.CreateExpense)
	expenses.Get("/", controllers.GetExpenses)

	balances := api.Group("/groups/:groupId")
	balances.Get("/balances", controllers.GetBalances)
}
