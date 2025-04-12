package models

import "time"

type DoctorSchedule struct {
	ID          int       `json:"id" db:"id"`
	DoctorID    int       `json:"doctor_id" db:"doctor_id"`
	DayOfWeek   string    `json:"day_of_week" db:"day_of_week"` // e.g., "Monday"
	StartTime   time.Time `json:"start_time" db:"start_time"`   // only time part is relevant
	EndTime     time.Time `json:"end_time" db:"end_time"`       // only time part is relevant
	IsAvailable bool      `json:"is_available" db:"is_available"`
	CreatedAt   time.Time `json:"created_at" db:"created_at"`
	Slots       []Slot    // Last update time
}

type Slot struct {
	ID              int       `json:"id" db:"id"`                                      // Unique identifier for each slot
	StartTime       time.Time `json:"startTime" db:"start_time"`                       // Format: "HH:MM" (time of day)
	EndTime         time.Time `json:"endTime" db:"end_time"`                           // Format: "HH:MM" (time of day)
	IsBooked        bool      `json:"isBooked" db:"is_booked"`                         // If slot is taken
	PatientID       string    `json:"patientId,omitempty" db:"patient_id"`             // If booked, who booked it
	AppointmentType string    `json:"appointmentType,omitempty" db:"appointment_type"` // "checkup", "consultation", etc.
	CreatedAt       time.Time `json:"createdAt" db:"created_at"`                       // When the record was created
	UpdatedAt       time.Time `json:"updatedAt" db:"updated_at"`                       // Last update time
}

// Enum for days of week
const (
	Sunday = iota
	Monday
	Tuesday
	Wednesday
	Thursday
	Friday
	Saturday
)

// Example of appointment types
const (
	GeneralCheckup = "general_checkup"
	Consultation   = "consultation"
	FollowUp       = "follow_up"
	Emergency      = "emergency"
)
