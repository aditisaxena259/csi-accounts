package routers

import (

	"github.com/gofiber/fiber/v2"
	// Assuming routes are defined in the "internal/routes" package
)



// SetUp initializes the routes for the application
func SetUp(app *fiber.App) {
	EventRouter(app)
}
