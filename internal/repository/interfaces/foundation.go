package interfaces

import (
	"boilerplate/internal/models"
	"context"
)

// UserRepository interface for user operations
type UserRepository interface {
	Create(ctx context.Context, user *models.User) error
	GetByID(ctx context.Context, id uint) (*models.User, error)
	GetByEmail(ctx context.Context, email string) (*models.User, error)
	Update(ctx context.Context, user *models.User) error
	Delete(ctx context.Context, id uint) error
	List(ctx context.Context, limit, offset int) ([]*models.User, error)
	GetByOutletID(ctx context.Context, outletID uint) ([]*models.User, error)
}

// OutletRepository interface for outlet operations
type OutletRepository interface {
	Create(ctx context.Context, outlet *models.Outlet) error
	GetByID(ctx context.Context, id uint) (*models.Outlet, error)
	Update(ctx context.Context, outlet *models.Outlet) error
	Delete(ctx context.Context, id uint) error
	List(ctx context.Context, limit, offset int) ([]*models.Outlet, error)
	GetByStatus(ctx context.Context, status models.StatusUmum) ([]*models.Outlet, error)
}

// RoleRepository interface for role operations
type RoleRepository interface {
	Create(ctx context.Context, role *models.Role) error
	GetByID(ctx context.Context, id uint) (*models.Role, error)
	GetByName(ctx context.Context, name string) (*models.Role, error)
	Update(ctx context.Context, role *models.Role) error
	Delete(ctx context.Context, id uint) error
	List(ctx context.Context, limit, offset int) ([]*models.Role, error)
	AttachPermissions(ctx context.Context, roleID uint, permissionIDs []uint) error
	DetachPermissions(ctx context.Context, roleID uint, permissionIDs []uint) error
}

// PermissionRepository interface for permission operations
type PermissionRepository interface {
	Create(ctx context.Context, permission *models.Permission) error
	GetByID(ctx context.Context, id uint) (*models.Permission, error)
	GetByName(ctx context.Context, name string) (*models.Permission, error)
	Update(ctx context.Context, permission *models.Permission) error
	Delete(ctx context.Context, id uint) error
	List(ctx context.Context, limit, offset int) ([]*models.Permission, error)
	GetByRoleID(ctx context.Context, roleID uint) ([]*models.Permission, error)
}