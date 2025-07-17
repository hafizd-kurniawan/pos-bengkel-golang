package models

import (
	"time"

	"gorm.io/gorm"
)

// Users table
type User struct {
	UserID    uint           `gorm:"primaryKey;autoIncrement" json:"user_id"`
	Name      string         `gorm:"size:255;not null" json:"name"`
	Email     string         `gorm:"size:255;unique;not null" json:"email"`
	Password  string         `gorm:"size:255;not null" json:"-"`
	OutletID  *uint          `gorm:"index" json:"outlet_id"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at"`
	CreatedBy *uint          `json:"created_by"`

	// Relationships
	Outlet *Outlet `gorm:"foreignKey:OutletID" json:"outlet,omitempty"`
}

// Outlets table
type Outlet struct {
	OutletID     uint           `gorm:"primaryKey;autoIncrement" json:"outlet_id"`
	OutletName   string         `gorm:"size:255;not null" json:"outlet_name"`
	BranchType   string         `gorm:"size:50;not null" json:"branch_type"`
	City         string         `gorm:"size:100;not null" json:"city"`
	Address      *string        `gorm:"type:text" json:"address"`
	PhoneNumber  *string        `gorm:"size:20" json:"phone_number"`
	Status       StatusUmum     `gorm:"not null;default:'Aktif'" json:"status"`
	CreatedAt    time.Time      `json:"created_at"`
	UpdatedAt    time.Time      `json:"updated_at"`
	DeletedAt    gorm.DeletedAt `gorm:"index" json:"deleted_at"`
	CreatedBy    *uint          `json:"created_by"`

	// Relationships
	Users []User `gorm:"foreignKey:OutletID" json:"users,omitempty"`
}

// Roles table (from Spatie Permission)
type Role struct {
	ID        uint           `gorm:"primaryKey;autoIncrement" json:"id"`
	Name      string         `gorm:"size:255;not null" json:"name"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at"`
	CreatedBy *uint          `json:"created_by"`

	// Relationships
	Permissions []Permission `gorm:"many2many:role_has_permissions" json:"permissions,omitempty"`
}

// Permissions table (from Spatie Permission)
type Permission struct {
	ID        uint           `gorm:"primaryKey;autoIncrement" json:"id"`
	Name      string         `gorm:"size:255;not null" json:"name"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at"`
	CreatedBy *uint          `json:"created_by"`

	// Relationships
	Roles []Role `gorm:"many2many:role_has_permissions" json:"roles,omitempty"`
}

// RoleHasPermissions table (pivot table)
type RoleHasPermission struct {
	PermissionID uint `gorm:"primaryKey" json:"permission_id"`
	RoleID       uint `gorm:"primaryKey" json:"role_id"`

	// Relationships
	Permission Permission `gorm:"foreignKey:PermissionID" json:"permission,omitempty"`
	Role       Role       `gorm:"foreignKey:RoleID" json:"role,omitempty"`
}