package routes

import (
	"doctor-on-demand/handlers"

	"github.com/labstack/echo"
)

func Routes(e *echo.Echo, DoctorHandler *handlers.DoctorHandler) {
	e.POST("/doctors", DoctorHandler.Create())
	e.GET("/doctor/:id", DoctorHandler.GetById())
	// e.PUT("/users/:id", DoctorHandler.UpdateDoctor)
	// e.DELETE("/users/:id", DoctorHandler.DeleteDoctor)
}
