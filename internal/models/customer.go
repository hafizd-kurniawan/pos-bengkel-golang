package models

import (
	"database/sql/driver"
	"time"
)

// Customers table
type Customer struct {
	CustomerID  uint       `db:"customer_id" json:"customer_id"`
	Name        string     `db:"name" json:"name"`
	PhoneNumber string     `db:"phone_number" json:"phone_number"`
	Address     *string    `db:"address" json:"address"`
	Status      StatusUmum `db:"status" json:"status"`
	CreatedAt   time.Time  `db:"created_at" json:"created_at"`
	UpdatedAt   time.Time  `db:"updated_at" json:"updated_at"`
	DeletedAt   *time.Time `db:"deleted_at" json:"deleted_at"`
	CreatedBy   *uint      `db:"created_by" json:"created_by"`

	// Relationships (populated separately)
	Vehicles []CustomerVehicle `json:"vehicles,omitempty"`
}

// CustomerVehicles table
type CustomerVehicle struct {
	VehicleID      uint       `db:"vehicle_id" json:"vehicle_id"`
	CustomerID     uint       `db:"customer_id" json:"customer_id"`
	PlateNumber    string     `db:"plate_number" json:"plate_number"`
	Brand          string     `db:"brand" json:"brand"`
	Model          string     `db:"model" json:"model"`
	Type           string     `db:"type" json:"type"`
	ProductionYear int        `db:"production_year" json:"production_year"`
	ChassisNumber  string     `db:"chassis_number" json:"chassis_number"`
	EngineNumber   string     `db:"engine_number" json:"engine_number"`
	Color          string     `db:"color" json:"color"`
	Notes          *string    `db:"notes" json:"notes"`
	CreatedAt      time.Time  `db:"created_at" json:"created_at"`
	UpdatedAt      time.Time  `db:"updated_at" json:"updated_at"`
	DeletedAt      *time.Time `db:"deleted_at" json:"deleted_at"`
	CreatedBy      *uint      `db:"created_by" json:"created_by"`

	// Relationships (populated separately)
	Customer *Customer `json:"customer,omitempty"`
}

// Implement database/sql interfaces for StatusUmum
func (s *StatusUmum) Scan(value interface{}) error {
	if value == nil {
		*s = ""
		return nil
	}
	if str, ok := value.(string); ok {
		*s = StatusUmum(str)
		return nil
	}
	return nil
}

func (s StatusUmum) Value() (driver.Value, error) {
	return string(s), nil
}