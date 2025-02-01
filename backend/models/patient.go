package models


type Patient struct{
	BaseModel
	Patientname string `json:"patientname"`
	Patientphone string `gorm:"primaryKey" json:"patientphone"`
}