package implementations

import (
	"boilerplate/internal/models"
	"boilerplate/internal/repository/interfaces"
	"context"

	"gorm.io/gorm"
)

// UserRepositoryImpl implements UserRepository interface
type UserRepositoryImpl struct {
	db *gorm.DB
}

// NewUserRepository creates a new user repository
func NewUserRepository(db *gorm.DB) interfaces.UserRepository {
	return &UserRepositoryImpl{db: db}
}

func (r *UserRepositoryImpl) Create(ctx context.Context, user *models.User) error {
	return r.db.WithContext(ctx).Create(user).Error
}

func (r *UserRepositoryImpl) GetByID(ctx context.Context, id uint) (*models.User, error) {
	var user models.User
	err := r.db.WithContext(ctx).Preload("Outlet").First(&user, id).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *UserRepositoryImpl) GetByEmail(ctx context.Context, email string) (*models.User, error) {
	var user models.User
	err := r.db.WithContext(ctx).Where("email = ?", email).Preload("Outlet").First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *UserRepositoryImpl) Update(ctx context.Context, user *models.User) error {
	return r.db.WithContext(ctx).Save(user).Error
}

func (r *UserRepositoryImpl) Delete(ctx context.Context, id uint) error {
	return r.db.WithContext(ctx).Delete(&models.User{}, id).Error
}

func (r *UserRepositoryImpl) List(ctx context.Context, limit, offset int) ([]*models.User, error) {
	var users []*models.User
	err := r.db.WithContext(ctx).Preload("Outlet").Limit(limit).Offset(offset).Find(&users).Error
	return users, err
}

func (r *UserRepositoryImpl) GetByOutletID(ctx context.Context, outletID uint) ([]*models.User, error) {
	var users []*models.User
	err := r.db.WithContext(ctx).Where("outlet_id = ?", outletID).Preload("Outlet").Find(&users).Error
	return users, err
}

// OutletRepositoryImpl implements OutletRepository interface
type OutletRepositoryImpl struct {
	db *gorm.DB
}

// NewOutletRepository creates a new outlet repository
func NewOutletRepository(db *gorm.DB) interfaces.OutletRepository {
	return &OutletRepositoryImpl{db: db}
}

func (r *OutletRepositoryImpl) Create(ctx context.Context, outlet *models.Outlet) error {
	return r.db.WithContext(ctx).Create(outlet).Error
}

func (r *OutletRepositoryImpl) GetByID(ctx context.Context, id uint) (*models.Outlet, error) {
	var outlet models.Outlet
	err := r.db.WithContext(ctx).Preload("Users").First(&outlet, id).Error
	if err != nil {
		return nil, err
	}
	return &outlet, nil
}

func (r *OutletRepositoryImpl) Update(ctx context.Context, outlet *models.Outlet) error {
	return r.db.WithContext(ctx).Save(outlet).Error
}

func (r *OutletRepositoryImpl) Delete(ctx context.Context, id uint) error {
	return r.db.WithContext(ctx).Delete(&models.Outlet{}, id).Error
}

func (r *OutletRepositoryImpl) List(ctx context.Context, limit, offset int) ([]*models.Outlet, error) {
	var outlets []*models.Outlet
	err := r.db.WithContext(ctx).Limit(limit).Offset(offset).Find(&outlets).Error
	return outlets, err
}

func (r *OutletRepositoryImpl) GetByStatus(ctx context.Context, status models.StatusUmum) ([]*models.Outlet, error) {
	var outlets []*models.Outlet
	err := r.db.WithContext(ctx).Where("status = ?", status).Find(&outlets).Error
	return outlets, err
}

// RoleRepositoryImpl implements RoleRepository interface
type RoleRepositoryImpl struct {
	db *gorm.DB
}

// NewRoleRepository creates a new role repository
func NewRoleRepository(db *gorm.DB) interfaces.RoleRepository {
	return &RoleRepositoryImpl{db: db}
}

func (r *RoleRepositoryImpl) Create(ctx context.Context, role *models.Role) error {
	return r.db.WithContext(ctx).Create(role).Error
}

