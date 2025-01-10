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

// GetClients retrieves all clients
func GetClients(c *fiber.Ctx) error {
	var clients []models.Client
	if err := initializers.DB.Find(&clients).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return &fiber.Error{Code: 404, Message: "No clients found"}
		}
		return &helpers.AppError{Code: 500, Message: config.DATABASE_ERROR, LogMessage: err.Error(), Err: err}
	}

	return c.Status(200).JSON(fiber.Map{
		"status":  "success",
		"message": "",
		"clients": clients,
	})
}

// GetClient retrieves a single client by ID
func GetClient(c *fiber.Ctx) error {
	parsedClientID, err := uuid.Parse(c.Params("clientID"))
	if err != nil {
		return &fiber.Error{Code: 400, Message: "Invalid client ID"}
	}

	var client models.Client
	if err := initializers.DB.Where("id = ?", parsedClientID).First(&client).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return &fiber.Error{Code: 404, Message: "Client not found"}
		}
		return &helpers.AppError{Code: 500, Message: config.DATABASE_ERROR, LogMessage: err.Error(), Err: err}
	}

	return c.Status(200).JSON(fiber.Map{
		"status":  "success",
		"message": "",
		"client":  client,
	})
}

// CreateClient creates a new client
func CreateClient(c *fiber.Ctx) error {
	var reqBody models.Client
	if err := c.BodyParser(&reqBody); err != nil {
		return &fiber.Error{Code: 400, Message: "Invalid request body"}
	}

	reqBody.ID = uuid.New()
	reqBody.ClientID = helpers.GenerateRandomString(32)    // Replace with your method for generating ClientID
	reqBody.ClientSecret = helpers.GenerateRandomString(64) // Replace with your method for generating ClientSecret

	if err := initializers.DB.Create(&reqBody).Error; err != nil {
		return &helpers.AppError{Code: 500, Message: config.DATABASE_ERROR, LogMessage: err.Error(), Err: err}
	}

	return c.Status(201).JSON(fiber.Map{
		"status":  "success",
		"message": "Client created successfully",
		"client":  reqBody,
	})
}

// UpdateClient updates an existing client by ID
func UpdateClient(c *fiber.Ctx) error {
	parsedClientID, err := uuid.Parse(c.Params("clientID"))
	if err != nil {
		return &fiber.Error{Code: 400, Message: "Invalid client ID"}
	}

	var reqBody models.Client
	if err := c.BodyParser(&reqBody); err != nil {
		return &fiber.Error{Code: 400, Message: "Invalid request body"}
	}

	var client models.Client
	if err := initializers.DB.Where("id = ?", parsedClientID).First(&client).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return &fiber.Error{Code: 404, Message: "Client not found"}
		}
		return &helpers.AppError{Code: 500, Message: config.DATABASE_ERROR, LogMessage: err.Error(), Err: err}
	}

	client.Name = reqBody.Name
	client.Description = reqBody.Description
	client.RedirectURIs = reqBody.RedirectURIs

	if err := initializers.DB.Save(&client).Error; err != nil {
		return &helpers.AppError{Code: 500, Message: config.DATABASE_ERROR, LogMessage: err.Error(), Err: err}
	}

	return c.Status(200).JSON(fiber.Map{
		"status":  "success",
		"message": "Client updated successfully",
		"client":  client,
	})
}

// DeleteClient deletes a client by ID
func DeleteClient(c *fiber.Ctx) error {
	parsedClientID, err := uuid.Parse(c.Params("clientID"))
	if err != nil {
		return &fiber.Error{Code: 400, Message: "Invalid client ID"}
	}

	if err := initializers.DB.Delete(&models.Client{}, "id = ?", parsedClientID).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return &fiber.Error{Code: 404, Message: "Client not found"}
		}
		return &helpers.AppError{Code: 500, Message: config.DATABASE_ERROR, LogMessage: err.Error(), Err: err}
	}

	return c.Status(200).JSON(fiber.Map{
		"status":  "success",
		"message": "Client deleted successfully",
	})
}
