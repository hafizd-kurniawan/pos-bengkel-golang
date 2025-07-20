package implementations

import (
	"boilerplate/internal/models"
	"boilerplate/internal/repository/interfaces"
	"context"

	"gorm.io/gorm"
)

// vehicleRepository implements interfaces.VehicleRepository
type vehicleRepository struct {
	db *gorm.DB
}

// NewVehicleRepository creates a new vehicle repository
func NewVehicleRepository(db *gorm.DB) interfaces.VehicleRepository {
	return &vehicleRepository{db: db}
}

func (r *vehicleRepository) Create(ctx context.Context, vehicle *models.Vehicle) error {
	return r.db.WithContext(ctx).Create(vehicle).Error
}

func (r *vehicleRepository) GetByID(ctx context.Context, id uint) (*models.Vehicle, error) {
	var vehicle models.Vehicle
	err := r.db.WithContext(ctx).Preload("Customer").Preload("PurchaseTransactions").Preload("ReconditioningJobs").Preload("SalesTransactions").First(&vehicle, id).Error
	if err != nil {
		return nil, err
	}
	return &vehicle, nil
}

func (r *vehicleRepository) GetByPlateNumber(ctx context.Context, plateNumber string) (*models.Vehicle, error) {
	var vehicle models.Vehicle
	err := r.db.WithContext(ctx).Preload("Customer").Where("plate_number = ?", plateNumber).First(&vehicle).Error
	if err != nil {
		return nil, err
	}
	return &vehicle, nil
}

func (r *vehicleRepository) GetByChassisNumber(ctx context.Context, chassisNumber string) (*models.Vehicle, error) {
	var vehicle models.Vehicle
	err := r.db.WithContext(ctx).Preload("Customer").Where("chassis_number = ?", chassisNumber).First(&vehicle).Error
	if err != nil {
		return nil, err
	}
	return &vehicle, nil
}

func (r *vehicleRepository) GetByEngineNumber(ctx context.Context, engineNumber string) (*models.Vehicle, error) {
	var vehicle models.Vehicle
	err := r.db.WithContext(ctx).Preload("Customer").Where("engine_number = ?", engineNumber).First(&vehicle).Error
	if err != nil {
		return nil, err
	}
	return &vehicle, nil
}

func (r *vehicleRepository) Update(ctx context.Context, vehicle *models.Vehicle) error {
	return r.db.WithContext(ctx).Save(vehicle).Error
}

func (r *vehicleRepository) Delete(ctx context.Context, id uint) error {
	return r.db.WithContext(ctx).Delete(&models.Vehicle{}, id).Error
}

func (r *vehicleRepository) List(ctx context.Context, limit, offset int) ([]*models.Vehicle, error) {
	var vehicles []*models.Vehicle
	err := r.db.WithContext(ctx).Preload("Customer").Limit(limit).Offset(offset).Find(&vehicles).Error
	return vehicles, err
}

func (r *vehicleRepository) GetByCustomerID(ctx context.Context, customerID uint) ([]*models.Vehicle, error) {
	var vehicles []*models.Vehicle
	err := r.db.WithContext(ctx).Preload("Customer").Where("customer_id = ?", customerID).Find(&vehicles).Error
	return vehicles, err
}

func (r *vehicleRepository) GetByOwnershipStatus(ctx context.Context, status models.VehicleOwnershipStatus) ([]*models.Vehicle, error) {
	var vehicles []*models.Vehicle
	err := r.db.WithContext(ctx).Preload("Customer").Where("ownership_status = ?", status).Find(&vehicles).Error
	return vehicles, err
}

func (r *vehicleRepository) GetBySaleStatus(ctx context.Context, status models.VehicleSaleStatus) ([]*models.Vehicle, error) {
	var vehicles []*models.Vehicle
	err := r.db.WithContext(ctx).Preload("Customer").Where("sale_status = ?", status).Find(&vehicles).Error
	return vehicles, err
}

