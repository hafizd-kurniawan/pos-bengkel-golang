package handlers

import (
	"boilerplate/internal/delivery/http/responses"
	"boilerplate/internal/usecase"
	"boilerplate/internal/usecase/interfaces"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

// FoundationHandler handles foundation-related HTTP requests
type FoundationHandler struct {
	usecase *usecase.UsecaseManager
}

// NewFoundationHandler creates a new foundation handler
func NewFoundationHandler(usecase *usecase.UsecaseManager) *FoundationHandler {
	return &FoundationHandler{usecase: usecase}
}

// User handlers

// CreateUser handles user creation
func (h *FoundationHandler) CreateUser(c *fiber.Ctx) error {
	var req interfaces.CreateUserRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(responses.Response{
			Status:  "error",
			Message: "Invalid request body",
			Error:   err.Error(),
		})
	}

	user, err := h.usecase.User.CreateUser(c.Context(), req)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(responses.Response{
			Status:  "error",
			Message: "Failed to create user",
			Error:   err.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(responses.Response{
		Status:  "success",
		Message: "User created successfully",
		Data:    responses.ToUserResponse(user),
	})
}

// GetUser handles getting a single user
func (h *FoundationHandler) GetUser(c *fiber.Ctx) error {
	idParam := c.Params("id")
	id, err := strconv.ParseUint(idParam, 10, 32)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(responses.Response{
			Status:  "error",
			Message: "Invalid user ID",
			Error:   err.Error(),
		})
	}

	user, err := h.usecase.User.GetUser(c.Context(), uint(id))
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(responses.Response{
			Status:  "error",
			Message: "User not found",
			Error:   err.Error(),
		})
	}

	return c.JSON(responses.Response{
		Status:  "success",
		Message: "User retrieved successfully",
		Data:    responses.ToUserResponse(user),
	})
}

// UpdateUser handles user updates
func (h *FoundationHandler) UpdateUser(c *fiber.Ctx) error {
	idParam := c.Params("id")
	id, err := strconv.ParseUint(idParam, 10, 32)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(responses.Response{
			Status:  "error",
			Message: "Invalid user ID",
			Error:   err.Error(),
		})
	}

	var req interfaces.UpdateUserRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(responses.Response{
			Status:  "error",
			Message: "Invalid request body",
			Error:   err.Error(),
		})
	}

	user, err := h.usecase.User.UpdateUser(c.Context(), uint(id), req)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(responses.Response{
			Status:  "error",
			Message: "Failed to update user",
			Error:   err.Error(),
		})
	}

	return c.JSON(responses.Response{
		Status:  "success",
		Message: "User updated successfully",
		Data:    responses.ToUserResponse(user),
	})
}

// DeleteUser handles user deletion
func (h *FoundationHandler) DeleteUser(c *fiber.Ctx) error {
	idParam := c.Params("id")
	id, err := strconv.ParseUint(idParam, 10, 32)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(responses.Response{
			Status:  "error",
			Message: "Invalid user ID",
			Error:   err.Error(),
		})
	}

	err = h.usecase.User.DeleteUser(c.Context(), uint(id))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(responses.Response{
			Status:  "error",
			Message: "Failed to delete user",
			Error:   err.Error(),
		})
	}

	return c.JSON(responses.Response{
		Status:  "success",
		Message: "User deleted successfully",
	})
}

// ListUsers handles listing users with pagination
func (h *FoundationHandler) ListUsers(c *fiber.Ctx) error {
	page := c.QueryInt("page", 1)
	limit := c.QueryInt("limit", 10)
	offset := (page - 1) * limit

	users, err := h.usecase.User.ListUsers(c.Context(), limit, offset)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(responses.Response{
			Status:  "error",
			Message: "Failed to retrieve users",
			Error:   err.Error(),
		})
	}

	var userResponses []responses.UserResponse
	for _, user := range users {
		userResponses = append(userResponses, *responses.ToUserResponse(user))
	}

	return c.JSON(responses.Response{
		Status:  "success",
		Message: "Users retrieved successfully",
		Data:    userResponses,
	})
}

// Outlet handlers

// CreateOutlet handles outlet creation
func (h *FoundationHandler) CreateOutlet(c *fiber.Ctx) error {
	var req interfaces.CreateOutletRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(responses.Response{
			Status:  "error",
			Message: "Invalid request body",
			Error:   err.Error(),
		})
	}

	outlet, err := h.usecase.Outlet.CreateOutlet(c.Context(), req)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(responses.Response{
			Status:  "error",
			Message: "Failed to create outlet",
			Error:   err.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(responses.Response{
		Status:  "success",
		Message: "Outlet created successfully",
		Data:    responses.ToOutletResponse(outlet),
	})
}

// GetOutlet handles getting a single outlet
func (h *FoundationHandler) GetOutlet(c *fiber.Ctx) error {
	idParam := c.Params("id")
	id, err := strconv.ParseUint(idParam, 10, 32)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(responses.Response{
			Status:  "error",
			Message: "Invalid outlet ID",
			Error:   err.Error(),
		})
	}

	outlet, err := h.usecase.Outlet.GetOutlet(c.Context(), uint(id))
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(responses.Response{
			Status:  "error",
			Message: "Outlet not found",
			Error:   err.Error(),
		})
	}

	return c.JSON(responses.Response{
		Status:  "success",
		Message: "Outlet retrieved successfully",
		Data:    responses.ToOutletResponse(outlet),
	})
}

// ListOutlets handles listing outlets
func (h *FoundationHandler) ListOutlets(c *fiber.Ctx) error {
	page := c.QueryInt("page", 1)
	limit := c.QueryInt("limit", 10)
	offset := (page - 1) * limit

	outlets, err := h.usecase.Outlet.ListOutlets(c.Context(), limit, offset)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(responses.Response{
			Status:  "error",
			Message: "Failed to retrieve outlets",
			Error:   err.Error(),
		})
	}

	var outletResponses []responses.OutletResponse
	for _, outlet := range outlets {
		outletResponses = append(outletResponses, *responses.ToOutletResponse(outlet))
	}

	return c.JSON(responses.Response{
		Status:  "success",
		Message: "Outlets retrieved successfully",
		Data:    outletResponses,
	})
}