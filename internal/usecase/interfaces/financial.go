package interfaces

import (
	"boilerplate/internal/models"
	"context"
	"time"
)

// Financial module request structures

// PaymentMethod request structures
type CreatePaymentMethodRequest struct {
	Name      string            `json:"name" validate:"required,min=2,max=100"`
	Status    models.StatusUmum `json:"status,omitempty"`
	CreatedBy *uint             `json:"created_by,omitempty"`
}

type UpdatePaymentMethodRequest struct {
	Name   *string            `json:"name,omitempty" validate:"omitempty,min=2,max=100"`
	Status *models.StatusUmum `json:"status,omitempty"`
}

// Payment request structures
type CreatePaymentRequest struct {
	TransactionID uint                      `json:"transaction_id" validate:"required"`
	MethodID      uint                      `json:"method_id" validate:"required"`
	Amount        float64                   `json:"amount" validate:"required,min=0"`
	Status        models.TransactionStatus  `json:"status,omitempty"`
	PaymentDate   *time.Time                `json:"payment_date,omitempty"`
	CreatedBy     *uint                     `json:"created_by,omitempty"`
}

type UpdatePaymentRequest struct {
	TransactionID *uint                      `json:"transaction_id,omitempty"`
	MethodID      *uint                      `json:"method_id,omitempty"`
	Amount        *float64                   `json:"amount,omitempty" validate:"omitempty,min=0"`
	Status        *models.TransactionStatus  `json:"status,omitempty"`
	PaymentDate   *time.Time                 `json:"payment_date,omitempty"`
}

// CashFlow request structures
type CreateCashFlowRequest struct {
	UserID      uint                `json:"user_id" validate:"required"`
	OutletID    uint                `json:"outlet_id" validate:"required"`
	FlowType    models.CashFlowType `json:"flow_type" validate:"required"`
	Amount      float64             `json:"amount" validate:"required,min=0"`
	Description string              `json:"description" validate:"required,min=2,max=255"`
	FlowDate    time.Time           `json:"flow_date" validate:"required"`
	CreatedBy   *uint               `json:"created_by,omitempty"`
}

type UpdateCashFlowRequest struct {
	UserID      *uint                `json:"user_id,omitempty"`
	OutletID    *uint                `json:"outlet_id,omitempty"`
	FlowType    *models.CashFlowType `json:"flow_type,omitempty"`
	Amount      *float64             `json:"amount,omitempty" validate:"omitempty,min=0"`
	Description *string              `json:"description,omitempty" validate:"omitempty,min=2,max=255"`
	FlowDate    *time.Time           `json:"flow_date,omitempty"`
}

// Transaction request structures
type CreateTransactionRequest struct {
	InvoiceNumber   string                   `json:"invoice_number" validate:"required,min=3,max=255"`
	TransactionDate time.Time                `json:"transaction_date" validate:"required"`
	UserID          uint                     `json:"user_id" validate:"required"`
	CustomerID      *uint                    `json:"customer_id,omitempty"`
	OutletID        uint                     `json:"outlet_id" validate:"required"`
	TransactionType string                   `json:"transaction_type" validate:"required"`
	Status          models.TransactionStatus `json:"status,omitempty"`
	CreatedBy       *uint                    `json:"created_by,omitempty"`
}

type UpdateTransactionRequest struct {
	InvoiceNumber   *string                   `json:"invoice_number,omitempty" validate:"omitempty,min=3,max=255"`
	TransactionDate *time.Time                `json:"transaction_date,omitempty"`
	UserID          *uint                     `json:"user_id,omitempty"`
	CustomerID      *uint                     `json:"customer_id,omitempty"`
	OutletID        *uint                     `json:"outlet_id,omitempty"`
	TransactionType *string                   `json:"transaction_type,omitempty"`
	Status          *models.TransactionStatus `json:"status,omitempty"`
}

// Transaction Detail request structures
type CreateTransactionDetailRequest struct {
	TransactionType string  `json:"transaction_type" validate:"required"`
	TransactionID   uint    `json:"transaction_id" validate:"required"`
	ProductID       *uint   `json:"product_id,omitempty"`
	SerialNumberID  *uint   `json:"serial_number_id,omitempty"`
	Quantity        int     `json:"quantity" validate:"required,min=1"`
	UnitPrice       float64 `json:"unit_price" validate:"required,min=0"`
	TotalPrice      float64 `json:"total_price" validate:"required,min=0"`
	CreatedBy       *uint   `json:"created_by,omitempty"`
}

type UpdateTransactionDetailRequest struct {
	TransactionType *string  `json:"transaction_type,omitempty"`
	TransactionID   *uint    `json:"transaction_id,omitempty"`
	ProductID       *uint    `json:"product_id,omitempty"`
	SerialNumberID  *uint    `json:"serial_number_id,omitempty"`
	Quantity        *int     `json:"quantity,omitempty" validate:"omitempty,min=1"`
	UnitPrice       *float64 `json:"unit_price,omitempty" validate:"omitempty,min=0"`
	TotalPrice      *float64 `json:"total_price,omitempty" validate:"omitempty,min=0"`
}

