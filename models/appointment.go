package models

import (
	"time"

	"gorm.io/gorm"
)

const (
	StatusPending   = "pending"
	StatusConfirmed = "confirmed"
	StatusCancelled = "cancelled"
	StatusCompleted = "completed"
)

type Appointment struct {
	gorm.Model
	AppointmentID   uint      `gorm:"primaryKey;autoIncrement" json:"appointment_id"` // Changed to primary key
	PatientID       uint      `gorm:"not null" json:"patient_id"`
	DoctorID        uint      `gorm:"not null" json:"doctor_id"`
	AppointmentDate time.Time `gorm:"type:timestamp;not null" json:"appointment_date"` // Fixed type
	Status          string    `gorm:"type:varchar(20);not null" json:"status"`         // Increased size

	// Relationships
	Doctor  DoctorList `gorm:"foreignKey:DoctorID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"doctor"`
	Patient Patient    `gorm:"foreignKey:PatientID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"patient"`
}
