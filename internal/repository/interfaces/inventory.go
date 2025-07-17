package interfaces

import (
	"boilerplate/internal/models"
	"context"
)

// ProductRepository interface for product operations
type ProductRepository interface {
	Create(ctx context.Context, product *models.Product) error
	GetByID(ctx context.Context, id uint) (*models.Product, error)
	GetBySKU(ctx context.Context, sku string) (*models.Product, error)
	GetByBarcode(ctx context.Context, barcode string) (*models.Product, error)
	Update(ctx context.Context, product *models.Product) error
	Delete(ctx context.Context, id uint) error
	List(ctx context.Context, limit, offset int) ([]*models.Product, error)
	GetByCategoryID(ctx context.Context, categoryID uint) ([]*models.Product, error)
	GetBySupplierID(ctx context.Context, supplierID uint) ([]*models.Product, error)
	GetByUsageStatus(ctx context.Context, status models.ProductUsageStatus) ([]*models.Product, error)
	Search(ctx context.Context, query string, limit, offset int) ([]*models.Product, error)
	UpdateStock(ctx context.Context, productID uint, quantity int) error
	GetLowStock(ctx context.Context, threshold int) ([]*models.Product, error)
}

// ProductSerialNumberRepository interface for product serial number operations
type ProductSerialNumberRepository interface {
	Create(ctx context.Context, serialNumber *models.ProductSerialNumber) error
	GetByID(ctx context.Context, id uint) (*models.ProductSerialNumber, error)
	GetBySerialNumber(ctx context.Context, serialNumber string) (*models.ProductSerialNumber, error)
	Update(ctx context.Context, serialNumber *models.ProductSerialNumber) error
	Delete(ctx context.Context, id uint) error
	List(ctx context.Context, limit, offset int) ([]*models.ProductSerialNumber, error)
	GetByProductID(ctx context.Context, productID uint) ([]*models.ProductSerialNumber, error)
	GetByStatus(ctx context.Context, status models.SNStatus) ([]*models.ProductSerialNumber, error)
	UpdateStatus(ctx context.Context, id uint, status models.SNStatus) error
}

// CategoryRepository interface for category operations
type CategoryRepository interface {
	Create(ctx context.Context, category *models.Category) error
	GetByID(ctx context.Context, id uint) (*models.Category, error)
	GetByName(ctx context.Context, name string) (*models.Category, error)
	Update(ctx context.Context, category *models.Category) error
	Delete(ctx context.Context, id uint) error
	List(ctx context.Context, limit, offset int) ([]*models.Category, error)
	GetByStatus(ctx context.Context, status models.StatusUmum) ([]*models.Category, error)
}

// SupplierRepository interface for supplier operations
type SupplierRepository interface {
	Create(ctx context.Context, supplier *models.Supplier) error
	GetByID(ctx context.Context, id uint) (*models.Supplier, error)
	GetByName(ctx context.Context, name string) (*models.Supplier, error)
	Update(ctx context.Context, supplier *models.Supplier) error
	Delete(ctx context.Context, id uint) error
	List(ctx context.Context, limit, offset int) ([]*models.Supplier, error)
	GetByStatus(ctx context.Context, status models.StatusUmum) ([]*models.Supplier, error)
	Search(ctx context.Context, query string, limit, offset int) ([]*models.Supplier, error)
}

// UnitTypeRepository interface for unit type operations
type UnitTypeRepository interface {
	Create(ctx context.Context, unitType *models.UnitType) error
	GetByID(ctx context.Context, id uint) (*models.UnitType, error)
	GetByName(ctx context.Context, name string) (*models.UnitType, error)
	Update(ctx context.Context, unitType *models.UnitType) error
	Delete(ctx context.Context, id uint) error
	List(ctx context.Context, limit, offset int) ([]*models.UnitType, error)
	GetByStatus(ctx context.Context, status models.StatusUmum) ([]*models.UnitType, error)
}