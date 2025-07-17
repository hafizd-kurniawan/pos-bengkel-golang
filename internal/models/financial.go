package models

import (
	"time"

	"gorm.io/gorm"
)

// PaymentMethods table
type PaymentMethod struct {
	MethodID  uint           `gorm:"primaryKey;autoIncrement" json:"method_id"`
	Name      string         `gorm:"size:100;not null" json:"name"`
	Status    StatusUmum     `gorm:"not null;default:'Aktif'" json:"status"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at"`
	CreatedBy *uint          `json:"created_by"`

	// Relationships
	Payments []Payment `gorm:"foreignKey:MethodID" json:"payments,omitempty"`
}

// Payments table
type Payment struct {
	PaymentID     uint              `gorm:"primaryKey;autoIncrement" json:"payment_id"`
	TransactionID uint              `gorm:"not null;index" json:"transaction_id"`
	MethodID      uint              `gorm:"not null;index" json:"method_id"`
	Amount        float64           `gorm:"type:decimal(15,2);not null" json:"amount"`
	Status        TransactionStatus `gorm:"not null;default:'sukses'" json:"status"`
	PaymentDate   *time.Time        `json:"payment_date"`
	CreatedAt     time.Time         `json:"created_at"`
	UpdatedAt     time.Time         `json:"updated_at"`
	DeletedAt     gorm.DeletedAt    `gorm:"index" json:"deleted_at"`
	CreatedBy     *uint             `json:"created_by"`

	// Relationships
	Transaction   *Transaction   `gorm:"foreignKey:TransactionID" json:"transaction,omitempty"`
	PaymentMethod *PaymentMethod `gorm:"foreignKey:MethodID" json:"payment_method,omitempty"`
}

// AccountsPayables table (Hutang)
type AccountsPayable struct {
	PayableID       uint           `gorm:"primaryKey;autoIncrement" json:"payable_id"`
	PurchaseOrderID uint           `gorm:"not null;index" json:"purchase_order_id"`
	SupplierID      uint           `gorm:"not null;index" json:"supplier_id"`
	TotalAmount     float64        `gorm:"type:decimal(15,2);not null" json:"total_amount"`
	AmountPaid      float64        `gorm:"type:decimal(15,2);not null;default:0" json:"amount_paid"`
	DueDate         time.Time      `gorm:"type:date;not null" json:"due_date"`
	Status          APARStatus     `gorm:"not null;default:'Belum Lunas'" json:"status"`
	CreatedAt       time.Time      `json:"created_at"`
	UpdatedAt       time.Time      `json:"updated_at"`
	DeletedAt       gorm.DeletedAt `gorm:"index" json:"deleted_at"`
	CreatedBy       *uint          `json:"created_by"`

	// Relationships
	PurchaseOrder    *PurchaseOrder    `gorm:"foreignKey:PurchaseOrderID" json:"purchase_order,omitempty"`
	Supplier         *Supplier         `gorm:"foreignKey:SupplierID" json:"supplier,omitempty"`
	PayablePayments  []PayablePayment  `gorm:"foreignKey:PayableID" json:"payable_payments,omitempty"`
}

// PayablePayments table (Cicilan Hutang)
type PayablePayment struct {
	PaymentID   uint           `gorm:"primaryKey;autoIncrement" json:"payment_id"`
	PayableID   uint           `gorm:"not null;index" json:"payable_id"`
	PaymentDate time.Time      `gorm:"type:date;not null" json:"payment_date"`
	Amount      float64        `gorm:"type:decimal(15,2);not null" json:"amount"`
	Notes       *string        `gorm:"type:text" json:"notes"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"deleted_at"`
	CreatedBy   *uint          `json:"created_by"`

	// Relationships
	AccountsPayable *AccountsPayable `gorm:"foreignKey:PayableID" json:"accounts_payable,omitempty"`
}

// AccountsReceivables table (Piutang)
type AccountsReceivable struct {
	ReceivableID  uint           `gorm:"primaryKey;autoIncrement" json:"receivable_id"`
	TransactionID uint           `gorm:"not null;index" json:"transaction_id"`
	CustomerID    uint           `gorm:"not null;index" json:"customer_id"`
	TotalAmount   float64        `gorm:"type:decimal(15,2);not null" json:"total_amount"`
	AmountPaid    float64        `gorm:"type:decimal(15,2);not null;default:0" json:"amount_paid"`
	DueDate       time.Time      `gorm:"type:date;not null" json:"due_date"`
	Status        APARStatus     `gorm:"not null;default:'Belum Lunas'" json:"status"`
	CreatedAt     time.Time      `json:"created_at"`
	UpdatedAt     time.Time      `json:"updated_at"`
	DeletedAt     gorm.DeletedAt `gorm:"index" json:"deleted_at"`
	CreatedBy     *uint          `json:"created_by"`

	// Relationships
	Transaction         *Transaction         `gorm:"foreignKey:TransactionID" json:"transaction,omitempty"`
	Customer            *Customer            `gorm:"foreignKey:CustomerID" json:"customer,omitempty"`
	ReceivablePayments  []ReceivablePayment  `gorm:"foreignKey:ReceivableID" json:"receivable_payments,omitempty"`
}

// ReceivablePayments table (Cicilan Piutang)
type ReceivablePayment struct {
	PaymentID    uint           `gorm:"primaryKey;autoIncrement" json:"payment_id"`
	ReceivableID uint           `gorm:"not null;index" json:"receivable_id"`
	PaymentDate  time.Time      `gorm:"type:date;not null" json:"payment_date"`
	Amount       float64        `gorm:"type:decimal(15,2);not null" json:"amount"`
	Notes        *string        `gorm:"type:text" json:"notes"`
	CreatedAt    time.Time      `json:"created_at"`
	UpdatedAt    time.Time      `json:"updated_at"`
	DeletedAt    gorm.DeletedAt `gorm:"index" json:"deleted_at"`
	CreatedBy    *uint          `json:"created_by"`

	// Relationships
	AccountsReceivable *AccountsReceivable `gorm:"foreignKey:ReceivableID" json:"accounts_receivable,omitempty"`
}

// CashFlows table
type CashFlow struct {
	CashFlowID uint         `gorm:"primaryKey;autoIncrement" json:"cash_flow_id"`
	Type       CashFlowType `gorm:"not null" json:"type"`
	Source     string       `gorm:"size:255;not null" json:"source"`
	Amount     float64      `gorm:"type:decimal(15,2);not null" json:"amount"`
	Date       time.Time    `gorm:"type:date;not null" json:"date"`
	Notes      *string      `gorm:"type:text" json:"notes"`
	UserID     uint         `gorm:"not null;index" json:"user_id"`
	CreatedAt  time.Time    `json:"created_at"`
	UpdatedAt  time.Time    `json:"updated_at"`
	DeletedAt  gorm.DeletedAt `gorm:"index" json:"deleted_at"`
	CreatedBy  *uint        `json:"created_by"`

	// Relationships
	User *User `gorm:"foreignKey:UserID" json:"user,omitempty"`
}