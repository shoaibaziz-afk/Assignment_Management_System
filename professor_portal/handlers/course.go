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
COURSES — list & create
========================
*/

// List all courses created by the logged-in professor
func ListCourses(c *gin.Context) {
	profIDStr, err := c.Cookie("professor_id")
	if err != nil {
		c.Redirect(http.StatusFound, "/ui/login")
		return
	}

	profID, _ := strconv.Atoi(profIDStr)

	var courses []models.Course
	db.DB.Where("created_by_professor = ?", profID).Find(&courses)

	c.HTML(http.StatusOK, "courses.html", gin.H{
		"Courses": courses,
	})
}

// Show the "Add Course" form
func ShowAddCourse(c *gin.Context) {
	c.HTML(http.StatusOK, "add_course.html", nil)
}

// Handle course creation
func HandleAddCourse(c *gin.Context) {
	profIDStr, err := c.Cookie("professor_id")
	if err != nil {
		c.Redirect(http.StatusFound, "/ui/login")
		return
	}

	profID, _ := strconv.Atoi(profIDStr)

	course := models.Course{
		Title:              c.PostForm("title"),
		Description:        c.PostForm("description"),
		CourseType:         c.PostForm("course_type"),
		OwningDepartment:   c.PostForm("department"),
		CreatedByProfessor: uint(profID),
	}

	db.DB.Create(&course)
	c.Redirect(http.StatusFound, "/ui/courses")
}

/*
========================
COURSE DETAIL + OFFERING
========================
*/

// Show a single course and its offerings
func ShowCourse(c *gin.Context) {
	courseID, _ := strconv.Atoi(c.Param("id"))

	var course models.Course
	var offerings []models.CourseOffering

	db.DB.First(&course, courseID)
	db.DB.Where("course_id = ?", courseID).Find(&offerings)

	c.HTML(http.StatusOK, "course_detail.html", gin.H{
		"Course":    course,
		"Offerings": offerings,
	})
}

// Offer a course in a new semester (create CourseOffering)
func OfferCourse(c *gin.Context) {
	profIDStr, err := c.Cookie("professor_id")
	if err != nil {
		c.Redirect(http.StatusFound, "/ui/login")
		return
	}

	profID, _ := strconv.Atoi(profIDStr)
	courseID, _ := strconv.Atoi(c.Param("id"))

	offering := models.CourseOffering{
		CourseID:     uint(courseID),
		ProfessorID:  uint(profID),
		AcademicTerm: c.PostForm("academic_term"),
		Semester:     c.PostForm("semester"),
	}

	db.DB.Create(&offering)
	c.Redirect(http.StatusFound, "/ui/courses/"+strconv.Itoa(courseID))
}
