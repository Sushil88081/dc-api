package handlers

import (
	"doctor-on-demand/models"
	repository "doctor-on-demand/repositories"
	"net/http"
	"strings"

	"github.com/labstack/echo"
	"github.com/sirupsen/logrus"
)

type IAppointmentHandler interface {
	Create() echo.HandlerFunc
}

type AppointmentHandler struct {
	repo repository.IAppointmentRepository
}

func NewAppointmentHandler(repo repository.IAppointmentRepository) *AppointmentHandler {
	return &AppointmentHandler{
		repo: repo,
	}
}
func (h *AppointmentHandler) Create() echo.HandlerFunc {
	return func(c echo.Context) error {
		var appointment models.Appointment

		// Bind JSON to struct
		if err := c.Bind(&appointment); err != nil {
			logrus.WithField("error", err).Error("Failed to bind appointment")
			return c.JSON(http.StatusBadRequest, echo.Map{"error": "Invalid request payload"})
		}
		appointment.Schedule.DayOfWeek = strings.ToLower(appointment.AppointmentDate.Weekday().String())
		// Basic field validation
		if appointment.DoctorID == 0 || appointment.PatientID == 0 {
			logrus.Error("Doctor ID and Patient ID cannot be zero")
			return c.JSON(http.StatusBadRequest, echo.Map{"error": "Doctor ID and Patient ID are required"})
		}

		result, err := h.repo.BookAppointment(c.Request().Context(), appointment)
		if err != nil {
			logrus.WithField("error", err).Error("Failed to book appointment")

			// Custom error handling
			if err == repository.ErrScheduleNotAvailable {
				return c.JSON(http.StatusConflict, echo.Map{"error": "Selected time slot is not available"})
			}

			return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Could not book appointment"})
		}

		return c.JSON(http.StatusOK, result)
	}
}
