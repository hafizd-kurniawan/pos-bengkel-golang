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

// CustomerUsecase implements the customer usecase interface
type CustomerUsecase struct {
	repo *repository.RepositoryManager
}

// NewCustomerUsecase creates a new customer usecase
func NewCustomerUsecase(repo *repository.RepositoryManager) interfaces.CustomerUsecase {
	return &CustomerUsecase{repo: repo}
}

// CreateCustomer creates a new customer
func (u *CustomerUsecase) CreateCustomer(ctx context.Context, req interfaces.CreateCustomerRequest) (*models.Customer, error) {
	// Check if phone number already exists
	existingCustomer, err := u.repo.Customer.GetByPhoneNumber(ctx, req.PhoneNumber)
	if err == nil && existingCustomer != nil {
		return nil, errors.New("customer with this phone number already exists")
	}
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}

	// Set default status if not provided
	status := req.Status
	if status == "" {
		status = models.StatusAktif
	}

	customer := &models.Customer{
		Name:        req.Name,
		PhoneNumber: req.PhoneNumber,
		Address:     req.Address,
		Status:      status,
		CreatedBy:   req.CreatedBy,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	if err := u.repo.Customer.Create(ctx, customer); err != nil {
		return nil, err
	}

	return customer, nil
}

// GetCustomer retrieves a customer by ID
func (u *CustomerUsecase) GetCustomer(ctx context.Context, id uint) (*models.Customer, error) {
	customer, err := u.repo.Customer.GetByID(ctx, id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("customer not found")
		}
		return nil, err
	}
	return customer, nil
}

// GetCustomerByPhoneNumber retrieves a customer by phone number
func (u *CustomerUsecase) GetCustomerByPhoneNumber(ctx context.Context, phoneNumber string) (*models.Customer, error) {
	customer, err := u.repo.Customer.GetByPhoneNumber(ctx, phoneNumber)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("customer not found")
		}
		return nil, err
	}
	return customer, nil
}

// UpdateCustomer updates a customer
func (u *CustomerUsecase) UpdateCustomer(ctx context.Context, id uint, req interfaces.UpdateCustomerRequest) (*models.Customer, error) {
	customer, err := u.repo.Customer.GetByID(ctx, id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("customer not found")
		}
		return nil, err
	}

	// Check if phone number already exists (if being updated)
	if req.PhoneNumber != nil && *req.PhoneNumber != customer.PhoneNumber {
		existingCustomer, err := u.repo.Customer.GetByPhoneNumber(ctx, *req.PhoneNumber)
		if err == nil && existingCustomer != nil && existingCustomer.CustomerID != id {
			return nil, errors.New("customer with this phone number already exists")
		}
		if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, err
		}
	}

	// Update fields if provided
	if req.Name != nil {
		customer.Name = *req.Name
	}
	if req.PhoneNumber != nil {
		customer.PhoneNumber = *req.PhoneNumber
	}
	if req.Address != nil {
		customer.Address = req.Address
	}
	if req.Status != nil {
		customer.Status = *req.Status
	}
	customer.UpdatedAt = time.Now()

	if err := u.repo.Customer.Update(ctx, customer); err != nil {
		return nil, err
	}

	return customer, nil
}

// DeleteCustomer deletes a customer
func (u *CustomerUsecase) DeleteCustomer(ctx context.Context, id uint) error {
	_, err := u.repo.Customer.GetByID(ctx, id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New("customer not found")
		}
		return err
	}

	// Check if customer has vehicles
	vehicles, err := u.repo.CustomerVehicle.GetByCustomerID(ctx, id)
	if err != nil {
		return err
	}
	if len(vehicles) > 0 {
		return errors.New("cannot delete customer with existing vehicles")
	}

	return u.repo.Customer.Delete(ctx, id)
}

// ListCustomers retrieves customers with pagination
func (u *CustomerUsecase) ListCustomers(ctx context.Context, limit, offset int) ([]*models.Customer, error) {
	if limit <= 0 {
		limit = 10
	}
	if offset < 0 {
		offset = 0
	}

	return u.repo.Customer.List(ctx, limit, offset)
}

// GetCustomersByStatus retrieves customers by status
func (u *CustomerUsecase) GetCustomersByStatus(ctx context.Context, status models.StatusUmum) ([]*models.Customer, error) {
	return u.repo.Customer.GetByStatus(ctx, status)
}

