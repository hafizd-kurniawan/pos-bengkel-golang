package models

// This file serves as an index for all models in the POS Bengkel system
// Import this package to get access to all model definitions

// Re-export all models for easier importing
type (
	// Foundation & Security
	UserModel            = User
	OutletModel          = Outlet
	RoleModel            = Role
	PermissionModel      = Permission
	RoleHasPermissionModel = RoleHasPermission

	// Customer & Vehicle
	CustomerModel        = Customer
	CustomerVehicleModel = CustomerVehicle

	// Master Data & Inventory
	ProductModel             = Product
	ProductSerialNumberModel = ProductSerialNumber
	CategoryModel            = Category
	SupplierModel            = Supplier
	UnitTypeModel            = UnitType

	// Services
	ServiceModel         = Service
	ServiceCategoryModel = ServiceCategory

	// Core Operations
	ServiceJobModel        = ServiceJob
	ServiceDetailModel     = ServiceDetail
	ServiceJobHistoryModel = ServiceJobHistory

	// Transactions
	TransactionModel        = Transaction
	TransactionDetailModel  = TransactionDetail
	PurchaseOrderModel      = PurchaseOrder
	PurchaseOrderDetailModel = PurchaseOrderDetail
	VehiclePurchaseModel    = VehiclePurchase

	// Financial
	PaymentMethodModel      = PaymentMethod
	PaymentModel            = Payment
	AccountsPayableModel    = AccountsPayable
	PayablePaymentModel     = PayablePayment
	AccountsReceivableModel = AccountsReceivable
	ReceivablePaymentModel  = ReceivablePayment
	CashFlowModel           = CashFlow

	// Reporting & Promotions
	ReportModel    = Report
	PromotionModel = Promotion
)

// GetAllModels returns a slice of all model types for migration purposes
func GetAllModels() []interface{} {
	return []interface{}{
		// Foundation & Security
		&User{},
		&Outlet{},
		&Role{},
		&Permission{},
		&RoleHasPermission{},

		// Customer & Vehicle
		&Customer{},
		&CustomerVehicle{},

		// Master Data & Inventory
		&Product{},
		&ProductSerialNumber{},
		&Category{},
		&Supplier{},
		&UnitType{},

		// Services
		&Service{},
		&ServiceCategory{},

		// Core Operations
		&ServiceJob{},
		&ServiceDetail{},
		&ServiceJobHistory{},

		// Transactions
		&Transaction{},
		&TransactionDetail{},
		&PurchaseOrder{},
		&PurchaseOrderDetail{},
		&VehiclePurchase{},

		// Financial
		&PaymentMethod{},
		&Payment{},
		&AccountsPayable{},
		&PayablePayment{},
		&AccountsReceivable{},
		&ReceivablePayment{},
		&CashFlow{},

		// Reporting & Promotions
		&Report{},
		&Promotion{},
	}
}