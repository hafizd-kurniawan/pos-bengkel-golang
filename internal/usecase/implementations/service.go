package implementations

import (
	"boilerplate/internal/models"
	"boilerplate/internal/repository"
	"boilerplate/internal/usecase/interfaces"
	"context"
	"errors"
	"fmt"
	"time"

	"gorm.io/gorm"
)

// ServiceUsecase implements the service usecase interface
type ServiceUsecase struct {
	repo *repository.RepositoryManager
}

// NewServiceUsecase creates a new service usecase
func NewServiceUsecase(repo *repository.RepositoryManager) interfaces.ServiceUsecase {
	return &ServiceUsecase{repo: repo}
}

// CreateService creates a new service
func (u *ServiceUsecase) CreateService(ctx context.Context, req interfaces.CreateServiceRequest) (*models.Service, error) {
	// Validate service category exists
	_, err := u.repo.ServiceCategory.GetByID(ctx, req.ServiceCategoryID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("service category not found")
		}
		return nil, err
	}

	// Check if service code already exists
	existingService, err := u.repo.Service.GetByServiceCode(ctx, req.ServiceCode)
	if err == nil && existingService != nil {
		return nil, errors.New("service with this service code already exists")
	}
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}

	// Set default status if not provided
	status := req.Status
	if status == "" {
		status = models.StatusAktif
	}

	service := &models.Service{
		ServiceCode:       req.ServiceCode,
		Name:              req.Name,
		ServiceCategoryID: req.ServiceCategoryID,
		Fee:               req.Fee,
		Status:            status,
		CreatedBy:         req.CreatedBy,
		CreatedAt:         time.Now(),
		UpdatedAt:         time.Now(),
	}

	if err := u.repo.Service.Create(ctx, service); err != nil {
		return nil, err
	}

	return service, nil
}

// GetService retrieves a service by ID
func (u *ServiceUsecase) GetService(ctx context.Context, id uint) (*models.Service, error) {
	service, err := u.repo.Service.GetByID(ctx, id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("service not found")
		}
		return nil, err
	}
	return service, nil
}

// GetServiceByServiceCode retrieves a service by service code
func (u *ServiceUsecase) GetServiceByServiceCode(ctx context.Context, serviceCode string) (*models.Service, error) {
	service, err := u.repo.Service.GetByServiceCode(ctx, serviceCode)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("service not found")
		}
		return nil, err
	}
	return service, nil
}

// UpdateService updates a service
func (u *ServiceUsecase) UpdateService(ctx context.Context, id uint, req interfaces.UpdateServiceRequest) (*models.Service, error) {
	service, err := u.repo.Service.GetByID(ctx, id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("service not found")
		}
		return nil, err
	}

	// Validate service category exists if provided
	if req.ServiceCategoryID != nil {
		_, err := u.repo.ServiceCategory.GetByID(ctx, *req.ServiceCategoryID)
		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return nil, errors.New("service category not found")
			}
			return nil, err
		}
	}

	// Check if service code already exists (if being updated)
	if req.ServiceCode != nil && *req.ServiceCode != service.ServiceCode {
		existingService, err := u.repo.Service.GetByServiceCode(ctx, *req.ServiceCode)
		if err == nil && existingService != nil && existingService.ServiceID != id {
			return nil, errors.New("service with this service code already exists")
		}
		if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, err
		}
	}

	// Update fields if provided
	if req.ServiceCode != nil {
		service.ServiceCode = *req.ServiceCode
	}
	if req.Name != nil {
		service.Name = *req.Name
	}
	if req.ServiceCategoryID != nil {
		service.ServiceCategoryID = *req.ServiceCategoryID
	}
	if req.Fee != nil {
		service.Fee = *req.Fee
	}
	if req.Status != nil {
		service.Status = *req.Status
	}
	service.UpdatedAt = time.Now()

	if err := u.repo.Service.Update(ctx, service); err != nil {
		return nil, err
	}

	return service, nil
}

// DeleteService deletes a service
func (u *ServiceUsecase) DeleteService(ctx context.Context, id uint) error {
	_, err := u.repo.Service.GetByID(ctx, id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New("service not found")
		}
		return err
	}

	// TODO: Check if service is used in service jobs

	return u.repo.Service.Delete(ctx, id)
}

// ListServices retrieves services with pagination
func (u *ServiceUsecase) ListServices(ctx context.Context, limit, offset int) ([]*models.Service, error) {
	if limit <= 0 {
		limit = 10
	}
	if offset < 0 {
		offset = 0
	}

	return u.repo.Service.List(ctx, limit, offset)
}

