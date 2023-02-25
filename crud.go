package main

import (
	"P9/crud_gin/student"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func main() {
	r := gin.Default()

	// Create Student
	r.POST("/students", func(c *gin.Context) {
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
	})

	// Get Student
	r.GET("/students/:msv", func(c *gin.Context) {
		msv := c.Param("msv")
		s, err := student.GetStudent(msv)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, s)
	})

	// Get All Students
	r.GET("/students", func(c *gin.Context) {
		students, err := student.GetAllStudents()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, students)
	})

	// Update Student
	r.PUT("/students/:msv", func(c *gin.Context) {
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
	})

	// Delete Student
	r.DELETE("/students/:msv", func(c *gin.Context) {
		msv := c.Param("msv")

		if err := student.DeleteStudent(msv); err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "Student deleted successfully"})
	})

	// Run server
	if err := r.Run(":8080"); err != nil {
		log.Fatalf("Failed to run server: %v", err)
	}
}
