package models

import (
	"time"

	"gorm.io/gorm"
)

// Services table
type Service struct {
	ServiceID         uint           `gorm:"primaryKey;autoIncrement" json:"service_id"`
	ServiceCode       string         `gorm:"size:50;unique;not null" json:"service_code"`
	Name              string         `gorm:"size:255;not null" json:"name"`
	ServiceCategoryID uint           `gorm:"not null;index" json:"service_category_id"`
	Fee               float64        `gorm:"type:decimal(15,2);not null" json:"fee"`
	Status            StatusUmum     `gorm:"not null;default:'Aktif'" json:"status"`
	CreatedAt         time.Time      `json:"created_at"`
	UpdatedAt         time.Time      `json:"updated_at"`
	DeletedAt         gorm.DeletedAt `gorm:"index" json:"deleted_at"`
	CreatedBy         *uint          `json:"created_by"`

	// Relationships
	ServiceCategory *ServiceCategory `gorm:"foreignKey:ServiceCategoryID" json:"service_category,omitempty"`
}

// ServiceCategories table
type ServiceCategory struct {
	ServiceCategoryID uint           `gorm:"primaryKey;autoIncrement" json:"service_category_id"`
	Name              string         `gorm:"size:255;not null" json:"name"`
	Status            StatusUmum     `gorm:"not null;default:'Aktif'" json:"status"`
	CreatedAt         time.Time      `json:"created_at"`
	UpdatedAt         time.Time      `json:"updated_at"`
	DeletedAt         gorm.DeletedAt `gorm:"index" json:"deleted_at"`
	CreatedBy         *uint          `json:"created_by"`

	// Relationships
	Services []Service `gorm:"foreignKey:ServiceCategoryID" json:"services,omitempty"`
}