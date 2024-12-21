package controllers

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/kittituchdev/pos-guide/models"
)

func CreateProduct(c *fiber.Ctx) error {
	var product models.Product

	// Parse incoming JSON
	if err := c.BodyParser(&product); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "Cannot parse JSON",
		})
	}

	// Validate required fields
	if product.Name == "" || product.Price <= 0 || product.Stock < 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "Missing or invalid required fields",
			"errors": fiber.Map{
				"name":  "Name is required",
				"price": "Price must be greater than zero",
				"stock": "Stock cannot be negative",
			},
		})
	}

	// Set default values and timestamps
	if product.CreatedBy == "" {
		product.CreatedBy = "Admin"
	}
	if product.UpdatedBy == "" {
		product.UpdatedBy = "Admin"
	}
	product.CreatedAt = time.Now().UnixMilli()
	product.UpdatedAt = time.Now().UnixMilli()
	product.IsActive = true  // Default to active
	product.IsDelete = false // Default to not deleted

	// Insert into database
	err := models.InsertOneProduct(product)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"message": "Failed to insert product",
		})
	}

	// Return success response
	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"success": true,
		"message": "Product inserted successfully",
	})
}

func GetAllProduct(c *fiber.Ctx) error {
	// Fetch all products
	result, err := models.FindAllProduct()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"message": "Failed to get products",
		})
	}

	// Return success response with data
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": true,
		"message": "Products fetched successfully",
		"data":    result,
	})
}

func UpdateProductPrice(c *fiber.Ctx) error {
	var product models.Product
	id := c.Params("id")

	// Parse incoming JSON
	if err := c.BodyParser(&product); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "Cannot parse JSON",
		})
	}

	// Validate required fields
	if id == "" || product.Price <= 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "Missing or invalid required fields",
			"errors": fiber.Map{
				"id":    "ID is required",
				"price": "Price must be greater than zero",
			},
		})
	}

	// Set updated timestamp
	product.UpdatedAt = time.Now().UnixMilli()

	// Update product price
	err := models.UpdateProductPrice(id, product.Price)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"message": "Failed to update product price",
		})
	}

	// Return success response
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": true,
		"message": "Product price updated successfully",
	})
}

func UpdateProduct(c *fiber.Ctx) error {
	var product models.Product
	id := c.Params("id")

	// Parse incoming JSON
	if err := c.BodyParser(&product); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "Cannot parse JSON",
		})
	}

	// Validate required fields
	if id == "" || product.Name == "" || product.Price <= 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "Missing or invalid required fields",
			"errors": fiber.Map{
				"id":    "ID is required",
				"name":  "Name is required",
				"price": "Price must be greater than zero",
				"stock": "Stock cannot be negative",
			},
		})
	}

	// Set updated timestamp
	product.UpdatedAt = time.Now().UnixMilli()

	// Update product
	err := models.UpdateProduct(id, product)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"message": "Failed to update product",
		})
	}

	// Return success response
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": true,
		"message": "Product updated successfully",
	})
}

// func GetUser(c *fiber.Ctx) error {
// 	id := c.Params("id")
// 	return c.JSON(fiber.Map{
// 		"message": "Get user by ID",
// 		"id":      id,
// 	})
// }

// func UpdateUser(c *fiber.Ctx) error {
// 	id := c.Params("id")
// 	return c.JSON(fiber.Map{
// 		"message": "User updated successfully",
// 		"id":      id,
// 	})
// }

// func DeleteUser(c *fiber.Ctx) error {
// 	id := c.Params("id")
// 	return c.JSON(fiber.Map{
// 		"message": "User deleted successfully",
// 		"id":      id,
// 	})
// }
