package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/kittituchdev/pos-guide/controllers"
)

func UserRoutes(app *fiber.App) {
	api := app.Group("/api")

	api.Get("/products", controllers.GetAllProduct)
	api.Post("/products", controllers.CreateProduct)
	api.Patch("/products/:id", controllers.UpdateProduct)

}
