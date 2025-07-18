package models

import (
	"database/sql/driver"
	"time"
)

// Products table
type Product struct {
	ProductID          uint               `db:"product_id" json:"product_id"`
	ProductName        string             `db:"product_name" json:"product_name"`
	ProductDescription *string            `db:"product_description" json:"product_description"`
	ProductImage       *string            `db:"product_image" json:"product_image"`
	CostPrice          float64            `db:"cost_price" json:"cost_price"`
	SellingPrice       float64            `db:"selling_price" json:"selling_price"`
	Stock              int                `db:"stock" json:"stock"`
	SKU                *string            `db:"sku" json:"sku"`
	Barcode            *string            `db:"barcode" json:"barcode"`
	HasSerialNumber    bool               `db:"has_serial_number" json:"has_serial_number"`
	ShelfLocation      *string            `db:"shelf_location" json:"shelf_location"`
	UsageStatus        ProductUsageStatus `db:"usage_status" json:"usage_status"`
	IsActive           bool               `db:"is_active" json:"is_active"`
	CategoryID         *uint              `db:"category_id" json:"category_id"`
	SupplierID         *uint              `db:"supplier_id" json:"supplier_id"`
	UnitTypeID         *uint              `db:"unit_type_id" json:"unit_type_id"`
	SourceableID       *uint              `db:"sourceable_id" json:"sourceable_id"`
	SourceableType     *string            `db:"sourceable_type" json:"sourceable_type"`
	CreatedAt          time.Time          `db:"created_at" json:"created_at"`
	UpdatedAt          time.Time          `db:"updated_at" json:"updated_at"`
	DeletedAt          *time.Time         `db:"deleted_at" json:"deleted_at"`
	CreatedBy          *uint              `db:"created_by" json:"created_by"`

	// Relationships (populated separately)
	Category      *Category             `json:"category,omitempty"`
	Supplier      *Supplier             `json:"supplier,omitempty"`
	UnitType      *UnitType             `json:"unit_type,omitempty"`
	SerialNumbers []ProductSerialNumber `json:"serial_numbers,omitempty"`
}

// ProductSerialNumbers table
type ProductSerialNumber struct {
	SerialNumberID uint       `db:"serial_number_id" json:"serial_number_id"`
	ProductID      uint       `db:"product_id" json:"product_id"`
	SerialNumber   string     `db:"serial_number" json:"serial_number"`
	Status         SNStatus   `db:"status" json:"status"`
	CreatedAt      time.Time  `db:"created_at" json:"created_at"`
	UpdatedAt      time.Time  `db:"updated_at" json:"updated_at"`
	DeletedAt      *time.Time `db:"deleted_at" json:"deleted_at"`
	CreatedBy      *uint      `db:"created_by" json:"created_by"`

	// Relationships (populated separately)
	Product *Product `json:"product,omitempty"`
}

// Categories table
type Category struct {
	CategoryID uint       `db:"category_id" json:"category_id"`
	Name       string     `db:"name" json:"name"`
	Status     StatusUmum `db:"status" json:"status"`
	CreatedAt  time.Time  `db:"created_at" json:"created_at"`
	UpdatedAt  time.Time  `db:"updated_at" json:"updated_at"`
	DeletedAt  *time.Time `db:"deleted_at" json:"deleted_at"`
	CreatedBy  *uint      `db:"created_by" json:"created_by"`

	// Relationships (populated separately)
	Products []Product `json:"products,omitempty"`
}

// Suppliers table
type Supplier struct {
	SupplierID        uint       `db:"supplier_id" json:"supplier_id"`
	SupplierName      string     `db:"supplier_name" json:"supplier_name"`
	ContactPersonName string     `db:"contact_person_name" json:"contact_person_name"`
	PhoneNumber       string     `db:"phone_number" json:"phone_number"`
	Address           *string    `db:"address" json:"address"`
	Status            StatusUmum `db:"status" json:"status"`
	CreatedAt         time.Time  `db:"created_at" json:"created_at"`
	UpdatedAt         time.Time  `db:"updated_at" json:"updated_at"`
	DeletedAt         *time.Time `db:"deleted_at" json:"deleted_at"`
	CreatedBy         *uint      `db:"created_by" json:"created_by"`

	// Relationships (populated separately)
	Products []Product `json:"products,omitempty"`
}

// UnitTypes table
type UnitType struct {
	UnitTypeID uint       `db:"unit_type_id" json:"unit_type_id"`
	Name       string     `db:"name" json:"name"`
	Status     StatusUmum `db:"status" json:"status"`
	CreatedAt  time.Time  `db:"created_at" json:"created_at"`
	UpdatedAt  time.Time  `db:"updated_at" json:"updated_at"`
	DeletedAt  *time.Time `db:"deleted_at" json:"deleted_at"`
	CreatedBy  *uint      `db:"created_by" json:"created_by"`

	// Relationships (populated separately)
	Products []Product `json:"products,omitempty"`
}

// Implement database/sql interfaces for ProductUsageStatus
func (p *ProductUsageStatus) Scan(value interface{}) error {
	if value == nil {
		*p = ""
		return nil
	}
	if str, ok := value.(string); ok {
		*p = ProductUsageStatus(str)
		return nil
	}
	return nil
}

func (p ProductUsageStatus) Value() (driver.Value, error) {
	return string(p), nil
}

// Implement database/sql interfaces for SNStatus
func (s *SNStatus) Scan(value interface{}) error {
	if value == nil {
		*s = ""
		return nil
	}
	if str, ok := value.(string); ok {
		*s = SNStatus(str)
		return nil
	}
	return nil
}

func (s SNStatus) Value() (driver.Value, error) {
	return string(s), nil
}