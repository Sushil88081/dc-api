package models

import "time"

type DoctorList struct {
	ID             int       `db:"id" json:"id"`
	Name           string    `db:"name" json:"name"`
	Specialization string    `db:"specialization" json:"specialization"`
	Phone          string    `db:"phone" json:"phone"`
	Email          string    `db:"email" json:"email"`
	CreatedAt      time.Time `db:"created_at" json:"created_at"`
}
