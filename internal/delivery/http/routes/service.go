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

	// ============= Service Job Routes =============
	
	// Service Job routes
	serviceJobs := api.Group("/service-jobs")
	serviceJobs.Post("/", serviceHandler.CreateServiceJob)
	serviceJobs.Get("/", serviceHandler.ListServiceJobs)
	serviceJobs.Get("/service-code", serviceHandler.GetServiceJobByServiceCode)
	serviceJobs.Get("/status", serviceHandler.GetServiceJobsByStatus)
	serviceJobs.Get("/:id", serviceHandler.GetServiceJob)
	serviceJobs.Put("/:id", serviceHandler.UpdateServiceJob)
	serviceJobs.Put("/:id/status", serviceHandler.UpdateServiceJobStatus)
	serviceJobs.Delete("/:id", serviceHandler.DeleteServiceJob)

	// Customer-specific service job routes
	customers := api.Group("/customers")
	customers.Get("/:customer_id/service-jobs", serviceHandler.GetServiceJobsByCustomer)

	// ============= Service Detail Routes =============
	
	// Service Detail routes
	serviceDetails := api.Group("/service-details")
	serviceDetails.Post("/", serviceHandler.CreateServiceDetail)
	serviceDetails.Put("/:id", serviceHandler.UpdateServiceDetail)
	serviceDetails.Delete("/:id", serviceHandler.DeleteServiceDetail)

	// Service job specific service details
	serviceJobs.Get("/:service_job_id/details", serviceHandler.GetServiceDetailsByServiceJob)

	// ============= Service Job History Routes =============
	
	// Service job specific histories
	serviceJobs.Get("/:service_job_id/histories", serviceHandler.GetServiceJobHistoriesByServiceJob)

	// ============= Queue Management Routes =============
	
	// Queue management endpoints
	api.Get("/queue/:outlet_id", serviceHandler.GetServiceJobQueue)
	api.Get("/queue/:outlet_id/today", serviceHandler.GetTodayServiceJobQueue)
	api.Put("/queue/:outlet_id/reorder", serviceHandler.ReorderServiceJobQueue)
}