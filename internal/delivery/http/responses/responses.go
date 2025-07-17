package responses

import (
	"boilerplate/internal/models"
	"time"
)

// Response represents a standard API response
type Response struct {
	Status  string      `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
	Error   string      `json:"error,omitempty"`
}

// PaginatedResponse represents a paginated API response
type PaginatedResponse struct {
	Status     string      `json:"status"`
	Message    string      `json:"message"`
	Data       interface{} `json:"data"`
	Pagination Pagination  `json:"pagination"`
}

// Pagination contains pagination metadata
type Pagination struct {
	Page    int   `json:"page"`
	Limit   int   `json:"limit"`
	Total   int64 `json:"total"`
	Pages   int   `json:"pages"`
}

// UserResponse represents user data in API response
type UserResponse struct {
	UserID    uint          `json:"user_id"`
	Name      string        `json:"name"`
	Email     string        `json:"email"`
	OutletID  *uint         `json:"outlet_id"`
	Outlet    *OutletResponse `json:"outlet,omitempty"`
	CreatedAt time.Time     `json:"created_at"`
	UpdatedAt time.Time     `json:"updated_at"`
}

// OutletResponse represents outlet data in API response
type OutletResponse struct {
	OutletID    uint              `json:"outlet_id"`
	OutletName  string            `json:"outlet_name"`
	BranchType  string            `json:"branch_type"`
	City        string            `json:"city"`
	Address     *string           `json:"address"`
	PhoneNumber *string           `json:"phone_number"`
	Status      models.StatusUmum `json:"status"`
	CreatedAt   time.Time         `json:"created_at"`
	UpdatedAt   time.Time         `json:"updated_at"`
}

// RoleResponse represents role data in API response
type RoleResponse struct {
	ID          uint                `json:"id"`
	Name        string              `json:"name"`
	Permissions []PermissionResponse `json:"permissions,omitempty"`
	CreatedAt   time.Time           `json:"created_at"`
	UpdatedAt   time.Time           `json:"updated_at"`
}

// PermissionResponse represents permission data in API response
type PermissionResponse struct {
	ID        uint      `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// CustomerResponse represents customer data in API response
type CustomerResponse struct {
	CustomerID  uint              `json:"customer_id"`
	Name        string            `json:"name"`
	PhoneNumber string            `json:"phone_number"`
	Address     *string           `json:"address"`
	Status      models.StatusUmum `json:"status"`
	Vehicles    []CustomerVehicleResponse `json:"vehicles,omitempty"`
	CreatedAt   time.Time         `json:"created_at"`
	UpdatedAt   time.Time         `json:"updated_at"`
}

// CustomerVehicleResponse represents customer vehicle data in API response
type CustomerVehicleResponse struct {
	VehicleID      uint              `json:"vehicle_id"`
	CustomerID     uint              `json:"customer_id"`
	PlateNumber    string            `json:"plate_number"`
	Brand          string            `json:"brand"`
	Model          string            `json:"model"`
	Type           string            `json:"type"`
	ProductionYear int               `json:"production_year"`
	ChassisNumber  string            `json:"chassis_number"`
	EngineNumber   string            `json:"engine_number"`
	Color          string            `json:"color"`
	Notes          *string           `json:"notes"`
	Customer       *CustomerResponse `json:"customer,omitempty"`
	CreatedAt      time.Time         `json:"created_at"`
	UpdatedAt      time.Time         `json:"updated_at"`
}

// Helper functions to convert models to responses
func ToUserResponse(user *models.User) *UserResponse {
	response := &UserResponse{
		UserID:    user.UserID,
		Name:      user.Name,
		Email:     user.Email,
		OutletID:  user.OutletID,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}

	if user.Outlet != nil {
		response.Outlet = ToOutletResponse(user.Outlet)
	}

	return response
}

func ToOutletResponse(outlet *models.Outlet) *OutletResponse {
	return &OutletResponse{
		OutletID:    outlet.OutletID,
		OutletName:  outlet.OutletName,
		BranchType:  outlet.BranchType,
		City:        outlet.City,
		Address:     outlet.Address,
		PhoneNumber: outlet.PhoneNumber,
		Status:      outlet.Status,
		CreatedAt:   outlet.CreatedAt,
		UpdatedAt:   outlet.UpdatedAt,
	}
}

func ToRoleResponse(role *models.Role) *RoleResponse {
	response := &RoleResponse{
		ID:        role.ID,
		Name:      role.Name,
		CreatedAt: role.CreatedAt,
		UpdatedAt: role.UpdatedAt,
	}

	if role.Permissions != nil {
		for _, permission := range role.Permissions {
			response.Permissions = append(response.Permissions, *ToPermissionResponse(&permission))
		}
	}

	return response
}

func ToPermissionResponse(permission *models.Permission) *PermissionResponse {
	return &PermissionResponse{
		ID:        permission.ID,
		Name:      permission.Name,
		CreatedAt: permission.CreatedAt,
		UpdatedAt: permission.UpdatedAt,
	}
}

func ToCustomerResponse(customer *models.Customer) *CustomerResponse {
	response := &CustomerResponse{
		CustomerID:  customer.CustomerID,
		Name:        customer.Name,
		PhoneNumber: customer.PhoneNumber,
		Address:     customer.Address,
		Status:      customer.Status,
		CreatedAt:   customer.CreatedAt,
		UpdatedAt:   customer.UpdatedAt,
	}

	if customer.Vehicles != nil {
		for _, vehicle := range customer.Vehicles {
			response.Vehicles = append(response.Vehicles, *ToCustomerVehicleResponse(&vehicle))
		}
	}

	return response
}

func ToCustomerVehicleResponse(vehicle *models.CustomerVehicle) *CustomerVehicleResponse {
	response := &CustomerVehicleResponse{
		VehicleID:      vehicle.VehicleID,
		CustomerID:     vehicle.CustomerID,
		PlateNumber:    vehicle.PlateNumber,
		Brand:          vehicle.Brand,
		Model:          vehicle.Model,
		Type:           vehicle.Type,
		ProductionYear: vehicle.ProductionYear,
		ChassisNumber:  vehicle.ChassisNumber,
		EngineNumber:   vehicle.EngineNumber,
		Color:          vehicle.Color,
		Notes:          vehicle.Notes,
		CreatedAt:      vehicle.CreatedAt,
		UpdatedAt:      vehicle.UpdatedAt,
	}

	if vehicle.Customer != nil {
		response.Customer = &CustomerResponse{
			CustomerID:  vehicle.Customer.CustomerID,
			Name:        vehicle.Customer.Name,
			PhoneNumber: vehicle.Customer.PhoneNumber,
			Address:     vehicle.Customer.Address,
			Status:      vehicle.Customer.Status,
			CreatedAt:   vehicle.Customer.CreatedAt,
			UpdatedAt:   vehicle.Customer.UpdatedAt,
		}
	}

	return response
}