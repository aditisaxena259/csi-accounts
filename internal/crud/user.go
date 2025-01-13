package crud

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

// GetUsers retrieves all users
func GetUsers(c *fiber.Ctx) error {
	var users []models.User
	if err := initializers.DB.Find(&users).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return &fiber.Error{Code: 404, Message: "Users not found"}
		}
		return &helpers.AppError{Code: 500, Message: config.DATABASE_ERROR, LogMessage: err.Error(), Err: err}
	}

	return c.Status(200).JSON(fiber.Map{
		"status": "success",
		"message": "",
		"users": users,
	})
}

// GetUser retrieves a single user by ID
func GetUser(c *fiber.Ctx) error {
	parsedUserID, err := uuid.Parse(c.Params("userID"))
	if err != nil {
		return &fiber.Error{Code: 400, Message: "Invalid user ID"}
	}

	var user models.User
	if err := initializers.DB.Where("id = ?", parsedUserID).First(&user).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return &fiber.Error{Code: 404, Message: "User not found"}
		}
		return &helpers.AppError{Code: 500, Message: config.DATABASE_ERROR, LogMessage: err.Error(), Err: err}
	}

	return c.Status(200).JSON(fiber.Map{
		"status": "success",
		"message": "",
		"user": user,
	})
}

// CreateUser creates a new user
func CreateUser(c *fiber.Ctx) error {
	var reqBody schemas.CreateUserBody
	if err := c.BodyParser(&reqBody); err != nil {
		return &fiber.Error{Code: 400, Message: "Invalid request body"}
	}

	user := models.User{
		ID:           uuid.New(),
		Name:         reqBody.Name,
		Username:     reqBody.Username,
		Email:        reqBody.Email,
		PasswordHash: reqBody.PasswordHash, // You can hash the password here if needed
		Status:       models.ACTIVE,
		RoleID:       reqBody.RoleID,
	}

	if err := initializers.DB.Create(&user).Error; err != nil {
		return &helpers.AppError{Code: 500, Message: config.DATABASE_ERROR, LogMessage: err.Error(), Err: err}
	}

	return c.Status(201).JSON(fiber.Map{
		"status": "success",
		"message": "",
		"user": user,
	})
}

// UpdateUser updates an existing user
func UpdateUser(c *fiber.Ctx) error {
	parsedUserID, err := uuid.Parse(c.Params("userID"))
	if err != nil {
		return &fiber.Error{Code: 400, Message: "Invalid user ID"}
	}

	var reqBody schemas.UpdateUserBody
	if err := c.BodyParser(&reqBody); err != nil {
		return &fiber.Error{Code: 400, Message: "Invalid request body"}
	}

	var user models.User
	if err := initializers.DB.Where("id = ?", parsedUserID).First(&user).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return &fiber.Error{Code: 404, Message: "User not found"}
		}
		return &helpers.AppError{Code: 500, Message: config.DATABASE_ERROR, LogMessage: err.Error(), Err: err}
	}

	// Update user fields
	if reqBody.Name != "" {
		user.Name = reqBody.Name
	}

	if reqBody.Username != "" {
		user.Username = reqBody.Username
	}

	if reqBody.Email != "" {
		user.Email = reqBody.Email
	}

	if reqBody.Status != "" {
		user.Status = models.Status(reqBody.Status) 
	}

	if err := initializers.DB.Save(&user).Error; err != nil {
		return &helpers.AppError{Code: 500, Message: config.DATABASE_ERROR, LogMessage: err.Error(), Err: err}
	}

	return c.Status(200).JSON(fiber.Map{
		"status": "success",
		"message": "",
		"user": user,
	})
}

// DeleteUser deletes a user by ID
func DeleteUser(c *fiber.Ctx) error {
	parsedUserID, err := uuid.Parse(c.Params("userID"))
	if err != nil {
		return &fiber.Error{Code: 400, Message: "Invalid user ID"}
	}

	if err := initializers.DB.Delete(&models.User{}, "id = ?", parsedUserID).Error; err != nil {
		return &helpers.AppError{Code: 500, Message: config.DATABASE_ERROR, LogMessage: err.Error(), Err: err}
	}

	return c.Status(200).JSON(fiber.Map{
		"status": "success",
		"message": "User Deleted",
	})
}
