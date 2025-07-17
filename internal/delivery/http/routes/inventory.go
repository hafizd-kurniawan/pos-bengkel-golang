package routes

import (
	"boilerplate/internal/delivery/http/handlers"
	"boilerplate/internal/usecase"

	"github.com/gofiber/fiber/v2"
)

// SetupInventoryRoutes sets up routes for inventory endpoints
func SetupInventoryRoutes(app *fiber.App, usecase *usecase.UsecaseManager) {
	// Create handler
	inventoryHandler := handlers.NewInventoryHandler(usecase)

	// API group
	api := app.Group("/api/v1")

	// Product routes
	products := api.Group("/products")
	products.Post("/", inventoryHandler.CreateProduct)
	products.Get("/", inventoryHandler.ListProducts)
	products.Get("/search", inventoryHandler.SearchProducts)
	products.Get("/sku", inventoryHandler.GetProductBySKU)
	products.Get("/barcode", inventoryHandler.GetProductByBarcode)
	products.Get("/usage-status", inventoryHandler.GetProductsByUsageStatus)
	products.Get("/low-stock", inventoryHandler.GetLowStockProducts)
	products.Get("/:id", inventoryHandler.GetProduct)
	products.Put("/:id", inventoryHandler.UpdateProduct)
	products.Delete("/:id", inventoryHandler.DeleteProduct)
	products.Post("/:id/stock", inventoryHandler.UpdateProductStock)

	// Category routes
	categories := api.Group("/categories")
	categories.Post("/", inventoryHandler.CreateCategory)
	categories.Get("/", inventoryHandler.ListCategories)
	categories.Get("/:id", inventoryHandler.GetCategory)
	categories.Put("/:id", inventoryHandler.UpdateCategory)
	categories.Delete("/:id", inventoryHandler.DeleteCategory)

	// Supplier routes
	suppliers := api.Group("/suppliers")
	suppliers.Post("/", inventoryHandler.CreateSupplier)
	suppliers.Get("/", inventoryHandler.ListSuppliers)
	suppliers.Get("/search", inventoryHandler.SearchSuppliers)
	suppliers.Get("/:id", inventoryHandler.GetSupplier)
	suppliers.Put("/:id", inventoryHandler.UpdateSupplier)
	suppliers.Delete("/:id", inventoryHandler.DeleteSupplier)

	// Unit Type routes
	unitTypes := api.Group("/unit-types")
	unitTypes.Post("/", inventoryHandler.CreateUnitType)
	unitTypes.Get("/", inventoryHandler.ListUnitTypes)
	unitTypes.Get("/:id", inventoryHandler.GetUnitType)
	unitTypes.Put("/:id", inventoryHandler.UpdateUnitType)
	unitTypes.Delete("/:id", inventoryHandler.DeleteUnitType)

	// Category-specific product routes
	categories.Get("/:category_id/products", inventoryHandler.GetProductsByCategory)

	// Supplier-specific product routes
	suppliers.Get("/:supplier_id/products", inventoryHandler.GetProductsBySupplier)
}