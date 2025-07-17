package interfaces

import (
	"boilerplate/internal/models"
	"context"
	"time"
)

// Service request structures
type CreateServiceRequest struct {
	ServiceCode       string            `json:"service_code" validate:"required,min=3,max=50"`
	Name              string            `json:"name" validate:"required,min=2,max=255"`
	ServiceCategoryID uint              `json:"service_category_id" validate:"required"`
	Fee               float64           `json:"fee" validate:"required,min=0"`
	Status            models.StatusUmum `json:"status,omitempty"`
	CreatedBy         *uint             `json:"created_by,omitempty"`
}

type UpdateServiceRequest struct {
	ServiceCode       *string            `json:"service_code,omitempty" validate:"omitempty,min=3,max=50"`
	Name              *string            `json:"name,omitempty" validate:"omitempty,min=2,max=255"`
	ServiceCategoryID *uint              `json:"service_category_id,omitempty"`
	Fee               *float64           `json:"fee,omitempty" validate:"omitempty,min=0"`
	Status            *models.StatusUmum `json:"status,omitempty"`
}

// Service Category request structures
type CreateServiceCategoryRequest struct {
	Name      string            `json:"name" validate:"required,min=2,max=255"`
	Status    models.StatusUmum `json:"status,omitempty"`
	CreatedBy *uint             `json:"created_by,omitempty"`
}

type UpdateServiceCategoryRequest struct {
	Name   *string            `json:"name,omitempty" validate:"omitempty,min=2,max=255"`
	Status *models.StatusUmum `json:"status,omitempty"`
}

// Service Job request structures
type CreateServiceJobRequest struct {
	CustomerID                 uint                      `json:"customer_id" validate:"required"`
	VehicleID                  uint                      `json:"vehicle_id" validate:"required"`
	TechnicianID               *uint                     `json:"technician_id,omitempty"`
	ReceivedByUserID           uint                      `json:"received_by_user_id" validate:"required"`
	OutletID                   uint                      `json:"outlet_id" validate:"required"`
	ProblemDescription         string                    `json:"problem_description" validate:"required,min=10"`
	TechnicianNotes            *string                   `json:"technician_notes,omitempty"`
	Status                     models.ServiceStatusEnum  `json:"status,omitempty"`
	ServiceInDate              time.Time                 `json:"service_in_date" validate:"required"`
	WarrantyExpiresAt          *time.Time                `json:"warranty_expires_at,omitempty"`
	NextServiceReminderDate    *time.Time                `json:"next_service_reminder_date,omitempty"`
	DownPayment                float64                   `json:"down_payment" validate:"min=0"`
	CreatedBy                  *uint                     `json:"created_by,omitempty"`
}

type UpdateServiceJobRequest struct {
	CustomerID                 *uint                     `json:"customer_id,omitempty"`
	VehicleID                  *uint                     `json:"vehicle_id,omitempty"`
	TechnicianID               *uint                     `json:"technician_id,omitempty"`
	ReceivedByUserID           *uint                     `json:"received_by_user_id,omitempty"`
	OutletID                   *uint                     `json:"outlet_id,omitempty"`
	ProblemDescription         *string                   `json:"problem_description,omitempty" validate:"omitempty,min=10"`
	TechnicianNotes            *string                   `json:"technician_notes,omitempty"`
	Status                     *models.ServiceStatusEnum `json:"status,omitempty"`
	ServiceInDate              *time.Time                `json:"service_in_date,omitempty"`
	PickedUpDate               *time.Time                `json:"picked_up_date,omitempty"`
	ComplainDate               *time.Time                `json:"complain_date,omitempty"`
	WarrantyExpiresAt          *time.Time                `json:"warranty_expires_at,omitempty"`
	NextServiceReminderDate    *time.Time                `json:"next_service_reminder_date,omitempty"`
	DownPayment                *float64                  `json:"down_payment,omitempty" validate:"omitempty,min=0"`
	GrandTotal                 *float64                  `json:"grand_total,omitempty" validate:"omitempty,min=0"`
	TechnicianCommission       *float64                  `json:"technician_commission,omitempty" validate:"omitempty,min=0"`
	ShopProfit                 *float64                  `json:"shop_profit,omitempty" validate:"omitempty,min=0"`
}

// Service Detail request structures
type CreateServiceDetailRequest struct {
	ServiceJobID     uint    `json:"service_job_id" validate:"required"`
	ItemID           uint    `json:"item_id" validate:"required"`
	ItemType         string  `json:"item_type" validate:"required,oneof=service product"`
	Description      string  `json:"description" validate:"required,min=2,max=255"`
	SerialNumberUsed *string `json:"serial_number_used,omitempty"`
	Quantity         int     `json:"quantity" validate:"required,min=1"`
	PricePerItem     float64 `json:"price_per_item" validate:"required,min=0"`
	CostPerItem      float64 `json:"cost_per_item" validate:"required,min=0"`
}

