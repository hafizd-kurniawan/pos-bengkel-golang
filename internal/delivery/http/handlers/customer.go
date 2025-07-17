package handlers

import (
	"boilerplate/internal/delivery/http/responses"
	"boilerplate/internal/usecase"
	"boilerplate/internal/usecase/interfaces"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

// CustomerHandler handles customer-related HTTP requests
type CustomerHandler struct {
	usecase *usecase.UsecaseManager
}

// NewCustomerHandler creates a new customer handler
func NewCustomerHandler(usecase *usecase.UsecaseManager) *CustomerHandler {
	return &CustomerHandler{usecase: usecase}
}

// Customer handlers

// CreateCustomer handles customer creation
func (h *CustomerHandler) CreateCustomer(c *fiber.Ctx) error {
	var req interfaces.CreateCustomerRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(responses.Response{
			Status:  "error",
			Message: "Invalid request body",
			Error:   err.Error(),
		})
	}

	customer, err := h.usecase.Customer.CreateCustomer(c.Context(), req)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(responses.Response{
			Status:  "error",
			Message: "Failed to create customer",
			Error:   err.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(responses.Response{
		Status:  "success",
		Message: "Customer created successfully",
		Data:    responses.ToCustomerResponse(customer),
	})
}

// GetCustomer handles getting a single customer
func (h *CustomerHandler) GetCustomer(c *fiber.Ctx) error {
	idParam := c.Params("id")
	id, err := strconv.ParseUint(idParam, 10, 32)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(responses.Response{
			Status:  "error",
			Message: "Invalid customer ID",
			Error:   err.Error(),
		})
	}

	customer, err := h.usecase.Customer.GetCustomer(c.Context(), uint(id))
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(responses.Response{
			Status:  "error",
			Message: "Customer not found",
			Error:   err.Error(),
		})
	}

	return c.JSON(responses.Response{
		Status:  "success",
		Message: "Customer retrieved successfully",
		Data:    responses.ToCustomerResponse(customer),
	})
}

// GetCustomerByPhoneNumber handles getting a customer by phone number
func (h *CustomerHandler) GetCustomerByPhoneNumber(c *fiber.Ctx) error {
	phoneNumber := c.Query("phone_number")
	if phoneNumber == "" {
		return c.Status(fiber.StatusBadRequest).JSON(responses.Response{
			Status:  "error",
			Message: "Phone number is required",
		})
	}

	customer, err := h.usecase.Customer.GetCustomerByPhoneNumber(c.Context(), phoneNumber)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(responses.Response{
			Status:  "error",
			Message: "Customer not found",
			Error:   err.Error(),
		})
	}

	return c.JSON(responses.Response{
		Status:  "success",
		Message: "Customer retrieved successfully",
		Data:    responses.ToCustomerResponse(customer),
	})
}

// UpdateCustomer handles customer updates
func (h *CustomerHandler) UpdateCustomer(c *fiber.Ctx) error {
	idParam := c.Params("id")
	id, err := strconv.ParseUint(idParam, 10, 32)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(responses.Response{
			Status:  "error",
			Message: "Invalid customer ID",
			Error:   err.Error(),
		})
	}

	var req interfaces.UpdateCustomerRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(responses.Response{
			Status:  "error",
			Message: "Invalid request body",
			Error:   err.Error(),
		})
	}

	customer, err := h.usecase.Customer.UpdateCustomer(c.Context(), uint(id), req)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(responses.Response{
			Status:  "error",
			Message: "Failed to update customer",
			Error:   err.Error(),
		})
	}

	return c.JSON(responses.Response{
		Status:  "success",
		Message: "Customer updated successfully",
		Data:    responses.ToCustomerResponse(customer),
	})
}

// DeleteCustomer handles customer deletion
func (h *CustomerHandler) DeleteCustomer(c *fiber.Ctx) error {
	idParam := c.Params("id")
	id, err := strconv.ParseUint(idParam, 10, 32)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(responses.Response{
			Status:  "error",
			Message: "Invalid customer ID",
			Error:   err.Error(),
		})
	}

	err = h.usecase.Customer.DeleteCustomer(c.Context(), uint(id))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(responses.Response{
			Status:  "error",
			Message: "Failed to delete customer",
			Error:   err.Error(),
		})
	}

	return c.JSON(responses.Response{
		Status:  "success",
		Message: "Customer deleted successfully",
	})
}

