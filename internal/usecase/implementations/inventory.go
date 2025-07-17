package implementations

import (
	"boilerplate/internal/models"
	"boilerplate/internal/repository"
	"boilerplate/internal/usecase/interfaces"
	"context"
	"errors"
	"time"

	"gorm.io/gorm"
)

// ProductUsecase implements the product usecase interface
type ProductUsecase struct {
	repo *repository.RepositoryManager
}

// NewProductUsecase creates a new product usecase
func NewProductUsecase(repo *repository.RepositoryManager) interfaces.ProductUsecase {
	return &ProductUsecase{repo: repo}
}

// CreateProduct creates a new product
func (u *ProductUsecase) CreateProduct(ctx context.Context, req interfaces.CreateProductRequest) (*models.Product, error) {
	// Validate category exists if provided
	if req.CategoryID != nil {
		_, err := u.repo.Category.GetByID(ctx, *req.CategoryID)
		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return nil, errors.New("category not found")
			}
			return nil, err
		}
	}

	// Validate supplier exists if provided
	if req.SupplierID != nil {
		_, err := u.repo.Supplier.GetByID(ctx, *req.SupplierID)
		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return nil, errors.New("supplier not found")
			}
			return nil, err
		}
	}

	// Validate unit type exists if provided
	if req.UnitTypeID != nil {
		_, err := u.repo.UnitType.GetByID(ctx, *req.UnitTypeID)
		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return nil, errors.New("unit type not found")
			}
			return nil, err
		}
	}

	// Check if SKU already exists
	if req.SKU != nil {
		existingProduct, err := u.repo.Product.GetBySKU(ctx, *req.SKU)
		if err == nil && existingProduct != nil {
			return nil, errors.New("product with this SKU already exists")
		}
		if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, err
		}
	}

	// Check if barcode already exists
	if req.Barcode != nil {
		existingProduct, err := u.repo.Product.GetByBarcode(ctx, *req.Barcode)
		if err == nil && existingProduct != nil {
			return nil, errors.New("product with this barcode already exists")
		}
		if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, err
		}
	}

	product := &models.Product{
		ProductName:        req.ProductName,
		ProductDescription: req.ProductDescription,
		ProductImage:       req.ProductImage,
		CostPrice:          req.CostPrice,
		SellingPrice:       req.SellingPrice,
		Stock:              req.Stock,
		SKU:                req.SKU,
		Barcode:            req.Barcode,
		HasSerialNumber:    req.HasSerialNumber,
		ShelfLocation:      req.ShelfLocation,
		UsageStatus:        req.UsageStatus,
		IsActive:           req.IsActive,
		CategoryID:         req.CategoryID,
		SupplierID:         req.SupplierID,
		UnitTypeID:         req.UnitTypeID,
		CreatedBy:          req.CreatedBy,
		CreatedAt:          time.Now(),
		UpdatedAt:          time.Now(),
	}

	if err := u.repo.Product.Create(ctx, product); err != nil {
		return nil, err
	}

	return product, nil
}

// GetProduct retrieves a product by ID
func (u *ProductUsecase) GetProduct(ctx context.Context, id uint) (*models.Product, error) {
	product, err := u.repo.Product.GetByID(ctx, id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("product not found")
		}
		return nil, err
	}
	return product, nil
}

// GetProductBySKU retrieves a product by SKU
func (u *ProductUsecase) GetProductBySKU(ctx context.Context, sku string) (*models.Product, error) {
	product, err := u.repo.Product.GetBySKU(ctx, sku)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("product not found")
		}
		return nil, err
	}
	return product, nil
}

// GetProductByBarcode retrieves a product by barcode
func (u *ProductUsecase) GetProductByBarcode(ctx context.Context, barcode string) (*models.Product, error) {
	product, err := u.repo.Product.GetByBarcode(ctx, barcode)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("product not found")
		}
		return nil, err
	}
	return product, nil
}

