package main

import (
	"log"
	"os"

	
	"csi-accounts/internal/initializers"
	"csi-accounts/internal/routers"
	"csi-accounts/pkg/helpers"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func init() {
	initializers.LoadEnv()
	initializers.ConnectToDB()
	initializers.RunMigrations()
	initializers.CustomZerologLogger()
}

func main() {
	// ✅ Create a new Fiber app with a custom error handler
	app := fiber.New(fiber.Config{
		ErrorHandler: helpers.ErrorHandler,
	})

	// ✅ Logging middleware
	app.Use(logger.New())

	// ✅ CORS Middleware - Allow frontend requests
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*", // Allow all origins (change for production)
		AllowMethods: "GET, POST, PUT, DELETE, OPTIONS",
		AllowHeaders: "Content-Type, Authorization",
	}))

	// ✅ Handle OPTIONS preflight requests globally
	// app.Options("/*", func(c *fiber.Ctx) error {
	// 	return c.SendStatus(fiber.StatusNoContent)
	// })

	// // ✅ Allow both GET and POST on "/"
	// app.Get("/", func(c *fiber.Ctx) error {
	// 	return c.SendString("Hello, World!")
	// })
	// app.Post("/auth", func(c *fiber.Ctx) error {
	// 	authCode := helpers.GenerateAuthorizationCode() // ✅ Generate a real auth code
	// 	return c.JSON(fiber.Map{
	// 		"message":            "✅ POST request received!",
	// 		"status":             "success",
	// 		"authorization_code": authCode, // ✅ Correct response
	// 	})
	// })	
	// app.Get("/auth", func(c *fiber.Ctx) error {
	// 	authCode := helpers.GenerateAuthorizationCode() // Generate code
	// 	return c.JSON(fiber.Map{"authorization_code": authCode}) // ✅ Return correct key
	// })
	routers.SetUpAuthRoutes(app)

	// ✅ Register API routes
	routers.SetUp(app)

	// ✅ Validate PORT before starting server
	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("❌ PORT is not set in environment variables")
	}

	log.Println("🚀 Server running on port:", port)
	log.Fatal(app.Listen(":" + port))
}
