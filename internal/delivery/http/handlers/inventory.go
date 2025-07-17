package handlers

import (
	"boilerplate/internal/delivery/http/responses"
	"boilerplate/internal/models"
	"boilerplate/internal/usecase"
	"boilerplate/internal/usecase/interfaces"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

// InventoryHandler handles inventory-related HTTP requests
type InventoryHandler struct {
	usecase *usecase.UsecaseManager
}

// NewInventoryHandler creates a new inventory handler
func NewInventoryHandler(usecase *usecase.UsecaseManager) *InventoryHandler {
	return &InventoryHandler{usecase: usecase}
}

// Product handlers

// CreateProduct handles product creation
func (h *InventoryHandler) CreateProduct(c *fiber.Ctx) error {
	var req interfaces.CreateProductRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(responses.Response{
			Status:  "error",
			Message: "Invalid request body",
			Error:   err.Error(),
		})
	}

	product, err := h.usecase.Product.CreateProduct(c.Context(), req)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(responses.Response{
			Status:  "error",
			Message: "Failed to create product",
			Error:   err.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(responses.Response{
		Status:  "success",
		Message: "Product created successfully",
		Data:    responses.ToProductResponse(product),
	})
}

// GetProduct handles getting a single product
func (h *InventoryHandler) GetProduct(c *fiber.Ctx) error {
	idParam := c.Params("id")
	id, err := strconv.ParseUint(idParam, 10, 32)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(responses.Response{
			Status:  "error",
			Message: "Invalid product ID",
			Error:   err.Error(),
		})
	}

	product, err := h.usecase.Product.GetProduct(c.Context(), uint(id))
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(responses.Response{
			Status:  "error",
			Message: "Product not found",
			Error:   err.Error(),
		})
	}

	return c.JSON(responses.Response{
		Status:  "success",
		Message: "Product retrieved successfully",
		Data:    responses.ToProductResponse(product),
	})
}

// GetProductBySKU handles getting a product by SKU
func (h *InventoryHandler) GetProductBySKU(c *fiber.Ctx) error {
	sku := c.Query("sku")
	if sku == "" {
		return c.Status(fiber.StatusBadRequest).JSON(responses.Response{
			Status:  "error",
			Message: "SKU is required",
		})
	}

	product, err := h.usecase.Product.GetProductBySKU(c.Context(), sku)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(responses.Response{
			Status:  "error",
			Message: "Product not found",
			Error:   err.Error(),
		})
	}

	return c.JSON(responses.Response{
		Status:  "success",
		Message: "Product retrieved successfully",
		Data:    responses.ToProductResponse(product),
	})
}

// GetProductByBarcode handles getting a product by barcode
func (h *InventoryHandler) GetProductByBarcode(c *fiber.Ctx) error {
	barcode := c.Query("barcode")
	if barcode == "" {
		return c.Status(fiber.StatusBadRequest).JSON(responses.Response{
			Status:  "error",
			Message: "Barcode is required",
		})
	}

	product, err := h.usecase.Product.GetProductByBarcode(c.Context(), barcode)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(responses.Response{
			Status:  "error",
			Message: "Product not found",
			Error:   err.Error(),
		})
	}

	return c.JSON(responses.Response{
		Status:  "success",
		Message: "Product retrieved successfully",
		Data:    responses.ToProductResponse(product),
	})
}

// UpdateProduct handles product updates
func (h *InventoryHandler) UpdateProduct(c *fiber.Ctx) error {
	idParam := c.Params("id")
	id, err := strconv.ParseUint(idParam, 10, 32)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(responses.Response{
			Status:  "error",
			Message: "Invalid product ID",
			Error:   err.Error(),
		})
	}

	var req interfaces.UpdateProductRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(responses.Response{
			Status:  "error",
			Message: "Invalid request body",
			Error:   err.Error(),
		})
	}

	product, err := h.usecase.Product.UpdateProduct(c.Context(), uint(id), req)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(responses.Response{
			Status:  "error",
			Message: "Failed to update product",
			Error:   err.Error(),
		})
	}

	return c.JSON(responses.Response{
		Status:  "success",
		Message: "Product updated successfully",
		Data:    responses.ToProductResponse(product),
	})
}

// DeleteProduct handles product deletion
func (h *InventoryHandler) DeleteProduct(c *fiber.Ctx) error {
	idParam := c.Params("id")
	id, err := strconv.ParseUint(idParam, 10, 32)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(responses.Response{
			Status:  "error",
			Message: "Invalid product ID",
			Error:   err.Error(),
		})
	}

	err = h.usecase.Product.DeleteProduct(c.Context(), uint(id))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(responses.Response{
			Status:  "error",
			Message: "Failed to delete product",
			Error:   err.Error(),
		})
	}

	return c.JSON(responses.Response{
		Status:  "success",
		Message: "Product deleted successfully",
	})
}

