package routes

import (
	"hi-doctor-be/controllers"
	"hi-doctor-be/middlewares"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App, patientController *controllers.PatientController, doctorController *controllers.DoctorController) {
	// API v1 group
	v1 := app.Group("/api/v1")

	// Patient routes
	patients := v1.Group("/patients")
	patients.Post("/verify", patientController.VerifyGoogleAuth)
	patients.Get("/:phone", patientController.GetProfile)
	patients.Put("/profile", patientController.UpdateProfile)

	// Doctor routes
	doctors := v1.Group("/doctors")
	doctors.Post("/register", doctorController.RegisterDoctor) // New endpoint for registering doctors
	doctors.Post("/login", doctorController.Login)

	// Protected doctor routes
	protectedDoctors := doctors.Group("/", middlewares.DoctorAuth())
	protectedDoctors.Get("/me", doctorController.GetProfile)
	protectedDoctors.Put("/me", doctorController.UpdateProfile)
}