// SearchCustomers searches customers
func (u *CustomerUsecase) SearchCustomers(ctx context.Context, query string, limit, offset int) ([]*models.Customer, error) {
	if limit <= 0 {
		limit = 10
	}
	if offset < 0 {
		offset = 0
	}

	return u.repo.Customer.Search(ctx, query, limit, offset)
}

// CustomerVehicleUsecase implements the customer vehicle usecase interface
type CustomerVehicleUsecase struct {
	repo *repository.RepositoryManager
}

// NewCustomerVehicleUsecase creates a new customer vehicle usecase
func NewCustomerVehicleUsecase(repo *repository.RepositoryManager) interfaces.CustomerVehicleUsecase {
	return &CustomerVehicleUsecase{repo: repo}
}

// CreateCustomerVehicle creates a new customer vehicle
func (u *CustomerVehicleUsecase) CreateCustomerVehicle(ctx context.Context, req interfaces.CreateCustomerVehicleRequest) (*models.CustomerVehicle, error) {
	// Check if customer exists
	_, err := u.repo.Customer.GetByID(ctx, req.CustomerID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("customer not found")
		}
		return nil, err
	}

	// Check if plate number already exists
	existingVehicle, err := u.repo.CustomerVehicle.GetByPlateNumber(ctx, req.PlateNumber)
	if err == nil && existingVehicle != nil {
		return nil, errors.New("vehicle with this plate number already exists")
	}
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}

	// Check if chassis number already exists
	existingVehicle, err = u.repo.CustomerVehicle.GetByChassisNumber(ctx, req.ChassisNumber)
	if err == nil && existingVehicle != nil {
		return nil, errors.New("vehicle with this chassis number already exists")
	}
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}

	// Check if engine number already exists
	existingVehicle, err = u.repo.CustomerVehicle.GetByEngineNumber(ctx, req.EngineNumber)
	if err == nil && existingVehicle != nil {
		return nil, errors.New("vehicle with this engine number already exists")
	}
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}

	vehicle := &models.CustomerVehicle{
		CustomerID:     req.CustomerID,
		PlateNumber:    req.PlateNumber,
		Brand:          req.Brand,
		Model:          req.Model,
		Type:           req.Type,
		ProductionYear: req.ProductionYear,
		ChassisNumber:  req.ChassisNumber,
		EngineNumber:   req.EngineNumber,
		Color:          req.Color,
		Notes:          req.Notes,
		CreatedBy:      req.CreatedBy,
		CreatedAt:      time.Now(),
		UpdatedAt:      time.Now(),
	}

	if err := u.repo.CustomerVehicle.Create(ctx, vehicle); err != nil {
		return nil, err
	}

	return vehicle, nil
}

// GetCustomerVehicle retrieves a customer vehicle by ID
func (u *CustomerVehicleUsecase) GetCustomerVehicle(ctx context.Context, id uint) (*models.CustomerVehicle, error) {
	vehicle, err := u.repo.CustomerVehicle.GetByID(ctx, id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("customer vehicle not found")
		}
		return nil, err
	}
	return vehicle, nil
}

// GetCustomerVehicleByPlateNumber retrieves a customer vehicle by plate number
func (u *CustomerVehicleUsecase) GetCustomerVehicleByPlateNumber(ctx context.Context, plateNumber string) (*models.CustomerVehicle, error) {
	vehicle, err := u.repo.CustomerVehicle.GetByPlateNumber(ctx, plateNumber)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("customer vehicle not found")
		}
		return nil, err
	}
	return vehicle, nil
}

// GetCustomerVehicleByChassisNumber retrieves a customer vehicle by chassis number
func (u *CustomerVehicleUsecase) GetCustomerVehicleByChassisNumber(ctx context.Context, chassisNumber string) (*models.CustomerVehicle, error) {
	vehicle, err := u.repo.CustomerVehicle.GetByChassisNumber(ctx, chassisNumber)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("customer vehicle not found")
		}
		return nil, err
	}
	return vehicle, nil
}

// GetCustomerVehicleByEngineNumber retrieves a customer vehicle by engine number
func (u *CustomerVehicleUsecase) GetCustomerVehicleByEngineNumber(ctx context.Context, engineNumber string) (*models.CustomerVehicle, error) {
	vehicle, err := u.repo.CustomerVehicle.GetByEngineNumber(ctx, engineNumber)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("customer vehicle not found")
		}
		return nil, err
	}
	return vehicle, nil
}

