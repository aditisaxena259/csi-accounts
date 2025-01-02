package controllers

import (
	"csi-accounts/internal/initializers"
	"csi-accounts/internal/models"
	"csi-accounts/internal/schemas"
	"csi-accounts/pkg/config"
	"csi-accounts/pkg/helpers"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func GetEvents(c *fiber.Ctx) error {

	var events []models.Event
	if err := initializers.DB.Find(&events).Error; err != nil {
		if err == gorm.ErrRecordNotFound{
			return &fiber.Error{Code: 404, Message: "Events not found"}
		}
		return &helpers.AppError{Code:500, Message: config.DATABASE_ERROR, LogMessage: err.Error(), Err: err}
	}

	return c.Status(200).JSON(fiber.Map{
		"status": "success",
		"message": "",
		"events": events,
	})
}

func GetEvent(c *fiber.Ctx) error {
	parsedEventID, err := uuid.Parse(c.Params("eventID"))
	if err != nil {
		return &fiber.Error{Code: 400, Message: "Invalid event ID"}
	}

	var event models.Event
	if err := initializers.DB.Where("id = ?", parsedEventID).First(&event).Error; err != nil {
		if err == gorm.ErrRecordNotFound{
			return &fiber.Error{Code: 404, Message: "Event not found"}
		}
		return &helpers.AppError{Code:500, Message: config.DATABASE_ERROR, LogMessage: err.Error(), Err: err}
	}

	return c.Status(200).JSON(fiber.Map{
		"status": "success",
		"message": "",
		"event": event,
	})
}

func CreateEvent(c *fiber.Ctx) error {
	loggedInUserID := c.GetRespHeader("loggedInUserID")
	parsedUserID, _ := uuid.Parse(loggedInUserID)

	var reqBody schemas.CreateEventBody
	if err := c.BodyParser(&reqBody); err != nil {
		return &fiber.Error{Code: 400, Message: "Invalid request body"}
	}

	var user models.User
	if err := initializers.DB.Where("id = ?", parsedUserID).First(&user).Error; err != nil {
		if err == gorm.ErrRecordNotFound{
			return &fiber.Error{Code: 404, Message: "User not found"}
		}
		return helpers.AppError{Code:500, Message: config.DATABASE_ERROR, LogMessage: err.Error(), Err: err}
	}

	tx := initializers.DB.Begin()
	if tx.Error != nil {
		return tx.Error
	}

	defer func() {
		if tx.Error != nil {
			tx.Rollback()
			go helpers.LogDatabaseError("Transaction rolled back due to error", tx.Error, "CreateEvent")
		}
	}()

	event := models.Event{
		Name: reqBody.Name,
		Date: reqBody.Date,
		Location: reqBody.Location,
	}

	result := tx.Create(&event)
	if result.Error != nil {
		return helpers.AppError{Code: 500, Message: config.DATABASE_ERROR, LogMessage: result.Error.Error(), Err: result.Error}
	}

	eventMemership := models.EventMembership{
		EventID: event.ID,
		UserID: user.ID,
		Role: models.Admin,
	}

	result = tx.Create(&eventMemership)
	if result.Error != nil {
		return helpers.AppError{Code: 500, Message: config.DATABASE_ERROR, LogMessage: result.Error.Error(), Err: result.Error}
	}

	if err := tx.Commit().Error; err != nil {
		return helpers.AppError{Code: 500, Message: config.DATABASE_ERROR, LogMessage: err.Error(), Err: err}
	}
	
	return c.Status(201).JSON(fiber.Map{
		"status": "success",
		"message": "",
		"event": event,
	})
}

func UpdateEvent(c *fiber.Ctx) error {
	parsedEventID, err := uuid.Parse(c.Params("eventID"))
	if err != nil {
		return &fiber.Error{Code: 400, Message: "Invalid event ID"}
	}

	var reqBody schemas.UpdateEventBody
	if err := c.BodyParser(&reqBody); err != nil {
		return &fiber.Error{Code: 400, Message: "Invalid request body"}		
	}

	var event models.Event
	if err := initializers.DB.Where("id = ?", parsedEventID).First(&event).Error; err != nil {
		if err == gorm.ErrRecordNotFound{
			return &fiber.Error{Code: 404, Message: "Event not found"}
		}
		return &helpers.AppError{Code:500, Message: config.DATABASE_ERROR, LogMessage: err.Error(), Err: err}
	}

	if reqBody.Name != "" {
		event.Name = reqBody.Name
	}

	if reqBody.Date != nil {
		event.Date = *reqBody.Date
	}

	if reqBody.Location != "" {
		event.Location = reqBody.Location
	}

	if err := initializers.DB.Save(&event).Error; err != nil {
		return &helpers.AppError{Code:500, Message: config.DATABASE_ERROR, LogMessage: err.Error(), Err: err}
	}

	return c.Status(200).JSON(fiber.Map{
		"status": "success",
		"message": "",
		"event": event,
	})
}

func DeleteEvent(c *fiber.Ctx) error {
	parsedEventID, _ := uuid.Parse(c.Params("eventID"))
	parsedLoggedInUserID, _ := uuid.Parse(c.GetRespHeader("loggedInUserID"))

	var eventMemership models.EventMembership
	if err := initializers.DB.Where("event_id = ? AND user_id = ? AND role = ?", parsedEventID, parsedLoggedInUserID, models.Admin).First(&eventMemership).Error; err != nil {
		if err == gorm.ErrRecordNotFound{
			return &fiber.Error{Code: 400, Message: "You are not an admin of this event"}
		}
		return helpers.AppError{Code:500, Message: config.DATABASE_ERROR, LogMessage: err.Error(), Err: err}
	}

	if err := initializers.DB.Delete(&models.Event{}, "id = ?", parsedEventID).Error; err != nil {
		return helpers.AppError{Code:500, Message: config.DATABASE_ERROR, LogMessage: err.Error(), Err: err}
	}

	return c.Status(200).JSON(fiber.Map{
		"status": "success",
		"message": "Event Deleted",
	})
}