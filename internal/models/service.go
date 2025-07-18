package models

import (
	"database/sql/driver"
	"time"
)

// Services table
type Service struct {
	ServiceID         uint       `db:"service_id" json:"service_id"`
	ServiceCode       string     `db:"service_code" json:"service_code"`
	Name              string     `db:"name" json:"name"`
	ServiceCategoryID uint       `db:"service_category_id" json:"service_category_id"`
	Fee               float64    `db:"fee" json:"fee"`
	Status            StatusUmum `db:"status" json:"status"`
	CreatedAt         time.Time  `db:"created_at" json:"created_at"`
	UpdatedAt         time.Time  `db:"updated_at" json:"updated_at"`
	DeletedAt         *time.Time `db:"deleted_at" json:"deleted_at"`
	CreatedBy         *uint      `db:"created_by" json:"created_by"`

	// Relationships (populated separately)
	ServiceCategory *ServiceCategory `json:"service_category,omitempty"`
}

// ServiceCategories table
type ServiceCategory struct {
	ServiceCategoryID uint       `db:"service_category_id" json:"service_category_id"`
	Name              string     `db:"name" json:"name"`
	Status            StatusUmum `db:"status" json:"status"`
	CreatedAt         time.Time  `db:"created_at" json:"created_at"`
	UpdatedAt         time.Time  `db:"updated_at" json:"updated_at"`
	DeletedAt         *time.Time `db:"deleted_at" json:"deleted_at"`
	CreatedBy         *uint      `db:"created_by" json:"created_by"`

	// Relationships (populated separately)
	Services []Service `json:"services,omitempty"`
}

// Implement database/sql interfaces for ServiceStatusEnum
func (s *ServiceStatusEnum) Scan(value interface{}) error {
	if value == nil {
		*s = ""
		return nil
	}
	if str, ok := value.(string); ok {
		*s = ServiceStatusEnum(str)
		return nil
	}
	return nil
}

func (s ServiceStatusEnum) Value() (driver.Value, error) {
	return string(s), nil
}