// UpdateCustomerVehicle updates a customer vehicle
func (u *CustomerVehicleUsecase) UpdateCustomerVehicle(ctx context.Context, id uint, req interfaces.UpdateCustomerVehicleRequest) (*models.CustomerVehicle, error) {
	vehicle, err := u.repo.CustomerVehicle.GetByID(ctx, id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("customer vehicle not found")
		}
		return nil, err
	}

	// Check if customer exists (if being updated)
	if req.CustomerID != nil && *req.CustomerID != vehicle.CustomerID {
		_, err := u.repo.Customer.GetByID(ctx, *req.CustomerID)
		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return nil, errors.New("customer not found")
			}
			return nil, err
		}
	}

	// Check unique fields if being updated
	if req.PlateNumber != nil && *req.PlateNumber != vehicle.PlateNumber {
		existingVehicle, err := u.repo.CustomerVehicle.GetByPlateNumber(ctx, *req.PlateNumber)
		if err == nil && existingVehicle != nil && existingVehicle.VehicleID != id {
			return nil, errors.New("vehicle with this plate number already exists")
		}
		if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, err
		}
	}

	if req.ChassisNumber != nil && *req.ChassisNumber != vehicle.ChassisNumber {
		existingVehicle, err := u.repo.CustomerVehicle.GetByChassisNumber(ctx, *req.ChassisNumber)
		if err == nil && existingVehicle != nil && existingVehicle.VehicleID != id {
			return nil, errors.New("vehicle with this chassis number already exists")
		}
		if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, err
		}
	}

	if req.EngineNumber != nil && *req.EngineNumber != vehicle.EngineNumber {
		existingVehicle, err := u.repo.CustomerVehicle.GetByEngineNumber(ctx, *req.EngineNumber)
		if err == nil && existingVehicle != nil && existingVehicle.VehicleID != id {
			return nil, errors.New("vehicle with this engine number already exists")
		}
		if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, err
		}
	}

	// Update fields if provided
	if req.CustomerID != nil {
		vehicle.CustomerID = *req.CustomerID
	}
	if req.PlateNumber != nil {
		vehicle.PlateNumber = *req.PlateNumber
	}
	if req.Brand != nil {
		vehicle.Brand = *req.Brand
	}
	if req.Model != nil {
		vehicle.Model = *req.Model
	}
	if req.Type != nil {
		vehicle.Type = *req.Type
	}
	if req.ProductionYear != nil {
		vehicle.ProductionYear = *req.ProductionYear
	}
	if req.ChassisNumber != nil {
		vehicle.ChassisNumber = *req.ChassisNumber
	}
	if req.EngineNumber != nil {
		vehicle.EngineNumber = *req.EngineNumber
	}
	if req.Color != nil {
		vehicle.Color = *req.Color
	}
	if req.Notes != nil {
		vehicle.Notes = req.Notes
	}
	vehicle.UpdatedAt = time.Now()

	if err := u.repo.CustomerVehicle.Update(ctx, vehicle); err != nil {
		return nil, err
	}

	return vehicle, nil
}

// DeleteCustomerVehicle deletes a customer vehicle
func (u *CustomerVehicleUsecase) DeleteCustomerVehicle(ctx context.Context, id uint) error {
	_, err := u.repo.CustomerVehicle.GetByID(ctx, id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New("customer vehicle not found")
		}
		return err
	}

	// TODO: Add checks for existing service jobs or transactions

	return u.repo.CustomerVehicle.Delete(ctx, id)
}

// ListCustomerVehicles retrieves customer vehicles with pagination
func (u *CustomerVehicleUsecase) ListCustomerVehicles(ctx context.Context, limit, offset int) ([]*models.CustomerVehicle, error) {
	if limit <= 0 {
		limit = 10
	}
	if offset < 0 {
		offset = 0
	}

	return u.repo.CustomerVehicle.List(ctx, limit, offset)
}

// GetCustomerVehiclesByCustomerID retrieves customer vehicles by customer ID
func (u *CustomerVehicleUsecase) GetCustomerVehiclesByCustomerID(ctx context.Context, customerID uint) ([]*models.CustomerVehicle, error) {
	return u.repo.CustomerVehicle.GetByCustomerID(ctx, customerID)
}

// SearchCustomerVehicles searches customer vehicles
func (u *CustomerVehicleUsecase) SearchCustomerVehicles(ctx context.Context, query string, limit, offset int) ([]*models.CustomerVehicle, error) {
	if limit <= 0 {
		limit = 10
	}
	if offset < 0 {
		offset = 0
	}

	return u.repo.CustomerVehicle.Search(ctx, query, limit, offset)
}