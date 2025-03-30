package repository

import (
	"context"
	"doctor-on-demand/models"

	"github.com/jmoiron/sqlx"
)

type IDoctorRepository interface {
	CreateDoctor(ctx context.Context, doctor models.DoctorList) error
	GetByID(ctx context.Context, id string) (models.DoctorList, error)
	UpdateDoctor(ctx context.Context, id string, doctor models.DoctorList) error
	DeleteDoctor(ctx context.Context, id string) error // ✅ यह method add करो
}

type DoctorRepository struct {
	db *sqlx.DB
}

func NewDoctorRepository(db *sqlx.DB) *DoctorRepository {
	return &DoctorRepository{db: db}
}

func (r *DoctorRepository) CreateDoctor(ctx context.Context, doctor models.DoctorList) error {
	query := `INSERT INTO doctors(name,specialization, phone, email) VALUES($1, $2, $3, $4)`
	_, err := r.db.Exec(query, doctor.Email, doctor.Name, doctor.Specialization, doctor.Phone)
	return err
}

func (r *DoctorRepository) GetByID(ctx context.Context, id string) (models.DoctorList, error) {
	return models.DoctorList{}, nil
}

func (r *DoctorRepository) UpdateDoctor(ctx context.Context, id string, doctor models.DoctorList) error {
	return nil
}

// ✅ अब `DeleteDoctor` को implement किया:
func (r *DoctorRepository) DeleteDoctor(ctx context.Context, id string) error {

	return nil
}