// GetServicesByCategory retrieves services by category
func (u *ServiceUsecase) GetServicesByCategory(ctx context.Context, categoryID uint) ([]*models.Service, error) {
	return u.repo.Service.GetByCategoryID(ctx, categoryID)
}

// GetServicesByStatus retrieves services by status
func (u *ServiceUsecase) GetServicesByStatus(ctx context.Context, status models.StatusUmum) ([]*models.Service, error) {
	return u.repo.Service.GetByStatus(ctx, status)
}

// SearchServices searches services
func (u *ServiceUsecase) SearchServices(ctx context.Context, query string, limit, offset int) ([]*models.Service, error) {
	if limit <= 0 {
		limit = 10
	}
	if offset < 0 {
		offset = 0
	}

	return u.repo.Service.Search(ctx, query, limit, offset)
}

// ServiceCategoryUsecase implements the service category usecase interface
type ServiceCategoryUsecase struct {
	repo *repository.RepositoryManager
}

// NewServiceCategoryUsecase creates a new service category usecase
func NewServiceCategoryUsecase(repo *repository.RepositoryManager) interfaces.ServiceCategoryUsecase {
	return &ServiceCategoryUsecase{repo: repo}
}

// CreateServiceCategory creates a new service category
func (u *ServiceCategoryUsecase) CreateServiceCategory(ctx context.Context, req interfaces.CreateServiceCategoryRequest) (*models.ServiceCategory, error) {
	// Check if category name already exists
	existingCategory, err := u.repo.ServiceCategory.GetByName(ctx, req.Name)
	if err == nil && existingCategory != nil {
		return nil, errors.New("service category with this name already exists")
	}
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}

	// Set default status if not provided
	status := req.Status
	if status == "" {
		status = models.StatusAktif
	}

	category := &models.ServiceCategory{
		Name:      req.Name,
		Status:    status,
		CreatedBy: req.CreatedBy,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	if err := u.repo.ServiceCategory.Create(ctx, category); err != nil {
		return nil, err
	}

	return category, nil
}

// GetServiceCategory retrieves a service category by ID
func (u *ServiceCategoryUsecase) GetServiceCategory(ctx context.Context, id uint) (*models.ServiceCategory, error) {
	category, err := u.repo.ServiceCategory.GetByID(ctx, id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("service category not found")
		}
		return nil, err
	}
	return category, nil
}

// GetServiceCategoryByName retrieves a service category by name
func (u *ServiceCategoryUsecase) GetServiceCategoryByName(ctx context.Context, name string) (*models.ServiceCategory, error) {
	category, err := u.repo.ServiceCategory.GetByName(ctx, name)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("service category not found")
		}
		return nil, err
	}
	return category, nil
}

// UpdateServiceCategory updates a service category
func (u *ServiceCategoryUsecase) UpdateServiceCategory(ctx context.Context, id uint, req interfaces.UpdateServiceCategoryRequest) (*models.ServiceCategory, error) {
	category, err := u.repo.ServiceCategory.GetByID(ctx, id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("service category not found")
		}
		return nil, err
	}

	// Check if category name already exists (if being updated)
	if req.Name != nil && *req.Name != category.Name {
		existingCategory, err := u.repo.ServiceCategory.GetByName(ctx, *req.Name)
		if err == nil && existingCategory != nil && existingCategory.ServiceCategoryID != id {
			return nil, errors.New("service category with this name already exists")
		}
		if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, err
		}
	}

	// Update fields if provided
	if req.Name != nil {
		category.Name = *req.Name
	}
	if req.Status != nil {
		category.Status = *req.Status
	}
	category.UpdatedAt = time.Now()

	if err := u.repo.ServiceCategory.Update(ctx, category); err != nil {
		return nil, err
	}

	return category, nil
}

// DeleteServiceCategory deletes a service category
func (u *ServiceCategoryUsecase) DeleteServiceCategory(ctx context.Context, id uint) error {
	_, err := u.repo.ServiceCategory.GetByID(ctx, id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New("service category not found")
		}
		return err
	}

	// Check if category has services
	services, err := u.repo.Service.GetByCategoryID(ctx, id)
	if err != nil {
		return err
	}
	if len(services) > 0 {
		return errors.New("cannot delete service category with existing services")
	}

	return u.repo.ServiceCategory.Delete(ctx, id)
}

// ListServiceCategories retrieves service categories with pagination
func (u *ServiceCategoryUsecase) ListServiceCategories(ctx context.Context, limit, offset int) ([]*models.ServiceCategory, error) {
	if limit <= 0 {
		limit = 10
	}
	if offset < 0 {
		offset = 0
	}

	return u.repo.ServiceCategory.List(ctx, limit, offset)
}

