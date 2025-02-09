package models

type Patient struct {
	BaseModel
	Patientname   string `json:"patientname"`
	Patientphone  string `gorm:"primaryKey" json:"patientphone"`
	Gender        string `json:"gender"`
	Age           int    `json:"age"`
	Address       string `json:"address"`
	Healthissues  string `json:"healthissues"`
	Preferredtime string `json:"preferredtime"`
}
