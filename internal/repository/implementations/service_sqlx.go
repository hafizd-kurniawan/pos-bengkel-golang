package implementations

import (
	"boilerplate/internal/models"
	"boilerplate/internal/repository/interfaces"
	"context"
	"database/sql"

	"github.com/jmoiron/sqlx"
)

// ServiceRepositorySQLX implements ServiceRepository interface using SQLx
type ServiceRepositorySQLX struct {
	db *sqlx.DB
}

// NewServiceRepositorySQLX creates a new service repository using SQLx
func NewServiceRepositorySQLX(db *sqlx.DB) interfaces.ServiceRepository {
	return &ServiceRepositorySQLX{db: db}
}

func (r *ServiceRepositorySQLX) Create(ctx context.Context, service *models.Service) error {
	query := `
		INSERT INTO services (service_code, name, service_category_id, fee, status, created_by)
		VALUES ($1, $2, $3, $4, $5, $6)
		RETURNING service_id, created_at, updated_at
	`
	
	row := r.db.QueryRowxContext(ctx, query,
		service.ServiceCode,
		service.Name,
		service.ServiceCategoryID,
		service.Fee,
		service.Status,
		service.CreatedBy,
	)
	
	return row.Scan(&service.ServiceID, &service.CreatedAt, &service.UpdatedAt)
}

func (r *ServiceRepositorySQLX) GetByID(ctx context.Context, id uint) (*models.Service, error) {
	var service models.Service
	query := `
		SELECT service_id, service_code, name, service_category_id, fee, status, 
			   created_at, updated_at, deleted_at, created_by
		FROM services 
		WHERE service_id = $1 AND deleted_at IS NULL
	`
	
	err := r.db.GetContext(ctx, &service, query, id)
	if err != nil {
		return nil, err
	}

	// Load service category
	category, err := r.getServiceCategoryByID(ctx, service.ServiceCategoryID)
	if err != nil && err != sql.ErrNoRows {
		return nil, err
	}
	service.ServiceCategory = category

	return &service, nil
}

func (r *ServiceRepositorySQLX) GetByServiceCode(ctx context.Context, serviceCode string) (*models.Service, error) {
	var service models.Service
	query := `
		SELECT service_id, service_code, name, service_category_id, fee, status, 
			   created_at, updated_at, deleted_at, created_by
		FROM services 
		WHERE service_code = $1 AND deleted_at IS NULL
	`
	
	err := r.db.GetContext(ctx, &service, query, serviceCode)
	if err != nil {
		return nil, err
	}

	// Load service category
	category, err := r.getServiceCategoryByID(ctx, service.ServiceCategoryID)
	if err != nil && err != sql.ErrNoRows {
		return nil, err
	}
	service.ServiceCategory = category

	return &service, nil
}

func (r *ServiceRepositorySQLX) Update(ctx context.Context, service *models.Service) error {
	query := `
		UPDATE services 
		SET service_code = $1, name = $2, service_category_id = $3, fee = $4, 
			status = $5, updated_at = NOW()
		WHERE service_id = $6 AND deleted_at IS NULL
	`
	
	_, err := r.db.ExecContext(ctx, query,
		service.ServiceCode,
		service.Name,
		service.ServiceCategoryID,
		service.Fee,
		service.Status,
		service.ServiceID,
	)
	
	return err
}

func (r *ServiceRepositorySQLX) Delete(ctx context.Context, id uint) error {
	query := `
		UPDATE services 
		SET deleted_at = NOW(), updated_at = NOW()
		WHERE service_id = $1 AND deleted_at IS NULL
	`
	
	_, err := r.db.ExecContext(ctx, query, id)
	return err
}

func (r *ServiceRepositorySQLX) List(ctx context.Context, limit, offset int) ([]*models.Service, error) {
	var services []*models.Service
	query := `
		SELECT service_id, service_code, name, service_category_id, fee, status, 
			   created_at, updated_at, deleted_at, created_by
		FROM services 
		WHERE deleted_at IS NULL
		ORDER BY created_at DESC
		LIMIT $1 OFFSET $2
	`
	
	err := r.db.SelectContext(ctx, &services, query, limit, offset)
	if err != nil {
		return nil, err
	}

	// Load service categories for each service
	for i := range services {
		category, err := r.getServiceCategoryByID(ctx, services[i].ServiceCategoryID)
		if err != nil && err != sql.ErrNoRows {
			return nil, err
		}
		services[i].ServiceCategory = category
	}

	return services, nil
}

