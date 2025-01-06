package routers

import (
	"csi-accounts/internal/controllers"

	"github.com/gofiber/fiber/v2"
)

func EventMembershipRoutes(app *fiber.App) {
	eventMemberships := app.Group("/event-memberships")

	eventMemberships.Get("/:eventID", controllers.GetEventMemberships)         // GET /event-memberships/:eventID
	eventMemberships.Put("/:eventID/coordinator/:userID", controllers.AddEventCoordinator) // PUT /event-memberships/:eventID/coordinator/:userID
}
