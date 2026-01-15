package db

import (
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	database, err := gorm.Open(
		sqlite.Open("../professor_portal/professor.db"),
		&gorm.Config{},
	)

	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	DB = database
}
