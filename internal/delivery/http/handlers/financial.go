package handlers

import (
	"boilerplate/internal/delivery/http/responses"
	"boilerplate/internal/models"
	"boilerplate/internal/usecase"
	"boilerplate/internal/usecase/interfaces"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
)

// FinancialHandler handles financial-related HTTP requests
type FinancialHandler struct {
	usecase *usecase.UsecaseManager
}

// NewFinancialHandler creates a new financial handler
func NewFinancialHandler(usecase *usecase.UsecaseManager) *FinancialHandler {
	return &FinancialHandler{usecase: usecase}
}

// ============= Payment Method Handlers =============

// CreatePaymentMethod creates a new payment method
func (h *FinancialHandler) CreatePaymentMethod(c *fiber.Ctx) error {
	var req interfaces.CreatePaymentMethodRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(responses.Response{
			Status:  "error",
			Message: "Invalid request body",
			Error:   err.Error(),
		})
	}

	paymentMethod, err := h.usecase.PaymentMethod.CreatePaymentMethod(c.Context(), req)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(responses.Response{
			Status:  "error",
			Message: "Failed to create payment method",
			Error:   err.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(responses.Response{
		Status:  "success",
		Message: "Payment method created successfully",
		Data:    paymentMethod,
	})
}

// ListPaymentMethods lists all payment methods with pagination
func (h *FinancialHandler) ListPaymentMethods(c *fiber.Ctx) error {
	limit, _ := strconv.Atoi(c.Query("limit", "10"))
	offset, _ := strconv.Atoi(c.Query("offset", "0"))

	paymentMethods, err := h.usecase.PaymentMethod.ListPaymentMethods(c.Context(), limit, offset)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(responses.Response{
			Status:  "error",
			Message: "Failed to retrieve payment methods",
			Error:   err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(responses.Response{
		Status:  "success",
		Message: "Payment methods retrieved successfully",
		Data:    paymentMethods,
	})
}

// GetPaymentMethod retrieves a payment method by ID
func (h *FinancialHandler) GetPaymentMethod(c *fiber.Ctx) error {
	id, err := strconv.ParseUint(c.Params("id"), 10, 32)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(responses.Response{
			Status:  "error",
			Message: "Invalid payment method ID",
			Error:   err.Error(),
		})
	}

	paymentMethod, err := h.usecase.PaymentMethod.GetPaymentMethod(c.Context(), uint(id))
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(responses.Response{
			Status:  "error",
			Message: "Payment method not found",
			Error:   err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(responses.Response{
		Status:  "success",
		Message: "Payment method retrieved successfully",
		Data:    paymentMethod,
	})
}

// UpdatePaymentMethod updates a payment method
func (h *FinancialHandler) UpdatePaymentMethod(c *fiber.Ctx) error {
	id, err := strconv.ParseUint(c.Params("id"), 10, 32)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(responses.Response{
			Status:  "error",
			Message: "Invalid payment method ID",
			Error:   err.Error(),
		})
	}

	var req interfaces.UpdatePaymentMethodRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(responses.Response{
			Status:  "error",
			Message: "Invalid request body",
			Error:   err.Error(),
		})
	}

	paymentMethod, err := h.usecase.PaymentMethod.UpdatePaymentMethod(c.Context(), uint(id), req)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(responses.Response{
			Status:  "error",
			Message: "Failed to update payment method",
			Error:   err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(responses.Response{
		Status:  "success",
		Message: "Payment method updated successfully",
		Data:    paymentMethod,
	})
}

// DeletePaymentMethod deletes a payment method
func (h *FinancialHandler) DeletePaymentMethod(c *fiber.Ctx) error {
	id, err := strconv.ParseUint(c.Params("id"), 10, 32)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(responses.Response{
			Status:  "error",
			Message: "Invalid payment method ID",
			Error:   err.Error(),
		})
	}

	err = h.usecase.PaymentMethod.DeletePaymentMethod(c.Context(), uint(id))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(responses.Response{
			Status:  "error",
			Message: "Failed to delete payment method",
			Error:   err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(responses.Response{
		Status:  "success",
		Message: "Payment method deleted successfully",
		Data:    nil,
	})
}

// ============= Transaction Handlers =============

// CreateTransaction creates a new transaction
func (h *FinancialHandler) CreateTransaction(c *fiber.Ctx) error {
	var req interfaces.CreateTransactionRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(responses.Response{
			Status:  "error",
			Message: "Invalid request body",
			Error:   err.Error(),
		})
	}

	transaction, err := h.usecase.Transaction.CreateTransaction(c.Context(), req)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(responses.Response{
			Status:  "error",
			Message: "Failed to create transaction",
			Error:   err.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(responses.Response{
		Status:  "success",
		Message: "Transaction created successfully",
		Data:    transaction,
	})
}

// ListTransactions lists all transactions with pagination
func (h *FinancialHandler) ListTransactions(c *fiber.Ctx) error {
	limit, _ := strconv.Atoi(c.Query("limit", "10"))
	offset, _ := strconv.Atoi(c.Query("offset", "0"))

	transactions, err := h.usecase.Transaction.ListTransactions(c.Context(), limit, offset)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(responses.Response{
			Status:  "error",
			Message: "Failed to retrieve transactions",
			Error:   err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(responses.Response{
		Status:  "success",
		Message: "Transactions retrieved successfully",
		Data:    transactions,
	})
}

// GetTransaction retrieves a transaction by ID
func (h *FinancialHandler) GetTransaction(c *fiber.Ctx) error {
	id, err := strconv.ParseUint(c.Params("id"), 10, 32)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(responses.Response{
			Status:  "error",
			Message: "Invalid transaction ID",
			Error:   err.Error(),
		})
	}

	transaction, err := h.usecase.Transaction.GetTransaction(c.Context(), uint(id))
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(responses.Response{
			Status:  "error",
			Message: "Transaction not found",
			Error:   err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(responses.Response{
		Status:  "success",
		Message: "Transaction retrieved successfully",
		Data:    transaction,
	})
}

// GetTransactionByInvoiceNumber retrieves a transaction by invoice number
func (h *FinancialHandler) GetTransactionByInvoiceNumber(c *fiber.Ctx) error {
	invoiceNumber := c.Query("invoice_number")
	if invoiceNumber == "" {
		return c.Status(fiber.StatusBadRequest).JSON(responses.Response{
			Status:  "error",
			Message: "Invoice number is required",
			Error:   "missing invoice_number query parameter",
		})
	}

	transaction, err := h.usecase.Transaction.GetTransactionByInvoiceNumber(c.Context(), invoiceNumber)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(responses.Response{
			Status:  "error",
			Message: "Transaction not found",
			Error:   err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(responses.Response{
		Status:  "success",
		Message: "Transaction retrieved successfully",
		Data:    transaction,
	})
}

// UpdateTransaction updates a transaction
func (h *FinancialHandler) UpdateTransaction(c *fiber.Ctx) error {
	id, err := strconv.ParseUint(c.Params("id"), 10, 32)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(responses.Response{
			Status:  "error",
			Message: "Invalid transaction ID",
			Error:   err.Error(),
		})
	}

	var req interfaces.UpdateTransactionRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(responses.Response{
			Status:  "error",
			Message: "Invalid request body",
			Error:   err.Error(),
		})
	}

	transaction, err := h.usecase.Transaction.UpdateTransaction(c.Context(), uint(id), req)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(responses.Response{
			Status:  "error",
			Message: "Failed to update transaction",
			Error:   err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(responses.Response{
		Status:  "success",
		Message: "Transaction updated successfully",
		Data:    transaction,
	})
}

// DeleteTransaction deletes a transaction
func (h *FinancialHandler) DeleteTransaction(c *fiber.Ctx) error {
	id, err := strconv.ParseUint(c.Params("id"), 10, 32)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(responses.Response{
			Status:  "error",
			Message: "Invalid transaction ID",
			Error:   err.Error(),
		})
	}

	err = h.usecase.Transaction.DeleteTransaction(c.Context(), uint(id))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(responses.Response{
			Status:  "error",
			Message: "Failed to delete transaction",
			Error:   err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(responses.Response{
		Status:  "success",
		Message: "Transaction deleted successfully",
		Data:    nil,
	})
}

// GetTransactionsByCustomer retrieves transactions by customer ID
func (h *FinancialHandler) GetTransactionsByCustomer(c *fiber.Ctx) error {
	customerID, err := strconv.ParseUint(c.Params("customer_id"), 10, 32)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(responses.Response{
			Status:  "error",
			Message: "Invalid customer ID",
			Error:   err.Error(),
		})
	}

	transactions, err := h.usecase.Transaction.GetTransactionsByCustomer(c.Context(), uint(customerID))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(responses.Response{
			Status:  "error",
			Message: "Failed to retrieve transactions",
			Error:   err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(responses.Response{
		Status:  "success",
		Message: "Transactions retrieved successfully",
		Data:    transactions,
	})
}

// GetTransactionsByOutlet retrieves transactions by outlet ID
func (h *FinancialHandler) GetTransactionsByOutlet(c *fiber.Ctx) error {
	outletID, err := strconv.ParseUint(c.Params("outlet_id"), 10, 32)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(responses.Response{
			Status:  "error",
			Message: "Invalid outlet ID",
			Error:   err.Error(),
		})
	}

	transactions, err := h.usecase.Transaction.GetTransactionsByOutlet(c.Context(), uint(outletID))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(responses.Response{
			Status:  "error",
			Message: "Failed to retrieve transactions",
			Error:   err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(responses.Response{
		Status:  "success",
		Message: "Transactions retrieved successfully",
		Data:    transactions,
	})
}

// GetTransactionsByStatus retrieves transactions by status
func (h *FinancialHandler) GetTransactionsByStatus(c *fiber.Ctx) error {
	status := c.Query("status")
	if status == "" {
		return c.Status(fiber.StatusBadRequest).JSON(responses.Response{
			Status:  "error",
			Message: "Status is required",
			Error:   "missing status query parameter",
		})
	}

	transactions, err := h.usecase.Transaction.GetTransactionsByStatus(c.Context(), models.TransactionStatus(status))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(responses.Response{
			Status:  "error",
			Message: "Failed to retrieve transactions",
			Error:   err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(responses.Response{
		Status:  "success",
		Message: "Transactions retrieved successfully",
		Data:    transactions,
	})
}

// GetTransactionsByDateRange retrieves transactions by date range
func (h *FinancialHandler) GetTransactionsByDateRange(c *fiber.Ctx) error {
	startDateStr := c.Query("start_date")
	endDateStr := c.Query("end_date")

	if startDateStr == "" || endDateStr == "" {
		return c.Status(fiber.StatusBadRequest).JSON(responses.Response{
			Status:  "error",
			Message: "Start date and end date are required",
			Error:   "missing start_date or end_date query parameters",
		})
	}

	startDate, err := time.Parse("2006-01-02", startDateStr)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(responses.Response{
			Status:  "error",
			Message: "Invalid start date format",
			Error:   err.Error(),
		})
	}

	endDate, err := time.Parse("2006-01-02", endDateStr)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(responses.Response{
			Status:  "error",
			Message: "Invalid end date format",
			Error:   err.Error(),
		})
	}

	transactions, err := h.usecase.Transaction.GetTransactionsByDateRange(c.Context(), startDate, endDate)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(responses.Response{
			Status:  "error",
			Message: "Failed to retrieve transactions",
			Error:   err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(responses.Response{
		Status:  "success",
		Message: "Transactions retrieved successfully",
		Data:    transactions,
	})
}

// ============= Cash Flow Handlers =============

// CreateCashFlow creates a new cash flow
func (h *FinancialHandler) CreateCashFlow(c *fiber.Ctx) error {
	var req interfaces.CreateCashFlowRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(responses.Response{
			Status:  "error",
			Message: "Invalid request body",
			Error:   err.Error(),
		})
	}

	cashFlow, err := h.usecase.CashFlow.CreateCashFlow(c.Context(), req)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(responses.Response{
			Status:  "error",
			Message: "Failed to create cash flow",
			Error:   err.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(responses.Response{
		Status:  "success",
		Message: "Cash flow created successfully",
		Data:    cashFlow,
	})
}

// ListCashFlows lists all cash flows with pagination
func (h *FinancialHandler) ListCashFlows(c *fiber.Ctx) error {
	limit, _ := strconv.Atoi(c.Query("limit", "10"))
	offset, _ := strconv.Atoi(c.Query("offset", "0"))

	cashFlows, err := h.usecase.CashFlow.ListCashFlows(c.Context(), limit, offset)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(responses.Response{
			Status:  "error",
			Message: "Failed to retrieve cash flows",
			Error:   err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(responses.Response{
		Status:  "success",
		Message: "Cash flows retrieved successfully",
		Data:    cashFlows,
	})
}

// GetCashFlow retrieves a cash flow by ID
func (h *FinancialHandler) GetCashFlow(c *fiber.Ctx) error {
	id, err := strconv.ParseUint(c.Params("id"), 10, 32)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(responses.Response{
			Status:  "error",
			Message: "Invalid cash flow ID",
			Error:   err.Error(),
		})
	}

	cashFlow, err := h.usecase.CashFlow.GetCashFlow(c.Context(), uint(id))
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(responses.Response{
			Status:  "error",
			Message: "Cash flow not found",
			Error:   err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(responses.Response{
		Status:  "success",
		Message: "Cash flow retrieved successfully",
		Data:    cashFlow,
	})
}

// UpdateCashFlow updates a cash flow
func (h *FinancialHandler) UpdateCashFlow(c *fiber.Ctx) error {
	id, err := strconv.ParseUint(c.Params("id"), 10, 32)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(responses.Response{
			Status:  "error",
			Message: "Invalid cash flow ID",
			Error:   err.Error(),
		})
	}

	var req interfaces.UpdateCashFlowRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(responses.Response{
			Status:  "error",
			Message: "Invalid request body",
			Error:   err.Error(),
		})
	}

	cashFlow, err := h.usecase.CashFlow.UpdateCashFlow(c.Context(), uint(id), req)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(responses.Response{
			Status:  "error",
			Message: "Failed to update cash flow",
			Error:   err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(responses.Response{
		Status:  "success",
		Message: "Cash flow updated successfully",
		Data:    cashFlow,
	})
}

// DeleteCashFlow deletes a cash flow
func (h *FinancialHandler) DeleteCashFlow(c *fiber.Ctx) error {
	id, err := strconv.ParseUint(c.Params("id"), 10, 32)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(responses.Response{
			Status:  "error",
			Message: "Invalid cash flow ID",
			Error:   err.Error(),
		})
	}

	err = h.usecase.CashFlow.DeleteCashFlow(c.Context(), uint(id))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(responses.Response{
			Status:  "error",
			Message: "Failed to delete cash flow",
			Error:   err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(responses.Response{
		Status:  "success",
		Message: "Cash flow deleted successfully",
		Data:    nil,
	})
}

// GetCashFlowsByType retrieves cash flows by type
func (h *FinancialHandler) GetCashFlowsByType(c *fiber.Ctx) error {
	flowType := c.Query("type")
	if flowType == "" {
		return c.Status(fiber.StatusBadRequest).JSON(responses.Response{
			Status:  "error",
			Message: "Flow type is required",
			Error:   "missing type query parameter",
		})
	}

	cashFlows, err := h.usecase.CashFlow.GetCashFlowsByType(c.Context(), models.CashFlowType(flowType))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(responses.Response{
			Status:  "error",
			Message: "Failed to retrieve cash flows",
			Error:   err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(responses.Response{
		Status:  "success",
		Message: "Cash flows retrieved successfully",
		Data:    cashFlows,
	})
}