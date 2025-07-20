package implementations

import (
	"boilerplate/internal/models"
	"boilerplate/internal/repository"
	"boilerplate/internal/usecase/interfaces"
	"context"
	"errors"
	"time"
)

// vehicleSalesUseCase implements interfaces.VehicleSalesUseCase
type vehicleSalesUseCase struct {
	repos *repository.RepositoryManager
}

// NewVehicleSalesUseCase creates a new vehicle sales use case
func NewVehicleSalesUseCase(repos *repository.RepositoryManager) interfaces.VehicleSalesUseCase {
	return &vehicleSalesUseCase{repos: repos}
}

// SellVehicle handles the complete vehicle sale process
func (uc *vehicleSalesUseCase) SellVehicle(ctx context.Context, req *interfaces.CreateVehicleSaleRequest) (*interfaces.VehicleSaleResponse, error) {
	// Validate vehicle exists and is for sale
	vehicle, err := uc.repos.Vehicle.GetByID(ctx, req.VehicleID)
	if err != nil {
		return nil, errors.New("vehicle not found")
	}

	if vehicle.SaleStatus != models.VehicleSaleStatusForSale {
		return nil, errors.New("vehicle is not available for sale")
	}

	// Validate customer exists
	customer, err := uc.repos.Customer.GetByID(ctx, req.CustomerID)
	if err != nil {
		return nil, errors.New("customer not found")
	}

	// Validate installment configuration if required
	if req.TransactionType == models.SalesTransactionTypeInstallment {
		if req.InstallmentConfig == nil {
			return nil, errors.New("installment configuration is required for installment sales")
		}
		if req.DownPayment == nil {
			return nil, errors.New("down payment is required for installment sales")
		}
		if *req.DownPayment >= req.SalePrice {
			return nil, errors.New("down payment cannot be greater than or equal to sale price")
		}
	}

	// Calculate profit
	profitAmount := uc.calculateProfitAmount(ctx, vehicle, req.SalePrice)

	// Create sales transaction
	salesTransaction := &models.VehicleSalesTransaction{
		VehicleID:         req.VehicleID,
		CustomerID:        req.CustomerID,
		SalePrice:         req.SalePrice,
		DownPayment:       req.DownPayment,
		SaleDate:          time.Now(),
		TransactionType:   req.TransactionType,
		PaymentMethod:     req.PaymentMethod,
		TransactionStatus: models.TransactionStatusSukses,
		PaymentReference:  req.PaymentReference,
		ProfitAmount:      &profitAmount,
		SalesPersonID:     req.SalesPersonID,
		Notes:             req.Notes,
		CreatedBy:         req.CreatedBy,
	}

	if err := uc.repos.VehicleSalesTransaction.Create(ctx, salesTransaction); err != nil {
		return nil, err
	}

	// Update vehicle status
	vehicle.SaleStatus = models.VehicleSaleStatusSold
	vehicle.OwnershipStatus = models.VehicleOwnershipCustomer
	vehicle.CustomerID = &req.CustomerID
	vehicle.SellingPrice = &req.SalePrice
	if err := uc.repos.Vehicle.Update(ctx, vehicle); err != nil {
		return nil, err
	}

	response := &interfaces.VehicleSaleResponse{
		Vehicle:          vehicle,
		SalesTransaction: salesTransaction,
		ProfitAmount:     &profitAmount,
	}

	// Create installment if required
	if req.TransactionType == models.SalesTransactionTypeInstallment {
		installment, err := uc.createInstallment(ctx, salesTransaction, req)
		if err != nil {
			return nil, err
		}
		response.Installment = installment
	}

	// Load relationships
	salesTransaction.Vehicle = vehicle
	salesTransaction.Customer = customer

	return response, nil
}

// MarkVehicleForSale marks a vehicle as available for sale
func (uc *vehicleSalesUseCase) MarkVehicleForSale(ctx context.Context, vehicleID uint, sellingPrice float64) error {
	vehicle, err := uc.repos.Vehicle.GetByID(ctx, vehicleID)
	if err != nil {
		return errors.New("vehicle not found")
	}

	if vehicle.OwnershipStatus != models.VehicleOwnershipShowroom {
		return errors.New("vehicle must be owned by showroom to be marked for sale")
	}

	vehicle.SaleStatus = models.VehicleSaleStatusForSale
	vehicle.SellingPrice = &sellingPrice

	return uc.repos.Vehicle.Update(ctx, vehicle)
}

