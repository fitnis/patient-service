package handlers

import (
	"net/http"
	"strconv"

	"github.com/fitnis/patient-service/internal/database"
	"github.com/fitnis/patient-service/internal/models"
	"github.com/labstack/echo/v4"
)

// PatientHandler handles patient operations
type PatientHandler struct{}

// NewPatientHandler creates a new PatientHandler
func NewPatientHandler() *PatientHandler {
	return &PatientHandler{}
}

// RegisterPatient handles patient registration
func (h *PatientHandler) RegisterPatient(c echo.Context) error {
	var req models.PatientRegistrationRequest
	if err := c.Bind(&req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	patient := models.Patient{
		Name:      req.Name,
		Email:     req.Email,
		Phone:     req.Phone,
		Address:   req.Address,
		BirthDate: req.BirthDate,
		Gender:    req.Gender,
	}

	result := database.GetDB().Create(&patient)
	if result.Error != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, result.Error.Error())
	}

	return c.JSON(http.StatusCreated, patient)
}

// GetPatients retrieves all patients
func (h *PatientHandler) GetPatients(c echo.Context) error {
	var patients []models.Patient
	result := database.GetDB().Find(&patients)
	if result.Error != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, result.Error.Error())
	}

	return c.JSON(http.StatusOK, patients)
}

// GetPatient retrieves a patient by ID
func (h *PatientHandler) GetPatient(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid patient ID")
	}

	var patient models.Patient
	result := database.GetDB().First(&patient, id)
	if result.Error != nil {
		return echo.NewHTTPError(http.StatusNotFound, "Patient not found")
	}

	return c.JSON(http.StatusOK, patient)
}

// UpdatePatient updates a patient
func (h *PatientHandler) UpdatePatient(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid patient ID")
	}

	var patient models.Patient
	result := database.GetDB().First(&patient, id)
	if result.Error != nil {
		return echo.NewHTTPError(http.StatusNotFound, "Patient not found")
	}

	var req models.PatientUpdateRequest
	if err := c.Bind(&req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	// Update patient data with non-empty fields from request
	if req.Name != "" {
		patient.Name = req.Name
	}
	if req.Email != "" {
		patient.Email = req.Email
	}
	if req.Phone != "" {
		patient.Phone = req.Phone
	}
	if req.Address != "" {
		patient.Address = req.Address
	}
	if !req.BirthDate.IsZero() {
		patient.BirthDate = req.BirthDate
	}
	if req.Gender != "" {
		patient.Gender = req.Gender
	}

	result = database.GetDB().Save(&patient)
	if result.Error != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, result.Error.Error())
	}

	return c.JSON(http.StatusOK, patient)
}

// DeletePatient deletes a patient
func (h *PatientHandler) DeletePatient(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid patient ID")
	}

	var patient models.Patient
	result := database.GetDB().First(&patient, id)
	if result.Error != nil {
		return echo.NewHTTPError(http.StatusNotFound, "Patient not found")
	}

	result = database.GetDB().Delete(&patient)
	if result.Error != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, result.Error.Error())
	}

	return c.NoContent(http.StatusNoContent)
}
