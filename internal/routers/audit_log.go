package routers

import (
	"csi-accounts/internal/controllers"

	"github.com/gofiber/fiber/v2"
)

func AuditLogRouter(app *fiber.App) {
	auditLogRouter := app.Group("/audit-logs")

	auditLogRouter.Get("/", controllers.GetAuditLogs)
	auditLogRouter.Get("/:auditLogID", controllers.GetAuditLog)
	auditLogRouter.Post("/", controllers.CreateAuditLog)
	auditLogRouter.Patch("/:auditLogID", controllers.UpdateAuditLog)
	auditLogRouter.Delete("/:auditLogID", controllers.DeleteAuditLog)
}
