package implementations

import (
	"boilerplate/internal/models"
	"boilerplate/internal/repository/interfaces"
	"context"
	"database/sql"

	"github.com/jmoiron/sqlx"
)

// ProductRepositorySQLX implements ProductRepository interface using SQLx
type ProductRepositorySQLX struct {
	db *sqlx.DB
}

// NewProductRepositorySQLX creates a new product repository using SQLx
func NewProductRepositorySQLX(db *sqlx.DB) interfaces.ProductRepository {
	return &ProductRepositorySQLX{db: db}
}

func (r *ProductRepositorySQLX) Create(ctx context.Context, product *models.Product) error {
	query := `
		INSERT INTO products (product_name, product_description, product_image, cost_price, 
							 selling_price, stock, sku, barcode, has_serial_number, shelf_location, 
							 usage_status, is_active, category_id, supplier_id, unit_type_id, 
							 sourceable_id, sourceable_type, created_by)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18)
		RETURNING product_id, created_at, updated_at
	`
	
	row := r.db.QueryRowxContext(ctx, query,
		product.ProductName, product.ProductDescription, product.ProductImage,
		product.CostPrice, product.SellingPrice, product.Stock, product.SKU,
		product.Barcode, product.HasSerialNumber, product.ShelfLocation,
		product.UsageStatus, product.IsActive, product.CategoryID,
		product.SupplierID, product.UnitTypeID, product.SourceableID,
		product.SourceableType, product.CreatedBy,
	)
	
	return row.Scan(&product.ProductID, &product.CreatedAt, &product.UpdatedAt)
}

func (r *ProductRepositorySQLX) GetByID(ctx context.Context, id uint) (*models.Product, error) {
	var product models.Product
	query := `
		SELECT product_id, product_name, product_description, product_image, cost_price, 
			   selling_price, stock, sku, barcode, has_serial_number, shelf_location, 
			   usage_status, is_active, category_id, supplier_id, unit_type_id, 
			   sourceable_id, sourceable_type, created_at, updated_at, deleted_at, created_by
		FROM products 
		WHERE product_id = $1 AND deleted_at IS NULL
	`
	
	err := r.db.GetContext(ctx, &product, query, id)
	if err != nil {
		return nil, err
	}

	// Load relationships
	if err := r.loadProductRelationships(ctx, &product); err != nil {
		return nil, err
	}

	return &product, nil
}

func (r *ProductRepositorySQLX) Update(ctx context.Context, product *models.Product) error {
	query := `
		UPDATE products 
		SET product_name = $1, product_description = $2, product_image = $3, cost_price = $4, 
			selling_price = $5, stock = $6, sku = $7, barcode = $8, has_serial_number = $9, 
			shelf_location = $10, usage_status = $11, is_active = $12, category_id = $13, 
			supplier_id = $14, unit_type_id = $15, sourceable_id = $16, sourceable_type = $17, 
			updated_at = NOW()
		WHERE product_id = $18 AND deleted_at IS NULL
	`
	
	_, err := r.db.ExecContext(ctx, query,
		product.ProductName, product.ProductDescription, product.ProductImage,
		product.CostPrice, product.SellingPrice, product.Stock, product.SKU,
		product.Barcode, product.HasSerialNumber, product.ShelfLocation,
		product.UsageStatus, product.IsActive, product.CategoryID,
		product.SupplierID, product.UnitTypeID, product.SourceableID,
		product.SourceableType, product.ProductID,
	)
	
	return err
}

func (r *ProductRepositorySQLX) Delete(ctx context.Context, id uint) error {
	query := `
		UPDATE products 
		SET deleted_at = NOW(), updated_at = NOW()
		WHERE product_id = $1 AND deleted_at IS NULL
	`
	
	_, err := r.db.ExecContext(ctx, query, id)
	return err
}

