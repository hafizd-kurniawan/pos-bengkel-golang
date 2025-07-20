package interfaces

import (
	"boilerplate/internal/models"
	"context"
	"time"
)

// Vehicle Purchase Use Case Requests and Responses

// CreateVehiclePurchaseRequest represents a request to purchase a vehicle from customer
type CreateVehiclePurchaseRequest struct {
	VehicleData         CreateVehicleRequest             `json:"vehicle_data" validate:"required"`
	PurchasePrice       float64                          `json:"purchase_price" validate:"required,min=0"`
	PaymentMethod       models.PaymentTypeEnum           `json:"payment_method" validate:"required"`
	EvaluationNotes     *string                          `json:"evaluation_notes,omitempty"`
	PaymentReference    *string                          `json:"payment_reference,omitempty"`
	CreatedBy           *uint                            `json:"created_by,omitempty"`
}

// CreateVehicleRequest represents vehicle data for purchase
type CreateVehicleRequest struct {
	CustomerID        uint                              `json:"customer_id" validate:"required"`
	PlateNumber       string                            `json:"plate_number" validate:"required,min=3,max=20"`
	Brand             string                            `json:"brand" validate:"required,min=2,max=100"`
	Model             string                            `json:"model" validate:"required,min=2,max=100"`
	Type              string                            `json:"type" validate:"required,min=2,max=100"`
	ProductionYear    int                               `json:"production_year" validate:"required,min=1900,max=2030"`
	ChassisNumber     string                            `json:"chassis_number" validate:"required,min=5,max=100"`
	EngineNumber      string                            `json:"engine_number" validate:"required,min=5,max=100"`
	Color             string                            `json:"color" validate:"required,min=2,max=50"`
	Mileage           *int                              `json:"mileage,omitempty"`
	FuelType          *string                           `json:"fuel_type,omitempty"`
	Transmission      *string                           `json:"transmission,omitempty"`
	ConditionStatus   models.VehicleConditionStatus     `json:"condition_status" validate:"required"`
	EstimatedValue    *float64                          `json:"estimated_value,omitempty"`
	ConditionNotes    *string                           `json:"condition_notes,omitempty"`
	InternalNotes     *string                           `json:"internal_notes,omitempty"`
}

// CreateReconditioningJobRequest represents a request to create reconditioning job
type CreateReconditioningJobRequest struct {
	VehicleID            uint                              `json:"vehicle_id" validate:"required"`
	JobTitle             string                            `json:"job_title" validate:"required,min=3,max=255"`
	JobDescription       *string                           `json:"job_description,omitempty"`
	EstimatedCost        *float64                          `json:"estimated_cost,omitempty"`
	AssignedTechnicianID *uint                            `json:"assigned_technician_id,omitempty"`
	Notes                *string                           `json:"notes,omitempty"`
	CreatedBy            *uint                             `json:"created_by,omitempty"`
}

// AddReconditioningDetailRequest represents a request to add reconditioning detail
type AddReconditioningDetailRequest struct {
	ReconditioningJobID uint                             `json:"reconditioning_job_id" validate:"required"`
	DetailType          models.ReconditioningDetailType  `json:"detail_type" validate:"required"`
	ProductID           *uint                            `json:"product_id,omitempty"`
	ServiceID           *uint                            `json:"service_id,omitempty"`
	Description         string                           `json:"description" validate:"required,min=3,max=255"`
	Quantity            int                              `json:"quantity" validate:"required,min=1"`
	UnitPrice           float64                          `json:"unit_price" validate:"required,min=0"`
	Notes               *string                          `json:"notes,omitempty"`
	CreatedBy           *uint                            `json:"created_by,omitempty"`
}

// CompleteReconditioningJobRequest represents a request to complete reconditioning job
type CompleteReconditioningJobRequest struct {
	ReconditioningJobID uint                             `json:"reconditioning_job_id" validate:"required"`
	ActualCost          *float64                         `json:"actual_cost,omitempty"`
	CompletionNotes     *string                          `json:"completion_notes,omitempty"`
}

// Vehicle Sales Use Case Requests and Responses

// CreateVehicleSaleRequest represents a request to sell a vehicle
type CreateVehicleSaleRequest struct {
	VehicleID         uint                             `json:"vehicle_id" validate:"required"`
	CustomerID        uint                             `json:"customer_id" validate:"required"`
	SalePrice         float64                          `json:"sale_price" validate:"required,min=0"`
	TransactionType   models.SalesTransactionType      `json:"transaction_type" validate:"required"`
	PaymentMethod     models.PaymentTypeEnum           `json:"payment_method" validate:"required"`
	DownPayment       *float64                         `json:"down_payment,omitempty"`
	PaymentReference  *string                          `json:"payment_reference,omitempty"`
	SalesPersonID     *uint                            `json:"sales_person_id,omitempty"`
	Notes             *string                          `json:"notes,omitempty"`
	InstallmentConfig *InstallmentConfigRequest        `json:"installment_config,omitempty"`
	CreatedBy         *uint                            `json:"created_by,omitempty"`
}

