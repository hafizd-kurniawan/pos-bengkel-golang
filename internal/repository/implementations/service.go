package implementations

import (
	"boilerplate/internal/models"
	"boilerplate/internal/repository/interfaces"
	"context"

	"gorm.io/gorm"
)

// ServiceRepository implements the service repository interface
type ServiceRepository struct {
	db *gorm.DB
}

// NewServiceRepository creates a new service repository
func NewServiceRepository(db *gorm.DB) interfaces.ServiceRepository {
	return &ServiceRepository{db: db}
}

// Create creates a new service
func (r *ServiceRepository) Create(ctx context.Context, service *models.Service) error {
	return r.db.WithContext(ctx).Create(service).Error
}

// GetByID retrieves a service by ID
func (r *ServiceRepository) GetByID(ctx context.Context, id uint) (*models.Service, error) {
	var service models.Service
	err := r.db.WithContext(ctx).Preload("ServiceCategory").First(&service, id).Error
	if err != nil {
		return nil, err
	}
	return &service, nil
}

// GetByServiceCode retrieves a service by service code
func (r *ServiceRepository) GetByServiceCode(ctx context.Context, serviceCode string) (*models.Service, error) {
	var service models.Service
	err := r.db.WithContext(ctx).Preload("ServiceCategory").Where("service_code = ?", serviceCode).First(&service).Error
	if err != nil {
		return nil, err
	}
	return &service, nil
}

// Update updates a service
func (r *ServiceRepository) Update(ctx context.Context, service *models.Service) error {
	return r.db.WithContext(ctx).Save(service).Error
}

// Delete soft deletes a service
func (r *ServiceRepository) Delete(ctx context.Context, id uint) error {
	return r.db.WithContext(ctx).Delete(&models.Service{}, id).Error
}

// List retrieves services with pagination
func (r *ServiceRepository) List(ctx context.Context, limit, offset int) ([]*models.Service, error) {
	var services []*models.Service
	err := r.db.WithContext(ctx).Preload("ServiceCategory").Limit(limit).Offset(offset).Find(&services).Error
	if err != nil {
		return nil, err
	}
	return services, nil
}

// GetByCategoryID retrieves services by category ID
func (r *ServiceRepository) GetByCategoryID(ctx context.Context, categoryID uint) ([]*models.Service, error) {
	var services []*models.Service
	err := r.db.WithContext(ctx).Preload("ServiceCategory").Where("service_category_id = ?", categoryID).Find(&services).Error
	if err != nil {
		return nil, err
	}
	return services, nil
}

// GetByStatus retrieves services by status
func (r *ServiceRepository) GetByStatus(ctx context.Context, status models.StatusUmum) ([]*models.Service, error) {
	var services []*models.Service
	err := r.db.WithContext(ctx).Preload("ServiceCategory").Where("status = ?", status).Find(&services).Error
	if err != nil {
		return nil, err
	}
	return services, nil
}

// Search searches services by name or service code
func (r *ServiceRepository) Search(ctx context.Context, query string, limit, offset int) ([]*models.Service, error) {
	var services []*models.Service
	searchQuery := "%" + query + "%"
	
	err := r.db.WithContext(ctx).
		Preload("ServiceCategory").
		Where("name LIKE ? OR service_code LIKE ?", searchQuery, searchQuery).
		Limit(limit).
		Offset(offset).
		Find(&services).Error
	
	if err != nil {
		return nil, err
	}
	return services, nil
}

// ServiceCategoryRepository implements the service category repository interface
type ServiceCategoryRepository struct {
	db *gorm.DB
}

// NewServiceCategoryRepository creates a new service category repository
func NewServiceCategoryRepository(db *gorm.DB) interfaces.ServiceCategoryRepository {
	return &ServiceCategoryRepository{db: db}
}

// Create creates a new service category
func (r *ServiceCategoryRepository) Create(ctx context.Context, category *models.ServiceCategory) error {
	return r.db.WithContext(ctx).Create(category).Error
}

// GetByID retrieves a service category by ID
func (r *ServiceCategoryRepository) GetByID(ctx context.Context, id uint) (*models.ServiceCategory, error) {
	var category models.ServiceCategory
	err := r.db.WithContext(ctx).Preload("Services").First(&category, id).Error
	if err != nil {
		return nil, err
	}
	return &category, nil
}

