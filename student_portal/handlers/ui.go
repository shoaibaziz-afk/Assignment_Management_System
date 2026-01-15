package handlers

import (
	"net/http"

	"student_portal/db"
	"student_portal/models"

	"github.com/gin-gonic/gin"
)

// Show student signup page (GET)
func ShowSignup(c *gin.Context) {
	c.HTML(http.StatusOK, "signup.html", nil)
}

// Handle student signup (POST)
func HandleSignup(c *gin.Context) {
	student := models.Student{
		Name:       c.PostForm("name"),
		Email:      c.PostForm("email"),
		Password:   c.PostForm("password"), // password hashing later
		Department: c.PostForm("department"),
		Semester:   c.PostForm("semester"),
	}

	db.DB.Create(&student)

	// After successful signup, redirect to login page
	c.Redirect(http.StatusFound, "/ui/login")
}

// Show student login page (GET)
func ShowLogin(c *gin.Context) {
	c.HTML(http.StatusOK, "login.html", nil)
}
