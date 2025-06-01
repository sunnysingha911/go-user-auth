// main.go
package main

import (
	"log"

	"github.com/gofiber/fiber/v2"

	"github.com/sunnysingha911/user-service/config"
	"github.com/sunnysingha911/user-service/database"
	"github.com/sunnysingha911/user-service/routes/v1"
)

func main() {

	if err := config.LoadEnv(); err != nil {
		log.Fatalf("Failed to load env: %v", err)
	}

	// Connect to DB
	if err := database.Connect(); err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	app := fiber.New()

	api := app.Group("/api/v1")

	routes.UserRoutes(api)

	app.Listen(":3000")
}
