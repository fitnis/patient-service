package handlers

import (
	"net/http"

	"github.com/fitnis/patient-service/models"
	"github.com/fitnis/patient-service/services"

	"github.com/gin-gonic/gin"
)

func RegisterPatientRoutes(rg *gin.RouterGroup) {
	patients := rg.Group("/patients")
	patients.POST("/admit", admitPatient)
	patients.GET("/admit", getPatients)
	patients.DELETE("/admit/:patientId", dischargePatient)
	patients.POST("/register", registerPatient)
}

func admitPatient(c *gin.Context) {
	var req models.PatientRequest
	_ = c.ShouldBindJSON(&req)
	c.JSON(http.StatusCreated, services.AdmitPatient(req))
}

func getPatients(c *gin.Context) {
	c.JSON(http.StatusOK, services.GetAdmittedPatients())
}

func dischargePatient(c *gin.Context) {
	id := c.Param("patientId")
	services.DischargePatient(id)
	c.Status(http.StatusNoContent)
}

func registerPatient(c *gin.Context) {
	var req models.PatientRequest
	_ = c.ShouldBindJSON(&req)
	c.JSON(http.StatusCreated, services.RegisterPatient(req))
}
