package handlers

import (
	"net/http"

	"professor_portal/db"
	"professor_portal/models"

	"github.com/gin-gonic/gin"
)

/*
AddAssignment creates an assignment under a course.
*/
func AddAssignment(c *gin.Context) {
	var assignment models.Assignment
	c.BindJSON(&assignment)

	db.DB.Create(&assignment)

	c.JSON(http.StatusOK, gin.H{
		"message": "Assignment created successfully",
	})
}
