package responses

import (
	"boilerplate/internal/models"
	"time"
)

// Response represents a standard API response
type Response struct {
	Status  string      `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
	Error   string      `json:"error,omitempty"`
}

// PaginatedResponse represents a paginated API response
type PaginatedResponse struct {
	Status     string      `json:"status"`
	Message    string      `json:"message"`
	Data       interface{} `json:"data"`
	Pagination Pagination  `json:"pagination"`
}

// Pagination contains pagination metadata
type Pagination struct {
	Page    int   `json:"page"`
	Limit   int   `json:"limit"`
	Total   int64 `json:"total"`
	Pages   int   `json:"pages"`
}

// UserResponse represents user data in API response
type UserResponse struct {
	UserID    uint          `json:"user_id"`
	Name      string        `json:"name"`
	Email     string        `json:"email"`
	OutletID  *uint         `json:"outlet_id"`
	Outlet    *OutletResponse `json:"outlet,omitempty"`
	CreatedAt time.Time     `json:"created_at"`
	UpdatedAt time.Time     `json:"updated_at"`
}

// OutletResponse represents outlet data in API response
type OutletResponse struct {
	OutletID    uint              `json:"outlet_id"`
	OutletName  string            `json:"outlet_name"`
	BranchType  string            `json:"branch_type"`
	City        string            `json:"city"`
	Address     *string           `json:"address"`
	PhoneNumber *string           `json:"phone_number"`
	Status      models.StatusUmum `json:"status"`
	CreatedAt   time.Time         `json:"created_at"`
	UpdatedAt   time.Time         `json:"updated_at"`
}

// RoleResponse represents role data in API response
type RoleResponse struct {
	ID          uint                `json:"id"`
	Name        string              `json:"name"`
	Permissions []PermissionResponse `json:"permissions,omitempty"`
	CreatedAt   time.Time           `json:"created_at"`
	UpdatedAt   time.Time           `json:"updated_at"`
}

// PermissionResponse represents permission data in API response
type PermissionResponse struct {
	ID        uint      `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// CustomerResponse represents customer data in API response
type CustomerResponse struct {
	CustomerID  uint              `json:"customer_id"`
	Name        string            `json:"name"`
	PhoneNumber string            `json:"phone_number"`
	Address     *string           `json:"address"`
	Status      models.StatusUmum `json:"status"`
	Vehicles    []CustomerVehicleResponse `json:"vehicles,omitempty"`
	CreatedAt   time.Time         `json:"created_at"`
	UpdatedAt   time.Time         `json:"updated_at"`
}

// CustomerVehicleResponse represents customer vehicle data in API response
type CustomerVehicleResponse struct {
	VehicleID      uint              `json:"vehicle_id"`
	CustomerID     uint              `json:"customer_id"`
	PlateNumber    string            `json:"plate_number"`
	Brand          string            `json:"brand"`
	Model          string            `json:"model"`
	Type           string            `json:"type"`
	ProductionYear int               `json:"production_year"`
	ChassisNumber  string            `json:"chassis_number"`
	EngineNumber   string            `json:"engine_number"`
	Color          string            `json:"color"`
	Notes          *string           `json:"notes"`
	Customer       *CustomerResponse `json:"customer,omitempty"`
	CreatedAt      time.Time         `json:"created_at"`
	UpdatedAt      time.Time         `json:"updated_at"`
}

// ProductResponse represents product data in API response
type ProductResponse struct {
	ProductID          uint                      `json:"product_id"`
	ProductName        string                    `json:"product_name"`
	ProductDescription *string                   `json:"product_description"`
	ProductImage       *string                   `json:"product_image"`
	CostPrice          float64                   `json:"cost_price"`
	SellingPrice       float64                   `json:"selling_price"`
	Stock              int                       `json:"stock"`
	SKU                *string                   `json:"sku"`
	Barcode            *string                   `json:"barcode"`
	HasSerialNumber    bool                      `json:"has_serial_number"`
	ShelfLocation      *string                   `json:"shelf_location"`
	UsageStatus        models.ProductUsageStatus `json:"usage_status"`
	IsActive           bool                      `json:"is_active"`
	CategoryID         *uint                     `json:"category_id"`
	SupplierID         *uint                     `json:"supplier_id"`
	UnitTypeID         *uint                     `json:"unit_type_id"`
	Category           *CategoryResponse         `json:"category,omitempty"`
	Supplier           *SupplierResponse         `json:"supplier,omitempty"`
	UnitType           *UnitTypeResponse         `json:"unit_type,omitempty"`
	SerialNumbers      []ProductSerialNumberResponse `json:"serial_numbers,omitempty"`
	CreatedAt          time.Time                 `json:"created_at"`
	UpdatedAt          time.Time                 `json:"updated_at"`
}

// ProductSerialNumberResponse represents product serial number data in API response
type ProductSerialNumberResponse struct {
	SerialNumberID uint            `json:"serial_number_id"`
	ProductID      uint            `json:"product_id"`
	SerialNumber   string          `json:"serial_number"`
	Status         models.SNStatus `json:"status"`
	Product        *ProductResponse `json:"product,omitempty"`
	CreatedAt      time.Time       `json:"created_at"`
	UpdatedAt      time.Time       `json:"updated_at"`
}

// CategoryResponse represents category data in API response
type CategoryResponse struct {
	CategoryID uint              `json:"category_id"`
	Name       string            `json:"name"`
	Status     models.StatusUmum `json:"status"`
	Products   []ProductResponse `json:"products,omitempty"`
	CreatedAt  time.Time         `json:"created_at"`
	UpdatedAt  time.Time         `json:"updated_at"`
}

// SupplierResponse represents supplier data in API response
type SupplierResponse struct {
	SupplierID        uint              `json:"supplier_id"`
	SupplierName      string            `json:"supplier_name"`
	ContactPersonName string            `json:"contact_person_name"`
	PhoneNumber       string            `json:"phone_number"`
	Address           *string           `json:"address"`
	Status            models.StatusUmum `json:"status"`
	Products          []ProductResponse `json:"products,omitempty"`
	CreatedAt         time.Time         `json:"created_at"`
	UpdatedAt         time.Time         `json:"updated_at"`
}

// UnitTypeResponse represents unit type data in API response
type UnitTypeResponse struct {
	UnitTypeID uint              `json:"unit_type_id"`
	Name       string            `json:"name"`
	Status     models.StatusUmum `json:"status"`
	Products   []ProductResponse `json:"products,omitempty"`
	CreatedAt  time.Time         `json:"created_at"`
	UpdatedAt  time.Time         `json:"updated_at"`
}

// Helper functions to convert models to responses
func ToUserResponse(user *models.User) *UserResponse {
	response := &UserResponse{
		UserID:    user.UserID,
		Name:      user.Name,
		Email:     user.Email,
		OutletID:  user.OutletID,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}

	if user.Outlet != nil {
		response.Outlet = ToOutletResponse(user.Outlet)
	}

	return response
}

func ToOutletResponse(outlet *models.Outlet) *OutletResponse {
	return &OutletResponse{
		OutletID:    outlet.OutletID,
		OutletName:  outlet.OutletName,
		BranchType:  outlet.BranchType,
		City:        outlet.City,
		Address:     outlet.Address,
		PhoneNumber: outlet.PhoneNumber,
		Status:      outlet.Status,
		CreatedAt:   outlet.CreatedAt,
		UpdatedAt:   outlet.UpdatedAt,
	}
}

func ToRoleResponse(role *models.Role) *RoleResponse {
	response := &RoleResponse{
		ID:        role.ID,
		Name:      role.Name,
		CreatedAt: role.CreatedAt,
		UpdatedAt: role.UpdatedAt,
	}

	if role.Permissions != nil {
		for _, permission := range role.Permissions {
			response.Permissions = append(response.Permissions, *ToPermissionResponse(&permission))
		}
	}

	return response
}

func ToPermissionResponse(permission *models.Permission) *PermissionResponse {
	return &PermissionResponse{
		ID:        permission.ID,
		Name:      permission.Name,
		CreatedAt: permission.CreatedAt,
		UpdatedAt: permission.UpdatedAt,
	}
}

func ToCustomerResponse(customer *models.Customer) *CustomerResponse {
	response := &CustomerResponse{
		CustomerID:  customer.CustomerID,
		Name:        customer.Name,
		PhoneNumber: customer.PhoneNumber,
		Address:     customer.Address,
		Status:      customer.Status,
		CreatedAt:   customer.CreatedAt,
		UpdatedAt:   customer.UpdatedAt,
	}

	if customer.Vehicles != nil {
		for _, vehicle := range customer.Vehicles {
			response.Vehicles = append(response.Vehicles, *ToCustomerVehicleResponse(&vehicle))
		}
	}

	return response
}

func ToCustomerVehicleResponse(vehicle *models.CustomerVehicle) *CustomerVehicleResponse {
	response := &CustomerVehicleResponse{
		VehicleID:      vehicle.VehicleID,
		CustomerID:     vehicle.CustomerID,
		PlateNumber:    vehicle.PlateNumber,
		Brand:          vehicle.Brand,
		Model:          vehicle.Model,
		Type:           vehicle.Type,
		ProductionYear: vehicle.ProductionYear,
		ChassisNumber:  vehicle.ChassisNumber,
		EngineNumber:   vehicle.EngineNumber,
		Color:          vehicle.Color,
		Notes:          vehicle.Notes,
		CreatedAt:      vehicle.CreatedAt,
		UpdatedAt:      vehicle.UpdatedAt,
	}

	if vehicle.Customer != nil {
		response.Customer = &CustomerResponse{
			CustomerID:  vehicle.Customer.CustomerID,
			Name:        vehicle.Customer.Name,
			PhoneNumber: vehicle.Customer.PhoneNumber,
			Address:     vehicle.Customer.Address,
			Status:      vehicle.Customer.Status,
			CreatedAt:   vehicle.Customer.CreatedAt,
			UpdatedAt:   vehicle.Customer.UpdatedAt,
		}
	}

	return response
}
func ToProductResponse(product *models.Product) *ProductResponse {
response := &ProductResponse{
ProductID:          product.ProductID,
ProductName:        product.ProductName,
ProductDescription: product.ProductDescription,
ProductImage:       product.ProductImage,
CostPrice:          product.CostPrice,
SellingPrice:       product.SellingPrice,
Stock:              product.Stock,
SKU:                product.SKU,
Barcode:            product.Barcode,
HasSerialNumber:    product.HasSerialNumber,
ShelfLocation:      product.ShelfLocation,
UsageStatus:        product.UsageStatus,
IsActive:           product.IsActive,
CategoryID:         product.CategoryID,
SupplierID:         product.SupplierID,
UnitTypeID:         product.UnitTypeID,
CreatedAt:          product.CreatedAt,
UpdatedAt:          product.UpdatedAt,
}

if product.Category != nil {
response.Category = &CategoryResponse{
CategoryID: product.Category.CategoryID,
Name:       product.Category.Name,
Status:     product.Category.Status,
CreatedAt:  product.Category.CreatedAt,
UpdatedAt:  product.Category.UpdatedAt,
}
}

if product.Supplier != nil {
response.Supplier = &SupplierResponse{
SupplierID:        product.Supplier.SupplierID,
SupplierName:      product.Supplier.SupplierName,
ContactPersonName: product.Supplier.ContactPersonName,
PhoneNumber:       product.Supplier.PhoneNumber,
Address:           product.Supplier.Address,
Status:            product.Supplier.Status,
CreatedAt:         product.Supplier.CreatedAt,
UpdatedAt:         product.Supplier.UpdatedAt,
}
}

if product.UnitType != nil {
response.UnitType = &UnitTypeResponse{
UnitTypeID: product.UnitType.UnitTypeID,
Name:       product.UnitType.Name,
Status:     product.UnitType.Status,
CreatedAt:  product.UnitType.CreatedAt,
UpdatedAt:  product.UnitType.UpdatedAt,
}
}

if product.SerialNumbers != nil {
for _, sn := range product.SerialNumbers {
response.SerialNumbers = append(response.SerialNumbers, *ToProductSerialNumberResponse(&sn))
}
}

return response
}

func ToProductSerialNumberResponse(sn *models.ProductSerialNumber) *ProductSerialNumberResponse {
response := &ProductSerialNumberResponse{
SerialNumberID: sn.SerialNumberID,
ProductID:      sn.ProductID,
SerialNumber:   sn.SerialNumber,
Status:         sn.Status,
CreatedAt:      sn.CreatedAt,
UpdatedAt:      sn.UpdatedAt,
}

if sn.Product != nil {
response.Product = &ProductResponse{
ProductID:          sn.Product.ProductID,
ProductName:        sn.Product.ProductName,
ProductDescription: sn.Product.ProductDescription,
ProductImage:       sn.Product.ProductImage,
CostPrice:          sn.Product.CostPrice,
SellingPrice:       sn.Product.SellingPrice,
Stock:              sn.Product.Stock,
SKU:                sn.Product.SKU,
Barcode:            sn.Product.Barcode,
HasSerialNumber:    sn.Product.HasSerialNumber,
ShelfLocation:      sn.Product.ShelfLocation,
UsageStatus:        sn.Product.UsageStatus,
IsActive:           sn.Product.IsActive,
CategoryID:         sn.Product.CategoryID,
SupplierID:         sn.Product.SupplierID,
UnitTypeID:         sn.Product.UnitTypeID,
CreatedAt:          sn.Product.CreatedAt,
UpdatedAt:          sn.Product.UpdatedAt,
}
}

return response
}

func ToCategoryResponse(category *models.Category) *CategoryResponse {
response := &CategoryResponse{
CategoryID: category.CategoryID,
Name:       category.Name,
Status:     category.Status,
CreatedAt:  category.CreatedAt,
UpdatedAt:  category.UpdatedAt,
}

// Don't populate products to avoid circular references in API responses
return response
}

func ToSupplierResponse(supplier *models.Supplier) *SupplierResponse {
response := &SupplierResponse{
SupplierID:        supplier.SupplierID,
SupplierName:      supplier.SupplierName,
ContactPersonName: supplier.ContactPersonName,
PhoneNumber:       supplier.PhoneNumber,
Address:           supplier.Address,
Status:            supplier.Status,
CreatedAt:         supplier.CreatedAt,
UpdatedAt:         supplier.UpdatedAt,
}

// Don't populate products to avoid circular references in API responses
return response
}

func ToUnitTypeResponse(unitType *models.UnitType) *UnitTypeResponse {
response := &UnitTypeResponse{
UnitTypeID: unitType.UnitTypeID,
Name:       unitType.Name,
Status:     unitType.Status,
CreatedAt:  unitType.CreatedAt,
UpdatedAt:  unitType.UpdatedAt,
}

// Don't populate products to avoid circular references in API responses
return response
}

// ServiceResponse represents service data in API response
type ServiceResponse struct {
ServiceID         uint                      `json:"service_id"`
ServiceCode       string                    `json:"service_code"`
Name              string                    `json:"name"`
ServiceCategoryID uint                      `json:"service_category_id"`
Fee               float64                   `json:"fee"`
Status            models.StatusUmum         `json:"status"`
ServiceCategory   *ServiceCategoryResponse  `json:"service_category,omitempty"`
CreatedAt         time.Time                 `json:"created_at"`
UpdatedAt         time.Time                 `json:"updated_at"`
}

// ServiceCategoryResponse represents service category data in API response
type ServiceCategoryResponse struct {
ServiceCategoryID uint              `json:"service_category_id"`
Name              string            `json:"name"`
Status            models.StatusUmum `json:"status"`
Services          []ServiceResponse `json:"services,omitempty"`
CreatedAt         time.Time         `json:"created_at"`
UpdatedAt         time.Time         `json:"updated_at"`
}

// ServiceJobResponse represents service job data in API response
type ServiceJobResponse struct {
ServiceJobID               uint                      `json:"service_job_id"`
ServiceCode                string                    `json:"service_code"`
QueueNumber                int                       `json:"queue_number"`
CustomerID                 uint                      `json:"customer_id"`
VehicleID                  uint                      `json:"vehicle_id"`
TechnicianID               *uint                     `json:"technician_id"`
ReceivedByUserID           uint                      `json:"received_by_user_id"`
OutletID                   uint                      `json:"outlet_id"`
ProblemDescription         string                    `json:"problem_description"`
TechnicianNotes            *string                   `json:"technician_notes"`
Status                     models.ServiceStatusEnum  `json:"status"`
ServiceInDate              time.Time                 `json:"service_in_date"`
PickedUpDate               *time.Time                `json:"picked_up_date"`
ComplainDate               *time.Time                `json:"complain_date"`
WarrantyExpiresAt          *time.Time                `json:"warranty_expires_at"`
NextServiceReminderDate    *time.Time                `json:"next_service_reminder_date"`
DownPayment                float64                   `json:"down_payment"`
GrandTotal                 float64                   `json:"grand_total"`
TechnicianCommission       float64                   `json:"technician_commission"`
ShopProfit                 float64                   `json:"shop_profit"`
Customer                   *CustomerResponse         `json:"customer,omitempty"`
Vehicle                    *CustomerVehicleResponse  `json:"vehicle,omitempty"`
Technician                 *UserResponse             `json:"technician,omitempty"`
ReceivedByUser             *UserResponse             `json:"received_by_user,omitempty"`
Outlet                     *OutletResponse           `json:"outlet,omitempty"`
ServiceDetails             []ServiceDetailResponse   `json:"service_details,omitempty"`
Histories                  []ServiceJobHistoryResponse `json:"histories,omitempty"`
CreatedAt                  time.Time                 `json:"created_at"`
UpdatedAt                  time.Time                 `json:"updated_at"`
}

// ServiceDetailResponse represents service detail data in API response
type ServiceDetailResponse struct {
DetailID         uint               `json:"detail_id"`
ServiceJobID     uint               `json:"service_job_id"`
ItemID           uint               `json:"item_id"`
ItemType         string             `json:"item_type"`
Description      string             `json:"description"`
SerialNumberUsed *string            `json:"serial_number_used"`
Quantity         int                `json:"quantity"`
PricePerItem     float64            `json:"price_per_item"`
CostPerItem      float64            `json:"cost_per_item"`
ServiceJob       *ServiceJobResponse `json:"service_job,omitempty"`
}

// ServiceJobHistoryResponse represents service job history data in API response
type ServiceJobHistoryResponse struct {
HistoryID    uint               `json:"history_id"`
ServiceJobID uint               `json:"service_job_id"`
UserID       uint               `json:"user_id"`
Notes        *string            `json:"notes"`
ChangedAt    time.Time          `json:"changed_at"`
ServiceJob   *ServiceJobResponse `json:"service_job,omitempty"`
User         *UserResponse      `json:"user,omitempty"`
}

// TransactionResponse represents transaction data in API response
type TransactionResponse struct {
TransactionID      uint                      `json:"transaction_id"`
InvoiceNumber      string                    `json:"invoice_number"`
TransactionDate    time.Time                 `json:"transaction_date"`
UserID             uint                      `json:"user_id"`
CustomerID         *uint                     `json:"customer_id"`
OutletID           uint                      `json:"outlet_id"`
TransactionType    string                    `json:"transaction_type"`
Status             models.TransactionStatus  `json:"status"`
User               *UserResponse             `json:"user,omitempty"`
Customer           *CustomerResponse         `json:"customer,omitempty"`
Outlet             *OutletResponse           `json:"outlet,omitempty"`
TransactionDetails []TransactionDetailResponse `json:"transaction_details,omitempty"`
CreatedAt          time.Time                 `json:"created_at"`
UpdatedAt          time.Time                 `json:"updated_at"`
}

// TransactionDetailResponse represents transaction detail data in API response
type TransactionDetailResponse struct {
DetailID        uint                        `json:"detail_id"`
TransactionType string                      `json:"transaction_type"`
TransactionID   uint                        `json:"transaction_id"`
ProductID       *uint                       `json:"product_id"`
SerialNumberID  *uint                       `json:"serial_number_id"`
Quantity        int                         `json:"quantity"`
UnitPrice       float64                     `json:"unit_price"`
TotalPrice      float64                     `json:"total_price"`
Transaction     *TransactionResponse        `json:"transaction,omitempty"`
Product         *ProductResponse            `json:"product,omitempty"`
SerialNumber    *ProductSerialNumberResponse `json:"serial_number,omitempty"`
CreatedAt       time.Time                   `json:"created_at"`
UpdatedAt       time.Time                   `json:"updated_at"`
}

// PaymentMethodResponse represents payment method data in API response
type PaymentMethodResponse struct {
MethodID  uint              `json:"method_id"`
Name      string            `json:"name"`
Status    models.StatusUmum `json:"status"`
CreatedAt time.Time         `json:"created_at"`
UpdatedAt time.Time         `json:"updated_at"`
}

// CashFlowResponse represents cash flow data in API response
type CashFlowResponse struct {
CashFlowID uint                   `json:"cash_flow_id"`
Type       models.CashFlowType    `json:"type"`
Source     string                 `json:"source"`
Amount     float64                `json:"amount"`
Date       time.Time              `json:"date"`
Notes      *string                `json:"notes"`
UserID     uint                   `json:"user_id"`
User       *UserResponse          `json:"user,omitempty"`
CreatedAt  time.Time              `json:"created_at"`
UpdatedAt  time.Time              `json:"updated_at"`
}

// Service response transformers
func ToServiceResponse(service *models.Service) *ServiceResponse {
response := &ServiceResponse{
ServiceID:         service.ServiceID,
ServiceCode:       service.ServiceCode,
Name:              service.Name,
ServiceCategoryID: service.ServiceCategoryID,
Fee:               service.Fee,
Status:            service.Status,
CreatedAt:         service.CreatedAt,
UpdatedAt:         service.UpdatedAt,
}

if service.ServiceCategory != nil {
response.ServiceCategory = &ServiceCategoryResponse{
ServiceCategoryID: service.ServiceCategory.ServiceCategoryID,
Name:              service.ServiceCategory.Name,
Status:            service.ServiceCategory.Status,
CreatedAt:         service.ServiceCategory.CreatedAt,
UpdatedAt:         service.ServiceCategory.UpdatedAt,
}
}

return response
}

func ToServiceCategoryResponse(category *models.ServiceCategory) *ServiceCategoryResponse {
response := &ServiceCategoryResponse{
ServiceCategoryID: category.ServiceCategoryID,
Name:              category.Name,
Status:            category.Status,
CreatedAt:         category.CreatedAt,
UpdatedAt:         category.UpdatedAt,
}

if category.Services != nil {
for _, service := range category.Services {
response.Services = append(response.Services, *ToServiceResponse(&service))
}
}

return response
}

func ToServiceJobResponse(serviceJob *models.ServiceJob) *ServiceJobResponse {
response := &ServiceJobResponse{
ServiceJobID:               serviceJob.ServiceJobID,
ServiceCode:                serviceJob.ServiceCode,
QueueNumber:                serviceJob.QueueNumber,
CustomerID:                 serviceJob.CustomerID,
VehicleID:                  serviceJob.VehicleID,
TechnicianID:               serviceJob.TechnicianID,
ReceivedByUserID:           serviceJob.ReceivedByUserID,
OutletID:                   serviceJob.OutletID,
ProblemDescription:         serviceJob.ProblemDescription,
TechnicianNotes:            serviceJob.TechnicianNotes,
Status:                     serviceJob.Status,
ServiceInDate:              serviceJob.ServiceInDate,
PickedUpDate:               serviceJob.PickedUpDate,
ComplainDate:               serviceJob.ComplainDate,
WarrantyExpiresAt:          serviceJob.WarrantyExpiresAt,
NextServiceReminderDate:    serviceJob.NextServiceReminderDate,
DownPayment:                serviceJob.DownPayment,
GrandTotal:                 serviceJob.GrandTotal,
TechnicianCommission:       serviceJob.TechnicianCommission,
ShopProfit:                 serviceJob.ShopProfit,
CreatedAt:                  serviceJob.CreatedAt,
UpdatedAt:                  serviceJob.UpdatedAt,
}

if serviceJob.Customer != nil {
response.Customer = ToCustomerResponse(serviceJob.Customer)
}

if serviceJob.Vehicle != nil {
response.Vehicle = ToCustomerVehicleResponse(serviceJob.Vehicle)
}

if serviceJob.Technician != nil {
response.Technician = ToUserResponse(serviceJob.Technician)
}

if serviceJob.ReceivedByUser != nil {
response.ReceivedByUser = ToUserResponse(serviceJob.ReceivedByUser)
}

if serviceJob.Outlet != nil {
response.Outlet = ToOutletResponse(serviceJob.Outlet)
}

if serviceJob.ServiceDetails != nil {
for _, detail := range serviceJob.ServiceDetails {
response.ServiceDetails = append(response.ServiceDetails, *ToServiceDetailResponse(&detail))
}
}

if serviceJob.Histories != nil {
for _, history := range serviceJob.Histories {
response.Histories = append(response.Histories, *ToServiceJobHistoryResponse(&history))
}
}

return response
}

func ToServiceDetailResponse(detail *models.ServiceDetail) *ServiceDetailResponse {
response := &ServiceDetailResponse{
DetailID:         detail.DetailID,
ServiceJobID:     detail.ServiceJobID,
ItemID:           detail.ItemID,
ItemType:         detail.ItemType,
Description:      detail.Description,
SerialNumberUsed: detail.SerialNumberUsed,
Quantity:         detail.Quantity,
PricePerItem:     detail.PricePerItem,
CostPerItem:      detail.CostPerItem,
}

// Avoid circular reference by not including full ServiceJob
return response
}

func ToServiceJobHistoryResponse(history *models.ServiceJobHistory) *ServiceJobHistoryResponse {
response := &ServiceJobHistoryResponse{
HistoryID:    history.HistoryID,
ServiceJobID: history.ServiceJobID,
UserID:       history.UserID,
Notes:        history.Notes,
ChangedAt:    history.ChangedAt,
}

if history.User != nil {
response.User = ToUserResponse(history.User)
}

// Avoid circular reference by not including full ServiceJob
return response
}

func ToTransactionResponse(transaction *models.Transaction) *TransactionResponse {
response := &TransactionResponse{
TransactionID:   transaction.TransactionID,
InvoiceNumber:   transaction.InvoiceNumber,
TransactionDate: transaction.TransactionDate,
UserID:          transaction.UserID,
CustomerID:      transaction.CustomerID,
OutletID:        transaction.OutletID,
TransactionType: transaction.TransactionType,
Status:          transaction.Status,
CreatedAt:       transaction.CreatedAt,
UpdatedAt:       transaction.UpdatedAt,
}

if transaction.User != nil {
response.User = ToUserResponse(transaction.User)
}

if transaction.Customer != nil {
response.Customer = ToCustomerResponse(transaction.Customer)
}

if transaction.Outlet != nil {
response.Outlet = ToOutletResponse(transaction.Outlet)
}

if transaction.TransactionDetails != nil {
for _, detail := range transaction.TransactionDetails {
response.TransactionDetails = append(response.TransactionDetails, *ToTransactionDetailResponse(&detail))
}
}

return response
}

func ToTransactionDetailResponse(detail *models.TransactionDetail) *TransactionDetailResponse {
response := &TransactionDetailResponse{
DetailID:        detail.DetailID,
TransactionType: detail.TransactionType,
TransactionID:   detail.TransactionID,
ProductID:       detail.ProductID,
SerialNumberID:  detail.SerialNumberID,
Quantity:        detail.Quantity,
UnitPrice:       detail.UnitPrice,
TotalPrice:      detail.TotalPrice,
CreatedAt:       detail.CreatedAt,
UpdatedAt:       detail.UpdatedAt,
}

if detail.Product != nil {
response.Product = ToProductResponse(detail.Product)
}

if detail.SerialNumber != nil {
response.SerialNumber = ToProductSerialNumberResponse(detail.SerialNumber)
}

// Avoid circular reference by not including full Transaction
return response
}

func ToPaymentMethodResponse(paymentMethod *models.PaymentMethod) *PaymentMethodResponse {
return &PaymentMethodResponse{
MethodID:  paymentMethod.MethodID,
Name:      paymentMethod.Name,
Status:    paymentMethod.Status,
CreatedAt: paymentMethod.CreatedAt,
UpdatedAt: paymentMethod.UpdatedAt,
}
}

func ToCashFlowResponse(cashFlow *models.CashFlow) *CashFlowResponse {
response := &CashFlowResponse{
CashFlowID: cashFlow.CashFlowID,
Type:       cashFlow.Type,
Source:     cashFlow.Source,
Amount:     cashFlow.Amount,
Date:       cashFlow.Date,
Notes:      cashFlow.Notes,
UserID:     cashFlow.UserID,
CreatedAt:  cashFlow.CreatedAt,
UpdatedAt:  cashFlow.UpdatedAt,
}

if cashFlow.User != nil {
response.User = ToUserResponse(cashFlow.User)
}

return response
}
