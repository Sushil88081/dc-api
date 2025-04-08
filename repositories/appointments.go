package repository

import (
	"context"
	"doctor-on-demand/models"
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
)

type IAppointmentsRepository interface {
	GetByID(ctx context.Context, id string) (models.Appointment, error)
	Create(ctx context.Context, appointments models.Appointment) (models.Appointment, error)
	Update(ctx context.Context, appointments models.Appointment) (models.Appointment, error)
	Delete(ctx context.Context, id string) error
}

type AppointmentRepository struct {
	db *sqlx.DB
}

func NewAppointmentRepository(db *sqlx.DB) *AppointmentRepository {
	return &AppointmentRepository{db: db}
}

func (r *AppointmentRepository) Create(ctx context.Context, appointment models.Appointment) (models.Appointment, error) {
	// 1. First verify the foreign keys exist
	if err := r.verifyDoctorExists(ctx, appointment.DoctorID); err != nil {
		return models.Appointment{}, fmt.Errorf("invalid doctor: %w", err)
	}

	if err := r.verifyPatientExists(ctx, appointment.PatientID); err != nil {
		return models.Appointment{}, fmt.Errorf("invalid patient: %w", err)
	}

	// 2. Insert with RETURNING to get complete record
	query := `
        INSERT INTO Appointments 
            (appointment_date, doctor_id, patient_id, status)
        VALUES 
            ($1, $2, $3, $4)
        RETURNING id, appointment_date, doctor_id, patient_id, status`

	var createdAppointment models.Appointment
	err := r.db.QueryRowContext(ctx, query,
		appointment.AppointmentDate,
		appointment.DoctorID,
		appointment.PatientID,
		appointment.Status,
	).Scan(
		&createdAppointment.ID,
		&createdAppointment.AppointmentDate,
		&createdAppointment.DoctorID,
		&createdAppointment.PatientID,
		&createdAppointment.Status,
	)

	if err != nil {
		logrus.WithContext(ctx).WithFields(logrus.Fields{
			"doctor_id":  appointment.DoctorID,
			"patient_id": appointment.PatientID,
			"error":      err,
		}).Error("failed to create appointment")
		return models.Appointment{}, fmt.Errorf("failed to create appointment: %w", err)
	}

	return createdAppointment, nil
}

// Helper functions to verify references
func (r *AppointmentRepository) verifyDoctorExists(ctx context.Context, doctorID int) error {
	var exists bool

	err := r.db.QueryRowContext(ctx, "SELECT EXISTS(SELECT * FROM doctors WHERE id = $1)", doctorID).Scan(&exists)
	if err != nil || !exists {
		return fmt.Errorf("doctor with ID %d not found", doctorID)
	}
	return nil
}

func (r *AppointmentRepository) verifyPatientExists(ctx context.Context, patientID int) error {
	var exists bool
	err := r.db.QueryRowContext(ctx, "SELECT EXISTS(SELECT * FROM patients WHERE id = $1)", patientID).Scan(&exists)
	if err != nil || !exists {
		return fmt.Errorf("patient with ID %d not found", patientID)
	}
	return nil
}

// func (r *AppointmentRepository) GetByID(ctx context.Context, id string) (models.Appointment,error){

// }
