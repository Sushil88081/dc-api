package repository

import (
	"context"
	"doctor-on-demand/models"

	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
)

type IDoctorScheduleRepository interface {
	// GetByID(ctx context.Context, id string) (models.DoctorSchedule, error)
	Create(ctx context.Context, doctor_schedule models.DoctorSchedule) (models.DoctorSchedule, error)
	// Update(ctx context.Context, doctor_schedule models.DoctorSchedule) (models.DoctorSchedule, error)
	// Delete(ctx context.Context, id string) error
	// GetByDoctorId(ctx context.Context, doctor_id string) models.DoctorSchedule
}

type DoctorScheduleRepository struct {
	db *sqlx.DB
}

func NewDoctorScheduleRepository(db *sqlx.DB) *DoctorScheduleRepository {
	return &DoctorScheduleRepository{db: db}
}

func (ds *DoctorScheduleRepository) Create(ctx context.Context, doctor_schedule models.DoctorSchedule) (models.DoctorSchedule, error) {
	query := `INSERT INTO doctor_schedules(id,doctor_id,day_of_week,start_time,end_time,is_available,created_at) VALUES ($1, $2, $3, $4, $5, $6)
		RETURNING id`
	row, err := ds.db.ExecContext(ctx, query, doctor_schedule.ID, doctor_schedule.DoctorID, doctor_schedule.DayOfWeek, doctor_schedule.StartTime, doctor_schedule.EndTime, doctor_schedule.IsAvailable, doctor_schedule.CreatedAt)
	if err != nil {
		logrus.Info("failed in creating doctor schedule", row)

	}

	return models.DoctorSchedule{}, err
}
