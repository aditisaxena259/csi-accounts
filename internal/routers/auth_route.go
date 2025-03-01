package routers

import (
	"csi-accounts/internal/controllers"
	"github.com/gofiber/fiber/v2"
)

func SetUpAuthRoutes(app *fiber.App) {
	auth := app.Group("/auth")
	auth.Post("/signup", controllers.Signup)
	auth.Post("/login", controllers.Login)
}
