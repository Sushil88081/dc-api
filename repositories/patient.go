package repository

import (
	"context"
	"doctor-on-demand/models"

	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
)

type IPatientRepository interface {
	GetById(ctx context.Context, id string) (models.Patient, error)
	Create(ctx context.Context, patient models.Patient) error
	Update(ctx context.Context, id string, patient models.Patient) error
	Delete(ctx context.Context, id string) error
}

type PatientRepository struct {
	db *sqlx.DB
}

func NewPatientRepository(db *sqlx.DB) *PatientRepository {
	return &PatientRepository{db: db}
}

func (r *PatientRepository) GetById(ctx context.Context, id string) (models.Patient, error) {
	query := `SELECT * FROM patients WHERE id=$1`
	var patient models.Patient
	err := r.db.Get(&patient, query, id)
	if err != nil {
		logrus.Info("Error getting patient", id)
	}
	return patient, err

}

func (r *PatientRepository) Create(ctx context.Context, patient models.Patient) error {
	return nil
}
func (r *PatientRepository) Update(ctx context.Context, id string, patient models.Patient) error {
	return nil
}
func (r *PatientRepository) Delete(ctx context.Context, id string) error {
	return nil
}
