package routers

import (
	"csi-accounts/internal/controllers"
	"csi-accounts/internal/middlewares"

	"github.com/gofiber/fiber/v2"
)

func EventRouter(app *fiber.App) {
	eventRouter := app.Group("/events")

	eventRouter.Get("/", controllers.GetEvents)
	eventRouter.Get("/:eventID", controllers.GetEvent)
	eventRouter.Post("/", controllers.CreateEvent)
	eventRouter.Patch("/:eventID", controllers.UpdateEvent)
	eventRouter.Delete("/:eventID", controllers.DeleteEvent)

	eventMembershipRouter := eventRouter.Group("/:eventID/memberships")
	eventMembershipRouter.Get("/", controllers.GetEventMemberships)
	eventMembershipRouter.Post("/coordinators/:userID", middlewares.EventAdminAuthorization, controllers.AddEventCoordinator)
}