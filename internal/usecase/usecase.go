package usecase

import (
	"boilerplate/internal/repository"
	"boilerplate/internal/usecase/implementations"
	"boilerplate/internal/usecase/interfaces"
)

// UsecaseManager contains all usecase interfaces
type UsecaseManager struct {
	// Foundation & Security
	User       interfaces.UserUsecase
	Outlet     interfaces.OutletUsecase
	Role       interfaces.RoleUsecase
	Permission interfaces.PermissionUsecase

	// Customer & Vehicle
	Customer        interfaces.CustomerUsecase
	CustomerVehicle interfaces.CustomerVehicleUsecase

	// Master Data & Inventory
	Product             interfaces.ProductUsecase
	ProductSerialNumber interfaces.ProductSerialNumberUsecase
	Category            interfaces.CategoryUsecase
	Supplier            interfaces.SupplierUsecase
	UnitType            interfaces.UnitTypeUsecase

	// Services
	Service           interfaces.ServiceUsecase
	ServiceCategory   interfaces.ServiceCategoryUsecase
	ServiceJob        interfaces.ServiceJobUsecase
	ServiceDetail     interfaces.ServiceDetailUsecase
	ServiceJobHistory interfaces.ServiceJobHistoryUsecase

	// Transactions
	Transaction       interfaces.TransactionUsecase
	TransactionDetail interfaces.TransactionDetailUsecase

	// Financial
	PaymentMethod interfaces.PaymentMethodUsecase
	Payment       interfaces.PaymentUsecase
	CashFlow      interfaces.CashFlowUsecase

	// Add other usecases as they are implemented
}

// NewUsecaseManager creates a new usecase manager with all usecases
func NewUsecaseManager(repo *repository.RepositoryManager) *UsecaseManager {
	return &UsecaseManager{
		// Foundation & Security
		User:   implementations.NewUserUsecase(repo),
		Outlet: implementations.NewOutletUsecase(repo),
		Role:       implementations.NewRoleUsecase(repo),
		Permission: implementations.NewPermissionUsecase(repo),

		// Customer & Vehicle
		Customer:        implementations.NewCustomerUsecase(repo),
		CustomerVehicle: implementations.NewCustomerVehicleUsecase(repo),

		// Master Data & Inventory
		Product:             implementations.NewProductUsecase(repo),
		ProductSerialNumber: implementations.NewProductSerialNumberUsecase(repo),
		Category:            implementations.NewCategoryUsecase(repo),
		Supplier:            implementations.NewSupplierUsecase(repo),
		UnitType:            implementations.NewUnitTypeUsecase(repo),

		// Services
		Service:           implementations.NewServiceUsecase(repo),
		ServiceCategory:   implementations.NewServiceCategoryUsecase(repo),
		ServiceJob:        implementations.NewServiceJobUsecase(repo),
		ServiceDetail:     implementations.NewServiceDetailUsecase(repo),
		ServiceJobHistory: implementations.NewServiceJobHistoryUsecase(repo),

		// Transactions
		Transaction:       implementations.NewTransactionUsecase(repo),
		TransactionDetail: implementations.NewTransactionDetailUsecase(repo),

		// Financial
		PaymentMethod: implementations.NewPaymentMethodUsecase(repo),
		Payment:       implementations.NewPaymentUsecase(repo),
		CashFlow:      implementations.NewCashFlowUsecase(repo),

		// Add other usecases as they are implemented
	}
}