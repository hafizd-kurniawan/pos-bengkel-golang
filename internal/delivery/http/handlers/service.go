package handlers

import (
	"boilerplate/internal/delivery/http/responses"
	"boilerplate/internal/usecase"
	"boilerplate/internal/usecase/interfaces"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

// ServiceHandler handles service-related HTTP requests
type ServiceHandler struct {
	usecase *usecase.UsecaseManager
}

// NewServiceHandler creates a new service handler
func NewServiceHandler(usecase *usecase.UsecaseManager) *ServiceHandler {
	return &ServiceHandler{usecase: usecase}
}

// ============= Service Category Handlers =============

// CreateServiceCategory creates a new service category
func (h *ServiceHandler) CreateServiceCategory(c *fiber.Ctx) error {
	var req interfaces.CreateServiceCategoryRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(responses.Response{
			Status:  "error",
			Message: "Invalid request body",
			Error:   err.Error(),
		})
	}

	serviceCategory, err := h.usecase.ServiceCategory.CreateServiceCategory(c.Context(), req)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(responses.Response{
			Status:  "error",
			Message: "Failed to create service category",
			Error:   err.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(responses.Response{
		Status:  "success",
		Message: "Service category created successfully",
		Data:    serviceCategory,
	})
}

// ListServiceCategories lists all service categories with pagination
func (h *ServiceHandler) ListServiceCategories(c *fiber.Ctx) error {
	limit, _ := strconv.Atoi(c.Query("limit", "10"))
	offset, _ := strconv.Atoi(c.Query("offset", "0"))

	serviceCategories, err := h.usecase.ServiceCategory.ListServiceCategories(c.Context(), limit, offset)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(responses.Response{
			Status:  "error",
			Message: "Failed to retrieve service categories",
			Error:   err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(responses.Response{
		Status:  "success",
		Message: "Service categories retrieved successfully",
		Data:    serviceCategories,
	})
}

// GetServiceCategory retrieves a service category by ID
func (h *ServiceHandler) GetServiceCategory(c *fiber.Ctx) error {
	id, err := strconv.ParseUint(c.Params("id"), 10, 32)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(responses.Response{
			Status:  "error",
			Message: "Invalid service category ID",
			Error:   err.Error(),
		})
	}

	serviceCategory, err := h.usecase.ServiceCategory.GetServiceCategory(c.Context(), uint(id))
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(responses.Response{
			Status:  "error",
			Message: "Service category not found",
			Error:   err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(responses.Response{
		Status:  "success",
		Message: "Service category retrieved successfully",
		Data:    serviceCategory,
	})
}

// UpdateServiceCategory updates a service category
func (h *ServiceHandler) UpdateServiceCategory(c *fiber.Ctx) error {
	id, err := strconv.ParseUint(c.Params("id"), 10, 32)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(responses.Response{
			Status:  "error",
			Message: "Invalid service category ID",
			Error:   err.Error(),
		})
	}

	var req interfaces.UpdateServiceCategoryRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(responses.Response{
			Status:  "error",
			Message: "Invalid request body",
			Error:   err.Error(),
		})
	}

	serviceCategory, err := h.usecase.ServiceCategory.UpdateServiceCategory(c.Context(), uint(id), req)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(responses.Response{
			Status:  "error",
			Message: "Failed to update service category",
			Error:   err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(responses.Response{
		Status:  "success",
		Message: "Service category updated successfully",
		Data:    serviceCategory,
	})
}

// DeleteServiceCategory deletes a service category
func (h *ServiceHandler) DeleteServiceCategory(c *fiber.Ctx) error {
	id, err := strconv.ParseUint(c.Params("id"), 10, 32)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(responses.Response{
			Status:  "error",
			Message: "Invalid service category ID",
			Error:   err.Error(),
		})
	}

	err = h.usecase.ServiceCategory.DeleteServiceCategory(c.Context(), uint(id))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(responses.Response{
			Status:  "error",
			Message: "Failed to delete service category",
			Error:   err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(responses.Response{
		Status:  "success",
		Message: "Service category deleted successfully",
		Data:    nil,
	})
}

// ============= Service Handlers =============

// CreateService creates a new service
func (h *ServiceHandler) CreateService(c *fiber.Ctx) error {
	var req interfaces.CreateServiceRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(responses.Response{
			Status:  "error",
			Message: "Invalid request body",
			Error:   err.Error(),
		})
	}

	service, err := h.usecase.Service.CreateService(c.Context(), req)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(responses.Response{
			Status:  "error",
			Message: "Failed to create service",
			Error:   err.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(responses.Response{
		Status:  "success",
		Message: "Service created successfully",
		Data:    service,
	})
}

// ListServices lists all services with pagination
func (h *ServiceHandler) ListServices(c *fiber.Ctx) error {
	limit, _ := strconv.Atoi(c.Query("limit", "10"))
	offset, _ := strconv.Atoi(c.Query("offset", "0"))

	services, err := h.usecase.Service.ListServices(c.Context(), limit, offset)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(responses.Response{
			Status:  "error",
			Message: "Failed to retrieve services",
			Error:   err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(responses.Response{
		Status:  "success",
		Message: "Services retrieved successfully",
		Data:    services,
	})
}

// GetService retrieves a service by ID
func (h *ServiceHandler) GetService(c *fiber.Ctx) error {
	id, err := strconv.ParseUint(c.Params("id"), 10, 32)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(responses.Response{
			Status:  "error",
			Message: "Invalid service ID",
			Error:   err.Error(),
		})
	}

	service, err := h.usecase.Service.GetService(c.Context(), uint(id))
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(responses.Response{
			Status:  "error",
			Message: "Service not found",
			Error:   err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(responses.Response{
		Status:  "success",
		Message: "Service retrieved successfully",
		Data:    service,
	})
}

// GetServiceByCode retrieves a service by service code
func (h *ServiceHandler) GetServiceByCode(c *fiber.Ctx) error {
	serviceCode := c.Query("service_code")
	if serviceCode == "" {
		return c.Status(fiber.StatusBadRequest).JSON(responses.Response{
			Status:  "error",
			Message: "Service code is required",
			Error:   "missing service_code query parameter",
		})
	}

	service, err := h.usecase.Service.GetServiceByServiceCode(c.Context(), serviceCode)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(responses.Response{
			Status:  "error",
			Message: "Service not found",
			Error:   err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(responses.Response{
		Status:  "success",
		Message: "Service retrieved successfully",
		Data:    service,
	})
}

// UpdateService updates a service
func (h *ServiceHandler) UpdateService(c *fiber.Ctx) error {
	id, err := strconv.ParseUint(c.Params("id"), 10, 32)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(responses.Response{
			Status:  "error",
			Message: "Invalid service ID",
			Error:   err.Error(),
		})
	}

	var req interfaces.UpdateServiceRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(responses.Response{
			Status:  "error",
			Message: "Invalid request body",
			Error:   err.Error(),
		})
	}

	service, err := h.usecase.Service.UpdateService(c.Context(), uint(id), req)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(responses.Response{
			Status:  "error",
			Message: "Failed to update service",
			Error:   err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(responses.Response{
		Status:  "success",
		Message: "Service updated successfully",
		Data:    service,
	})
}

// DeleteService deletes a service
func (h *ServiceHandler) DeleteService(c *fiber.Ctx) error {
	id, err := strconv.ParseUint(c.Params("id"), 10, 32)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(responses.Response{
			Status:  "error",
			Message: "Invalid service ID",
			Error:   err.Error(),
		})
	}

	err = h.usecase.Service.DeleteService(c.Context(), uint(id))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(responses.Response{
			Status:  "error",
			Message: "Failed to delete service",
			Error:   err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(responses.Response{
		Status:  "success",
		Message: "Service deleted successfully",
		Data:    nil,
	})
}

// SearchServices searches services by name or service code
func (h *ServiceHandler) SearchServices(c *fiber.Ctx) error {
	query := c.Query("q")
	if query == "" {
		return c.Status(fiber.StatusBadRequest).JSON(responses.Response{
			Status:  "error",
			Message: "Search query is required",
			Error:   "missing q query parameter",
		})
	}

	limit, _ := strconv.Atoi(c.Query("limit", "10"))
	offset, _ := strconv.Atoi(c.Query("offset", "0"))

	services, err := h.usecase.Service.SearchServices(c.Context(), query, limit, offset)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(responses.Response{
			Status:  "error",
			Message: "Failed to search services",
			Error:   err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(responses.Response{
		Status:  "success",
		Message: "Services search completed successfully",
		Data:    services,
	})
}

// GetServicesByCategory retrieves services by category ID
func (h *ServiceHandler) GetServicesByCategory(c *fiber.Ctx) error {
	categoryID, err := strconv.ParseUint(c.Params("category_id"), 10, 32)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(responses.Response{
			Status:  "error",
			Message: "Invalid category ID",
			Error:   err.Error(),
		})
	}

	services, err := h.usecase.Service.GetServicesByCategory(c.Context(), uint(categoryID))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(responses.Response{
			Status:  "error",
			Message: "Failed to retrieve services",
			Error:   err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(responses.Response{
		Status:  "success",
		Message: "Services retrieved successfully",
		Data:    services,
	})
}