// createInstallment creates an installment plan for a sales transaction
func (uc *vehicleSalesUseCase) createInstallment(ctx context.Context, transaction *models.VehicleSalesTransaction, req *interfaces.CreateVehicleSaleRequest) (*models.VehicleInstallment, error) {
	config := req.InstallmentConfig
	
	// Parse start date
	startDate, err := time.Parse("2006-01-02", config.StartDate)
	if err != nil {
		return nil, errors.New("invalid start date format")
	}

	// Calculate installment amount
	remainingAmount := transaction.SalePrice - *transaction.DownPayment
	installmentAmount := remainingAmount / float64(config.NumberOfInstallments)

	// Apply interest if specified
	if config.InterestRate != nil && *config.InterestRate > 0 {
		interestAmount := remainingAmount * (*config.InterestRate / 100)
		installmentAmount = (remainingAmount + interestAmount) / float64(config.NumberOfInstallments)
	}

	// Calculate end date
	endDate := startDate.AddDate(0, config.NumberOfInstallments, 0)

	// Create installment record
	installment := &models.VehicleInstallment{
		SalesTransactionID:   transaction.SalesTransactionID,
		TotalAmount:          transaction.SalePrice,
		DownPayment:          *transaction.DownPayment,
		InstallmentAmount:    installmentAmount,
		NumberOfInstallments: config.NumberOfInstallments,
		InterestRate:         config.InterestRate,
		StartDate:            startDate,
		EndDate:              endDate,
		Status:               models.InstallmentStatusActive,
		RemainingBalance:     remainingAmount,
		CreatedBy:            req.CreatedBy,
	}

	if err := uc.repos.VehicleInstallment.Create(ctx, installment); err != nil {
		return nil, err
	}

	// Generate payment schedule
	if err := uc.generatePaymentSchedule(ctx, installment); err != nil {
		return nil, err
	}

	return installment, nil
}

// generatePaymentSchedule creates the payment schedule for an installment
func (uc *vehicleSalesUseCase) generatePaymentSchedule(ctx context.Context, installment *models.VehicleInstallment) error {
	for i := 1; i <= installment.NumberOfInstallments; i++ {
		dueDate := installment.StartDate.AddDate(0, i, 0)
		
		payment := &models.InstallmentPayment{
			InstallmentID:   installment.InstallmentID,
			PaymentNumber:   i,
			DueDate:         dueDate,
			DueAmount:       installment.InstallmentAmount,
			PaymentStatus:   models.InstallmentPaymentStatusPending,
			CreatedBy:       installment.CreatedBy,
		}

		if err := uc.repos.InstallmentPayment.Create(ctx, payment); err != nil {
			return err
		}
	}
	return nil
}

// ProcessInstallmentPayment processes a payment towards an installment
func (uc *vehicleSalesUseCase) ProcessInstallmentPayment(ctx context.Context, req *interfaces.ProcessInstallmentPaymentRequest) (*interfaces.InstallmentPaymentResponse, error) {
	// Get payment record
	payment, err := uc.repos.InstallmentPayment.GetByID(ctx, req.PaymentID)
	if err != nil {
		return nil, errors.New("payment not found")
	}

	if payment.PaymentStatus == models.InstallmentPaymentStatusPaid {
		return nil, errors.New("payment has already been processed")
	}

	// Get installment
	installment, err := uc.repos.VehicleInstallment.GetByID(ctx, payment.InstallmentID)
	if err != nil {
		return nil, errors.New("installment not found")
	}

	// Calculate late fee if payment is overdue
	var lateFee float64
	now := time.Now()
	if now.After(payment.DueDate) && payment.PaymentStatus == models.InstallmentPaymentStatusPending {
		lateFee, _ = uc.CalculateLateFee(ctx, req.PaymentID)
		payment.PaymentStatus = models.InstallmentPaymentStatusLate
	}

	// Update payment record
	payment.PaymentDate = &now
	payment.PaidAmount = &req.PaidAmount
	payment.LateFee = &lateFee
	payment.PaymentMethod = &req.PaymentMethod
	payment.PaymentReference = req.PaymentReference
	payment.Notes = req.Notes
	payment.PaymentStatus = models.InstallmentPaymentStatusPaid

	if err := uc.repos.InstallmentPayment.Update(ctx, payment); err != nil {
		return nil, err
	}

	// Update installment remaining balance
	installment.RemainingBalance -= req.PaidAmount
	if installment.RemainingBalance <= 0 {
		installment.Status = models.InstallmentStatusCompleted
		installment.RemainingBalance = 0
	}

	if err := uc.repos.VehicleInstallment.Update(ctx, installment); err != nil {
		return nil, err
	}

	// Get next payment due
	var nextPaymentDue *time.Time
	if installment.Status == models.InstallmentStatusActive {
		nextPayments, err := uc.repos.InstallmentPayment.GetByInstallmentID(ctx, installment.InstallmentID)
		if err == nil {
			for _, np := range nextPayments {
				if np.PaymentStatus == models.InstallmentPaymentStatusPending {
					nextPaymentDue = &np.DueDate
					break
				}
			}
		}
	}

	return &interfaces.InstallmentPaymentResponse{
		Payment:            payment,
		RemainingBalance:   installment.RemainingBalance,
		NextPaymentDue:     nextPaymentDue,
	}, nil
}

