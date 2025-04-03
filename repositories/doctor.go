package repository

import (
	"context"
	"doctor-on-demand/models"
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
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
	_, err := r.db.Exec(query, doctor.Name, doctor.Specialization, doctor.Phone, doctor.Email)
	return err
}

func (r *DoctorRepository) GetByID(ctx context.Context, id string) (models.DoctorList, error) {
	var doctor models.DoctorList
	query := `SELECT * FROM doctors WHERE id =$1;`
	err := r.db.GetContext(ctx, &doctor, query, id)
	if err != nil {
		logrus.Info("error getting doctor")
	} else if id == "" {
		logrus.Info("id must be entered", id)
	}
	return doctor, nil

}

func (r *DoctorRepository) UpdateDoctor(ctx context.Context, id string, doctor models.DoctorList) error {
	query := `UPDATE doctors SET name = $1, email = $2, specialization =$3, phone=$4 WHERE id = $5`
	_, err := r.db.Exec(query, doctor.Name, doctor.Email, doctor.Specialization, doctor.Phone, id)
	if err != nil {
		return fmt.Errorf("failed to update doctor: %w", err)

	}
	return nil
}

// ✅ अब `DeleteDoctor` को implement किया:
func (r *DoctorRepository) DeleteDoctor(ctx context.Context, id string) error {
	query := `DELETE FROM doctors WHERE id = $1;`
	_, err := r.db.Exec(query, id)
	return err

}
