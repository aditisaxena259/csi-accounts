package routes

import (
    "csi-accounts/database"
    "csi-accounts/internal/models"
    "net/http"

    "github.com/gofiber/fiber/v2"
    "github.com/google/uuid"
)

func CreateAuditLog(c *fiber.Ctx) error {
    var auditLog models.AuditLog

    if err := c.BodyParser(&auditLog); err != nil {
        return c.Status(http.StatusBadRequest).JSON(fiber.Map{
            "error": "Invalid request body",
        })
    }

    if auditLog.UserID == uuid.Nil || auditLog.ClientID == uuid.Nil || auditLog.Action == "" {
        return c.Status(http.StatusBadRequest).JSON(fiber.Map{
            "error": "UserID, ClientID, and Action are required fields",
        })
    }

    if err := database.DB.Create(&auditLog).Error; err != nil {
        return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
            "error": "Failed to create audit log",
        })
    }

    return c.Status(http.StatusOK).JSON(auditLog)
}

func GetAuditLog(c *fiber.Ctx) error {
    id := c.Params("id")
    var auditLog models.AuditLog

    if err := database.DB.First(&auditLog, "id = ?", id).Error; err != nil {
        return c.Status(http.StatusNotFound).JSON(fiber.Map{
            "error": "Audit log not found",
        })
    }

    return c.Status(http.StatusOK).JSON(auditLog)
}

func UpdateAuditLog(c *fiber.Ctx) error {
    id := c.Params("id")
    var auditLog models.AuditLog

    if err := database.DB.First(&auditLog, "id = ?", id).Error; err != nil {
        return c.Status(http.StatusNotFound).JSON(fiber.Map{
            "error": "Audit log not found",
        })
    }

    if err := c.BodyParser(&auditLog); err != nil {
        return c.Status(http.StatusBadRequest).JSON(fiber.Map{
            "error": "Invalid request body",
        })
    }

    if auditLog.Action == "" {
        return c.Status(http.StatusBadRequest).JSON(fiber.Map{
            "error": "Action is a required field",
        })
    }

    if err := database.DB.Save(&auditLog).Error; err != nil {
        return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
            "error": "Failed to update audit log",
        })
    }

    return c.Status(http.StatusOK).JSON(auditLog)
}

func DeleteAuditLog(c *fiber.Ctx) error {
    id := c.Params("id")

    if err := database.DB.Delete(&models.AuditLog{}, "id = ?", id).Error; err != nil {
        return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
            "error": "Failed to delete audit log",
        })
    }

    return c.Status(http.StatusOK).JSON(fiber.Map{
        "message": "Audit log deleted successfully",
    })
}

