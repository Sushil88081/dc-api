package handlers

import (
	"doctor-on-demand/models"
	repository "doctor-on-demand/repositories"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/labstack/echo"
	"github.com/sirupsen/logrus"
)

// DoctorScheduleHandler defines the interface for doctor schedule handlers
type IDoctorScheduleHandler interface {
	Create() echo.HandlerFunc
	GetByID() echo.HandlerFunc
	Update() echo.HandlerFunc
	Delete() echo.HandlerFunc
	GetByDoctorID() echo.HandlerFunc
}

// doctorScheduleHandler implements DoctorScheduleHandler
type DoctorScheduleHandler struct {
	repo repository.IDoctorScheduleRepository
}

// NewDoctorScheduleHandler creates a new handler instance
func NewDoctorScheduleHandler(repo repository.IDoctorScheduleRepository) *DoctorScheduleHandler {
	return &DoctorScheduleHandler{repo: repo}
}

func (h *DoctorScheduleHandler) Create() echo.HandlerFunc {
	return func(c echo.Context) error {
		req := models.DoctorSchedule{
			CreatedAt: time.Now().UTC(),
			UpdatedAt: time.Now().UTC(),
		}

		if err := c.Bind(&req); err != nil {
			logrus.WithFields(logrus.Fields{
				"error": err.Error(),
				"route": c.Path(),
			}).Error("Failed to bind request")
			return c.JSON(http.StatusBadRequest, map[string]interface{}{
				"error":   "Invalid request payload",
				"details": err.Error(),
			})
		}

		// Validate required fields
		if req.DoctorID == 0 {
			return c.JSON(http.StatusBadRequest, map[string]string{
				"error": "doctor_id is required",
			})
		}

		if req.DayOfWeek == "" {
			return c.JSON(http.StatusBadRequest, map[string]string{
				"error": "day_of_week is required",
			})
		}

		if req.StartTime.IsZero() || req.EndTime.IsZero() {
			return c.JSON(http.StatusBadRequest, map[string]string{
				"error": "both start_time and end_time are required",
			})
		}

		createdSchedule, err := h.repo.Create(c.Request().Context(), req)
		if err != nil {
			logrus.WithFields(logrus.Fields{
				"error":   err.Error(),
				"request": req,
				"route":   c.Path(),
			}).Error("Failed to create schedule")

			return c.JSON(http.StatusInternalServerError, map[string]string{
				"error": "Failed to create schedule",
			})
		}

		return c.JSON(http.StatusCreated, createdSchedule)
	}
}

func (h *DoctorScheduleHandler) GetByID() echo.HandlerFunc {
	return func(c echo.Context) error {
		idStr := c.Param("id")
		if idStr == "" {
			return c.JSON(http.StatusBadRequest, map[string]string{
				"error": "schedule id is required",
			})
		}

		id, err := strconv.ParseUint(idStr, 10, 64)
		if err != nil {
			return c.JSON(http.StatusBadRequest, map[string]string{
				"error": "invalid schedule id format",
			})
		}

		schedule, err := h.repo.GetByID(c.Request().Context(), uint(id))
		if err != nil {
			logrus.WithFields(logrus.Fields{
				"error":      err.Error(),
				"scheduleID": id,
				"route":      c.Path(),
			}).Error("Failed to get schedule")

			return c.JSON(http.StatusNotFound, map[string]string{
				"error": "Schedule not found",
			})
		}

		return c.JSON(http.StatusOK, schedule)
	}
}

// Implement other interface methods similarly
func (h *DoctorScheduleHandler) Update() echo.HandlerFunc {
	return func(c echo.Context) error {
		// 1. Get and validate ID
		idStr := c.Param("id")
		if idStr == "" {
			return c.JSON(http.StatusBadRequest, map[string]string{
				"error": "schedule id is required",
			})
		}

		id, err := strconv.ParseUint(idStr, 10, 64)
		if err != nil {
			return c.JSON(http.StatusBadRequest, map[string]string{
				"error": "invalid id format",
			})
		}

		// 2. Bind request data
		var updateData models.DoctorSchedule
		if err := c.Bind(&updateData); err != nil {
			logrus.WithError(err).Error("Failed to bind request")
			return c.JSON(http.StatusBadRequest, map[string]interface{}{
				"error":   "Invalid request payload",
				"details": err.Error(),
			})
		}

		// 3. Validate at least one field is being updated
		if updateData.DayOfWeek == "" && updateData.StartTime.IsZero() &&
			updateData.EndTime.IsZero() && !updateData.IsAvailable {
			return c.JSON(http.StatusBadRequest, map[string]string{
				"error": "no fields provided for update",
			})
		}

		// 4. Perform the update
		updatedSchedule, err := h.repo.Update(c.Request().Context(), uint(id), updateData)
		if err != nil {
			logrus.WithFields(logrus.Fields{
				"error": err.Error(),
				"id":    id,
			}).Error("Failed to update schedule")

			if strings.Contains(err.Error(), "not found") {
				return c.JSON(http.StatusNotFound, map[string]string{
					"error": "Schedule not found",
				})
			}
			return c.JSON(http.StatusInternalServerError, map[string]string{
				"error": "Failed to update schedule",
			})
		}

		return c.JSON(http.StatusOK, updatedSchedule)
	}
}

func (h *DoctorScheduleHandler) Delete() echo.HandlerFunc {
	return func(c echo.Context) error {
		return nil
	}
}

func (h *DoctorScheduleHandler) GetByDoctorID() echo.HandlerFunc {
	return func(c echo.Context) error {
		return nil
	}
}
