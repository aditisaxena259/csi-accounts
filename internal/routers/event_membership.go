package routers

import (
	"csi-accounts/internal/crud"

	"github.com/gofiber/fiber/v2"
)

func EventMembershipRoutes(app *fiber.App) {
	eventMemberships := app.Group("/event-memberships")

	eventMemberships.Get("/:eventID", crud.GetEventMemberships)         // GET /event-memberships/:eventID
	eventMemberships.Put("/:eventID/coordinator/:userID", crud.AddEventCoordinator) // PUT /event-memberships/:eventID/coordinator/:userID
}
