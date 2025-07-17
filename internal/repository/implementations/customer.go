package implementations

import (
	"boilerplate/internal/models"
	"boilerplate/internal/repository/interfaces"
	"context"

	"gorm.io/gorm"
)

// CustomerRepository implements the customer repository interface
type CustomerRepository struct {
	db *gorm.DB
}

// NewCustomerRepository creates a new customer repository
func NewCustomerRepository(db *gorm.DB) interfaces.CustomerRepository {
	return &CustomerRepository{db: db}
}

// Create creates a new customer
func (r *CustomerRepository) Create(ctx context.Context, customer *models.Customer) error {
	return r.db.WithContext(ctx).Create(customer).Error
}

// GetByID retrieves a customer by ID
func (r *CustomerRepository) GetByID(ctx context.Context, id uint) (*models.Customer, error) {
	var customer models.Customer
	err := r.db.WithContext(ctx).Preload("Vehicles").First(&customer, id).Error
	if err != nil {
		return nil, err
	}
	return &customer, nil
}

// GetByPhoneNumber retrieves a customer by phone number
func (r *CustomerRepository) GetByPhoneNumber(ctx context.Context, phoneNumber string) (*models.Customer, error) {
	var customer models.Customer
	err := r.db.WithContext(ctx).Preload("Vehicles").Where("phone_number = ?", phoneNumber).First(&customer).Error
	if err != nil {
		return nil, err
	}
	return &customer, nil
}

// Update updates a customer
func (r *CustomerRepository) Update(ctx context.Context, customer *models.Customer) error {
	return r.db.WithContext(ctx).Save(customer).Error
}

// Delete soft deletes a customer
func (r *CustomerRepository) Delete(ctx context.Context, id uint) error {
	return r.db.WithContext(ctx).Delete(&models.Customer{}, id).Error
}

// List retrieves customers with pagination
func (r *CustomerRepository) List(ctx context.Context, limit, offset int) ([]*models.Customer, error) {
	var customers []*models.Customer
	err := r.db.WithContext(ctx).Preload("Vehicles").Limit(limit).Offset(offset).Find(&customers).Error
	if err != nil {
		return nil, err
	}
	return customers, nil
}

// GetByStatus retrieves customers by status
func (r *CustomerRepository) GetByStatus(ctx context.Context, status models.StatusUmum) ([]*models.Customer, error) {
	var customers []*models.Customer
	err := r.db.WithContext(ctx).Preload("Vehicles").Where("status = ?", status).Find(&customers).Error
	if err != nil {
		return nil, err
	}
	return customers, nil
}

// Search searches customers by name, phone number, or address
func (r *CustomerRepository) Search(ctx context.Context, query string, limit, offset int) ([]*models.Customer, error) {
	var customers []*models.Customer
	searchQuery := "%" + query + "%"
	
	err := r.db.WithContext(ctx).
		Preload("Vehicles").
		Where("name LIKE ? OR phone_number LIKE ? OR address LIKE ?", searchQuery, searchQuery, searchQuery).
		Limit(limit).
		Offset(offset).
		Find(&customers).Error
	
	if err != nil {
		return nil, err
	}
	return customers, nil
}

// CustomerVehicleRepository implements the customer vehicle repository interface
type CustomerVehicleRepository struct {
	db *gorm.DB
}

// NewCustomerVehicleRepository creates a new customer vehicle repository
func NewCustomerVehicleRepository(db *gorm.DB) interfaces.CustomerVehicleRepository {
	return &CustomerVehicleRepository{db: db}
}

// Create creates a new customer vehicle
func (r *CustomerVehicleRepository) Create(ctx context.Context, vehicle *models.CustomerVehicle) error {
	return r.db.WithContext(ctx).Create(vehicle).Error
}

// GetByID retrieves a customer vehicle by ID
func (r *CustomerVehicleRepository) GetByID(ctx context.Context, id uint) (*models.CustomerVehicle, error) {
	var vehicle models.CustomerVehicle
	err := r.db.WithContext(ctx).Preload("Customer").First(&vehicle, id).Error
	if err != nil {
		return nil, err
	}
	return &vehicle, nil
}

// GetByPlateNumber retrieves a customer vehicle by plate number
func (r *CustomerVehicleRepository) GetByPlateNumber(ctx context.Context, plateNumber string) (*models.CustomerVehicle, error) {
	var vehicle models.CustomerVehicle
	err := r.db.WithContext(ctx).Preload("Customer").Where("plate_number = ?", plateNumber).First(&vehicle).Error
	if err != nil {
		return nil, err
	}
	return &vehicle, nil
}

// GetByChassisNumber retrieves a customer vehicle by chassis number
func (r *CustomerVehicleRepository) GetByChassisNumber(ctx context.Context, chassisNumber string) (*models.CustomerVehicle, error) {
	var vehicle models.CustomerVehicle
	err := r.db.WithContext(ctx).Preload("Customer").Where("chassis_number = ?", chassisNumber).First(&vehicle).Error
	if err != nil {
		return nil, err
	}
	return &vehicle, nil
}

// GetByEngineNumber retrieves a customer vehicle by engine number
func (r *CustomerVehicleRepository) GetByEngineNumber(ctx context.Context, engineNumber string) (*models.CustomerVehicle, error) {
	var vehicle models.CustomerVehicle
	err := r.db.WithContext(ctx).Preload("Customer").Where("engine_number = ?", engineNumber).First(&vehicle).Error
	if err != nil {
		return nil, err
	}
	return &vehicle, nil
}

// Update updates a customer vehicle
func (r *CustomerVehicleRepository) Update(ctx context.Context, vehicle *models.CustomerVehicle) error {
	return r.db.WithContext(ctx).Save(vehicle).Error
}

// Delete soft deletes a customer vehicle
func (r *CustomerVehicleRepository) Delete(ctx context.Context, id uint) error {
	return r.db.WithContext(ctx).Delete(&models.CustomerVehicle{}, id).Error
}

// List retrieves customer vehicles with pagination
func (r *CustomerVehicleRepository) List(ctx context.Context, limit, offset int) ([]*models.CustomerVehicle, error) {
	var vehicles []*models.CustomerVehicle
	err := r.db.WithContext(ctx).Preload("Customer").Limit(limit).Offset(offset).Find(&vehicles).Error
	if err != nil {
		return nil, err
	}
	return vehicles, nil
}

// GetByCustomerID retrieves customer vehicles by customer ID
func (r *CustomerVehicleRepository) GetByCustomerID(ctx context.Context, customerID uint) ([]*models.CustomerVehicle, error) {
	var vehicles []*models.CustomerVehicle
	err := r.db.WithContext(ctx).Preload("Customer").Where("customer_id = ?", customerID).Find(&vehicles).Error
	if err != nil {
		return nil, err
	}
	return vehicles, nil
}

// Search searches customer vehicles by plate number, brand, model, or type
func (r *CustomerVehicleRepository) Search(ctx context.Context, query string, limit, offset int) ([]*models.CustomerVehicle, error) {
	var vehicles []*models.CustomerVehicle
	searchQuery := "%" + query + "%"
	
	err := r.db.WithContext(ctx).
		Preload("Customer").
		Where("plate_number LIKE ? OR brand LIKE ? OR model LIKE ? OR type LIKE ?", searchQuery, searchQuery, searchQuery, searchQuery).
		Limit(limit).
		Offset(offset).
		Find(&vehicles).Error
	
	if err != nil {
		return nil, err
	}
	return vehicles, nil
}