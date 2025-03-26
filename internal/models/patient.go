package models

import (
	"time"

	"gorm.io/gorm"
)

// Patient represents a patient in the system
type Patient struct {
	ID        uint           `json:"id" gorm:"primaryKey"`
	Name      string         `json:"name" gorm:"size:255;not null"`
	Email     string         `json:"email" gorm:"size:255;not null;uniqueIndex"`
	Phone     string         `json:"phone" gorm:"size:255;not null"`
	Address   string         `json:"address" gorm:"size:500"`
	BirthDate time.Time      `json:"birth_date"`
	Gender    string         `json:"gender" gorm:"size:50"`
	CreatedAt time.Time      `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time      `json:"updated_at" gorm:"autoUpdateTime"`
	DeletedAt gorm.DeletedAt `json:"deleted_at" gorm:"index"`
}

// PatientRegistrationRequest represents the request body for patient registration
type PatientRegistrationRequest struct {
	Name      string    `json:"name" validate:"required"`
	Email     string    `json:"email" validate:"required,email"`
	Phone     string    `json:"phone" validate:"required"`
	Address   string    `json:"address"`
	BirthDate time.Time `json:"birth_date" validate:"required"`
	Gender    string    `json:"gender" validate:"required"`
}

// PatientUpdateRequest represents the request body for updating a patient
type PatientUpdateRequest struct {
	Name      string    `json:"name"`
	Email     string    `json:"email" validate:"email"`
	Phone     string    `json:"phone"`
	Address   string    `json:"address"`
	BirthDate time.Time `json:"birth_date"`
	Gender    string    `json:"gender"`
}
