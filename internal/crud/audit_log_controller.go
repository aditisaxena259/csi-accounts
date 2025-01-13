package crud

import (
	"csi-accounts/internal/initializers"
	"csi-accounts/internal/models"
	"csi-accounts/pkg/config"
	"csi-accounts/pkg/helpers"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func GetAuditLogs(c *fiber.Ctx) error {
	var auditLogs []models.AuditLog
	if err := initializers.DB.Find(&auditLogs).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return &fiber.Error{Code: 404, Message: "Audit logs not found"}
		}
		return &helpers.AppError{Code: 500, Message: config.DATABASE_ERROR, LogMessage: err.Error(), Err: err}
	}

	return c.Status(200).JSON(fiber.Map{
		"status":  "success",
		"message": "",
		"auditLogs": auditLogs,
	})
}

func GetAuditLog(c *fiber.Ctx) error {
	parsedAuditLogID, err := uuid.Parse(c.Params("auditLogID"))
	if err != nil {
		return &fiber.Error{Code: 400, Message: "Invalid audit log ID"}
	}

	var auditLog models.AuditLog
	if err := initializers.DB.Where("id = ?", parsedAuditLogID).First(&auditLog).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return &fiber.Error{Code: 404, Message: "Audit log not found"}
		}
		return &helpers.AppError{Code: 500, Message: config.DATABASE_ERROR, LogMessage: err.Error(), Err: err}
	}

	return c.Status(200).JSON(fiber.Map{
		"status":  "success",
		"message": "",
		"auditLog": auditLog,
	})
}

func CreateAuditLog(c *fiber.Ctx) error {
	var reqBody models.AuditLog
	if err := c.BodyParser(&reqBody); err != nil {
		return &fiber.Error{Code: 400, Message: "Invalid request body"}
	}

	reqBody.ID = uuid.New()
	if err := initializers.DB.Create(&reqBody).Error; err != nil {
		return &helpers.AppError{Code: 500, Message: config.DATABASE_ERROR, LogMessage: err.Error(), Err: err}
	}

	return c.Status(201).JSON(fiber.Map{
		"status":  "success",
		"message": "Audit log created",
		"auditLog": reqBody,
	})
}

func UpdateAuditLog(c *fiber.Ctx) error {
	parsedAuditLogID, err := uuid.Parse(c.Params("auditLogID"))
	if err != nil {
		return &fiber.Error{Code: 400, Message: "Invalid audit log ID"}
	}

	var reqBody models.AuditLog
	if err := c.BodyParser(&reqBody); err != nil {
		return &fiber.Error{Code: 400, Message: "Invalid request body"}
	}

	var auditLog models.AuditLog
	if err := initializers.DB.Where("id = ?", parsedAuditLogID).First(&auditLog).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return &fiber.Error{Code: 404, Message: "Audit log not found"}
		}
		return &helpers.AppError{Code: 500, Message: config.DATABASE_ERROR, LogMessage: err.Error(), Err: err}
	}

	// Update fields
	if reqBody.Action != "" {
		auditLog.Action = reqBody.Action
	}
	if !reqBody.Timestamp.IsZero() {
		auditLog.Timestamp = reqBody.Timestamp
	}

	if err := initializers.DB.Save(&auditLog).Error; err != nil {
		return &helpers.AppError{Code: 500, Message: config.DATABASE_ERROR, LogMessage: err.Error(), Err: err}
	}

	return c.Status(200).JSON(fiber.Map{
		"status":  "success",
		"message": "Audit log updated",
		"auditLog": auditLog,
	})
}

func DeleteAuditLog(c *fiber.Ctx) error {
	parsedAuditLogID, err := uuid.Parse(c.Params("auditLogID"))
	if err != nil {
		return &fiber.Error{Code: 400, Message: "Invalid audit log ID"}
	}

	if err := initializers.DB.Delete(&models.AuditLog{}, "id = ?", parsedAuditLogID).Error; err != nil {
		return &helpers.AppError{Code: 500, Message: config.DATABASE_ERROR, LogMessage: err.Error(), Err: err}
	}

	return c.Status(200).JSON(fiber.Map{
		"status":  "success",
		"message": "Audit log deleted",
	})
}
