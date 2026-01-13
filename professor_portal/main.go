package main

import (
	"professor_portal/db"
	"professor_portal/handlers"
	"professor_portal/middleware"
	"professor_portal/models"

	"github.com/gin-gonic/gin"
)

func main() {
	// Connect to DB
	db.Connect()

	// Auto-create tables
	db.DB.AutoMigrate(
		&models.Professor{},
		&models.Course{},
		&models.Assignment{},
	)

	// Initialize router
	r := gin.Default()

	// Public routes
	r.POST("/professor/signup", handlers.ProfessorSignup)
	r.POST("/professor/login", handlers.ProfessorLogin)

	// Protected professor routes
	prof := r.Group("/professor")
	prof.Use(middleware.AuthMiddleware())

	prof.POST("/course", handlers.AddCourse)
	prof.GET("/courses", handlers.ListCourses)
	prof.POST("/assignment", handlers.AddAssignment)

	// Start server
	r.Run(":8000")
}
