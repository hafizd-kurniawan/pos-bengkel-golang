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
