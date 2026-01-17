package main

import (
	"professor_portal/db"
	"professor_portal/handlers"
	"professor_portal/models"

	"github.com/gin-gonic/gin"
)

func main() {
	// --------------------------------------------------
	// 1. Connect to database
	// --------------------------------------------------
	db.Connect()

	// --------------------------------------------------
	// 2. Auto-migrate database tables
	// --------------------------------------------------
	db.DB.AutoMigrate(
		&models.Professor{},
		&models.Course{},
		&models.CourseOffering{},
	)

	// --------------------------------------------------
	// 3. Initialize Gin router
	// --------------------------------------------------
	r := gin.Default()

	// --------------------------------------------------
	// 4. Load professor templates
	// --------------------------------------------------
	r.LoadHTMLGlob("templates/*")

	// --------------------------------------------------
	// 5. UI routes (Professor Portal)
	// --------------------------------------------------
	ui := r.Group("/ui")

	// ---------- Authentication ----------
	ui.GET("/signup", handlers.ShowSignup)
	ui.POST("/signup", handlers.HandleSignup)

	ui.GET("/login", handlers.ShowLogin)
	ui.POST("/login", handlers.HandleLogin)

	// ---------- Dashboard ----------
	ui.GET("/dashboard", handlers.ShowDashboard)

	// ---------- Courses (ALL course logic lives in courses.go) ----------
	ui.GET("/courses", handlers.ListCourses)

	ui.GET("/courses/new", handlers.ShowAddCourse)
	ui.POST("/courses/new", handlers.HandleAddCourse)

	ui.GET("/courses/:id", handlers.ShowCourse)
	ui.POST("/courses/:id/offer", handlers.OfferCourse)

	// --------------------------------------------------
	// 6. Start server
	// --------------------------------------------------
	r.Run(":8000")
}
