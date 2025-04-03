package handlers

import (
	"doctor-on-demand/models"
	repository "doctor-on-demand/repositories"
	"net/http"

	"github.com/labstack/echo"
	"github.com/sirupsen/logrus"
)

type IPatient interface {
	GetById() echo.HandlerFunc
	Create() echo.HandlerFunc
	Update() echo.HandlerFunc
	Delete() echo.HandlerFunc
}

type PatientHandler struct {
	patient models.Patient
	repo    repository.IPatientRepository
}

func NewPatientHandler(repo repository.IPatientRepository) *PatientHandler {
	return &PatientHandler{
		repo: repo,
	}
}

func (p *PatientHandler) GetById() echo.HandlerFunc {
	return func(c echo.Context) error {
		id := c.Param("id")
		logrus.Info("request received for id: ", id)
		patient, err := p.repo.GetById(c.Request().Context(), id)
		if err != nil {
			logrus.Info("error getting patient", err)
		} else if id == "" {
			logrus.Info("please enter a id")
		}
		return c.JSON(http.StatusOK, patient)
	}
}
func (p *PatientHandler) Create() echo.HandlerFunc {
	return func(c echo.Context) error {
		return nil
	}
}
func (p *PatientHandler) Update() echo.HandlerFunc {
	return func(c echo.Context) error {
		return nil
	}
}

func (p *PatientHandler) Delete() echo.HandlerFunc {
	return func(c echo.Context) error {
		return nil
	}
}
