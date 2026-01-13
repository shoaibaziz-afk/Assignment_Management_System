package handlers

import (
	"net/http"

	"professor_portal/db"
	"professor_portal/models"

	"github.com/gin-gonic/gin"
)

/*
AddCourse allows professor to create a new course.
*/
func AddCourse(c *gin.Context) {
	var course models.Course
	c.BindJSON(&course)

	db.DB.Create(&course)

	c.JSON(http.StatusOK, gin.H{
		"message": "Course added successfully",
	})
}

/*
ListCourses shows all courses taught by professor.
*/
func ListCourses(c *gin.Context) {
	var courses []models.Course
	db.DB.Find(&courses)

	c.JSON(http.StatusOK, courses)
}
