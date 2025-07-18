package implementations

import (
	"boilerplate/internal/models"
	"boilerplate/internal/repository/interfaces"
	"context"
	"database/sql"

	"github.com/jmoiron/sqlx"
)

// CustomerRepository implements the customer repository interface
type CustomerRepository struct {
	db *sqlx.DB
}

// NewCustomerRepository creates a new customer repository
func NewCustomerRepository(db *sqlx.DB) interfaces.CustomerRepository {
	return &CustomerRepository{db: db}
}

// Create creates a new customer
func (r *CustomerRepository) Create(ctx context.Context, customer *models.Customer) error {
	query := `
		INSERT INTO customers (name, phone_number, address, status, created_by)
		VALUES ($1, $2, $3, $4, $5)
		RETURNING customer_id, created_at, updated_at
	`
	
	row := r.db.QueryRowxContext(ctx, query,
		customer.Name,
		customer.PhoneNumber,
		customer.Address,
		customer.Status,
		customer.CreatedBy,
	)
	
	return row.Scan(&customer.CustomerID, &customer.CreatedAt, &customer.UpdatedAt)
}

// GetByID retrieves a customer by ID
func (r *CustomerRepository) GetByID(ctx context.Context, id uint) (*models.Customer, error) {
	var customer models.Customer
	query := `
		SELECT customer_id, name, phone_number, address, status, 
			   created_at, updated_at, deleted_at, created_by
		FROM customers 
		WHERE customer_id = $1 AND deleted_at IS NULL
	`
	
	err := r.db.GetContext(ctx, &customer, query, id)
	if err != nil {
		return nil, err
	}

	// Load vehicles separately
	vehicles, err := r.getVehiclesByCustomerID(ctx, id)
	if err != nil {
		return nil, err
	}
	customer.Vehicles = vehicles

	return &customer, nil
}

// GetByPhoneNumber retrieves a customer by phone number
func (r *CustomerRepository) GetByPhoneNumber(ctx context.Context, phoneNumber string) (*models.Customer, error) {
	var customer models.Customer
	query := `
		SELECT customer_id, name, phone_number, address, status, 
			   created_at, updated_at, deleted_at, created_by
		FROM customers 
		WHERE phone_number = $1 AND deleted_at IS NULL
	`
	
	err := r.db.GetContext(ctx, &customer, query, phoneNumber)
	if err != nil {
		return nil, err
	}

	// Load vehicles separately
	vehicles, err := r.getVehiclesByCustomerID(ctx, customer.CustomerID)
	if err != nil {
		return nil, err
	}
	customer.Vehicles = vehicles

	return &customer, nil
}

// Update updates a customer
func (r *CustomerRepository) Update(ctx context.Context, customer *models.Customer) error {
	query := `
		UPDATE customers 
		SET name = $1, phone_number = $2, address = $3, status = $4, updated_at = NOW()
		WHERE customer_id = $5 AND deleted_at IS NULL
	`
	
	_, err := r.db.ExecContext(ctx, query,
		customer.Name,
		customer.PhoneNumber,
		customer.Address,
		customer.Status,
		customer.CustomerID,
	)
	
	return err
}

// Delete soft deletes a customer
func (r *CustomerRepository) Delete(ctx context.Context, id uint) error {
	query := `
		UPDATE customers 
		SET deleted_at = NOW(), updated_at = NOW()
		WHERE customer_id = $1 AND deleted_at IS NULL
	`
	
	_, err := r.db.ExecContext(ctx, query, id)
	return err
}

// List retrieves customers with pagination
func (r *CustomerRepository) List(ctx context.Context, limit, offset int) ([]*models.Customer, error) {
	var customers []*models.Customer
	query := `
		SELECT customer_id, name, phone_number, address, status, 
			   created_at, updated_at, deleted_at, created_by
		FROM customers 
		WHERE deleted_at IS NULL
		ORDER BY created_at DESC
		LIMIT $1 OFFSET $2
	`
	
	err := r.db.SelectContext(ctx, &customers, query, limit, offset)
	if err != nil {
		return nil, err
	}

	// Load vehicles for each customer
	for i := range customers {
		vehicles, err := r.getVehiclesByCustomerID(ctx, customers[i].CustomerID)
		if err != nil {
			return nil, err
		}
		customers[i].Vehicles = vehicles
	}

	return customers, nil
}

