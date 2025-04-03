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
	query := `INSERT INTO patients (id,name,age,phone,email)
	       VALUES($1, $2, $3, $4, $5)`

	_, err := r.db.Exec(query, patient.ID, patient.Name, patient.Age, patient.Phone, patient.Email)
	if err != nil {
		logrus.Info("Error creating patient", err)

	}

	return err
}
func (r *PatientRepository) Update(ctx context.Context, id string, patient models.Patient) error {

	query := `UPDATE patients SET name = $2, age =$3, phone=$4, email=$5 WHERE id=$1`
	_, err := r.db.Exec(query, patient.ID, patient.Name, patient.Age, patient.Phone, patient.Email)
	if err != nil {
		logrus.Info("Error updating patient", err)

	}
	return err
}
func (r *PatientRepository) Delete(ctx context.Context, id string) error {

	query := `DELETE FROM patients WHERE id=$1`
	_, err := r.db.Exec(query, id)
	if err != nil {
		logrus.Info("error deleting patient", err)
	}
	return err
}
