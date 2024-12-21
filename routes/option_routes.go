package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/kittituchdev/pos-guide/controllers"
)

func OptionRoutes(app *fiber.App) {
	api := app.Group("/api")

	api.Get("/options", controllers.GetAllOption)
	api.Post("/options", controllers.CreateOption)
}
