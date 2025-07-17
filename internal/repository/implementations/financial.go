package implementations

import (
	"boilerplate/internal/models"
	"boilerplate/internal/repository/interfaces"
	"context"
	"time"

	"gorm.io/gorm"
)

// PaymentMethodRepository implements the payment method repository interface
type PaymentMethodRepository struct {
	db *gorm.DB
}

// NewPaymentMethodRepository creates a new payment method repository
func NewPaymentMethodRepository(db *gorm.DB) interfaces.PaymentMethodRepository {
	return &PaymentMethodRepository{db: db}
}

// Create creates a new payment method
func (r *PaymentMethodRepository) Create(ctx context.Context, paymentMethod *models.PaymentMethod) error {
	return r.db.WithContext(ctx).Create(paymentMethod).Error
}

// GetByID retrieves a payment method by ID
func (r *PaymentMethodRepository) GetByID(ctx context.Context, id uint) (*models.PaymentMethod, error) {
	var paymentMethod models.PaymentMethod
	err := r.db.WithContext(ctx).Preload("Payments").First(&paymentMethod, id).Error
	if err != nil {
		return nil, err
	}
	return &paymentMethod, nil
}

// GetByName retrieves a payment method by name
func (r *PaymentMethodRepository) GetByName(ctx context.Context, name string) (*models.PaymentMethod, error) {
	var paymentMethod models.PaymentMethod
	err := r.db.WithContext(ctx).Preload("Payments").Where("name = ?", name).First(&paymentMethod).Error
	if err != nil {
		return nil, err
	}
	return &paymentMethod, nil
}

// Update updates a payment method
func (r *PaymentMethodRepository) Update(ctx context.Context, paymentMethod *models.PaymentMethod) error {
	return r.db.WithContext(ctx).Save(paymentMethod).Error
}

// Delete soft deletes a payment method
func (r *PaymentMethodRepository) Delete(ctx context.Context, id uint) error {
	return r.db.WithContext(ctx).Delete(&models.PaymentMethod{}, id).Error
}

// List retrieves payment methods with pagination
func (r *PaymentMethodRepository) List(ctx context.Context, limit, offset int) ([]*models.PaymentMethod, error) {
	var paymentMethods []*models.PaymentMethod
	err := r.db.WithContext(ctx).Preload("Payments").Limit(limit).Offset(offset).Find(&paymentMethods).Error
	if err != nil {
		return nil, err
	}
	return paymentMethods, nil
}

// GetByStatus retrieves payment methods by status
func (r *PaymentMethodRepository) GetByStatus(ctx context.Context, status models.StatusUmum) ([]*models.PaymentMethod, error) {
	var paymentMethods []*models.PaymentMethod
	err := r.db.WithContext(ctx).Preload("Payments").Where("status = ?", status).Find(&paymentMethods).Error
	if err != nil {
		return nil, err
	}
	return paymentMethods, nil
}

// PaymentRepository implements the payment repository interface
type PaymentRepository struct {
	db *gorm.DB
}

// NewPaymentRepository creates a new payment repository
func NewPaymentRepository(db *gorm.DB) interfaces.PaymentRepository {
	return &PaymentRepository{db: db}
}

// Create creates a new payment
func (r *PaymentRepository) Create(ctx context.Context, payment *models.Payment) error {
	return r.db.WithContext(ctx).Create(payment).Error
}

// GetByID retrieves a payment by ID
func (r *PaymentRepository) GetByID(ctx context.Context, id uint) (*models.Payment, error) {
	var payment models.Payment
	err := r.db.WithContext(ctx).
		Preload("Transaction").
		Preload("PaymentMethod").
		First(&payment, id).Error
	if err != nil {
		return nil, err
	}
	return &payment, nil
}

// Update updates a payment
func (r *PaymentRepository) Update(ctx context.Context, payment *models.Payment) error {
	return r.db.WithContext(ctx).Save(payment).Error
}

// Delete soft deletes a payment
func (r *PaymentRepository) Delete(ctx context.Context, id uint) error {
	return r.db.WithContext(ctx).Delete(&models.Payment{}, id).Error
}

// List retrieves payments with pagination
func (r *PaymentRepository) List(ctx context.Context, limit, offset int) ([]*models.Payment, error) {
	var payments []*models.Payment
	err := r.db.WithContext(ctx).
		Preload("Transaction").
		Preload("PaymentMethod").
		Limit(limit).
		Offset(offset).
		Find(&payments).Error
	if err != nil {
		return nil, err
	}
	return payments, nil
}

// GetByTransactionID retrieves payments by transaction ID
func (r *PaymentRepository) GetByTransactionID(ctx context.Context, transactionID uint) ([]*models.Payment, error) {
	var payments []*models.Payment
	err := r.db.WithContext(ctx).
		Preload("Transaction").
		Preload("PaymentMethod").
		Where("transaction_id = ?", transactionID).
		Find(&payments).Error
	if err != nil {
		return nil, err
	}
	return payments, nil
}

