package implementations

import (
	"boilerplate/internal/repository"
	"boilerplate/internal/models"
	"boilerplate/internal/usecase/interfaces"
	"context"
	"errors"
	"time"

	"golang.org/x/crypto/bcrypt"
)

// UserUsecaseImpl implements UserUsecase interface
type UserUsecaseImpl struct {
	repo *repository.RepositoryManager
}

// NewUserUsecase creates a new user usecase
func NewUserUsecase(repo *repository.RepositoryManager) interfaces.UserUsecase {
	return &UserUsecaseImpl{repo: repo}
}

func (u *UserUsecaseImpl) CreateUser(ctx context.Context, req interfaces.CreateUserRequest) (*models.User, error) {
	// Check if email already exists
	existingUser, err := u.repo.User.GetByEmail(ctx, req.Email)
	if err == nil && existingUser != nil {
		return nil, errors.New("email already exists")
	}

	// Hash password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	user := &models.User{
		Name:      req.Name,
		Email:     req.Email,
		Password:  string(hashedPassword),
		OutletID:  req.OutletID,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	err = u.repo.User.Create(ctx, user)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (u *UserUsecaseImpl) GetUser(ctx context.Context, id uint) (*models.User, error) {
	return u.repo.User.GetByID(ctx, id)
}

func (u *UserUsecaseImpl) GetUserByEmail(ctx context.Context, email string) (*models.User, error) {
	return u.repo.User.GetByEmail(ctx, email)
}

func (u *UserUsecaseImpl) UpdateUser(ctx context.Context, id uint, req interfaces.UpdateUserRequest) (*models.User, error) {
	user, err := u.repo.User.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}

	if req.Name != nil {
		user.Name = *req.Name
	}
	if req.Email != nil {
		// Check if email is already taken by another user
		existingUser, err := u.repo.User.GetByEmail(ctx, *req.Email)
		if err == nil && existingUser != nil && existingUser.UserID != id {
			return nil, errors.New("email already exists")
		}
		user.Email = *req.Email
	}
	if req.OutletID != nil {
		user.OutletID = req.OutletID
	}

	user.UpdatedAt = time.Now()

	err = u.repo.User.Update(ctx, user)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (u *UserUsecaseImpl) DeleteUser(ctx context.Context, id uint) error {
	return u.repo.User.Delete(ctx, id)
}

func (u *UserUsecaseImpl) ListUsers(ctx context.Context, limit, offset int) ([]*models.User, error) {
	return u.repo.User.List(ctx, limit, offset)
}

func (u *UserUsecaseImpl) GetUsersByOutlet(ctx context.Context, outletID uint) ([]*models.User, error) {
	return u.repo.User.GetByOutletID(ctx, outletID)
}

func (u *UserUsecaseImpl) ChangePassword(ctx context.Context, userID uint, oldPassword, newPassword string) error {
	user, err := u.repo.User.GetByID(ctx, userID)
	if err != nil {
		return err
	}

	// Verify old password
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(oldPassword))
	if err != nil {
		return errors.New("invalid old password")
	}

	// Hash new password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(newPassword), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	user.Password = string(hashedPassword)
	user.UpdatedAt = time.Now()

	return u.repo.User.Update(ctx, user)
}

// OutletUsecaseImpl implements OutletUsecase interface
type OutletUsecaseImpl struct {
	repo *repository.RepositoryManager
}

// NewOutletUsecase creates a new outlet usecase
func NewOutletUsecase(repo *repository.RepositoryManager) interfaces.OutletUsecase {
	return &OutletUsecaseImpl{repo: repo}
}

func (u *OutletUsecaseImpl) CreateOutlet(ctx context.Context, req interfaces.CreateOutletRequest) (*models.Outlet, error) {
	outlet := &models.Outlet{
		OutletName:  req.OutletName,
		BranchType:  req.BranchType,
		City:        req.City,
		Address:     req.Address,
		PhoneNumber: req.PhoneNumber,
		Status:      req.Status,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	// Set default status if not provided
	if outlet.Status == "" {
		outlet.Status = models.StatusAktif
	}

	err := u.repo.Outlet.Create(ctx, outlet)
	if err != nil {
		return nil, err
	}

	return outlet, nil
}

func (u *OutletUsecaseImpl) GetOutlet(ctx context.Context, id uint) (*models.Outlet, error) {
	return u.repo.Outlet.GetByID(ctx, id)
}

func (u *OutletUsecaseImpl) UpdateOutlet(ctx context.Context, id uint, req interfaces.UpdateOutletRequest) (*models.Outlet, error) {
	outlet, err := u.repo.Outlet.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}

	if req.OutletName != nil {
		outlet.OutletName = *req.OutletName
	}
	if req.BranchType != nil {
		outlet.BranchType = *req.BranchType
	}
	if req.City != nil {
		outlet.City = *req.City
	}
	if req.Address != nil {
		outlet.Address = req.Address
	}
	if req.PhoneNumber != nil {
		outlet.PhoneNumber = req.PhoneNumber
	}
	if req.Status != nil {
		outlet.Status = *req.Status
	}

	outlet.UpdatedAt = time.Now()

	err = u.repo.Outlet.Update(ctx, outlet)
	if err != nil {
		return nil, err
	}

	return outlet, nil
}

func (u *OutletUsecaseImpl) DeleteOutlet(ctx context.Context, id uint) error {
	return u.repo.Outlet.Delete(ctx, id)
}

func (u *OutletUsecaseImpl) ListOutlets(ctx context.Context, limit, offset int) ([]*models.Outlet, error) {
	return u.repo.Outlet.List(ctx, limit, offset)
}

