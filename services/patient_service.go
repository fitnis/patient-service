package services

import (
	"github.com/fitnis/patient-service/models"
)

var patients []models.PatientRequest

func AdmitPatient(req models.PatientRequest) models.GenericResponse {
	patients = append(patients, req)
	return models.GenericResponse{Message: "Patient admitted"}
}

func GetAdmittedPatients() []models.PatientRequest {
	return patients
}

func DischargePatient(id string) {
	patients = []models.PatientRequest{}
}

func RegisterPatient(req models.PatientRequest) models.GenericResponse {
	return models.GenericResponse{Message: "Patient registered"}
}