// InstallmentConfigRequest represents installment configuration
type InstallmentConfigRequest struct {
	NumberOfInstallments int      `json:"number_of_installments" validate:"required,min=1,max=60"`
	InterestRate         *float64 `json:"interest_rate,omitempty"`
	StartDate            string   `json:"start_date" validate:"required"`
}

// ProcessInstallmentPaymentRequest represents a request to process installment payment
type ProcessInstallmentPaymentRequest struct {
	PaymentID        uint                             `json:"payment_id" validate:"required"`
	PaidAmount       float64                          `json:"paid_amount" validate:"required,min=0"`
	PaymentMethod    models.PaymentTypeEnum           `json:"payment_method" validate:"required"`
	PaymentReference *string                          `json:"payment_reference,omitempty"`
	Notes            *string                          `json:"notes,omitempty"`
}

// Response Types

// VehiclePurchaseResponse represents response after vehicle purchase
type VehiclePurchaseResponse struct {
	Vehicle             *models.Vehicle                   `json:"vehicle"`
	PurchaseTransaction *models.VehiclePurchaseTransaction `json:"purchase_transaction"`
}

// ReconditioningJobResponse represents reconditioning job response
type ReconditioningJobResponse struct {
	Job     *models.VehicleReconditioningJob `json:"job"`
	Details []*models.ReconditioningDetail   `json:"details,omitempty"`
}

// VehicleSaleResponse represents response after vehicle sale
type VehicleSaleResponse struct {
	Vehicle          *models.Vehicle              `json:"vehicle"`
	SalesTransaction *models.VehicleSalesTransaction `json:"sales_transaction"`
	Installment      *models.VehicleInstallment   `json:"installment,omitempty"`
	ProfitAmount     *float64                     `json:"profit_amount,omitempty"`
}

// InstallmentPaymentResponse represents installment payment response
type InstallmentPaymentResponse struct {
	Payment            *models.InstallmentPayment `json:"payment"`
	RemainingBalance   float64                    `json:"remaining_balance"`
	NextPaymentDue     *time.Time                 `json:"next_payment_due,omitempty"`
}

// Use Case Interfaces

// VehiclePurchaseUseCase defines methods for vehicle purchase operations
type VehiclePurchaseUseCase interface {
	// Purchase vehicle from customer
	PurchaseVehicle(ctx context.Context, req *CreateVehiclePurchaseRequest) (*VehiclePurchaseResponse, error)
	
	// Reconditioning operations
	CreateReconditioningJob(ctx context.Context, req *CreateReconditioningJobRequest) (*models.VehicleReconditioningJob, error)
	AddReconditioningDetail(ctx context.Context, req *AddReconditioningDetailRequest) (*models.ReconditioningDetail, error)
	CompleteReconditioningJob(ctx context.Context, req *CompleteReconditioningJobRequest) (*ReconditioningJobResponse, error)
	
	// Query operations
	GetVehicleByID(ctx context.Context, id uint) (*models.Vehicle, error)
	GetReconditioningJobsByVehicleID(ctx context.Context, vehicleID uint) ([]*models.VehicleReconditioningJob, error)
	GetReconditioningJobByID(ctx context.Context, id uint) (*ReconditioningJobResponse, error)
	GetReconditioningDetailsByJobID(ctx context.Context, jobID uint) ([]*models.ReconditioningDetail, error)
}

// VehicleSalesUseCase defines methods for vehicle sales operations
type VehicleSalesUseCase interface {
	// Sales operations
	SellVehicle(ctx context.Context, req *CreateVehicleSaleRequest) (*VehicleSaleResponse, error)
	MarkVehicleForSale(ctx context.Context, vehicleID uint, sellingPrice float64) error
	
	// Installment operations
	ProcessInstallmentPayment(ctx context.Context, req *ProcessInstallmentPaymentRequest) (*InstallmentPaymentResponse, error)
	GenerateInstallmentSchedule(ctx context.Context, installmentID uint) ([]*models.InstallmentPayment, error)
	GetOverduePayments(ctx context.Context) ([]*models.InstallmentPayment, error)
	CalculateLateFee(ctx context.Context, paymentID uint) (float64, error)
	
	// Query operations
	GetVehiclesByStatus(ctx context.Context, status models.VehicleSaleStatus) ([]*models.Vehicle, error)
	GetSalesTransactionByID(ctx context.Context, id uint) (*models.VehicleSalesTransaction, error)
	GetInstallmentByID(ctx context.Context, id uint) (*models.VehicleInstallment, error)
	GetInstallmentPaymentsByInstallmentID(ctx context.Context, installmentID uint) ([]*models.InstallmentPayment, error)
	CalculateProfitForVehicle(ctx context.Context, vehicleID uint) (float64, error)
}