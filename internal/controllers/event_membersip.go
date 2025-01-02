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

func GetEventMemberships(c *fiber.Ctx) error {
	parsedEventID, err := uuid.Parse(c.Params("eventID"))
	if err != nil {
		return &fiber.Error{Code: 400, Message: "Invalid event ID"}
	}

	var eventMemberships []models.EventMembership
	if err := initializers.DB.Where("event_id = ?", parsedEventID).Find(&eventMemberships).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return &fiber.Error{Code: 404, Message: "Event not found"}
		}
		return &helpers.AppError{Code:500, Message: config.DATABASE_ERROR, LogMessage: err.Error(), Err: err}
	}

	return c.Status(200).JSON(fiber.Map{
		"status": "success",
		"message": "",
		"eventMemberships": eventMemberships,
	})
}

func AddEventCoordinator(c *fiber.Ctx) error {
	parsedUserID , err := uuid.Parse(c.Params("userID"))
	if err != nil {
		return &fiber.Error{Code: 400, Message: "Invalid user ID"}
	}

	parsedEventID, err := uuid.Parse(c.Params("eventID"))
	if err != nil {
		return &fiber.Error{Code: 400, Message: "Invalid event ID"}
	}

	var eventMembership models.EventMembership
	if err := initializers.DB.Where("event_id = ? AND user_id = ?", parsedEventID, parsedUserID).First(&eventMembership).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return &fiber.Error{Code: 404, Message: "User not found"}
		}
		return &helpers.AppError{Code:500, Message: config.DATABASE_ERROR, LogMessage: err.Error(), Err: err}
	}

	eventMembership.Role = models.Coordinator

	if err := initializers.DB.Save(&eventMembership).Error; err != nil {
		return &helpers.AppError{Code:500, Message: config.DATABASE_ERROR, LogMessage: err.Error(), Err: err}
	}

	return nil
}