// GetByMethodID retrieves payments by payment method ID
func (r *PaymentRepository) GetByMethodID(ctx context.Context, methodID uint) ([]*models.Payment, error) {
	var payments []*models.Payment
	err := r.db.WithContext(ctx).
		Preload("Transaction").
		Preload("PaymentMethod").
		Where("method_id = ?", methodID).
		Find(&payments).Error
	if err != nil {
		return nil, err
	}
	return payments, nil
}

// GetByStatus retrieves payments by status
func (r *PaymentRepository) GetByStatus(ctx context.Context, status models.TransactionStatus) ([]*models.Payment, error) {
	var payments []*models.Payment
	err := r.db.WithContext(ctx).
		Preload("Transaction").
		Preload("PaymentMethod").
		Where("status = ?", status).
		Find(&payments).Error
	if err != nil {
		return nil, err
	}
	return payments, nil
}

// GetByDateRange retrieves payments by date range
func (r *PaymentRepository) GetByDateRange(ctx context.Context, startDate, endDate time.Time) ([]*models.Payment, error) {
	var payments []*models.Payment
	err := r.db.WithContext(ctx).
		Preload("Transaction").
		Preload("PaymentMethod").
		Where("payment_date BETWEEN ? AND ?", startDate, endDate).
		Find(&payments).Error
	if err != nil {
		return nil, err
	}
	return payments, nil
}

// CashFlowRepository implements the cash flow repository interface
type CashFlowRepository struct {
	db *gorm.DB
}

// NewCashFlowRepository creates a new cash flow repository
func NewCashFlowRepository(db *gorm.DB) interfaces.CashFlowRepository {
	return &CashFlowRepository{db: db}
}

// Create creates a new cash flow
func (r *CashFlowRepository) Create(ctx context.Context, cashFlow *models.CashFlow) error {
	return r.db.WithContext(ctx).Create(cashFlow).Error
}

// GetByID retrieves a cash flow by ID
func (r *CashFlowRepository) GetByID(ctx context.Context, id uint) (*models.CashFlow, error) {
	var cashFlow models.CashFlow
	err := r.db.WithContext(ctx).Preload("User").First(&cashFlow, id).Error
	if err != nil {
		return nil, err
	}
	return &cashFlow, nil
}

// Update updates a cash flow
func (r *CashFlowRepository) Update(ctx context.Context, cashFlow *models.CashFlow) error {
	return r.db.WithContext(ctx).Save(cashFlow).Error
}

// Delete soft deletes a cash flow
func (r *CashFlowRepository) Delete(ctx context.Context, id uint) error {
	return r.db.WithContext(ctx).Delete(&models.CashFlow{}, id).Error
}

// List retrieves cash flows with pagination
func (r *CashFlowRepository) List(ctx context.Context, limit, offset int) ([]*models.CashFlow, error) {
	var cashFlows []*models.CashFlow
	err := r.db.WithContext(ctx).
		Preload("User").
		Limit(limit).
		Offset(offset).
		Find(&cashFlows).Error
	if err != nil {
		return nil, err
	}
	return cashFlows, nil
}

// GetByUserID retrieves cash flows by user ID
func (r *CashFlowRepository) GetByUserID(ctx context.Context, userID uint) ([]*models.CashFlow, error) {
	var cashFlows []*models.CashFlow
	err := r.db.WithContext(ctx).
		Preload("User").
		Where("user_id = ?", userID).
		Find(&cashFlows).Error
	if err != nil {
		return nil, err
	}
	return cashFlows, nil
}

// GetByOutletID retrieves cash flows by outlet ID
func (r *CashFlowRepository) GetByOutletID(ctx context.Context, outletID uint) ([]*models.CashFlow, error) {
	var cashFlows []*models.CashFlow
	err := r.db.WithContext(ctx).
		Preload("User").
		Find(&cashFlows).Error
	if err != nil {
		return nil, err
	}
	return cashFlows, nil
}

// GetByType retrieves cash flows by type
func (r *CashFlowRepository) GetByType(ctx context.Context, flowType models.CashFlowType) ([]*models.CashFlow, error) {
	var cashFlows []*models.CashFlow
	err := r.db.WithContext(ctx).
		Preload("User").
		Where("type = ?", flowType).
		Find(&cashFlows).Error
	if err != nil {
		return nil, err
	}
	return cashFlows, nil
}

// GetByDateRange retrieves cash flows by date range
func (r *CashFlowRepository) GetByDateRange(ctx context.Context, startDate, endDate time.Time) ([]*models.CashFlow, error) {
	var cashFlows []*models.CashFlow
	err := r.db.WithContext(ctx).
		Preload("User").
		Where("date BETWEEN ? AND ?", startDate, endDate).
		Find(&cashFlows).Error
	if err != nil {
		return nil, err
	}
	return cashFlows, nil
}

// GetTotalByType retrieves total cash flows by type and date range
func (r *CashFlowRepository) GetTotalByType(ctx context.Context, flowType models.CashFlowType, startDate, endDate time.Time) (float64, error) {
	var total float64
	err := r.db.WithContext(ctx).
		Model(&models.CashFlow{}).
		Where("type = ? AND date BETWEEN ? AND ?", flowType, startDate, endDate).
		Select("COALESCE(SUM(amount), 0)").
		Scan(&total).Error
	if err != nil {
		return 0, err
	}
	return total, nil
}