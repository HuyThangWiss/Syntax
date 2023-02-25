package router

import (
	"P9/crud_gin/handlers"
	"github.com/gin-gonic/gin"
)

func Router() *gin.Engine {
	r := gin.Default()
	// Create Student
	r.POST("/students", handlers.CreateStudent)

	// Get Student
	r.GET("/students/:msv", handlers.GetStudent)

	// Get All Students
	r.GET("/students", handlers.GetAllStudents)

	// Update Student
	r.PUT("/students/:msv", handlers.UpdateStudent)

	// Delete Student
	r.DELETE("/students/:msv", handlers.DeleteStudent)
}
