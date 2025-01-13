package routers

import (
	"csi-accounts/internal/crud"

	"github.com/gofiber/fiber/v2"
)

func RegisterUserRoutes(app *fiber.App) {
	userGroup := app.Group("/users")
	userGroup.Get("/", crud.GetUsers)
	userGroup.Get("/:userID", crud.GetUser)
	userGroup.Post("/", crud.CreateUser)
	userGroup.Put("/:userID", crud.UpdateUser)
	userGroup.Delete("/:userID", crud.DeleteUser)
}
