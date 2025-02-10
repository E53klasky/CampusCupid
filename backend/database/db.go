package database

// Import necessary packages
import (
	"datingapp/models" // Import the models package for User model
	"fmt"              // For formatted I/O
	"os"               // For working with environment variables

	"gorm.io/driver/postgres" // GORM PostgreSQL driver
	"gorm.io/gorm"            // GORM ORM library
)

var DB *gorm.DB // Global variable to hold the database connection

// Connect function to establish a connection to the PostgreSQL database
func Connect() {
	// Build the database connection string (DSN) using environment variables
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", // Format the DSN for connecting to PostgreSQL
		os.Getenv("DB_HOST"),     // Get DB host from environment variable
		os.Getenv("DB_USER"),     // Get DB user from environment variable
		os.Getenv("DB_PASSWORD"), // Get DB password from environment variable
		os.Getenv("DB_NAME"),     // Get DB name from environment variable
		os.Getenv("DB_PORT"),     // Get DB port from environment variable
	)

	// Variable to capture potential error during database connection
	var err error
	// Attempt to open the database connection using GORM
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	// If connection fails, panic and display an error message
	if err != nil {
		panic("Failed to connect to database")
	}

	// Print a success message if connection is established
	fmt.Println("Database connection established")

	// Auto-migrate the User model to ensure the schema is up-to-date in the database
	// This will create the tables, columns, and constraints based on the User model
	if err := DB.AutoMigrate(&models.User{}); err != nil {
		// If migration fails, panic and display an error message
		panic("Failed to auto-migrate database")
	}

	// Print a success message if migration is completed successfully
	fmt.Println("Database migration completed")
}
