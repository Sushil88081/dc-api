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

		err := d.repo.CreateDoctor(c.Request().Context(), req)
		if err != nil {
			logrus.Error("Failed to create doctor: ", err)
			return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to create doctor"})
		}

		return c.JSON(http.StatusCreated, map[string]string{"message": "Doctor created successfully"}) // ✅ Success Response
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

		err := d.repo.UpdateDoctor(c.Request().Context(), id, req)
		if err != nil {
			return err
		}
		return c.JSON(http.StatusOK, map[string]string{"message": "Doctor updated successfully"}) // ✅ Success Response
	}
}

func (d *DoctorHandler) Delete() echo.HandlerFunc {
	return func(c echo.Context) error {
		id := c.Param("id")
		logrus.Info("request received for id", id)
		err := d.repo.DeleteDoctor(c.Request().Context(), id)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to delete doctor"})
		} else if id == "" {
			return c.JSON(http.StatusBadRequest, map[string]string{"error": "id must be provided"})
		}
		return c.JSON(http.StatusOK, map[string]string{"message": "Doctor deleted successfully"}) // ✅ Success Response
	}
}
