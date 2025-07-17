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

	// Run migrations for foundation and customer models
	err = db.AutoMigrate(
		&models.User{},
		&models.Outlet{},
		&models.Customer{},
		&models.CustomerVehicle{},
	)
	if err != nil {
		log.Fatalf("Failed to migrate database: %v", err)
	}

	// Create repository and usecase managers
	repoManager := repository.NewRepositoryManager(db)
	usecaseManager := usecase.NewUsecaseManager(repoManager)

	ctx := context.Background()

	// Test Customer Creation
	fmt.Println("Testing Customer Creation...")
	customerReq := interfaces.CreateCustomerRequest{
		Name:        "John Doe",
		PhoneNumber: "081234567890",
		Address:     stringPtr("Jl. Merdeka No. 123"),
		Status:      models.StatusAktif,
	}

	customer, err := usecaseManager.Customer.CreateCustomer(ctx, customerReq)
	if err != nil {
		log.Fatalf("Failed to create customer: %v", err)
	}
	fmt.Printf("Customer created: %+v\n", customer)

	// Test Customer Retrieval
	fmt.Println("\nTesting Customer Retrieval...")
	retrievedCustomer, err := usecaseManager.Customer.GetCustomer(ctx, customer.CustomerID)
	if err != nil {
		log.Fatalf("Failed to retrieve customer: %v", err)
	}
	fmt.Printf("Customer retrieved: %+v\n", retrievedCustomer)

	// Test Customer Vehicle Creation
	fmt.Println("\nTesting Customer Vehicle Creation...")
	vehicleReq := interfaces.CreateCustomerVehicleRequest{
		CustomerID:     customer.CustomerID,
		PlateNumber:    "B1234XYZ",
		Brand:          "Toyota",
		Model:          "Avanza",
		Type:           "MPV",
		ProductionYear: 2020,
		ChassisNumber:  "CH1234567890123456",
		EngineNumber:   "ENG1234567890",
		Color:          "Silver",
		Notes:          stringPtr("Customer vehicle in good condition"),
	}

	vehicle, err := usecaseManager.CustomerVehicle.CreateCustomerVehicle(ctx, vehicleReq)
	if err != nil {
		log.Fatalf("Failed to create customer vehicle: %v", err)
	}
	fmt.Printf("Customer vehicle created: %+v\n", vehicle)

	// Test Customer Vehicle Retrieval
	fmt.Println("\nTesting Customer Vehicle Retrieval...")
	retrievedVehicle, err := usecaseManager.CustomerVehicle.GetCustomerVehicle(ctx, vehicle.VehicleID)
	if err != nil {
		log.Fatalf("Failed to retrieve customer vehicle: %v", err)
	}
	fmt.Printf("Customer vehicle retrieved: %+v\n", retrievedVehicle)

	// Test Customer Vehicles by Customer ID
	fmt.Println("\nTesting Customer Vehicles by Customer ID...")
	customerVehicles, err := usecaseManager.CustomerVehicle.GetCustomerVehiclesByCustomerID(ctx, customer.CustomerID)
	if err != nil {
		log.Fatalf("Failed to retrieve customer vehicles: %v", err)
	}
	fmt.Printf("Customer vehicles: %+v\n", customerVehicles)

	// Test Customer List
	fmt.Println("\nTesting Customer List...")
	customers, err := usecaseManager.Customer.ListCustomers(ctx, 10, 0)
	if err != nil {
		log.Fatalf("Failed to list customers: %v", err)
	}
	fmt.Printf("Customers list: %+v\n", customers)

	// Test Customer Search
	fmt.Println("\nTesting Customer Search...")
	searchResults, err := usecaseManager.Customer.SearchCustomers(ctx, "John", 10, 0)
	if err != nil {
		log.Fatalf("Failed to search customers: %v", err)
	}
	fmt.Printf("Search results: %+v\n", searchResults)

	fmt.Println("\nAll tests completed successfully!")
}

func stringPtr(s string) *string {
	return &s
}