// GetServiceCategoriesByStatus retrieves service categories by status
func (u *ServiceCategoryUsecase) GetServiceCategoriesByStatus(ctx context.Context, status models.StatusUmum) ([]*models.ServiceCategory, error) {
	return u.repo.ServiceCategory.GetByStatus(ctx, status)
}

// ServiceJobUsecase implements the service job usecase interface
type ServiceJobUsecase struct {
	repo *repository.RepositoryManager
}

// NewServiceJobUsecase creates a new service job usecase
func NewServiceJobUsecase(repo *repository.RepositoryManager) interfaces.ServiceJobUsecase {
	return &ServiceJobUsecase{repo: repo}
}

// CreateServiceJob creates a new service job
func (u *ServiceJobUsecase) CreateServiceJob(ctx context.Context, req interfaces.CreateServiceJobRequest) (*models.ServiceJob, error) {
	// Validate customer exists
	_, err := u.repo.Customer.GetByID(ctx, req.CustomerID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("customer not found")
		}
		return nil, err
	}

	// Validate vehicle exists
	_, err = u.repo.CustomerVehicle.GetByID(ctx, req.VehicleID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("vehicle not found")
		}
		return nil, err
	}

	// Validate technician exists if provided
	if req.TechnicianID != nil {
		_, err := u.repo.User.GetByID(ctx, *req.TechnicianID)
		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return nil, errors.New("technician not found")
			}
			return nil, err
		}
	}

	// Validate received by user exists
	_, err = u.repo.User.GetByID(ctx, req.ReceivedByUserID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("received by user not found")
		}
		return nil, err
	}

	// Validate outlet exists
	_, err = u.repo.Outlet.GetByID(ctx, req.OutletID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("outlet not found")
		}
		return nil, err
	}

	// Generate service code
	serviceCode := fmt.Sprintf("SJ-%d-%d", req.OutletID, time.Now().Unix())

	// Get queue number
	queueNumber, err := u.repo.ServiceJob.GetQueueNumber(ctx, req.OutletID)
	if err != nil {
		return nil, err
	}

	// Set default status if not provided
	status := req.Status
	if status == "" {
		status = models.ServiceStatusAntri
	}

	serviceJob := &models.ServiceJob{
		ServiceCode:             serviceCode,
		QueueNumber:             queueNumber,
		CustomerID:              req.CustomerID,
		VehicleID:               req.VehicleID,
		TechnicianID:            req.TechnicianID,
		ReceivedByUserID:        req.ReceivedByUserID,
		OutletID:                req.OutletID,
		ProblemDescription:      req.ProblemDescription,
		TechnicianNotes:         req.TechnicianNotes,
		Status:                  status,
		ServiceInDate:           req.ServiceInDate,
		WarrantyExpiresAt:       req.WarrantyExpiresAt,
		NextServiceReminderDate: req.NextServiceReminderDate,
		DownPayment:             req.DownPayment,
		CreatedBy:               req.CreatedBy,
		CreatedAt:               time.Now(),
		UpdatedAt:               time.Now(),
	}

	if err := u.repo.ServiceJob.Create(ctx, serviceJob); err != nil {
		return nil, err
	}

	// Create history entry
	historyReq := interfaces.CreateServiceJobHistoryRequest{
		ServiceJobID: serviceJob.ServiceJobID,
		UserID:       req.ReceivedByUserID,
		Notes:        &req.ProblemDescription,
	}
	_, err = u.createServiceJobHistory(ctx, historyReq)
	if err != nil {
		// Don't fail the entire operation for history creation failure
		// but log it
		fmt.Printf("Failed to create service job history: %v\n", err)
	}

	return serviceJob, nil
}

// CreateServiceJobHistory creates a service job history entry
func (u *ServiceJobUsecase) createServiceJobHistory(ctx context.Context, req interfaces.CreateServiceJobHistoryRequest) (*models.ServiceJobHistory, error) {
	history := &models.ServiceJobHistory{
		ServiceJobID: req.ServiceJobID,
		UserID:       req.UserID,
		Notes:        req.Notes,
		ChangedAt:    time.Now(),
	}

	if err := u.repo.ServiceJobHistory.Create(ctx, history); err != nil {
		return nil, err
	}

	return history, nil
}

