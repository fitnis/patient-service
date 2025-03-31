package main

import (
	"github.com/fitnis/patient-service/handlers"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	p := router.Group("/patients")
	{
		p.POST("/admit", handlers.AdmitPatient)
		p.GET("/admit", handlers.GetPatients)
		p.DELETE("/admit/:patientId", handlers.DischargePatient)
		p.POST("/register", handlers.RegisterPatient)
	}

	router.Run(":8082")
}
