package handlers

import (
	"doctor-on-demand/models"
	repository "doctor-on-demand/repositories"
	"net/http"

	"github.com/labstack/echo"
	"github.com/sirupsen/logrus"
)

type Idoctor interface {
	Create() echo.HandlerFunc
	Get() echo.HandlerFunc
	Delete() echo.HandlerFunc
	Update() echo.HandlerFunc
}

type DoctorHandler struct {
	doctor models.DoctorList
	repo   repository.IDoctorRepository
}

func NewDoctorHandler(repo repository.IDoctorRepository) *DoctorHandler {
	return &DoctorHandler{repo: repo}
}

func (d *DoctorHandler) Create() echo.HandlerFunc {
	return func(c echo.Context) error {
		req := models.DoctorList{}
		if err := c.Bind(&req); err != nil {
			return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request payload"})
		}

		err := d.repo.CreateDoctor(c.Request().Context(), req) // ✅ सिर्फ error को ही assign करो
		if err != nil {
			logrus.Error("Failed to create doctor: ", err)
			return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to create doctor"})
		}

		return c.JSON(http.StatusCreated, map[string]string{"message": "Doctor created successfully"}) // ✅ Success Response
	}
}