// GetServiceJob retrieves a service job by ID
func (u *ServiceJobUsecase) GetServiceJob(ctx context.Context, id uint) (*models.ServiceJob, error) {
	serviceJob, err := u.repo.ServiceJob.GetByID(ctx, id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("service job not found")
		}
		return nil, err
	}
	return serviceJob, nil
}

// GetServiceJobByServiceCode retrieves a service job by service code
func (u *ServiceJobUsecase) GetServiceJobByServiceCode(ctx context.Context, serviceCode string) (*models.ServiceJob, error) {
	serviceJob, err := u.repo.ServiceJob.GetByServiceCode(ctx, serviceCode)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("service job not found")
		}
		return nil, err
	}
	return serviceJob, nil
}

// UpdateServiceJob updates a service job
func (u *ServiceJobUsecase) UpdateServiceJob(ctx context.Context, id uint, req interfaces.UpdateServiceJobRequest) (*models.ServiceJob, error) {
	serviceJob, err := u.repo.ServiceJob.GetByID(ctx, id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("service job not found")
		}
		return nil, err
	}

	// Validate entities exist if being updated
	if req.CustomerID != nil && *req.CustomerID != serviceJob.CustomerID {
		_, err := u.repo.Customer.GetByID(ctx, *req.CustomerID)
		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return nil, errors.New("customer not found")
			}
			return nil, err
		}
	}

	if req.VehicleID != nil && *req.VehicleID != serviceJob.VehicleID {
		_, err := u.repo.CustomerVehicle.GetByID(ctx, *req.VehicleID)
		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return nil, errors.New("vehicle not found")
			}
			return nil, err
		}
	}

	if req.TechnicianID != nil && (serviceJob.TechnicianID == nil || *req.TechnicianID != *serviceJob.TechnicianID) {
		_, err := u.repo.User.GetByID(ctx, *req.TechnicianID)
		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return nil, errors.New("technician not found")
			}
			return nil, err
		}
	}

	if req.ReceivedByUserID != nil && *req.ReceivedByUserID != serviceJob.ReceivedByUserID {
		_, err := u.repo.User.GetByID(ctx, *req.ReceivedByUserID)
		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return nil, errors.New("received by user not found")
			}
			return nil, err
		}
	}

	if req.OutletID != nil && *req.OutletID != serviceJob.OutletID {
		_, err := u.repo.Outlet.GetByID(ctx, *req.OutletID)
		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return nil, errors.New("outlet not found")
			}
			return nil, err
		}
	}

	// Update fields if provided
	if req.CustomerID != nil {
		serviceJob.CustomerID = *req.CustomerID
	}
	if req.VehicleID != nil {
		serviceJob.VehicleID = *req.VehicleID
	}
	if req.TechnicianID != nil {
		serviceJob.TechnicianID = req.TechnicianID
	}
	if req.ReceivedByUserID != nil {
		serviceJob.ReceivedByUserID = *req.ReceivedByUserID
	}
	if req.OutletID != nil {
		serviceJob.OutletID = *req.OutletID
	}
	if req.ProblemDescription != nil {
		serviceJob.ProblemDescription = *req.ProblemDescription
	}
	if req.TechnicianNotes != nil {
		serviceJob.TechnicianNotes = req.TechnicianNotes
	}
	if req.Status != nil {
		serviceJob.Status = *req.Status
	}
	if req.ServiceInDate != nil {
		serviceJob.ServiceInDate = *req.ServiceInDate
	}
	if req.PickedUpDate != nil {
		serviceJob.PickedUpDate = req.PickedUpDate
	}
	if req.ComplainDate != nil {
		serviceJob.ComplainDate = req.ComplainDate
	}
	if req.WarrantyExpiresAt != nil {
		serviceJob.WarrantyExpiresAt = req.WarrantyExpiresAt
	}
	if req.NextServiceReminderDate != nil {
		serviceJob.NextServiceReminderDate = req.NextServiceReminderDate
	}
	if req.DownPayment != nil {
		serviceJob.DownPayment = *req.DownPayment
	}
	if req.GrandTotal != nil {
		serviceJob.GrandTotal = *req.GrandTotal
	}
	if req.TechnicianCommission != nil {
		serviceJob.TechnicianCommission = *req.TechnicianCommission
	}
	if req.ShopProfit != nil {
		serviceJob.ShopProfit = *req.ShopProfit
	}
	serviceJob.UpdatedAt = time.Now()

	if err := u.repo.ServiceJob.Update(ctx, serviceJob); err != nil {
		return nil, err
	}

	return serviceJob, nil
}

