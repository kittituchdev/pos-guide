package controllers

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/kittituchdev/pos-guide/models"
	"github.com/kittituchdev/pos-guide/utils"
)

func CreateOrder(c *fiber.Ctx) error {
	// Create an order
	var order models.Order

	// Parse incoming JSON
	if err := c.BodyParser(&order); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "Cannot parse JSON",
		})
	}

	// Generate Order Number
	order.OrderNumber = utils.GenerateOrderNumber()

	// Set default values
	order.CreatedAt = time.Now().UnixMilli()
	order.UpdatedAt = time.Now().UnixMilli()
	order.IsActive = true
	order.IsDelete = false

	// Insert into database
	err := models.InsertOneOrder(order)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"message": "Failed to insert order",
		})
	}

	// Return success response
	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"success": true,
		"message": "Order inserted successfully",
	})
}
