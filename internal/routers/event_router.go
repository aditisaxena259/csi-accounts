package routers

import (
	"csi-accounts/internal/crud"
	"csi-accounts/internal/middlewares"
	"github.com/gofiber/fiber/v2"
)

func EventRouter(app *fiber.App) {
	eventRouter := app.Group("/events")

	eventRouter.Get("/", crud.GetEvents)
	eventRouter.Get("/:eventID", crud.GetEvent)
	eventRouter.Post("/", crud.CreateEvent)
	eventRouter.Patch("/:eventID", crud.UpdateEvent)
	eventRouter.Delete("/:eventID", crud.DeleteEvent)

	eventMembershipRouter := eventRouter.Group("/:eventID/memberships")
	eventMembershipRouter.Get("/", crud.GetEventMemberships)
	eventMembershipRouter.Post("/coordinators/:userID", middlewares.EventAdminAuthorization, crud.AddEventCoordinator)
}