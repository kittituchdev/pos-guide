package controllers

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/kittituchdev/pos-guide/models"
)

func CreateOption(c *fiber.Ctx) error {

	var option models.Option

	// Parse incoming JSON
	if err := c.BodyParser(&option); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "Cannot parse JSON",
		})
	}

	// Validate required fields
	if option.Name == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "Missing or invalid required fields",
			"errors": fiber.Map{
				"name": "Name is required",
			},
		})
	}

	// Set default values and timestamps
	if option.CreatedBy == "" {
		option.CreatedBy = "Admin"
	}
	if option.UpdatedBy == "" {
		option.UpdatedBy = "Admin"
	}
	option.CreatedAt = time.Now().UnixMilli()
	option.UpdatedAt = time.Now().UnixMilli()
	option.IsActive = true  // Default to active
	option.IsDelete = false // Default to not deleted

	// Insert into database
	err := models.InsertOneOption(option)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"message": "Failed to insert option",
		})
	}

	// Return success response
	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"success": true,
		"message": "Option inserted successfully",
	})
}

func GetAllOption(c *fiber.Ctx) error {
	options, err := models.FindAllOption()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"message": "Failed to fetch options",
		})
	}
	return c.JSON(fiber.Map{
		"success": true,
		"message": "Options found",
		"data":    options,
	})
}