func (r *ProductRepositorySQLX) List(ctx context.Context, limit, offset int) ([]*models.Product, error) {
	var products []*models.Product
	query := `
		SELECT product_id, product_name, product_description, product_image, cost_price, 
			   selling_price, stock, sku, barcode, has_serial_number, shelf_location, 
			   usage_status, is_active, category_id, supplier_id, unit_type_id, 
			   sourceable_id, sourceable_type, created_at, updated_at, deleted_at, created_by
		FROM products 
		WHERE deleted_at IS NULL
		ORDER BY created_at DESC
		LIMIT $1 OFFSET $2
	`
	
	err := r.db.SelectContext(ctx, &products, query, limit, offset)
	if err != nil {
		return nil, err
	}

	// Load relationships for each product
	for i := range products {
		if err := r.loadProductRelationships(ctx, products[i]); err != nil {
			return nil, err
		}
	}

	return products, nil
}

func (r *ProductRepositorySQLX) GetBySKU(ctx context.Context, sku string) (*models.Product, error) {
	var product models.Product
	query := `
		SELECT product_id, product_name, product_description, product_image, cost_price, 
			   selling_price, stock, sku, barcode, has_serial_number, shelf_location, 
			   usage_status, is_active, category_id, supplier_id, unit_type_id, 
			   sourceable_id, sourceable_type, created_at, updated_at, deleted_at, created_by
		FROM products 
		WHERE sku = $1 AND deleted_at IS NULL
	`
	
	err := r.db.GetContext(ctx, &product, query, sku)
	if err != nil {
		return nil, err
	}

	if err := r.loadProductRelationships(ctx, &product); err != nil {
		return nil, err
	}

	return &product, nil
}

func (r *ProductRepositorySQLX) GetByBarcode(ctx context.Context, barcode string) (*models.Product, error) {
	var product models.Product
	query := `
		SELECT product_id, product_name, product_description, product_image, cost_price, 
			   selling_price, stock, sku, barcode, has_serial_number, shelf_location, 
			   usage_status, is_active, category_id, supplier_id, unit_type_id, 
			   sourceable_id, sourceable_type, created_at, updated_at, deleted_at, created_by
		FROM products 
		WHERE barcode = $1 AND deleted_at IS NULL
	`
	
	err := r.db.GetContext(ctx, &product, query, barcode)
	if err != nil {
		return nil, err
	}

	if err := r.loadProductRelationships(ctx, &product); err != nil {
		return nil, err
	}

	return &product, nil
}

func (r *ProductRepositorySQLX) GetByCategoryID(ctx context.Context, categoryID uint) ([]*models.Product, error) {
	var products []*models.Product
	query := `
		SELECT product_id, product_name, product_description, product_image, cost_price, 
			   selling_price, stock, sku, barcode, has_serial_number, shelf_location, 
			   usage_status, is_active, category_id, supplier_id, unit_type_id, 
			   sourceable_id, sourceable_type, created_at, updated_at, deleted_at, created_by
		FROM products 
		WHERE category_id = $1 AND deleted_at IS NULL
		ORDER BY created_at DESC
	`
	
	err := r.db.SelectContext(ctx, &products, query, categoryID)
	if err != nil {
		return nil, err
	}

	for i := range products {
		if err := r.loadProductRelationships(ctx, products[i]); err != nil {
			return nil, err
		}
	}

	return products, nil
}

func (r *ProductRepositorySQLX) GetBySupplierID(ctx context.Context, supplierID uint) ([]*models.Product, error) {
	var products []*models.Product
	query := `
		SELECT product_id, product_name, product_description, product_image, cost_price, 
			   selling_price, stock, sku, barcode, has_serial_number, shelf_location, 
			   usage_status, is_active, category_id, supplier_id, unit_type_id, 
			   sourceable_id, sourceable_type, created_at, updated_at, deleted_at, created_by
		FROM products 
		WHERE supplier_id = $1 AND deleted_at IS NULL
		ORDER BY created_at DESC
	`
	
	err := r.db.SelectContext(ctx, &products, query, supplierID)
	if err != nil {
		return nil, err
	}

	for i := range products {
		if err := r.loadProductRelationships(ctx, products[i]); err != nil {
			return nil, err
		}
	}

	return products, nil
}