func (r *vehicleRepository) Search(ctx context.Context, query string, limit, offset int) ([]*models.Vehicle, error) {
	var vehicles []*models.Vehicle
	searchQuery := "%" + query + "%"
	err := r.db.WithContext(ctx).Preload("Customer").Where(
		"plate_number ILIKE ? OR brand ILIKE ? OR model ILIKE ? OR chassis_number ILIKE ? OR engine_number ILIKE ?",
		searchQuery, searchQuery, searchQuery, searchQuery, searchQuery,
	).Limit(limit).Offset(offset).Find(&vehicles).Error
	return vehicles, err
}

// vehiclePurchaseTransactionRepository implements interfaces.VehiclePurchaseTransactionRepository
type vehiclePurchaseTransactionRepository struct {
	db *gorm.DB
}

// NewVehiclePurchaseTransactionRepository creates a new purchase transaction repository
func NewVehiclePurchaseTransactionRepository(db *gorm.DB) interfaces.VehiclePurchaseTransactionRepository {
	return &vehiclePurchaseTransactionRepository{db: db}
}

func (r *vehiclePurchaseTransactionRepository) Create(ctx context.Context, transaction *models.VehiclePurchaseTransaction) error {
	return r.db.WithContext(ctx).Create(transaction).Error
}

func (r *vehiclePurchaseTransactionRepository) GetByID(ctx context.Context, id uint) (*models.VehiclePurchaseTransaction, error) {
	var transaction models.VehiclePurchaseTransaction
	err := r.db.WithContext(ctx).Preload("Vehicle").Preload("Customer").First(&transaction, id).Error
	if err != nil {
		return nil, err
	}
	return &transaction, nil
}

func (r *vehiclePurchaseTransactionRepository) Update(ctx context.Context, transaction *models.VehiclePurchaseTransaction) error {
	return r.db.WithContext(ctx).Save(transaction).Error
}

func (r *vehiclePurchaseTransactionRepository) Delete(ctx context.Context, id uint) error {
	return r.db.WithContext(ctx).Delete(&models.VehiclePurchaseTransaction{}, id).Error
}

func (r *vehiclePurchaseTransactionRepository) List(ctx context.Context, limit, offset int) ([]*models.VehiclePurchaseTransaction, error) {
	var transactions []*models.VehiclePurchaseTransaction
	err := r.db.WithContext(ctx).Preload("Vehicle").Preload("Customer").Limit(limit).Offset(offset).Find(&transactions).Error
	return transactions, err
}

func (r *vehiclePurchaseTransactionRepository) GetByVehicleID(ctx context.Context, vehicleID uint) ([]*models.VehiclePurchaseTransaction, error) {
	var transactions []*models.VehiclePurchaseTransaction
	err := r.db.WithContext(ctx).Preload("Vehicle").Preload("Customer").Where("vehicle_id = ?", vehicleID).Find(&transactions).Error
	return transactions, err
}

func (r *vehiclePurchaseTransactionRepository) GetByCustomerID(ctx context.Context, customerID uint) ([]*models.VehiclePurchaseTransaction, error) {
	var transactions []*models.VehiclePurchaseTransaction
	err := r.db.WithContext(ctx).Preload("Vehicle").Preload("Customer").Where("customer_id = ?", customerID).Find(&transactions).Error
	return transactions, err
}

func (r *vehiclePurchaseTransactionRepository) GetByDateRange(ctx context.Context, startDate, endDate string) ([]*models.VehiclePurchaseTransaction, error) {
	var transactions []*models.VehiclePurchaseTransaction
	err := r.db.WithContext(ctx).Preload("Vehicle").Preload("Customer").Where("purchase_date BETWEEN ? AND ?", startDate, endDate).Find(&transactions).Error
	return transactions, err
}

