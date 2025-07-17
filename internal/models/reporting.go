package models

import (
	"time"

	"gorm.io/gorm"
)

// Reports table
type Report struct {
	ReportID         uint           `gorm:"primaryKey;autoIncrement" json:"report_id"`
	ReportName       string         `gorm:"size:255;unique;not null" json:"report_name"`
	ReportType       ReportTypeEnum `gorm:"not null" json:"report_type"`
	StartDate        time.Time      `gorm:"type:date;not null" json:"start_date"`
	EndDate          time.Time      `gorm:"type:date;not null" json:"end_date"`
	OutletID         *uint          `gorm:"index" json:"outlet_id"`
	UserID           uint           `gorm:"not null;index" json:"user_id"`
	GeneratedFilePath *string       `gorm:"size:255" json:"generated_file_path"`
	Status           ReportStatus   `gorm:"default:'Pending'" json:"status"`
	CreatedAt        time.Time      `json:"created_at"`
	UpdatedAt        time.Time      `json:"updated_at"`
	DeletedAt        gorm.DeletedAt `gorm:"index" json:"deleted_at"`
	CreatedBy        *uint          `json:"created_by"`

	// Relationships
	Outlet *Outlet `gorm:"foreignKey:OutletID" json:"outlet,omitempty"`
	User   *User   `gorm:"foreignKey:UserID" json:"user,omitempty"`
}

// Promotions table
type Promotion struct {
	PromotionID   uint           `gorm:"primaryKey;autoIncrement" json:"promotion_id"`
	PromotionName string         `gorm:"size:255;not null" json:"promotion_name"`
	StartDate     time.Time      `gorm:"not null" json:"start_date"`
	EndDate       time.Time      `gorm:"not null" json:"end_date"`
	Type          string         `gorm:"type:enum('percentage','fixed');not null" json:"type"`
	Value         float64        `gorm:"type:decimal(15,2);not null" json:"value"`
	CreatedAt     time.Time      `json:"created_at"`
	UpdatedAt     time.Time      `json:"updated_at"`
	DeletedAt     gorm.DeletedAt `gorm:"index" json:"deleted_at"`
	CreatedBy     *uint          `json:"created_by"`
}