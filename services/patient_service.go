package services

import (
	"github.com/fitnis/patient-service/models"
)

var admittedPatients = make(map[string]models.PatientRequest)

func AdmitPatient(id string, req models.PatientRequest) models.GenericResponse {
	admittedPatients[id] = req
	return models.GenericResponse{Message: "Patient admitted"}
}

func GetAdmittedPatients() []models.PatientRequest {
	var list []models.PatientRequest
	for _, p := range admittedPatients {
		list = append(list, p)
	}
	return list
}

func DischargePatient(id string) {
	delete(admittedPatients, id)
}

func RegisterPatient(req models.PatientRequest) models.GenericResponse {
	// no storage needed for demo
	return models.GenericResponse{Message: "Patient registered"}
}
