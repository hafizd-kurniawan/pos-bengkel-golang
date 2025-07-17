package interfaces

import (
	"boilerplate/internal/models"
	"context"
)

// CustomerRepository interface for customer operations
type CustomerRepository interface {
	Create(ctx context.Context, customer *models.Customer) error
	GetByID(ctx context.Context, id uint) (*models.Customer, error)
	GetByPhoneNumber(ctx context.Context, phoneNumber string) (*models.Customer, error)
	Update(ctx context.Context, customer *models.Customer) error
	Delete(ctx context.Context, id uint) error
	List(ctx context.Context, limit, offset int) ([]*models.Customer, error)
	GetByStatus(ctx context.Context, status models.StatusUmum) ([]*models.Customer, error)
	Search(ctx context.Context, query string, limit, offset int) ([]*models.Customer, error)
}

// CustomerVehicleRepository interface for customer vehicle operations
type CustomerVehicleRepository interface {
	Create(ctx context.Context, vehicle *models.CustomerVehicle) error
	GetByID(ctx context.Context, id uint) (*models.CustomerVehicle, error)
	GetByPlateNumber(ctx context.Context, plateNumber string) (*models.CustomerVehicle, error)
	GetByChassisNumber(ctx context.Context, chassisNumber string) (*models.CustomerVehicle, error)
	GetByEngineNumber(ctx context.Context, engineNumber string) (*models.CustomerVehicle, error)
	Update(ctx context.Context, vehicle *models.CustomerVehicle) error
	Delete(ctx context.Context, id uint) error
	List(ctx context.Context, limit, offset int) ([]*models.CustomerVehicle, error)
	GetByCustomerID(ctx context.Context, customerID uint) ([]*models.CustomerVehicle, error)
	Search(ctx context.Context, query string, limit, offset int) ([]*models.CustomerVehicle, error)
}