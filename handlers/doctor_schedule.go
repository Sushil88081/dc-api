package handlers

import (
	"doctor-on-demand/models"
	repository "doctor-on-demand/repositories"
	"net/http"
	"time"

	"github.com/labstack/echo"
	"github.com/sirupsen/logrus"
)

type IDoctorScheduleHandler interface {
	Create() echo.HandlerFunc
	// GetById() echo.HandlerFunc
	// GetByDoctorId() echo.HandlerFunc
	// Delete() echo.HandlerFunc
	// Update() echo.HandlerFunc
}

type DoctorScheduleHandler struct {
	repo repository.IDoctorScheduleRepository
}

func NewDoctorScheduleHandler(repo repository.IDoctorScheduleRepository) *DoctorScheduleHandler {
	return &DoctorScheduleHandler{
		repo: repo,
	}
}

func (ds *DoctorScheduleHandler) Create() echo.HandlerFunc {
	return func(c echo.Context) error {
		req := models.DoctorSchedule{}
		req.CreatedAt = time.Now()
		if err := c.Bind(&req); err != nil {
			return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request payload"})
		}

		result, err := ds.repo.Create(c.Request().Context(), req)
		if err != nil {
			logrus.Info("Error creating schedule", err)
		}
		return c.JSON(http.StatusOK, result)

	}

}
