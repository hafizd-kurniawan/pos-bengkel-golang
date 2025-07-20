package interfaces

import (
	"boilerplate/internal/models"
	"context"
)

// VehicleRepository interface for vehicle operations
type VehicleRepository interface {
	Create(ctx context.Context, vehicle *models.Vehicle) error
	GetByID(ctx context.Context, id uint) (*models.Vehicle, error)
	GetByPlateNumber(ctx context.Context, plateNumber string) (*models.Vehicle, error)
	GetByChassisNumber(ctx context.Context, chassisNumber string) (*models.Vehicle, error)
	GetByEngineNumber(ctx context.Context, engineNumber string) (*models.Vehicle, error)
	Update(ctx context.Context, vehicle *models.Vehicle) error
	Delete(ctx context.Context, id uint) error
	List(ctx context.Context, limit, offset int) ([]*models.Vehicle, error)
	GetByCustomerID(ctx context.Context, customerID uint) ([]*models.Vehicle, error)
	GetByOwnershipStatus(ctx context.Context, status models.VehicleOwnershipStatus) ([]*models.Vehicle, error)
	GetBySaleStatus(ctx context.Context, status models.VehicleSaleStatus) ([]*models.Vehicle, error)
	Search(ctx context.Context, query string, limit, offset int) ([]*models.Vehicle, error)
}

// VehiclePurchaseTransactionRepository interface for purchase transaction operations
type VehiclePurchaseTransactionRepository interface {
	Create(ctx context.Context, transaction *models.VehiclePurchaseTransaction) error
	GetByID(ctx context.Context, id uint) (*models.VehiclePurchaseTransaction, error)
	Update(ctx context.Context, transaction *models.VehiclePurchaseTransaction) error
	Delete(ctx context.Context, id uint) error
	List(ctx context.Context, limit, offset int) ([]*models.VehiclePurchaseTransaction, error)
	GetByVehicleID(ctx context.Context, vehicleID uint) ([]*models.VehiclePurchaseTransaction, error)
	GetByCustomerID(ctx context.Context, customerID uint) ([]*models.VehiclePurchaseTransaction, error)
	GetByDateRange(ctx context.Context, startDate, endDate string) ([]*models.VehiclePurchaseTransaction, error)
	GetByStatus(ctx context.Context, status models.TransactionStatus) ([]*models.VehiclePurchaseTransaction, error)
}

// VehicleReconditioningJobRepository interface for reconditioning job operations
type VehicleReconditioningJobRepository interface {
	Create(ctx context.Context, job *models.VehicleReconditioningJob) error
	GetByID(ctx context.Context, id uint) (*models.VehicleReconditioningJob, error)
	Update(ctx context.Context, job *models.VehicleReconditioningJob) error
	Delete(ctx context.Context, id uint) error
	List(ctx context.Context, limit, offset int) ([]*models.VehicleReconditioningJob, error)
	GetByVehicleID(ctx context.Context, vehicleID uint) ([]*models.VehicleReconditioningJob, error)
	GetByTechnicianID(ctx context.Context, technicianID uint) ([]*models.VehicleReconditioningJob, error)
	GetByStatus(ctx context.Context, status models.ReconditioningJobStatus) ([]*models.VehicleReconditioningJob, error)
	GetActiveJobsForVehicle(ctx context.Context, vehicleID uint) ([]*models.VehicleReconditioningJob, error)
}

// ReconditioningDetailRepository interface for reconditioning detail operations
type ReconditioningDetailRepository interface {
	Create(ctx context.Context, detail *models.ReconditioningDetail) error
	GetByID(ctx context.Context, id uint) (*models.ReconditioningDetail, error)
	Update(ctx context.Context, detail *models.ReconditioningDetail) error
	Delete(ctx context.Context, id uint) error
	List(ctx context.Context, limit, offset int) ([]*models.ReconditioningDetail, error)
	GetByReconditioningJobID(ctx context.Context, jobID uint) ([]*models.ReconditioningDetail, error)
	GetByProductID(ctx context.Context, productID uint) ([]*models.ReconditioningDetail, error)
	GetByServiceID(ctx context.Context, serviceID uint) ([]*models.ReconditioningDetail, error)
	GetByType(ctx context.Context, detailType models.ReconditioningDetailType) ([]*models.ReconditioningDetail, error)
}

// VehicleSalesTransactionRepository interface for sales transaction operations
type VehicleSalesTransactionRepository interface {
	Create(ctx context.Context, transaction *models.VehicleSalesTransaction) error
	GetByID(ctx context.Context, id uint) (*models.VehicleSalesTransaction, error)
	Update(ctx context.Context, transaction *models.VehicleSalesTransaction) error
	Delete(ctx context.Context, id uint) error
	List(ctx context.Context, limit, offset int) ([]*models.VehicleSalesTransaction, error)
	GetByVehicleID(ctx context.Context, vehicleID uint) ([]*models.VehicleSalesTransaction, error)
	GetByCustomerID(ctx context.Context, customerID uint) ([]*models.VehicleSalesTransaction, error)
	GetBySalesPersonID(ctx context.Context, salesPersonID uint) ([]*models.VehicleSalesTransaction, error)
	GetByDateRange(ctx context.Context, startDate, endDate string) ([]*models.VehicleSalesTransaction, error)
	GetByTransactionType(ctx context.Context, transactionType models.SalesTransactionType) ([]*models.VehicleSalesTransaction, error)
	GetByStatus(ctx context.Context, status models.TransactionStatus) ([]*models.VehicleSalesTransaction, error)
}

// VehicleInstallmentRepository interface for installment operations
type VehicleInstallmentRepository interface {
	Create(ctx context.Context, installment *models.VehicleInstallment) error
	GetByID(ctx context.Context, id uint) (*models.VehicleInstallment, error)
	Update(ctx context.Context, installment *models.VehicleInstallment) error
	Delete(ctx context.Context, id uint) error
	List(ctx context.Context, limit, offset int) ([]*models.VehicleInstallment, error)
	GetBySalesTransactionID(ctx context.Context, salesTransactionID uint) (*models.VehicleInstallment, error)
	GetByStatus(ctx context.Context, status models.InstallmentStatus) ([]*models.VehicleInstallment, error)
	GetActiveInstallments(ctx context.Context) ([]*models.VehicleInstallment, error)
}

// InstallmentPaymentRepository interface for installment payment operations
type InstallmentPaymentRepository interface {
	Create(ctx context.Context, payment *models.InstallmentPayment) error
	GetByID(ctx context.Context, id uint) (*models.InstallmentPayment, error)
	Update(ctx context.Context, payment *models.InstallmentPayment) error
	Delete(ctx context.Context, id uint) error
	List(ctx context.Context, limit, offset int) ([]*models.InstallmentPayment, error)
	GetByInstallmentID(ctx context.Context, installmentID uint) ([]*models.InstallmentPayment, error)
	GetByPaymentNumber(ctx context.Context, installmentID uint, paymentNumber int) (*models.InstallmentPayment, error)
	GetByStatus(ctx context.Context, status models.InstallmentPaymentStatus) ([]*models.InstallmentPayment, error)
	GetOverduePayments(ctx context.Context) ([]*models.InstallmentPayment, error)
	GetPaymentsDueInRange(ctx context.Context, startDate, endDate string) ([]*models.InstallmentPayment, error)
}