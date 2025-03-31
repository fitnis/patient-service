package models

type PatientRequest struct {
	Name   string `json:"name"`
	DOB    string `json:"dob"`
	Reason string `json:"reason"`
}

type GenericResponse struct {
	Message string `json:"message"`
}