// GetByStatus retrieves customers by status
func (r *CustomerRepository) GetByStatus(ctx context.Context, status models.StatusUmum) ([]*models.Customer, error) {
	var customers []*models.Customer
	query := `
		SELECT customer_id, name, phone_number, address, status, 
			   created_at, updated_at, deleted_at, created_by
		FROM customers 
		WHERE status = $1 AND deleted_at IS NULL
		ORDER BY created_at DESC
	`
	
	err := r.db.SelectContext(ctx, &customers, query, status)
	if err != nil {
		return nil, err
	}

	// Load vehicles for each customer
	for i := range customers {
		vehicles, err := r.getVehiclesByCustomerID(ctx, customers[i].CustomerID)
		if err != nil {
			return nil, err
		}
		customers[i].Vehicles = vehicles
	}

	return customers, nil
}

// Search searches customers by name, phone number, or address
func (r *CustomerRepository) Search(ctx context.Context, query string, limit, offset int) ([]*models.Customer, error) {
	var customers []*models.Customer
	searchQuery := "%" + query + "%"
	
	sqlQuery := `
		SELECT customer_id, name, phone_number, address, status, 
			   created_at, updated_at, deleted_at, created_by
		FROM customers 
		WHERE (name ILIKE $1 OR phone_number ILIKE $2 OR address ILIKE $3) 
		  AND deleted_at IS NULL
		ORDER BY created_at DESC
		LIMIT $4 OFFSET $5
	`
	
	err := r.db.SelectContext(ctx, &customers, sqlQuery, searchQuery, searchQuery, searchQuery, limit, offset)
	if err != nil {
		return nil, err
	}

	// Load vehicles for each customer
	for i := range customers {
		vehicles, err := r.getVehiclesByCustomerID(ctx, customers[i].CustomerID)
		if err != nil {
			return nil, err
		}
		customers[i].Vehicles = vehicles
	}

	return customers, nil
}

// Helper method to get vehicles by customer ID
func (r *CustomerRepository) getVehiclesByCustomerID(ctx context.Context, customerID uint) ([]models.CustomerVehicle, error) {
	var vehicles []models.CustomerVehicle
	query := `
		SELECT vehicle_id, customer_id, plate_number, brand, model, type, 
			   production_year, chassis_number, engine_number, color, notes,
			   created_at, updated_at, deleted_at, created_by
		FROM customer_vehicles 
		WHERE customer_id = $1 AND deleted_at IS NULL
		ORDER BY created_at DESC
	`
	
	err := r.db.SelectContext(ctx, &vehicles, query, customerID)
	return vehicles, err
}

// CustomerVehicleRepository implements the customer vehicle repository interface
type CustomerVehicleRepository struct {
	db *sqlx.DB
}

// NewCustomerVehicleRepository creates a new customer vehicle repository
func NewCustomerVehicleRepository(db *sqlx.DB) interfaces.CustomerVehicleRepository {
	return &CustomerVehicleRepository{db: db}
}

// Create creates a new customer vehicle
func (r *CustomerVehicleRepository) Create(ctx context.Context, vehicle *models.CustomerVehicle) error {
	query := `
		INSERT INTO customer_vehicles (customer_id, plate_number, brand, model, type, 
									  production_year, chassis_number, engine_number, color, notes, created_by)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)
		RETURNING vehicle_id, created_at, updated_at
	`
	
	row := r.db.QueryRowxContext(ctx, query,
		vehicle.CustomerID,
		vehicle.PlateNumber,
		vehicle.Brand,
		vehicle.Model,
		vehicle.Type,
		vehicle.ProductionYear,
		vehicle.ChassisNumber,
		vehicle.EngineNumber,
		vehicle.Color,
		vehicle.Notes,
		vehicle.CreatedBy,
	)
	
	return row.Scan(&vehicle.VehicleID, &vehicle.CreatedAt, &vehicle.UpdatedAt)
}