// UpdateProduct updates a product
func (u *ProductUsecase) UpdateProduct(ctx context.Context, id uint, req interfaces.UpdateProductRequest) (*models.Product, error) {
	product, err := u.repo.Product.GetByID(ctx, id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("product not found")
		}
		return nil, err
	}

	// Validate category exists if provided
	if req.CategoryID != nil {
		_, err := u.repo.Category.GetByID(ctx, *req.CategoryID)
		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return nil, errors.New("category not found")
			}
			return nil, err
		}
	}

	// Validate supplier exists if provided
	if req.SupplierID != nil {
		_, err := u.repo.Supplier.GetByID(ctx, *req.SupplierID)
		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return nil, errors.New("supplier not found")
			}
			return nil, err
		}
	}

	// Validate unit type exists if provided
	if req.UnitTypeID != nil {
		_, err := u.repo.UnitType.GetByID(ctx, *req.UnitTypeID)
		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return nil, errors.New("unit type not found")
			}
			return nil, err
		}
	}

	// Check if SKU already exists (if being updated)
	if req.SKU != nil && (product.SKU == nil || *req.SKU != *product.SKU) {
		existingProduct, err := u.repo.Product.GetBySKU(ctx, *req.SKU)
		if err == nil && existingProduct != nil && existingProduct.ProductID != id {
			return nil, errors.New("product with this SKU already exists")
		}
		if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, err
		}
	}

	// Check if barcode already exists (if being updated)
	if req.Barcode != nil && (product.Barcode == nil || *req.Barcode != *product.Barcode) {
		existingProduct, err := u.repo.Product.GetByBarcode(ctx, *req.Barcode)
		if err == nil && existingProduct != nil && existingProduct.ProductID != id {
			return nil, errors.New("product with this barcode already exists")
		}
		if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, err
		}
	}

	// Update fields if provided
	if req.ProductName != nil {
		product.ProductName = *req.ProductName
	}
	if req.ProductDescription != nil {
		product.ProductDescription = req.ProductDescription
	}
	if req.ProductImage != nil {
		product.ProductImage = req.ProductImage
	}
	if req.CostPrice != nil {
		product.CostPrice = *req.CostPrice
	}
	if req.SellingPrice != nil {
		product.SellingPrice = *req.SellingPrice
	}
	if req.Stock != nil {
		product.Stock = *req.Stock
	}
	if req.SKU != nil {
		product.SKU = req.SKU
	}
	if req.Barcode != nil {
		product.Barcode = req.Barcode
	}
	if req.HasSerialNumber != nil {
		product.HasSerialNumber = *req.HasSerialNumber
	}
	if req.ShelfLocation != nil {
		product.ShelfLocation = req.ShelfLocation
	}
	if req.UsageStatus != nil {
		product.UsageStatus = *req.UsageStatus
	}
	if req.IsActive != nil {
		product.IsActive = *req.IsActive
	}
	if req.CategoryID != nil {
		product.CategoryID = req.CategoryID
	}
	if req.SupplierID != nil {
		product.SupplierID = req.SupplierID
	}
	if req.UnitTypeID != nil {
		product.UnitTypeID = req.UnitTypeID
	}
	product.UpdatedAt = time.Now()

	if err := u.repo.Product.Update(ctx, product); err != nil {
		return nil, err
	}

	return product, nil
}

// DeleteProduct deletes a product
func (u *ProductUsecase) DeleteProduct(ctx context.Context, id uint) error {
	_, err := u.repo.Product.GetByID(ctx, id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New("product not found")
		}
		return err
	}

	// TODO: Check if product has serial numbers or is used in transactions

	return u.repo.Product.Delete(ctx, id)
}

// ListProducts retrieves products with pagination
func (u *ProductUsecase) ListProducts(ctx context.Context, limit, offset int) ([]*models.Product, error) {
	if limit <= 0 {
		limit = 10
	}
	if offset < 0 {
		offset = 0
	}

	return u.repo.Product.List(ctx, limit, offset)
}

