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
	GetAll(ctx context.Context) ([]models.DoctorList, error)
	Count(ctx context.Context) (int, error)
}

type DoctorRepository struct {
	db *sqlx.DB
}

func NewDoctorRepository(db *sqlx.DB) *DoctorRepository {
	return &DoctorRepository{db: db}
}

func (r *DoctorRepository) CreateDoctor(ctx context.Context, doctor models.DoctorList) error {
	query := `INSERT INTO doctors(name,specialization, phone, email,image,fee,availability,schedule) VALUES($1, $2, $3, $4,$5,$6,$7,$8)`
	_, err := r.db.Exec(query, doctor.Name, doctor.Specialization, doctor.Phone, doctor.Email, doctor.ImageUrl, doctor.Fee, doctor.Availability, doctor.Schedule)
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
	query := `UPDATE doctors SET name = $1, email = $2, specialization =$3, phone=$4 ,availability=$5,fee=$6,image=$7,Schedule=$8 WHERE id = $9`
	_, err := r.db.Exec(query, doctor.Name, doctor.Email, doctor.Specialization, doctor.Phone, doctor.Availability, doctor.Fee, doctor.ImageUrl, doctor.Schedule, id)
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

func (r *DoctorRepository) GetAll(ctx context.Context) ([]models.DoctorList, error) {
	var doctors []models.DoctorList
	query := `SELECT * FROM doctors;`
	err := r.db.SelectContext(ctx, &doctors, query)
	if err != nil {
		logrus.Info("Eror fetching the doctors from the database")
	}
	return doctors, err
}

func (r *DoctorRepository) Count(ctx context.Context) (int, error) {
	query := `SELECT COUNT(*) FROM doctors;`

	var count int
	err := r.db.QueryRowContext(ctx, query).Scan(&count)
	if err != nil {
		return 0, err
	}

	return count, nil
}
