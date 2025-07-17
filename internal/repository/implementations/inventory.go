package implementations

import (
	"boilerplate/internal/models"
	"boilerplate/internal/repository/interfaces"
	"context"

	"gorm.io/gorm"
)

// ProductRepository implements the product repository interface
type ProductRepository struct {
	db *gorm.DB
}

// NewProductRepository creates a new product repository
func NewProductRepository(db *gorm.DB) interfaces.ProductRepository {
	return &ProductRepository{db: db}
}

// Create creates a new product
func (r *ProductRepository) Create(ctx context.Context, product *models.Product) error {
	return r.db.WithContext(ctx).Create(product).Error
}

// GetByID retrieves a product by ID
func (r *ProductRepository) GetByID(ctx context.Context, id uint) (*models.Product, error) {
	var product models.Product
	err := r.db.WithContext(ctx).
		Preload("Category").
		Preload("Supplier").
		Preload("UnitType").
		Preload("SerialNumbers").
		First(&product, id).Error
	if err != nil {
		return nil, err
	}
	return &product, nil
}

// GetBySKU retrieves a product by SKU
func (r *ProductRepository) GetBySKU(ctx context.Context, sku string) (*models.Product, error) {
	var product models.Product
	err := r.db.WithContext(ctx).
		Preload("Category").
		Preload("Supplier").
		Preload("UnitType").
		Preload("SerialNumbers").
		Where("sku = ?", sku).
		First(&product).Error
	if err != nil {
		return nil, err
	}
	return &product, nil
}

// GetByBarcode retrieves a product by barcode
func (r *ProductRepository) GetByBarcode(ctx context.Context, barcode string) (*models.Product, error) {
	var product models.Product
	err := r.db.WithContext(ctx).
		Preload("Category").
		Preload("Supplier").
		Preload("UnitType").
		Preload("SerialNumbers").
		Where("barcode = ?", barcode).
		First(&product).Error
	if err != nil {
		return nil, err
	}
	return &product, nil
}

// Update updates a product
func (r *ProductRepository) Update(ctx context.Context, product *models.Product) error {
	return r.db.WithContext(ctx).Save(product).Error
}

// Delete soft deletes a product
func (r *ProductRepository) Delete(ctx context.Context, id uint) error {
	return r.db.WithContext(ctx).Delete(&models.Product{}, id).Error
}

// List retrieves products with pagination
func (r *ProductRepository) List(ctx context.Context, limit, offset int) ([]*models.Product, error) {
	var products []*models.Product
	err := r.db.WithContext(ctx).
		Preload("Category").
		Preload("Supplier").
		Preload("UnitType").
		Preload("SerialNumbers").
		Limit(limit).
		Offset(offset).
		Find(&products).Error
	if err != nil {
		return nil, err
	}
	return products, nil
}

// GetByCategoryID retrieves products by category ID
func (r *ProductRepository) GetByCategoryID(ctx context.Context, categoryID uint) ([]*models.Product, error) {
	var products []*models.Product
	err := r.db.WithContext(ctx).
		Preload("Category").
		Preload("Supplier").
		Preload("UnitType").
		Preload("SerialNumbers").
		Where("category_id = ?", categoryID).
		Find(&products).Error
	if err != nil {
		return nil, err
	}
	return products, nil
}

// GetBySupplierID retrieves products by supplier ID
func (r *ProductRepository) GetBySupplierID(ctx context.Context, supplierID uint) ([]*models.Product, error) {
	var products []*models.Product
	err := r.db.WithContext(ctx).
		Preload("Category").
		Preload("Supplier").
		Preload("UnitType").
		Preload("SerialNumbers").
		Where("supplier_id = ?", supplierID).
		Find(&products).Error
	if err != nil {
		return nil, err
	}
	return products, nil
}

// GetByUsageStatus retrieves products by usage status
func (r *ProductRepository) GetByUsageStatus(ctx context.Context, status models.ProductUsageStatus) ([]*models.Product, error) {
	var products []*models.Product
	err := r.db.WithContext(ctx).
		Preload("Category").
		Preload("Supplier").
		Preload("UnitType").
		Preload("SerialNumbers").
		Where("usage_status = ?", status).
		Find(&products).Error
	if err != nil {
		return nil, err
	}
	return products, nil
}