// GetByName retrieves a service category by name
func (r *ServiceCategoryRepository) GetByName(ctx context.Context, name string) (*models.ServiceCategory, error) {
	var category models.ServiceCategory
	err := r.db.WithContext(ctx).Preload("Services").Where("name = ?", name).First(&category).Error
	if err != nil {
		return nil, err
	}
	return &category, nil
}

// Update updates a service category
func (r *ServiceCategoryRepository) Update(ctx context.Context, category *models.ServiceCategory) error {
	return r.db.WithContext(ctx).Save(category).Error
}

// Delete soft deletes a service category
func (r *ServiceCategoryRepository) Delete(ctx context.Context, id uint) error {
	return r.db.WithContext(ctx).Delete(&models.ServiceCategory{}, id).Error
}

// List retrieves service categories with pagination
func (r *ServiceCategoryRepository) List(ctx context.Context, limit, offset int) ([]*models.ServiceCategory, error) {
	var categories []*models.ServiceCategory
	err := r.db.WithContext(ctx).Preload("Services").Limit(limit).Offset(offset).Find(&categories).Error
	if err != nil {
		return nil, err
	}
	return categories, nil
}

// GetByStatus retrieves service categories by status
func (r *ServiceCategoryRepository) GetByStatus(ctx context.Context, status models.StatusUmum) ([]*models.ServiceCategory, error) {
	var categories []*models.ServiceCategory
	err := r.db.WithContext(ctx).Preload("Services").Where("status = ?", status).Find(&categories).Error
	if err != nil {
		return nil, err
	}
	return categories, nil
}

// ServiceJobRepository implements the service job repository interface
type ServiceJobRepository struct {
	db *gorm.DB
}

// NewServiceJobRepository creates a new service job repository
func NewServiceJobRepository(db *gorm.DB) interfaces.ServiceJobRepository {
	return &ServiceJobRepository{db: db}
}

// Create creates a new service job
func (r *ServiceJobRepository) Create(ctx context.Context, serviceJob *models.ServiceJob) error {
	return r.db.WithContext(ctx).Create(serviceJob).Error
}

// GetByID retrieves a service job by ID
func (r *ServiceJobRepository) GetByID(ctx context.Context, id uint) (*models.ServiceJob, error) {
	var serviceJob models.ServiceJob
	err := r.db.WithContext(ctx).
		Preload("Customer").
		Preload("Vehicle").
		Preload("Technician").
		Preload("ReceivedByUser").
		Preload("Outlet").
		Preload("ServiceDetails").
		Preload("Histories").
		First(&serviceJob, id).Error
	if err != nil {
		return nil, err
	}
	return &serviceJob, nil
}

// GetByServiceCode retrieves a service job by service code
func (r *ServiceJobRepository) GetByServiceCode(ctx context.Context, serviceCode string) (*models.ServiceJob, error) {
	var serviceJob models.ServiceJob
	err := r.db.WithContext(ctx).
		Preload("Customer").
		Preload("Vehicle").
		Preload("Technician").
		Preload("ReceivedByUser").
		Preload("Outlet").
		Preload("ServiceDetails").
		Preload("Histories").
		Where("service_code = ?", serviceCode).
		First(&serviceJob).Error
	if err != nil {
		return nil, err
	}
	return &serviceJob, nil
}

// Update updates a service job
func (r *ServiceJobRepository) Update(ctx context.Context, serviceJob *models.ServiceJob) error {
	return r.db.WithContext(ctx).Save(serviceJob).Error
}

// Delete soft deletes a service job
func (r *ServiceJobRepository) Delete(ctx context.Context, id uint) error {
	return r.db.WithContext(ctx).Delete(&models.ServiceJob{}, id).Error
}

// List retrieves service jobs with pagination
func (r *ServiceJobRepository) List(ctx context.Context, limit, offset int) ([]*models.ServiceJob, error) {
	var serviceJobs []*models.ServiceJob
	err := r.db.WithContext(ctx).
		Preload("Customer").
		Preload("Vehicle").
		Preload("Technician").
		Preload("ReceivedByUser").
		Preload("Outlet").
		Preload("ServiceDetails").
		Preload("Histories").
		Limit(limit).
		Offset(offset).
		Find(&serviceJobs).Error
	if err != nil {
		return nil, err
	}
	return serviceJobs, nil
}

