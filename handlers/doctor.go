package handlers

import (
	"doctor-on-demand/models"

	"github.com/labstack/echo"
)

type Idoctor interface {
	createDoctor(c echo.Context) error
	getDoctor(c echo.Context) error
	deleteDoctor(c echo.Context) error
	updateDoctor(c echo.Context) error
}

type DoctorHandler struct {
	doctor models.DoctorList
}

func NewDoctorHandler(doctor models.DoctorList) *DoctorHandler {
	return &DoctorHandler{
		doctor: doctor,
	}
}
