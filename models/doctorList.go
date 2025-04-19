package models

import (
	"time"

	"gorm.io/gorm"
)

type DoctorList struct {
	gorm.Model
	Name           string    `gorm:"type=varchar(100);not null" json:"name"`
	Specialization string    `gorm:"type=varchar(100);not null" json:"specialization"`
	Phone          string    `gorm:"type=varchar(20);not null;unique" json:"phone"`
	Email          string    `gorm:"type=varchar(100);not null;unique" json:"email"`
	Image          string    `gorm:"type=text" json:"image"`
	Availability   string    `gorm:"type=varchar(50)" json:"availability"`
	Fee            int       `gorm:"type=int;not null" json:"fee"`
	Schedule       time.Time `gorm:"type=timestamp without time zone" json:"schedule"`
}

// TableName to map model to "doctors" table
func (DoctorList) TableName() string {
	return "doctors"
}
