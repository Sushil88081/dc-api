package routes

import (
	"doctor-on-demand/handlers"

	"github.com/labstack/echo"
)

func DoctorRoutes(e *echo.Echo, DoctorHandler *handlers.DoctorHandler) {
	e.POST("/users", DoctorHandler.CreateDoctor)
	e.GET("/users/:id", DoctorHandler.GetDoctor)
	e.PUT("/users/:id", DoctorHandler.UpdateDoctor)
	e.DELETE("/users/:id", DoctorHandler.DeleteDoctor)
}
