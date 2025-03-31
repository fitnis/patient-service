package handlers

import (
	"net/http"

	"github.com/fitnis/patient-service/models"
	"github.com/fitnis/patient-service/services"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// curl -X POST -H "Content-Type: application/json" -d '{"name":"John Doe","dob":"1990-01-01","reason":"Chest pain"}' http://localhost:8082/patients/admit
func AdmitPatient(c *gin.Context) {
	var req models.PatientRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid input"})
		return
	}
	id := uuid.New().String()
	c.JSON(http.StatusCreated, services.AdmitPatient(id, req))
}

// curl http://localhost:8082/patients/admit
func GetPatients(c *gin.Context) {
	c.JSON(http.StatusOK, services.GetAdmittedPatients())
}

// curl -X DELETE http://localhost:8082/patients/admit/{patientId}
func DischargePatient(c *gin.Context) {
	id := c.Param("patientId")
	services.DischargePatient(id)
	c.Status(http.StatusNoContent)
}

// curl -X POST -H "Content-Type: application/json" -d '{"name":"Jane Smith","dob":"1985-11-20","reason":"Checkup"}' http://localhost:8082/patients/register
func RegisterPatient(c *gin.Context) {
	var req models.PatientRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid input"})
		return
	}
	c.JSON(http.StatusCreated, services.RegisterPatient(req))
}