// DeleteServiceJob deletes a service job
func (u *ServiceJobUsecase) DeleteServiceJob(ctx context.Context, id uint) error {
	_, err := u.repo.ServiceJob.GetByID(ctx, id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New("service job not found")
		}
		return err
	}

	// TODO: Add business logic checks (e.g., can't delete if status is completed)

	return u.repo.ServiceJob.Delete(ctx, id)
}

// ListServiceJobs retrieves service jobs with pagination
func (u *ServiceJobUsecase) ListServiceJobs(ctx context.Context, limit, offset int) ([]*models.ServiceJob, error) {
	if limit <= 0 {
		limit = 10
	}
	if offset < 0 {
		offset = 0
	}

	return u.repo.ServiceJob.List(ctx, limit, offset)
}

// GetServiceJobsByCustomer retrieves service jobs by customer
func (u *ServiceJobUsecase) GetServiceJobsByCustomer(ctx context.Context, customerID uint) ([]*models.ServiceJob, error) {
	return u.repo.ServiceJob.GetByCustomerID(ctx, customerID)
}

// GetServiceJobsByVehicle retrieves service jobs by vehicle
func (u *ServiceJobUsecase) GetServiceJobsByVehicle(ctx context.Context, vehicleID uint) ([]*models.ServiceJob, error) {
	return u.repo.ServiceJob.GetByVehicleID(ctx, vehicleID)
}

// GetServiceJobsByTechnician retrieves service jobs by technician
func (u *ServiceJobUsecase) GetServiceJobsByTechnician(ctx context.Context, technicianID uint) ([]*models.ServiceJob, error) {
	return u.repo.ServiceJob.GetByTechnicianID(ctx, technicianID)
}

// GetServiceJobsByOutlet retrieves service jobs by outlet
func (u *ServiceJobUsecase) GetServiceJobsByOutlet(ctx context.Context, outletID uint) ([]*models.ServiceJob, error) {
	return u.repo.ServiceJob.GetByOutletID(ctx, outletID)
}

// GetServiceJobsByStatus retrieves service jobs by status
func (u *ServiceJobUsecase) GetServiceJobsByStatus(ctx context.Context, status models.ServiceStatusEnum) ([]*models.ServiceJob, error) {
	return u.repo.ServiceJob.GetByStatus(ctx, status)
}

// UpdateServiceJobStatus updates service job status and creates history
func (u *ServiceJobUsecase) UpdateServiceJobStatus(ctx context.Context, id uint, status models.ServiceStatusEnum, userID uint, notes *string) error {
	// Validate service job exists
	serviceJob, err := u.repo.ServiceJob.GetByID(ctx, id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New("service job not found")
		}
		return err
	}

	// Validate user exists
	_, err = u.repo.User.GetByID(ctx, userID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New("user not found")
		}
		return err
	}

	// Handle status transition logic
	switch status {
	case models.ServiceStatusDikerjakan:
		// When starting work, validate technician is assigned
		if serviceJob.TechnicianID == nil {
			return errors.New("technician must be assigned before starting work")
		}
	case models.ServiceStatusSelesai:
		// When completing, ensure all details are calculated
		if err := u.CalculateServiceJobTotals(ctx, id); err != nil {
			fmt.Printf("Warning: Failed to calculate totals: %v\n", err)
		}
	case models.ServiceStatusDiambil:
		// When picked up, ensure all payments are settled
		// This is where you could add payment validation logic
		// Update picked up date
		now := time.Now()
		updateReq := interfaces.UpdateServiceJobRequest{
			PickedUpDate: &now,
		}
		if _, err := u.UpdateServiceJob(ctx, id, updateReq); err != nil {
			fmt.Printf("Warning: Failed to update picked up date: %v\n", err)
		}
	}

	// Update status
	if err := u.repo.ServiceJob.UpdateStatus(ctx, id, status); err != nil {
		return err
	}

	// Create history entry
	historyNotes := notes
	if historyNotes == nil {
		statusNote := fmt.Sprintf("Status changed to %s", string(status))
		historyNotes = &statusNote
	}
	
	historyReq := interfaces.CreateServiceJobHistoryRequest{
		ServiceJobID: id,
		UserID:       userID,
		Notes:        historyNotes,
	}
	_, err = u.createServiceJobHistory(ctx, historyReq)
	if err != nil {
		// Don't fail the entire operation for history creation failure
		fmt.Printf("Failed to create service job history: %v\n", err)
	}

	return nil
}

// ============= Queue Management Methods =============

