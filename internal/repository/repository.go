package repository

import (
	"boilerplate/internal/repository/implementations"
	"boilerplate/internal/repository/interfaces"

	"github.com/jmoiron/sqlx"
	"gorm.io/gorm"
)

// RepositoryManager contains all repository interfaces
type RepositoryManager struct {
	// Foundation & Security
	User       interfaces.UserRepository
	Outlet     interfaces.OutletRepository
	Role       interfaces.RoleRepository
	Permission interfaces.PermissionRepository

	// Customer & Vehicle
	Customer        interfaces.CustomerRepository
	CustomerVehicle interfaces.CustomerVehicleRepository

	// Master Data & Inventory
	Product             interfaces.ProductRepository
	ProductSerialNumber interfaces.ProductSerialNumberRepository
	Category            interfaces.CategoryRepository
	Supplier            interfaces.SupplierRepository
	UnitType            interfaces.UnitTypeRepository

	// Services
	Service         interfaces.ServiceRepository
	ServiceCategory interfaces.ServiceCategoryRepository
	ServiceJob      interfaces.ServiceJobRepository
	ServiceDetail   interfaces.ServiceDetailRepository
	ServiceJobHistory interfaces.ServiceJobHistoryRepository

	// Transactions
	Transaction           interfaces.TransactionRepository
	TransactionDetail     interfaces.TransactionDetailRepository
	PurchaseOrder         interfaces.PurchaseOrderRepository
	PurchaseOrderDetail   interfaces.PurchaseOrderDetailRepository
	VehiclePurchase       interfaces.VehiclePurchaseRepository

	// Financial
	PaymentMethod       interfaces.PaymentMethodRepository
	Payment             interfaces.PaymentRepository
	AccountsPayable     interfaces.AccountsPayableRepository
	PayablePayment      interfaces.PayablePaymentRepository
	AccountsReceivable  interfaces.AccountsReceivableRepository
	ReceivablePayment   interfaces.ReceivablePaymentRepository
	CashFlow            interfaces.CashFlowRepository

	// Reporting & Promotions
	Report    interfaces.ReportRepository
	Promotion interfaces.PromotionRepository
}

// NewRepositoryManager creates a new repository manager with all repositories
func NewRepositoryManager(db *gorm.DB) *RepositoryManager {
	return &RepositoryManager{
		// Foundation & Security
		User:       implementations.NewUserRepository(db),
		Outlet:     implementations.NewOutletRepository(db),
		Role:       implementations.NewRoleRepository(db),
		Permission: implementations.NewPermissionRepository(db),

		// Customer & Vehicle
		Customer:        implementations.NewCustomerRepositoryGORM(db),
		CustomerVehicle: implementations.NewCustomerVehicleRepositoryGORM(db),

		// Master Data & Inventory
		Product:             implementations.NewProductRepository(db),
		ProductSerialNumber: implementations.NewProductSerialNumberRepository(db),
		Category:            implementations.NewCategoryRepository(db),
		Supplier:            implementations.NewSupplierRepository(db),
		UnitType:            implementations.NewUnitTypeRepository(db),

		// Services
		Service:           implementations.NewServiceRepository(db),
		ServiceCategory:   implementations.NewServiceCategoryRepository(db),
		ServiceJob:        implementations.NewServiceJobRepository(db),
		ServiceDetail:     implementations.NewServiceDetailRepository(db),
		ServiceJobHistory: implementations.NewServiceJobHistoryRepository(db),

		// Transactions
		Transaction:           implementations.NewTransactionRepository(db),
		TransactionDetail:     implementations.NewTransactionDetailRepository(db),

		// Financial
		PaymentMethod:       implementations.NewPaymentMethodRepository(db),
		Payment:             implementations.NewPaymentRepository(db),
		CashFlow:            implementations.NewCashFlowRepository(db),

		// Add other repositories as they are implemented
	}
}

// NewSQLXRepositoryManager creates a new repository manager with SQLx repositories
func NewSQLXRepositoryManager(db *sqlx.DB, gormDB *gorm.DB) *RepositoryManager {
	return &RepositoryManager{
		// Foundation & Security (converted to SQLx)
		User:       implementations.NewUserRepositorySQLX(db),
		Outlet:     implementations.NewOutletRepositorySQLX(db),
		Role:       implementations.NewRoleRepositorySQLX(db),
		Permission: implementations.NewPermissionRepositorySQLX(db),

		// Customer & Vehicle (converted to SQLx)
		Customer:        implementations.NewCustomerRepository(db),
		CustomerVehicle: implementations.NewCustomerVehicleRepository(db),

		// Master Data & Inventory (converted to SQLx)
		Product:             implementations.NewProductRepositorySQLX(db),
		ProductSerialNumber: implementations.NewProductSerialNumberRepository(gormDB),
		Category:            implementations.NewCategoryRepositorySQLX(db),
		Supplier:            implementations.NewSupplierRepository(gormDB),
		UnitType:            implementations.NewUnitTypeRepository(gormDB),

		// Services (using GORM for now)
		Service:           implementations.NewServiceRepository(gormDB),
		ServiceCategory:   implementations.NewServiceCategoryRepository(gormDB),
		ServiceJob:        implementations.NewServiceJobRepository(gormDB),
		ServiceDetail:     implementations.NewServiceDetailRepository(gormDB),
		ServiceJobHistory: implementations.NewServiceJobHistoryRepository(gormDB),

		// Transactions (using GORM for now)
		Transaction:           implementations.NewTransactionRepository(gormDB),
		TransactionDetail:     implementations.NewTransactionDetailRepository(gormDB),

		// Financial (using GORM for now)
		PaymentMethod:       implementations.NewPaymentMethodRepository(gormDB),
		Payment:             implementations.NewPaymentRepository(gormDB),
		CashFlow:            implementations.NewCashFlowRepository(gormDB),
	}
}