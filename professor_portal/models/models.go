package models

import "gorm.io/gorm"

// --------------------
// Professor
// --------------------
type Professor struct {
	gorm.Model
	Name       string
	Email      string `gorm:"unique"`
	Password   string
	Department string
}

// --------------------
// Course (catalog-level)
// --------------------
type Course struct {
	gorm.Model
	Title              string
	Description        string
	OwningDepartment   string
	CourseType         string // MANDATORY | DEPARTMENT_ELECTIVE | OPEN_ELECTIVE
	CreatedByProfessor uint
}

// --------------------
// Course Offering (semester-specific)
// --------------------
type CourseOffering struct {
	gorm.Model
	CourseID     uint
	ProfessorID  uint
	AcademicTerm string // e.g. Fall 2026
	Semester     string // e.g. 6
}
