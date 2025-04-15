package models

import (
	"time"

	"gorm.io/gorm"
)

type Appointment struct {
	gorm.Model
	AppointmentId   uint      `gorm:"not null" json:"appointment_id"`
	PatientID       uint      `gorm:"not null" json:"patient_id"`
	DoctorID        uint      `gorm:"not null" json:"doctor_id"`
	AppointmentDate time.Time ` gorm:"TIMPSTAMP"`
	Status          string    `gorm:"type:varchar(10);not null" json:"status"`

	//belongs to Doctor and patient table
	Doctor  DoctorList `gorm:"foreignKey:DoctorID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"doctor"`
	Patient Patient    `gorm:"foreignKey:PatientID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"patient"`
}