// GetProductsByCategory retrieves products by category
func (u *ProductUsecase) GetProductsByCategory(ctx context.Context, categoryID uint) ([]*models.Product, error) {
	return u.repo.Product.GetByCategoryID(ctx, categoryID)
}

// GetProductsBySupplier retrieves products by supplier
func (u *ProductUsecase) GetProductsBySupplier(ctx context.Context, supplierID uint) ([]*models.Product, error) {
	return u.repo.Product.GetBySupplierID(ctx, supplierID)
}

// GetProductsByUsageStatus retrieves products by usage status
func (u *ProductUsecase) GetProductsByUsageStatus(ctx context.Context, status models.ProductUsageStatus) ([]*models.Product, error) {
	return u.repo.Product.GetByUsageStatus(ctx, status)
}

// SearchProducts searches products
func (u *ProductUsecase) SearchProducts(ctx context.Context, query string, limit, offset int) ([]*models.Product, error) {
	if limit <= 0 {
		limit = 10
	}
	if offset < 0 {
		offset = 0
	}

	return u.repo.Product.Search(ctx, query, limit, offset)
}

// UpdateProductStock updates product stock
func (u *ProductUsecase) UpdateProductStock(ctx context.Context, productID uint, quantity int) error {
	_, err := u.repo.Product.GetByID(ctx, productID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New("product not found")
		}
		return err
	}

	return u.repo.Product.UpdateStock(ctx, productID, quantity)
}

// GetLowStockProducts retrieves products with low stock
func (u *ProductUsecase) GetLowStockProducts(ctx context.Context, threshold int) ([]*models.Product, error) {
	return u.repo.Product.GetLowStock(ctx, threshold)
}

// CategoryUsecase implements the category usecase interface
type CategoryUsecase struct {
	repo *repository.RepositoryManager
}

// NewCategoryUsecase creates a new category usecase
func NewCategoryUsecase(repo *repository.RepositoryManager) interfaces.CategoryUsecase {
	return &CategoryUsecase{repo: repo}
}

// CreateCategory creates a new category
func (u *CategoryUsecase) CreateCategory(ctx context.Context, req interfaces.CreateCategoryRequest) (*models.Category, error) {
	// Check if category name already exists
	existingCategory, err := u.repo.Category.GetByName(ctx, req.Name)
	if err == nil && existingCategory != nil {
		return nil, errors.New("category with this name already exists")
	}
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}

	// Set default status if not provided
	status := req.Status
	if status == "" {
		status = models.StatusAktif
	}

	category := &models.Category{
		Name:      req.Name,
		Status:    status,
		CreatedBy: req.CreatedBy,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	if err := u.repo.Category.Create(ctx, category); err != nil {
		return nil, err
	}

	return category, nil
}

// GetCategory retrieves a category by ID
func (u *CategoryUsecase) GetCategory(ctx context.Context, id uint) (*models.Category, error) {
	category, err := u.repo.Category.GetByID(ctx, id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("category not found")
		}
		return nil, err
	}
	return category, nil
}

// GetCategoryByName retrieves a category by name
func (u *CategoryUsecase) GetCategoryByName(ctx context.Context, name string) (*models.Category, error) {
	category, err := u.repo.Category.GetByName(ctx, name)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("category not found")
		}
		return nil, err
	}
	return category, nil
}

// UpdateCategory updates a category
func (u *CategoryUsecase) UpdateCategory(ctx context.Context, id uint, req interfaces.UpdateCategoryRequest) (*models.Category, error) {
	category, err := u.repo.Category.GetByID(ctx, id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("category not found")
		}
		return nil, err
	}

	// Check if category name already exists (if being updated)
	if req.Name != nil && *req.Name != category.Name {
		existingCategory, err := u.repo.Category.GetByName(ctx, *req.Name)
		if err == nil && existingCategory != nil && existingCategory.CategoryID != id {
			return nil, errors.New("category with this name already exists")
		}
		if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, err
		}
	}

	// Update fields if provided
	if req.Name != nil {
		category.Name = *req.Name
	}
	if req.Status != nil {
		category.Status = *req.Status
	}
	category.UpdatedAt = time.Now()

	if err := u.repo.Category.Update(ctx, category); err != nil {
		return nil, err
	}

	return category, nil
}