func (r *ProductRepositorySQLX) GetByUsageStatus(ctx context.Context, status models.ProductUsageStatus) ([]*models.Product, error) {
	var products []*models.Product
	query := `
		SELECT product_id, product_name, product_description, product_image, cost_price, 
			   selling_price, stock, sku, barcode, has_serial_number, shelf_location, 
			   usage_status, is_active, category_id, supplier_id, unit_type_id, 
			   sourceable_id, sourceable_type, created_at, updated_at, deleted_at, created_by
		FROM products 
		WHERE usage_status = $1 AND deleted_at IS NULL
		ORDER BY created_at DESC
	`
	
	err := r.db.SelectContext(ctx, &products, query, status)
	if err != nil {
		return nil, err
	}

	for i := range products {
		if err := r.loadProductRelationships(ctx, products[i]); err != nil {
			return nil, err
		}
	}

	return products, nil
}

func (r *ProductRepositorySQLX) Search(ctx context.Context, query string, limit, offset int) ([]*models.Product, error) {
	var products []*models.Product
	searchQuery := "%" + query + "%"
	
	sqlQuery := `
		SELECT product_id, product_name, product_description, product_image, cost_price, 
			   selling_price, stock, sku, barcode, has_serial_number, shelf_location, 
			   usage_status, is_active, category_id, supplier_id, unit_type_id, 
			   sourceable_id, sourceable_type, created_at, updated_at, deleted_at, created_by
		FROM products 
		WHERE (product_name ILIKE $1 OR product_description ILIKE $2 OR sku ILIKE $3 OR barcode ILIKE $4) 
		  AND deleted_at IS NULL
		ORDER BY created_at DESC
		LIMIT $5 OFFSET $6
	`
	
	err := r.db.SelectContext(ctx, &products, sqlQuery, searchQuery, searchQuery, searchQuery, searchQuery, limit, offset)
	if err != nil {
		return nil, err
	}

	for i := range products {
		if err := r.loadProductRelationships(ctx, products[i]); err != nil {
			return nil, err
		}
	}

	return products, nil
}

func (r *ProductRepositorySQLX) UpdateStock(ctx context.Context, productID uint, newStock int) error {
	query := `
		UPDATE products 
		SET stock = $1, updated_at = NOW()
		WHERE product_id = $2 AND deleted_at IS NULL
	`
	
	_, err := r.db.ExecContext(ctx, query, newStock, productID)
	return err
}

func (r *ProductRepositorySQLX) GetLowStock(ctx context.Context, threshold int) ([]*models.Product, error) {
	var products []*models.Product
	query := `
		SELECT product_id, product_name, product_description, product_image, cost_price, 
			   selling_price, stock, sku, barcode, has_serial_number, shelf_location, 
			   usage_status, is_active, category_id, supplier_id, unit_type_id, 
			   sourceable_id, sourceable_type, created_at, updated_at, deleted_at, created_by
		FROM products 
		WHERE stock <= $1 AND deleted_at IS NULL AND is_active = true
		ORDER BY stock ASC, created_at DESC
	`
	
	err := r.db.SelectContext(ctx, &products, query, threshold)
	if err != nil {
		return nil, err
	}

	for i := range products {
		if err := r.loadProductRelationships(ctx, products[i]); err != nil {
			return nil, err
		}
	}

	return products, nil
}

// Helper method to load product relationships
func (r *ProductRepositorySQLX) loadProductRelationships(ctx context.Context, product *models.Product) error {
	// Load category
	if product.CategoryID != nil {
		category, err := r.getCategoryByID(ctx, *product.CategoryID)
		if err != nil && err != sql.ErrNoRows {
			return err
		}
		product.Category = category
	}

	// Load supplier
	if product.SupplierID != nil {
		supplier, err := r.getSupplierByID(ctx, *product.SupplierID)
		if err != nil && err != sql.ErrNoRows {
			return err
		}
		product.Supplier = supplier
	}

	// Load unit type
	if product.UnitTypeID != nil {
		unitType, err := r.getUnitTypeByID(ctx, *product.UnitTypeID)
		if err != nil && err != sql.ErrNoRows {
			return err
		}
		product.UnitType = unitType
	}

	// Load serial numbers
	serialNumbers, err := r.getSerialNumbersByProductID(ctx, product.ProductID)
	if err != nil {
		return err
	}
	product.SerialNumbers = serialNumbers

	return nil
}

