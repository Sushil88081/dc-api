package main

import (
	"doctor-on-demand/config"
	"doctor-on-demand/handlers"
	repository "doctor-on-demand/repositories"
	"doctor-on-demand/routes"
	"log"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	// Initialize Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"}, // Restrict this in production
		AllowMethods: []string{echo.GET, echo.POST, echo.PUT, echo.DELETE},
	}))

	// Connect to DB
	db := config.ConnectDB()

	// Close DB connection gracefully
	sqlDB, err := db.DB()
	if err != nil {
		log.Fatalf("Failed to get DB from GORM: %v", err)
	}
	defer sqlDB.Close()

	// Initialize repositories and handlers
	doctorRepo := repository.NewDoctorRepository(db)
	doctorHandler := handlers.NewDoctorHandler(doctorRepo)

	patientRepo := repository.NewPatientRepository(db)
	patientHandler := handlers.NewPatientHandler(patientRepo)

	// appointmentRepo := repository.NewAppointmentRepository(db)
	// appointmentHandler := handlers.NewAppointementHandler(*appointmentRepo)

	// scheduleRepo := repository.NewDoctorScheduleRepository(db)
	// scheduleHandler := handlers.NewDoctorScheduleHandler(scheduleRepo)

	// Define Routes
	routes.Routes(e, doctorHandler)
	routes.PatientRoutes(e, patientHandler)
	// routes.AppointmentRoutes(e, appointmentHandler)
	// routes.DoctorSchedule(e, scheduleHandler)

	// Start server
	e.Logger.Fatal(e.Start("0.0.0.0:8080"))
}