// GetServiceJobQueue retrieves all service jobs in queue for an outlet
func (u *ServiceJobUsecase) GetServiceJobQueue(ctx context.Context, outletID uint) ([]*models.ServiceJob, error) {
	// Validate outlet exists
	_, err := u.repo.Outlet.GetByID(ctx, outletID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("outlet not found")
		}
		return nil, err
	}

	// Get all service jobs in queue order (Antri and Dikerjakan status)
	return u.repo.ServiceJob.GetQueueByOutletID(ctx, outletID)
}

// GetTodayServiceJobQueue retrieves today's service jobs in queue for an outlet
func (u *ServiceJobUsecase) GetTodayServiceJobQueue(ctx context.Context, outletID uint) ([]*models.ServiceJob, error) {
	// Validate outlet exists
	_, err := u.repo.Outlet.GetByID(ctx, outletID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("outlet not found")
		}
		return nil, err
	}

	// Get today's service jobs in queue order
	return u.repo.ServiceJob.GetTodayQueueByOutletID(ctx, outletID)
}

// ReorderServiceJobQueue reorders service jobs in the queue
func (u *ServiceJobUsecase) ReorderServiceJobQueue(ctx context.Context, outletID uint, serviceJobIDs []uint) error {
	// Validate outlet exists
	_, err := u.repo.Outlet.GetByID(ctx, outletID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New("outlet not found")
		}
		return err
	}

	// Validate all service jobs exist and belong to the outlet
	for i, serviceJobID := range serviceJobIDs {
		serviceJob, err := u.repo.ServiceJob.GetByID(ctx, serviceJobID)
		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return fmt.Errorf("service job with ID %d not found", serviceJobID)
			}
			return err
		}
		
		if serviceJob.OutletID != outletID {
			return fmt.Errorf("service job with ID %d does not belong to outlet %d", serviceJobID, outletID)
		}

		// Update queue number (1-based)
		if err := u.repo.ServiceJob.UpdateQueueNumber(ctx, serviceJobID, i+1); err != nil {
			return fmt.Errorf("failed to update queue number for service job %d: %v", serviceJobID, err)
		}
	}

	return nil
}

// UpdateServiceJobStatusWithTechnician updates service job status and assigns technician
func (u *ServiceJobUsecase) UpdateServiceJobStatusWithTechnician(ctx context.Context, id uint, status models.ServiceStatusEnum, userID uint, technicianID *uint, notes *string) error {
	// Validate service job exists
	_, err := u.repo.ServiceJob.GetByID(ctx, id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New("service job not found")
		}
		return err
	}

	// Validate user exists
	_, err = u.repo.User.GetByID(ctx, userID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New("user not found")
		}
		return err
	}

	// Validate technician exists if provided
	if technicianID != nil {
		_, err := u.repo.User.GetByID(ctx, *technicianID)
		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return errors.New("technician not found")
			}
			return err
		}
	}

	// Update technician assignment if provided
	if technicianID != nil {
		updateReq := interfaces.UpdateServiceJobRequest{
			TechnicianID: technicianID,
		}
		if _, err := u.UpdateServiceJob(ctx, id, updateReq); err != nil {
			return err
		}
	}

	// Update status using existing method
	return u.UpdateServiceJobStatus(ctx, id, status, userID, notes)
}

// CalculateServiceJobTotals calculates and updates service job totals
func (u *ServiceJobUsecase) CalculateServiceJobTotals(ctx context.Context, serviceJobID uint) error {
	// Get service job
	_, err := u.repo.ServiceJob.GetByID(ctx, serviceJobID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New("service job not found")
		}
		return err
	}

	// Get service details
	serviceDetails, err := u.repo.ServiceDetail.GetByServiceJobID(ctx, serviceJobID)
	if err != nil {
		return err
	}

	// Calculate totals
	var grandTotal float64
	var totalCost float64
	var technicianCommission float64

	for _, detail := range serviceDetails {
		itemTotal := detail.PricePerItem * float64(detail.Quantity)
		grandTotal += itemTotal
		totalCost += detail.CostPerItem * float64(detail.Quantity)
		
		// Calculate technician commission (example: 10% of service items)
		if detail.ItemType == "service" {
			technicianCommission += itemTotal * 0.10
		}
	}

	// Calculate shop profit
	shopProfit := grandTotal - totalCost - technicianCommission

	// Update service job
	updateReq := interfaces.UpdateServiceJobRequest{
		GrandTotal:           &grandTotal,
		TechnicianCommission: &technicianCommission,
		ShopProfit:           &shopProfit,
	}

	_, err = u.UpdateServiceJob(ctx, serviceJobID, updateReq)
	return err
}