// GetByID retrieves a customer vehicle by ID
func (r *CustomerVehicleRepository) GetByID(ctx context.Context, id uint) (*models.CustomerVehicle, error) {
	var vehicle models.CustomerVehicle
	query := `
		SELECT vehicle_id, customer_id, plate_number, brand, model, type, 
			   production_year, chassis_number, engine_number, color, notes,
			   created_at, updated_at, deleted_at, created_by
		FROM customer_vehicles 
		WHERE vehicle_id = $1 AND deleted_at IS NULL
	`
	
	err := r.db.GetContext(ctx, &vehicle, query, id)
	if err != nil {
		return nil, err
	}

	// Load customer separately
	customer, err := r.getCustomerByID(ctx, vehicle.CustomerID)
	if err != nil && err != sql.ErrNoRows {
		return nil, err
	}
	vehicle.Customer = customer

	return &vehicle, nil
}

// GetByPlateNumber retrieves a customer vehicle by plate number
func (r *CustomerVehicleRepository) GetByPlateNumber(ctx context.Context, plateNumber string) (*models.CustomerVehicle, error) {
	var vehicle models.CustomerVehicle
	query := `
		SELECT vehicle_id, customer_id, plate_number, brand, model, type, 
			   production_year, chassis_number, engine_number, color, notes,
			   created_at, updated_at, deleted_at, created_by
		FROM customer_vehicles 
		WHERE plate_number = $1 AND deleted_at IS NULL
	`
	
	err := r.db.GetContext(ctx, &vehicle, query, plateNumber)
	if err != nil {
		return nil, err
	}

	// Load customer separately
	customer, err := r.getCustomerByID(ctx, vehicle.CustomerID)
	if err != nil && err != sql.ErrNoRows {
		return nil, err
	}
	vehicle.Customer = customer

	return &vehicle, nil
}

// GetByChassisNumber retrieves a customer vehicle by chassis number
func (r *CustomerVehicleRepository) GetByChassisNumber(ctx context.Context, chassisNumber string) (*models.CustomerVehicle, error) {
	var vehicle models.CustomerVehicle
	query := `
		SELECT vehicle_id, customer_id, plate_number, brand, model, type, 
			   production_year, chassis_number, engine_number, color, notes,
			   created_at, updated_at, deleted_at, created_by
		FROM customer_vehicles 
		WHERE chassis_number = $1 AND deleted_at IS NULL
	`
	
	err := r.db.GetContext(ctx, &vehicle, query, chassisNumber)
	if err != nil {
		return nil, err
	}

	// Load customer separately
	customer, err := r.getCustomerByID(ctx, vehicle.CustomerID)
	if err != nil && err != sql.ErrNoRows {
		return nil, err
	}
	vehicle.Customer = customer

	return &vehicle, nil
}

// GetByEngineNumber retrieves a customer vehicle by engine number
func (r *CustomerVehicleRepository) GetByEngineNumber(ctx context.Context, engineNumber string) (*models.CustomerVehicle, error) {
	var vehicle models.CustomerVehicle
	query := `
		SELECT vehicle_id, customer_id, plate_number, brand, model, type, 
			   production_year, chassis_number, engine_number, color, notes,
			   created_at, updated_at, deleted_at, created_by
		FROM customer_vehicles 
		WHERE engine_number = $1 AND deleted_at IS NULL
	`
	
	err := r.db.GetContext(ctx, &vehicle, query, engineNumber)
	if err != nil {
		return nil, err
	}

	// Load customer separately
	customer, err := r.getCustomerByID(ctx, vehicle.CustomerID)
	if err != nil && err != sql.ErrNoRows {
		return nil, err
	}
	vehicle.Customer = customer

	return &vehicle, nil
}

// Update updates a customer vehicle
func (r *CustomerVehicleRepository) Update(ctx context.Context, vehicle *models.CustomerVehicle) error {
	query := `
		UPDATE customer_vehicles 
		SET customer_id = $1, plate_number = $2, brand = $3, model = $4, type = $5, 
			production_year = $6, chassis_number = $7, engine_number = $8, color = $9, 
			notes = $10, updated_at = NOW()
		WHERE vehicle_id = $11 AND deleted_at IS NULL
	`
	
	_, err := r.db.ExecContext(ctx, query,
		vehicle.CustomerID,
		vehicle.PlateNumber,
		vehicle.Brand,
		vehicle.Model,
		vehicle.Type,
		vehicle.ProductionYear,
		vehicle.ChassisNumber,
		vehicle.EngineNumber,
		vehicle.Color,
		vehicle.Notes,
		vehicle.VehicleID,
	)
	
	return err
}

