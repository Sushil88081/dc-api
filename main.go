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
	doctorCollection := db.Collection("Doctors")
	doctorRepo := repository.NewDoctorRepository(doctorCollection)
	doctorHandler := handlers.NewDoctorHandler(doctorRepo)
	routes.Routes(e, doctorHandler)

	e.Logger.Fatal(e.Start(":8080"))
}