// GenerateInstallmentSchedule generates the payment schedule for an installment
func (uc *vehicleSalesUseCase) GenerateInstallmentSchedule(ctx context.Context, installmentID uint) ([]*models.InstallmentPayment, error) {
	return uc.repos.InstallmentPayment.GetByInstallmentID(ctx, installmentID)
}

// GetOverduePayments retrieves all overdue payments
func (uc *vehicleSalesUseCase) GetOverduePayments(ctx context.Context) ([]*models.InstallmentPayment, error) {
	return uc.repos.InstallmentPayment.GetOverduePayments(ctx)
}

// CalculateLateFee calculates late fee for an overdue payment
func (uc *vehicleSalesUseCase) CalculateLateFee(ctx context.Context, paymentID uint) (float64, error) {
	payment, err := uc.repos.InstallmentPayment.GetByID(ctx, paymentID)
	if err != nil {
		return 0, err
	}

	now := time.Now()
	if !now.After(payment.DueDate) {
		return 0, nil // Not overdue
	}

	// Calculate days overdue
	daysOverdue := int(now.Sub(payment.DueDate).Hours() / 24)
	
	// Simple late fee calculation: 1% of due amount per month overdue (minimum 1 day = 1 month)
	monthsOverdue := (daysOverdue / 30) + 1
	lateFeeRate := 0.01 // 1% per month
	lateFee := payment.DueAmount * float64(monthsOverdue) * lateFeeRate

	return lateFee, nil
}

// GetVehiclesByStatus retrieves vehicles by sale status
func (uc *vehicleSalesUseCase) GetVehiclesByStatus(ctx context.Context, status models.VehicleSaleStatus) ([]*models.Vehicle, error) {
	return uc.repos.Vehicle.GetBySaleStatus(ctx, status)
}

// GetSalesTransactionByID retrieves a sales transaction by ID
func (uc *vehicleSalesUseCase) GetSalesTransactionByID(ctx context.Context, id uint) (*models.VehicleSalesTransaction, error) {
	return uc.repos.VehicleSalesTransaction.GetByID(ctx, id)
}

// GetInstallmentByID retrieves an installment by ID
func (uc *vehicleSalesUseCase) GetInstallmentByID(ctx context.Context, id uint) (*models.VehicleInstallment, error) {
	return uc.repos.VehicleInstallment.GetByID(ctx, id)
}

// GetInstallmentPaymentsByInstallmentID retrieves payments for an installment
func (uc *vehicleSalesUseCase) GetInstallmentPaymentsByInstallmentID(ctx context.Context, installmentID uint) ([]*models.InstallmentPayment, error) {
	return uc.repos.InstallmentPayment.GetByInstallmentID(ctx, installmentID)
}

// CalculateProfitForVehicle calculates total profit for a vehicle
func (uc *vehicleSalesUseCase) CalculateProfitForVehicle(ctx context.Context, vehicleID uint) (float64, error) {
	vehicle, err := uc.repos.Vehicle.GetByID(ctx, vehicleID)
	if err != nil {
		return 0, err
	}

	if vehicle.SellingPrice == nil || vehicle.PurchasePrice == nil {
		return 0, errors.New("vehicle must have both purchase and selling prices")
	}

	return uc.calculateProfitAmount(ctx, vehicle, *vehicle.SellingPrice), nil
}

// calculateProfitAmount calculates profit considering purchase price and reconditioning costs
func (uc *vehicleSalesUseCase) calculateProfitAmount(ctx context.Context, vehicle *models.Vehicle, salePrice float64) float64 {
	profit := salePrice
	
	// Subtract purchase price
	if vehicle.PurchasePrice != nil {
		profit -= *vehicle.PurchasePrice
	}

	// Subtract reconditioning costs
	jobs, err := uc.repos.VehicleReconditioningJob.GetByVehicleID(ctx, vehicle.VehicleID)
	if err == nil {
		for _, job := range jobs {
			if job.ActualCost != nil {
				profit -= *job.ActualCost
			}
		}
	}

	return profit
}