package main

import (
	"student_portal/db"
	"student_portal/handlers"
	"student_portal/models"

	"github.com/gin-gonic/gin"
)

func main() {
	// ---------------------------------------------
	// 1. Connect to shared database
	// ---------------------------------------------
	db.Connect()

	// ---------------------------------------------
	// 2. Auto-migrate student-related tables
	// ---------------------------------------------
	db.DB.AutoMigrate(
		&models.Student{},
		&models.Enrollment{},
	)

	// ---------------------------------------------
	// 3. Initialize Gin
	// ---------------------------------------------
	r := gin.Default()

	// ---------------------------------------------
	// 4. Load HTML templates
	// ---------------------------------------------
	r.LoadHTMLGlob("templates/*")

	// =============================================
	// ========== UI ROUTES =========================
	// =============================================

	ui := r.Group("/ui")
	ui.GET("/signup", handlers.ShowSignup)
	ui.POST("/signup", handlers.HandleSignup)
	ui.GET("/login", handlers.ShowLogin)

	// ---------------------------------------------
	// 5. Start server
	// ---------------------------------------------
	r.Run(":8001")
}
