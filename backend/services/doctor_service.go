package services

import (
    "errors"
    "hi-doctor-be/models"
    "golang.org/x/crypto/bcrypt"
    "gorm.io/gorm"
)

type DoctorService struct {
    db *gorm.DB
}

func NewDoctorService(db *gorm.DB) *DoctorService {
    return &DoctorService{
        db: db,
    }
}

func (s *DoctorService) ValidateCredentials(phone, password string) (*models.Doctor, error) {
    var doctor models.Doctor
    if err := s.db.Where("doctorphone = ?", phone).First(&doctor).Error; err != nil {
        if err == gorm.ErrRecordNotFound {
            return nil, errors.New("doctor not found")
        }
        return nil, err
    }

    if err := bcrypt.CompareHashAndPassword([]byte(doctor.Password), []byte(password)); err != nil {
        return nil, errors.New("invalid credentials")
    }

    return &doctor, nil
}


func (s *DoctorService) GetDoctorProfile(id uint) (*models.Doctor, error) {
    var doctor models.Doctor
    if err := s.db.First(&doctor, id).Error; err != nil {
        return nil, err
    }
    return &doctor, nil
}


func (s *DoctorService) UpdateDoctorProfile(id uint, updates map[string]interface{}) error {
    return s.db.Model(&models.Doctor{}).Where("id = ?", id).Updates(updates).Error
}