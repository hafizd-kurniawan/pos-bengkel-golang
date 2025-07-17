package interfaces

import (
	"boilerplate/internal/models"
	"context"
)

// UserUsecase interface for user business logic
type UserUsecase interface {
	CreateUser(ctx context.Context, req CreateUserRequest) (*models.User, error)
	GetUser(ctx context.Context, id uint) (*models.User, error)
	GetUserByEmail(ctx context.Context, email string) (*models.User, error)
	UpdateUser(ctx context.Context, id uint, req UpdateUserRequest) (*models.User, error)
	DeleteUser(ctx context.Context, id uint) error
	ListUsers(ctx context.Context, limit, offset int) ([]*models.User, error)
	GetUsersByOutlet(ctx context.Context, outletID uint) ([]*models.User, error)
	ChangePassword(ctx context.Context, userID uint, oldPassword, newPassword string) error
}

// OutletUsecase interface for outlet business logic
type OutletUsecase interface {
	CreateOutlet(ctx context.Context, req CreateOutletRequest) (*models.Outlet, error)
	GetOutlet(ctx context.Context, id uint) (*models.Outlet, error)
	UpdateOutlet(ctx context.Context, id uint, req UpdateOutletRequest) (*models.Outlet, error)
	DeleteOutlet(ctx context.Context, id uint) error
	ListOutlets(ctx context.Context, limit, offset int) ([]*models.Outlet, error)
	GetActiveOutlets(ctx context.Context) ([]*models.Outlet, error)
}

// RoleUsecase interface for role business logic
type RoleUsecase interface {
	CreateRole(ctx context.Context, req CreateRoleRequest) (*models.Role, error)
	GetRole(ctx context.Context, id uint) (*models.Role, error)
	GetRoleByName(ctx context.Context, name string) (*models.Role, error)
	UpdateRole(ctx context.Context, id uint, req UpdateRoleRequest) (*models.Role, error)
	DeleteRole(ctx context.Context, id uint) error
	ListRoles(ctx context.Context, limit, offset int) ([]*models.Role, error)
	AttachPermissions(ctx context.Context, roleID uint, permissionIDs []uint) error
	DetachPermissions(ctx context.Context, roleID uint, permissionIDs []uint) error
}

// PermissionUsecase interface for permission business logic
type PermissionUsecase interface {
	CreatePermission(ctx context.Context, req CreatePermissionRequest) (*models.Permission, error)
	GetPermission(ctx context.Context, id uint) (*models.Permission, error)
	GetPermissionByName(ctx context.Context, name string) (*models.Permission, error)
	UpdatePermission(ctx context.Context, id uint, req UpdatePermissionRequest) (*models.Permission, error)
	DeletePermission(ctx context.Context, id uint) error
	ListPermissions(ctx context.Context, limit, offset int) ([]*models.Permission, error)
	GetPermissionsByRole(ctx context.Context, roleID uint) ([]*models.Permission, error)
}

// Request structs
type CreateUserRequest struct {
	Name     string `json:"name" validate:"required"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=6"`
	OutletID *uint  `json:"outlet_id"`
}

type UpdateUserRequest struct {
	Name     *string `json:"name"`
	Email    *string `json:"email" validate:"omitempty,email"`
	OutletID *uint   `json:"outlet_id"`
}

type CreateOutletRequest struct {
	OutletName  string             `json:"outlet_name" validate:"required"`
	BranchType  string             `json:"branch_type" validate:"required"`
	City        string             `json:"city" validate:"required"`
	Address     *string            `json:"address"`
	PhoneNumber *string            `json:"phone_number"`
	Status      models.StatusUmum  `json:"status"`
}

type UpdateOutletRequest struct {
	OutletName  *string            `json:"outlet_name"`
	BranchType  *string            `json:"branch_type"`
	City        *string            `json:"city"`
	Address     *string            `json:"address"`
	PhoneNumber *string            `json:"phone_number"`
	Status      *models.StatusUmum `json:"status"`
}

type CreateRoleRequest struct {
	Name string `json:"name" validate:"required"`
}

type UpdateRoleRequest struct {
	Name *string `json:"name"`
}

type CreatePermissionRequest struct {
	Name string `json:"name" validate:"required"`
}

type UpdatePermissionRequest struct {
	Name *string `json:"name"`
}