// Search searches products by name, description, or SKU
func (r *ProductRepository) Search(ctx context.Context, query string, limit, offset int) ([]*models.Product, error) {
	var products []*models.Product
	searchQuery := "%" + query + "%"
	
	err := r.db.WithContext(ctx).
		Preload("Category").
		Preload("Supplier").
		Preload("UnitType").
		Preload("SerialNumbers").
		Where("product_name LIKE ? OR product_description LIKE ? OR sku LIKE ?", searchQuery, searchQuery, searchQuery).
		Limit(limit).
		Offset(offset).
		Find(&products).Error
	
	if err != nil {
		return nil, err
	}
	return products, nil
}

// UpdateStock updates product stock
func (r *ProductRepository) UpdateStock(ctx context.Context, productID uint, quantity int) error {
	return r.db.WithContext(ctx).
		Model(&models.Product{}).
		Where("product_id = ?", productID).
		Update("stock", gorm.Expr("stock + ?", quantity)).Error
}

// GetLowStock retrieves products with stock below threshold
func (r *ProductRepository) GetLowStock(ctx context.Context, threshold int) ([]*models.Product, error) {
	var products []*models.Product
	err := r.db.WithContext(ctx).
		Preload("Category").
		Preload("Supplier").
		Preload("UnitType").
		Preload("SerialNumbers").
		Where("stock <= ?", threshold).
		Find(&products).Error
	if err != nil {
		return nil, err
	}
	return products, nil
}

// ProductSerialNumberRepository implements the product serial number repository interface
type ProductSerialNumberRepository struct {
	db *gorm.DB
}

// NewProductSerialNumberRepository creates a new product serial number repository
func NewProductSerialNumberRepository(db *gorm.DB) interfaces.ProductSerialNumberRepository {
	return &ProductSerialNumberRepository{db: db}
}

// Create creates a new product serial number
func (r *ProductSerialNumberRepository) Create(ctx context.Context, serialNumber *models.ProductSerialNumber) error {
	return r.db.WithContext(ctx).Create(serialNumber).Error
}

// GetByID retrieves a product serial number by ID
func (r *ProductSerialNumberRepository) GetByID(ctx context.Context, id uint) (*models.ProductSerialNumber, error) {
	var serialNumber models.ProductSerialNumber
	err := r.db.WithContext(ctx).Preload("Product").First(&serialNumber, id).Error
	if err != nil {
		return nil, err
	}
	return &serialNumber, nil
}

// GetBySerialNumber retrieves a product serial number by serial number
func (r *ProductSerialNumberRepository) GetBySerialNumber(ctx context.Context, serialNumber string) (*models.ProductSerialNumber, error) {
	var sn models.ProductSerialNumber
	err := r.db.WithContext(ctx).Preload("Product").Where("serial_number = ?", serialNumber).First(&sn).Error
	if err != nil {
		return nil, err
	}
	return &sn, nil
}

// Update updates a product serial number
func (r *ProductSerialNumberRepository) Update(ctx context.Context, serialNumber *models.ProductSerialNumber) error {
	return r.db.WithContext(ctx).Save(serialNumber).Error
}

// Delete soft deletes a product serial number
func (r *ProductSerialNumberRepository) Delete(ctx context.Context, id uint) error {
	return r.db.WithContext(ctx).Delete(&models.ProductSerialNumber{}, id).Error
}

// List retrieves product serial numbers with pagination
func (r *ProductSerialNumberRepository) List(ctx context.Context, limit, offset int) ([]*models.ProductSerialNumber, error) {
	var serialNumbers []*models.ProductSerialNumber
	err := r.db.WithContext(ctx).Preload("Product").Limit(limit).Offset(offset).Find(&serialNumbers).Error
	if err != nil {
		return nil, err
	}
	return serialNumbers, nil
}

// GetByProductID retrieves product serial numbers by product ID
func (r *ProductSerialNumberRepository) GetByProductID(ctx context.Context, productID uint) ([]*models.ProductSerialNumber, error) {
	var serialNumbers []*models.ProductSerialNumber
	err := r.db.WithContext(ctx).Preload("Product").Where("product_id = ?", productID).Find(&serialNumbers).Error
	if err != nil {
		return nil, err
	}
	return serialNumbers, nil
}