// DeleteCategory deletes a category
func (u *CategoryUsecase) DeleteCategory(ctx context.Context, id uint) error {
	_, err := u.repo.Category.GetByID(ctx, id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New("category not found")
		}
		return err
	}

	// Check if category has products
	products, err := u.repo.Product.GetByCategoryID(ctx, id)
	if err != nil {
		return err
	}
	if len(products) > 0 {
		return errors.New("cannot delete category with existing products")
	}

	return u.repo.Category.Delete(ctx, id)
}

// ListCategories retrieves categories with pagination
func (u *CategoryUsecase) ListCategories(ctx context.Context, limit, offset int) ([]*models.Category, error) {
	if limit <= 0 {
		limit = 10
	}
	if offset < 0 {
		offset = 0
	}

	return u.repo.Category.List(ctx, limit, offset)
}

// GetCategoriesByStatus retrieves categories by status
func (u *CategoryUsecase) GetCategoriesByStatus(ctx context.Context, status models.StatusUmum) ([]*models.Category, error) {
	return u.repo.Category.GetByStatus(ctx, status)
}

// SupplierUsecase implements the supplier usecase interface
type SupplierUsecase struct {
	repo *repository.RepositoryManager
}

// NewSupplierUsecase creates a new supplier usecase
func NewSupplierUsecase(repo *repository.RepositoryManager) interfaces.SupplierUsecase {
	return &SupplierUsecase{repo: repo}
}

// CreateSupplier creates a new supplier
func (u *SupplierUsecase) CreateSupplier(ctx context.Context, req interfaces.CreateSupplierRequest) (*models.Supplier, error) {
	// Check if supplier name already exists
	existingSupplier, err := u.repo.Supplier.GetByName(ctx, req.SupplierName)
	if err == nil && existingSupplier != nil {
		return nil, errors.New("supplier with this name already exists")
	}
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}

	// Set default status if not provided
	status := req.Status
	if status == "" {
		status = models.StatusAktif
	}

	supplier := &models.Supplier{
		SupplierName:      req.SupplierName,
		ContactPersonName: req.ContactPersonName,
		PhoneNumber:       req.PhoneNumber,
		Address:           req.Address,
		Status:            status,
		CreatedBy:         req.CreatedBy,
		CreatedAt:         time.Now(),
		UpdatedAt:         time.Now(),
	}

	if err := u.repo.Supplier.Create(ctx, supplier); err != nil {
		return nil, err
	}

	return supplier, nil
}

// GetSupplier retrieves a supplier by ID
func (u *SupplierUsecase) GetSupplier(ctx context.Context, id uint) (*models.Supplier, error) {
	supplier, err := u.repo.Supplier.GetByID(ctx, id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("supplier not found")
		}
		return nil, err
	}
	return supplier, nil
}

// GetSupplierByName retrieves a supplier by name
func (u *SupplierUsecase) GetSupplierByName(ctx context.Context, name string) (*models.Supplier, error) {
	supplier, err := u.repo.Supplier.GetByName(ctx, name)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("supplier not found")
		}
		return nil, err
	}
	return supplier, nil
}

// UpdateSupplier updates a supplier
func (u *SupplierUsecase) UpdateSupplier(ctx context.Context, id uint, req interfaces.UpdateSupplierRequest) (*models.Supplier, error) {
	supplier, err := u.repo.Supplier.GetByID(ctx, id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("supplier not found")
		}
		return nil, err
	}

	// Check if supplier name already exists (if being updated)
	if req.SupplierName != nil && *req.SupplierName != supplier.SupplierName {
		existingSupplier, err := u.repo.Supplier.GetByName(ctx, *req.SupplierName)
		if err == nil && existingSupplier != nil && existingSupplier.SupplierID != id {
			return nil, errors.New("supplier with this name already exists")
		}
		if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, err
		}
	}

	// Update fields if provided
	if req.SupplierName != nil {
		supplier.SupplierName = *req.SupplierName
	}
	if req.ContactPersonName != nil {
		supplier.ContactPersonName = *req.ContactPersonName
	}
	if req.PhoneNumber != nil {
		supplier.PhoneNumber = *req.PhoneNumber
	}
	if req.Address != nil {
		supplier.Address = req.Address
	}
	if req.Status != nil {
		supplier.Status = *req.Status
	}
	supplier.UpdatedAt = time.Now()

	if err := u.repo.Supplier.Update(ctx, supplier); err != nil {
		return nil, err
	}

	return supplier, nil
}

