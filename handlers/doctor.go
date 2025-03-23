package handlers

import (
	"doctor-on-demand/models"
	"net/http"

	"github.com/labstack/echo"
	"github.com/sirupsen/logrus"
)

type Idoctor interface {
	CreateDoctor(c echo.Context) error
	GetDoctor(c echo.Context) error
	DeleteDoctor(c echo.Context) error
	UpdateDoctor(c echo.Context) error
}

type DoctorHandler struct {
	doctor models.DoctorList
}

func NewDoctorHandler(doctor models.DoctorList) *DoctorHandler {
	return &DoctorHandler{
		doctor: doctor,
	}
}

func (d *DoctorHandler) CreateDoctor(c echo.Context) error {
	var doctor models.DoctorList
	err := c.Bind(&doctor)

	if err != nil {
		logrus.Error("Failed to bind doctor data: ", err)
		return err
	}
	logrus.WithFields(logrus.Fields{
		"doctor": doctor,
	}).Info("Doctor created successfully")
	return c.JSON(http.StatusCreated, doctor)
}

func (d *DoctorHandler) GetDoctor(c echo.Context) error {
	return nil
}
func (d *DoctorHandler) DeleteDoctor(c echo.Context) error {
	return nil
}

func (d *DoctorHandler) UpdateDoctor(c echo.Context) error {
	return nil
}