// Helper methods for relationships
func (r *ProductRepositorySQLX) getCategoryByID(ctx context.Context, id uint) (*models.Category, error) {
	var category models.Category
	query := `
		SELECT category_id, name, status, created_at, updated_at, deleted_at, created_by
		FROM categories 
		WHERE category_id = $1 AND deleted_at IS NULL
	`
	
	err := r.db.GetContext(ctx, &category, query, id)
	if err != nil {
		return nil, err
	}
	return &category, nil
}

func (r *ProductRepositorySQLX) getSupplierByID(ctx context.Context, id uint) (*models.Supplier, error) {
	var supplier models.Supplier
	query := `
		SELECT supplier_id, supplier_name, contact_person_name, phone_number, address, 
			   status, created_at, updated_at, deleted_at, created_by
		FROM suppliers 
		WHERE supplier_id = $1 AND deleted_at IS NULL
	`
	
	err := r.db.GetContext(ctx, &supplier, query, id)
	if err != nil {
		return nil, err
	}
	return &supplier, nil
}

func (r *ProductRepositorySQLX) getUnitTypeByID(ctx context.Context, id uint) (*models.UnitType, error) {
	var unitType models.UnitType
	query := `
		SELECT unit_type_id, name, status, created_at, updated_at, deleted_at, created_by
		FROM unit_types 
		WHERE unit_type_id = $1 AND deleted_at IS NULL
	`
	
	err := r.db.GetContext(ctx, &unitType, query, id)
	if err != nil {
		return nil, err
	}
	return &unitType, nil
}

func (r *ProductRepositorySQLX) getSerialNumbersByProductID(ctx context.Context, productID uint) ([]models.ProductSerialNumber, error) {
	var serialNumbers []models.ProductSerialNumber
	query := `
		SELECT serial_number_id, product_id, serial_number, status, 
			   created_at, updated_at, deleted_at, created_by
		FROM product_serial_numbers 
		WHERE product_id = $1 AND deleted_at IS NULL
		ORDER BY created_at DESC
	`
	
	err := r.db.SelectContext(ctx, &serialNumbers, query, productID)
	return serialNumbers, err
}

// CategoryRepositorySQLX implements CategoryRepository interface using SQLx
type CategoryRepositorySQLX struct {
	db *sqlx.DB
}

// NewCategoryRepositorySQLX creates a new category repository using SQLx
func NewCategoryRepositorySQLX(db *sqlx.DB) interfaces.CategoryRepository {
	return &CategoryRepositorySQLX{db: db}
}

func (r *CategoryRepositorySQLX) Create(ctx context.Context, category *models.Category) error {
	query := `
		INSERT INTO categories (name, status, created_by)
		VALUES ($1, $2, $3)
		RETURNING category_id, created_at, updated_at
	`
	
	row := r.db.QueryRowxContext(ctx, query, category.Name, category.Status, category.CreatedBy)
	return row.Scan(&category.CategoryID, &category.CreatedAt, &category.UpdatedAt)
}

func (r *CategoryRepositorySQLX) GetByID(ctx context.Context, id uint) (*models.Category, error) {
	var category models.Category
	query := `
		SELECT category_id, name, status, created_at, updated_at, deleted_at, created_by
		FROM categories 
		WHERE category_id = $1 AND deleted_at IS NULL
	`
	
	err := r.db.GetContext(ctx, &category, query, id)
	if err != nil {
		return nil, err
	}

	// Load products
	products, err := r.getProductsByCategoryID(ctx, id)
	if err != nil {
		return nil, err
	}
	category.Products = products

	return &category, nil
}

func (r *CategoryRepositorySQLX) Update(ctx context.Context, category *models.Category) error {
	query := `
		UPDATE categories 
		SET name = $1, status = $2, updated_at = NOW()
		WHERE category_id = $3 AND deleted_at IS NULL
	`
	
	_, err := r.db.ExecContext(ctx, query, category.Name, category.Status, category.CategoryID)
	return err
}

