package implementations

import (
	"boilerplate/internal/models"
	"boilerplate/internal/repository/interfaces"
	"context"
	"database/sql"

	"github.com/jmoiron/sqlx"
)

// UserRepositorySQLX implements UserRepository interface using SQLx
type UserRepositorySQLX struct {
	db *sqlx.DB
}

// NewUserRepositorySQLX creates a new user repository using SQLx
func NewUserRepositorySQLX(db *sqlx.DB) interfaces.UserRepository {
	return &UserRepositorySQLX{db: db}
}

func (r *UserRepositorySQLX) Create(ctx context.Context, user *models.User) error {
	query := `
		INSERT INTO users (name, email, password, outlet_id, created_by)
		VALUES ($1, $2, $3, $4, $5)
		RETURNING user_id, created_at, updated_at
	`
	
	row := r.db.QueryRowxContext(ctx, query,
		user.Name,
		user.Email,
		user.Password,
		user.OutletID,
		user.CreatedBy,
	)
	
	return row.Scan(&user.UserID, &user.CreatedAt, &user.UpdatedAt)
}

func (r *UserRepositorySQLX) GetByID(ctx context.Context, id uint) (*models.User, error) {
	var user models.User
	query := `
		SELECT user_id, name, email, password, outlet_id, 
			   created_at, updated_at, deleted_at, created_by
		FROM users 
		WHERE user_id = $1 AND deleted_at IS NULL
	`
	
	err := r.db.GetContext(ctx, &user, query, id)
	if err != nil {
		return nil, err
	}

	// Load outlet separately
	if user.OutletID != nil {
		outlet, err := r.getOutletByID(ctx, *user.OutletID)
		if err != nil && err != sql.ErrNoRows {
			return nil, err
		}
		user.Outlet = outlet
	}

	return &user, nil
}

func (r *UserRepositorySQLX) GetByEmail(ctx context.Context, email string) (*models.User, error) {
	var user models.User
	query := `
		SELECT user_id, name, email, password, outlet_id, 
			   created_at, updated_at, deleted_at, created_by
		FROM users 
		WHERE email = $1 AND deleted_at IS NULL
	`
	
	err := r.db.GetContext(ctx, &user, query, email)
	if err != nil {
		return nil, err
	}

	// Load outlet separately
	if user.OutletID != nil {
		outlet, err := r.getOutletByID(ctx, *user.OutletID)
		if err != nil && err != sql.ErrNoRows {
			return nil, err
		}
		user.Outlet = outlet
	}

	return &user, nil
}

func (r *UserRepositorySQLX) Update(ctx context.Context, user *models.User) error {
	query := `
		UPDATE users 
		SET name = $1, email = $2, password = $3, outlet_id = $4, updated_at = NOW()
		WHERE user_id = $5 AND deleted_at IS NULL
	`
	
	_, err := r.db.ExecContext(ctx, query,
		user.Name,
		user.Email,
		user.Password,
		user.OutletID,
		user.UserID,
	)
	
	return err
}

func (r *UserRepositorySQLX) Delete(ctx context.Context, id uint) error {
	query := `
		UPDATE users 
		SET deleted_at = NOW(), updated_at = NOW()
		WHERE user_id = $1 AND deleted_at IS NULL
	`
	
	_, err := r.db.ExecContext(ctx, query, id)
	return err
}

func (r *UserRepositorySQLX) List(ctx context.Context, limit, offset int) ([]*models.User, error) {
	var users []*models.User
	query := `
		SELECT user_id, name, email, password, outlet_id, 
			   created_at, updated_at, deleted_at, created_by
		FROM users 
		WHERE deleted_at IS NULL
		ORDER BY created_at DESC
		LIMIT $1 OFFSET $2
	`
	
	err := r.db.SelectContext(ctx, &users, query, limit, offset)
	if err != nil {
		return nil, err
	}

	// Load outlets for each user
	for i := range users {
		if users[i].OutletID != nil {
			outlet, err := r.getOutletByID(ctx, *users[i].OutletID)
			if err != nil && err != sql.ErrNoRows {
				return nil, err
			}
			users[i].Outlet = outlet
		}
	}

	return users, nil
}

