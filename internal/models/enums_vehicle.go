package models

// Vehicle ownership status
type VehicleOwnershipStatus string

const (
	VehicleOwnershipCustomer   VehicleOwnershipStatus = "Customer"
	VehicleOwnershipShowroom   VehicleOwnershipStatus = "Showroom"
	VehicleOwnershipWorkshop   VehicleOwnershipStatus = "Workshop"
)

// Vehicle condition status
type VehicleConditionStatus string

const (
	VehicleConditionExcellent VehicleConditionStatus = "Excellent"
	VehicleConditionGood      VehicleConditionStatus = "Good"
	VehicleConditionFair      VehicleConditionStatus = "Fair"
	VehicleConditionPoor      VehicleConditionStatus = "Poor"
)

// Vehicle sale status
type VehicleSaleStatus string

const (
	VehicleSaleStatusNotForSale VehicleSaleStatus = "Not For Sale"
	VehicleSaleStatusForSale    VehicleSaleStatus = "For Sale"
	VehicleSaleStatusSold       VehicleSaleStatus = "Sold"
	VehicleSaleStatusReserved   VehicleSaleStatus = "Reserved"
)

// Reconditioning job status
type ReconditioningJobStatus string

const (
	ReconditioningJobStatusPending    ReconditioningJobStatus = "Pending"
	ReconditioningJobStatusInProgress ReconditioningJobStatus = "In Progress"
	ReconditioningJobStatusCompleted  ReconditioningJobStatus = "Completed"
	ReconditioningJobStatusCancelled  ReconditioningJobStatus = "Cancelled"
)

// Reconditioning detail type
type ReconditioningDetailType string

const (
	ReconditioningDetailTypePart    ReconditioningDetailType = "Part"
	ReconditioningDetailTypeService ReconditioningDetailType = "Service"
)

// Sales transaction type
type SalesTransactionType string

const (
	SalesTransactionTypeCash        SalesTransactionType = "Cash"
	SalesTransactionTypeInstallment SalesTransactionType = "Installment"
)

// Installment status
type InstallmentStatus string

const (
	InstallmentStatusActive    InstallmentStatus = "Active"
	InstallmentStatusCompleted InstallmentStatus = "Completed"
	InstallmentStatusDefaulted InstallmentStatus = "Defaulted"
)

// Installment payment status
type InstallmentPaymentStatus string

const (
	InstallmentPaymentStatusPending InstallmentPaymentStatus = "Pending"
	InstallmentPaymentStatusPaid    InstallmentPaymentStatus = "Paid"
	InstallmentPaymentStatusLate    InstallmentPaymentStatus = "Late"
	InstallmentPaymentStatusSkipped InstallmentPaymentStatus = "Skipped"
)