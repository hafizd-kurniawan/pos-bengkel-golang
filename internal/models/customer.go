package models

import (
	"time"

	"gorm.io/gorm"
)

// Customers table
type Customer struct {
	CustomerID  uint           `gorm:"primaryKey;autoIncrement" json:"customer_id"`
	Name        string         `gorm:"size:255;not null" json:"name"`
	PhoneNumber string         `gorm:"size:20;unique;not null" json:"phone_number"`
	Address     *string        `gorm:"type:text" json:"address"`
	Status      StatusUmum     `gorm:"not null;default:'Aktif'" json:"status"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"deleted_at"`
	CreatedBy   *uint          `json:"created_by"`

	// Relationships
	Vehicles []CustomerVehicle `gorm:"foreignKey:CustomerID" json:"vehicles,omitempty"`
}

// CustomerVehicles table
type CustomerVehicle struct {
	VehicleID      uint           `gorm:"primaryKey;autoIncrement" json:"vehicle_id"`
	CustomerID     uint           `gorm:"not null;index" json:"customer_id"`
	PlateNumber    string         `gorm:"size:20;unique;not null" json:"plate_number"`
	Brand          string         `gorm:"size:100;not null" json:"brand"`
	Model          string         `gorm:"size:100;not null" json:"model"`
	Type           string         `gorm:"size:100;not null" json:"type"`
	ProductionYear int            `gorm:"not null" json:"production_year"`
	ChassisNumber  string         `gorm:"size:100;unique;not null" json:"chassis_number"`
	EngineNumber   string         `gorm:"size:100;unique;not null" json:"engine_number"`
	Color          string         `gorm:"size:50;not null" json:"color"`
	Notes          *string        `gorm:"type:text" json:"notes"`
	CreatedAt      time.Time      `json:"created_at"`
	UpdatedAt      time.Time      `json:"updated_at"`
	DeletedAt      gorm.DeletedAt `gorm:"index" json:"deleted_at"`
	CreatedBy      *uint          `json:"created_by"`

	// Relationships
	Customer *Customer `gorm:"foreignKey:CustomerID" json:"customer,omitempty"`
}