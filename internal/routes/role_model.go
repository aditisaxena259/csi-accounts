package routes
//role_model
import (
    "csi-accounts/database"
    "csi-accounts/internal/models"
    "net/http"

    "github.com/gofiber/fiber/v2"
)

// Create a Role
func CreateRole(c *fiber.Ctx) error {
    var role models.Role

    if err := c.BodyParser(&role); err != nil {
        return c.Status(http.StatusBadRequest).JSON(fiber.Map{
            "error": "Invalid request body",
        })
    }

    if err := database.DB.Create(&role).Error; err != nil {
        return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
            "error": "Failed to create role",
        })
    }

    return c.Status(http.StatusOK).JSON(role)
}

// Get a Role
func GetRole(c *fiber.Ctx) error {
    id := c.Params("id")
    var role models.Role

    if err := database.DB.First(&role, "id = ?", id).Error; err != nil {
        return c.Status(http.StatusNotFound).JSON(fiber.Map{
            "error": "Role not found",
        })
    }

    return c.Status(http.StatusOK).JSON(role)
}

// Update a Role
func UpdateRole(c *fiber.Ctx) error {
    id := c.Params("id")
    var role models.Role

    if err := database.DB.First(&role, "id = ?", id).Error; err != nil {
        return c.Status(http.StatusNotFound).JSON(fiber.Map{
            "error": "Role not found",
        })
    }

    if err := c.BodyParser(&role); err != nil {
        return c.Status(http.StatusBadRequest).JSON(fiber.Map{
            "error": "Invalid request body",
        })
    }

    if err := database.DB.Save(&role).Error; err != nil {
        return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
            "error": "Failed to update role",
        })
    }

    return c.Status(http.StatusOK).JSON(role)
}

// Delete a Role
func DeleteRole(c *fiber.Ctx) error {
    id := c.Params("id")

    if err := database.DB.Delete(&models.Role{}, "id = ?", id).Error; err != nil {
        return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
            "error": "Failed to delete role",
        })
    }

    return c.Status(http.StatusOK).JSON(fiber.Map{
        "message": "Role deleted successfully",
    })
}

// Create a RolePermission
func CreateRolePermission(c *fiber.Ctx) error {
    var rolePermission models.RolePermission

    if err := c.BodyParser(&rolePermission); err != nil {
        return c.Status(http.StatusBadRequest).JSON(fiber.Map{
            "error": "Invalid request body",
        })
    }

    if err := database.DB.Create(&rolePermission).Error; err != nil {
        return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
            "error": "Failed to create role permission",
        })
    }

    return c.Status(http.StatusOK).JSON(rolePermission)
}

// Get RolePermission by RoleID
func GetRolePermissions(c *fiber.Ctx) error {
    roleID := c.Params("roleID")
    var rolePermissions []models.RolePermission

    if err := database.DB.Where("role_id = ?", roleID).Find(&rolePermissions).Error; err != nil {
        return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
            "error": "Failed to retrieve role permissions",
        })
    }

    return c.Status(http.StatusOK).JSON(rolePermissions)
}

// Delete RolePermission
func DeleteRolePermission(c *fiber.Ctx) error {
    id := c.Params("id")

    if err := database.DB.Delete(&models.RolePermission{}, "id = ?", id).Error; err != nil {
        return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
            "error": "Failed to delete role permission",
        })
    }

    return c.Status(http.StatusOK).JSON(fiber.Map{
        "message": "Role permission deleted successfully",
    })
}

