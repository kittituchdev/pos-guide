package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/kittituchdev/pos-guide/controllers"
)

// OrderRoutes function
func OrderRoutes(app *fiber.App) {
	api := app.Group("/api")

	api.Post("/orders", controllers.CreateOrder)
}
