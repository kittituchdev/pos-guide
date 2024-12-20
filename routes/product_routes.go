package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/kittituchdev/pos-guide/controllers"
)

func UserRoutes(app *fiber.App) {
	api := app.Group("/api")

	// User routes
	api.Get("/products", controllers.GetProducts)
	api.Get("/products/:id", controllers.GetUser)
	api.Post("/products", controllers.CreateProduct)
	api.Put("/products/:id", controllers.UpdateUser)
	api.Delete("/products/:id", controllers.DeleteUser)
}
