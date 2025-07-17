package interfaces

import (
	"boilerplate/internal/models"
	"context"
)

// CreateCustomerRequest represents the request to create a customer
type CreateCustomerRequest struct {
	Name        string             `json:"name" validate:"required,min=2,max=255"`
	PhoneNumber string             `json:"phone_number" validate:"required,min=10,max=20"`
	Address     *string            `json:"address,omitempty"`
	Status      models.StatusUmum  `json:"status,omitempty"`
	CreatedBy   *uint              `json:"created_by,omitempty"`
}

// UpdateCustomerRequest represents the request to update a customer
type UpdateCustomerRequest struct {
	Name        *string            `json:"name,omitempty" validate:"omitempty,min=2,max=255"`
	PhoneNumber *string            `json:"phone_number,omitempty" validate:"omitempty,min=10,max=20"`
	Address     *string            `json:"address,omitempty"`
	Status      *models.StatusUmum `json:"status,omitempty"`
}

// CreateCustomerVehicleRequest represents the request to create a customer vehicle
type CreateCustomerVehicleRequest struct {
	CustomerID     uint    `json:"customer_id" validate:"required"`
	PlateNumber    string  `json:"plate_number" validate:"required,min=3,max=20"`
	Brand          string  `json:"brand" validate:"required,min=2,max=100"`
	Model          string  `json:"model" validate:"required,min=2,max=100"`
	Type           string  `json:"type" validate:"required,min=2,max=100"`
	ProductionYear int     `json:"production_year" validate:"required,min=1900,max=2030"`
	ChassisNumber  string  `json:"chassis_number" validate:"required,min=10,max=100"`
	EngineNumber   string  `json:"engine_number" validate:"required,min=5,max=100"`
	Color          string  `json:"color" validate:"required,min=2,max=50"`
	Notes          *string `json:"notes,omitempty"`
	CreatedBy      *uint   `json:"created_by,omitempty"`
}

// UpdateCustomerVehicleRequest represents the request to update a customer vehicle
type UpdateCustomerVehicleRequest struct {
	CustomerID     *uint   `json:"customer_id,omitempty"`
	PlateNumber    *string `json:"plate_number,omitempty" validate:"omitempty,min=3,max=20"`
	Brand          *string `json:"brand,omitempty" validate:"omitempty,min=2,max=100"`
	Model          *string `json:"model,omitempty" validate:"omitempty,min=2,max=100"`
	Type           *string `json:"type,omitempty" validate:"omitempty,min=2,max=100"`
	ProductionYear *int    `json:"production_year,omitempty" validate:"omitempty,min=1900,max=2030"`
	ChassisNumber  *string `json:"chassis_number,omitempty" validate:"omitempty,min=10,max=100"`
	EngineNumber   *string `json:"engine_number,omitempty" validate:"omitempty,min=5,max=100"`
	Color          *string `json:"color,omitempty" validate:"omitempty,min=2,max=50"`
	Notes          *string `json:"notes,omitempty"`
}

// CustomerUsecase interface for customer business logic
type CustomerUsecase interface {
	CreateCustomer(ctx context.Context, req CreateCustomerRequest) (*models.Customer, error)
	GetCustomer(ctx context.Context, id uint) (*models.Customer, error)
	GetCustomerByPhoneNumber(ctx context.Context, phoneNumber string) (*models.Customer, error)
	UpdateCustomer(ctx context.Context, id uint, req UpdateCustomerRequest) (*models.Customer, error)
	DeleteCustomer(ctx context.Context, id uint) error
	ListCustomers(ctx context.Context, limit, offset int) ([]*models.Customer, error)
	GetCustomersByStatus(ctx context.Context, status models.StatusUmum) ([]*models.Customer, error)
	SearchCustomers(ctx context.Context, query string, limit, offset int) ([]*models.Customer, error)
}

// CustomerVehicleUsecase interface for customer vehicle business logic
type CustomerVehicleUsecase interface {
	CreateCustomerVehicle(ctx context.Context, req CreateCustomerVehicleRequest) (*models.CustomerVehicle, error)
	GetCustomerVehicle(ctx context.Context, id uint) (*models.CustomerVehicle, error)
	GetCustomerVehicleByPlateNumber(ctx context.Context, plateNumber string) (*models.CustomerVehicle, error)
	GetCustomerVehicleByChassisNumber(ctx context.Context, chassisNumber string) (*models.CustomerVehicle, error)
	GetCustomerVehicleByEngineNumber(ctx context.Context, engineNumber string) (*models.CustomerVehicle, error)
	UpdateCustomerVehicle(ctx context.Context, id uint, req UpdateCustomerVehicleRequest) (*models.CustomerVehicle, error)
	DeleteCustomerVehicle(ctx context.Context, id uint) error
	ListCustomerVehicles(ctx context.Context, limit, offset int) ([]*models.CustomerVehicle, error)
	GetCustomerVehiclesByCustomerID(ctx context.Context, customerID uint) ([]*models.CustomerVehicle, error)
	SearchCustomerVehicles(ctx context.Context, query string, limit, offset int) ([]*models.CustomerVehicle, error)
}