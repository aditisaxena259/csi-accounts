package main

import (
	"csi-accounts/internal/initializers"
	"csi-accounts/internal/routers"
	"csi-accounts/pkg/helpers"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func init() {
	initializers.LoadEnv()
	initializers.ConnectToDB()
	initializers.RunMigrations()
	initializers.CustomZerologLogger()
}
/*func SetUpRoutes(app *fiber.App)
{
	app.Get("/api",welcome)
	app.Post("/api/users",routes.createUser)
}*/
func main() {

	app := fiber.New(fiber.Config{
		ErrorHandler: helpers.ErrorHandler,
	})
	app.Use(logger.New())
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	routers.SetUp(app)

	app.Listen(":" + initializers.CONFIG.PORT)
}
