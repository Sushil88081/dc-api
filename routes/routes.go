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
	e.GET("/doctors", DoctorHandler.GetAll())
	e.GET("/doctorCount", DoctorHandler.Count())

}

// func PatientRoutes(e *echo.Echo, patientHandler *handlers.PatientHandler) {
// 	e.GET("/patient/:id", patientHandler.GetById())
// 	e.POST("/patient", patientHandler.Create())
// 	e.PUT("/patient/:id", patientHandler.Update())
// 	e.DELETE("/patient/:id", patientHandler.Delete())
// }

// func AppointmentRoutes(e *echo.Echo, appointmentHandler *handlers.AppointementHandler) {
// 	//  e.GET("/appointment/:id", appointmentHandler.GetById())
// 	e.POST("/appointment", appointmentHandler.Create())
// 	// e.PUT("/appointment/:id", appointmentHandler.Update())
// 	// e.DELETE("/appointment/:id", appointmentHandler.Delete())
// }

// func DoctorSchedule(e *echo.Echo, ScheduleHandler *handlers.DoctorScheduleHandler) {
// 	e.POST("/schedule", ScheduleHandler.Create())
// }
