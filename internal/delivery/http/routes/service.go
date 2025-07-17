package routes

import (
	"boilerplate/internal/delivery/http/handlers"
	"boilerplate/internal/usecase"

	"github.com/gofiber/fiber/v2"
)

// SetupServiceRoutes sets up routes for service endpoints
func SetupServiceRoutes(app *fiber.App, usecase *usecase.UsecaseManager) {
	// Create handler
	serviceHandler := handlers.NewServiceHandler(usecase)

	// API group
	api := app.Group("/api/v1")

	// Service Category routes
	serviceCategories := api.Group("/service-categories")
	serviceCategories.Post("/", serviceHandler.CreateServiceCategory)
	serviceCategories.Get("/", serviceHandler.ListServiceCategories)
	serviceCategories.Get("/:id", serviceHandler.GetServiceCategory)
	serviceCategories.Put("/:id", serviceHandler.UpdateServiceCategory)
	serviceCategories.Delete("/:id", serviceHandler.DeleteServiceCategory)

	// Service routes
	services := api.Group("/services")
	services.Post("/", serviceHandler.CreateService)
	services.Get("/", serviceHandler.ListServices)
	services.Get("/search", serviceHandler.SearchServices)
	services.Get("/code", serviceHandler.GetServiceByCode)
	services.Get("/:id", serviceHandler.GetService)
	services.Put("/:id", serviceHandler.UpdateService)
	services.Delete("/:id", serviceHandler.DeleteService)

	// Category-specific service routes
	serviceCategories.Get("/:category_id/services", serviceHandler.GetServicesByCategory)
}