// GetByCustomerID retrieves service jobs by customer ID
func (r *ServiceJobRepository) GetByCustomerID(ctx context.Context, customerID uint) ([]*models.ServiceJob, error) {
	var serviceJobs []*models.ServiceJob
	err := r.db.WithContext(ctx).
		Preload("Customer").
		Preload("Vehicle").
		Preload("Technician").
		Preload("ReceivedByUser").
		Preload("Outlet").
		Preload("ServiceDetails").
		Preload("Histories").
		Where("customer_id = ?", customerID).
		Find(&serviceJobs).Error
	if err != nil {
		return nil, err
	}
	return serviceJobs, nil
}

// GetByVehicleID retrieves service jobs by vehicle ID
func (r *ServiceJobRepository) GetByVehicleID(ctx context.Context, vehicleID uint) ([]*models.ServiceJob, error) {
	var serviceJobs []*models.ServiceJob
	err := r.db.WithContext(ctx).
		Preload("Customer").
		Preload("Vehicle").
		Preload("Technician").
		Preload("ReceivedByUser").
		Preload("Outlet").
		Preload("ServiceDetails").
		Preload("Histories").
		Where("vehicle_id = ?", vehicleID).
		Find(&serviceJobs).Error
	if err != nil {
		return nil, err
	}
	return serviceJobs, nil
}

// GetByTechnicianID retrieves service jobs by technician ID
func (r *ServiceJobRepository) GetByTechnicianID(ctx context.Context, technicianID uint) ([]*models.ServiceJob, error) {
	var serviceJobs []*models.ServiceJob
	err := r.db.WithContext(ctx).
		Preload("Customer").
		Preload("Vehicle").
		Preload("Technician").
		Preload("ReceivedByUser").
		Preload("Outlet").
		Preload("ServiceDetails").
		Preload("Histories").
		Where("technician_id = ?", technicianID).
		Find(&serviceJobs).Error
	if err != nil {
		return nil, err
	}
	return serviceJobs, nil
}

// GetByOutletID retrieves service jobs by outlet ID
func (r *ServiceJobRepository) GetByOutletID(ctx context.Context, outletID uint) ([]*models.ServiceJob, error) {
	var serviceJobs []*models.ServiceJob
	err := r.db.WithContext(ctx).
		Preload("Customer").
		Preload("Vehicle").
		Preload("Technician").
		Preload("ReceivedByUser").
		Preload("Outlet").
		Preload("ServiceDetails").
		Preload("Histories").
		Where("outlet_id = ?", outletID).
		Find(&serviceJobs).Error
	if err != nil {
		return nil, err
	}
	return serviceJobs, nil
}

// GetByStatus retrieves service jobs by status
func (r *ServiceJobRepository) GetByStatus(ctx context.Context, status models.ServiceStatusEnum) ([]*models.ServiceJob, error) {
	var serviceJobs []*models.ServiceJob
	err := r.db.WithContext(ctx).
		Preload("Customer").
		Preload("Vehicle").
		Preload("Technician").
		Preload("ReceivedByUser").
		Preload("Outlet").
		Preload("ServiceDetails").
		Preload("Histories").
		Where("status = ?", status).
		Find(&serviceJobs).Error
	if err != nil {
		return nil, err
	}
	return serviceJobs, nil
}

// UpdateStatus updates service job status
func (r *ServiceJobRepository) UpdateStatus(ctx context.Context, id uint, status models.ServiceStatusEnum) error {
	return r.db.WithContext(ctx).
		Model(&models.ServiceJob{}).
		Where("service_job_id = ?", id).
		Update("status", status).Error
}

// GetQueueNumber gets the next queue number for an outlet
func (r *ServiceJobRepository) GetQueueNumber(ctx context.Context, outletID uint) (int, error) {
	var maxQueue int
	err := r.db.WithContext(ctx).
		Model(&models.ServiceJob{}).
		Where("outlet_id = ? AND DATE(service_in_date) = CURRENT_DATE", outletID).
		Select("COALESCE(MAX(queue_number), 0)").
		Scan(&maxQueue).Error
	if err != nil {
		return 0, err
	}
	return maxQueue + 1, nil
}

// ServiceDetailRepository implements the service detail repository interface
type ServiceDetailRepository struct {
	db *gorm.DB
}

// NewServiceDetailRepository creates a new service detail repository
func NewServiceDetailRepository(db *gorm.DB) interfaces.ServiceDetailRepository {
	return &ServiceDetailRepository{db: db}
}

