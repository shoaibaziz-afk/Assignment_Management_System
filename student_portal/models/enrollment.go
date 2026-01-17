package models

import "gorm.io/gorm"

// Enrollment links a student to a course offering
type Enrollment struct {
	gorm.Model
	StudentID        uint
	CourseOfferingID uint
}
