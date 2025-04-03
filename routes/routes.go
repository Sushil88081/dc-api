package routes

import (
	"doctor-on-demand/handlers"

	"github.com/labstack/echo"
)

func Routes(e *echo.Echo, DoctorHandler *handlers.DoctorHandler) {
	e.POST("/doctor", DoctorHandler.Create())
	e.GET("/doctor/:id", DoctorHandler.GetById())
	e.PUT("/doctor/:id", DoctorHandler.Update())
	e.DELETE("/doctor/:id", DoctorHandler.Delete())
}

func PatientRoutes(e *echo.Echo, patientHandler *handlers.PatientHandler) {
	e.GET("/patient/:id", patientHandler.GetById())
	e.POST("/patient", patientHandler.Create())
	e.PUT("/patient/:id", patientHandler.Update())
	e.DELETE("/patient/:id", patientHandler.Delete())
}
