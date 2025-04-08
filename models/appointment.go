package models

import "time"

type Appointment struct {
	ID              int       `db:"id" json:"id"`
	AppointmentDate time.Time `db:"Appointment_Date" json:"Appointment_Date"`
	DoctorID        int       `db:"Doctor_ID" json:"Doctor_ID"`
	PatientID       int       `db:"Patient_ID" json:"Patient_ID"`
	Status          string    `db:"Status" json:"Status"`
}
