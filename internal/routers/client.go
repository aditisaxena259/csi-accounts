package routers

import (
	"csi-accounts/internal/crud"
	"csi-accounts/internal/middlewares"
	"github.com/gofiber/fiber/v2"
)

func ClientRouter(app *fiber.App) {
	eventRouter := app.Group("/events")

	eventRouter.Get("/", crud.GetClients)
	eventRouter.Get("/:eventID", crud.GetClient)
	eventRouter.Post("/", crud.CreateClient)
	eventRouter.Patch("/:eventID", crud.UpdateClient)
	eventRouter.Delete("/:eventID", crud.DeleteClient)

	eventMembershipRouter := eventRouter.Group("/:eventID/memberships")
	eventMembershipRouter.Get("/", crud.GetEventMemberships)
	eventMembershipRouter.Post("/coordinators/:userID", middlewares.EventAdminAuthorization, crud.AddEventCoordinator)
}