// Delete soft deletes a customer vehicle
func (r *CustomerVehicleRepository) Delete(ctx context.Context, id uint) error {
	query := `
		UPDATE customer_vehicles 
		SET deleted_at = NOW(), updated_at = NOW()
		WHERE vehicle_id = $1 AND deleted_at IS NULL
	`
	
	_, err := r.db.ExecContext(ctx, query, id)
	return err
}

// List retrieves customer vehicles with pagination
func (r *CustomerVehicleRepository) List(ctx context.Context, limit, offset int) ([]*models.CustomerVehicle, error) {
	var vehicles []*models.CustomerVehicle
	query := `
		SELECT vehicle_id, customer_id, plate_number, brand, model, type, 
			   production_year, chassis_number, engine_number, color, notes,
			   created_at, updated_at, deleted_at, created_by
		FROM customer_vehicles 
		WHERE deleted_at IS NULL
		ORDER BY created_at DESC
		LIMIT $1 OFFSET $2
	`
	
	err := r.db.SelectContext(ctx, &vehicles, query, limit, offset)
	if err != nil {
		return nil, err
	}

	// Load customer for each vehicle
	for i := range vehicles {
		customer, err := r.getCustomerByID(ctx, vehicles[i].CustomerID)
		if err != nil && err != sql.ErrNoRows {
			return nil, err
		}
		vehicles[i].Customer = customer
	}

	return vehicles, nil
}

// GetByCustomerID retrieves customer vehicles by customer ID
func (r *CustomerVehicleRepository) GetByCustomerID(ctx context.Context, customerID uint) ([]*models.CustomerVehicle, error) {
	var vehicles []*models.CustomerVehicle
	query := `
		SELECT vehicle_id, customer_id, plate_number, brand, model, type, 
			   production_year, chassis_number, engine_number, color, notes,
			   created_at, updated_at, deleted_at, created_by
		FROM customer_vehicles 
		WHERE customer_id = $1 AND deleted_at IS NULL
		ORDER BY created_at DESC
	`
	
	err := r.db.SelectContext(ctx, &vehicles, query, customerID)
	if err != nil {
		return nil, err
	}

	// Load customer for each vehicle
	for i := range vehicles {
		customer, err := r.getCustomerByID(ctx, vehicles[i].CustomerID)
		if err != nil && err != sql.ErrNoRows {
			return nil, err
		}
		vehicles[i].Customer = customer
	}

	return vehicles, nil
}

// Search searches customer vehicles by plate number, brand, model, or type
func (r *CustomerVehicleRepository) Search(ctx context.Context, query string, limit, offset int) ([]*models.CustomerVehicle, error) {
	var vehicles []*models.CustomerVehicle
	searchQuery := "%" + query + "%"
	
	sqlQuery := `
		SELECT vehicle_id, customer_id, plate_number, brand, model, type, 
			   production_year, chassis_number, engine_number, color, notes,
			   created_at, updated_at, deleted_at, created_by
		FROM customer_vehicles 
		WHERE (plate_number ILIKE $1 OR brand ILIKE $2 OR model ILIKE $3 OR type ILIKE $4) 
		  AND deleted_at IS NULL
		ORDER BY created_at DESC
		LIMIT $5 OFFSET $6
	`
	
	err := r.db.SelectContext(ctx, &vehicles, sqlQuery, searchQuery, searchQuery, searchQuery, searchQuery, limit, offset)
	if err != nil {
		return nil, err
	}

	// Load customer for each vehicle
	for i := range vehicles {
		customer, err := r.getCustomerByID(ctx, vehicles[i].CustomerID)
		if err != nil && err != sql.ErrNoRows {
			return nil, err
		}
		vehicles[i].Customer = customer
	}

	return vehicles, nil
}

// Helper method to get customer by ID
func (r *CustomerVehicleRepository) getCustomerByID(ctx context.Context, customerID uint) (*models.Customer, error) {
	var customer models.Customer
	query := `
		SELECT customer_id, name, phone_number, address, status, 
			   created_at, updated_at, deleted_at, created_by
		FROM customers 
		WHERE customer_id = $1 AND deleted_at IS NULL
	`
	
	err := r.db.GetContext(ctx, &customer, query, customerID)
	if err != nil {
		return nil, err
	}
	
	return &customer, nil
}