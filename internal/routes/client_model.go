package routes
//client_model
import (
    "csi-accounts/internal/models"
    "csi-accounts/database"
    "net/http"

    "github.com/gofiber/fiber/v2"
)

// CreateClient handles creating a new client
func CreateClient(c *fiber.Ctx) error {
    var client models.Client

    if err := c.BodyParser(&client); err != nil {
        return c.Status(http.StatusBadRequest).JSON(fiber.Map{
            "error": "Invalid request body",
        })
    }

    // Ensure required fields are provided
    if client.Name == "" || client.ClientID == "" || client.ClientSecret == "" {
        return c.Status(http.StatusBadRequest).JSON(fiber.Map{
            "error": "Name, ClientID, and ClientSecret are required fields",
        })
    }

    if err := database.DB.Create(&client).Error; err != nil {
        return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
            "error": "Failed to create client",
        })
    }

    return c.Status(http.StatusOK).JSON(client)
}

// GetClient retrieves a client by its ID
func GetClient(c *fiber.Ctx) error {
    id := c.Params("id")
    var client models.Client

    if err := database.DB.Preload("ClientScopes").First(&client, "id = ?", id).Error; err != nil {
        return c.Status(http.StatusNotFound).JSON(fiber.Map{
            "error": "Client not found",
        })
    }

    return c.Status(http.StatusOK).JSON(client)
}

// UpdateClient updates a client's details
func UpdateClient(c *fiber.Ctx) error {
    id := c.Params("id")
    var client models.Client

    if err := database.DB.First(&client, "id = ?", id).Error; err != nil {
        return c.Status(http.StatusNotFound).JSON(fiber.Map{
            "error": "Client not found",
        })
    }

    if err := c.BodyParser(&client); err != nil {
        return c.Status(http.StatusBadRequest).JSON(fiber.Map{
            "error": "Invalid request body",
        })
    }

    if client.Name == "" || client.ClientID == "" || client.ClientSecret == "" {
        return c.Status(http.StatusBadRequest).JSON(fiber.Map{
            "error": "Name, ClientID, and ClientSecret are required fields",
        })
    }

    if err := database.DB.Save(&client).Error; err != nil {
        return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
            "error": "Failed to update client",
        })
    }

    return c.Status(http.StatusOK).JSON(client)
}

// DeleteClient deletes a client by its ID
func DeleteClient(c *fiber.Ctx) error {
    id := c.Params("id")

    if err := database.DB.Delete(&models.Client{}, "id = ?", id).Error; err != nil {
        return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
            "error": "Failed to delete client",
        })
    }

    return c.Status(http.StatusOK).JSON(fiber.Map{
        "message": "Client deleted successfully",
    })
}

// GetAllClients retrieves all clients
func GetAllClients(c *fiber.Ctx) error {
    var clients []models.Client

    if err := database.DB.Preload("ClientScopes").Find(&clients).Error; err != nil {
        return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
            "error": "Failed to retrieve clients",
        })
    }

    return c.Status(http.StatusOK).JSON(clients)
}

