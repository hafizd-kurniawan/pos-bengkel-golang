package models

import (
	"time"

	"gorm.io/gorm"
)

// Products table
type Product struct {
	ProductID          uint               `gorm:"primaryKey;autoIncrement" json:"product_id"`
	ProductName        string             `gorm:"size:255;not null" json:"product_name"`
	ProductDescription *string            `gorm:"type:text" json:"product_description"`
	ProductImage       *string            `gorm:"size:255" json:"product_image"`
	CostPrice          float64            `gorm:"type:decimal(15,2);not null" json:"cost_price"`
	SellingPrice       float64            `gorm:"type:decimal(15,2);not null" json:"selling_price"`
	Stock              int                `gorm:"not null;default:0" json:"stock"`
	SKU                *string            `gorm:"size:100;unique" json:"sku"`
	Barcode            *string            `gorm:"size:100;unique" json:"barcode"`
	HasSerialNumber    bool               `gorm:"not null;default:false" json:"has_serial_number"`
	ShelfLocation      *string            `gorm:"size:100" json:"shelf_location"`
	UsageStatus        ProductUsageStatus `gorm:"not null" json:"usage_status"`
	IsActive           bool               `gorm:"not null;default:true" json:"is_active"`
	CategoryID         *uint              `gorm:"index" json:"category_id"`
	SupplierID         *uint              `gorm:"index" json:"supplier_id"`
	UnitTypeID         *uint              `gorm:"index" json:"unit_type_id"`
	SourceableID       *uint              `json:"sourceable_id"`
	SourceableType     *string            `json:"sourceable_type"`
	CreatedAt          time.Time          `json:"created_at"`
	UpdatedAt          time.Time          `json:"updated_at"`
	DeletedAt          gorm.DeletedAt     `gorm:"index" json:"deleted_at"`
	CreatedBy          *uint              `json:"created_by"`

	// Relationships
	Category     *Category              `gorm:"foreignKey:CategoryID" json:"category,omitempty"`
	Supplier     *Supplier              `gorm:"foreignKey:SupplierID" json:"supplier,omitempty"`
	UnitType     *UnitType              `gorm:"foreignKey:UnitTypeID" json:"unit_type,omitempty"`
	SerialNumbers []ProductSerialNumber `gorm:"foreignKey:ProductID" json:"serial_numbers,omitempty"`
}

// ProductSerialNumbers table
type ProductSerialNumber struct {
	SerialNumberID uint           `gorm:"primaryKey;autoIncrement" json:"serial_number_id"`
	ProductID      uint           `gorm:"not null;index" json:"product_id"`
	SerialNumber   string         `gorm:"size:255;unique;not null" json:"serial_number"`
	Status         SNStatus       `gorm:"not null;default:'Tersedia'" json:"status"`
	CreatedAt      time.Time      `json:"created_at"`
	UpdatedAt      time.Time      `json:"updated_at"`
	DeletedAt      gorm.DeletedAt `gorm:"index" json:"deleted_at"`
	CreatedBy      *uint          `json:"created_by"`

	// Relationships
	Product *Product `gorm:"foreignKey:ProductID" json:"product,omitempty"`
}

// Categories table
type Category struct {
	CategoryID uint           `gorm:"primaryKey;autoIncrement" json:"category_id"`
	Name       string         `gorm:"size:255;not null" json:"name"`
	Status     StatusUmum     `gorm:"not null;default:'Aktif'" json:"status"`
	CreatedAt  time.Time      `json:"created_at"`
	UpdatedAt  time.Time      `json:"updated_at"`
	DeletedAt  gorm.DeletedAt `gorm:"index" json:"deleted_at"`
	CreatedBy  *uint          `json:"created_by"`

	// Relationships
	Products []Product `gorm:"foreignKey:CategoryID" json:"products,omitempty"`
}

// Suppliers table
type Supplier struct {
	SupplierID        uint           `gorm:"primaryKey;autoIncrement" json:"supplier_id"`
	SupplierName      string         `gorm:"size:255;not null" json:"supplier_name"`
	ContactPersonName string         `gorm:"size:255;not null" json:"contact_person_name"`
	PhoneNumber       string         `gorm:"size:20;not null" json:"phone_number"`
	Address           *string        `gorm:"type:text" json:"address"`
	Status            StatusUmum     `gorm:"not null;default:'Aktif'" json:"status"`
	CreatedAt         time.Time      `json:"created_at"`
	UpdatedAt         time.Time      `json:"updated_at"`
	DeletedAt         gorm.DeletedAt `gorm:"index" json:"deleted_at"`
	CreatedBy         *uint          `json:"created_by"`

	// Relationships
	Products []Product `gorm:"foreignKey:SupplierID" json:"products,omitempty"`
}

// UnitTypes table
type UnitType struct {
	UnitTypeID uint           `gorm:"primaryKey;autoIncrement" json:"unit_type_id"`
	Name       string         `gorm:"size:50;not null" json:"name"`
	Status     StatusUmum     `gorm:"not null;default:'Aktif'" json:"status"`
	CreatedAt  time.Time      `json:"created_at"`
	UpdatedAt  time.Time      `json:"updated_at"`
	DeletedAt  gorm.DeletedAt `gorm:"index" json:"deleted_at"`
	CreatedBy  *uint          `json:"created_by"`

	// Relationships
	Products []Product `gorm:"foreignKey:UnitTypeID" json:"products,omitempty"`
}