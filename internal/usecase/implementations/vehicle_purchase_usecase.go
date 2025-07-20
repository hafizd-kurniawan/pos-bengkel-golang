package implementations

import (
	"boilerplate/internal/models"
	"boilerplate/internal/repository"
	"boilerplate/internal/usecase/interfaces"
	"context"
	"errors"
	"time"
)

// vehiclePurchaseUseCase implements interfaces.VehiclePurchaseUseCase
type vehiclePurchaseUseCase struct {
	repos *repository.RepositoryManager
}

// NewVehiclePurchaseUseCase creates a new vehicle purchase use case
func NewVehiclePurchaseUseCase(repos *repository.RepositoryManager) interfaces.VehiclePurchaseUseCase {
	return &vehiclePurchaseUseCase{repos: repos}
}

// PurchaseVehicle handles the complete vehicle purchase process
func (uc *vehiclePurchaseUseCase) PurchaseVehicle(ctx context.Context, req *interfaces.CreateVehiclePurchaseRequest) (*interfaces.VehiclePurchaseResponse, error) {
	// Validate customer exists
	customer, err := uc.repos.Customer.GetByID(ctx, req.VehicleData.CustomerID)
	if err != nil {
		return nil, errors.New("customer not found")
	}

	// Check for duplicate plate number, chassis number, engine number
	if existing, _ := uc.repos.Vehicle.GetByPlateNumber(ctx, req.VehicleData.PlateNumber); existing != nil {
		return nil, errors.New("vehicle with this plate number already exists")
	}
	if existing, _ := uc.repos.Vehicle.GetByChassisNumber(ctx, req.VehicleData.ChassisNumber); existing != nil {
		return nil, errors.New("vehicle with this chassis number already exists")
	}
	if existing, _ := uc.repos.Vehicle.GetByEngineNumber(ctx, req.VehicleData.EngineNumber); existing != nil {
		return nil, errors.New("vehicle with this engine number already exists")
	}

	// Create vehicle record
	vehicle := &models.Vehicle{
		CustomerID:       &req.VehicleData.CustomerID,
		PlateNumber:      req.VehicleData.PlateNumber,
		Brand:            req.VehicleData.Brand,
		Model:            req.VehicleData.Model,
		Type:             req.VehicleData.Type,
		ProductionYear:   req.VehicleData.ProductionYear,
		ChassisNumber:    req.VehicleData.ChassisNumber,
		EngineNumber:     req.VehicleData.EngineNumber,
		Color:            req.VehicleData.Color,
		Mileage:          req.VehicleData.Mileage,
		FuelType:         req.VehicleData.FuelType,
		Transmission:     req.VehicleData.Transmission,
		OwnershipStatus:  models.VehicleOwnershipShowroom, // Vehicle now owned by showroom
		ConditionStatus:  req.VehicleData.ConditionStatus,
		SaleStatus:       models.VehicleSaleStatusNotForSale, // Not for sale until reconditioning
		PurchasePrice:    &req.PurchasePrice,
		EstimatedValue:   req.VehicleData.EstimatedValue,
		ConditionNotes:   req.VehicleData.ConditionNotes,
		InternalNotes:    req.VehicleData.InternalNotes,
		CreatedBy:        req.CreatedBy,
	}

	if err := uc.repos.Vehicle.Create(ctx, vehicle); err != nil {
		return nil, err
	}

	// Create purchase transaction record
	purchaseTransaction := &models.VehiclePurchaseTransaction{
		VehicleID:         vehicle.VehicleID,
		CustomerID:        req.VehicleData.CustomerID,
		PurchasePrice:     req.PurchasePrice,
		PurchaseDate:      time.Now(),
		PaymentMethod:     req.PaymentMethod,
		TransactionStatus: models.TransactionStatusSukses,
		EvaluationNotes:   req.EvaluationNotes,
		PaymentReference:  req.PaymentReference,
		CreatedBy:         req.CreatedBy,
	}

	if err := uc.repos.VehiclePurchaseTransaction.Create(ctx, purchaseTransaction); err != nil {
		return nil, err
	}

	// Load relationships for response
	vehicle.Customer = customer
	purchaseTransaction.Vehicle = vehicle
	purchaseTransaction.Customer = customer

	return &interfaces.VehiclePurchaseResponse{
		Vehicle:             vehicle,
		PurchaseTransaction: purchaseTransaction,
	}, nil
}

