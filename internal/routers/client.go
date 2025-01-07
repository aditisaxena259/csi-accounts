package routers

import (
	"csi-accounts/internal/controllers"
	"csi-accounts/internal/middlewares"
	"github.com/gofiber/fiber/v2"
)

func ClientRouter(app *fiber.App) {
	eventRouter := app.Group("/events")

	eventRouter.Get("/", controllers.GetClients)
	eventRouter.Get("/:eventID", controllers.GetClient)
	eventRouter.Post("/", controllers.CreateClient)
	eventRouter.Patch("/:eventID", controllers.UpdateClient)
	eventRouter.Delete("/:eventID", controllers.DeleteClient)

	eventMembershipRouter := eventRouter.Group("/:eventID/memberships")
	eventMembershipRouter.Get("/", controllers.GetEventMemberships)
	eventMembershipRouter.Post("/coordinators/:userID", middlewares.EventAdminAuthorization, controllers.AddEventCoordinator)
}