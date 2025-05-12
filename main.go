package main

import (
	"github.com/fitnis/patient-service/handlers"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	p := router.Group("/patients")
	{
		handlers.RegisterPatientRoutes(p)
	}

	router.Run(":8082")
}