// ListCustomers handles listing customers with pagination
func (h *CustomerHandler) ListCustomers(c *fiber.Ctx) error {
	page := c.QueryInt("page", 1)
	limit := c.QueryInt("limit", 10)
	offset := (page - 1) * limit

	customers, err := h.usecase.Customer.ListCustomers(c.Context(), limit, offset)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(responses.Response{
			Status:  "error",
			Message: "Failed to retrieve customers",
			Error:   err.Error(),
		})
	}

	var customerResponses []responses.CustomerResponse
	for _, customer := range customers {
		customerResponses = append(customerResponses, *responses.ToCustomerResponse(customer))
	}

	return c.JSON(responses.Response{
		Status:  "success",
		Message: "Customers retrieved successfully",
		Data:    customerResponses,
	})
}

// SearchCustomers handles searching customers
func (h *CustomerHandler) SearchCustomers(c *fiber.Ctx) error {
	query := c.Query("q")
	if query == "" {
		return c.Status(fiber.StatusBadRequest).JSON(responses.Response{
			Status:  "error",
			Message: "Search query is required",
		})
	}

	page := c.QueryInt("page", 1)
	limit := c.QueryInt("limit", 10)
	offset := (page - 1) * limit

	customers, err := h.usecase.Customer.SearchCustomers(c.Context(), query, limit, offset)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(responses.Response{
			Status:  "error",
			Message: "Failed to search customers",
			Error:   err.Error(),
		})
	}

	var customerResponses []responses.CustomerResponse
	for _, customer := range customers {
		customerResponses = append(customerResponses, *responses.ToCustomerResponse(customer))
	}

	return c.JSON(responses.Response{
		Status:  "success",
		Message: "Customers search completed successfully",
		Data:    customerResponses,
	})
}

// Customer Vehicle handlers

// CreateCustomerVehicle handles customer vehicle creation
func (h *CustomerHandler) CreateCustomerVehicle(c *fiber.Ctx) error {
	var req interfaces.CreateCustomerVehicleRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(responses.Response{
			Status:  "error",
			Message: "Invalid request body",
			Error:   err.Error(),
		})
	}

	vehicle, err := h.usecase.CustomerVehicle.CreateCustomerVehicle(c.Context(), req)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(responses.Response{
			Status:  "error",
			Message: "Failed to create customer vehicle",
			Error:   err.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(responses.Response{
		Status:  "success",
		Message: "Customer vehicle created successfully",
		Data:    responses.ToCustomerVehicleResponse(vehicle),
	})
}

// GetCustomerVehicle handles getting a single customer vehicle
func (h *CustomerHandler) GetCustomerVehicle(c *fiber.Ctx) error {
	idParam := c.Params("id")
	id, err := strconv.ParseUint(idParam, 10, 32)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(responses.Response{
			Status:  "error",
			Message: "Invalid customer vehicle ID",
			Error:   err.Error(),
		})
	}

	vehicle, err := h.usecase.CustomerVehicle.GetCustomerVehicle(c.Context(), uint(id))
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(responses.Response{
			Status:  "error",
			Message: "Customer vehicle not found",
			Error:   err.Error(),
		})
	}

	return c.JSON(responses.Response{
		Status:  "success",
		Message: "Customer vehicle retrieved successfully",
		Data:    responses.ToCustomerVehicleResponse(vehicle),
	})
}

// UpdateCustomerVehicle handles customer vehicle updates
func (h *CustomerHandler) UpdateCustomerVehicle(c *fiber.Ctx) error {
	idParam := c.Params("id")
	id, err := strconv.ParseUint(idParam, 10, 32)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(responses.Response{
			Status:  "error",
			Message: "Invalid customer vehicle ID",
			Error:   err.Error(),
		})
	}

	var req interfaces.UpdateCustomerVehicleRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(responses.Response{
			Status:  "error",
			Message: "Invalid request body",
			Error:   err.Error(),
		})
	}

	vehicle, err := h.usecase.CustomerVehicle.UpdateCustomerVehicle(c.Context(), uint(id), req)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(responses.Response{
			Status:  "error",
			Message: "Failed to update customer vehicle",
			Error:   err.Error(),
		})
	}

	return c.JSON(responses.Response{
		Status:  "success",
		Message: "Customer vehicle updated successfully",
		Data:    responses.ToCustomerVehicleResponse(vehicle),
	})
}

