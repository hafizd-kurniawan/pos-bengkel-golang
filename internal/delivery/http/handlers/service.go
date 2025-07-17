package handlers

import (
	"boilerplate/internal/delivery/http/responses"
	"boilerplate/internal/models"
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
// ============= Service Job Handlers =============

// CreateServiceJob creates a new service job
func (h *ServiceHandler) CreateServiceJob(c *fiber.Ctx) error {
var req interfaces.CreateServiceJobRequest
if err := c.BodyParser(&req); err != nil {
return c.Status(fiber.StatusBadRequest).JSON(responses.Response{
Status:  "error",
Message: "Invalid request body",
Error:   err.Error(),
})
}

serviceJob, err := h.usecase.ServiceJob.CreateServiceJob(c.Context(), req)
if err != nil {
return c.Status(fiber.StatusInternalServerError).JSON(responses.Response{
Status:  "error",
Message: "Failed to create service job",
Error:   err.Error(),
})
}

return c.Status(fiber.StatusCreated).JSON(responses.Response{
Status:  "success",
Message: "Service job created successfully",
Data:    responses.ToServiceJobResponse(serviceJob),
})
}

// GetServiceJob retrieves a service job by ID
func (h *ServiceHandler) GetServiceJob(c *fiber.Ctx) error {
id, err := strconv.ParseUint(c.Params("id"), 10, 32)
if err != nil {
return c.Status(fiber.StatusBadRequest).JSON(responses.Response{
Status:  "error",
Message: "Invalid service job ID",
Error:   err.Error(),
})
}

serviceJob, err := h.usecase.ServiceJob.GetServiceJob(c.Context(), uint(id))
if err != nil {
return c.Status(fiber.StatusNotFound).JSON(responses.Response{
Status:  "error",
Message: "Service job not found",
Error:   err.Error(),
})
}

return c.Status(fiber.StatusOK).JSON(responses.Response{
Status:  "success",
Message: "Service job retrieved successfully",
Data:    responses.ToServiceJobResponse(serviceJob),
})
}

// ListServiceJobs retrieves service jobs with pagination
func (h *ServiceHandler) ListServiceJobs(c *fiber.Ctx) error {
limit, _ := strconv.Atoi(c.Query("limit", "10"))
offset, _ := strconv.Atoi(c.Query("offset", "0"))

serviceJobs, err := h.usecase.ServiceJob.ListServiceJobs(c.Context(), limit, offset)
if err != nil {
return c.Status(fiber.StatusInternalServerError).JSON(responses.Response{
Status:  "error",
Message: "Failed to retrieve service jobs",
Error:   err.Error(),
})
}

var serviceJobResponses []interface{}
for _, serviceJob := range serviceJobs {
serviceJobResponses = append(serviceJobResponses, responses.ToServiceJobResponse(serviceJob))
}

return c.Status(fiber.StatusOK).JSON(responses.Response{
Status:  "success",
Message: "Service jobs retrieved successfully",
Data:    serviceJobResponses,
})
}

// UpdateServiceJob updates a service job
func (h *ServiceHandler) UpdateServiceJob(c *fiber.Ctx) error {
id, err := strconv.ParseUint(c.Params("id"), 10, 32)
if err != nil {
return c.Status(fiber.StatusBadRequest).JSON(responses.Response{
Status:  "error",
Message: "Invalid service job ID",
Error:   err.Error(),
})
}

var req interfaces.UpdateServiceJobRequest
if err := c.BodyParser(&req); err != nil {
return c.Status(fiber.StatusBadRequest).JSON(responses.Response{
Status:  "error",
Message: "Invalid request body",
Error:   err.Error(),
})
}

serviceJob, err := h.usecase.ServiceJob.UpdateServiceJob(c.Context(), uint(id), req)
if err != nil {
return c.Status(fiber.StatusInternalServerError).JSON(responses.Response{
Status:  "error",
Message: "Failed to update service job",
Error:   err.Error(),
})
}

return c.Status(fiber.StatusOK).JSON(responses.Response{
Status:  "success",
Message: "Service job updated successfully",
Data:    responses.ToServiceJobResponse(serviceJob),
})
}

// DeleteServiceJob deletes a service job
func (h *ServiceHandler) DeleteServiceJob(c *fiber.Ctx) error {
id, err := strconv.ParseUint(c.Params("id"), 10, 32)
if err != nil {
return c.Status(fiber.StatusBadRequest).JSON(responses.Response{
Status:  "error",
Message: "Invalid service job ID",
Error:   err.Error(),
})
}

err = h.usecase.ServiceJob.DeleteServiceJob(c.Context(), uint(id))
if err != nil {
return c.Status(fiber.StatusInternalServerError).JSON(responses.Response{
Status:  "error",
Message: "Failed to delete service job",
Error:   err.Error(),
})
}

return c.Status(fiber.StatusOK).JSON(responses.Response{
Status:  "success",
Message: "Service job deleted successfully",
})
}

// GetServiceJobsByCustomer retrieves service jobs by customer ID
func (h *ServiceHandler) GetServiceJobsByCustomer(c *fiber.Ctx) error {
customerID, err := strconv.ParseUint(c.Params("customer_id"), 10, 32)
if err != nil {
return c.Status(fiber.StatusBadRequest).JSON(responses.Response{
Status:  "error",
Message: "Invalid customer ID",
Error:   err.Error(),
})
}

serviceJobs, err := h.usecase.ServiceJob.GetServiceJobsByCustomer(c.Context(), uint(customerID))
if err != nil {
return c.Status(fiber.StatusInternalServerError).JSON(responses.Response{
Status:  "error",
Message: "Failed to retrieve service jobs",
Error:   err.Error(),
})
}

var serviceJobResponses []interface{}
for _, serviceJob := range serviceJobs {
serviceJobResponses = append(serviceJobResponses, responses.ToServiceJobResponse(serviceJob))
}

return c.Status(fiber.StatusOK).JSON(responses.Response{
Status:  "success",
Message: "Service jobs retrieved successfully",
Data:    serviceJobResponses,
})
}

// GetServiceJobsByStatus retrieves service jobs by status
func (h *ServiceHandler) GetServiceJobsByStatus(c *fiber.Ctx) error {
status := c.Query("status")
if status == "" {
return c.Status(fiber.StatusBadRequest).JSON(responses.Response{
Status:  "error",
Message: "Status is required",
})
}

serviceJobs, err := h.usecase.ServiceJob.GetServiceJobsByStatus(c.Context(), models.ServiceStatusEnum(status))
if err != nil {
return c.Status(fiber.StatusInternalServerError).JSON(responses.Response{
Status:  "error",
Message: "Failed to retrieve service jobs",
Error:   err.Error(),
})
}

var serviceJobResponses []interface{}
for _, serviceJob := range serviceJobs {
serviceJobResponses = append(serviceJobResponses, responses.ToServiceJobResponse(serviceJob))
}

return c.Status(fiber.StatusOK).JSON(responses.Response{
Status:  "success",
Message: "Service jobs retrieved successfully",
Data:    serviceJobResponses,
})
}

// ============= Service Detail Handlers =============

// CreateServiceDetail creates a new service detail
func (h *ServiceHandler) CreateServiceDetail(c *fiber.Ctx) error {
var req interfaces.CreateServiceDetailRequest
if err := c.BodyParser(&req); err != nil {
return c.Status(fiber.StatusBadRequest).JSON(responses.Response{
Status:  "error",
Message: "Invalid request body",
Error:   err.Error(),
})
}

serviceDetail, err := h.usecase.ServiceDetail.CreateServiceDetail(c.Context(), req)
if err != nil {
return c.Status(fiber.StatusInternalServerError).JSON(responses.Response{
Status:  "error",
Message: "Failed to create service detail",
Error:   err.Error(),
})
}

return c.Status(fiber.StatusCreated).JSON(responses.Response{
Status:  "success",
Message: "Service detail created successfully",
Data:    responses.ToServiceDetailResponse(serviceDetail),
})
}

// GetServiceDetailsByServiceJob retrieves service details by service job ID
func (h *ServiceHandler) GetServiceDetailsByServiceJob(c *fiber.Ctx) error {
serviceJobID, err := strconv.ParseUint(c.Params("service_job_id"), 10, 32)
if err != nil {
return c.Status(fiber.StatusBadRequest).JSON(responses.Response{
Status:  "error",
Message: "Invalid service job ID",
Error:   err.Error(),
})
}

serviceDetails, err := h.usecase.ServiceDetail.GetServiceDetailsByServiceJob(c.Context(), uint(serviceJobID))
if err != nil {
return c.Status(fiber.StatusInternalServerError).JSON(responses.Response{
Status:  "error",
Message: "Failed to retrieve service details",
Error:   err.Error(),
})
}

var serviceDetailResponses []interface{}
for _, serviceDetail := range serviceDetails {
serviceDetailResponses = append(serviceDetailResponses, responses.ToServiceDetailResponse(serviceDetail))
}

return c.Status(fiber.StatusOK).JSON(responses.Response{
Status:  "success",
Message: "Service details retrieved successfully",
Data:    serviceDetailResponses,
})
}

// UpdateServiceDetail updates a service detail
func (h *ServiceHandler) UpdateServiceDetail(c *fiber.Ctx) error {
id, err := strconv.ParseUint(c.Params("id"), 10, 32)
if err != nil {
return c.Status(fiber.StatusBadRequest).JSON(responses.Response{
Status:  "error",
Message: "Invalid service detail ID",
Error:   err.Error(),
})
}

var req interfaces.UpdateServiceDetailRequest
if err := c.BodyParser(&req); err != nil {
return c.Status(fiber.StatusBadRequest).JSON(responses.Response{
Status:  "error",
Message: "Invalid request body",
Error:   err.Error(),
})
}

serviceDetail, err := h.usecase.ServiceDetail.UpdateServiceDetail(c.Context(), uint(id), req)
if err != nil {
return c.Status(fiber.StatusInternalServerError).JSON(responses.Response{
Status:  "error",
Message: "Failed to update service detail",
Error:   err.Error(),
})
}

return c.Status(fiber.StatusOK).JSON(responses.Response{
Status:  "success",
Message: "Service detail updated successfully",
Data:    responses.ToServiceDetailResponse(serviceDetail),
})
}

// DeleteServiceDetail deletes a service detail
func (h *ServiceHandler) DeleteServiceDetail(c *fiber.Ctx) error {
id, err := strconv.ParseUint(c.Params("id"), 10, 32)
if err != nil {
return c.Status(fiber.StatusBadRequest).JSON(responses.Response{
Status:  "error",
Message: "Invalid service detail ID",
Error:   err.Error(),
})
}

err = h.usecase.ServiceDetail.DeleteServiceDetail(c.Context(), uint(id))
if err != nil {
return c.Status(fiber.StatusInternalServerError).JSON(responses.Response{
Status:  "error",
Message: "Failed to delete service detail",
Error:   err.Error(),
})
}

return c.Status(fiber.StatusOK).JSON(responses.Response{
Status:  "success",
Message: "Service detail deleted successfully",
})
}

// ============= Service Job History Handlers =============

// GetServiceJobHistoriesByServiceJob retrieves service job histories by service job ID
func (h *ServiceHandler) GetServiceJobHistoriesByServiceJob(c *fiber.Ctx) error {
serviceJobID, err := strconv.ParseUint(c.Params("service_job_id"), 10, 32)
if err != nil {
return c.Status(fiber.StatusBadRequest).JSON(responses.Response{
Status:  "error",
Message: "Invalid service job ID",
Error:   err.Error(),
})
}

histories, err := h.usecase.ServiceJobHistory.GetServiceJobHistoriesByServiceJob(c.Context(), uint(serviceJobID))
if err != nil {
return c.Status(fiber.StatusInternalServerError).JSON(responses.Response{
Status:  "error",
Message: "Failed to retrieve service job histories",
Error:   err.Error(),
})
}

var historyResponses []interface{}
for _, history := range histories {
historyResponses = append(historyResponses, responses.ToServiceJobHistoryResponse(history))
}

return c.Status(fiber.StatusOK).JSON(responses.Response{
Status:  "success",
Message: "Service job histories retrieved successfully",
Data:    historyResponses,
})
}

// GetServiceJobByServiceCode retrieves a service job by service code
func (h *ServiceHandler) GetServiceJobByServiceCode(c *fiber.Ctx) error {
serviceCode := c.Query("service_code")
if serviceCode == "" {
return c.Status(fiber.StatusBadRequest).JSON(responses.Response{
Status:  "error",
Message: "Service code is required",
})
}

serviceJob, err := h.usecase.ServiceJob.GetServiceJobByServiceCode(c.Context(), serviceCode)
if err != nil {
return c.Status(fiber.StatusNotFound).JSON(responses.Response{
Status:  "error",
Message: "Service job not found",
Error:   err.Error(),
})
}

return c.Status(fiber.StatusOK).JSON(responses.Response{
Status:  "success",
Message: "Service job retrieved successfully",
Data:    responses.ToServiceJobResponse(serviceJob),
})
}

// UpdateServiceJobStatus updates service job status and creates history
func (h *ServiceHandler) UpdateServiceJobStatus(c *fiber.Ctx) error {
id, err := strconv.ParseUint(c.Params("id"), 10, 32)
if err != nil {
return c.Status(fiber.StatusBadRequest).JSON(responses.Response{
Status:  "error",
Message: "Invalid service job ID",
Error:   err.Error(),
})
}

var req struct {
Status string  `json:"status" validate:"required"`
UserID uint    `json:"user_id" validate:"required"`
Notes  *string `json:"notes,omitempty"`
}
if err := c.BodyParser(&req); err != nil {
return c.Status(fiber.StatusBadRequest).JSON(responses.Response{
Status:  "error",
Message: "Invalid request body",
Error:   err.Error(),
})
}

err = h.usecase.ServiceJob.UpdateServiceJobStatus(c.Context(), uint(id), models.ServiceStatusEnum(req.Status), req.UserID, req.Notes)
if err != nil {
return c.Status(fiber.StatusInternalServerError).JSON(responses.Response{
Status:  "error",
Message: "Failed to update service job status",
Error:   err.Error(),
})
}

return c.Status(fiber.StatusOK).JSON(responses.Response{
Status:  "success",
Message: "Service job status updated successfully",
})
}
