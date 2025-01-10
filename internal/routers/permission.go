package routers

import (
	"csi-accounts/internal/crud"

	"github.com/gofiber/fiber/v2"
)

func RegisterPermissionRoutes(app *fiber.App) {
	permissionGroup := app.Group("/permissions")

	permissionGroup.Get("/", crud.GetPermissions)
	permissionGroup.Get("/:permissionID", crud.GetPermission)
	permissionGroup.Post("/", crud.CreatePermission)
	permissionGroup.Put("/:permissionID", crud.UpdatePermission)
	permissionGroup.Delete("/:permissionID", crud.DeletePermission)
}