// DeleteCustomerVehicle handles customer vehicle deletion
func (h *CustomerHandler) DeleteCustomerVehicle(c *fiber.Ctx) error {
	idParam := c.Params("id")
	id, err := strconv.ParseUint(idParam, 10, 32)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(responses.Response{
			Status:  "error",
			Message: "Invalid customer vehicle ID",
			Error:   err.Error(),
		})
	}

	err = h.usecase.CustomerVehicle.DeleteCustomerVehicle(c.Context(), uint(id))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(responses.Response{
			Status:  "error",
			Message: "Failed to delete customer vehicle",
			Error:   err.Error(),
		})
	}

	return c.JSON(responses.Response{
		Status:  "success",
		Message: "Customer vehicle deleted successfully",
	})
}

// ListCustomerVehicles handles listing customer vehicles with pagination
func (h *CustomerHandler) ListCustomerVehicles(c *fiber.Ctx) error {
	page := c.QueryInt("page", 1)
	limit := c.QueryInt("limit", 10)
	offset := (page - 1) * limit

	vehicles, err := h.usecase.CustomerVehicle.ListCustomerVehicles(c.Context(), limit, offset)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(responses.Response{
			Status:  "error",
			Message: "Failed to retrieve customer vehicles",
			Error:   err.Error(),
		})
	}

	var vehicleResponses []responses.CustomerVehicleResponse
	for _, vehicle := range vehicles {
		vehicleResponses = append(vehicleResponses, *responses.ToCustomerVehicleResponse(vehicle))
	}

	return c.JSON(responses.Response{
		Status:  "success",
		Message: "Customer vehicles retrieved successfully",
		Data:    vehicleResponses,
	})
}

// GetCustomerVehiclesByCustomerID handles getting customer vehicles by customer ID
func (h *CustomerHandler) GetCustomerVehiclesByCustomerID(c *fiber.Ctx) error {
	idParam := c.Params("customer_id")
	customerID, err := strconv.ParseUint(idParam, 10, 32)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(responses.Response{
			Status:  "error",
			Message: "Invalid customer ID",
			Error:   err.Error(),
		})
	}

	vehicles, err := h.usecase.CustomerVehicle.GetCustomerVehiclesByCustomerID(c.Context(), uint(customerID))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(responses.Response{
			Status:  "error",
			Message: "Failed to retrieve customer vehicles",
			Error:   err.Error(),
		})
	}

	var vehicleResponses []responses.CustomerVehicleResponse
	for _, vehicle := range vehicles {
		vehicleResponses = append(vehicleResponses, *responses.ToCustomerVehicleResponse(vehicle))
	}

	return c.JSON(responses.Response{
		Status:  "success",
		Message: "Customer vehicles retrieved successfully",
		Data:    vehicleResponses,
	})
}

// SearchCustomerVehicles handles searching customer vehicles
func (h *CustomerHandler) SearchCustomerVehicles(c *fiber.Ctx) error {
	query := c.Query("q")
	if query == "" {
		return c.Status(fiber.StatusBadRequest).JSON(responses.Response{
			Status:  "error",
			Message: "Search query is required",
		})
	}

	page := c.QueryInt("page", 1)
	limit := c.QueryInt("limit", 10)
	offset := (page - 1) * limit

	vehicles, err := h.usecase.CustomerVehicle.SearchCustomerVehicles(c.Context(), query, limit, offset)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(responses.Response{
			Status:  "error",
			Message: "Failed to search customer vehicles",
			Error:   err.Error(),
		})
	}

	var vehicleResponses []responses.CustomerVehicleResponse
	for _, vehicle := range vehicles {
		vehicleResponses = append(vehicleResponses, *responses.ToCustomerVehicleResponse(vehicle))
	}

	return c.JSON(responses.Response{
		Status:  "success",
		Message: "Customer vehicles search completed successfully",
		Data:    vehicleResponses,
	})
}