package main

import (
	"doctor-on-demand/config"
	"doctor-on-demand/handlers"
	repository "doctor-on-demand/repositories"
	"doctor-on-demand/routes"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	// Initialize Echo instance
	e := echo.New()
	e.Use(middleware.CORS())
	db := config.ConnectDB()
	//doctor reppo initialization
	doctorRepo := repository.NewDoctorRepository(db)
	doctorHandler := handlers.NewDoctorHandler(doctorRepo)
	//patient reppo initialization
	patientRepo := repository.NewPatientRepository(db)
	patientHandler := handlers.NewPatientHandler(patientRepo)
	//apointments reppo initialization
	appointmentRepo := repository.NewAppointmentRepository(db)
	appointmentHandler := handlers.NewAppointementHandler(*appointmentRepo)
	routes.Routes(e, doctorHandler)
	routes.PatientRoutes(e, patientHandler)
	routes.AppointmentRoutes(e, appointmentHandler)

	// e.Logger.Fatal(e.Start(":8080"))
	e.Logger.Fatal(e.Start("0.0.0.0:8080"))
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"}, // You can restrict this later
		AllowMethods: []string{echo.GET, echo.POST, echo.PUT, echo.DELETE},
	}))
	defer db.Close()
}