func (r *ServiceRepositorySQLX) GetByCategoryID(ctx context.Context, categoryID uint) ([]*models.Service, error) {
	var services []*models.Service
	query := `
		SELECT service_id, service_code, name, service_category_id, fee, status, 
			   created_at, updated_at, deleted_at, created_by
		FROM services 
		WHERE service_category_id = $1 AND deleted_at IS NULL
		ORDER BY created_at DESC
	`
	
	err := r.db.SelectContext(ctx, &services, query, categoryID)
	if err != nil {
		return nil, err
	}

	// Load service category for each service
	for i := range services {
		category, err := r.getServiceCategoryByID(ctx, services[i].ServiceCategoryID)
		if err != nil && err != sql.ErrNoRows {
			return nil, err
		}
		services[i].ServiceCategory = category
	}

	return services, nil
}

func (r *ServiceRepositorySQLX) GetByStatus(ctx context.Context, status models.StatusUmum) ([]*models.Service, error) {
	var services []*models.Service
	query := `
		SELECT service_id, service_code, name, service_category_id, fee, status, 
			   created_at, updated_at, deleted_at, created_by
		FROM services 
		WHERE status = $1 AND deleted_at IS NULL
		ORDER BY created_at DESC
	`
	
	err := r.db.SelectContext(ctx, &services, query, status)
	if err != nil {
		return nil, err
	}

	// Load service categories
	for i := range services {
		category, err := r.getServiceCategoryByID(ctx, services[i].ServiceCategoryID)
		if err != nil && err != sql.ErrNoRows {
			return nil, err
		}
		services[i].ServiceCategory = category
	}

	return services, nil
}

func (r *ServiceRepositorySQLX) Search(ctx context.Context, query string, limit, offset int) ([]*models.Service, error) {
	var services []*models.Service
	searchQuery := "%" + query + "%"
	
	sqlQuery := `
		SELECT service_id, service_code, name, service_category_id, fee, status, 
			   created_at, updated_at, deleted_at, created_by
		FROM services 
		WHERE (service_code ILIKE $1 OR name ILIKE $2) 
		  AND deleted_at IS NULL
		ORDER BY created_at DESC
		LIMIT $3 OFFSET $4
	`
	
	err := r.db.SelectContext(ctx, &services, sqlQuery, searchQuery, searchQuery, limit, offset)
	if err != nil {
		return nil, err
	}

	// Load service categories
	for i := range services {
		category, err := r.getServiceCategoryByID(ctx, services[i].ServiceCategoryID)
		if err != nil && err != sql.ErrNoRows {
			return nil, err
		}
		services[i].ServiceCategory = category
	}

	return services, nil
}

// Helper method to get service category by ID
func (r *ServiceRepositorySQLX) getServiceCategoryByID(ctx context.Context, categoryID uint) (*models.ServiceCategory, error) {
	var category models.ServiceCategory
	query := `
		SELECT service_category_id, name, status, created_at, updated_at, deleted_at, created_by
		FROM service_categories 
		WHERE service_category_id = $1 AND deleted_at IS NULL
	`
	
	err := r.db.GetContext(ctx, &category, query, categoryID)
	if err != nil {
		return nil, err
	}
	
	return &category, nil
}

// ServiceCategoryRepositorySQLX implements ServiceCategoryRepository interface using SQLx
type ServiceCategoryRepositorySQLX struct {
	db *sqlx.DB
}

// NewServiceCategoryRepositorySQLX creates a new service category repository using SQLx
func NewServiceCategoryRepositorySQLX(db *sqlx.DB) interfaces.ServiceCategoryRepository {
	return &ServiceCategoryRepositorySQLX{db: db}
}

func (r *ServiceCategoryRepositorySQLX) Create(ctx context.Context, category *models.ServiceCategory) error {
	query := `
		INSERT INTO service_categories (name, status, created_by)
		VALUES ($1, $2, $3)
		RETURNING service_category_id, created_at, updated_at
	`
	
	row := r.db.QueryRowxContext(ctx, query, category.Name, category.Status, category.CreatedBy)
	return row.Scan(&category.ServiceCategoryID, &category.CreatedAt, &category.UpdatedAt)
}

func (r *ServiceCategoryRepositorySQLX) GetByID(ctx context.Context, id uint) (*models.ServiceCategory, error) {
	var category models.ServiceCategory
	query := `
		SELECT service_category_id, name, status, created_at, updated_at, deleted_at, created_by
		FROM service_categories 
		WHERE service_category_id = $1 AND deleted_at IS NULL
	`
	
	err := r.db.GetContext(ctx, &category, query, id)
	if err != nil {
		return nil, err
	}

	// Load services
	services, err := r.getServicesByCategoryID(ctx, id)
	if err != nil {
		return nil, err
	}
	category.Services = services

	return &category, nil
}

