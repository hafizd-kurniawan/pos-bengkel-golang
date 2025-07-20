package implementations

import (
	"boilerplate/internal/models"
	"boilerplate/internal/repository/interfaces"
	"context"
	"time"

	"gorm.io/gorm"
)

// reconditioningDetailRepository implements interfaces.ReconditioningDetailRepository
type reconditioningDetailRepository struct {
	db *gorm.DB
}

// NewReconditioningDetailRepository creates a new reconditioning detail repository
func NewReconditioningDetailRepository(db *gorm.DB) interfaces.ReconditioningDetailRepository {
	return &reconditioningDetailRepository{db: db}
}

func (r *reconditioningDetailRepository) Create(ctx context.Context, detail *models.ReconditioningDetail) error {
	return r.db.WithContext(ctx).Create(detail).Error
}

func (r *reconditioningDetailRepository) GetByID(ctx context.Context, id uint) (*models.ReconditioningDetail, error) {
	var detail models.ReconditioningDetail
	err := r.db.WithContext(ctx).Preload("ReconditioningJob").Preload("Product").Preload("Service").First(&detail, id).Error
	if err != nil {
		return nil, err
	}
	return &detail, nil
}

func (r *reconditioningDetailRepository) Update(ctx context.Context, detail *models.ReconditioningDetail) error {
	return r.db.WithContext(ctx).Save(detail).Error
}

func (r *reconditioningDetailRepository) Delete(ctx context.Context, id uint) error {
	return r.db.WithContext(ctx).Delete(&models.ReconditioningDetail{}, id).Error
}

func (r *reconditioningDetailRepository) List(ctx context.Context, limit, offset int) ([]*models.ReconditioningDetail, error) {
	var details []*models.ReconditioningDetail
	err := r.db.WithContext(ctx).Preload("ReconditioningJob").Preload("Product").Preload("Service").Limit(limit).Offset(offset).Find(&details).Error
	return details, err
}

func (r *reconditioningDetailRepository) GetByReconditioningJobID(ctx context.Context, jobID uint) ([]*models.ReconditioningDetail, error) {
	var details []*models.ReconditioningDetail
	err := r.db.WithContext(ctx).Preload("ReconditioningJob").Preload("Product").Preload("Service").Where("reconditioning_job_id = ?", jobID).Find(&details).Error
	return details, err
}

func (r *reconditioningDetailRepository) GetByProductID(ctx context.Context, productID uint) ([]*models.ReconditioningDetail, error) {
	var details []*models.ReconditioningDetail
	err := r.db.WithContext(ctx).Preload("ReconditioningJob").Preload("Product").Where("product_id = ?", productID).Find(&details).Error
	return details, err
}

func (r *reconditioningDetailRepository) GetByServiceID(ctx context.Context, serviceID uint) ([]*models.ReconditioningDetail, error) {
	var details []*models.ReconditioningDetail
	err := r.db.WithContext(ctx).Preload("ReconditioningJob").Preload("Service").Where("service_id = ?", serviceID).Find(&details).Error
	return details, err
}

func (r *reconditioningDetailRepository) GetByType(ctx context.Context, detailType models.ReconditioningDetailType) ([]*models.ReconditioningDetail, error) {
	var details []*models.ReconditioningDetail
	err := r.db.WithContext(ctx).Preload("ReconditioningJob").Preload("Product").Preload("Service").Where("detail_type = ?", detailType).Find(&details).Error
	return details, err
}

// vehicleSalesTransactionRepository implements interfaces.VehicleSalesTransactionRepository
type vehicleSalesTransactionRepository struct {
	db *gorm.DB
}

// NewVehicleSalesTransactionRepository creates a new sales transaction repository
func NewVehicleSalesTransactionRepository(db *gorm.DB) interfaces.VehicleSalesTransactionRepository {
	return &vehicleSalesTransactionRepository{db: db}
}

func (r *vehicleSalesTransactionRepository) Create(ctx context.Context, transaction *models.VehicleSalesTransaction) error {
	return r.db.WithContext(ctx).Create(transaction).Error
}

func (r *vehicleSalesTransactionRepository) GetByID(ctx context.Context, id uint) (*models.VehicleSalesTransaction, error) {
	var transaction models.VehicleSalesTransaction
	err := r.db.WithContext(ctx).Preload("Vehicle").Preload("Customer").Preload("SalesPerson").Preload("Installments").First(&transaction, id).Error
	if err != nil {
		return nil, err
	}
	return &transaction, nil
}

func (r *vehicleSalesTransactionRepository) Update(ctx context.Context, transaction *models.VehicleSalesTransaction) error {
	return r.db.WithContext(ctx).Save(transaction).Error
}