func (r *UserRepositorySQLX) GetByOutletID(ctx context.Context, outletID uint) ([]*models.User, error) {
	var users []*models.User
	query := `
		SELECT user_id, name, email, password, outlet_id, 
			   created_at, updated_at, deleted_at, created_by
		FROM users 
		WHERE outlet_id = $1 AND deleted_at IS NULL
		ORDER BY created_at DESC
	`
	
	err := r.db.SelectContext(ctx, &users, query, outletID)
	if err != nil {
		return nil, err
	}

	// Load outlet for each user
	for i := range users {
		outlet, err := r.getOutletByID(ctx, outletID)
		if err != nil && err != sql.ErrNoRows {
			return nil, err
		}
		users[i].Outlet = outlet
	}

	return users, nil
}

// Helper method to get outlet by ID
func (r *UserRepositorySQLX) getOutletByID(ctx context.Context, outletID uint) (*models.Outlet, error) {
	var outlet models.Outlet
	query := `
		SELECT outlet_id, outlet_name, branch_type, city, address, phone_number, status, 
			   created_at, updated_at, deleted_at, created_by
		FROM outlets 
		WHERE outlet_id = $1 AND deleted_at IS NULL
	`
	
	err := r.db.GetContext(ctx, &outlet, query, outletID)
	if err != nil {
		return nil, err
	}
	
	return &outlet, nil
}

// OutletRepositorySQLX implements OutletRepository interface using SQLx
type OutletRepositorySQLX struct {
	db *sqlx.DB
}

// NewOutletRepositorySQLX creates a new outlet repository using SQLx
func NewOutletRepositorySQLX(db *sqlx.DB) interfaces.OutletRepository {
	return &OutletRepositorySQLX{db: db}
}

func (r *OutletRepositorySQLX) Create(ctx context.Context, outlet *models.Outlet) error {
	query := `
		INSERT INTO outlets (outlet_name, branch_type, city, address, phone_number, status, created_by)
		VALUES ($1, $2, $3, $4, $5, $6, $7)
		RETURNING outlet_id, created_at, updated_at
	`
	
	row := r.db.QueryRowxContext(ctx, query,
		outlet.OutletName,
		outlet.BranchType,
		outlet.City,
		outlet.Address,
		outlet.PhoneNumber,
		outlet.Status,
		outlet.CreatedBy,
	)
	
	return row.Scan(&outlet.OutletID, &outlet.CreatedAt, &outlet.UpdatedAt)
}

func (r *OutletRepositorySQLX) GetByID(ctx context.Context, id uint) (*models.Outlet, error) {
	var outlet models.Outlet
	query := `
		SELECT outlet_id, outlet_name, branch_type, city, address, phone_number, status, 
			   created_at, updated_at, deleted_at, created_by
		FROM outlets 
		WHERE outlet_id = $1 AND deleted_at IS NULL
	`
	
	err := r.db.GetContext(ctx, &outlet, query, id)
	if err != nil {
		return nil, err
	}

	// Load users separately
	users, err := r.getUsersByOutletID(ctx, id)
	if err != nil {
		return nil, err
	}
	outlet.Users = users

	return &outlet, nil
}

func (r *OutletRepositorySQLX) Update(ctx context.Context, outlet *models.Outlet) error {
	query := `
		UPDATE outlets 
		SET outlet_name = $1, branch_type = $2, city = $3, address = $4, 
			phone_number = $5, status = $6, updated_at = NOW()
		WHERE outlet_id = $7 AND deleted_at IS NULL
	`
	
	_, err := r.db.ExecContext(ctx, query,
		outlet.OutletName,
		outlet.BranchType,
		outlet.City,
		outlet.Address,
		outlet.PhoneNumber,
		outlet.Status,
		outlet.OutletID,
	)
	
	return err
}

func (r *OutletRepositorySQLX) Delete(ctx context.Context, id uint) error {
	query := `
		UPDATE outlets 
		SET deleted_at = NOW(), updated_at = NOW()
		WHERE outlet_id = $1 AND deleted_at IS NULL
	`
	
	_, err := r.db.ExecContext(ctx, query, id)
	return err
}

