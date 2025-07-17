package db

import (
	"boilerplate/internal/models"
	"log"

	"gorm.io/gorm"
)

// RunMigrations runs all model migrations using GORM AutoMigrate
func RunMigrations(db *gorm.DB) error {
	log.Println("Starting database migrations...")

	// Get all models to migrate
	modelsToMigrate := models.GetAllModels()

	// Run AutoMigrate for all models
	err := db.AutoMigrate(modelsToMigrate...)
	if err != nil {
		log.Printf("Error running migrations: %v", err)
		return err
	}

	log.Println("Database migrations completed successfully")
	return nil
}

// DropAllTables drops all tables (useful for testing)
func DropAllTables(db *gorm.DB) error {
	log.Println("Dropping all tables...")

	// Get all models to drop
	modelsToMigrate := models.GetAllModels()

	// Drop tables in reverse order
	for i := len(modelsToMigrate) - 1; i >= 0; i-- {
		err := db.Migrator().DropTable(modelsToMigrate[i])
		if err != nil {
			log.Printf("Error dropping table: %v", err)
			// Continue with other tables even if one fails
		}
	}

	log.Println("All tables dropped successfully")
	return nil
}