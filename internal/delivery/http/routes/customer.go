package routes

import (
	"boilerplate/internal/delivery/http/handlers"
	"boilerplate/internal/usecase"

	"github.com/gofiber/fiber/v2"
)

// SetupCustomerRoutes sets up routes for customer endpoints
func SetupCustomerRoutes(app *fiber.App, usecase *usecase.UsecaseManager) {
	// Create handler
	customerHandler := handlers.NewCustomerHandler(usecase)

	// API group
	api := app.Group("/api/v1")

	// Customer routes
	customers := api.Group("/customers")
	customers.Post("/", customerHandler.CreateCustomer)
	customers.Get("/", customerHandler.ListCustomers)
	customers.Get("/search", customerHandler.SearchCustomers)
	customers.Get("/phone", customerHandler.GetCustomerByPhoneNumber)
	customers.Get("/:id", customerHandler.GetCustomer)
	customers.Put("/:id", customerHandler.UpdateCustomer)
	customers.Delete("/:id", customerHandler.DeleteCustomer)

	// Customer vehicle routes
	customerVehicles := api.Group("/customer-vehicles")
	customerVehicles.Post("/", customerHandler.CreateCustomerVehicle)
	customerVehicles.Get("/", customerHandler.ListCustomerVehicles)
	customerVehicles.Get("/search", customerHandler.SearchCustomerVehicles)
	customerVehicles.Get("/:id", customerHandler.GetCustomerVehicle)
	customerVehicles.Put("/:id", customerHandler.UpdateCustomerVehicle)
	customerVehicles.Delete("/:id", customerHandler.DeleteCustomerVehicle)

	// Customer-specific vehicle routes
	customers.Get("/:customer_id/vehicles", customerHandler.GetCustomerVehiclesByCustomerID)
}