// CreateReconditioningJob creates a new reconditioning job for a vehicle
func (uc *vehiclePurchaseUseCase) CreateReconditioningJob(ctx context.Context, req *interfaces.CreateReconditioningJobRequest) (*models.VehicleReconditioningJob, error) {
	// Validate vehicle exists and is owned by showroom
	vehicle, err := uc.repos.Vehicle.GetByID(ctx, req.VehicleID)
	if err != nil {
		return nil, errors.New("vehicle not found")
	}

	if vehicle.OwnershipStatus != models.VehicleOwnershipShowroom {
		return nil, errors.New("vehicle must be owned by showroom to start reconditioning")
	}

	// Create reconditioning job
	job := &models.VehicleReconditioningJob{
		VehicleID:            req.VehicleID,
		JobTitle:             req.JobTitle,
		JobDescription:       req.JobDescription,
		EstimatedCost:        req.EstimatedCost,
		Status:               models.ReconditioningJobStatusPending,
		AssignedTechnicianID: req.AssignedTechnicianID,
		Notes:                req.Notes,
		CreatedBy:            req.CreatedBy,
	}

	if err := uc.repos.VehicleReconditioningJob.Create(ctx, job); err != nil {
		return nil, err
	}

	// Update vehicle ownership to workshop
	vehicle.OwnershipStatus = models.VehicleOwnershipWorkshop
	if err := uc.repos.Vehicle.Update(ctx, vehicle); err != nil {
		return nil, err
	}

	return job, nil
}

// AddReconditioningDetail adds parts or services to a reconditioning job
func (uc *vehiclePurchaseUseCase) AddReconditioningDetail(ctx context.Context, req *interfaces.AddReconditioningDetailRequest) (*models.ReconditioningDetail, error) {
	// Validate reconditioning job exists and is active
	job, err := uc.repos.VehicleReconditioningJob.GetByID(ctx, req.ReconditioningJobID)
	if err != nil {
		return nil, errors.New("reconditioning job not found")
	}

	if job.Status != models.ReconditioningJobStatusPending && job.Status != models.ReconditioningJobStatusInProgress {
		return nil, errors.New("cannot add details to completed or cancelled job")
	}

	// Validate product or service exists based on detail type
	if req.DetailType == models.ReconditioningDetailTypePart {
		if req.ProductID == nil {
			return nil, errors.New("product_id is required for part details")
		}
		if _, err := uc.repos.Product.GetByID(ctx, *req.ProductID); err != nil {
			return nil, errors.New("product not found")
		}
	} else if req.DetailType == models.ReconditioningDetailTypeService {
		if req.ServiceID == nil {
			return nil, errors.New("service_id is required for service details")
		}
		if _, err := uc.repos.Service.GetByID(ctx, *req.ServiceID); err != nil {
			return nil, errors.New("service not found")
		}
	}

	// Calculate total price
	totalPrice := req.UnitPrice * float64(req.Quantity)

	// Create reconditioning detail
	detail := &models.ReconditioningDetail{
		ReconditioningJobID: req.ReconditioningJobID,
		DetailType:          req.DetailType,
		ProductID:           req.ProductID,
		ServiceID:           req.ServiceID,
		Description:         req.Description,
		Quantity:            req.Quantity,
		UnitPrice:           req.UnitPrice,
		TotalPrice:          totalPrice,
		UsageDate:           time.Now(),
		Notes:               req.Notes,
		CreatedBy:           req.CreatedBy,
	}

	if err := uc.repos.ReconditioningDetail.Create(ctx, detail); err != nil {
		return nil, err
	}

	// Update job status to in progress if it was pending
	if job.Status == models.ReconditioningJobStatusPending {
		job.Status = models.ReconditioningJobStatusInProgress
		job.StartDate = &[]time.Time{time.Now()}[0]
		if err := uc.repos.VehicleReconditioningJob.Update(ctx, job); err != nil {
			return nil, err
		}
	}

	// If it's a part, reduce inventory
	if req.DetailType == models.ReconditioningDetailTypePart && req.ProductID != nil {
		product, err := uc.repos.Product.GetByID(ctx, *req.ProductID)
		if err == nil && product.Stock >= req.Quantity {
			product.Stock -= req.Quantity
			uc.repos.Product.Update(ctx, product)
		}
	}

	return detail, nil
}

