package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/kittituchdev/pos-guide/config"
	"github.com/kittituchdev/pos-guide/routes"
)

func main() {
	app := fiber.New()

	config.ConnectDatabase()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(&fiber.Map{"data": "Hello from Fiber & mongoDB"})
	})

	// Register routes
	routes.UserRoutes(app)
	routes.CategoryRoutes(app)
	routes.OptionRoutes(app)
	routes.OrderRoutes(app)

	app.Listen(":3000")
}
