package models

import (
	"time"

	"gorm.io/gorm"
)

// Transactions table
type Transaction struct {
	TransactionID   uint              `gorm:"primaryKey;autoIncrement" json:"transaction_id"`
	InvoiceNumber   string            `gorm:"size:255;unique;not null" json:"invoice_number"`
	TransactionDate time.Time         `gorm:"not null" json:"transaction_date"`
	UserID          uint              `gorm:"not null;index" json:"user_id"`
	CustomerID      *uint             `gorm:"index" json:"customer_id"`
	OutletID        uint              `gorm:"not null;index" json:"outlet_id"`
	TransactionType string            `gorm:"size:255;not null" json:"transaction_type"`
	Status          TransactionStatus `gorm:"not null;default:'sukses'" json:"status"`
	CreatedAt       time.Time         `json:"created_at"`
	UpdatedAt       time.Time         `json:"updated_at"`
	DeletedAt       gorm.DeletedAt    `gorm:"index" json:"deleted_at"`
	CreatedBy       *uint             `json:"created_by"`

	// Relationships
	User               *User                `gorm:"foreignKey:UserID" json:"user,omitempty"`
	Customer           *Customer            `gorm:"foreignKey:CustomerID" json:"customer,omitempty"`
	Outlet             *Outlet              `gorm:"foreignKey:OutletID" json:"outlet,omitempty"`
	TransactionDetails []TransactionDetail  `gorm:"foreignKey:TransactionID" json:"transaction_details,omitempty"`
	Payments           []Payment            `gorm:"foreignKey:TransactionID" json:"payments,omitempty"`
}

// TransactionDetails table
type TransactionDetail struct {
	DetailID         uint           `gorm:"primaryKey;autoIncrement" json:"detail_id"`
	TransactionType  string         `gorm:"size:255;not null" json:"transaction_type"`
	TransactionID    uint           `gorm:"not null;index" json:"transaction_id"`
	ProductID        *uint          `gorm:"index" json:"product_id"`
	SerialNumberID   *uint          `gorm:"index" json:"serial_number_id"`
	Quantity         int            `gorm:"not null" json:"quantity"`
	UnitPrice        float64        `gorm:"type:decimal(15,2);not null" json:"unit_price"`
	TotalPrice       float64        `gorm:"type:decimal(15,2);not null" json:"total_price"`
	CreatedAt        time.Time      `json:"created_at"`
	UpdatedAt        time.Time      `json:"updated_at"`
	DeletedAt        gorm.DeletedAt `gorm:"index" json:"deleted_at"`
	CreatedBy        *uint          `json:"created_by"`

	// Relationships
	Transaction    *Transaction         `gorm:"foreignKey:TransactionID" json:"transaction,omitempty"`
	Product        *Product             `gorm:"foreignKey:ProductID" json:"product,omitempty"`
	SerialNumber   *ProductSerialNumber `gorm:"foreignKey:SerialNumberID" json:"serial_number,omitempty"`
}

// PurchaseOrders table
type PurchaseOrder struct {
	PurchaseOrderID uint            `gorm:"primaryKey;autoIncrement" json:"purchase_order_id"`
	POCode          string          `gorm:"size:50;unique;not null" json:"po_code"`
	SupplierID      uint            `gorm:"not null;index" json:"supplier_id"`
	OutletID        uint            `gorm:"not null;index" json:"outlet_id"`
	PODate          time.Time       `gorm:"type:date;not null" json:"po_date"`
	TotalAmount     float64         `gorm:"type:decimal(15,2);not null" json:"total_amount"`
	AmountPaid      float64         `gorm:"type:decimal(15,2);not null;default:0" json:"amount_paid"`
	ChangeAmount    float64         `gorm:"type:decimal(15,2);not null;default:0" json:"change_amount"`
	PaymentType     PaymentTypeEnum `gorm:"not null" json:"payment_type"`
	Status          PurchaseStatus  `gorm:"not null;default:'Selesai'" json:"status"`
	Notes           *string         `gorm:"type:text" json:"notes"`
	CreatedAt       time.Time       `json:"created_at"`
	UpdatedAt       time.Time       `json:"updated_at"`

	// Relationships
	Supplier              *Supplier              `gorm:"foreignKey:SupplierID" json:"supplier,omitempty"`
	Outlet                *Outlet                `gorm:"foreignKey:OutletID" json:"outlet,omitempty"`
	PurchaseOrderDetails  []PurchaseOrderDetail  `gorm:"foreignKey:PurchaseOrderID" json:"purchase_order_details,omitempty"`
}

// PurchaseOrderDetails table
type PurchaseOrderDetail struct {
	DetailID        uint    `gorm:"primaryKey;autoIncrement" json:"detail_id"`
	PurchaseOrderID uint    `gorm:"not null;index" json:"purchase_order_id"`
	ProductID       uint    `gorm:"not null;index" json:"product_id"`
	Quantity        int     `gorm:"not null" json:"quantity"`
	CostPrice       float64 `gorm:"type:decimal(15,2);not null" json:"cost_price"`

	// Relationships
	PurchaseOrder *PurchaseOrder `gorm:"foreignKey:PurchaseOrderID" json:"purchase_order,omitempty"`
	Product       *Product       `gorm:"foreignKey:ProductID" json:"product,omitempty"`
}

// VehiclePurchases table
type VehiclePurchase struct {
	PurchaseID      uint           `gorm:"primaryKey;autoIncrement" json:"purchase_id"`
	PurchaseCode    string         `gorm:"size:50;unique;not null" json:"purchase_code"`
	CustomerID      *uint          `gorm:"index" json:"customer_id"`
	UserID          uint           `gorm:"not null;index" json:"user_id"`
	OutletID        uint           `gorm:"not null;index" json:"outlet_id"`
	PurchaseDate    time.Time      `gorm:"type:date;not null" json:"purchase_date"`
	PurchasePrice   float64        `gorm:"type:decimal(15,2);not null" json:"purchase_price"`
	VehicleSnapshot string         `gorm:"type:text;not null" json:"vehicle_snapshot"`
	Notes           *string        `gorm:"type:text" json:"notes"`
	CreatedAt       time.Time      `json:"created_at"`
	UpdatedAt       time.Time      `json:"updated_at"`
	DeletedAt       gorm.DeletedAt `gorm:"index" json:"deleted_at"`
	CreatedBy       *uint          `json:"created_by"`

	// Relationships
	Customer *Customer `gorm:"foreignKey:CustomerID" json:"customer,omitempty"`
	User     *User     `gorm:"foreignKey:UserID" json:"user,omitempty"`
	Outlet   *Outlet   `gorm:"foreignKey:OutletID" json:"outlet,omitempty"`
}