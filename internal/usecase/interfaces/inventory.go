package interfaces

import (
	"boilerplate/internal/models"
	"context"
)

// Product request structures
type CreateProductRequest struct {
	ProductName        string                      `json:"product_name" validate:"required,min=2,max=255"`
	ProductDescription *string                     `json:"product_description,omitempty"`
	ProductImage       *string                     `json:"product_image,omitempty"`
	CostPrice          float64                     `json:"cost_price" validate:"required,min=0"`
	SellingPrice       float64                     `json:"selling_price" validate:"required,min=0"`
	Stock              int                         `json:"stock" validate:"required,min=0"`
	SKU                *string                     `json:"sku,omitempty"`
	Barcode            *string                     `json:"barcode,omitempty"`
	HasSerialNumber    bool                        `json:"has_serial_number"`
	ShelfLocation      *string                     `json:"shelf_location,omitempty"`
	UsageStatus        models.ProductUsageStatus   `json:"usage_status" validate:"required"`
	IsActive           bool                        `json:"is_active"`
	CategoryID         *uint                       `json:"category_id,omitempty"`
	SupplierID         *uint                       `json:"supplier_id,omitempty"`
	UnitTypeID         *uint                       `json:"unit_type_id,omitempty"`
	CreatedBy          *uint                       `json:"created_by,omitempty"`
}

type UpdateProductRequest struct {
	ProductName        *string                     `json:"product_name,omitempty" validate:"omitempty,min=2,max=255"`
	ProductDescription *string                     `json:"product_description,omitempty"`
	ProductImage       *string                     `json:"product_image,omitempty"`
	CostPrice          *float64                    `json:"cost_price,omitempty" validate:"omitempty,min=0"`
	SellingPrice       *float64                    `json:"selling_price,omitempty" validate:"omitempty,min=0"`
	Stock              *int                        `json:"stock,omitempty" validate:"omitempty,min=0"`
	SKU                *string                     `json:"sku,omitempty"`
	Barcode            *string                     `json:"barcode,omitempty"`
	HasSerialNumber    *bool                       `json:"has_serial_number,omitempty"`
	ShelfLocation      *string                     `json:"shelf_location,omitempty"`
	UsageStatus        *models.ProductUsageStatus  `json:"usage_status,omitempty"`
	IsActive           *bool                       `json:"is_active,omitempty"`
	CategoryID         *uint                       `json:"category_id,omitempty"`
	SupplierID         *uint                       `json:"supplier_id,omitempty"`
	UnitTypeID         *uint                       `json:"unit_type_id,omitempty"`
}

// Product Serial Number request structures
type CreateProductSerialNumberRequest struct {
	ProductID    uint              `json:"product_id" validate:"required"`
	SerialNumber string            `json:"serial_number" validate:"required,min=3,max=255"`
	Status       models.SNStatus   `json:"status,omitempty"`
	CreatedBy    *uint             `json:"created_by,omitempty"`
}

type UpdateProductSerialNumberRequest struct {
	ProductID    *uint             `json:"product_id,omitempty"`
	SerialNumber *string           `json:"serial_number,omitempty" validate:"omitempty,min=3,max=255"`
	Status       *models.SNStatus  `json:"status,omitempty"`
}

// Category request structures
type CreateCategoryRequest struct {
	Name      string            `json:"name" validate:"required,min=2,max=255"`
	Status    models.StatusUmum `json:"status,omitempty"`
	CreatedBy *uint             `json:"created_by,omitempty"`
}

type UpdateCategoryRequest struct {
	Name   *string            `json:"name,omitempty" validate:"omitempty,min=2,max=255"`
	Status *models.StatusUmum `json:"status,omitempty"`
}

// Supplier request structures
type CreateSupplierRequest struct {
	SupplierName      string            `json:"supplier_name" validate:"required,min=2,max=255"`
	ContactPersonName string            `json:"contact_person_name" validate:"required,min=2,max=255"`
	PhoneNumber       string            `json:"phone_number" validate:"required,min=10,max=20"`
	Address           *string           `json:"address,omitempty"`
	Status            models.StatusUmum `json:"status,omitempty"`
	CreatedBy         *uint             `json:"created_by,omitempty"`
}

type UpdateSupplierRequest struct {
	SupplierName      *string            `json:"supplier_name,omitempty" validate:"omitempty,min=2,max=255"`
	ContactPersonName *string            `json:"contact_person_name,omitempty" validate:"omitempty,min=2,max=255"`
	PhoneNumber       *string            `json:"phone_number,omitempty" validate:"omitempty,min=10,max=20"`
	Address           *string            `json:"address,omitempty"`
	Status            *models.StatusUmum `json:"status,omitempty"`
}

