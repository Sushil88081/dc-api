package main

import (
	"doctor-on-demand/handlers"
	"doctor-on-demand/models"
	"doctor-on-demand/routes"

	"github.com/labstack/echo"
)

func main() {
	// Initialize Echo instance
	e := echo.New()

	// Initialize logger
	// logger := logrus.New()

	// Initialize DoctorHandler with a proper models.DoctorList instance
	doctorList := models.DoctorList{} // Ensure this is properly initialized
	doctorHandler := handlers.NewDoctorHandler(doctorList)

	// Setup routes
	routes.DoctorRoutes(e, doctorHandler)

	// Start server
	e.Logger.Fatal(e.Start(":8080"))
}