func (u *OutletUsecaseImpl) GetActiveOutlets(ctx context.Context) ([]*models.Outlet, error) {
	return u.repo.Outlet.GetByStatus(ctx, models.StatusAktif)
}
// RoleUsecaseImpl implements RoleUsecase interface  
type RoleUsecaseImpl struct {
repo *repository.RepositoryManager
}

// NewRoleUsecase creates a new role usecase
func NewRoleUsecase(repo *repository.RepositoryManager) interfaces.RoleUsecase {
return &RoleUsecaseImpl{repo: repo}
}

func (u *RoleUsecaseImpl) CreateRole(ctx context.Context, req interfaces.CreateRoleRequest) (*models.Role, error) {
// Check if role name already exists
existingRole, err := u.repo.Role.GetByName(ctx, req.Name)
if err == nil && existingRole != nil {
return nil, errors.New("role with this name already exists")
}

role := &models.Role{
Name:      req.Name,
CreatedAt: time.Now(),
UpdatedAt: time.Now(),
}

err = u.repo.Role.Create(ctx, role)
if err != nil {
return nil, err
}

return role, nil
}

func (u *RoleUsecaseImpl) GetRole(ctx context.Context, id uint) (*models.Role, error) {
return u.repo.Role.GetByID(ctx, id)
}

func (u *RoleUsecaseImpl) GetRoleByName(ctx context.Context, name string) (*models.Role, error) {
return u.repo.Role.GetByName(ctx, name)
}

func (u *RoleUsecaseImpl) UpdateRole(ctx context.Context, id uint, req interfaces.UpdateRoleRequest) (*models.Role, error) {
role, err := u.repo.Role.GetByID(ctx, id)
if err != nil {
return nil, err
}

if req.Name != nil {
// Check if new name already exists
existingRole, err := u.repo.Role.GetByName(ctx, *req.Name)
if err == nil && existingRole != nil && existingRole.ID != id {
return nil, errors.New("role with this name already exists")
}
role.Name = *req.Name
}

role.UpdatedAt = time.Now()

err = u.repo.Role.Update(ctx, role)
if err != nil {
return nil, err
}

return role, nil
}

func (u *RoleUsecaseImpl) DeleteRole(ctx context.Context, id uint) error {
return u.repo.Role.Delete(ctx, id)
}

func (u *RoleUsecaseImpl) ListRoles(ctx context.Context, limit, offset int) ([]*models.Role, error) {
return u.repo.Role.List(ctx, limit, offset)
}

func (u *RoleUsecaseImpl) AttachPermissions(ctx context.Context, roleID uint, permissionIDs []uint) error {
	return u.repo.Role.AttachPermissions(ctx, roleID, permissionIDs)
}

func (u *RoleUsecaseImpl) DetachPermissions(ctx context.Context, roleID uint, permissionIDs []uint) error {
	return u.repo.Role.DetachPermissions(ctx, roleID, permissionIDs)
}

// PermissionUsecaseImpl implements PermissionUsecase interface
type PermissionUsecaseImpl struct {
repo *repository.RepositoryManager
}

// NewPermissionUsecase creates a new permission usecase
func NewPermissionUsecase(repo *repository.RepositoryManager) interfaces.PermissionUsecase {
return &PermissionUsecaseImpl{repo: repo}
}

func (u *PermissionUsecaseImpl) CreatePermission(ctx context.Context, req interfaces.CreatePermissionRequest) (*models.Permission, error) {
// Check if permission name already exists
existingPermission, err := u.repo.Permission.GetByName(ctx, req.Name)
if err == nil && existingPermission != nil {
return nil, errors.New("permission with this name already exists")
}

permission := &models.Permission{
Name:      req.Name,
CreatedAt: time.Now(),
UpdatedAt: time.Now(),
}

err = u.repo.Permission.Create(ctx, permission)
if err != nil {
return nil, err
}

return permission, nil
}

func (u *PermissionUsecaseImpl) GetPermission(ctx context.Context, id uint) (*models.Permission, error) {
return u.repo.Permission.GetByID(ctx, id)
}

func (u *PermissionUsecaseImpl) GetPermissionByName(ctx context.Context, name string) (*models.Permission, error) {
return u.repo.Permission.GetByName(ctx, name)
}

func (u *PermissionUsecaseImpl) UpdatePermission(ctx context.Context, id uint, req interfaces.UpdatePermissionRequest) (*models.Permission, error) {
permission, err := u.repo.Permission.GetByID(ctx, id)
if err != nil {
return nil, err
}

if req.Name != nil {
// Check if new name already exists
existingPermission, err := u.repo.Permission.GetByName(ctx, *req.Name)
if err == nil && existingPermission != nil && existingPermission.ID != id {
return nil, errors.New("permission with this name already exists")
}
permission.Name = *req.Name
}

permission.UpdatedAt = time.Now()

err = u.repo.Permission.Update(ctx, permission)
if err != nil {
return nil, err
}

return permission, nil
}

func (u *PermissionUsecaseImpl) DeletePermission(ctx context.Context, id uint) error {
return u.repo.Permission.Delete(ctx, id)
}

func (u *PermissionUsecaseImpl) ListPermissions(ctx context.Context, limit, offset int) ([]*models.Permission, error) {
return u.repo.Permission.List(ctx, limit, offset)
}



func (u *PermissionUsecaseImpl) GetPermissionsByRole(ctx context.Context, roleID uint) ([]*models.Permission, error) {
return u.repo.Permission.GetByRoleID(ctx, roleID)
}
