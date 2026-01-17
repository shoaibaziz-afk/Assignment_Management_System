package db

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// Global DB object used across professor portal
var DB *gorm.DB

// Connect initializes the database connection
func Connect() {
	database, err := gorm.Open(
		sqlite.Open("/home/shoaibaziz/assignment_system/db/professor.db"),
		&gorm.Config{},
	)
	if err != nil {
		panic("Failed to connect to database")
	}
	DB = database
}
