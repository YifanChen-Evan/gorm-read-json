package database

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// global variable
var DB *gorm.DB

// create a function to connect gorm with database
func ConnectDatabase() error {
	dsn := "host=localhost user=postgres dbname=postgres port=5432 sslmode=disable" // dsn: data source name
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil { // fail
		return fmt.Errorf("failed to connect database: %w", err)
	}

	DB = db
	fmt.Println("database connection successful.")
	return nil // success
}
