package routers 
import (
	"csi-accounts/internal/controllers"

	"github.com/gofiber/fiber/v2"
)

func RegisterRoleRoutes(app *fiber.App) {
	roleGroup := app.Group("/roles")

	roleGroup.Get("/", controllers.GetRoles)
	roleGroup.Get("/:roleID", controllers.GetRole)
	roleGroup.Post("/", controllers.CreateRole)
	roleGroup.Post("/:roleID/permissions", controllers.AssignPermissionsToRole)
	roleGroup.Delete("/:roleID", controllers.DeleteRole)
}
