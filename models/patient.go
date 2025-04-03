package models

import "time"

type Patient struct {
	ID        int       `db:"id" json:"id"`
	Name      string    `db:"name" json:"name" `
	Age       int       `db:"age" json:"age"`
	Phone     int       `db:"phone" json:"phone"`
	Email     string    `db:"email" json:"email" `
	CreatedAt time.Time `db:"created_at" json:"created_at"`
}
