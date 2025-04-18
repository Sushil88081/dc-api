package repository

import (
	"context"
	"doctor-on-demand/models"
	"errors"
	"time"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

var (
	ErrScheduleNotAvailable = errors.New("doctor schedule is not available")
	ErrInvalidAppointment   = errors.New("invalid appointment details")
)

type IAppointmentRepository interface {
	BookAppointment(ctx context.Context, appointment models.Appointment) (models.Appointment, error)
	// GetAppointmentByID(ctx context.Context, id string) (models.Appointment, error)
	// GetAllAppointments(ctx context.Context, filter models.AppointmentFilter) ([]models.Appointment, error)
	// GetAppointmentsByDoctorID(ctx context.Context, doctorID string, date time.Time) ([]models.Appointment, error)
	// CancelAppointment(ctx context.Context, appointmentID string) error
	// UpdateAppointmentStatus(ctx context.Context, appointmentID string, status string) error
}

type AppointmentRepository struct {
	db *gorm.DB
}

func NewAppointmentRepository(db *gorm.DB) IAppointmentRepository {
	return &AppointmentRepository{
		db: db,
	}
}

// BookAppointment creates a new appointment after checking doctor's availability
func (r *AppointmentRepository) BookAppointment(ctx context.Context, appointment models.Appointment) (models.Appointment, error) {
	tx := r.db.WithContext(ctx).Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	// 1. Check if the schedule is available
	var schedule models.DoctorSchedule
	if err := tx.Where("id = ? AND is_available = ?", appointment.ScheduleID, true).First(&schedule).Error; err != nil {
		tx.Rollback()
		// Add more detailed logging
		if errors.Is(err, gorm.ErrRecordNotFound) {
			logrus.WithFields(logrus.Fields{
				"schedule_id": appointment.ScheduleID,
			}).Error("Doctor schedule is not available")
			return models.Appointment{}, ErrScheduleNotAvailable
		}
		return models.Appointment{}, err
	}

	// 2. Mark schedule as booked
	if err := tx.Model(&models.DoctorSchedule{}).
		Where("id = ?", appointment.ScheduleID).
		Update("is_available", false).Error; err != nil {
		tx.Rollback()
		return models.Appointment{}, err
	}

	// 3. Create the appointment
	appointment.Status = models.StatusConfirmed // or "booked"
	appointment.CreatedAt = time.Now().UTC()
	if err := tx.Create(&appointment).Error; err != nil {
		tx.Rollback()
		return models.Appointment{}, err
	}

	// 4. Commit transaction
	if err := tx.Commit().Error; err != nil {
		return models.Appointment{}, err
	}

	return appointment, nil
}

// // GetAppointmentByID fetches a single appointment by its ID
// func (r *AppointmentRepository) GetAppointmentByID(ctx context.Context, id string) (models.Appointment, error) {
// 	var appointment models.Appointment
// 	err := r.db.WithContext(ctx).
// 		Preload("Doctor").
// 		Preload("Patient").
// 		Preload("Schedule").
// 		Where("id = ?", id).
// 		First(&appointment).Error

// 	if err != nil {
// 		if errors.Is(err, gorm.ErrRecordNotFound) {
// 			return models.Appointment{}, nil
// 		}
// 		return models.Appointment{}, err
// 	}
// 	return appointment, nil
// }

// // GetAllAppointments fetches appointments with optional filters
// func (r *AppointmentRepository) GetAllAppointments(ctx context.Context, filter models.AppointmentFilter) ([]models.Appointment, error) {
// 	var appointments []models.Appointment

// 	query := r.db.WithContext(ctx).
// 		Preload("Doctor").
// 		Preload("Patient").
// 		Preload("Schedule")

// 	if filter.DoctorID != "" {
// 		query = query.Where("doctor_id = ?", filter.DoctorID)
// 	}
// 	if filter.PatientID != "" {
// 		query = query.Where("patient_id = ?", filter.PatientID)
// 	}
// 	if filter.Status != "" {
// 		query = query.Where("status = ?", filter.Status)
// 	}
// 	if !filter.StartDate.IsZero() {
// 		query = query.Where("created_at >= ?", filter.StartDate)
// 	}
// 	if !filter.EndDate.IsZero() {
// 		query = query.Where("created_at <= ?", filter.EndDate)
// 	}

// 	err := query.Find(&appointments).Error
// 	if err != nil {
// 		return nil, err
// 	}

// 	return appointments, nil
// }

// // GetAppointmentsByDoctorID fetches all appointments for a doctor on a specific date
// func (r *AppointmentRepository) GetAppointmentsByDoctorID(ctx context.Context, doctorID string, date time.Time) ([]models.Appointment, error) {
// 	var appointments []models.Appointment

// 	startOfDay := time.Date(date.Year(), date.Month(), date.Day(), 0, 0, 0, 0, date.Location())
// 	endOfDay := startOfDay.Add(24 * time.Hour)

// 	err := r.db.WithContext(ctx).
// 		Preload("Patient").
// 		Preload("Schedule").
// 		Where("doctor_id = ?", doctorID).
// 		Where("created_at BETWEEN ? AND ?", startOfDay, endOfDay).
// 		Find(&appointments).Error

// 	if err != nil {
// 		return nil, err
// 	}

// 	return appointments, nil
// }

// // CancelAppointment cancels an existing appointment
// func (r *AppointmentRepository) CancelAppointment(ctx context.Context, appointmentID string) error {
// 	tx := r.db.WithContext(ctx).Begin()
// 	defer func() {
// 		if r := recover(); r != nil {
// 			tx.Rollback()
// 		}
// 	}()

// 	// 1. Get the appointment to find the schedule ID
// 	var appointment models.Appointment
// 	if err := tx.Where("id = ?", appointmentID).First(&appointment).Error; err != nil {
// 		tx.Rollback()
// 		return err
// 	}

// 	// 2. Update appointment status
// 	if err := tx.Model(&models.Appointment{}).
// 		Where("id = ?", appointmentID).
// 		Update("status", "cancelled").Error; err != nil {
// 		tx.Rollback()
// 		return err
// 	}

// 	// 3. Mark schedule as available again
// 	if appointment.ScheduleID != "" {
// 		if err := tx.Model(&models.DoctorSchedule{}).
// 			Where("id = ?", appointment.ScheduleID).
// 			Update("is_available", true).Error; err != nil {
// 			tx.Rollback()
// 			return err
// 		}
// 	}

// 	return tx.Commit().Error
// }

// // UpdateAppointmentStatus updates the status of an appointment
// func (r *AppointmentRepository) UpdateAppointmentStatus(ctx context.Context, appointmentID string, status string) error {
// 	return r.db.WithContext(ctx).
// 		Model(&models.Appointment{}).
// 		Where("id = ?", appointmentID).
// 		Update("status", status).Error
// }