// CompleteReconditioningJob completes a reconditioning job
func (uc *vehiclePurchaseUseCase) CompleteReconditioningJob(ctx context.Context, req *interfaces.CompleteReconditioningJobRequest) (*interfaces.ReconditioningJobResponse, error) {
	// Get reconditioning job
	job, err := uc.repos.VehicleReconditioningJob.GetByID(ctx, req.ReconditioningJobID)
	if err != nil {
		return nil, errors.New("reconditioning job not found")
	}

	if job.Status != models.ReconditioningJobStatusInProgress {
		return nil, errors.New("only in-progress jobs can be completed")
	}

	// Get all details to calculate actual cost if not provided
	details, err := uc.repos.ReconditioningDetail.GetByReconditioningJobID(ctx, req.ReconditioningJobID)
	if err != nil {
		return nil, err
	}

	actualCost := req.ActualCost
	if actualCost == nil {
		// Calculate from details
		var totalCost float64
		for _, detail := range details {
			totalCost += detail.TotalPrice
		}
		actualCost = &totalCost
	}

	// Update job completion
	now := time.Now()
	job.Status = models.ReconditioningJobStatusCompleted
	job.CompletionDate = &now
	job.ActualCost = actualCost
	if req.CompletionNotes != nil {
		job.Notes = req.CompletionNotes
	}

	if err := uc.repos.VehicleReconditioningJob.Update(ctx, job); err != nil {
		return nil, err
	}

	// Update vehicle ownership back to showroom and ready for sale
	vehicle, err := uc.repos.Vehicle.GetByID(ctx, job.VehicleID)
	if err != nil {
		return nil, err
	}

	vehicle.OwnershipStatus = models.VehicleOwnershipShowroom
	vehicle.SaleStatus = models.VehicleSaleStatusForSale
	if err := uc.repos.Vehicle.Update(ctx, vehicle); err != nil {
		return nil, err
	}

	return &interfaces.ReconditioningJobResponse{
		Job:     job,
		Details: details,
	}, nil
}

// GetVehicleByID retrieves a vehicle by ID
func (uc *vehiclePurchaseUseCase) GetVehicleByID(ctx context.Context, id uint) (*models.Vehicle, error) {
	return uc.repos.Vehicle.GetByID(ctx, id)
}

// GetReconditioningJobsByVehicleID retrieves reconditioning jobs for a vehicle
func (uc *vehiclePurchaseUseCase) GetReconditioningJobsByVehicleID(ctx context.Context, vehicleID uint) ([]*models.VehicleReconditioningJob, error) {
	return uc.repos.VehicleReconditioningJob.GetByVehicleID(ctx, vehicleID)
}

// GetReconditioningJobByID retrieves a reconditioning job with details
func (uc *vehiclePurchaseUseCase) GetReconditioningJobByID(ctx context.Context, id uint) (*interfaces.ReconditioningJobResponse, error) {
	job, err := uc.repos.VehicleReconditioningJob.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}

	details, err := uc.repos.ReconditioningDetail.GetByReconditioningJobID(ctx, id)
	if err != nil {
		return nil, err
	}

	return &interfaces.ReconditioningJobResponse{
		Job:     job,
		Details: details,
	}, nil
}

// GetReconditioningDetailsByJobID retrieves reconditioning details for a job
func (uc *vehiclePurchaseUseCase) GetReconditioningDetailsByJobID(ctx context.Context, jobID uint) ([]*models.ReconditioningDetail, error) {
	return uc.repos.ReconditioningDetail.GetByReconditioningJobID(ctx, jobID)
}