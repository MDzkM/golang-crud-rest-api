package database

import (
	"golang-crud-rest-api/internal/product"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// Create database instance and variable to store errors
var Instance *gorm.DB
var err error

// Connect to mysql using GORM
func Connect(connectionString string) {
	Instance, err = gorm.Open(mysql.Open(connectionString), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
		panic("Cannot connect to DB")
	}
	log.Println("Connected to Database...")
}

// Ensure existing entities in backend has a corresponding table in the database
func Migrate(){	
	Instance.AutoMigrate(&entities.Product{})
	log.Println("Database Migration Completed...")
}