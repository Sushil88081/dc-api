package initializers

import (
	"doctor-on-demand/config"
	"doctor-on-demand/handlers"
	repository "doctor-on-demand/repositories"
	"doctor-on-demand/routes"

	"github.com/labstack/echo"
	"gorm.io/gorm"
)

type App struct {
	DB       *gorm.DB
	Handlers struct {
		Doctor         *handlers.DoctorHandler
		Patient        *handlers.PatientHandler
		DoctorSchedule *handlers.DoctorScheduleHandler
	}
}

func Initializers() *App {
	// Initialize database connection
	db := config.ConnectDB()
	// 2. Initialize Repositories
	doctorRepo := repository.NewDoctorRepository(db)
	patientRepo := repository.NewPatientRepository(db)
	scheduleRepo := repository.NewDoctorScheduleRepository(db)
	app := &App{
		DB: db,
		Handlers: struct {
			Doctor         *handlers.DoctorHandler
			Patient        *handlers.PatientHandler
			DoctorSchedule *handlers.DoctorScheduleHandler
		}{
			Doctor:         handlers.NewDoctorHandler(doctorRepo),
			Patient:        handlers.NewPatientHandler(patientRepo),
			DoctorSchedule: handlers.NewDoctorScheduleHandler(scheduleRepo),
		},
	}
	return app
}
func (a *App) SetupRoutes(e *echo.Echo) {
	routes.Routes(e, a.Handlers.Doctor)
	routes.PatientRoutes(e, a.Handlers.Patient)
	routes.DoctorSchedule(e, a.Handlers.DoctorSchedule)
}
