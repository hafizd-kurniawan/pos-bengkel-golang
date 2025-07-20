package models

import (
	"time"

	"gorm.io/gorm"
)

// Vehicle represents an enhanced vehicle model for trading operations
type Vehicle struct {
	VehicleID         uint                   `gorm:"primaryKey;autoIncrement" json:"vehicle_id"`
	CustomerID        *uint                  `gorm:"index" json:"customer_id"` // Current owner, nullable for showroom inventory
	PlateNumber       string                 `gorm:"size:20;unique;not null" json:"plate_number"`
	Brand             string                 `gorm:"size:100;not null" json:"brand"`
	Model             string                 `gorm:"size:100;not null" json:"model"`
	Type              string                 `gorm:"size:100;not null" json:"type"`
	ProductionYear    int                    `gorm:"not null" json:"production_year"`
	ChassisNumber     string                 `gorm:"size:100;unique;not null" json:"chassis_number"`
	EngineNumber      string                 `gorm:"size:100;unique;not null" json:"engine_number"`
	Color             string                 `gorm:"size:50;not null" json:"color"`
	Mileage           *int                   `json:"mileage"`
	FuelType          *string                `gorm:"size:20" json:"fuel_type"`
	Transmission      *string                `gorm:"size:20" json:"transmission"`
	OwnershipStatus   VehicleOwnershipStatus `gorm:"not null;default:'Customer'" json:"ownership_status"`
	ConditionStatus   VehicleConditionStatus `gorm:"not null;default:'Good'" json:"condition_status"`
	SaleStatus        VehicleSaleStatus      `gorm:"not null;default:'Not For Sale'" json:"sale_status"`
	PurchasePrice     *float64               `gorm:"type:decimal(15,2)" json:"purchase_price"`
	SellingPrice      *float64               `gorm:"type:decimal(15,2)" json:"selling_price"`
	EstimatedValue    *float64               `gorm:"type:decimal(15,2)" json:"estimated_value"`
	ConditionNotes    *string                `gorm:"type:text" json:"condition_notes"`
	InternalNotes     *string                `gorm:"type:text" json:"internal_notes"`
	CreatedAt         time.Time              `json:"created_at"`
	UpdatedAt         time.Time              `json:"updated_at"`
	DeletedAt         gorm.DeletedAt         `gorm:"index" json:"deleted_at"`
	CreatedBy         *uint                  `json:"created_by"`

	// Relationships
	Customer                 *Customer                  `gorm:"foreignKey:CustomerID" json:"customer,omitempty"`
	PurchaseTransactions     []VehiclePurchaseTransaction `gorm:"foreignKey:VehicleID" json:"purchase_transactions,omitempty"`
	ReconditioningJobs       []VehicleReconditioningJob `gorm:"foreignKey:VehicleID" json:"reconditioning_jobs,omitempty"`
	SalesTransactions        []VehicleSalesTransaction  `gorm:"foreignKey:VehicleID" json:"sales_transactions,omitempty"`
}

// VehiclePurchaseTransaction represents a transaction when buying a vehicle from a customer
type VehiclePurchaseTransaction struct {
	PurchaseTransactionID uint               `gorm:"primaryKey;autoIncrement" json:"purchase_transaction_id"`
	VehicleID             uint               `gorm:"not null;index" json:"vehicle_id"`
	CustomerID            uint               `gorm:"not null;index" json:"customer_id"`
	PurchasePrice         float64            `gorm:"type:decimal(15,2);not null" json:"purchase_price"`
	PurchaseDate          time.Time          `gorm:"not null" json:"purchase_date"`
	PaymentMethod         PaymentTypeEnum    `gorm:"not null" json:"payment_method"`
	TransactionStatus     TransactionStatus  `gorm:"not null;default:'pending'" json:"transaction_status"`
	EvaluationNotes       *string            `gorm:"type:text" json:"evaluation_notes"`
	PaymentReference      *string            `gorm:"size:255" json:"payment_reference"`
	CreatedAt             time.Time          `json:"created_at"`
	UpdatedAt             time.Time          `json:"updated_at"`
	DeletedAt             gorm.DeletedAt     `gorm:"index" json:"deleted_at"`
	CreatedBy             *uint              `json:"created_by"`

	// Relationships
	Vehicle  *Vehicle  `gorm:"foreignKey:VehicleID" json:"vehicle,omitempty"`
	Customer *Customer `gorm:"foreignKey:CustomerID" json:"customer,omitempty"`
}

