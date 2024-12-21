package controllers

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/kittituchdev/pos-guide/models"
)

func CreateCategory(c *fiber.Ctx) error {
	var category models.Category

	// Parse incoming JSON
	if err := c.BodyParser(&category); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "Cannot parse JSON",
		})
	}

	// Validate required fields
	if category.Name == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "Missing or invalid required fields",
			"errors": fiber.Map{
				"name": "Name is required",
			},
		})
	}

	// Set default values and timestamps
	if category.CreatedBy == "" {
		category.CreatedBy = "Admin"
	}
	if category.UpdatedBy == "" {
		category.UpdatedBy = "Admin"
	}
	category.CreatedAt = time.Now().UnixMilli()
	category.UpdatedAt = time.Now().UnixMilli()
	category.IsActive = true  // Default to active
	category.IsDelete = false // Default to not deleted

	// Insert into database
	err := models.InsertOneCategory(category)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"message": "Failed to insert category",
		})
	}

	// Return success response
	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"success": true,
		"message": "Category inserted successfully",
	})
}

func GetAllCategory(c *fiber.Ctx) error {
	categories, err := models.FindAllCategory()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"message": "Failed to fetch categories",
		})
	}

	return c.JSON(fiber.Map{
		"success":    true,
		"message":    "Categories fetched successfully",
		"categories": categories,
	})
}