// ListProducts handles listing products with pagination
func (h *InventoryHandler) ListProducts(c *fiber.Ctx) error {
	page := c.QueryInt("page", 1)
	limit := c.QueryInt("limit", 10)
	offset := (page - 1) * limit

	products, err := h.usecase.Product.ListProducts(c.Context(), limit, offset)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(responses.Response{
			Status:  "error",
			Message: "Failed to retrieve products",
			Error:   err.Error(),
		})
	}

	var productResponses []responses.ProductResponse
	for _, product := range products {
		productResponses = append(productResponses, *responses.ToProductResponse(product))
	}

	return c.JSON(responses.Response{
		Status:  "success",
		Message: "Products retrieved successfully",
		Data:    productResponses,
	})
}

// GetProductsByCategory handles getting products by category
func (h *InventoryHandler) GetProductsByCategory(c *fiber.Ctx) error {
	idParam := c.Params("category_id")
	categoryID, err := strconv.ParseUint(idParam, 10, 32)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(responses.Response{
			Status:  "error",
			Message: "Invalid category ID",
			Error:   err.Error(),
		})
	}

	products, err := h.usecase.Product.GetProductsByCategory(c.Context(), uint(categoryID))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(responses.Response{
			Status:  "error",
			Message: "Failed to retrieve products",
			Error:   err.Error(),
		})
	}

	var productResponses []responses.ProductResponse
	for _, product := range products {
		productResponses = append(productResponses, *responses.ToProductResponse(product))
	}

	return c.JSON(responses.Response{
		Status:  "success",
		Message: "Products retrieved successfully",
		Data:    productResponses,
	})
}

// GetProductsBySupplier handles getting products by supplier
func (h *InventoryHandler) GetProductsBySupplier(c *fiber.Ctx) error {
	idParam := c.Params("supplier_id")
	supplierID, err := strconv.ParseUint(idParam, 10, 32)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(responses.Response{
			Status:  "error",
			Message: "Invalid supplier ID",
			Error:   err.Error(),
		})
	}

	products, err := h.usecase.Product.GetProductsBySupplier(c.Context(), uint(supplierID))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(responses.Response{
			Status:  "error",
			Message: "Failed to retrieve products",
			Error:   err.Error(),
		})
	}

	var productResponses []responses.ProductResponse
	for _, product := range products {
		productResponses = append(productResponses, *responses.ToProductResponse(product))
	}

	return c.JSON(responses.Response{
		Status:  "success",
		Message: "Products retrieved successfully",
		Data:    productResponses,
	})
}

// GetProductsByUsageStatus handles getting products by usage status
func (h *InventoryHandler) GetProductsByUsageStatus(c *fiber.Ctx) error {
	statusParam := c.Query("usage_status")
	if statusParam == "" {
		return c.Status(fiber.StatusBadRequest).JSON(responses.Response{
			Status:  "error",
			Message: "Usage status is required",
		})
	}

	status := models.ProductUsageStatus(statusParam)
	products, err := h.usecase.Product.GetProductsByUsageStatus(c.Context(), status)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(responses.Response{
			Status:  "error",
			Message: "Failed to retrieve products",
			Error:   err.Error(),
		})
	}

	var productResponses []responses.ProductResponse
	for _, product := range products {
		productResponses = append(productResponses, *responses.ToProductResponse(product))
	}

	return c.JSON(responses.Response{
		Status:  "success",
		Message: "Products retrieved successfully",
		Data:    productResponses,
	})
}

// SearchProducts handles searching products
func (h *InventoryHandler) SearchProducts(c *fiber.Ctx) error {
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

	products, err := h.usecase.Product.SearchProducts(c.Context(), query, limit, offset)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(responses.Response{
			Status:  "error",
			Message: "Failed to search products",
			Error:   err.Error(),
		})
	}

	var productResponses []responses.ProductResponse
	for _, product := range products {
		productResponses = append(productResponses, *responses.ToProductResponse(product))
	}

	return c.JSON(responses.Response{
		Status:  "success",
		Message: "Products search completed successfully",
		Data:    productResponses,
	})
}

