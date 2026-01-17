package handlers

import (
	"net/http"
	"strconv"

	"student_portal/db"
	"student_portal/models"

	"github.com/gin-gonic/gin"
)

// Show signup page
func ShowSignup(c *gin.Context) {
	c.HTML(http.StatusOK, "signup.html", nil)
}

// Handle signup
func HandleSignup(c *gin.Context) {
	student := models.Student{
		Name:            c.PostForm("name"),
		Email:           c.PostForm("email"),
		Password:        c.PostForm("password"), // hashing later
		Department:      c.PostForm("department"),
		CurrentSemester: c.PostForm("semester"),
	}

	db.DB.Create(&student)
	c.Redirect(http.StatusFound, "/ui/login")
}

// Show login page
func ShowLogin(c *gin.Context) {
	c.HTML(http.StatusOK, "login.html", nil)
}

// Handle login
func HandleLogin(c *gin.Context) {
	email := c.PostForm("email")

	var student models.Student
	err := db.DB.Where("email = ?", email).First(&student).Error
	if err != nil {
		c.String(http.StatusUnauthorized, "Invalid credentials")
		return
	}

	c.SetCookie(
		"student_id",
		strconv.Itoa(int(student.ID)),
		3600,
		"/",
		"",
		false,
		true,
	)

	c.Redirect(http.StatusFound, "/ui/dashboard")
}