func (r *OutletRepositorySQLX) List(ctx context.Context, limit, offset int) ([]*models.Outlet, error) {
	var outlets []*models.Outlet
	query := `
		SELECT outlet_id, outlet_name, branch_type, city, address, phone_number, status, 
			   created_at, updated_at, deleted_at, created_by
		FROM outlets 
		WHERE deleted_at IS NULL
		ORDER BY created_at DESC
		LIMIT $1 OFFSET $2
	`
	
	err := r.db.SelectContext(ctx, &outlets, query, limit, offset)
	if err != nil {
		return nil, err
	}

	// Load users for each outlet
	for i := range outlets {
		users, err := r.getUsersByOutletID(ctx, outlets[i].OutletID)
		if err != nil {
			return nil, err
		}
		outlets[i].Users = users
	}

	return outlets, nil
}

func (r *OutletRepositorySQLX) GetByStatus(ctx context.Context, status models.StatusUmum) ([]*models.Outlet, error) {
	var outlets []*models.Outlet
	query := `
		SELECT outlet_id, outlet_name, branch_type, city, address, phone_number, status, 
			   created_at, updated_at, deleted_at, created_by
		FROM outlets 
		WHERE status = $1 AND deleted_at IS NULL
		ORDER BY created_at DESC
	`
	
	err := r.db.SelectContext(ctx, &outlets, query, status)
	if err != nil {
		return nil, err
	}

	// Load users for each outlet
	for i := range outlets {
		users, err := r.getUsersByOutletID(ctx, outlets[i].OutletID)
		if err != nil {
			return nil, err
		}
		outlets[i].Users = users
	}

	return outlets, nil
}

func (r *OutletRepositorySQLX) GetByCity(ctx context.Context, city string) ([]*models.Outlet, error) {
	var outlets []*models.Outlet
	query := `
		SELECT outlet_id, outlet_name, branch_type, city, address, phone_number, status, 
			   created_at, updated_at, deleted_at, created_by
		FROM outlets 
		WHERE city = $1 AND deleted_at IS NULL
		ORDER BY created_at DESC
	`
	
	err := r.db.SelectContext(ctx, &outlets, query, city)
	if err != nil {
		return nil, err
	}

	// Load users for each outlet
	for i := range outlets {
		users, err := r.getUsersByOutletID(ctx, outlets[i].OutletID)
		if err != nil {
			return nil, err
		}
		outlets[i].Users = users
	}

	return outlets, nil
}

// Helper method to get users by outlet ID
func (r *OutletRepositorySQLX) getUsersByOutletID(ctx context.Context, outletID uint) ([]models.User, error) {
	var users []models.User
	query := `
		SELECT user_id, name, email, password, outlet_id, 
			   created_at, updated_at, deleted_at, created_by
		FROM users 
		WHERE outlet_id = $1 AND deleted_at IS NULL
		ORDER BY created_at DESC
	`
	
	err := r.db.SelectContext(ctx, &users, query, outletID)
	return users, err
}

// RoleRepositorySQLX implements RoleRepository interface using SQLx
type RoleRepositorySQLX struct {
	db *sqlx.DB
}

// NewRoleRepositorySQLX creates a new role repository using SQLx
func NewRoleRepositorySQLX(db *sqlx.DB) interfaces.RoleRepository {
	return &RoleRepositorySQLX{db: db}
}

func (r *RoleRepositorySQLX) Create(ctx context.Context, role *models.Role) error {
	query := `
		INSERT INTO roles (name, created_by)
		VALUES ($1, $2)
		RETURNING id, created_at, updated_at
	`
	
	row := r.db.QueryRowxContext(ctx, query, role.Name, role.CreatedBy)
	return row.Scan(&role.ID, &role.CreatedAt, &role.UpdatedAt)
}

func (r *RoleRepositorySQLX) GetByID(ctx context.Context, id uint) (*models.Role, error) {
	var role models.Role
	query := `
		SELECT id, name, created_at, updated_at, deleted_at, created_by
		FROM roles 
		WHERE id = $1 AND deleted_at IS NULL
	`
	
	err := r.db.GetContext(ctx, &role, query, id)
	if err != nil {
		return nil, err
	}

	// Load permissions separately
	permissions, err := r.getPermissionsByRoleID(ctx, id)
	if err != nil {
		return nil, err
	}
	role.Permissions = permissions

	return &role, nil
}

