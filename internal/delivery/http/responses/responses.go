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