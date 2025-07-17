package routes

import (
	"boilerplate/internal/delivery/http/handlers"
	"boilerplate/internal/usecase"

	"github.com/gofiber/fiber/v2"
)

// SetupFinancialRoutes sets up routes for financial endpoints
func SetupFinancialRoutes(app *fiber.App, usecase *usecase.UsecaseManager) {
	// Create handler
	financialHandler := handlers.NewFinancialHandler(usecase)

	// API group
	api := app.Group("/api/v1")

	// Payment Method routes
	paymentMethods := api.Group("/payment-methods")
	paymentMethods.Post("/", financialHandler.CreatePaymentMethod)
	paymentMethods.Get("/", financialHandler.ListPaymentMethods)
	paymentMethods.Get("/:id", financialHandler.GetPaymentMethod)
	paymentMethods.Put("/:id", financialHandler.UpdatePaymentMethod)
	paymentMethods.Delete("/:id", financialHandler.DeletePaymentMethod)

	// Transaction routes
	transactions := api.Group("/transactions")
	transactions.Post("/", financialHandler.CreateTransaction)
	transactions.Get("/", financialHandler.ListTransactions)
	transactions.Get("/invoice", financialHandler.GetTransactionByInvoiceNumber)
	transactions.Get("/status", financialHandler.GetTransactionsByStatus)
	transactions.Get("/date-range", financialHandler.GetTransactionsByDateRange)
	transactions.Get("/:id", financialHandler.GetTransaction)
	transactions.Put("/:id", financialHandler.UpdateTransaction)
	transactions.Delete("/:id", financialHandler.DeleteTransaction)

	// Cash Flow routes
	cashFlows := api.Group("/cash-flows")
	cashFlows.Post("/", financialHandler.CreateCashFlow)
	cashFlows.Get("/", financialHandler.ListCashFlows)
	cashFlows.Get("/type", financialHandler.GetCashFlowsByType)
	cashFlows.Get("/:id", financialHandler.GetCashFlow)
	cashFlows.Put("/:id", financialHandler.UpdateCashFlow)
	cashFlows.Delete("/:id", financialHandler.DeleteCashFlow)

	// Customer-specific transaction routes
	customers := api.Group("/customers")
	customers.Get("/:customer_id/transactions", financialHandler.GetTransactionsByCustomer)

	// Outlet-specific transaction routes
	outlets := api.Group("/outlets")
	outlets.Get("/:outlet_id/transactions", financialHandler.GetTransactionsByOutlet)
}