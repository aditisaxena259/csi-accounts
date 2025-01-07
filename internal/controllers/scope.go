package controllers

import (
	"csi-accounts/internal/initializers"
	"csi-accounts/internal/models"
	"csi-accounts/pkg/helpers"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

// Get all scopes
func GetScopes(c *fiber.Ctx) error {
	var scopes []models.Scope
	if err := initializers.DB.Find(&scopes).Error; err != nil {
		return helpers.InternalServerError(c, err, "Failed to fetch scopes")
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status": "success",
		"scopes": scopes,
	})
}

// Get a single scope by ID
func GetScope(c *fiber.Ctx) error {
	scopeID, err := uuid.Parse(c.Params("scopeID"))
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid scope ID")
	}

	var scope models.Scope
	if err := initializers.DB.First(&scope, "id = ?", scopeID).Error; err != nil {
		return helpers.RecordNotFoundError(c, err, "Scope not found")
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status": "success",
		"scope":  scope,
	})
}

// Create a new scope
func CreateScope(c *fiber.Ctx) error {
	var reqBody models.Scope
	if err := c.BodyParser(&reqBody); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid request body")
	}

	reqBody.ID = uuid.New() // Generate a new UUID for the scope

	if err := initializers.DB.Create(&reqBody).Error; err != nil {
		return helpers.InternalServerError(c, err, "Failed to create scope")
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"status": "success",
		"scope":  reqBody,
	})
}

// Update an existing scope
func UpdateScope(c *fiber.Ctx) error {
	scopeID, err := uuid.Parse(c.Params("scopeID"))
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid scope ID")
	}

	var reqBody models.Scope
	if err := c.BodyParser(&reqBody); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid request body")
	}

	var scope models.Scope
	if err := initializers.DB.First(&scope, "id = ?", scopeID).Error; err != nil {
		return helpers.RecordNotFoundError(c, err, "Scope not found")
	}

	// Update scope fields
	scope.Name = reqBody.Name
	scope.Description = reqBody.Description

	if err := initializers.DB.Save(&scope).Error; err != nil {
		return helpers.InternalServerError(c, err, "Failed to update scope")
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status": "success",
		"scope":  scope,
	})
}

// Delete a scope by ID
func DeleteScope(c *fiber.Ctx) error {
	scopeID, err := uuid.Parse(c.Params("scopeID"))
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid scope ID")
	}

	if err := initializers.DB.Delete(&models.Scope{}, "id = ?", scopeID).Error; err != nil {
		return helpers.InternalServerError(c, err, "Failed to delete scope")
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":  "success",
		"message": "Scope deleted successfully",
	})
}
