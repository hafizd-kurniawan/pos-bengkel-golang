package implementations

import (
	"boilerplate/internal/models"
	"boilerplate/internal/repository/interfaces"
	"context"
	"time"

	"gorm.io/gorm"
)

// TransactionRepository implements the transaction repository interface
type TransactionRepository struct {
	db *gorm.DB
}

// NewTransactionRepository creates a new transaction repository
func NewTransactionRepository(db *gorm.DB) interfaces.TransactionRepository {
	return &TransactionRepository{db: db}
}

// Create creates a new transaction
func (r *TransactionRepository) Create(ctx context.Context, transaction *models.Transaction) error {
	return r.db.WithContext(ctx).Create(transaction).Error
}

// GetByID retrieves a transaction by ID
func (r *TransactionRepository) GetByID(ctx context.Context, id uint) (*models.Transaction, error) {
	var transaction models.Transaction
	err := r.db.WithContext(ctx).
		Preload("User").
		Preload("Customer").
		Preload("Outlet").
		Preload("TransactionDetails").
		Preload("Payments").
		First(&transaction, id).Error
	if err != nil {
		return nil, err
	}
	return &transaction, nil
}

// GetByInvoiceNumber retrieves a transaction by invoice number
func (r *TransactionRepository) GetByInvoiceNumber(ctx context.Context, invoiceNumber string) (*models.Transaction, error) {
	var transaction models.Transaction
	err := r.db.WithContext(ctx).
		Preload("User").
		Preload("Customer").
		Preload("Outlet").
		Preload("TransactionDetails").
		Preload("Payments").
		Where("invoice_number = ?", invoiceNumber).
		First(&transaction).Error
	if err != nil {
		return nil, err
	}
	return &transaction, nil
}

// Update updates a transaction
func (r *TransactionRepository) Update(ctx context.Context, transaction *models.Transaction) error {
	return r.db.WithContext(ctx).Save(transaction).Error
}

// Delete soft deletes a transaction
func (r *TransactionRepository) Delete(ctx context.Context, id uint) error {
	return r.db.WithContext(ctx).Delete(&models.Transaction{}, id).Error
}

// List retrieves transactions with pagination
func (r *TransactionRepository) List(ctx context.Context, limit, offset int) ([]*models.Transaction, error) {
	var transactions []*models.Transaction
	err := r.db.WithContext(ctx).
		Preload("User").
		Preload("Customer").
		Preload("Outlet").
		Preload("TransactionDetails").
		Preload("Payments").
		Limit(limit).
		Offset(offset).
		Find(&transactions).Error
	if err != nil {
		return nil, err
	}
	return transactions, nil
}

// GetByCustomerID retrieves transactions by customer ID
func (r *TransactionRepository) GetByCustomerID(ctx context.Context, customerID uint) ([]*models.Transaction, error) {
	var transactions []*models.Transaction
	err := r.db.WithContext(ctx).
		Preload("User").
		Preload("Customer").
		Preload("Outlet").
		Preload("TransactionDetails").
		Preload("Payments").
		Where("customer_id = ?", customerID).
		Find(&transactions).Error
	if err != nil {
		return nil, err
	}
	return transactions, nil
}

// GetByUserID retrieves transactions by user ID
func (r *TransactionRepository) GetByUserID(ctx context.Context, userID uint) ([]*models.Transaction, error) {
	var transactions []*models.Transaction
	err := r.db.WithContext(ctx).
		Preload("User").
		Preload("Customer").
		Preload("Outlet").
		Preload("TransactionDetails").
		Preload("Payments").
		Where("user_id = ?", userID).
		Find(&transactions).Error
	if err != nil {
		return nil, err
	}
	return transactions, nil
}

// GetByOutletID retrieves transactions by outlet ID
func (r *TransactionRepository) GetByOutletID(ctx context.Context, outletID uint) ([]*models.Transaction, error) {
	var transactions []*models.Transaction
	err := r.db.WithContext(ctx).
		Preload("User").
		Preload("Customer").
		Preload("Outlet").
		Preload("TransactionDetails").
		Preload("Payments").
		Where("outlet_id = ?", outletID).
		Find(&transactions).Error
	if err != nil {
		return nil, err
	}
	return transactions, nil
}

