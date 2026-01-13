package handlers

import (
	"net/http"

	"professor_portal/auth"
	"professor_portal/db"
	"professor_portal/models"

	"github.com/gin-gonic/gin"
)

/*
ProfessorSignup registers a new professor.
*/
func ProfessorSignup(c *gin.Context) {
	var prof models.Professor
	c.BindJSON(&prof)

	db.DB.Create(&prof)

	c.JSON(http.StatusOK, gin.H{
		"message": "Professor registered successfully",
	})
}

/*
ProfessorLogin logs in a professor and returns JWT.
*/
func ProfessorLogin(c *gin.Context) {
	var req models.Professor
	var prof models.Professor

	c.BindJSON(&req)
	db.DB.Where("email = ?", req.Email).First(&prof)

	token, _ := auth.GenerateToken(prof.Email)

	c.JSON(http.StatusOK, gin.H{
		"token": token,
	})
}