// ServiceDetailUsecase implements the service detail usecase interface
type ServiceDetailUsecase struct {
	repo *repository.RepositoryManager
}

// NewServiceDetailUsecase creates a new service detail usecase
func NewServiceDetailUsecase(repo *repository.RepositoryManager) interfaces.ServiceDetailUsecase {
	return &ServiceDetailUsecase{repo: repo}
}

// CreateServiceDetail creates a new service detail
func (u *ServiceDetailUsecase) CreateServiceDetail(ctx context.Context, req interfaces.CreateServiceDetailRequest) (*models.ServiceDetail, error) {
	// Validate service job exists
	_, err := u.repo.ServiceJob.GetByID(ctx, req.ServiceJobID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("service job not found")
		}
		return nil, err
	}

	// Validate item exists based on type
	if req.ItemType == "service" {
		_, err := u.repo.Service.GetByID(ctx, req.ItemID)
		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return nil, errors.New("service not found")
			}
			return nil, err
		}
	} else if req.ItemType == "product" {
		_, err := u.repo.Product.GetByID(ctx, req.ItemID)
		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return nil, errors.New("product not found")
			}
			return nil, err
		}
	}

	// Validate serial number if provided
	if req.SerialNumberUsed != nil {
		_, err := u.repo.ProductSerialNumber.GetBySerialNumber(ctx, *req.SerialNumberUsed)
		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return nil, errors.New("serial number not found")
			}
			return nil, err
		}
	}

	serviceDetail := &models.ServiceDetail{
		ServiceJobID:     req.ServiceJobID,
		ItemID:           req.ItemID,
		ItemType:         req.ItemType,
		Description:      req.Description,
		SerialNumberUsed: req.SerialNumberUsed,
		Quantity:         req.Quantity,
		PricePerItem:     req.PricePerItem,
		CostPerItem:      req.CostPerItem,
	}

	if err := u.repo.ServiceDetail.Create(ctx, serviceDetail); err != nil {
		return nil, err
	}

	return serviceDetail, nil
}

// GetServiceDetail retrieves a service detail by ID
func (u *ServiceDetailUsecase) GetServiceDetail(ctx context.Context, id uint) (*models.ServiceDetail, error) {
	serviceDetail, err := u.repo.ServiceDetail.GetByID(ctx, id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("service detail not found")
		}
		return nil, err
	}
	return serviceDetail, nil
}

// UpdateServiceDetail updates a service detail
func (u *ServiceDetailUsecase) UpdateServiceDetail(ctx context.Context, id uint, req interfaces.UpdateServiceDetailRequest) (*models.ServiceDetail, error) {
	serviceDetail, err := u.repo.ServiceDetail.GetByID(ctx, id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("service detail not found")
		}
		return nil, err
	}

	// Validate service job exists if being updated
	if req.ServiceJobID != nil && *req.ServiceJobID != serviceDetail.ServiceJobID {
		_, err := u.repo.ServiceJob.GetByID(ctx, *req.ServiceJobID)
		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return nil, errors.New("service job not found")
			}
			return nil, err
		}
	}

	// Validate item exists based on type if being updated
	if req.ItemID != nil && req.ItemType != nil {
		if *req.ItemType == "service" {
			_, err := u.repo.Service.GetByID(ctx, *req.ItemID)
			if err != nil {
				if errors.Is(err, gorm.ErrRecordNotFound) {
					return nil, errors.New("service not found")
				}
				return nil, err
			}
		} else if *req.ItemType == "product" {
			_, err := u.repo.Product.GetByID(ctx, *req.ItemID)
			if err != nil {
				if errors.Is(err, gorm.ErrRecordNotFound) {
					return nil, errors.New("product not found")
				}
				return nil, err
			}
		}
	}

	// Validate serial number if provided
	if req.SerialNumberUsed != nil {
		_, err := u.repo.ProductSerialNumber.GetBySerialNumber(ctx, *req.SerialNumberUsed)
		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return nil, errors.New("serial number not found")
			}
			return nil, err
		}
	}

	// Update fields if provided
	if req.ServiceJobID != nil {
		serviceDetail.ServiceJobID = *req.ServiceJobID
	}
	if req.ItemID != nil {
		serviceDetail.ItemID = *req.ItemID
	}
	if req.ItemType != nil {
		serviceDetail.ItemType = *req.ItemType
	}
	if req.Description != nil {
		serviceDetail.Description = *req.Description
	}
	if req.SerialNumberUsed != nil {
		serviceDetail.SerialNumberUsed = req.SerialNumberUsed
	}
	if req.Quantity != nil {
		serviceDetail.Quantity = *req.Quantity
	}
	if req.PricePerItem != nil {
		serviceDetail.PricePerItem = *req.PricePerItem
	}
	if req.CostPerItem != nil {
		serviceDetail.CostPerItem = *req.CostPerItem
	}

	if err := u.repo.ServiceDetail.Update(ctx, serviceDetail); err != nil {
		return nil, err
	}

	return serviceDetail, nil
}