// VehicleReconditioningJob represents a reconditioning job for a vehicle
type VehicleReconditioningJob struct {
	ReconditioningJobID uint                    `gorm:"primaryKey;autoIncrement" json:"reconditioning_job_id"`
	VehicleID           uint                    `gorm:"not null;index" json:"vehicle_id"`
	JobTitle            string                  `gorm:"size:255;not null" json:"job_title"`
	JobDescription      *string                 `gorm:"type:text" json:"job_description"`
	EstimatedCost       *float64                `gorm:"type:decimal(15,2)" json:"estimated_cost"`
	ActualCost          *float64                `gorm:"type:decimal(15,2)" json:"actual_cost"`
	StartDate           *time.Time              `json:"start_date"`
	CompletionDate      *time.Time              `json:"completion_date"`
	Status              ReconditioningJobStatus `gorm:"not null;default:'Pending'" json:"status"`
	AssignedTechnicianID *uint                  `gorm:"index" json:"assigned_technician_id"`
	Notes               *string                 `gorm:"type:text" json:"notes"`
	CreatedAt           time.Time               `json:"created_at"`
	UpdatedAt           time.Time               `json:"updated_at"`
	DeletedAt           gorm.DeletedAt          `gorm:"index" json:"deleted_at"`
	CreatedBy           *uint                   `json:"created_by"`

	// Relationships
	Vehicle             *Vehicle              `gorm:"foreignKey:VehicleID" json:"vehicle,omitempty"`
	AssignedTechnician  *User                 `gorm:"foreignKey:AssignedTechnicianID" json:"assigned_technician,omitempty"`
	ReconditioningDetails []ReconditioningDetail `gorm:"foreignKey:ReconditioningJobID" json:"reconditioning_details,omitempty"`
}

// ReconditioningDetail represents parts or services used in reconditioning
type ReconditioningDetail struct {
	DetailID            uint                     `gorm:"primaryKey;autoIncrement" json:"detail_id"`
	ReconditioningJobID uint                     `gorm:"not null;index" json:"reconditioning_job_id"`
	DetailType          ReconditioningDetailType `gorm:"not null" json:"detail_type"`
	ProductID           *uint                    `gorm:"index" json:"product_id"` // For parts
	ServiceID           *uint                    `gorm:"index" json:"service_id"` // For services
	Description         string                   `gorm:"size:255;not null" json:"description"`
	Quantity            int                      `gorm:"not null" json:"quantity"`
	UnitPrice           float64                  `gorm:"type:decimal(15,2);not null" json:"unit_price"`
	TotalPrice          float64                  `gorm:"type:decimal(15,2);not null" json:"total_price"`
	UsageDate           time.Time                `gorm:"not null" json:"usage_date"`
	Notes               *string                  `gorm:"type:text" json:"notes"`
	CreatedAt           time.Time                `json:"created_at"`
	UpdatedAt           time.Time                `json:"updated_at"`
	DeletedAt           gorm.DeletedAt           `gorm:"index" json:"deleted_at"`
	CreatedBy           *uint                    `json:"created_by"`

	// Relationships
	ReconditioningJob *VehicleReconditioningJob `gorm:"foreignKey:ReconditioningJobID" json:"reconditioning_job,omitempty"`
	Product           *Product                  `gorm:"foreignKey:ProductID" json:"product,omitempty"`
	Service           *Service                  `gorm:"foreignKey:ServiceID" json:"service,omitempty"`
}

