package db

import (
	"boilerplate/config"

	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
	_ "github.com/lib/pq"
)

type SQLXDatabase struct {
	DB     *sqlx.DB
	Logger *logrus.Logger
}

// NewSQLXConnection creates a new SQLx database connection for PostgreSQL
func NewSQLXConnection(conf *config.DatabaseAccount, appLogger *logrus.Logger) *SQLXDatabase {
	// Use DriverSource directly as DSN for PostgreSQL
	dsn := conf.DriverSource

	db, err := sqlx.Connect("postgres", dsn)
	if err != nil {
		appLogger.Fatalf("Failed to connect to database: %v", err)
	}

	// Set connection pool settings from config
	db.SetMaxOpenConns(conf.MaxOpenConns)
	db.SetMaxIdleConns(conf.MaxIdleConns)
	db.SetConnMaxLifetime(conf.ConnMaxLifetime)
	db.SetConnMaxIdleTime(conf.ConnMaxIdleTime)

	// Test the connection
	if err := db.Ping(); err != nil {
		appLogger.Fatalf("Failed to ping database: %v", err)
	}

	appLogger.Info("SQLx PostgreSQL database connected successfully")

	return &SQLXDatabase{
		DB:     db,
		Logger: appLogger,
	}
}

// Close closes the database connection
func (s *SQLXDatabase) Close() error {
	return s.DB.Close()
}

// GetDB returns the underlying sqlx.DB instance
func (s *SQLXDatabase) GetDB() *sqlx.DB {
	return s.DB
}