func (r *RoleRepositoryImpl) GetByID(ctx context.Context, id uint) (*models.Role, error) {
	var role models.Role
	err := r.db.WithContext(ctx).Preload("Permissions").First(&role, id).Error
	if err != nil {
		return nil, err
	}
	return &role, nil
}

func (r *RoleRepositoryImpl) GetByName(ctx context.Context, name string) (*models.Role, error) {
	var role models.Role
	err := r.db.WithContext(ctx).Where("name = ?", name).Preload("Permissions").First(&role).Error
	if err != nil {
		return nil, err
	}
	return &role, nil
}

func (r *RoleRepositoryImpl) Update(ctx context.Context, role *models.Role) error {
	return r.db.WithContext(ctx).Save(role).Error
}

func (r *RoleRepositoryImpl) Delete(ctx context.Context, id uint) error {
	return r.db.WithContext(ctx).Delete(&models.Role{}, id).Error
}

func (r *RoleRepositoryImpl) List(ctx context.Context, limit, offset int) ([]*models.Role, error) {
	var roles []*models.Role
	err := r.db.WithContext(ctx).Preload("Permissions").Limit(limit).Offset(offset).Find(&roles).Error
	return roles, err
}

func (r *RoleRepositoryImpl) AttachPermissions(ctx context.Context, roleID uint, permissionIDs []uint) error {
	role := models.Role{ID: roleID}
	var permissions []models.Permission
	for _, id := range permissionIDs {
		permissions = append(permissions, models.Permission{ID: id})
	}
	return r.db.WithContext(ctx).Model(&role).Association("Permissions").Append(permissions)
}

func (r *RoleRepositoryImpl) DetachPermissions(ctx context.Context, roleID uint, permissionIDs []uint) error {
	role := models.Role{ID: roleID}
	var permissions []models.Permission
	for _, id := range permissionIDs {
		permissions = append(permissions, models.Permission{ID: id})
	}
	return r.db.WithContext(ctx).Model(&role).Association("Permissions").Delete(permissions)
}

// PermissionRepositoryImpl implements PermissionRepository interface
type PermissionRepositoryImpl struct {
	db *gorm.DB
}

// NewPermissionRepository creates a new permission repository
func NewPermissionRepository(db *gorm.DB) interfaces.PermissionRepository {
	return &PermissionRepositoryImpl{db: db}
}

func (r *PermissionRepositoryImpl) Create(ctx context.Context, permission *models.Permission) error {
	return r.db.WithContext(ctx).Create(permission).Error
}

func (r *PermissionRepositoryImpl) GetByID(ctx context.Context, id uint) (*models.Permission, error) {
	var permission models.Permission
	err := r.db.WithContext(ctx).Preload("Roles").First(&permission, id).Error
	if err != nil {
		return nil, err
	}
	return &permission, nil
}

func (r *PermissionRepositoryImpl) GetByName(ctx context.Context, name string) (*models.Permission, error) {
	var permission models.Permission
	err := r.db.WithContext(ctx).Where("name = ?", name).Preload("Roles").First(&permission).Error
	if err != nil {
		return nil, err
	}
	return &permission, nil
}

func (r *PermissionRepositoryImpl) Update(ctx context.Context, permission *models.Permission) error {
	return r.db.WithContext(ctx).Save(permission).Error
}

func (r *PermissionRepositoryImpl) Delete(ctx context.Context, id uint) error {
	return r.db.WithContext(ctx).Delete(&models.Permission{}, id).Error
}

func (r *PermissionRepositoryImpl) List(ctx context.Context, limit, offset int) ([]*models.Permission, error) {
	var permissions []*models.Permission
	err := r.db.WithContext(ctx).Limit(limit).Offset(offset).Find(&permissions).Error
	return permissions, err
}

func (r *PermissionRepositoryImpl) GetByRoleID(ctx context.Context, roleID uint) ([]*models.Permission, error) {
	var permissions []*models.Permission
	err := r.db.WithContext(ctx).
		Joins("JOIN role_has_permissions ON permissions.id = role_has_permissions.permission_id").
		Where("role_has_permissions.role_id = ?", roleID).
		Find(&permissions).Error
	return permissions, err
}