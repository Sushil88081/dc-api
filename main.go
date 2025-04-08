package main

import (
	"doctor-on-demand/config"
	"doctor-on-demand/handlers"
	repository "doctor-on-demand/repositories"
	"doctor-on-demand/routes"

	"github.com/labstack/echo"
)

func main() {
	// Initialize Echo instance
	e := echo.New()
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

	e.Logger.Fatal(e.Start(":8080"))
	defer db.Close()
}
