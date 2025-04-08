package handlers

import (
	"doctor-on-demand/models"
	repository "doctor-on-demand/repositories"
	"net/http"

	"github.com/labstack/echo"
	"github.com/sirupsen/logrus"
)

type IAppointement interface {
	Create() echo.HandlerFunc
	GetById() echo.HandlerFunc
	Delete() echo.HandlerFunc
	Update() echo.HandlerFunc
}

type AppointementHandler struct {
	appointement models.Appointment
	repo         repository.AppointmentRepository
}

func NewAppointementHandler(repo repository.AppointmentRepository) *AppointementHandler {
	return &AppointementHandler{repo: repo}
}

func (r *AppointementHandler) Create() echo.HandlerFunc {
	return func(c echo.Context) error {
		var appointement models.Appointment
		if err := c.Bind(&appointement); err != nil {
			return c.JSON(http.StatusBadRequest, err)
		}
		result, err := r.repo.Create(c.Request().Context(), appointement)
		if err != nil {
			logrus.Info("Appointement creation failed", err)
		}
		return c.JSON(http.StatusOK, result)
	}
}

func (r *AppointementHandler) GetById() echo.HandlerFunc {
	return func(c echo.Context) error {
		return nil
	}
}
func (r *AppointementHandler) Delete() echo.HandlerFunc {
	return func(c echo.Context) error {
		return nil
	}
}
func (r *AppointementHandler) Update() echo.HandlerFunc {
	return func(c echo.Context) error {
		return nil
	}
}
