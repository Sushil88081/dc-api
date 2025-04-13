package models

import (
	"time"
)

type DoctorList struct {
	ID             string    `gorm:"primaryKey;autoIncrement" json:"id"`
	Name           string    `gorm:"type=varchar(100);not null" json:"name"`
	Specialization string    `gorm:"type=varchar(100);not null" json:"specialization"`
	Phone          string    `gorm:"type=varchar(20);not null;unique" json:"phone"`
	Email          string    `gorm:"type=varchar(100);not null;unique" json:"email"`
	Image          string    `gorm:"type=text" json:"image"`
	Availability   string    `gorm:"type=varchar(50)" json:"availability"`
	Fee            int       `gorm:"type=int;not null" json:"fee"`
	Schedule       time.Time `gorm:"type=timestamp without time zone" json:"schedule"`
	CreatedAt      time.Time `gorm:"autoCreateTime" json:"created_at"`
}

// TableName to map model to "doctors" table
func (DoctorList) TableName() string {
	return "doctors"
}