// GetByStatus retrieves product serial numbers by status
func (r *ProductSerialNumberRepository) GetByStatus(ctx context.Context, status models.SNStatus) ([]*models.ProductSerialNumber, error) {
	var serialNumbers []*models.ProductSerialNumber
	err := r.db.WithContext(ctx).Preload("Product").Where("status = ?", status).Find(&serialNumbers).Error
	if err != nil {
		return nil, err
	}
	return serialNumbers, nil
}

// UpdateStatus updates product serial number status
func (r *ProductSerialNumberRepository) UpdateStatus(ctx context.Context, id uint, status models.SNStatus) error {
	return r.db.WithContext(ctx).
		Model(&models.ProductSerialNumber{}).
		Where("serial_number_id = ?", id).
		Update("status", status).Error
}

// CategoryRepository implements the category repository interface
type CategoryRepository struct {
	db *gorm.DB
}

// NewCategoryRepository creates a new category repository
func NewCategoryRepository(db *gorm.DB) interfaces.CategoryRepository {
	return &CategoryRepository{db: db}
}

// Create creates a new category
func (r *CategoryRepository) Create(ctx context.Context, category *models.Category) error {
	return r.db.WithContext(ctx).Create(category).Error
}

// GetByID retrieves a category by ID
func (r *CategoryRepository) GetByID(ctx context.Context, id uint) (*models.Category, error) {
	var category models.Category
	err := r.db.WithContext(ctx).Preload("Products").First(&category, id).Error
	if err != nil {
		return nil, err
	}
	return &category, nil
}

// GetByName retrieves a category by name
func (r *CategoryRepository) GetByName(ctx context.Context, name string) (*models.Category, error) {
	var category models.Category
	err := r.db.WithContext(ctx).Preload("Products").Where("name = ?", name).First(&category).Error
	if err != nil {
		return nil, err
	}
	return &category, nil
}

// Update updates a category
func (r *CategoryRepository) Update(ctx context.Context, category *models.Category) error {
	return r.db.WithContext(ctx).Save(category).Error
}

// Delete soft deletes a category
func (r *CategoryRepository) Delete(ctx context.Context, id uint) error {
	return r.db.WithContext(ctx).Delete(&models.Category{}, id).Error
}

// List retrieves categories with pagination
func (r *CategoryRepository) List(ctx context.Context, limit, offset int) ([]*models.Category, error) {
	var categories []*models.Category
	err := r.db.WithContext(ctx).Preload("Products").Limit(limit).Offset(offset).Find(&categories).Error
	if err != nil {
		return nil, err
	}
	return categories, nil
}

// GetByStatus retrieves categories by status
func (r *CategoryRepository) GetByStatus(ctx context.Context, status models.StatusUmum) ([]*models.Category, error) {
	var categories []*models.Category
	err := r.db.WithContext(ctx).Preload("Products").Where("status = ?", status).Find(&categories).Error
	if err != nil {
		return nil, err
	}
	return categories, nil
}

// SupplierRepository implements the supplier repository interface
type SupplierRepository struct {
	db *gorm.DB
}

// NewSupplierRepository creates a new supplier repository
func NewSupplierRepository(db *gorm.DB) interfaces.SupplierRepository {
	return &SupplierRepository{db: db}
}

// Create creates a new supplier
func (r *SupplierRepository) Create(ctx context.Context, supplier *models.Supplier) error {
	return r.db.WithContext(ctx).Create(supplier).Error
}

// GetByID retrieves a supplier by ID
func (r *SupplierRepository) GetByID(ctx context.Context, id uint) (*models.Supplier, error) {
	var supplier models.Supplier
	err := r.db.WithContext(ctx).Preload("Products").First(&supplier, id).Error
	if err != nil {
		return nil, err
	}
	return &supplier, nil
}

// GetByName retrieves a supplier by name
func (r *SupplierRepository) GetByName(ctx context.Context, name string) (*models.Supplier, error) {
	var supplier models.Supplier
	err := r.db.WithContext(ctx).Preload("Products").Where("supplier_name = ?", name).First(&supplier).Error
	if err != nil {
		return nil, err
	}
	return &supplier, nil
}

