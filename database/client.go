package database

import (
	"log"

	"rest-api-golang/entities"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var Instance *gorm.DB // Instance of the database
var err error

// Attemps to connect to the database with GORM and the connection string.
func Connect(connectionString string) {
	Instance, err = gorm.Open(mysql.Open(connectionString), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
		panic("Cannot connect to DB") // Stops exec of the current goroutine
	}
	log.Println("Connected to Database")
}

// Ensures that a table named documents is created on the database
func Migrate() {
	Instance.AutoMigrate(&entities.Document{})
	log.Println("Database migration completed...")
}