// UpdateProductStock handles updating product stock
func (h *InventoryHandler) UpdateProductStock(c *fiber.Ctx) error {
	idParam := c.Params("id")
	id, err := strconv.ParseUint(idParam, 10, 32)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(responses.Response{
			Status:  "error",
			Message: "Invalid product ID",
			Error:   err.Error(),
		})
	}

	var req struct {
		Quantity int `json:"quantity" validate:"required"`
	}
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(responses.Response{
			Status:  "error",
			Message: "Invalid request body",
			Error:   err.Error(),
		})
	}

	err = h.usecase.Product.UpdateProductStock(c.Context(), uint(id), req.Quantity)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(responses.Response{
			Status:  "error",
			Message: "Failed to update product stock",
			Error:   err.Error(),
		})
	}

	return c.JSON(responses.Response{
		Status:  "success",
		Message: "Product stock updated successfully",
	})
}

// GetLowStockProducts handles getting low stock products
func (h *InventoryHandler) GetLowStockProducts(c *fiber.Ctx) error {
	threshold := c.QueryInt("threshold", 5)

	products, err := h.usecase.Product.GetLowStockProducts(c.Context(), threshold)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(responses.Response{
			Status:  "error",
			Message: "Failed to retrieve low stock products",
			Error:   err.Error(),
		})
	}

	var productResponses []responses.ProductResponse
	for _, product := range products {
		productResponses = append(productResponses, *responses.ToProductResponse(product))
	}

	return c.JSON(responses.Response{
		Status:  "success",
		Message: "Low stock products retrieved successfully",
		Data:    productResponses,
	})
}

// Category handlers

// CreateCategory handles category creation
func (h *InventoryHandler) CreateCategory(c *fiber.Ctx) error {
	var req interfaces.CreateCategoryRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(responses.Response{
			Status:  "error",
			Message: "Invalid request body",
			Error:   err.Error(),
		})
	}

	category, err := h.usecase.Category.CreateCategory(c.Context(), req)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(responses.Response{
			Status:  "error",
			Message: "Failed to create category",
			Error:   err.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(responses.Response{
		Status:  "success",
		Message: "Category created successfully",
		Data:    responses.ToCategoryResponse(category),
	})
}

// GetCategory handles getting a single category
func (h *InventoryHandler) GetCategory(c *fiber.Ctx) error {
	idParam := c.Params("id")
	id, err := strconv.ParseUint(idParam, 10, 32)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(responses.Response{
			Status:  "error",
			Message: "Invalid category ID",
			Error:   err.Error(),
		})
	}

	category, err := h.usecase.Category.GetCategory(c.Context(), uint(id))
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(responses.Response{
			Status:  "error",
			Message: "Category not found",
			Error:   err.Error(),
		})
	}

	return c.JSON(responses.Response{
		Status:  "success",
		Message: "Category retrieved successfully",
		Data:    responses.ToCategoryResponse(category),
	})
}

// UpdateCategory handles category updates
func (h *InventoryHandler) UpdateCategory(c *fiber.Ctx) error {
	idParam := c.Params("id")
	id, err := strconv.ParseUint(idParam, 10, 32)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(responses.Response{
			Status:  "error",
			Message: "Invalid category ID",
			Error:   err.Error(),
		})
	}

	var req interfaces.UpdateCategoryRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(responses.Response{
			Status:  "error",
			Message: "Invalid request body",
			Error:   err.Error(),
		})
	}

	category, err := h.usecase.Category.UpdateCategory(c.Context(), uint(id), req)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(responses.Response{
			Status:  "error",
			Message: "Failed to update category",
			Error:   err.Error(),
		})
	}

	return c.JSON(responses.Response{
		Status:  "success",
		Message: "Category updated successfully",
		Data:    responses.ToCategoryResponse(category),
	})
}

// DeleteCategory handles category deletion
func (h *InventoryHandler) DeleteCategory(c *fiber.Ctx) error {
	idParam := c.Params("id")
	id, err := strconv.ParseUint(idParam, 10, 32)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(responses.Response{
			Status:  "error",
			Message: "Invalid category ID",
			Error:   err.Error(),
		})
	}

	err = h.usecase.Category.DeleteCategory(c.Context(), uint(id))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(responses.Response{
			Status:  "error",
			Message: "Failed to delete category",
			Error:   err.Error(),
		})
	}

	return c.JSON(responses.Response{
		Status:  "success",
		Message: "Category deleted successfully",
	})
}

