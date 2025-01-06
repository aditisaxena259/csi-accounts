package controllers

import (
	"csi-accounts/internal/initializers"
	"csi-accounts/internal/models"
	"csi-accounts/pkg/helpers"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

// Get all roles
func GetRoles(c *fiber.Ctx) error {
	var roles []models.Role
	if err := initializers.DB.Find(&roles).Error; err != nil {
		return helpers.InternalServerError(c, err, "Failed to fetch roles")
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status": "success",
		"roles":  roles,
	})
}

// Get a single role by ID
func GetRole(c *fiber.Ctx) error {
	roleID, err := uuid.Parse(c.Params("roleID"))
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid role ID")
	}

	var role models.Role
	if err := initializers.DB.First(&role, "id = ?", roleID).Error; err != nil {
		return helpers.RecordNotFoundError(c, err, "Role not found")
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status": "success",
		"role":   role,
	})
}

// Create a new role
func CreateRole(c *fiber.Ctx) error {
	var reqBody models.Role
	if err := c.BodyParser(&reqBody); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid request body")
	}

	reqBody.ID = uuid.New() // Generate a new UUID for the role

	if err := initializers.DB.Create(&reqBody).Error; err != nil {
		return helpers.InternalServerError(c, err, "Failed to create role")
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"status": "success",
		"role":   reqBody,
	})
}

// Assign permissions to a role
func AssignPermissionsToRole(c *fiber.Ctx) error {
	var reqBody struct {
		Permissions []uuid.UUID `json:"permissions"` // List of permission IDs to assign
	}
	roleID, err := uuid.Parse(c.Params("roleID"))
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid role ID")
	}

	if err := c.BodyParser(&reqBody); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid request body")
	}

	// Validate role existence
	var role models.Role
	if err := initializers.DB.First(&role, "id = ?", roleID).Error; err != nil {
		return helpers.RecordNotFoundError(c, err, "Role not found")
	}

	// Add permissions to the role
	for _, permissionID := range reqBody.Permissions {
		rolePermission := models.RolePermission{
			ID:         uuid.New(),
			RoleID:     roleID,
			Permissions: []models.Permission{
				{ID: permissionID},
			},
		}
		if err := initializers.DB.Create(&rolePermission).Error; err != nil {
			return helpers.InternalServerError(c, err, "Failed to assign permissions to role")
		}
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"status":  "success",
		"message": "Permissions assigned successfully",
	})
}

// Delete a role by ID
func DeleteRole(c *fiber.Ctx) error {
	roleID, err := uuid.Parse(c.Params("roleID"))
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid role ID")
	}

	if err := initializers.DB.Delete(&models.Role{}, "id = ?", roleID).Error; err != nil {
		return helpers.InternalServerError(c, err, "Failed to delete role")
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":  "success",
		"message": "Role deleted successfully",
	})
}