// Unit Type request structures
type CreateUnitTypeRequest struct {
	Name      string            `json:"name" validate:"required,min=1,max=50"`
	Status    models.StatusUmum `json:"status,omitempty"`
	CreatedBy *uint             `json:"created_by,omitempty"`
}

type UpdateUnitTypeRequest struct {
	Name   *string            `json:"name,omitempty" validate:"omitempty,min=1,max=50"`
	Status *models.StatusUmum `json:"status,omitempty"`
}

// Usecase interfaces
type ProductUsecase interface {
	CreateProduct(ctx context.Context, req CreateProductRequest) (*models.Product, error)
	GetProduct(ctx context.Context, id uint) (*models.Product, error)
	GetProductBySKU(ctx context.Context, sku string) (*models.Product, error)
	GetProductByBarcode(ctx context.Context, barcode string) (*models.Product, error)
	UpdateProduct(ctx context.Context, id uint, req UpdateProductRequest) (*models.Product, error)
	DeleteProduct(ctx context.Context, id uint) error
	ListProducts(ctx context.Context, limit, offset int) ([]*models.Product, error)
	GetProductsByCategory(ctx context.Context, categoryID uint) ([]*models.Product, error)
	GetProductsBySupplier(ctx context.Context, supplierID uint) ([]*models.Product, error)
	GetProductsByUsageStatus(ctx context.Context, status models.ProductUsageStatus) ([]*models.Product, error)
	SearchProducts(ctx context.Context, query string, limit, offset int) ([]*models.Product, error)
	UpdateProductStock(ctx context.Context, productID uint, quantity int) error
	GetLowStockProducts(ctx context.Context, threshold int) ([]*models.Product, error)
}

type ProductSerialNumberUsecase interface {
	CreateProductSerialNumber(ctx context.Context, req CreateProductSerialNumberRequest) (*models.ProductSerialNumber, error)
	GetProductSerialNumber(ctx context.Context, id uint) (*models.ProductSerialNumber, error)
	GetProductSerialNumberBySerial(ctx context.Context, serialNumber string) (*models.ProductSerialNumber, error)
	UpdateProductSerialNumber(ctx context.Context, id uint, req UpdateProductSerialNumberRequest) (*models.ProductSerialNumber, error)
	DeleteProductSerialNumber(ctx context.Context, id uint) error
	ListProductSerialNumbers(ctx context.Context, limit, offset int) ([]*models.ProductSerialNumber, error)
	GetProductSerialNumbersByProduct(ctx context.Context, productID uint) ([]*models.ProductSerialNumber, error)
	GetProductSerialNumbersByStatus(ctx context.Context, status models.SNStatus) ([]*models.ProductSerialNumber, error)
	UpdateProductSerialNumberStatus(ctx context.Context, id uint, status models.SNStatus) error
}

type CategoryUsecase interface {
	CreateCategory(ctx context.Context, req CreateCategoryRequest) (*models.Category, error)
	GetCategory(ctx context.Context, id uint) (*models.Category, error)
	GetCategoryByName(ctx context.Context, name string) (*models.Category, error)
	UpdateCategory(ctx context.Context, id uint, req UpdateCategoryRequest) (*models.Category, error)
	DeleteCategory(ctx context.Context, id uint) error
	ListCategories(ctx context.Context, limit, offset int) ([]*models.Category, error)
	GetCategoriesByStatus(ctx context.Context, status models.StatusUmum) ([]*models.Category, error)
}

type SupplierUsecase interface {
	CreateSupplier(ctx context.Context, req CreateSupplierRequest) (*models.Supplier, error)
	GetSupplier(ctx context.Context, id uint) (*models.Supplier, error)
	GetSupplierByName(ctx context.Context, name string) (*models.Supplier, error)
	UpdateSupplier(ctx context.Context, id uint, req UpdateSupplierRequest) (*models.Supplier, error)
	DeleteSupplier(ctx context.Context, id uint) error
	ListSuppliers(ctx context.Context, limit, offset int) ([]*models.Supplier, error)
	GetSuppliersByStatus(ctx context.Context, status models.StatusUmum) ([]*models.Supplier, error)
	SearchSuppliers(ctx context.Context, query string, limit, offset int) ([]*models.Supplier, error)
}

type UnitTypeUsecase interface {
	CreateUnitType(ctx context.Context, req CreateUnitTypeRequest) (*models.UnitType, error)
	GetUnitType(ctx context.Context, id uint) (*models.UnitType, error)
	GetUnitTypeByName(ctx context.Context, name string) (*models.UnitType, error)
	UpdateUnitType(ctx context.Context, id uint, req UpdateUnitTypeRequest) (*models.UnitType, error)
	DeleteUnitType(ctx context.Context, id uint) error
	ListUnitTypes(ctx context.Context, limit, offset int) ([]*models.UnitType, error)
	GetUnitTypesByStatus(ctx context.Context, status models.StatusUmum) ([]*models.UnitType, error)
}