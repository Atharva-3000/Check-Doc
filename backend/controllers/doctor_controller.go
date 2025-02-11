package controllers

import (
	"hi-doctor-be/models"
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

	// For login, we need to decrypt the password that was encrypted on frontend
	decryptedPassword, err := utils.DecryptPassword(req.Password)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "Invalid password format",
		})
	}

	doctor, err := dc.doctorService.ValidateCredentials(req.Phone, decryptedPassword)
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

type RegisterDoctorRequest struct {
	Doctorname     string   `json:"doctorname"`
	Doctorphone    string   `json:"doctorphone"`
	Password       string   `json:"password"`
	Email          string   `json:"email"`
	Gender         string   `json:"gender"`
	Age            int      `json:"age"`
	Experience     int      `json:"experience"`
	Designation    string   `json:"designation"`
	Specialisation []string `json:"specialisation"`
	RoomNumber     string   `json:"room_number"`
}

func (dc *DoctorController) RegisterDoctor(c *fiber.Ctx) error {
	var req RegisterDoctorRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	// No need to decrypt during registration
	// The password will be hashed before storing
	doctor := &models.Doctor{
		Doctorname:     req.Doctorname,
		Doctorphone:    req.Doctorphone,
		Password:       req.Password, // Plain password will be hashed in service
		Email:          req.Email,
		Gender:         req.Gender,
		Age:            req.Age,
		Experience:     req.Experience,
		Designation:    req.Designation,
		Specialisation: req.Specialisation,
		RoomNumber:     req.RoomNumber,
	}

	if err := dc.doctorService.RegisterDoctor(doctor); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	doctor.Password = "" // Don't send password back
	return c.Status(201).JSON(fiber.Map{
		"message": "Doctor registered successfully",
		"doctor":  doctor,
	})
}