// DeleteServiceDetail deletes a service detail
func (u *ServiceDetailUsecase) DeleteServiceDetail(ctx context.Context, id uint) error {
	_, err := u.repo.ServiceDetail.GetByID(ctx, id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New("service detail not found")
		}
		return err
	}

	return u.repo.ServiceDetail.Delete(ctx, id)
}

// ListServiceDetails retrieves service details with pagination
func (u *ServiceDetailUsecase) ListServiceDetails(ctx context.Context, limit, offset int) ([]*models.ServiceDetail, error) {
	if limit <= 0 {
		limit = 10
	}
	if offset < 0 {
		offset = 0
	}

	return u.repo.ServiceDetail.List(ctx, limit, offset)
}

// GetServiceDetailsByServiceJob retrieves service details by service job
func (u *ServiceDetailUsecase) GetServiceDetailsByServiceJob(ctx context.Context, serviceJobID uint) ([]*models.ServiceDetail, error) {
	return u.repo.ServiceDetail.GetByServiceJobID(ctx, serviceJobID)
}

// DeleteServiceDetailsByServiceJob deletes service details by service job
func (u *ServiceDetailUsecase) DeleteServiceDetailsByServiceJob(ctx context.Context, serviceJobID uint) error {
	return u.repo.ServiceDetail.DeleteByServiceJobID(ctx, serviceJobID)
}

// ServiceJobHistoryUsecase implements the service job history usecase interface
type ServiceJobHistoryUsecase struct {
	repo *repository.RepositoryManager
}

// NewServiceJobHistoryUsecase creates a new service job history usecase
func NewServiceJobHistoryUsecase(repo *repository.RepositoryManager) interfaces.ServiceJobHistoryUsecase {
	return &ServiceJobHistoryUsecase{repo: repo}
}

// CreateServiceJobHistory creates a new service job history
func (u *ServiceJobHistoryUsecase) CreateServiceJobHistory(ctx context.Context, req interfaces.CreateServiceJobHistoryRequest) (*models.ServiceJobHistory, error) {
	// Validate service job exists
	_, err := u.repo.ServiceJob.GetByID(ctx, req.ServiceJobID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("service job not found")
		}
		return nil, err
	}

	// Validate user exists
	_, err = u.repo.User.GetByID(ctx, req.UserID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("user not found")
		}
		return nil, err
	}

	history := &models.ServiceJobHistory{
		ServiceJobID: req.ServiceJobID,
		UserID:       req.UserID,
		Notes:        req.Notes,
		ChangedAt:    time.Now(),
	}

	if err := u.repo.ServiceJobHistory.Create(ctx, history); err != nil {
		return nil, err
	}

	return history, nil
}

// GetServiceJobHistory retrieves a service job history by ID
func (u *ServiceJobHistoryUsecase) GetServiceJobHistory(ctx context.Context, id uint) (*models.ServiceJobHistory, error) {
	history, err := u.repo.ServiceJobHistory.GetByID(ctx, id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("service job history not found")
		}
		return nil, err
	}
	return history, nil
}

// ListServiceJobHistories retrieves service job histories with pagination
func (u *ServiceJobHistoryUsecase) ListServiceJobHistories(ctx context.Context, limit, offset int) ([]*models.ServiceJobHistory, error) {
	if limit <= 0 {
		limit = 10
	}
	if offset < 0 {
		offset = 0
	}

	return u.repo.ServiceJobHistory.List(ctx, limit, offset)
}

// GetServiceJobHistoriesByServiceJob retrieves service job histories by service job
func (u *ServiceJobHistoryUsecase) GetServiceJobHistoriesByServiceJob(ctx context.Context, serviceJobID uint) ([]*models.ServiceJobHistory, error) {
	return u.repo.ServiceJobHistory.GetByServiceJobID(ctx, serviceJobID)
}

// GetServiceJobHistoriesByUser retrieves service job histories by user
func (u *ServiceJobHistoryUsecase) GetServiceJobHistoriesByUser(ctx context.Context, userID uint) ([]*models.ServiceJobHistory, error) {
	return u.repo.ServiceJobHistory.GetByUserID(ctx, userID)
}