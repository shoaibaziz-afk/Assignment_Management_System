package main

import (
	"professor_portal/db"
	"professor_portal/handlers"
	"professor_portal/middleware"
	"professor_portal/models"

	"github.com/gin-gonic/gin"
)

func main() {
	// ---------------------------------------------
	// 1. Connect to database
	// ---------------------------------------------
	db.Connect()

	// ---------------------------------------------
	// 2. Auto-migrate database tables
	// (Creates tables if they do not exist)
	// ---------------------------------------------
	db.DB.AutoMigrate(
		&models.Professor{},
		&models.Course{},
		&models.Assignment{},
	)

	// ---------------------------------------------
	// 3. Initialize Gin router
	// ---------------------------------------------
	r := gin.Default()

	// ---------------------------------------------
	// 4. Load HTML templates for GUI
	// ---------------------------------------------
	r.LoadHTMLGlob("templates/*")

	// =============================================
	// ========== UI ROUTES (GUI) ===================
	// =============================================

	ui := r.Group("/ui")

	// Signup pages
	ui.GET("/signup", handlers.ShowSignup)
	ui.POST("/signup", handlers.HandleSignup)

	// Login pages
	ui.GET("/login", handlers.ShowLogin)
	ui.POST("/login", handlers.HandleLogin)

	// Dashboard
	ui.GET("/dashboard", handlers.ShowDashboard)

	// Course management
	ui.GET("/add-course", handlers.ShowAddCourse)
	ui.POST("/add-course", handlers.HandleAddCourse)

	// Assignment management
	ui.GET("/add-assignment", handlers.ShowAddAssignment)
	ui.POST("/add-assignment", handlers.HandleAddAssignment)

	// =============================================
	// ========== API ROUTES (JSON) =================
	// =============================================

	// Professor authentication (API)
	r.POST("/professor/signup", handlers.ProfessorSignup)
	r.POST("/professor/login", handlers.ProfessorLogin)

	// Protected professor API routes
	prof := r.Group("/professor")
	prof.Use(middleware.AuthMiddleware())

	prof.POST("/course", handlers.AddCourse)
	prof.GET("/courses", handlers.ListCourses)
	prof.POST("/assignment", handlers.AddAssignment)

	// ---------------------------------------------
	// 5. Start the HTTP server
	// ---------------------------------------------
	r.Run(":8000")
}