// ListCategories handles listing categories with pagination
func (h *InventoryHandler) ListCategories(c *fiber.Ctx) error {
	page := c.QueryInt("page", 1)
	limit := c.QueryInt("limit", 10)
	offset := (page - 1) * limit

	categories, err := h.usecase.Category.ListCategories(c.Context(), limit, offset)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(responses.Response{
			Status:  "error",
			Message: "Failed to retrieve categories",
			Error:   err.Error(),
		})
	}

	var categoryResponses []responses.CategoryResponse
	for _, category := range categories {
		categoryResponses = append(categoryResponses, *responses.ToCategoryResponse(category))
	}

	return c.JSON(responses.Response{
		Status:  "success",
		Message: "Categories retrieved successfully",
		Data:    categoryResponses,
	})
}

// Supplier handlers

// CreateSupplier handles supplier creation
func (h *InventoryHandler) CreateSupplier(c *fiber.Ctx) error {
	var req interfaces.CreateSupplierRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(responses.Response{
			Status:  "error",
			Message: "Invalid request body",
			Error:   err.Error(),
		})
	}

	supplier, err := h.usecase.Supplier.CreateSupplier(c.Context(), req)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(responses.Response{
			Status:  "error",
			Message: "Failed to create supplier",
			Error:   err.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(responses.Response{
		Status:  "success",
		Message: "Supplier created successfully",
		Data:    responses.ToSupplierResponse(supplier),
	})
}

// GetSupplier handles getting a single supplier
func (h *InventoryHandler) GetSupplier(c *fiber.Ctx) error {
	idParam := c.Params("id")
	id, err := strconv.ParseUint(idParam, 10, 32)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(responses.Response{
			Status:  "error",
			Message: "Invalid supplier ID",
			Error:   err.Error(),
		})
	}

	supplier, err := h.usecase.Supplier.GetSupplier(c.Context(), uint(id))
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(responses.Response{
			Status:  "error",
			Message: "Supplier not found",
			Error:   err.Error(),
		})
	}

	return c.JSON(responses.Response{
		Status:  "success",
		Message: "Supplier retrieved successfully",
		Data:    responses.ToSupplierResponse(supplier),
	})
}

// UpdateSupplier handles supplier updates
func (h *InventoryHandler) UpdateSupplier(c *fiber.Ctx) error {
	idParam := c.Params("id")
	id, err := strconv.ParseUint(idParam, 10, 32)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(responses.Response{
			Status:  "error",
			Message: "Invalid supplier ID",
			Error:   err.Error(),
		})
	}

	var req interfaces.UpdateSupplierRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(responses.Response{
			Status:  "error",
			Message: "Invalid request body",
			Error:   err.Error(),
		})
	}

	supplier, err := h.usecase.Supplier.UpdateSupplier(c.Context(), uint(id), req)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(responses.Response{
			Status:  "error",
			Message: "Failed to update supplier",
			Error:   err.Error(),
		})
	}

	return c.JSON(responses.Response{
		Status:  "success",
		Message: "Supplier updated successfully",
		Data:    responses.ToSupplierResponse(supplier),
	})
}

// DeleteSupplier handles supplier deletion
func (h *InventoryHandler) DeleteSupplier(c *fiber.Ctx) error {
	idParam := c.Params("id")
	id, err := strconv.ParseUint(idParam, 10, 32)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(responses.Response{
			Status:  "error",
			Message: "Invalid supplier ID",
			Error:   err.Error(),
		})
	}

	err = h.usecase.Supplier.DeleteSupplier(c.Context(), uint(id))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(responses.Response{
			Status:  "error",
			Message: "Failed to delete supplier",
			Error:   err.Error(),
		})
	}

	return c.JSON(responses.Response{
		Status:  "success",
		Message: "Supplier deleted successfully",
	})
}

// ListSuppliers handles listing suppliers with pagination
func (h *InventoryHandler) ListSuppliers(c *fiber.Ctx) error {
	page := c.QueryInt("page", 1)
	limit := c.QueryInt("limit", 10)
	offset := (page - 1) * limit

	suppliers, err := h.usecase.Supplier.ListSuppliers(c.Context(), limit, offset)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(responses.Response{
			Status:  "error",
			Message: "Failed to retrieve suppliers",
			Error:   err.Error(),
		})
	}

	var supplierResponses []responses.SupplierResponse
	for _, supplier := range suppliers {
		supplierResponses = append(supplierResponses, *responses.ToSupplierResponse(supplier))
	}

	return c.JSON(responses.Response{
		Status:  "success",
		Message: "Suppliers retrieved successfully",
		Data:    supplierResponses,
	})
}

