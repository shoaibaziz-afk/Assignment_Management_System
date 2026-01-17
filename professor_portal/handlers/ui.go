package handlers

import (
	"net/http"
	"strconv"

	"professor_portal/db"
	"professor_portal/models"

	"github.com/gin-gonic/gin"
)

/*
========================
Render HTML pages
========================
*/

// Signup page
func ShowSignup(c *gin.Context) {
	c.HTML(http.StatusOK, "signup.html", nil)
}

// Login page
func ShowLogin(c *gin.Context) {
	c.HTML(http.StatusOK, "login.html", nil)
}

// Dashboard page (acts as a navigation hub)
func ShowDashboard(c *gin.Context) {
	profIDStr, err := c.Cookie("professor_id")
	if err != nil {
		c.Redirect(http.StatusFound, "/ui/login")
		return
	}

	profID, _ := strconv.Atoi(profIDStr)

	var professor models.Professor
	db.DB.First(&professor, profID)

	c.HTML(http.StatusOK, "dashboard.html", gin.H{
		"Professor": professor,
	})
}

/*
========================
Handle form submissions
========================
*/

// Handle professor signup
func HandleSignup(c *gin.Context) {
	prof := models.Professor{
		Name:       c.PostForm("name"),
		Email:      c.PostForm("email"),
		Password:   c.PostForm("password"), // password hashing later
		Department: c.PostForm("department"),
	}

	db.DB.Create(&prof)
	c.Redirect(http.StatusFound, "/ui/login")
}

// Handle professor login
func HandleLogin(c *gin.Context) {
	email := c.PostForm("email")

	var prof models.Professor
	err := db.DB.Where("email = ?", email).First(&prof).Error
	if err != nil {
		c.String(http.StatusUnauthorized, "Invalid credentials")
		return
	}

	// ⚠️ Password check intentionally skipped for MVP

	// Store professor ID in cookie
	c.SetCookie(
		"professor_id",
		strconv.Itoa(int(prof.ID)),
		3600, // 1 hour
		"/",
		"",
		false,
		true,
	)

	c.Redirect(http.StatusFound, "/ui/dashboard")
}
