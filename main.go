package main

import (
	"github.com/gin-gonic/gin"
)

type Student struct {
	Name      string `json:"name"`
	StudentID string `json:"student_id"`
}

type Teacher struct {
	Name       string `json:"name"`
	EmployeeID string `json:"employee_id"`
}

func main() {
	r := gin.Default()

	students := []Student{
		{"Juan Pérez", "A11001"},
		{"María López", "A11002"},
		{"Carlos Ruiz", "A11003"},
		{"Ana Torres", "A11004"},
	}

	teachers := []Teacher{
		{"Roberto Gómez", "P11001"},
		{"Sandra Méndez", "P11002"},
		{"Luis Martínez", "P11003"},
		{"Fernanda Castro", "P11004"},
	}

	r.GET("/students", func(c *gin.Context) {
		c.JSON(200, students)
	})

	r.GET("/teachers", func(c *gin.Context) {
		c.JSON(200, teachers)
	})

	r.Run(":8080")
}