func (r *CategoryRepositorySQLX) Delete(ctx context.Context, id uint) error {
	query := `
		UPDATE categories 
		SET deleted_at = NOW(), updated_at = NOW()
		WHERE category_id = $1 AND deleted_at IS NULL
	`
	
	_, err := r.db.ExecContext(ctx, query, id)
	return err
}

func (r *CategoryRepositorySQLX) List(ctx context.Context, limit, offset int) ([]*models.Category, error) {
	var categories []*models.Category
	query := `
		SELECT category_id, name, status, created_at, updated_at, deleted_at, created_by
		FROM categories 
		WHERE deleted_at IS NULL
		ORDER BY created_at DESC
		LIMIT $1 OFFSET $2
	`
	
	err := r.db.SelectContext(ctx, &categories, query, limit, offset)
	if err != nil {
		return nil, err
	}

	// Load products for each category
	for i := range categories {
		products, err := r.getProductsByCategoryID(ctx, categories[i].CategoryID)
		if err != nil {
			return nil, err
		}
		categories[i].Products = products
	}

	return categories, nil
}

func (r *CategoryRepositorySQLX) GetByStatus(ctx context.Context, status models.StatusUmum) ([]*models.Category, error) {
	var categories []*models.Category
	query := `
		SELECT category_id, name, status, created_at, updated_at, deleted_at, created_by
		FROM categories 
		WHERE status = $1 AND deleted_at IS NULL
		ORDER BY created_at DESC
	`
	
	err := r.db.SelectContext(ctx, &categories, query, status)
	if err != nil {
		return nil, err
	}

	// Load products for each category
	for i := range categories {
		products, err := r.getProductsByCategoryID(ctx, categories[i].CategoryID)
		if err != nil {
			return nil, err
		}
		categories[i].Products = products
	}

	return categories, nil
}

func (r *CategoryRepositorySQLX) GetByName(ctx context.Context, name string) (*models.Category, error) {
	var category models.Category
	query := `
		SELECT category_id, name, status, created_at, updated_at, deleted_at, created_by
		FROM categories 
		WHERE name = $1 AND deleted_at IS NULL
	`
	
	err := r.db.GetContext(ctx, &category, query, name)
	if err != nil {
		return nil, err
	}

	// Load products
	products, err := r.getProductsByCategoryID(ctx, category.CategoryID)
	if err != nil {
		return nil, err
	}
	category.Products = products

	return &category, nil
}

func (r *CategoryRepositorySQLX) Search(ctx context.Context, query string, limit, offset int) ([]*models.Category, error) {
	var categories []*models.Category
	searchQuery := "%" + query + "%"
	
	sqlQuery := `
		SELECT category_id, name, status, created_at, updated_at, deleted_at, created_by
		FROM categories 
		WHERE name ILIKE $1 AND deleted_at IS NULL
		ORDER BY created_at DESC
		LIMIT $2 OFFSET $3
	`
	
	err := r.db.SelectContext(ctx, &categories, sqlQuery, searchQuery, limit, offset)
	if err != nil {
		return nil, err
	}

	// Load products for each category
	for i := range categories {
		products, err := r.getProductsByCategoryID(ctx, categories[i].CategoryID)
		if err != nil {
			return nil, err
		}
		categories[i].Products = products
	}

	return categories, nil
}

// Helper method to get products by category ID
func (r *CategoryRepositorySQLX) getProductsByCategoryID(ctx context.Context, categoryID uint) ([]models.Product, error) {
	var products []models.Product
	query := `
		SELECT product_id, product_name, product_description, product_image, cost_price, 
			   selling_price, stock, sku, barcode, has_serial_number, shelf_location, 
			   usage_status, is_active, category_id, supplier_id, unit_type_id, 
			   sourceable_id, sourceable_type, created_at, updated_at, deleted_at, created_by
		FROM products 
		WHERE category_id = $1 AND deleted_at IS NULL
		ORDER BY created_at DESC
	`
	
	err := r.db.SelectContext(ctx, &products, query, categoryID)
	return products, err
}