// Update updates a supplier
func (r *SupplierRepository) Update(ctx context.Context, supplier *models.Supplier) error {
	return r.db.WithContext(ctx).Save(supplier).Error
}

// Delete soft deletes a supplier
func (r *SupplierRepository) Delete(ctx context.Context, id uint) error {
	return r.db.WithContext(ctx).Delete(&models.Supplier{}, id).Error
}

// List retrieves suppliers with pagination
func (r *SupplierRepository) List(ctx context.Context, limit, offset int) ([]*models.Supplier, error) {
	var suppliers []*models.Supplier
	err := r.db.WithContext(ctx).Preload("Products").Limit(limit).Offset(offset).Find(&suppliers).Error
	if err != nil {
		return nil, err
	}
	return suppliers, nil
}

// GetByStatus retrieves suppliers by status
func (r *SupplierRepository) GetByStatus(ctx context.Context, status models.StatusUmum) ([]*models.Supplier, error) {
	var suppliers []*models.Supplier
	err := r.db.WithContext(ctx).Preload("Products").Where("status = ?", status).Find(&suppliers).Error
	if err != nil {
		return nil, err
	}
	return suppliers, nil
}

// Search searches suppliers by name, contact person, or phone number
func (r *SupplierRepository) Search(ctx context.Context, query string, limit, offset int) ([]*models.Supplier, error) {
	var suppliers []*models.Supplier
	searchQuery := "%" + query + "%"
	
	err := r.db.WithContext(ctx).
		Preload("Products").
		Where("supplier_name LIKE ? OR contact_person_name LIKE ? OR phone_number LIKE ?", searchQuery, searchQuery, searchQuery).
		Limit(limit).
		Offset(offset).
		Find(&suppliers).Error
	
	if err != nil {
		return nil, err
	}
	return suppliers, nil
}

// UnitTypeRepository implements the unit type repository interface
type UnitTypeRepository struct {
	db *gorm.DB
}

// NewUnitTypeRepository creates a new unit type repository
func NewUnitTypeRepository(db *gorm.DB) interfaces.UnitTypeRepository {
	return &UnitTypeRepository{db: db}
}

// Create creates a new unit type
func (r *UnitTypeRepository) Create(ctx context.Context, unitType *models.UnitType) error {
	return r.db.WithContext(ctx).Create(unitType).Error
}

// GetByID retrieves a unit type by ID
func (r *UnitTypeRepository) GetByID(ctx context.Context, id uint) (*models.UnitType, error) {
	var unitType models.UnitType
	err := r.db.WithContext(ctx).Preload("Products").First(&unitType, id).Error
	if err != nil {
		return nil, err
	}
	return &unitType, nil
}

// GetByName retrieves a unit type by name
func (r *UnitTypeRepository) GetByName(ctx context.Context, name string) (*models.UnitType, error) {
	var unitType models.UnitType
	err := r.db.WithContext(ctx).Preload("Products").Where("name = ?", name).First(&unitType).Error
	if err != nil {
		return nil, err
	}
	return &unitType, nil
}

// Update updates a unit type
func (r *UnitTypeRepository) Update(ctx context.Context, unitType *models.UnitType) error {
	return r.db.WithContext(ctx).Save(unitType).Error
}

// Delete soft deletes a unit type
func (r *UnitTypeRepository) Delete(ctx context.Context, id uint) error {
	return r.db.WithContext(ctx).Delete(&models.UnitType{}, id).Error
}

// List retrieves unit types with pagination
func (r *UnitTypeRepository) List(ctx context.Context, limit, offset int) ([]*models.UnitType, error) {
	var unitTypes []*models.UnitType
	err := r.db.WithContext(ctx).Preload("Products").Limit(limit).Offset(offset).Find(&unitTypes).Error
	if err != nil {
		return nil, err
	}
	return unitTypes, nil
}

// GetByStatus retrieves unit types by status
func (r *UnitTypeRepository) GetByStatus(ctx context.Context, status models.StatusUmum) ([]*models.UnitType, error) {
	var unitTypes []*models.UnitType
	err := r.db.WithContext(ctx).Preload("Products").Where("status = ?", status).Find(&unitTypes).Error
	if err != nil {
		return nil, err
	}
	return unitTypes, nil
}