func (r *ServiceCategoryRepositorySQLX) Update(ctx context.Context, category *models.ServiceCategory) error {
	query := `
		UPDATE service_categories 
		SET name = $1, status = $2, updated_at = NOW()
		WHERE service_category_id = $3 AND deleted_at IS NULL
	`
	
	_, err := r.db.ExecContext(ctx, query, category.Name, category.Status, category.ServiceCategoryID)
	return err
}

func (r *ServiceCategoryRepositorySQLX) Delete(ctx context.Context, id uint) error {
	query := `
		UPDATE service_categories 
		SET deleted_at = NOW(), updated_at = NOW()
		WHERE service_category_id = $1 AND deleted_at IS NULL
	`
	
	_, err := r.db.ExecContext(ctx, query, id)
	return err
}

func (r *ServiceCategoryRepositorySQLX) List(ctx context.Context, limit, offset int) ([]*models.ServiceCategory, error) {
	var categories []*models.ServiceCategory
	query := `
		SELECT service_category_id, name, status, created_at, updated_at, deleted_at, created_by
		FROM service_categories 
		WHERE deleted_at IS NULL
		ORDER BY created_at DESC
		LIMIT $1 OFFSET $2
	`
	
	err := r.db.SelectContext(ctx, &categories, query, limit, offset)
	if err != nil {
		return nil, err
	}

	// Load services for each category
	for i := range categories {
		services, err := r.getServicesByCategoryID(ctx, categories[i].ServiceCategoryID)
		if err != nil {
			return nil, err
		}
		categories[i].Services = services
	}

	return categories, nil
}

func (r *ServiceCategoryRepositorySQLX) GetByStatus(ctx context.Context, status models.StatusUmum) ([]*models.ServiceCategory, error) {
	var categories []*models.ServiceCategory
	query := `
		SELECT service_category_id, name, status, created_at, updated_at, deleted_at, created_by
		FROM service_categories 
		WHERE status = $1 AND deleted_at IS NULL
		ORDER BY created_at DESC
	`
	
	err := r.db.SelectContext(ctx, &categories, query, status)
	if err != nil {
		return nil, err
	}

	// Load services for each category
	for i := range categories {
		services, err := r.getServicesByCategoryID(ctx, categories[i].ServiceCategoryID)
		if err != nil {
			return nil, err
		}
		categories[i].Services = services
	}

	return categories, nil
}

func (r *ServiceCategoryRepositorySQLX) GetByName(ctx context.Context, name string) (*models.ServiceCategory, error) {
	var category models.ServiceCategory
	query := `
		SELECT service_category_id, name, status, created_at, updated_at, deleted_at, created_by
		FROM service_categories 
		WHERE name = $1 AND deleted_at IS NULL
	`
	
	err := r.db.GetContext(ctx, &category, query, name)
	if err != nil {
		return nil, err
	}

	// Load services
	services, err := r.getServicesByCategoryID(ctx, category.ServiceCategoryID)
	if err != nil {
		return nil, err
	}
	category.Services = services

	return &category, nil
}

func (r *ServiceCategoryRepositorySQLX) Search(ctx context.Context, query string, limit, offset int) ([]*models.ServiceCategory, error) {
	var categories []*models.ServiceCategory
	searchQuery := "%" + query + "%"
	
	sqlQuery := `
		SELECT service_category_id, name, status, created_at, updated_at, deleted_at, created_by
		FROM service_categories 
		WHERE name ILIKE $1 AND deleted_at IS NULL
		ORDER BY created_at DESC
		LIMIT $2 OFFSET $3
	`
	
	err := r.db.SelectContext(ctx, &categories, sqlQuery, searchQuery, limit, offset)
	if err != nil {
		return nil, err
	}

	// Load services for each category
	for i := range categories {
		services, err := r.getServicesByCategoryID(ctx, categories[i].ServiceCategoryID)
		if err != nil {
			return nil, err
		}
		categories[i].Services = services
	}

	return categories, nil
}

// Helper method to get services by category ID
func (r *ServiceCategoryRepositorySQLX) getServicesByCategoryID(ctx context.Context, categoryID uint) ([]models.Service, error) {
	var services []models.Service
	query := `
		SELECT service_id, service_code, name, service_category_id, fee, status, 
			   created_at, updated_at, deleted_at, created_by
		FROM services 
		WHERE service_category_id = $1 AND deleted_at IS NULL
		ORDER BY created_at DESC
	`
	
	err := r.db.SelectContext(ctx, &services, query, categoryID)
	return services, err
}