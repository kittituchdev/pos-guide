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

	if len(order.OrderItems) == 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "Order items are required",
		})
	}

	// Generate Order Number
	order.OrderNumber = utils.GenerateOrderNumber()

	// Set default values
	order.CreatedAt = time.Now().UnixMilli()
	order.UpdatedAt = time.Now().UnixMilli()
	order.IsCancel = false

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

func GetAllOrder(c *fiber.Ctx) error {
	// Get all orders

	orders, err := models.FindAllOrder()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"message": "Failed to get orders",
		})
	}

	// Return success response
	return c.JSON(fiber.Map{
		"success": true,
		"message": "Orders retrieved successfully",
		"data":    orders,
	})
}
