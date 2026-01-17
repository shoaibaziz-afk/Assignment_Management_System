package models

import "gorm.io/gorm"

// AcademicTerm represents a semester like "Fall 2026"
type AcademicTerm struct {
	gorm.Model
	Name     string
	IsActive bool
}
