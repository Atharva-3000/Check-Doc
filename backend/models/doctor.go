package models


type Doctor struct{
	BaseModel
	Doctorname string `json:"doctorname"`
	Doctorphone string `gorm:"primaryKey" json:"doctorphone"`
}