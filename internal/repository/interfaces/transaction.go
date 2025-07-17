package interfaces

import (
	"boilerplate/internal/models"
	"context"
	"time"
)

// TransactionRepository interface for transaction operations
type TransactionRepository interface {
	Create(ctx context.Context, transaction *models.Transaction) error
	GetByID(ctx context.Context, id uint) (*models.Transaction, error)
	GetByInvoiceNumber(ctx context.Context, invoiceNumber string) (*models.Transaction, error)
	Update(ctx context.Context, transaction *models.Transaction) error
	Delete(ctx context.Context, id uint) error
	List(ctx context.Context, limit, offset int) ([]*models.Transaction, error)
	GetByCustomerID(ctx context.Context, customerID uint) ([]*models.Transaction, error)
	GetByUserID(ctx context.Context, userID uint) ([]*models.Transaction, error)
	GetByOutletID(ctx context.Context, outletID uint) ([]*models.Transaction, error)
	GetByStatus(ctx context.Context, status models.TransactionStatus) ([]*models.Transaction, error)
	GetByDateRange(ctx context.Context, startDate, endDate time.Time) ([]*models.Transaction, error)
}

// TransactionDetailRepository interface for transaction detail operations
type TransactionDetailRepository interface {
	Create(ctx context.Context, detail *models.TransactionDetail) error
	GetByID(ctx context.Context, id uint) (*models.TransactionDetail, error)
	Update(ctx context.Context, detail *models.TransactionDetail) error
	Delete(ctx context.Context, id uint) error
	List(ctx context.Context, limit, offset int) ([]*models.TransactionDetail, error)
	GetByTransactionID(ctx context.Context, transactionID uint) ([]*models.TransactionDetail, error)
	GetByProductID(ctx context.Context, productID uint) ([]*models.TransactionDetail, error)
	DeleteByTransactionID(ctx context.Context, transactionID uint) error
}

// PurchaseOrderRepository interface for purchase order operations
type PurchaseOrderRepository interface {
	Create(ctx context.Context, purchaseOrder *models.PurchaseOrder) error
	GetByID(ctx context.Context, id uint) (*models.PurchaseOrder, error)
	GetByPOCode(ctx context.Context, poCode string) (*models.PurchaseOrder, error)
	Update(ctx context.Context, purchaseOrder *models.PurchaseOrder) error
	Delete(ctx context.Context, id uint) error
	List(ctx context.Context, limit, offset int) ([]*models.PurchaseOrder, error)
	GetBySupplierID(ctx context.Context, supplierID uint) ([]*models.PurchaseOrder, error)
	GetByOutletID(ctx context.Context, outletID uint) ([]*models.PurchaseOrder, error)
	GetByStatus(ctx context.Context, status models.PurchaseStatus) ([]*models.PurchaseOrder, error)
	GetByDateRange(ctx context.Context, startDate, endDate time.Time) ([]*models.PurchaseOrder, error)
}

// PurchaseOrderDetailRepository interface for purchase order detail operations
type PurchaseOrderDetailRepository interface {
	Create(ctx context.Context, detail *models.PurchaseOrderDetail) error
	GetByID(ctx context.Context, id uint) (*models.PurchaseOrderDetail, error)
	Update(ctx context.Context, detail *models.PurchaseOrderDetail) error
	Delete(ctx context.Context, id uint) error
	List(ctx context.Context, limit, offset int) ([]*models.PurchaseOrderDetail, error)
	GetByPurchaseOrderID(ctx context.Context, purchaseOrderID uint) ([]*models.PurchaseOrderDetail, error)
	GetByProductID(ctx context.Context, productID uint) ([]*models.PurchaseOrderDetail, error)
	DeleteByPurchaseOrderID(ctx context.Context, purchaseOrderID uint) error
}

// VehiclePurchaseRepository interface for vehicle purchase operations
type VehiclePurchaseRepository interface {
	Create(ctx context.Context, purchase *models.VehiclePurchase) error
	GetByID(ctx context.Context, id uint) (*models.VehiclePurchase, error)
	GetByPurchaseCode(ctx context.Context, purchaseCode string) (*models.VehiclePurchase, error)
	Update(ctx context.Context, purchase *models.VehiclePurchase) error
	Delete(ctx context.Context, id uint) error
	List(ctx context.Context, limit, offset int) ([]*models.VehiclePurchase, error)
	GetByCustomerID(ctx context.Context, customerID uint) ([]*models.VehiclePurchase, error)
	GetByUserID(ctx context.Context, userID uint) ([]*models.VehiclePurchase, error)
	GetByOutletID(ctx context.Context, outletID uint) ([]*models.VehiclePurchase, error)
	GetByDateRange(ctx context.Context, startDate, endDate time.Time) ([]*models.VehiclePurchase, error)
}