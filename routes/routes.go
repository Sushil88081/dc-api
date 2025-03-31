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
