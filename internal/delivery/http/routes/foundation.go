package routes

import (
	"boilerplate/internal/delivery/http/handlers"
	"boilerplate/internal/usecase"

	"github.com/gofiber/fiber/v2"
)

// SetupFoundationRoutes sets up routes for foundation endpoints
func SetupFoundationRoutes(app *fiber.App, usecase *usecase.UsecaseManager) {
	// Create handler
	foundationHandler := handlers.NewFoundationHandler(usecase)

	// API group
	api := app.Group("/api/v1")

	// User routes
	users := api.Group("/users")
	users.Post("/", foundationHandler.CreateUser)
	users.Get("/", foundationHandler.ListUsers)
	users.Get("/:id", foundationHandler.GetUser)
	users.Put("/:id", foundationHandler.UpdateUser)
	users.Delete("/:id", foundationHandler.DeleteUser)

	// Outlet routes
	outlets := api.Group("/outlets")
	outlets.Post("/", foundationHandler.CreateOutlet)
	outlets.Get("/", foundationHandler.ListOutlets)
	outlets.Get("/:id", foundationHandler.GetOutlet)
}