// Create creates a new service detail
func (r *ServiceDetailRepository) Create(ctx context.Context, detail *models.ServiceDetail) error {
	return r.db.WithContext(ctx).Create(detail).Error
}

// GetByID retrieves a service detail by ID
func (r *ServiceDetailRepository) GetByID(ctx context.Context, id uint) (*models.ServiceDetail, error) {
	var detail models.ServiceDetail
	err := r.db.WithContext(ctx).Preload("ServiceJob").First(&detail, id).Error
	if err != nil {
		return nil, err
	}
	return &detail, nil
}

// Update updates a service detail
func (r *ServiceDetailRepository) Update(ctx context.Context, detail *models.ServiceDetail) error {
	return r.db.WithContext(ctx).Save(detail).Error
}

// Delete soft deletes a service detail
func (r *ServiceDetailRepository) Delete(ctx context.Context, id uint) error {
	return r.db.WithContext(ctx).Delete(&models.ServiceDetail{}, id).Error
}

// List retrieves service details with pagination
func (r *ServiceDetailRepository) List(ctx context.Context, limit, offset int) ([]*models.ServiceDetail, error) {
	var details []*models.ServiceDetail
	err := r.db.WithContext(ctx).Preload("ServiceJob").Limit(limit).Offset(offset).Find(&details).Error
	if err != nil {
		return nil, err
	}
	return details, nil
}

// GetByServiceJobID retrieves service details by service job ID
func (r *ServiceDetailRepository) GetByServiceJobID(ctx context.Context, serviceJobID uint) ([]*models.ServiceDetail, error) {
	var details []*models.ServiceDetail
	err := r.db.WithContext(ctx).Preload("ServiceJob").Where("service_job_id = ?", serviceJobID).Find(&details).Error
	if err != nil {
		return nil, err
	}
	return details, nil
}

// DeleteByServiceJobID deletes service details by service job ID
func (r *ServiceDetailRepository) DeleteByServiceJobID(ctx context.Context, serviceJobID uint) error {
	return r.db.WithContext(ctx).Where("service_job_id = ?", serviceJobID).Delete(&models.ServiceDetail{}).Error
}

// ServiceJobHistoryRepository implements the service job history repository interface
type ServiceJobHistoryRepository struct {
	db *gorm.DB
}

// NewServiceJobHistoryRepository creates a new service job history repository
func NewServiceJobHistoryRepository(db *gorm.DB) interfaces.ServiceJobHistoryRepository {
	return &ServiceJobHistoryRepository{db: db}
}

// Create creates a new service job history
func (r *ServiceJobHistoryRepository) Create(ctx context.Context, history *models.ServiceJobHistory) error {
	return r.db.WithContext(ctx).Create(history).Error
}

// GetByID retrieves a service job history by ID
func (r *ServiceJobHistoryRepository) GetByID(ctx context.Context, id uint) (*models.ServiceJobHistory, error) {
	var history models.ServiceJobHistory
	err := r.db.WithContext(ctx).Preload("ServiceJob").Preload("User").First(&history, id).Error
	if err != nil {
		return nil, err
	}
	return &history, nil
}

// List retrieves service job histories with pagination
func (r *ServiceJobHistoryRepository) List(ctx context.Context, limit, offset int) ([]*models.ServiceJobHistory, error) {
	var histories []*models.ServiceJobHistory
	err := r.db.WithContext(ctx).Preload("ServiceJob").Preload("User").Limit(limit).Offset(offset).Find(&histories).Error
	if err != nil {
		return nil, err
	}
	return histories, nil
}

// GetByServiceJobID retrieves service job histories by service job ID
func (r *ServiceJobHistoryRepository) GetByServiceJobID(ctx context.Context, serviceJobID uint) ([]*models.ServiceJobHistory, error) {
	var histories []*models.ServiceJobHistory
	err := r.db.WithContext(ctx).Preload("ServiceJob").Preload("User").Where("service_job_id = ?", serviceJobID).Find(&histories).Error
	if err != nil {
		return nil, err
	}
	return histories, nil
}

// GetByUserID retrieves service job histories by user ID
func (r *ServiceJobHistoryRepository) GetByUserID(ctx context.Context, userID uint) ([]*models.ServiceJobHistory, error) {
	var histories []*models.ServiceJobHistory
	err := r.db.WithContext(ctx).Preload("ServiceJob").Preload("User").Where("user_id = ?", userID).Find(&histories).Error
	if err != nil {
		return nil, err
	}
	return histories, nil
}