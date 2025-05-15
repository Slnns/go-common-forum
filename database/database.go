package database

import (
	"database/sql"
	"fmt"
	"github.com/Slnns/go-common-forum/logger"
	"go.uber.org/zap"
	"os"

	_ "github.com/jackc/pgx/v5/stdlib" // PostgreSQL driver
)

// ConnectDB establishes a database connection using the provided URL.
func ConnectDB(databaseURL string) (*sql.DB, error) {
	db, err := sql.Open("pgx", databaseURL)
	if err != nil {
		logger.Log.Error("Failed to open database connection", zap.Error(err))
		return nil, fmt.Errorf("failed to open database: %w", err)
	}

	err = db.Ping()
	if err != nil {
		logger.Log.Error("Failed to ping database", zap.Error(err))
		return nil, fmt.Errorf("failed to ping database: %w", err)
	}

	logger.Log.Info("Successfully connected to database")
	return db, nil
}

// GetDatabaseURL retrieves the database URL from the environment.
// It returns a default URL if the environment variable is not set.
func GetDatabaseURL() string {
	url := os.Getenv("DATABASE_URL")
	if url == "" {
		url = "postgres://postgres:12345@localhost:5432/myforum?sslmode=disable" // Replace with your default
		logger.Log.Warn("DATABASE_URL environment variable not set, using default")
	}
	return url
}

// CloseDB closes the database connection.
func CloseDB(db *sql.DB) error {
	err := db.Close()
	if err != nil {
		logger.Log.Error("Failed to close database connection", zap.Error(err))
		return fmt.Errorf("failed to close database: %w", err)
	}
	logger.Log.Info("Database connection closed")
	return nil
}