// Usecase interfaces
type PaymentMethodUsecase interface {
	CreatePaymentMethod(ctx context.Context, req CreatePaymentMethodRequest) (*models.PaymentMethod, error)
	GetPaymentMethod(ctx context.Context, id uint) (*models.PaymentMethod, error)
	GetPaymentMethodByName(ctx context.Context, name string) (*models.PaymentMethod, error)
	UpdatePaymentMethod(ctx context.Context, id uint, req UpdatePaymentMethodRequest) (*models.PaymentMethod, error)
	DeletePaymentMethod(ctx context.Context, id uint) error
	ListPaymentMethods(ctx context.Context, limit, offset int) ([]*models.PaymentMethod, error)
	GetPaymentMethodsByStatus(ctx context.Context, status models.StatusUmum) ([]*models.PaymentMethod, error)
}

type PaymentUsecase interface {
	CreatePayment(ctx context.Context, req CreatePaymentRequest) (*models.Payment, error)
	GetPayment(ctx context.Context, id uint) (*models.Payment, error)
	UpdatePayment(ctx context.Context, id uint, req UpdatePaymentRequest) (*models.Payment, error)
	DeletePayment(ctx context.Context, id uint) error
	ListPayments(ctx context.Context, limit, offset int) ([]*models.Payment, error)
	GetPaymentsByTransaction(ctx context.Context, transactionID uint) ([]*models.Payment, error)
	GetPaymentsByMethod(ctx context.Context, methodID uint) ([]*models.Payment, error)
	GetPaymentsByStatus(ctx context.Context, status models.TransactionStatus) ([]*models.Payment, error)
	GetPaymentsByDateRange(ctx context.Context, startDate, endDate time.Time) ([]*models.Payment, error)
}

type CashFlowUsecase interface {
	CreateCashFlow(ctx context.Context, req CreateCashFlowRequest) (*models.CashFlow, error)
	GetCashFlow(ctx context.Context, id uint) (*models.CashFlow, error)
	UpdateCashFlow(ctx context.Context, id uint, req UpdateCashFlowRequest) (*models.CashFlow, error)
	DeleteCashFlow(ctx context.Context, id uint) error
	ListCashFlows(ctx context.Context, limit, offset int) ([]*models.CashFlow, error)
	GetCashFlowsByUser(ctx context.Context, userID uint) ([]*models.CashFlow, error)
	GetCashFlowsByOutlet(ctx context.Context, outletID uint) ([]*models.CashFlow, error)
	GetCashFlowsByType(ctx context.Context, flowType models.CashFlowType) ([]*models.CashFlow, error)
	GetCashFlowsByDateRange(ctx context.Context, startDate, endDate time.Time) ([]*models.CashFlow, error)
	GetTotalByTypeAndDateRange(ctx context.Context, flowType models.CashFlowType, startDate, endDate time.Time) (float64, error)
}

type TransactionUsecase interface {
	CreateTransaction(ctx context.Context, req CreateTransactionRequest) (*models.Transaction, error)
	GetTransaction(ctx context.Context, id uint) (*models.Transaction, error)
	GetTransactionByInvoiceNumber(ctx context.Context, invoiceNumber string) (*models.Transaction, error)
	UpdateTransaction(ctx context.Context, id uint, req UpdateTransactionRequest) (*models.Transaction, error)
	DeleteTransaction(ctx context.Context, id uint) error
	ListTransactions(ctx context.Context, limit, offset int) ([]*models.Transaction, error)
	GetTransactionsByCustomer(ctx context.Context, customerID uint) ([]*models.Transaction, error)
	GetTransactionsByUser(ctx context.Context, userID uint) ([]*models.Transaction, error)
	GetTransactionsByOutlet(ctx context.Context, outletID uint) ([]*models.Transaction, error)
	GetTransactionsByStatus(ctx context.Context, status models.TransactionStatus) ([]*models.Transaction, error)
	GetTransactionsByDateRange(ctx context.Context, startDate, endDate time.Time) ([]*models.Transaction, error)
}

type TransactionDetailUsecase interface {
	CreateTransactionDetail(ctx context.Context, req CreateTransactionDetailRequest) (*models.TransactionDetail, error)
	GetTransactionDetail(ctx context.Context, id uint) (*models.TransactionDetail, error)
	UpdateTransactionDetail(ctx context.Context, id uint, req UpdateTransactionDetailRequest) (*models.TransactionDetail, error)
	DeleteTransactionDetail(ctx context.Context, id uint) error
	ListTransactionDetails(ctx context.Context, limit, offset int) ([]*models.TransactionDetail, error)
	GetTransactionDetailsByTransaction(ctx context.Context, transactionID uint) ([]*models.TransactionDetail, error)
	GetTransactionDetailsByProduct(ctx context.Context, productID uint) ([]*models.TransactionDetail, error)
	DeleteTransactionDetailsByTransaction(ctx context.Context, transactionID uint) error
}