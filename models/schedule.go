package models

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"strings"
	"time"
)

type DoctorSchedule struct {
	ID          uint       `gorm:"primaryKey;autoIncrement" json:"id"`
	DoctorID    uint       `gorm:"not null" json:"doctor_id"`
	DayOfWeek   string     `gorm:"type:varchar(10);not null" json:"day"`
	IsAvailable bool       `gorm:"not null" json:"is_available"`
	Date        FlexDate   `gorm:"type:date;not null" json:"date"`
	StartTime   FlexTime   `gorm:"type:time;not null" json:"start_time"`
	EndTime     FlexTime   `gorm:"type:time;not null" json:"end_time"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at"`
	Doctor      DoctorList `gorm:"foreignKey:DoctorID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"doctor"`
}

// FlexTime handles time fields with database and JSON conversion
type FlexTime struct {
	time.Time
}

// UnmarshalJSON handles JSON input
func (ft *FlexTime) UnmarshalJSON(b []byte) error {
	var timeStr string
	if err := json.Unmarshal(b, &timeStr); err != nil {
		return err
	}

	if timeStr == "" {
		ft.Time = time.Time{}
		return nil
	}

	// Supported time formats
	formats := []string{
		"15:04:05",   // 24-hour with seconds
		"15:04",      // 24-hour without seconds
		"3:04 PM",    // 12-hour with AM/PM
		"3PM",        // Compact 12-hour
		time.RFC3339, // Full timestamp
	}

	timeStr = strings.TrimSpace(timeStr)

	for _, format := range formats {
		if t, err := time.Parse(format, timeStr); err == nil {
			ft.Time = t
			return nil
		}
	}

	return fmt.Errorf("invalid time format, use HH:MM:SS, HH:MM, or 3PM")
}

// MarshalJSON handles JSON output
func (ft FlexTime) MarshalJSON() ([]byte, error) {
	if ft.IsZero() {
		return json.Marshal(nil)
	}
	return json.Marshal(ft.Format("15:04:05"))
}

// Value converts to database value
func (ft FlexTime) Value() (driver.Value, error) {
	if ft.IsZero() {
		return nil, nil
	}
	return ft.Format("15:04:05"), nil
}

// Scan reads from database
func (ft *FlexTime) Scan(value interface{}) error {
	if value == nil {
		ft.Time = time.Time{}
		return nil
	}

	switch v := value.(type) {
	case time.Time:
		ft.Time = v
	case []byte:
		return ft.parseTime(string(v))
	case string:
		return ft.parseTime(v)
	default:
		return fmt.Errorf("unsupported type for FlexTime: %T", value)
	}
	return nil
}

func (ft *FlexTime) parseTime(value string) error {
	value = strings.TrimSpace(value)
	if value == "" {
		ft.Time = time.Time{}
		return nil
	}

	parsedTime, err := time.Parse("15:04:05", value)
	if err != nil {
		return err
	}
	ft.Time = parsedTime
	return nil
}

// FlexDate handles date fields with database and JSON conversion
type FlexDate struct {
	time.Time
}

// UnmarshalJSON handles JSON input
func (fd *FlexDate) UnmarshalJSON(b []byte) error {
	var dateStr string
	if err := json.Unmarshal(b, &dateStr); err != nil {
		return err
	}

	if dateStr == "" {
		fd.Time = time.Time{}
		return nil
	}

	// Supported date formats
	formats := []string{
		"2006-01-02",  // ISO format
		"02/01/2006",  // DD/MM/YYYY
		"01/02/2006",  // MM/DD/YYYY
		time.RFC3339,  // Full timestamp
		"Jan 2, 2006", // Text format
	}

	dateStr = strings.TrimSpace(dateStr)

	for _, format := range formats {
		if t, err := time.Parse(format, dateStr); err == nil {
			fd.Time = t
			return nil
		}
	}

	return fmt.Errorf("invalid date format, use YYYY-MM-DD or DD/MM/YYYY")
}

// MarshalJSON handles JSON output
func (fd FlexDate) MarshalJSON() ([]byte, error) {
	if fd.IsZero() {
		return json.Marshal(nil)
	}
	return json.Marshal(fd.Format("2006-01-02"))
}

// Value converts to database value
func (fd FlexDate) Value() (driver.Value, error) {
	if fd.IsZero() {
		return nil, nil
	}
	return fd.Format("2006-01-02"), nil
}

// Scan reads from database
func (fd *FlexDate) Scan(value interface{}) error {
	if value == nil {
		fd.Time = time.Time{}
		return nil
	}

	switch v := value.(type) {
	case time.Time:
		fd.Time = v
	case []byte:
		return fd.parseDate(string(v))
	case string:
		return fd.parseDate(v)
	default:
		return fmt.Errorf("unsupported type for FlexDate: %T", value)
	}
	return nil
}

func (fd *FlexDate) parseDate(value string) error {
	value = strings.TrimSpace(value)
	if value == "" {
		fd.Time = time.Time{}
		return nil
	}

	parsedDate, err := time.Parse("2006-01-02", value)
	if err != nil {
		return err
	}
	fd.Time = parsedDate
	return nil
}