// VehicleSalesTransaction represents a transaction when selling a vehicle to a customer
type VehicleSalesTransaction struct {
	SalesTransactionID   uint                 `gorm:"primaryKey;autoIncrement" json:"sales_transaction_id"`
	VehicleID            uint                 `gorm:"not null;index" json:"vehicle_id"`
	CustomerID           uint                 `gorm:"not null;index" json:"customer_id"`
	SalePrice            float64              `gorm:"type:decimal(15,2);not null" json:"sale_price"`
	DownPayment          *float64             `gorm:"type:decimal(15,2)" json:"down_payment"`
	SaleDate             time.Time            `gorm:"not null" json:"sale_date"`
	TransactionType      SalesTransactionType `gorm:"not null" json:"transaction_type"`
	PaymentMethod        PaymentTypeEnum      `gorm:"not null" json:"payment_method"`
	TransactionStatus    TransactionStatus    `gorm:"not null;default:'pending'" json:"transaction_status"`
	PaymentReference     *string              `gorm:"size:255" json:"payment_reference"`
	ProfitAmount         *float64             `gorm:"type:decimal(15,2)" json:"profit_amount"`
	SalesPersonID        *uint                `gorm:"index" json:"sales_person_id"`
	Notes                *string              `gorm:"type:text" json:"notes"`
	CreatedAt            time.Time            `json:"created_at"`
	UpdatedAt            time.Time            `json:"updated_at"`
	DeletedAt            gorm.DeletedAt       `gorm:"index" json:"deleted_at"`
	CreatedBy            *uint                `json:"created_by"`

	// Relationships
	Vehicle      *Vehicle             `gorm:"foreignKey:VehicleID" json:"vehicle,omitempty"`
	Customer     *Customer            `gorm:"foreignKey:CustomerID" json:"customer,omitempty"`
	SalesPerson  *User                `gorm:"foreignKey:SalesPersonID" json:"sales_person,omitempty"`
	Installments []VehicleInstallment `gorm:"foreignKey:SalesTransactionID" json:"installments,omitempty"`
}

// VehicleInstallment represents an installment plan for vehicle purchase
type VehicleInstallment struct {
	InstallmentID       uint              `gorm:"primaryKey;autoIncrement" json:"installment_id"`
	SalesTransactionID  uint              `gorm:"not null;index" json:"sales_transaction_id"`
	TotalAmount         float64           `gorm:"type:decimal(15,2);not null" json:"total_amount"`
	DownPayment         float64           `gorm:"type:decimal(15,2);not null" json:"down_payment"`
	InstallmentAmount   float64           `gorm:"type:decimal(15,2);not null" json:"installment_amount"`
	NumberOfInstallments int              `gorm:"not null" json:"number_of_installments"`
	InterestRate        *float64          `gorm:"type:decimal(5,2)" json:"interest_rate"`
	StartDate           time.Time         `gorm:"not null" json:"start_date"`
	EndDate             time.Time         `gorm:"not null" json:"end_date"`
	Status              InstallmentStatus `gorm:"not null;default:'Active'" json:"status"`
	RemainingBalance    float64           `gorm:"type:decimal(15,2);not null" json:"remaining_balance"`
	CreatedAt           time.Time         `json:"created_at"`
	UpdatedAt           time.Time         `json:"updated_at"`
	DeletedAt           gorm.DeletedAt    `gorm:"index" json:"deleted_at"`
	CreatedBy           *uint             `json:"created_by"`

	// Relationships
	SalesTransaction    *VehicleSalesTransaction `gorm:"foreignKey:SalesTransactionID" json:"sales_transaction,omitempty"`
	InstallmentPayments []InstallmentPayment     `gorm:"foreignKey:InstallmentID" json:"installment_payments,omitempty"`
}

// InstallmentPayment represents individual payments made towards an installment
type InstallmentPayment struct {
	PaymentID         uint                     `gorm:"primaryKey;autoIncrement" json:"payment_id"`
	InstallmentID     uint                     `gorm:"not null;index" json:"installment_id"`
	PaymentNumber     int                      `gorm:"not null" json:"payment_number"`
	DueDate           time.Time                `gorm:"not null" json:"due_date"`
	PaymentDate       *time.Time               `json:"payment_date"`
	DueAmount         float64                  `gorm:"type:decimal(15,2);not null" json:"due_amount"`
	PaidAmount        *float64                 `gorm:"type:decimal(15,2)" json:"paid_amount"`
	LateFee           *float64                 `gorm:"type:decimal(15,2)" json:"late_fee"`
	PaymentStatus     InstallmentPaymentStatus `gorm:"not null;default:'Pending'" json:"payment_status"`
	PaymentMethod     *PaymentTypeEnum         `json:"payment_method"`
	PaymentReference  *string                  `gorm:"size:255" json:"payment_reference"`
	Notes             *string                  `gorm:"type:text" json:"notes"`
	CreatedAt         time.Time                `json:"created_at"`
	UpdatedAt         time.Time                `json:"updated_at"`
	DeletedAt         gorm.DeletedAt           `gorm:"index" json:"deleted_at"`
	CreatedBy         *uint                    `json:"created_by"`

	// Relationships
	Installment *VehicleInstallment `gorm:"foreignKey:InstallmentID" json:"installment,omitempty"`
}