// DeleteSupplier deletes a supplier
func (u *SupplierUsecase) DeleteSupplier(ctx context.Context, id uint) error {
	_, err := u.repo.Supplier.GetByID(ctx, id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New("supplier not found")
		}
		return err
	}

	// Check if supplier has products
	products, err := u.repo.Product.GetBySupplierID(ctx, id)
	if err != nil {
		return err
	}
	if len(products) > 0 {
		return errors.New("cannot delete supplier with existing products")
	}

	return u.repo.Supplier.Delete(ctx, id)
}

// ListSuppliers retrieves suppliers with pagination
func (u *SupplierUsecase) ListSuppliers(ctx context.Context, limit, offset int) ([]*models.Supplier, error) {
	if limit <= 0 {
		limit = 10
	}
	if offset < 0 {
		offset = 0
	}

	return u.repo.Supplier.List(ctx, limit, offset)
}

// GetSuppliersByStatus retrieves suppliers by status
func (u *SupplierUsecase) GetSuppliersByStatus(ctx context.Context, status models.StatusUmum) ([]*models.Supplier, error) {
	return u.repo.Supplier.GetByStatus(ctx, status)
}

// SearchSuppliers searches suppliers
func (u *SupplierUsecase) SearchSuppliers(ctx context.Context, query string, limit, offset int) ([]*models.Supplier, error) {
	if limit <= 0 {
		limit = 10
	}
	if offset < 0 {
		offset = 0
	}

	return u.repo.Supplier.Search(ctx, query, limit, offset)
}

// UnitTypeUsecase implements the unit type usecase interface
type UnitTypeUsecase struct {
	repo *repository.RepositoryManager
}

// NewUnitTypeUsecase creates a new unit type usecase
func NewUnitTypeUsecase(repo *repository.RepositoryManager) interfaces.UnitTypeUsecase {
	return &UnitTypeUsecase{repo: repo}
}

// CreateUnitType creates a new unit type
func (u *UnitTypeUsecase) CreateUnitType(ctx context.Context, req interfaces.CreateUnitTypeRequest) (*models.UnitType, error) {
	// Check if unit type name already exists
	existingUnitType, err := u.repo.UnitType.GetByName(ctx, req.Name)
	if err == nil && existingUnitType != nil {
		return nil, errors.New("unit type with this name already exists")
	}
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}

	// Set default status if not provided
	status := req.Status
	if status == "" {
		status = models.StatusAktif
	}

	unitType := &models.UnitType{
		Name:      req.Name,
		Status:    status,
		CreatedBy: req.CreatedBy,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	if err := u.repo.UnitType.Create(ctx, unitType); err != nil {
		return nil, err
	}

	return unitType, nil
}

// GetUnitType retrieves a unit type by ID
func (u *UnitTypeUsecase) GetUnitType(ctx context.Context, id uint) (*models.UnitType, error) {
	unitType, err := u.repo.UnitType.GetByID(ctx, id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("unit type not found")
		}
		return nil, err
	}
	return unitType, nil
}

// GetUnitTypeByName retrieves a unit type by name
func (u *UnitTypeUsecase) GetUnitTypeByName(ctx context.Context, name string) (*models.UnitType, error) {
	unitType, err := u.repo.UnitType.GetByName(ctx, name)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("unit type not found")
		}
		return nil, err
	}
	return unitType, nil
}