func (r *vehiclePurchaseTransactionRepository) GetByStatus(ctx context.Context, status models.TransactionStatus) ([]*models.VehiclePurchaseTransaction, error) {
	var transactions []*models.VehiclePurchaseTransaction
	err := r.db.WithContext(ctx).Preload("Vehicle").Preload("Customer").Where("transaction_status = ?", status).Find(&transactions).Error
	return transactions, err
}

// vehicleReconditioningJobRepository implements interfaces.VehicleReconditioningJobRepository
type vehicleReconditioningJobRepository struct {
	db *gorm.DB
}

// NewVehicleReconditioningJobRepository creates a new reconditioning job repository
func NewVehicleReconditioningJobRepository(db *gorm.DB) interfaces.VehicleReconditioningJobRepository {
	return &vehicleReconditioningJobRepository{db: db}
}

func (r *vehicleReconditioningJobRepository) Create(ctx context.Context, job *models.VehicleReconditioningJob) error {
	return r.db.WithContext(ctx).Create(job).Error
}

func (r *vehicleReconditioningJobRepository) GetByID(ctx context.Context, id uint) (*models.VehicleReconditioningJob, error) {
	var job models.VehicleReconditioningJob
	err := r.db.WithContext(ctx).Preload("Vehicle").Preload("AssignedTechnician").Preload("ReconditioningDetails").First(&job, id).Error
	if err != nil {
		return nil, err
	}
	return &job, nil
}

func (r *vehicleReconditioningJobRepository) Update(ctx context.Context, job *models.VehicleReconditioningJob) error {
	return r.db.WithContext(ctx).Save(job).Error
}

func (r *vehicleReconditioningJobRepository) Delete(ctx context.Context, id uint) error {
	return r.db.WithContext(ctx).Delete(&models.VehicleReconditioningJob{}, id).Error
}

func (r *vehicleReconditioningJobRepository) List(ctx context.Context, limit, offset int) ([]*models.VehicleReconditioningJob, error) {
	var jobs []*models.VehicleReconditioningJob
	err := r.db.WithContext(ctx).Preload("Vehicle").Preload("AssignedTechnician").Limit(limit).Offset(offset).Find(&jobs).Error
	return jobs, err
}

func (r *vehicleReconditioningJobRepository) GetByVehicleID(ctx context.Context, vehicleID uint) ([]*models.VehicleReconditioningJob, error) {
	var jobs []*models.VehicleReconditioningJob
	err := r.db.WithContext(ctx).Preload("Vehicle").Preload("AssignedTechnician").Preload("ReconditioningDetails").Where("vehicle_id = ?", vehicleID).Find(&jobs).Error
	return jobs, err
}

func (r *vehicleReconditioningJobRepository) GetByTechnicianID(ctx context.Context, technicianID uint) ([]*models.VehicleReconditioningJob, error) {
	var jobs []*models.VehicleReconditioningJob
	err := r.db.WithContext(ctx).Preload("Vehicle").Preload("AssignedTechnician").Where("assigned_technician_id = ?", technicianID).Find(&jobs).Error
	return jobs, err
}

func (r *vehicleReconditioningJobRepository) GetByStatus(ctx context.Context, status models.ReconditioningJobStatus) ([]*models.VehicleReconditioningJob, error) {
	var jobs []*models.VehicleReconditioningJob
	err := r.db.WithContext(ctx).Preload("Vehicle").Preload("AssignedTechnician").Where("status = ?", status).Find(&jobs).Error
	return jobs, err
}

func (r *vehicleReconditioningJobRepository) GetActiveJobsForVehicle(ctx context.Context, vehicleID uint) ([]*models.VehicleReconditioningJob, error) {
	var jobs []*models.VehicleReconditioningJob
	err := r.db.WithContext(ctx).Preload("Vehicle").Preload("AssignedTechnician").Preload("ReconditioningDetails").Where("vehicle_id = ? AND status IN (?)", vehicleID, []string{"Pending", "In Progress"}).Find(&jobs).Error
	return jobs, err
}