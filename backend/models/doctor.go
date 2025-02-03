package models

type Doctor struct {
    BaseModel
    Doctorname     string     `json:"doctorname"`
    Doctorphone    string     `gorm:"primaryKey" json:"doctorphone"`
    Gender         string     `json:"gender"`
    Age            int        `json:"age"`
    Experience     int        `json:"experience"`
    Designation    string     `json:"designation"`
    Specialisation []string   `json:"specialisation"`
    RoomNumber     string     `json:"room_number"`
    IsPresent      bool       `json:"is_present" gorm:"default:false"`
    Schedules      []Schedule `json:"schedules" gorm:"foreignKey:DoctorID"`
}