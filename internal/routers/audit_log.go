package routers

import (
	"csi-accounts/internal/crud"

	"github.com/gofiber/fiber/v2"
)

func AuditLogRouter(app *fiber.App) {
	auditLogRouter := app.Group("/audit-logs")

	auditLogRouter.Get("/", crud.GetAuditLogs)
	auditLogRouter.Get("/:auditLogID", crud.GetAuditLog)
	auditLogRouter.Post("/", crud.CreateAuditLog)
	auditLogRouter.Patch("/:auditLogID", crud.UpdateAuditLog)
	auditLogRouter.Delete("/:auditLogID", crud.DeleteAuditLog)
}
