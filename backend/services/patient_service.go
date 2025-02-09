package services

import (
	"encoding/json"
	"errors"
	"hi-doctor-be/models"
	"net/http"
	"gorm.io/gorm"
)


type PatientService struct {
	db *gorm.DB
}


type GoogleUserInfo struct {
	ID            string `json:"id"`
	Email         string `json:"email"`
	VerifiedEmail bool   `json:"verified_email"`
	Name          string `json:"name"`
	Picture       string `json:"picture"`
}


func NewPatientService(db *gorm.DB) *PatientService {
	return &PatientService{
		db: db,
	}
}

func (s *PatientService) CreatePatient(patient *models.Patient) error {
	// Check if phone already exists
	var existingPatient models.Patient
	result := s.db.Where("patientphone = ?", patient.Patientphone).First(&existingPatient)
	if result.Error == nil {
		return errors.New("phone number already registered")
	}

	return s.db.Create(patient).Error
}

func (s *PatientService) VerifyGoogleToken(token string) (*models.Patient, error) {
	// Verify token with Google
	resp, err := http.Get("https://www.googleapis.com/oauth2/v2/userinfo?access_token=" + token)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, errors.New("failed to verify Google token")
	}

	var userInfo GoogleUserInfo
	if err := json.NewDecoder(resp.Body).Decode(&userInfo); err != nil {
		return nil, err
	}

	// Find or create patient
	var patient models.Patient
	err = s.db.Where("patientphone = ?", userInfo.Email).FirstOrCreate(&patient, models.Patient{
		Patientname:  userInfo.Name,
		Patientphone: userInfo.Email,
	}).Error

	if err != nil {
		return nil, err
	}

	return &patient, nil
}

func (s *PatientService) UpdateProfile(patient *models.Patient) (*models.Patient, error) {
	var existingPatient models.Patient
	if err := s.db.Where("patientphone = ?", patient.Patientphone).First(&existingPatient).Error; err != nil {
		return nil, errors.New("patient not found")
	}

	if err := s.db.Model(&existingPatient).Updates(patient).Error; err != nil {
		return nil, err
	}

	return &existingPatient, nil
}

func (s *PatientService) GetPatientByPhone(phone string) (*models.Patient, error) {
	var patient models.Patient
	if err := s.db.Where("patientphone = ?", phone).First(&patient).Error; err != nil {
		return nil, err
	}
	return &patient, nil
}
