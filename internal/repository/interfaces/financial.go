package interfaces

import (
	"boilerplate/internal/models"
	"context"
	"time"
)

// PaymentMethodRepository interface for payment method operations
type PaymentMethodRepository interface {
	Create(ctx context.Context, method *models.PaymentMethod) error
	GetByID(ctx context.Context, id uint) (*models.PaymentMethod, error)
	GetByName(ctx context.Context, name string) (*models.PaymentMethod, error)
	Update(ctx context.Context, method *models.PaymentMethod) error
	Delete(ctx context.Context, id uint) error
	List(ctx context.Context, limit, offset int) ([]*models.PaymentMethod, error)
	GetByStatus(ctx context.Context, status models.StatusUmum) ([]*models.PaymentMethod, error)
}

// PaymentRepository interface for payment operations
type PaymentRepository interface {
	Create(ctx context.Context, payment *models.Payment) error
	GetByID(ctx context.Context, id uint) (*models.Payment, error)
	Update(ctx context.Context, payment *models.Payment) error
	Delete(ctx context.Context, id uint) error
	List(ctx context.Context, limit, offset int) ([]*models.Payment, error)
	GetByTransactionID(ctx context.Context, transactionID uint) ([]*models.Payment, error)
	GetByMethodID(ctx context.Context, methodID uint) ([]*models.Payment, error)
	GetByStatus(ctx context.Context, status models.TransactionStatus) ([]*models.Payment, error)
	GetByDateRange(ctx context.Context, startDate, endDate time.Time) ([]*models.Payment, error)
}

// AccountsPayableRepository interface for accounts payable operations
type AccountsPayableRepository interface {
	Create(ctx context.Context, payable *models.AccountsPayable) error
	GetByID(ctx context.Context, id uint) (*models.AccountsPayable, error)
	Update(ctx context.Context, payable *models.AccountsPayable) error
	Delete(ctx context.Context, id uint) error
	List(ctx context.Context, limit, offset int) ([]*models.AccountsPayable, error)
	GetByPurchaseOrderID(ctx context.Context, purchaseOrderID uint) ([]*models.AccountsPayable, error)
	GetBySupplierID(ctx context.Context, supplierID uint) ([]*models.AccountsPayable, error)
	GetByStatus(ctx context.Context, status models.APARStatus) ([]*models.AccountsPayable, error)
	GetOverdue(ctx context.Context) ([]*models.AccountsPayable, error)
	UpdateAmountPaid(ctx context.Context, id uint, amount float64) error
}

// PayablePaymentRepository interface for payable payment operations
type PayablePaymentRepository interface {
	Create(ctx context.Context, payment *models.PayablePayment) error
	GetByID(ctx context.Context, id uint) (*models.PayablePayment, error)
	Update(ctx context.Context, payment *models.PayablePayment) error
	Delete(ctx context.Context, id uint) error
	List(ctx context.Context, limit, offset int) ([]*models.PayablePayment, error)
	GetByPayableID(ctx context.Context, payableID uint) ([]*models.PayablePayment, error)
	GetByDateRange(ctx context.Context, startDate, endDate time.Time) ([]*models.PayablePayment, error)
}

// AccountsReceivableRepository interface for accounts receivable operations
type AccountsReceivableRepository interface {
	Create(ctx context.Context, receivable *models.AccountsReceivable) error
	GetByID(ctx context.Context, id uint) (*models.AccountsReceivable, error)
	Update(ctx context.Context, receivable *models.AccountsReceivable) error
	Delete(ctx context.Context, id uint) error
	List(ctx context.Context, limit, offset int) ([]*models.AccountsReceivable, error)
	GetByTransactionID(ctx context.Context, transactionID uint) ([]*models.AccountsReceivable, error)
	GetByCustomerID(ctx context.Context, customerID uint) ([]*models.AccountsReceivable, error)
	GetByStatus(ctx context.Context, status models.APARStatus) ([]*models.AccountsReceivable, error)
	GetOverdue(ctx context.Context) ([]*models.AccountsReceivable, error)
	UpdateAmountPaid(ctx context.Context, id uint, amount float64) error
}

// ReceivablePaymentRepository interface for receivable payment operations
type ReceivablePaymentRepository interface {
	Create(ctx context.Context, payment *models.ReceivablePayment) error
	GetByID(ctx context.Context, id uint) (*models.ReceivablePayment, error)
	Update(ctx context.Context, payment *models.ReceivablePayment) error
	Delete(ctx context.Context, id uint) error
	List(ctx context.Context, limit, offset int) ([]*models.ReceivablePayment, error)
	GetByReceivableID(ctx context.Context, receivableID uint) ([]*models.ReceivablePayment, error)
	GetByDateRange(ctx context.Context, startDate, endDate time.Time) ([]*models.ReceivablePayment, error)
}

// CashFlowRepository interface for cash flow operations
type CashFlowRepository interface {
	Create(ctx context.Context, cashFlow *models.CashFlow) error
	GetByID(ctx context.Context, id uint) (*models.CashFlow, error)
	Update(ctx context.Context, cashFlow *models.CashFlow) error
	Delete(ctx context.Context, id uint) error
	List(ctx context.Context, limit, offset int) ([]*models.CashFlow, error)
	GetByUserID(ctx context.Context, userID uint) ([]*models.CashFlow, error)
	GetByType(ctx context.Context, cashFlowType models.CashFlowType) ([]*models.CashFlow, error)
	GetByDateRange(ctx context.Context, startDate, endDate time.Time) ([]*models.CashFlow, error)
	GetTotalByType(ctx context.Context, cashFlowType models.CashFlowType, startDate, endDate time.Time) (float64, error)
}