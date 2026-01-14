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
	// 1. Connect to the database
	// ---------------------------------------------
	db.Connect()

	// ---------------------------------------------
	// 2. Auto-migrate tables (creates if not exist)
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
	// 4. Load HTML templates (GUI)
	// ---------------------------------------------
	r.LoadHTMLGlob("templates/*")

	// =============================================
	// ========== UI ROUTES (BROWSER) ===============
	// =============================================

	ui := r.Group("/ui")

	// ---------- Auth pages ----------
	ui.GET("/signup", handlers.ShowSignup)
	ui.POST("/signup", handlers.HandleSignup)

	ui.GET("/login", handlers.ShowLogin)
	ui.POST("/login", handlers.HandleLogin)

	// ---------- Dashboard ----------
	ui.GET("/dashboard", handlers.ShowDashboard)

	// ---------- Course pages ----------
	ui.GET("/add-course", handlers.ShowAddCourse)
	ui.POST("/add-course", handlers.HandleAddCourse)

	// 🔥 NEW ROUTE (THIS IS THE IMPORTANT ONE)
	// View a specific course and its assignments
	ui.GET("/course/:id", handlers.ShowCourseDetail)

	// ---------- Assignment creation ----------
	// Assignment is ALWAYS created inside a course
	ui.POST("/add-assignment", handlers.HandleAddAssignment)

	// =============================================
	// ========== API ROUTES (JSON) =================
	// =============================================

	// Professor API auth
	r.POST("/professor/signup", handlers.ProfessorSignup)
	r.POST("/professor/login", handlers.ProfessorLogin)

	// Protected professor APIs
	prof := r.Group("/professor")
	prof.Use(middleware.AuthMiddleware())

	prof.POST("/course", handlers.AddCourse)
	prof.GET("/courses", handlers.ListCourses)
	prof.POST("/assignment", handlers.AddAssignment)

	// ---------------------------------------------
	// 5. Start HTTP server
	// ---------------------------------------------
	r.Run(":8000")
}