func (r *RoleRepositorySQLX) GetByName(ctx context.Context, name string) (*models.Role, error) {
	var role models.Role
	query := `
		SELECT id, name, created_at, updated_at, deleted_at, created_by
		FROM roles 
		WHERE name = $1 AND deleted_at IS NULL
	`
	
	err := r.db.GetContext(ctx, &role, query, name)
	if err != nil {
		return nil, err
	}

	// Load permissions separately
	permissions, err := r.getPermissionsByRoleID(ctx, role.ID)
	if err != nil {
		return nil, err
	}
	role.Permissions = permissions

	return &role, nil
}

func (r *RoleRepositorySQLX) Update(ctx context.Context, role *models.Role) error {
	query := `
		UPDATE roles 
		SET name = $1, updated_at = NOW()
		WHERE id = $2 AND deleted_at IS NULL
	`
	
	_, err := r.db.ExecContext(ctx, query, role.Name, role.ID)
	return err
}

func (r *RoleRepositorySQLX) Delete(ctx context.Context, id uint) error {
	query := `
		UPDATE roles 
		SET deleted_at = NOW(), updated_at = NOW()
		WHERE id = $1 AND deleted_at IS NULL
	`
	
	_, err := r.db.ExecContext(ctx, query, id)
	return err
}

func (r *RoleRepositorySQLX) List(ctx context.Context, limit, offset int) ([]*models.Role, error) {
	var roles []*models.Role
	query := `
		SELECT id, name, created_at, updated_at, deleted_at, created_by
		FROM roles 
		WHERE deleted_at IS NULL
		ORDER BY created_at DESC
		LIMIT $1 OFFSET $2
	`
	
	err := r.db.SelectContext(ctx, &roles, query, limit, offset)
	if err != nil {
		return nil, err
	}

	// Load permissions for each role
	for i := range roles {
		permissions, err := r.getPermissionsByRoleID(ctx, roles[i].ID)
		if err != nil {
			return nil, err
		}
		roles[i].Permissions = permissions
	}

	return roles, nil
}

func (r *RoleRepositorySQLX) AttachPermissions(ctx context.Context, roleID uint, permissionIDs []uint) error {
	tx, err := r.db.Beginx()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	query := `
		INSERT INTO role_has_permissions (role_id, permission_id)
		VALUES ($1, $2)
		ON CONFLICT (role_id, permission_id) DO NOTHING
	`

	for _, permissionID := range permissionIDs {
		_, err := tx.ExecContext(ctx, query, roleID, permissionID)
		if err != nil {
			return err
		}
	}

	return tx.Commit()
}

func (r *RoleRepositorySQLX) DetachPermissions(ctx context.Context, roleID uint, permissionIDs []uint) error {
	tx, err := r.db.Beginx()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	query := `
		DELETE FROM role_has_permissions
		WHERE role_id = $1 AND permission_id = $2
	`

	for _, permissionID := range permissionIDs {
		_, err := tx.ExecContext(ctx, query, roleID, permissionID)
		if err != nil {
			return err
		}
	}

	return tx.Commit()
}

// Helper method to get permissions by role ID
func (r *RoleRepositorySQLX) getPermissionsByRoleID(ctx context.Context, roleID uint) ([]models.Permission, error) {
	var permissions []models.Permission
	query := `
		SELECT p.id, p.name, p.created_at, p.updated_at, p.deleted_at, p.created_by
		FROM permissions p
		INNER JOIN role_has_permissions rhp ON p.id = rhp.permission_id
		WHERE rhp.role_id = $1 AND p.deleted_at IS NULL
		ORDER BY p.created_at DESC
	`
	
	err := r.db.SelectContext(ctx, &permissions, query, roleID)
	return permissions, err
}

// PermissionRepositorySQLX implements PermissionRepository interface using SQLx
type PermissionRepositorySQLX struct {
	db *sqlx.DB
}

// NewPermissionRepositorySQLX creates a new permission repository using SQLx
func NewPermissionRepositorySQLX(db *sqlx.DB) interfaces.PermissionRepository {
	return &PermissionRepositorySQLX{db: db}
}

