package models

import "gorm.io/gorm"

// Student represents a student user
type Student struct {
	gorm.Model
	Name       string
	Email      string `gorm:"unique;not null"`
	Password   string
	Department string
	Semester   string
}

// Enrollment maps students to courses
type Enrollment struct {
	gorm.Model
	StudentID uint
	CourseID  uint
}