// SearchSuppliers handles searching suppliers
func (h *InventoryHandler) SearchSuppliers(c *fiber.Ctx) error {
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

	suppliers, err := h.usecase.Supplier.SearchSuppliers(c.Context(), query, limit, offset)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(responses.Response{
			Status:  "error",
			Message: "Failed to search suppliers",
			Error:   err.Error(),
		})
	}

	var supplierResponses []responses.SupplierResponse
	for _, supplier := range suppliers {
		supplierResponses = append(supplierResponses, *responses.ToSupplierResponse(supplier))
	}

	return c.JSON(responses.Response{
		Status:  "success",
		Message: "Suppliers search completed successfully",
		Data:    supplierResponses,
	})
}

// UnitType handlers

// CreateUnitType handles unit type creation
func (h *InventoryHandler) CreateUnitType(c *fiber.Ctx) error {
	var req interfaces.CreateUnitTypeRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(responses.Response{
			Status:  "error",
			Message: "Invalid request body",
			Error:   err.Error(),
		})
	}

	unitType, err := h.usecase.UnitType.CreateUnitType(c.Context(), req)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(responses.Response{
			Status:  "error",
			Message: "Failed to create unit type",
			Error:   err.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(responses.Response{
		Status:  "success",
		Message: "Unit type created successfully",
		Data:    responses.ToUnitTypeResponse(unitType),
	})
}

// GetUnitType handles getting a single unit type
func (h *InventoryHandler) GetUnitType(c *fiber.Ctx) error {
	idParam := c.Params("id")
	id, err := strconv.ParseUint(idParam, 10, 32)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(responses.Response{
			Status:  "error",
			Message: "Invalid unit type ID",
			Error:   err.Error(),
		})
	}

	unitType, err := h.usecase.UnitType.GetUnitType(c.Context(), uint(id))
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(responses.Response{
			Status:  "error",
			Message: "Unit type not found",
			Error:   err.Error(),
		})
	}

	return c.JSON(responses.Response{
		Status:  "success",
		Message: "Unit type retrieved successfully",
		Data:    responses.ToUnitTypeResponse(unitType),
	})
}

// UpdateUnitType handles unit type updates
func (h *InventoryHandler) UpdateUnitType(c *fiber.Ctx) error {
	idParam := c.Params("id")
	id, err := strconv.ParseUint(idParam, 10, 32)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(responses.Response{
			Status:  "error",
			Message: "Invalid unit type ID",
			Error:   err.Error(),
		})
	}

	var req interfaces.UpdateUnitTypeRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(responses.Response{
			Status:  "error",
			Message: "Invalid request body",
			Error:   err.Error(),
		})
	}

	unitType, err := h.usecase.UnitType.UpdateUnitType(c.Context(), uint(id), req)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(responses.Response{
			Status:  "error",
			Message: "Failed to update unit type",
			Error:   err.Error(),
		})
	}

	return c.JSON(responses.Response{
		Status:  "success",
		Message: "Unit type updated successfully",
		Data:    responses.ToUnitTypeResponse(unitType),
	})
}

// DeleteUnitType handles unit type deletion
func (h *InventoryHandler) DeleteUnitType(c *fiber.Ctx) error {
	idParam := c.Params("id")
	id, err := strconv.ParseUint(idParam, 10, 32)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(responses.Response{
			Status:  "error",
			Message: "Invalid unit type ID",
			Error:   err.Error(),
		})
	}

	err = h.usecase.UnitType.DeleteUnitType(c.Context(), uint(id))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(responses.Response{
			Status:  "error",
			Message: "Failed to delete unit type",
			Error:   err.Error(),
		})
	}

	return c.JSON(responses.Response{
		Status:  "success",
		Message: "Unit type deleted successfully",
	})
}

// ListUnitTypes handles listing unit types with pagination
func (h *InventoryHandler) ListUnitTypes(c *fiber.Ctx) error {
	page := c.QueryInt("page", 1)
	limit := c.QueryInt("limit", 10)
	offset := (page - 1) * limit

	unitTypes, err := h.usecase.UnitType.ListUnitTypes(c.Context(), limit, offset)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(responses.Response{
			Status:  "error",
			Message: "Failed to retrieve unit types",
			Error:   err.Error(),
		})
	}

	var unitTypeResponses []responses.UnitTypeResponse
	for _, unitType := range unitTypes {
		unitTypeResponses = append(unitTypeResponses, *responses.ToUnitTypeResponse(unitType))
	}

	return c.JSON(responses.Response{
		Status:  "success",
		Message: "Unit types retrieved successfully",
		Data:    unitTypeResponses,
	})
}