package controllers

import (
	"hi-doctor-be/models"
	"hi-doctor-be/services"

	"github.com/gofiber/fiber/v2"
)

type PatientController struct {
	patientService *services.PatientService
}

func NewPatientController(patientService *services.PatientService) *PatientController {
	return &PatientController{
		patientService: patientService,
	}
}

type GoogleAuthRequest struct {
	Token string `json:"token"`
}

func (pc *PatientController) VerifyGoogleAuth(c *fiber.Ctx) error {
	var req GoogleAuthRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	patient, err := pc.patientService.VerifyGoogleToken(req.Token)
	if err != nil {
		return c.Status(401).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"verified": true,
		"patient":  patient,
	})
}

func (pc *PatientController) UpdateProfile(c *fiber.Ctx) error {
	var patient models.Patient
	if err := c.BodyParser(&patient); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	updatedPatient, err := pc.patientService.UpdateProfile(&patient)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(updatedPatient)
}

func (pc *PatientController) GetProfile(c *fiber.Ctx) error {
	patientPhone := c.Params("phone")

	patient, err := pc.patientService.GetPatientByPhone(patientPhone)
	if err != nil {
		return c.Status(404).JSON(fiber.Map{
			"error": "Patient not found",
		})
	}

	return c.JSON(patient)
}
