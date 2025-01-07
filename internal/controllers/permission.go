package controllers

import (
	"csi-accounts/internal/initializers"
	"csi-accounts/internal/models"
	"csi-accounts/pkg/config"
	"csi-accounts/pkg/helpers"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

// GetPermissions retrieves all permissions from the database
func GetPermissions(c *fiber.Ctx) error {
	var permissions []models.Permission
	if err := initializers.DB.Find(&permissions).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return &fiber.Error{Code: 404, Message: "No permissions found"}
		}
		return &helpers.AppError{Code: 500, Message: config.DATABASE_ERROR, LogMessage: err.Error(), Err: err}
	}

	return c.Status(200).JSON(fiber.Map{
		"status": "success",
		"message": "Permissions retrieved successfully",
		"permissions": permissions,
	})
}

// GetPermission retrieves a specific permission by ID
func GetPermission(c *fiber.Ctx) error {
	parsedPermissionID, err := uuid.Parse(c.Params("permissionID"))
	if err != nil {
		return &fiber.Error{Code: 400, Message: "Invalid permission ID"}
	}

	var permission models.Permission
	if err := initializers.DB.First(&permission, "id = ?", parsedPermissionID).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return &fiber.Error{Code: 404, Message: "Permission not found"}
		}
		return &helpers.AppError{Code: 500, Message: config.DATABASE_ERROR, LogMessage: err.Error(), Err: err}
	}

	return c.Status(200).JSON(fiber.Map{
		"status": "success",
		"message": "Permission retrieved successfully",
		"permission": permission,
	})
}

// CreatePermission creates a new permission
func CreatePermission(c *fiber.Ctx) error {
	var reqBody struct {
		Name string `json:"name"`
	}

	if err := c.BodyParser(&reqBody); err != nil {
		return &fiber.Error{Code: 400, Message: "Invalid request body"}
	}

	if reqBody.Name == "" {
		return &fiber.Error{Code: 400, Message: "Permission name is required"}
	}

	permission := models.Permission{
		Name: reqBody.Name,
	}

	if err := initializers.DB.Create(&permission).Error; err != nil {
		return &helpers.AppError{Code: 500, Message: config.DATABASE_ERROR, LogMessage: err.Error(), Err: err}
	}

	return c.Status(201).JSON(fiber.Map{
		"status": "success",
		"message": "Permission created successfully",
		"permission": permission,
	})
}

// UpdatePermission updates an existing permission
func UpdatePermission(c *fiber.Ctx) error {
	parsedPermissionID, err := uuid.Parse(c.Params("permissionID"))
	if err != nil {
		return &fiber.Error{Code: 400, Message: "Invalid permission ID"}
	}

	var reqBody struct {
		Name string `json:"name"`
	}

	if err := c.BodyParser(&reqBody); err != nil {
		return &fiber.Error{Code: 400, Message: "Invalid request body"}
	}

	var permission models.Permission
	if err := initializers.DB.First(&permission, "id = ?", parsedPermissionID).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return &fiber.Error{Code: 404, Message: "Permission not found"}
		}
		return &helpers.AppError{Code: 500, Message: config.DATABASE_ERROR, LogMessage: err.Error(), Err: err}
	}

	if reqBody.Name != "" {
		permission.Name = reqBody.Name
	}

	if err := initializers.DB.Save(&permission).Error; err != nil {
		return &helpers.AppError{Code: 500, Message: config.DATABASE_ERROR, LogMessage: err.Error(), Err: err}
	}

	return c.Status(200).JSON(fiber.Map{
		"status": "success",
		"message": "Permission updated successfully",
		"permission": permission,
	})
}

// DeletePermission deletes a permission by ID
func DeletePermission(c *fiber.Ctx) error {
	parsedPermissionID, err := uuid.Parse(c.Params("permissionID"))
	if err != nil {
		return &fiber.Error{Code: 400, Message: "Invalid permission ID"}
	}

	if err := initializers.DB.Delete(&models.Permission{}, "id = ?", parsedPermissionID).Error; err != nil {
		return &helpers.AppError{Code: 500, Message: config.DATABASE_ERROR, LogMessage: err.Error(), Err: err}
	}

	return c.Status(200).JSON(fiber.Map{
		"status": "success",
		"message": "Permission deleted successfully",
	})
}
