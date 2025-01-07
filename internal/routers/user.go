package routers

import (
	"csi-accounts/internal/controllers"

	"github.com/gofiber/fiber/v2"
)

func UserRoutes(app *fiber.App) {
	userGroup := app.Group("/users")
	userGroup.Get("/", controllers.GetUsers)
	userGroup.Get("/:userID", controllers.GetUser)
	userGroup.Post("/", controllers.CreateUser)
	userGroup.Put("/:userID", controllers.UpdateUser)
	userGroup.Delete("/:userID", controllers.DeleteUser)
}