// UpdateUnitType updates a unit type
func (u *UnitTypeUsecase) UpdateUnitType(ctx context.Context, id uint, req interfaces.UpdateUnitTypeRequest) (*models.UnitType, error) {
	unitType, err := u.repo.UnitType.GetByID(ctx, id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("unit type not found")
		}
		return nil, err
	}

	// Check if unit type name already exists (if being updated)
	if req.Name != nil && *req.Name != unitType.Name {
		existingUnitType, err := u.repo.UnitType.GetByName(ctx, *req.Name)
		if err == nil && existingUnitType != nil && existingUnitType.UnitTypeID != id {
			return nil, errors.New("unit type with this name already exists")
		}
		if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, err
		}
	}

	// Update fields if provided
	if req.Name != nil {
		unitType.Name = *req.Name
	}
	if req.Status != nil {
		unitType.Status = *req.Status
	}
	unitType.UpdatedAt = time.Now()

	if err := u.repo.UnitType.Update(ctx, unitType); err != nil {
		return nil, err
	}

	return unitType, nil
}

// DeleteUnitType deletes a unit type
func (u *UnitTypeUsecase) DeleteUnitType(ctx context.Context, id uint) error {
	_, err := u.repo.UnitType.GetByID(ctx, id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New("unit type not found")
		}
		return err
	}

	// TODO: Check if unit type has products

	return u.repo.UnitType.Delete(ctx, id)
}

// ListUnitTypes retrieves unit types with pagination
func (u *UnitTypeUsecase) ListUnitTypes(ctx context.Context, limit, offset int) ([]*models.UnitType, error) {
	if limit <= 0 {
		limit = 10
	}
	if offset < 0 {
		offset = 0
	}

	return u.repo.UnitType.List(ctx, limit, offset)
}

// GetUnitTypesByStatus retrieves unit types by status
func (u *UnitTypeUsecase) GetUnitTypesByStatus(ctx context.Context, status models.StatusUmum) ([]*models.UnitType, error) {
	return u.repo.UnitType.GetByStatus(ctx, status)
}

// ProductSerialNumberUsecase implements the product serial number usecase interface
type ProductSerialNumberUsecase struct {
	repo *repository.RepositoryManager
}

// NewProductSerialNumberUsecase creates a new product serial number usecase
func NewProductSerialNumberUsecase(repo *repository.RepositoryManager) interfaces.ProductSerialNumberUsecase {
	return &ProductSerialNumberUsecase{repo: repo}
}

// CreateProductSerialNumber creates a new product serial number
func (u *ProductSerialNumberUsecase) CreateProductSerialNumber(ctx context.Context, req interfaces.CreateProductSerialNumberRequest) (*models.ProductSerialNumber, error) {
	// Check if product exists
	_, err := u.repo.Product.GetByID(ctx, req.ProductID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("product not found")
		}
		return nil, err
	}

	// Check if serial number already exists
	existingSerial, err := u.repo.ProductSerialNumber.GetBySerialNumber(ctx, req.SerialNumber)
	if err == nil && existingSerial != nil {
		return nil, errors.New("serial number already exists")
	}
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}

	// Set default status if not provided
	status := req.Status
	if status == "" {
		status = models.SNStatusTersedia
	}

	serialNumber := &models.ProductSerialNumber{
		ProductID:    req.ProductID,
		SerialNumber: req.SerialNumber,
		Status:       status,
		CreatedBy:    req.CreatedBy,
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	}

	if err := u.repo.ProductSerialNumber.Create(ctx, serialNumber); err != nil {
		return nil, err
	}

	return serialNumber, nil
}

// GetProductSerialNumber retrieves a product serial number by ID
func (u *ProductSerialNumberUsecase) GetProductSerialNumber(ctx context.Context, id uint) (*models.ProductSerialNumber, error) {
	serialNumber, err := u.repo.ProductSerialNumber.GetByID(ctx, id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("product serial number not found")
		}
		return nil, err
	}
	return serialNumber, nil
}

