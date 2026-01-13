package handlers

import (
	"net/http"
	"strconv"

	"professor_portal/db"
	"professor_portal/models"

	"github.com/gin-gonic/gin"
)

/*
Render HTML pages
*/

func ShowSignup(c *gin.Context) {
	c.HTML(http.StatusOK, "signup.html", nil)
}

func ShowLogin(c *gin.Context) {
	c.HTML(http.StatusOK, "login.html", nil)
}

func ShowDashboard(c *gin.Context) {
	c.HTML(http.StatusOK, "dashboard.html", nil)
}

func ShowAddCourse(c *gin.Context) {
	c.HTML(http.StatusOK, "add_course.html", nil)
}

func ShowAddAssignment(c *gin.Context) {
	c.HTML(http.StatusOK, "add_assignment.html", nil)
}

/*
Handle form submissions
*/

func HandleSignup(c *gin.Context) {
	prof := models.Professor{
		Email:      c.PostForm("email"),
		Password:   c.PostForm("password"),
		Department: c.PostForm("department"),
	}
	db.DB.Create(&prof)
	c.Redirect(http.StatusFound, "/ui/login")
}

func HandleLogin(c *gin.Context) {
	// For MVP, no password check yet
	c.Redirect(http.StatusFound, "/ui/dashboard")
}

func HandleAddCourse(c *gin.Context) {
	course := models.Course{
		Name:       c.PostForm("name"),
		Semester:   c.PostForm("semester"),
		Department: c.PostForm("department"),
	}
	db.DB.Create(&course)
	c.Redirect(http.StatusFound, "/ui/dashboard")
}

func HandleAddAssignment(c *gin.Context) {
	assignment := models.Assignment{
		Title:        c.PostForm("title"),
		Deadline:     c.PostForm("deadline"),
		TimeAllotted: atoi(c.PostForm("time_allotted")),
		AIUsageLimit: atoi(c.PostForm("ai_limit")),
		CourseID:     uint(atoi(c.PostForm("course_id"))),
	}
	db.DB.Create(&assignment)
	c.Redirect(http.StatusFound, "/ui/dashboard")
}

func atoi(s string) int {
	i, _ := strconv.Atoi(s)
	return i
}
