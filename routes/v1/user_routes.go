package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/sunnysingha911/user-service/handlers"
	"github.com/sunnysingha911/user-service/middlewares"
)

func UserRoutes(app fiber.Router) {
	// Public routes
	app.Post("/signup", handlers.SignUp)
	app.Post("/signin", handlers.SignIn)

	// Protected routes (JWT middleware)
	protected := app.Group("/", middlewares.AuthRequired)
	protected.Get("/me", handlers.GetMe)
	protected.Get("/all-users", handlers.GetAllUsers)
}
