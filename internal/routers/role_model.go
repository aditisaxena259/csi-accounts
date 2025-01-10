package routers 
import (
	"csi-accounts/internal/crud"

	"github.com/gofiber/fiber/v2"
)

func RegisterRoleRoutes(app *fiber.App) {
	roleGroup := app.Group("/roles")

	roleGroup.Get("/", crud.GetRoles)
	roleGroup.Get("/:roleID", crud.GetRole)
	roleGroup.Post("/", crud.CreateRole)
	roleGroup.Post("/:roleID/permissions", crud.AssignPermissionsToRole)
	roleGroup.Delete("/:roleID", crud.DeleteRole)
}