func (r *vehicleSalesTransactionRepository) Delete(ctx context.Context, id uint) error {
	return r.db.WithContext(ctx).Delete(&models.VehicleSalesTransaction{}, id).Error
}

func (r *vehicleSalesTransactionRepository) List(ctx context.Context, limit, offset int) ([]*models.VehicleSalesTransaction, error) {
	var transactions []*models.VehicleSalesTransaction
	err := r.db.WithContext(ctx).Preload("Vehicle").Preload("Customer").Preload("SalesPerson").Limit(limit).Offset(offset).Find(&transactions).Error
	return transactions, err
}

func (r *vehicleSalesTransactionRepository) GetByVehicleID(ctx context.Context, vehicleID uint) ([]*models.VehicleSalesTransaction, error) {
	var transactions []*models.VehicleSalesTransaction
	err := r.db.WithContext(ctx).Preload("Vehicle").Preload("Customer").Preload("SalesPerson").Where("vehicle_id = ?", vehicleID).Find(&transactions).Error
	return transactions, err
}

func (r *vehicleSalesTransactionRepository) GetByCustomerID(ctx context.Context, customerID uint) ([]*models.VehicleSalesTransaction, error) {
	var transactions []*models.VehicleSalesTransaction
	err := r.db.WithContext(ctx).Preload("Vehicle").Preload("Customer").Preload("SalesPerson").Where("customer_id = ?", customerID).Find(&transactions).Error
	return transactions, err
}

func (r *vehicleSalesTransactionRepository) GetBySalesPersonID(ctx context.Context, salesPersonID uint) ([]*models.VehicleSalesTransaction, error) {
	var transactions []*models.VehicleSalesTransaction
	err := r.db.WithContext(ctx).Preload("Vehicle").Preload("Customer").Preload("SalesPerson").Where("sales_person_id = ?", salesPersonID).Find(&transactions).Error
	return transactions, err
}

func (r *vehicleSalesTransactionRepository) GetByDateRange(ctx context.Context, startDate, endDate string) ([]*models.VehicleSalesTransaction, error) {
	var transactions []*models.VehicleSalesTransaction
	err := r.db.WithContext(ctx).Preload("Vehicle").Preload("Customer").Preload("SalesPerson").Where("sale_date BETWEEN ? AND ?", startDate, endDate).Find(&transactions).Error
	return transactions, err
}

func (r *vehicleSalesTransactionRepository) GetByTransactionType(ctx context.Context, transactionType models.SalesTransactionType) ([]*models.VehicleSalesTransaction, error) {
	var transactions []*models.VehicleSalesTransaction
	err := r.db.WithContext(ctx).Preload("Vehicle").Preload("Customer").Preload("SalesPerson").Where("transaction_type = ?", transactionType).Find(&transactions).Error
	return transactions, err
}

func (r *vehicleSalesTransactionRepository) GetByStatus(ctx context.Context, status models.TransactionStatus) ([]*models.VehicleSalesTransaction, error) {
	var transactions []*models.VehicleSalesTransaction
	err := r.db.WithContext(ctx).Preload("Vehicle").Preload("Customer").Preload("SalesPerson").Where("transaction_status = ?", status).Find(&transactions).Error
	return transactions, err
}

// vehicleInstallmentRepository implements interfaces.VehicleInstallmentRepository
type vehicleInstallmentRepository struct {
	db *gorm.DB
}

// NewVehicleInstallmentRepository creates a new installment repository
func NewVehicleInstallmentRepository(db *gorm.DB) interfaces.VehicleInstallmentRepository {
	return &vehicleInstallmentRepository{db: db}
}

func (r *vehicleInstallmentRepository) Create(ctx context.Context, installment *models.VehicleInstallment) error {
	return r.db.WithContext(ctx).Create(installment).Error
}

func (r *vehicleInstallmentRepository) GetByID(ctx context.Context, id uint) (*models.VehicleInstallment, error) {
	var installment models.VehicleInstallment
	err := r.db.WithContext(ctx).Preload("SalesTransaction").Preload("InstallmentPayments").First(&installment, id).Error
	if err != nil {
		return nil, err
	}
	return &installment, nil
}

func (r *vehicleInstallmentRepository) Update(ctx context.Context, installment *models.VehicleInstallment) error {
	return r.db.WithContext(ctx).Save(installment).Error
}

func (r *vehicleInstallmentRepository) Delete(ctx context.Context, id uint) error {
	return r.db.WithContext(ctx).Delete(&models.VehicleInstallment{}, id).Error
}

func (r *vehicleInstallmentRepository) List(ctx context.Context, limit, offset int) ([]*models.VehicleInstallment, error) {
	var installments []*models.VehicleInstallment
	err := r.db.WithContext(ctx).Preload("SalesTransaction").Limit(limit).Offset(offset).Find(&installments).Error
	return installments, err
}