// GetByStatus retrieves transactions by status
func (r *TransactionRepository) GetByStatus(ctx context.Context, status models.TransactionStatus) ([]*models.Transaction, error) {
	var transactions []*models.Transaction
	err := r.db.WithContext(ctx).
		Preload("User").
		Preload("Customer").
		Preload("Outlet").
		Preload("TransactionDetails").
		Preload("Payments").
		Where("status = ?", status).
		Find(&transactions).Error
	if err != nil {
		return nil, err
	}
	return transactions, nil
}

// GetByDateRange retrieves transactions by date range
func (r *TransactionRepository) GetByDateRange(ctx context.Context, startDate, endDate time.Time) ([]*models.Transaction, error) {
	var transactions []*models.Transaction
	err := r.db.WithContext(ctx).
		Preload("User").
		Preload("Customer").
		Preload("Outlet").
		Preload("TransactionDetails").
		Preload("Payments").
		Where("transaction_date BETWEEN ? AND ?", startDate, endDate).
		Find(&transactions).Error
	if err != nil {
		return nil, err
	}
	return transactions, nil
}

// TransactionDetailRepository implements the transaction detail repository interface
type TransactionDetailRepository struct {
	db *gorm.DB
}

// NewTransactionDetailRepository creates a new transaction detail repository
func NewTransactionDetailRepository(db *gorm.DB) interfaces.TransactionDetailRepository {
	return &TransactionDetailRepository{db: db}
}

// Create creates a new transaction detail
func (r *TransactionDetailRepository) Create(ctx context.Context, detail *models.TransactionDetail) error {
	return r.db.WithContext(ctx).Create(detail).Error
}

// GetByID retrieves a transaction detail by ID
func (r *TransactionDetailRepository) GetByID(ctx context.Context, id uint) (*models.TransactionDetail, error) {
	var detail models.TransactionDetail
	err := r.db.WithContext(ctx).
		Preload("Transaction").
		Preload("Product").
		Preload("SerialNumber").
		First(&detail, id).Error
	if err != nil {
		return nil, err
	}
	return &detail, nil
}

// Update updates a transaction detail
func (r *TransactionDetailRepository) Update(ctx context.Context, detail *models.TransactionDetail) error {
	return r.db.WithContext(ctx).Save(detail).Error
}

// Delete soft deletes a transaction detail
func (r *TransactionDetailRepository) Delete(ctx context.Context, id uint) error {
	return r.db.WithContext(ctx).Delete(&models.TransactionDetail{}, id).Error
}

// List retrieves transaction details with pagination
func (r *TransactionDetailRepository) List(ctx context.Context, limit, offset int) ([]*models.TransactionDetail, error) {
	var details []*models.TransactionDetail
	err := r.db.WithContext(ctx).
		Preload("Transaction").
		Preload("Product").
		Preload("SerialNumber").
		Limit(limit).
		Offset(offset).
		Find(&details).Error
	if err != nil {
		return nil, err
	}
	return details, nil
}

// GetByTransactionID retrieves transaction details by transaction ID
func (r *TransactionDetailRepository) GetByTransactionID(ctx context.Context, transactionID uint) ([]*models.TransactionDetail, error) {
	var details []*models.TransactionDetail
	err := r.db.WithContext(ctx).
		Preload("Transaction").
		Preload("Product").
		Preload("SerialNumber").
		Where("transaction_id = ?", transactionID).
		Find(&details).Error
	if err != nil {
		return nil, err
	}
	return details, nil
}

// GetByProductID retrieves transaction details by product ID
func (r *TransactionDetailRepository) GetByProductID(ctx context.Context, productID uint) ([]*models.TransactionDetail, error) {
	var details []*models.TransactionDetail
	err := r.db.WithContext(ctx).
		Preload("Transaction").
		Preload("Product").
		Preload("SerialNumber").
		Where("product_id = ?", productID).
		Find(&details).Error
	if err != nil {
		return nil, err
	}
	return details, nil
}

// DeleteByTransactionID deletes transaction details by transaction ID
func (r *TransactionDetailRepository) DeleteByTransactionID(ctx context.Context, transactionID uint) error {
	return r.db.WithContext(ctx).
		Where("transaction_id = ?", transactionID).
		Delete(&models.TransactionDetail{}).Error
}