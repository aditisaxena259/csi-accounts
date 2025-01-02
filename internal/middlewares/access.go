package middlewares

import (
	"csi-accounts/internal/initializers"
	"csi-accounts/internal/models"
	"csi-accounts/pkg/config"
	"csi-accounts/pkg/helpers"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func EventAdminAuthorization(c *fiber.Ctx) error {
	eventID := c.Params("eventID")
	loggedInUserID := c.GetRespHeader("loggedInUserID")

	parsedEventID, err := uuid.Parse(eventID)
	if err != nil {
		return &fiber.Error{Code: 400, Message: "Invalid event ID"}
	}
	parsedLoggedInUserID, err := uuid.Parse(loggedInUserID)
	if err != nil {
		return &fiber.Error{Code: 400, Message: "Invalid user ID"}
	}

	var eventMemership models.EventMembership
	if err := initializers.DB.Where("event_id = ? AND user_id = ? AND role = ?", parsedEventID, parsedLoggedInUserID, models.Admin).First(&eventMemership).Error; err != nil {
		if err == gorm.ErrRecordNotFound{
			return &fiber.Error{Code: 400, Message: "You are not an admin of this event"}
		}
		return helpers.AppError{Code:500, Message: config.DATABASE_ERROR, LogMessage: err.Error(), Err: err}
	}
	
	return c.Next()
}