type UpdateServiceDetailRequest struct {
	ServiceJobID     *uint    `json:"service_job_id,omitempty"`
	ItemID           *uint    `json:"item_id,omitempty"`
	ItemType         *string  `json:"item_type,omitempty" validate:"omitempty,oneof=service product"`
	Description      *string  `json:"description,omitempty" validate:"omitempty,min=2,max=255"`
	SerialNumberUsed *string  `json:"serial_number_used,omitempty"`
	Quantity         *int     `json:"quantity,omitempty" validate:"omitempty,min=1"`
	PricePerItem     *float64 `json:"price_per_item,omitempty" validate:"omitempty,min=0"`
	CostPerItem      *float64 `json:"cost_per_item,omitempty" validate:"omitempty,min=0"`
}

// Service Job History request structures
type CreateServiceJobHistoryRequest struct {
	ServiceJobID uint    `json:"service_job_id" validate:"required"`
	UserID       uint    `json:"user_id" validate:"required"`
	Notes        *string `json:"notes,omitempty"`
}

// Usecase interfaces
type ServiceUsecase interface {
	CreateService(ctx context.Context, req CreateServiceRequest) (*models.Service, error)
	GetService(ctx context.Context, id uint) (*models.Service, error)
	GetServiceByServiceCode(ctx context.Context, serviceCode string) (*models.Service, error)
	UpdateService(ctx context.Context, id uint, req UpdateServiceRequest) (*models.Service, error)
	DeleteService(ctx context.Context, id uint) error
	ListServices(ctx context.Context, limit, offset int) ([]*models.Service, error)
	GetServicesByCategory(ctx context.Context, categoryID uint) ([]*models.Service, error)
	GetServicesByStatus(ctx context.Context, status models.StatusUmum) ([]*models.Service, error)
	SearchServices(ctx context.Context, query string, limit, offset int) ([]*models.Service, error)
}

type ServiceCategoryUsecase interface {
	CreateServiceCategory(ctx context.Context, req CreateServiceCategoryRequest) (*models.ServiceCategory, error)
	GetServiceCategory(ctx context.Context, id uint) (*models.ServiceCategory, error)
	GetServiceCategoryByName(ctx context.Context, name string) (*models.ServiceCategory, error)
	UpdateServiceCategory(ctx context.Context, id uint, req UpdateServiceCategoryRequest) (*models.ServiceCategory, error)
	DeleteServiceCategory(ctx context.Context, id uint) error
	ListServiceCategories(ctx context.Context, limit, offset int) ([]*models.ServiceCategory, error)
	GetServiceCategoriesByStatus(ctx context.Context, status models.StatusUmum) ([]*models.ServiceCategory, error)
}

type ServiceJobUsecase interface {
	CreateServiceJob(ctx context.Context, req CreateServiceJobRequest) (*models.ServiceJob, error)
	GetServiceJob(ctx context.Context, id uint) (*models.ServiceJob, error)
	GetServiceJobByServiceCode(ctx context.Context, serviceCode string) (*models.ServiceJob, error)
	UpdateServiceJob(ctx context.Context, id uint, req UpdateServiceJobRequest) (*models.ServiceJob, error)
	DeleteServiceJob(ctx context.Context, id uint) error
	ListServiceJobs(ctx context.Context, limit, offset int) ([]*models.ServiceJob, error)
	GetServiceJobsByCustomer(ctx context.Context, customerID uint) ([]*models.ServiceJob, error)
	GetServiceJobsByVehicle(ctx context.Context, vehicleID uint) ([]*models.ServiceJob, error)
	GetServiceJobsByTechnician(ctx context.Context, technicianID uint) ([]*models.ServiceJob, error)
	GetServiceJobsByOutlet(ctx context.Context, outletID uint) ([]*models.ServiceJob, error)
	GetServiceJobsByStatus(ctx context.Context, status models.ServiceStatusEnum) ([]*models.ServiceJob, error)
	UpdateServiceJobStatus(ctx context.Context, id uint, status models.ServiceStatusEnum, userID uint, notes *string) error
	CalculateServiceJobTotals(ctx context.Context, serviceJobID uint) error
}

type ServiceDetailUsecase interface {
	CreateServiceDetail(ctx context.Context, req CreateServiceDetailRequest) (*models.ServiceDetail, error)
	GetServiceDetail(ctx context.Context, id uint) (*models.ServiceDetail, error)
	UpdateServiceDetail(ctx context.Context, id uint, req UpdateServiceDetailRequest) (*models.ServiceDetail, error)
	DeleteServiceDetail(ctx context.Context, id uint) error
	ListServiceDetails(ctx context.Context, limit, offset int) ([]*models.ServiceDetail, error)
	GetServiceDetailsByServiceJob(ctx context.Context, serviceJobID uint) ([]*models.ServiceDetail, error)
	DeleteServiceDetailsByServiceJob(ctx context.Context, serviceJobID uint) error
}

type ServiceJobHistoryUsecase interface {
	CreateServiceJobHistory(ctx context.Context, req CreateServiceJobHistoryRequest) (*models.ServiceJobHistory, error)
	GetServiceJobHistory(ctx context.Context, id uint) (*models.ServiceJobHistory, error)
	ListServiceJobHistories(ctx context.Context, limit, offset int) ([]*models.ServiceJobHistory, error)
	GetServiceJobHistoriesByServiceJob(ctx context.Context, serviceJobID uint) ([]*models.ServiceJobHistory, error)
	GetServiceJobHistoriesByUser(ctx context.Context, userID uint) ([]*models.ServiceJobHistory, error)
}