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

// ShowDashboard displays all courses for the professor
func ShowDashboard(c *gin.Context) {
	profIDStr, err := c.Cookie("professor_id")
	if err != nil {
		c.Redirect(http.StatusFound, "/ui/login")
		return
	}

	profID, _ := strconv.Atoi(profIDStr)

	var courses []models.Course
	db.DB.Where("professor_id = ?", profID).Find(&courses)

	c.HTML(http.StatusOK, "dashboard.html", gin.H{
		"Courses": courses,
	})
}

func ShowAddCourse(c *gin.Context) {
	c.HTML(http.StatusOK, "add_course.html", nil)
}

func ShowAddAssignment(c *gin.Context) {
	c.HTML(http.StatusOK, "add_assignment.html", nil)
}

// ShowCourseDetail displays a single course and its assignments
func ShowCourseDetail(c *gin.Context) {
	courseID, _ := strconv.Atoi(c.Param("id"))

	var course models.Course
	var assignments []models.Assignment
	var professor models.Professor

	db.DB.First(&course, courseID)
	db.DB.First(&professor, course.ProfessorID)
	db.DB.Where("course_id = ?", courseID).Find(&assignments)

	c.HTML(http.StatusOK, "course.html", gin.H{
		"Course":      course,
		"Assignments": assignments,
		"Professor":   professor,
	})
}

/*
Handle form submissions
*/

func HandleSignup(c *gin.Context) {
	prof := models.Professor{
		Name:       c.PostForm(("name")),
		Email:      c.PostForm("email"),
		Password:   c.PostForm("password"),
		Department: c.PostForm("department"),
	}
	db.DB.Create(&prof)
	c.Redirect(http.StatusFound, "/ui/login")
}

func HandleLogin(c *gin.Context) {
	email := c.PostForm("email")

	var prof models.Professor
	err := db.DB.Where("email = ?", email).First(&prof).Error
	if err != nil {
		c.String(http.StatusUnauthorized, "Invalid credentials")
		return
	}

	// ⚠️ Password check skipped for MVP

	// Store professor ID in cookie (session-like)
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

func HandleAddCourse(c *gin.Context) {
	profIDStr, err := c.Cookie("professor_id")
	if err != nil {
		c.Redirect(http.StatusFound, "/ui/login")
		return
	}

	profID, _ := strconv.Atoi(profIDStr)

	course := models.Course{
		Name:        c.PostForm("name"),
		Semester:    c.PostForm("semester"),
		Department:  c.PostForm("department"),
		ProfessorID: uint(profID),
	}

	db.DB.Create(&course)
	c.Redirect(http.StatusFound, "/ui/dashboard")
}

func HandleAddAssignment(c *gin.Context) {
	courseID, _ := strconv.Atoi(c.PostForm("course_id"))

	assignment := models.Assignment{
		Title:        c.PostForm("title"),
		Deadline:     c.PostForm("deadline"),
		TimeAllotted: atoi(c.PostForm("time_allotted")),
		AIUsageLimit: atoi(c.PostForm("ai_limit")),
		CourseID:     uint(courseID),
	}

	db.DB.Create(&assignment)

	// Redirect back to course page
	c.Redirect(http.StatusFound, "/ui/course/"+strconv.Itoa(courseID))
}

func atoi(s string) int {
	i, _ := strconv.Atoi(s)
	return i
}
