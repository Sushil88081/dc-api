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

	doctorRepo := repository.NewDoctorRepository(db)
	doctorHandler := handlers.NewDoctorHandler(doctorRepo)
	patientRepo := repository.NewPatientRepository(db)
	patientHandler := handlers.NewPatientHandler(patientRepo)
	routes.Routes(e, doctorHandler)
	routes.PatientRoutes(e, patientHandler)

	e.Logger.Fatal(e.Start(":8080"))
	defer db.Close()
}
