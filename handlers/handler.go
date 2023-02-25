package handlers

import (
	"P9/crud_gin/student"
	"github.com/gin-gonic/gin"
	"net/http"
)

func CreateStudent(c *gin.Context) {
	var s student.Student
	if err := c.BindJSON(&s); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := student.CreateStudent(s); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Student created successfully"})
}

func GetStudent(c *gin.Context) {
	msv := c.Param("msv")
	s, err := student.GetStudent(msv)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, s)
}

func GetAllStudents(c *gin.Context) {
	students, err := student.GetAllStudents()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, students)
}

func UpdateStudent(c *gin.Context) {
	msv := c.Param("msv")
	var s student.Student
	if err := c.BindJSON(&s); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	s.Msv = msv

	if err := student.UpdateStudent(s); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Student updated successfully"})
}

func DeleteStudent(c *gin.Context) {
	msv := c.Param("msv")

	if err := student.DeleteStudent(msv); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Student deleted successfully"})
}
