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
	GetById() echo.HandlerFunc
	Delete() echo.HandlerFunc
	Update() echo.HandlerFunc
	GetAll() echo.HandlerFunc
	Count() echo.HandlerFunc
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

		err := d.repo.CreateDoctor(c.Request().Context(), &req)
		if err != nil {
			logrus.Error("Failed to create doctor: ", err)
			return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to create doctor"})
		}

		return c.JSON(http.StatusCreated, map[string]string{"message": "Doctor created successfully"})
	}
}

func (d *DoctorHandler) GetById() echo.HandlerFunc {
	return func(c echo.Context) error {
		id := c.Param("id")
		logrus.Info("request received for id", id)

		doctor, err := d.repo.GetByID(c.Request().Context(), id)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to get doctor"})
		} else if id == "" {
			return c.JSON(http.StatusBadRequest, map[string]string{"error": " id must be provided"})
		}
		return c.JSON(http.StatusOK, doctor)
	}
}

func (d *DoctorHandler) Update() echo.HandlerFunc {
	return func(c echo.Context) error {
		id := c.Param("id")
		logrus.Info("request received for id", id)
		req := models.DoctorList{}
		if err := c.Bind(&req); err != nil {
			return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request payload"})
		}

		err := d.repo.UpdateDoctor(c.Request().Context(), id, &req)
		if err != nil {
			logrus.Error("Failed to update doctor: ", err)
			return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to update doctor"})
		}
		return c.JSON(http.StatusOK, map[string]string{"message": "Doctor updated successfully"})
	}
}

func (d *DoctorHandler) Delete() echo.HandlerFunc {
	return func(c echo.Context) error {
		id := c.Param("id")
		logrus.Info("request received for id ", id)

		if id == "" {
			return c.JSON(http.StatusBadRequest, map[string]string{"error": "id must be provided"})
		}

		err := d.repo.DeleteDoctor(c.Request().Context(), id)
		if err != nil {
			logrus.WithError(err).Error("failed to delete doctor")
			return c.JSON(http.StatusInternalServerError, map[string]string{
				"error": "Failed to delete doctor",
			})
		}

		return c.JSON(http.StatusOK, map[string]string{
			"message": "Doctor deleted successfully",
			"id":      id,
		})
	}
}

func (d *DoctorHandler) GetAll() echo.HandlerFunc {
	return func(c echo.Context) error {
		result, err := d.repo.GetAll(c.Request().Context())
		if err != nil {
			logrus.Error("Failed to get doctors: ", err)
			return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to get doctors"})
		}
		return c.JSON(http.StatusOK, result)
	}
}

func (d *DoctorHandler) Count() echo.HandlerFunc {
	return func(c echo.Context) error {
		result, err := d.repo.Count(c.Request().Context())
		if err != nil {
			logrus.Info("Error getting the count of doctors", err)
			return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to count doctors"})
		}
		return c.JSON(http.StatusOK, result)
	}
}
