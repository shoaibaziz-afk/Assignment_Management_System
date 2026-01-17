package models

import "gorm.io/gorm"

// Student represents a university student
type Student struct {
	gorm.Model
	Name            string
	Email           string `gorm:"unique"`
	Password        string
	Department      string
	CurrentSemester string
}
