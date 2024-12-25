package controllers

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/kittituchdev/pos-guide/models"
)

// CreateProduct handles creating a new product
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
	if product.Name == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "Name is required",
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
	product.IsActive = true
	product.IsDelete = false

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

// GetAllProduct handles fetching all products
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

// UpdateProduct handles updating a product with partial updates
func UpdateProduct(c *fiber.Ctx) error {
	id := c.Params("id")
	if id == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "ID is required",
		})
	}

	// Parse request body into UpdateProductInput
	var input models.UpdateProductInput
	admin := "Admin"
	input.UpdatedBy = &admin
	updatedAt := time.Now().UnixMilli()
	input.UpdatedAt = &updatedAt
	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "Cannot parse JSON",
		})
	}

	// Validate inputs
	if input.Price != nil && *input.Price < 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "Invalid price value",
		})
	}
	if input.Stock != nil && *input.Stock < 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "Invalid stock value",
		})
	}

	// Perform partial update
	err := models.UpdateProduct(id, input)
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
