package main

import (
	"student_portal/db"
	"student_portal/handlers"
	"student_portal/models"

	"github.com/gin-gonic/gin"
)

func main() {
	db.Connect()

	db.DB.AutoMigrate(
		&models.Student{},
		&models.Enrollment{},
		&models.AcademicTerm{},
	)

	r := gin.Default()
	r.LoadHTMLGlob("templates/*")

	ui := r.Group("/ui")

	ui.GET("/signup", handlers.ShowSignup)
	ui.POST("/signup", handlers.HandleSignup)

	ui.GET("/login", handlers.ShowLogin)
	ui.POST("/login", handlers.HandleLogin)

	ui.GET("/dashboard", handlers.ShowDashboard)

	r.Run(":9000")
}
