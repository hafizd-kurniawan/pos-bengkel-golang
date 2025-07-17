package interfaces

import (
	"boilerplate/internal/models"
	"context"
)

// ServiceRepository interface for service operations
type ServiceRepository interface {
	Create(ctx context.Context, service *models.Service) error
	GetByID(ctx context.Context, id uint) (*models.Service, error)
	GetByServiceCode(ctx context.Context, serviceCode string) (*models.Service, error)
	Update(ctx context.Context, service *models.Service) error
	Delete(ctx context.Context, id uint) error
	List(ctx context.Context, limit, offset int) ([]*models.Service, error)
	GetByCategoryID(ctx context.Context, categoryID uint) ([]*models.Service, error)
	GetByStatus(ctx context.Context, status models.StatusUmum) ([]*models.Service, error)
	Search(ctx context.Context, query string, limit, offset int) ([]*models.Service, error)
}

// ServiceCategoryRepository interface for service category operations
type ServiceCategoryRepository interface {
	Create(ctx context.Context, category *models.ServiceCategory) error
	GetByID(ctx context.Context, id uint) (*models.ServiceCategory, error)
	GetByName(ctx context.Context, name string) (*models.ServiceCategory, error)
	Update(ctx context.Context, category *models.ServiceCategory) error
	Delete(ctx context.Context, id uint) error
	List(ctx context.Context, limit, offset int) ([]*models.ServiceCategory, error)
	GetByStatus(ctx context.Context, status models.StatusUmum) ([]*models.ServiceCategory, error)
}

// ServiceJobRepository interface for service job operations
type ServiceJobRepository interface {
	Create(ctx context.Context, serviceJob *models.ServiceJob) error
	GetByID(ctx context.Context, id uint) (*models.ServiceJob, error)
	GetByServiceCode(ctx context.Context, serviceCode string) (*models.ServiceJob, error)
	Update(ctx context.Context, serviceJob *models.ServiceJob) error
	Delete(ctx context.Context, id uint) error
	List(ctx context.Context, limit, offset int) ([]*models.ServiceJob, error)
	GetByCustomerID(ctx context.Context, customerID uint) ([]*models.ServiceJob, error)
	GetByVehicleID(ctx context.Context, vehicleID uint) ([]*models.ServiceJob, error)
	GetByTechnicianID(ctx context.Context, technicianID uint) ([]*models.ServiceJob, error)
	GetByOutletID(ctx context.Context, outletID uint) ([]*models.ServiceJob, error)
	GetByStatus(ctx context.Context, status models.ServiceStatusEnum) ([]*models.ServiceJob, error)
	UpdateStatus(ctx context.Context, id uint, status models.ServiceStatusEnum) error
	GetQueueNumber(ctx context.Context, outletID uint) (int, error)
	
	// Queue management methods
	GetQueueByOutletID(ctx context.Context, outletID uint) ([]*models.ServiceJob, error)
	GetTodayQueueByOutletID(ctx context.Context, outletID uint) ([]*models.ServiceJob, error)
	UpdateQueueNumber(ctx context.Context, serviceJobID uint, queueNumber int) error
}

// ServiceDetailRepository interface for service detail operations
type ServiceDetailRepository interface {
	Create(ctx context.Context, detail *models.ServiceDetail) error
	GetByID(ctx context.Context, id uint) (*models.ServiceDetail, error)
	Update(ctx context.Context, detail *models.ServiceDetail) error
	Delete(ctx context.Context, id uint) error
	List(ctx context.Context, limit, offset int) ([]*models.ServiceDetail, error)
	GetByServiceJobID(ctx context.Context, serviceJobID uint) ([]*models.ServiceDetail, error)
	DeleteByServiceJobID(ctx context.Context, serviceJobID uint) error
}

// ServiceJobHistoryRepository interface for service job history operations
type ServiceJobHistoryRepository interface {
	Create(ctx context.Context, history *models.ServiceJobHistory) error
	GetByID(ctx context.Context, id uint) (*models.ServiceJobHistory, error)
	List(ctx context.Context, limit, offset int) ([]*models.ServiceJobHistory, error)
	GetByServiceJobID(ctx context.Context, serviceJobID uint) ([]*models.ServiceJobHistory, error)
	GetByUserID(ctx context.Context, userID uint) ([]*models.ServiceJobHistory, error)
}