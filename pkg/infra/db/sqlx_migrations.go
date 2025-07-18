package db

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
	"sort"
	"strings"

	"github.com/jmoiron/sqlx"
)

// RunSQLXMigrations runs SQL migration files using SQLx
func RunSQLXMigrations(db *sqlx.DB) error {
	// Create migrations table if it doesn't exist
	createMigrationsTable := `
		CREATE TABLE IF NOT EXISTS migrations (
			id SERIAL PRIMARY KEY,
			filename VARCHAR(255) UNIQUE NOT NULL,
			executed_at TIMESTAMP DEFAULT NOW()
		);
	`
	
	if _, err := db.Exec(createMigrationsTable); err != nil {
		return fmt.Errorf("failed to create migrations table: %v", err)
	}

	// Get migration files
	files, err := filepath.Glob("migrations/*.up.sql")
	if err != nil {
		return fmt.Errorf("failed to read migration files: %v", err)
	}

	// Sort files to ensure proper order
	sort.Strings(files)

	for _, file := range files {
		filename := filepath.Base(file)
		
		// Check if migration already executed
		var count int
		err := db.Get(&count, "SELECT COUNT(*) FROM migrations WHERE filename = $1", filename)
		if err != nil {
			return fmt.Errorf("failed to check migration status: %v", err)
		}

		if count > 0 {
			fmt.Printf("Migration %s already executed, skipping...\n", filename)
			continue
		}

		// Read and execute migration file
		content, err := ioutil.ReadFile(file)
		if err != nil {
			return fmt.Errorf("failed to read migration file %s: %v", filename, err)
		}

		// Split by semicolon and execute each statement
		statements := strings.Split(string(content), ";")
		
		tx, err := db.Beginx()
		if err != nil {
			return fmt.Errorf("failed to begin transaction: %v", err)
		}

		for _, stmt := range statements {
			stmt = strings.TrimSpace(stmt)
			if stmt == "" {
				continue
			}

			if _, err := tx.Exec(stmt); err != nil {
				tx.Rollback()
				return fmt.Errorf("failed to execute migration %s: %v", filename, err)
			}
		}

		// Record migration as executed
		if _, err := tx.Exec("INSERT INTO migrations (filename) VALUES ($1)", filename); err != nil {
			tx.Rollback()
			return fmt.Errorf("failed to record migration: %v", err)
		}

		if err := tx.Commit(); err != nil {
			return fmt.Errorf("failed to commit migration: %v", err)
		}

		fmt.Printf("Migration %s executed successfully\n", filename)
	}

	return nil
}