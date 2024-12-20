package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/kittituchdev/pos-guide/models"
)

func GetProducts(c *fiber.Ctx) error {
	result := models.FindAllProduct()
	return c.JSON(fiber.Map{
		"data": result,
	})
}

func GetUser(c *fiber.Ctx) error {
	id := c.Params("id")
	return c.JSON(fiber.Map{
		"message": "Get user by ID",
		"id":      id,
	})
}

func CreateProduct(c *fiber.Ctx) error {
	var product models.Product

	if err := c.BodyParser(&product); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Cannot parse JSON",
		})
	}
	err := models.InsertOneProduct(product)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to insert product",
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "Product inserted successfully",
	})
}

func UpdateUser(c *fiber.Ctx) error {
	id := c.Params("id")
	return c.JSON(fiber.Map{
		"message": "User updated successfully",
		"id":      id,
	})
}

func DeleteUser(c *fiber.Ctx) error {
	id := c.Params("id")
	return c.JSON(fiber.Map{
		"message": "User deleted successfully",
		"id":      id,
	})
}
