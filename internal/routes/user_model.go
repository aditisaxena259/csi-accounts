package routes
//user_model
import (
    "csi-accounts/database"
    "csi-accounts/internal/models"
    "net/http"

    "github.com/gofiber/fiber/v2"
    "github.com/google/uuid"
)

type User struct {
    ID    uuid.UUID `json:"id"`
    Name  string    `json:"name"`
    Email string    `json:"email"`
}

func CreateUser(c *fiber.Ctx) error {
    var user models.User

    if err := c.BodyParser(&user); err != nil {
        return c.Status(http.StatusBadRequest).JSON(fiber.Map{
            "error": "Invalid request body",
        })
    }

    if err := database.DB.Create(&user).Error; err != nil {
        return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
            "error": "Failed to create user",
        })
    }

    return c.Status(http.StatusOK).JSON(user)
}

func GetUser(c *fiber.Ctx) error {
    id := c.Params("id")
    var user models.User

    if err := database.DB.First(&user, "id = ?", id).Error; err != nil {
        return c.Status(http.StatusNotFound).JSON(fiber.Map{
            "error": "User not found",
        })
    }

    return c.Status(http.StatusOK).JSON(user)
}

func UpdateUser(c *fiber.Ctx) error {
    id := c.Params("id")
    var user models.User

    if err := database.DB.First(&user, "id = ?", id).Error; err != nil {
        return c.Status(http.StatusNotFound).JSON(fiber.Map{
            "error": "User not found",
        })
    }

    if err := c.BodyParser(&user); err != nil {
        return c.Status(http.StatusBadRequest).JSON(fiber.Map{
            "error": "Invalid request body",
        })
    }

    if err := database.DB.Save(&user).Error; err != nil {
        return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
            "error": "Failed to update user",
        })
    }

    return c.Status(http.StatusOK).JSON(user)
}

func DeleteUser(c *fiber.Ctx) error {
    id := c.Params("id")

    if err := database.DB.Delete(&models.User{}, "id = ?", id).Error; err != nil {
        return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
            "error": "Failed to delete user",
        })
    }

    return c.Status(http.StatusOK).JSON(fiber.Map{
        "message": "User deleted successfully",
    })
}

