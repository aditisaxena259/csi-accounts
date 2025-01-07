package routers

import (
	"csi-accounts/internal/controllers"

	"github.com/gofiber/fiber/v2"
)

func RegisterPermissionRoutes(app *fiber.App) {
	permissionGroup := app.Group("/permissions")

	permissionGroup.Get("/", controllers.GetPermissions)
	permissionGroup.Get("/:permissionID", controllers.GetPermission)
	permissionGroup.Post("/", controllers.CreatePermission)
	permissionGroup.Put("/:permissionID", controllers.UpdatePermission)
	permissionGroup.Delete("/:permissionID", controllers.DeletePermission)
}
