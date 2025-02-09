package main

import (
	"hi-doctor-be/config"
	"hi-doctor-be/controllers"
	"hi-doctor-be/routes"
	"hi-doctor-be/services"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	app := fiber.New()

	// Add CORS middleware
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "Origin, Content-Type, Accept, Authorization",
		AllowMethods: "GET,POST,PUT,DELETE",
	}))

	db, err := config.InitDB()
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	// Initialize services
	patientService := services.NewPatientService(db)
	doctorService := services.NewDoctorService(db)

	// Initialize controllers
	patientController := controllers.NewPatientController(patientService)
	doctorController := controllers.NewDoctorController(doctorService)

	// Setup routes
	routes.SetupRoutes(app, patientController, doctorController)

	log.Fatal(app.Listen(":3000"))
}
