package main

import (
	"boilerplate/internal/models"
	"boilerplate/internal/repository"
	"boilerplate/internal/usecase"
	"boilerplate/internal/usecase/interfaces"
	"context"
	"fmt"
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	// Create in-memory SQLite database
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	// Run migrations for inventory models
	err = db.AutoMigrate(
		&models.Category{},
		&models.Supplier{},
		&models.UnitType{},
		&models.Product{},
		&models.ProductSerialNumber{},
	)
	if err != nil {
		log.Fatalf("Failed to migrate database: %v", err)
	}

	// Create repository and usecase managers
	repoManager := repository.NewRepositoryManager(db)
	usecaseManager := usecase.NewUsecaseManager(repoManager)

	ctx := context.Background()

	// Test Category Creation
	fmt.Println("Testing Category Creation...")
	categoryReq := interfaces.CreateCategoryRequest{
		Name:   "Spare Parts",
		Status: models.StatusAktif,
	}

	category, err := usecaseManager.Category.CreateCategory(ctx, categoryReq)
	if err != nil {
		log.Fatalf("Failed to create category: %v", err)
	}
	fmt.Printf("Category created: %+v\n", category)

	// Test Supplier Creation
	fmt.Println("\nTesting Supplier Creation...")
	supplierReq := interfaces.CreateSupplierRequest{
		SupplierName:      "PT Auto Parts Indonesia",
		ContactPersonName: "Budi Santoso",
		PhoneNumber:       "021-87654321",
		Address:           stringPtr("Jl. Industri No. 45, Jakarta"),
		Status:            models.StatusAktif,
	}

	supplier, err := usecaseManager.Supplier.CreateSupplier(ctx, supplierReq)
	if err != nil {
		log.Fatalf("Failed to create supplier: %v", err)
	}
	fmt.Printf("Supplier created: %+v\n", supplier)

	// Test UnitType Creation
	fmt.Println("\nTesting UnitType Creation...")
	unitTypeReq := interfaces.CreateUnitTypeRequest{
		Name:   "Pieces",
		Status: models.StatusAktif,
	}

	unitType, err := usecaseManager.UnitType.CreateUnitType(ctx, unitTypeReq)
	if err != nil {
		log.Fatalf("Failed to create unit type: %v", err)
	}
	fmt.Printf("UnitType created: %+v\n", unitType)

	// Test Product Creation
	fmt.Println("\nTesting Product Creation...")
	productReq := interfaces.CreateProductRequest{
		ProductName:        "Brake Pad Toyota Avanza",
		ProductDescription: stringPtr("High quality brake pad for Toyota Avanza"),
		CostPrice:          150000,
		SellingPrice:       200000,
		Stock:              25,
		SKU:                stringPtr("BP-TOY-AVZ-001"),
		Barcode:            stringPtr("1234567890123"),
		HasSerialNumber:    false,
		ShelfLocation:      stringPtr("A1-B2"),
		UsageStatus:        models.ProductUsageJual,
		IsActive:           true,
		CategoryID:         &category.CategoryID,
		SupplierID:         &supplier.SupplierID,
		UnitTypeID:         &unitType.UnitTypeID,
	}

	product, err := usecaseManager.Product.CreateProduct(ctx, productReq)
	if err != nil {
		log.Fatalf("Failed to create product: %v", err)
	}
	fmt.Printf("Product created: %+v\n", product)

	// Test Product Retrieval
	fmt.Println("\nTesting Product Retrieval...")
	retrievedProduct, err := usecaseManager.Product.GetProduct(ctx, product.ProductID)
	if err != nil {
		log.Fatalf("Failed to retrieve product: %v", err)
	}
	fmt.Printf("Product retrieved: %+v\n", retrievedProduct)

	// Test Product Search
	fmt.Println("\nTesting Product Search...")
	searchResults, err := usecaseManager.Product.SearchProducts(ctx, "brake", 10, 0)
	if err != nil {
		log.Fatalf("Failed to search products: %v", err)
	}
	fmt.Printf("Search results: %+v\n", searchResults)

	// Test Product Stock Update
	fmt.Println("\nTesting Product Stock Update...")
	err = usecaseManager.Product.UpdateProductStock(ctx, product.ProductID, -5)
	if err != nil {
		log.Fatalf("Failed to update product stock: %v", err)
	}
	fmt.Println("Product stock updated successfully")

	// Test Product by SKU
	fmt.Println("\nTesting Product by SKU...")
	productBySKU, err := usecaseManager.Product.GetProductBySKU(ctx, "BP-TOY-AVZ-001")
	if err != nil {
		log.Fatalf("Failed to get product by SKU: %v", err)
	}
	fmt.Printf("Product by SKU: %+v\n", productBySKU)

	// Test Low Stock Products
	fmt.Println("\nTesting Low Stock Products...")
	lowStockProducts, err := usecaseManager.Product.GetLowStockProducts(ctx, 30)
	if err != nil {
		log.Fatalf("Failed to get low stock products: %v", err)
	}
	fmt.Printf("Low stock products: %+v\n", lowStockProducts)

	// Test Products by Category
	fmt.Println("\nTesting Products by Category...")
	productsByCategory, err := usecaseManager.Product.GetProductsByCategory(ctx, category.CategoryID)
	if err != nil {
		log.Fatalf("Failed to get products by category: %v", err)
	}
	fmt.Printf("Products by category: %+v\n", productsByCategory)

	// Test Products by Supplier
	fmt.Println("\nTesting Products by Supplier...")
	productsBySupplier, err := usecaseManager.Product.GetProductsBySupplier(ctx, supplier.SupplierID)
	if err != nil {
		log.Fatalf("Failed to get products by supplier: %v", err)
	}
	fmt.Printf("Products by supplier: %+v\n", productsBySupplier)

	// Test Category List
	fmt.Println("\nTesting Category List...")
	categories, err := usecaseManager.Category.ListCategories(ctx, 10, 0)
	if err != nil {
		log.Fatalf("Failed to list categories: %v", err)
	}
	fmt.Printf("Categories list: %+v\n", categories)

	// Test Supplier List
	fmt.Println("\nTesting Supplier List...")
	suppliers, err := usecaseManager.Supplier.ListSuppliers(ctx, 10, 0)
	if err != nil {
		log.Fatalf("Failed to list suppliers: %v", err)
	}
	fmt.Printf("Suppliers list: %+v\n", suppliers)

	// Test UnitType List
	fmt.Println("\nTesting UnitType List...")
	unitTypes, err := usecaseManager.UnitType.ListUnitTypes(ctx, 10, 0)
	if err != nil {
		log.Fatalf("Failed to list unit types: %v", err)
	}
	fmt.Printf("UnitTypes list: %+v\n", unitTypes)

	fmt.Println("\nAll inventory tests completed successfully!")
}

func stringPtr(s string) *string {
	return &s
}