package config

import (
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	db *gorm.DB // Global variable to store the database connection
)

// Connect establishes a connection to the MySQL database
func Connect() {
	// Get the DSN (Data Source Name) from environment variables
	dsn := os.Getenv("DB_DSN") // Change this to match the name in your .env file
	if dsn == "" {
		dsn = "user:pass@tcp(127.0.0.1:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local"
	}

	// Open a connection to the MySQL database using GORM
	d, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to database: " + err.Error())
	}

	// Assign the connection to the global variable
	db = d
}

// GetDB returns the database connection instance
func GetDB() *gorm.DB {
	return db
}
