package handlers

import (
	"net/http"
	"strconv"

	"student_portal/db"
	"student_portal/models"

	"github.com/gin-gonic/gin"
)

// CourseView is a helper struct for rendering
type CourseView struct {
	Title        string
	Department   string
	CourseType   string
	AcademicTerm string
	Semester     string
}

// Show student dashboard
func ShowDashboard(c *gin.Context) {
	studentIDStr, err := c.Cookie("student_id")
	if err != nil {
		c.Redirect(http.StatusFound, "/ui/login")
		return
	}

	studentID, _ := strconv.Atoi(studentIDStr)

	var student models.Student
	db.DB.First(&student, studentID)

	// Get active academic term
	var term struct {
		Name string
	}
	db.DB.Table("academic_terms").
		Select("name").
		Where("is_active = ?", true).
		First(&term)

	// Query offerings for student's semester and active term
	rows, _ := db.DB.Raw(`
		SELECT 
			c.title,
			c.owning_department,
			c.course_type,
			o.academic_term,
			o.semester
		FROM course_offerings o
		JOIN courses c ON c.id = o.course_id
		WHERE 
			o.semester = ?
			AND o.academic_term = ?
	`, student.CurrentSemester, term.Name).Rows()

	defer rows.Close()

	var mandatory []CourseView
	var deptElectives []CourseView
	var openElectives []CourseView

	for rows.Next() {
		var cv CourseView
		rows.Scan(
			&cv.Title,
			&cv.Department,
			&cv.CourseType,
			&cv.AcademicTerm,
			&cv.Semester,
		)

		switch cv.CourseType {
		case "MANDATORY":
			if cv.Department == student.Department {
				mandatory = append(mandatory, cv)
			}
		case "DEPARTMENT_ELECTIVE":
			if cv.Department == student.Department {
				deptElectives = append(deptElectives, cv)
			}
		case "OPEN_ELECTIVE":
			openElectives = append(openElectives, cv)
		}
	}

	c.HTML(http.StatusOK, "dashboard.html", gin.H{
		"Student":       student,
		"Mandatory":     mandatory,
		"DeptElectives": deptElectives,
		"OpenElectives": openElectives,
	})
}
