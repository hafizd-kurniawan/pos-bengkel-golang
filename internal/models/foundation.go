package models

import (
	"time"
)

// Users table
type User struct {
	UserID    uint       `db:"user_id" json:"user_id"`
	Name      string     `db:"name" json:"name"`
	Email     string     `db:"email" json:"email"`
	Password  string     `db:"password" json:"-"`
	OutletID  *uint      `db:"outlet_id" json:"outlet_id"`
	CreatedAt time.Time  `db:"created_at" json:"created_at"`
	UpdatedAt time.Time  `db:"updated_at" json:"updated_at"`
	DeletedAt *time.Time `db:"deleted_at" json:"deleted_at"`
	CreatedBy *uint      `db:"created_by" json:"created_by"`

	// Relationships (populated separately)
	Outlet *Outlet `json:"outlet,omitempty"`
}

// Outlets table
type Outlet struct {
	OutletID     uint       `db:"outlet_id" json:"outlet_id"`
	OutletName   string     `db:"outlet_name" json:"outlet_name"`
	BranchType   string     `db:"branch_type" json:"branch_type"`
	City         string     `db:"city" json:"city"`
	Address      *string    `db:"address" json:"address"`
	PhoneNumber  *string    `db:"phone_number" json:"phone_number"`
	Status       StatusUmum `db:"status" json:"status"`
	CreatedAt    time.Time  `db:"created_at" json:"created_at"`
	UpdatedAt    time.Time  `db:"updated_at" json:"updated_at"`
	DeletedAt    *time.Time `db:"deleted_at" json:"deleted_at"`
	CreatedBy    *uint      `db:"created_by" json:"created_by"`

	// Relationships (populated separately)
	Users []User `json:"users,omitempty"`
}

// Roles table (from Spatie Permission)
type Role struct {
	ID        uint       `db:"id" json:"id"`
	Name      string     `db:"name" json:"name"`
	CreatedAt time.Time  `db:"created_at" json:"created_at"`
	UpdatedAt time.Time  `db:"updated_at" json:"updated_at"`
	DeletedAt *time.Time `db:"deleted_at" json:"deleted_at"`
	CreatedBy *uint      `db:"created_by" json:"created_by"`

	// Relationships (populated separately)
	Permissions []Permission `json:"permissions,omitempty"`
}

// Permissions table (from Spatie Permission)
type Permission struct {
	ID        uint       `db:"id" json:"id"`
	Name      string     `db:"name" json:"name"`
	CreatedAt time.Time  `db:"created_at" json:"created_at"`
	UpdatedAt time.Time  `db:"updated_at" json:"updated_at"`
	DeletedAt *time.Time `db:"deleted_at" json:"deleted_at"`
	CreatedBy *uint      `db:"created_by" json:"created_by"`

	// Relationships (populated separately)
	Roles []Role `json:"roles,omitempty"`
}

// RoleHasPermissions table (pivot table)
type RoleHasPermission struct {
	PermissionID uint `db:"permission_id" json:"permission_id"`
	RoleID       uint `db:"role_id" json:"role_id"`

	// Relationships (populated separately)
	Permission Permission `json:"permission,omitempty"`
	Role       Role       `json:"role,omitempty"`
}