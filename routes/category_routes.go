package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/kittituchdev/pos-guide/controllers"
)

func CategoryRoutes(app *fiber.App) {
	api := app.Group("/api")

	api.Get("/categories", controllers.GetAllCategory)
	api.Post("/categories", controllers.CreateCategory)
}