// GetProductSerialNumberBySerial retrieves a product serial number by serial number
func (u *ProductSerialNumberUsecase) GetProductSerialNumberBySerial(ctx context.Context, serialNumber string) (*models.ProductSerialNumber, error) {
	sn, err := u.repo.ProductSerialNumber.GetBySerialNumber(ctx, serialNumber)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("product serial number not found")
		}
		return nil, err
	}
	return sn, nil
}

// UpdateProductSerialNumber updates a product serial number
func (u *ProductSerialNumberUsecase) UpdateProductSerialNumber(ctx context.Context, id uint, req interfaces.UpdateProductSerialNumberRequest) (*models.ProductSerialNumber, error) {
	serialNumber, err := u.repo.ProductSerialNumber.GetByID(ctx, id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("product serial number not found")
		}
		return nil, err
	}

	// Check if product exists (if being updated)
	if req.ProductID != nil && *req.ProductID != serialNumber.ProductID {
		_, err := u.repo.Product.GetByID(ctx, *req.ProductID)
		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return nil, errors.New("product not found")
			}
			return nil, err
		}
	}

	// Check if serial number already exists (if being updated)
	if req.SerialNumber != nil && *req.SerialNumber != serialNumber.SerialNumber {
		existingSerial, err := u.repo.ProductSerialNumber.GetBySerialNumber(ctx, *req.SerialNumber)
		if err == nil && existingSerial != nil && existingSerial.SerialNumberID != id {
			return nil, errors.New("serial number already exists")
		}
		if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, err
		}
	}

	// Update fields if provided
	if req.ProductID != nil {
		serialNumber.ProductID = *req.ProductID
	}
	if req.SerialNumber != nil {
		serialNumber.SerialNumber = *req.SerialNumber
	}
	if req.Status != nil {
		serialNumber.Status = *req.Status
	}
	serialNumber.UpdatedAt = time.Now()

	if err := u.repo.ProductSerialNumber.Update(ctx, serialNumber); err != nil {
		return nil, err
	}

	return serialNumber, nil
}

// DeleteProductSerialNumber deletes a product serial number
func (u *ProductSerialNumberUsecase) DeleteProductSerialNumber(ctx context.Context, id uint) error {
	_, err := u.repo.ProductSerialNumber.GetByID(ctx, id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New("product serial number not found")
		}
		return err
	}

	// TODO: Check if serial number is used in transactions

	return u.repo.ProductSerialNumber.Delete(ctx, id)
}

// ListProductSerialNumbers retrieves product serial numbers with pagination
func (u *ProductSerialNumberUsecase) ListProductSerialNumbers(ctx context.Context, limit, offset int) ([]*models.ProductSerialNumber, error) {
	if limit <= 0 {
		limit = 10
	}
	if offset < 0 {
		offset = 0
	}

	return u.repo.ProductSerialNumber.List(ctx, limit, offset)
}

// GetProductSerialNumbersByProduct retrieves product serial numbers by product
func (u *ProductSerialNumberUsecase) GetProductSerialNumbersByProduct(ctx context.Context, productID uint) ([]*models.ProductSerialNumber, error) {
	return u.repo.ProductSerialNumber.GetByProductID(ctx, productID)
}

// GetProductSerialNumbersByStatus retrieves product serial numbers by status
func (u *ProductSerialNumberUsecase) GetProductSerialNumbersByStatus(ctx context.Context, status models.SNStatus) ([]*models.ProductSerialNumber, error) {
	return u.repo.ProductSerialNumber.GetByStatus(ctx, status)
}

// UpdateProductSerialNumberStatus updates product serial number status
func (u *ProductSerialNumberUsecase) UpdateProductSerialNumberStatus(ctx context.Context, id uint, status models.SNStatus) error {
	_, err := u.repo.ProductSerialNumber.GetByID(ctx, id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New("product serial number not found")
		}
		return err
	}

	return u.repo.ProductSerialNumber.UpdateStatus(ctx, id, status)
}