func (r *vehicleInstallmentRepository) GetBySalesTransactionID(ctx context.Context, salesTransactionID uint) (*models.VehicleInstallment, error) {
	var installment models.VehicleInstallment
	err := r.db.WithContext(ctx).Preload("SalesTransaction").Preload("InstallmentPayments").Where("sales_transaction_id = ?", salesTransactionID).First(&installment).Error
	if err != nil {
		return nil, err
	}
	return &installment, nil
}

func (r *vehicleInstallmentRepository) GetByStatus(ctx context.Context, status models.InstallmentStatus) ([]*models.VehicleInstallment, error) {
	var installments []*models.VehicleInstallment
	err := r.db.WithContext(ctx).Preload("SalesTransaction").Where("status = ?", status).Find(&installments).Error
	return installments, err
}

func (r *vehicleInstallmentRepository) GetActiveInstallments(ctx context.Context) ([]*models.VehicleInstallment, error) {
	var installments []*models.VehicleInstallment
	err := r.db.WithContext(ctx).Preload("SalesTransaction").Preload("InstallmentPayments").Where("status = ?", "Active").Find(&installments).Error
	return installments, err
}

// installmentPaymentRepository implements interfaces.InstallmentPaymentRepository
type installmentPaymentRepository struct {
	db *gorm.DB
}

// NewInstallmentPaymentRepository creates a new installment payment repository
func NewInstallmentPaymentRepository(db *gorm.DB) interfaces.InstallmentPaymentRepository {
	return &installmentPaymentRepository{db: db}
}

func (r *installmentPaymentRepository) Create(ctx context.Context, payment *models.InstallmentPayment) error {
	return r.db.WithContext(ctx).Create(payment).Error
}

func (r *installmentPaymentRepository) GetByID(ctx context.Context, id uint) (*models.InstallmentPayment, error) {
	var payment models.InstallmentPayment
	err := r.db.WithContext(ctx).Preload("Installment").First(&payment, id).Error
	if err != nil {
		return nil, err
	}
	return &payment, nil
}

func (r *installmentPaymentRepository) Update(ctx context.Context, payment *models.InstallmentPayment) error {
	return r.db.WithContext(ctx).Save(payment).Error
}

func (r *installmentPaymentRepository) Delete(ctx context.Context, id uint) error {
	return r.db.WithContext(ctx).Delete(&models.InstallmentPayment{}, id).Error
}

func (r *installmentPaymentRepository) List(ctx context.Context, limit, offset int) ([]*models.InstallmentPayment, error) {
	var payments []*models.InstallmentPayment
	err := r.db.WithContext(ctx).Preload("Installment").Limit(limit).Offset(offset).Find(&payments).Error
	return payments, err
}

func (r *installmentPaymentRepository) GetByInstallmentID(ctx context.Context, installmentID uint) ([]*models.InstallmentPayment, error) {
	var payments []*models.InstallmentPayment
	err := r.db.WithContext(ctx).Preload("Installment").Where("installment_id = ?", installmentID).Order("payment_number ASC").Find(&payments).Error
	return payments, err
}

func (r *installmentPaymentRepository) GetByPaymentNumber(ctx context.Context, installmentID uint, paymentNumber int) (*models.InstallmentPayment, error) {
	var payment models.InstallmentPayment
	err := r.db.WithContext(ctx).Preload("Installment").Where("installment_id = ? AND payment_number = ?", installmentID, paymentNumber).First(&payment).Error
	if err != nil {
		return nil, err
	}
	return &payment, nil
}

func (r *installmentPaymentRepository) GetByStatus(ctx context.Context, status models.InstallmentPaymentStatus) ([]*models.InstallmentPayment, error) {
	var payments []*models.InstallmentPayment
	err := r.db.WithContext(ctx).Preload("Installment").Where("payment_status = ?", status).Find(&payments).Error
	return payments, err
}

func (r *installmentPaymentRepository) GetOverduePayments(ctx context.Context) ([]*models.InstallmentPayment, error) {
	var payments []*models.InstallmentPayment
	now := time.Now()
	err := r.db.WithContext(ctx).Preload("Installment").Where("due_date < ? AND payment_status IN (?)", now, []string{"Pending", "Late"}).Find(&payments).Error
	return payments, err
}

func (r *installmentPaymentRepository) GetPaymentsDueInRange(ctx context.Context, startDate, endDate string) ([]*models.InstallmentPayment, error) {
	var payments []*models.InstallmentPayment
	err := r.db.WithContext(ctx).Preload("Installment").Where("due_date BETWEEN ? AND ?", startDate, endDate).Find(&payments).Error
	return payments, err
}