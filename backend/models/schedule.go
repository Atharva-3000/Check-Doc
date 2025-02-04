// models/schedule.go
package models

import (
    "time"
)

type Schedule struct {
    BaseModel
    DoctorID    uint      `json:"doctor_id" gorm:"not null"`
    Doctor      Doctor    `json:"doctor" gorm:"constraint:OnDelete:CASCADE;"`
    StartTime   time.Time `json:"start_time" gorm:"not null"`
    EndTime     time.Time `json:"end_time" gorm:"not null"`
    Description string    `json:"description"`
    IsAvailable bool      `json:"is_available" gorm:"default:true"`
    Status      string    `json:"status" gorm:"type:schedule_status;default:'scheduled'"`
}