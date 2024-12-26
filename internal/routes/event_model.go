package routes
//event_model
import (
    "csi-accounts/database"
    "csi-accounts/internal/models"
    "net/http"
    "github.com/gofiber/fiber/v2"
)

// CreateEvent handles the creation of an event
func CreateEvent(c *fiber.Ctx) error {
    var event models.Event

    if err := c.BodyParser(&event); err != nil {
        return c.Status(http.StatusBadRequest).JSON(fiber.Map{
            "error": "Invalid request body",
        })
    }

    if err := database.DB.Create(&event).Error; err != nil {
        return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
            "error": "Failed to create event",
        })
    }

    return c.Status(http.StatusOK).JSON(event)
}

// GetEvent handles retrieving an event by ID
func GetEvent(c *fiber.Ctx) error {
    id := c.Params("id")
    var event models.Event

    if err := database.DB.First(&event, "id = ?", id).Error; err != nil {
        return c.Status(http.StatusNotFound).JSON(fiber.Map{
            "error": "Event not found",
        })
    }

    return c.Status(http.StatusOK).JSON(event)
}

// UpdateEvent handles updating an event by ID
func UpdateEvent(c *fiber.Ctx) error {
    id := c.Params("id")
    var event models.Event

    if err := database.DB.First(&event, "id = ?", id).Error; err != nil {
        return c.Status(http.StatusNotFound).JSON(fiber.Map{
            "error": "Event not found",
        })
    }

    if err := c.BodyParser(&event); err != nil {
        return c.Status(http.StatusBadRequest).JSON(fiber.Map{
            "error": "Invalid request body",
        })
    }

    if err := database.DB.Save(&event).Error; err != nil {
        return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
            "error": "Failed to update event",
        })
    }

    return c.Status(http.StatusOK).JSON(event)
}

// DeleteEvent handles deleting an event by ID
func DeleteEvent(c *fiber.Ctx) error {
    id := c.Params("id")

    if err := database.DB.Delete(&models.Event{}, "id = ?", id).Error; err != nil {
        return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
            "error": "Failed to delete event",
        })
    }

    return c.Status(http.StatusOK).JSON(fiber.Map{
        "message": "Event deleted successfully",
    })
}

// CreateEventMembership handles the creation of an event membership
func CreateEventMembership(c *fiber.Ctx) error {
    var membership models.EventMembership

    if err := c.BodyParser(&membership); err != nil {
        return c.Status(http.StatusBadRequest).JSON(fiber.Map{
            "error": "Invalid request body",
        })
    }

    if err := database.DB.Create(&membership).Error; err != nil {
        return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
            "error": "Failed to create event membership",
        })
    }

    return c.Status(http.StatusOK).JSON(membership)
}

// GetEventMembership handles retrieving an event membership by ID
func GetEventMembership(c *fiber.Ctx) error {
    id := c.Params("id")
    var membership models.EventMembership

    if err := database.DB.First(&membership, "id = ?", id).Error; err != nil {
        return c.Status(http.StatusNotFound).JSON(fiber.Map{
            "error": "Event membership not found",
        })
    }

    return c.Status(http.StatusOK).JSON(membership)
}

// GetEventMembershipsByEventID handles retrieving all event memberships for a specific event
func GetEventMembershipsByEventID(c *fiber.Ctx) error {
    eventID := c.Params("eventID")
    var memberships []models.EventMembership

    if err := database.DB.Find(&memberships, "event_id = ?", eventID).Error; err != nil {
        return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
            "error": "Failed to retrieve event memberships",
        })
    }

    return c.Status(http.StatusOK).JSON(memberships)
}

// UpdateEventMembership handles updating an event membership by ID
func UpdateEventMembership(c *fiber.Ctx) error {
    id := c.Params("id")
    var membership models.EventMembership

    if err := database.DB.First(&membership, "id = ?", id).Error; err != nil {
        return c.Status(http.StatusNotFound).JSON(fiber.Map{
            "error": "Event membership not found",
        })
    }

    if err := c.BodyParser(&membership); err != nil {
        return c.Status(http.StatusBadRequest).JSON(fiber.Map{
            "error": "Invalid request body",
        })
    }

    if err := database.DB.Save(&membership).Error; err != nil {
        return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
            "error": "Failed to update event membership",
        })
    }

    return c.Status(http.StatusOK).JSON(membership)
}

// DeleteEventMembership handles deleting an event membership by ID
func DeleteEventMembership(c *fiber.Ctx) error {
    id := c.Params("id")

    if err := database.DB.Delete(&models.EventMembership{}, "id = ?", id).Error; err != nil {
        return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
            "error": "Failed to delete event membership",
        })
    }

    return c.Status(http.StatusOK).JSON(fiber.Map{
        "message": "Event membership deleted successfully",
    })
}

