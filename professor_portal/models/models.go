package models

import "gorm.io/gorm"

/*
Professor represents a faculty member.
A professor can teach multiple courses.
*/
type Professor struct {
	gorm.Model
	Email      string `gorm:"unique;not null"`
	Password   string
	Department string
}

/*
Course represents a course taught by a professor.
*/
type Course struct {
	gorm.Model
	Name        string
	Semester    string
	Department  string
	ProfessorID uint
}

/*
Assignment belongs to a course.
Constraints are kept SIMPLE for MVP.
*/
type Assignment struct {
	gorm.Model
	Title        string
	Deadline     string
	TimeAllotted int // in minutes
	AIUsageLimit int
	CourseID     uint
}
