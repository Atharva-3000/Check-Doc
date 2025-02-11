// models/doctor.go
package models

import (
	"github.com/lib/pq"
)

type Doctor struct {
	BaseModel
	Doctorname     string         `json:"doctorname"`
	Doctorphone    string         `gorm:"uniqueIndex" json:"doctorphone"`
	Password       string         `json:"-"` // Password is hidden from JSON
	Email          string         `json:"email" gorm:"uniqueIndex"`
	Gender         string         `json:"gender"`
	Age            int            `json:"age"`
	Experience     int            `json:"experience"`
	Designation    string         `json:"designation"`
	Specialisation pq.StringArray `json:"specialisation" gorm:"type:text[]"`
	RoomNumber     string         `json:"room_number"`
	IsPresent      bool           `json:"is_present" gorm:"default:false"`
	IsActive       bool           `json:"is_active" gorm:"default:true"`
	Schedules      []Schedule     `json:"schedules" gorm:"foreignKey:DoctorID"`
}