func (r *PermissionRepositorySQLX) Create(ctx context.Context, permission *models.Permission) error {
	query := `
		INSERT INTO permissions (name, created_by)
		VALUES ($1, $2)
		RETURNING id, created_at, updated_at
	`
	
	row := r.db.QueryRowxContext(ctx, query, permission.Name, permission.CreatedBy)
	return row.Scan(&permission.ID, &permission.CreatedAt, &permission.UpdatedAt)
}

func (r *PermissionRepositorySQLX) GetByID(ctx context.Context, id uint) (*models.Permission, error) {
	var permission models.Permission
	query := `
		SELECT id, name, created_at, updated_at, deleted_at, created_by
		FROM permissions 
		WHERE id = $1 AND deleted_at IS NULL
	`
	
	err := r.db.GetContext(ctx, &permission, query, id)
	if err != nil {
		return nil, err
	}

	// Load roles separately
	roles, err := r.getRolesByPermissionID(ctx, id)
	if err != nil {
		return nil, err
	}
	permission.Roles = roles

	return &permission, nil
}

func (r *PermissionRepositorySQLX) GetByName(ctx context.Context, name string) (*models.Permission, error) {
	var permission models.Permission
	query := `
		SELECT id, name, created_at, updated_at, deleted_at, created_by
		FROM permissions 
		WHERE name = $1 AND deleted_at IS NULL
	`
	
	err := r.db.GetContext(ctx, &permission, query, name)
	if err != nil {
		return nil, err
	}

	// Load roles separately
	roles, err := r.getRolesByPermissionID(ctx, permission.ID)
	if err != nil {
		return nil, err
	}
	permission.Roles = roles

	return &permission, nil
}

func (r *PermissionRepositorySQLX) Update(ctx context.Context, permission *models.Permission) error {
	query := `
		UPDATE permissions 
		SET name = $1, updated_at = NOW()
		WHERE id = $2 AND deleted_at IS NULL
	`
	
	_, err := r.db.ExecContext(ctx, query, permission.Name, permission.ID)
	return err
}

func (r *PermissionRepositorySQLX) Delete(ctx context.Context, id uint) error {
	query := `
		UPDATE permissions 
		SET deleted_at = NOW(), updated_at = NOW()
		WHERE id = $1 AND deleted_at IS NULL
	`
	
	_, err := r.db.ExecContext(ctx, query, id)
	return err
}

func (r *PermissionRepositorySQLX) List(ctx context.Context, limit, offset int) ([]*models.Permission, error) {
	var permissions []*models.Permission
	query := `
		SELECT id, name, created_at, updated_at, deleted_at, created_by
		FROM permissions 
		WHERE deleted_at IS NULL
		ORDER BY created_at DESC
		LIMIT $1 OFFSET $2
	`
	
	err := r.db.SelectContext(ctx, &permissions, query, limit, offset)
	if err != nil {
		return nil, err
	}

	// Load roles for each permission
	for i := range permissions {
		roles, err := r.getRolesByPermissionID(ctx, permissions[i].ID)
		if err != nil {
			return nil, err
		}
		permissions[i].Roles = roles
	}

	return permissions, nil
}

func (r *PermissionRepositorySQLX) GetByRoleID(ctx context.Context, roleID uint) ([]*models.Permission, error) {
	var permissions []*models.Permission
	query := `
		SELECT p.id, p.name, p.created_at, p.updated_at, p.deleted_at, p.created_by
		FROM permissions p
		INNER JOIN role_has_permissions rhp ON p.id = rhp.permission_id
		WHERE rhp.role_id = $1 AND p.deleted_at IS NULL
		ORDER BY p.created_at DESC
	`
	
	err := r.db.SelectContext(ctx, &permissions, query, roleID)
	return permissions, err
}

// Helper method to get roles by permission ID
func (r *PermissionRepositorySQLX) getRolesByPermissionID(ctx context.Context, permissionID uint) ([]models.Role, error) {
	var roles []models.Role
	query := `
		SELECT r.id, r.name, r.created_at, r.updated_at, r.deleted_at, r.created_by
		FROM roles r
		INNER JOIN role_has_permissions rhp ON r.id = rhp.role_id
		WHERE rhp.permission_id = $1 AND r.deleted_at IS NULL
		ORDER BY r.created_at DESC
	`
	
	err := r.db.SelectContext(ctx, &roles, query, permissionID)
	return roles, err
}