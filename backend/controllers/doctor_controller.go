package controllers

import (
	"hi-doctor-be/services"
	"hi-doctor-be/utils"

	"github.com/gofiber/fiber/v2"
)

type DoctorController struct {
	doctorService *services.DoctorService
}

func NewDoctorController(doctorService *services.DoctorService) *DoctorController {
	return &DoctorController{
		doctorService: doctorService,
	}
}

type LoginRequest struct {
	Phone    string `json:"phone"`
	Password string `json:"password"`
}

func (dc *DoctorController) Login(c *fiber.Ctx) error {
	var req LoginRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	doctor, err := dc.doctorService.ValidateCredentials(req.Phone, req.Password)
	if err != nil {
		return c.Status(401).JSON(fiber.Map{
			"error": "Invalid credentials",
		})
	}

	token, err := utils.GenerateDoctorToken(doctor.ID, doctor.Doctorphone)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": "Failed to generate token",
		})
	}

	return c.JSON(fiber.Map{
		"token":  token,
		"doctor": doctor,
	})
}

func (dc *DoctorController) GetProfile(c *fiber.Ctx) error {
	doctorID := c.Locals("doctorID").(uint)

	doctor, err := dc.doctorService.GetDoctorProfile(doctorID)
	if err != nil {
		return c.Status(404).JSON(fiber.Map{
			"error": "Doctor not found",
		})
	}

	return c.JSON(doctor)
}

func (dc *DoctorController) UpdateProfile(c *fiber.Ctx) error {
	doctorID := c.Locals("doctorID").(uint)

	updates := make(map[string]interface{})
	if err := c.BodyParser(&updates); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	// Remove sensitive fields from updates
	delete(updates, "password")
	delete(updates, "id")

	if err := dc.doctorService.UpdateDoctorProfile(doctorID, updates); err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": "Failed to update profile",
		})
	}

	return c.JSON(fiber.Map{
		"message": "Profile updated successfully",
	})
}
